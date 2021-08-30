package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	api "github.com/ozoncp/ocp-certificate-api/internal/api"
	"github.com/ozoncp/ocp-certificate-api/internal/broker"
	cfg "github.com/ozoncp/ocp-certificate-api/internal/config"
	"github.com/ozoncp/ocp-certificate-api/internal/metrics"
	"github.com/ozoncp/ocp-certificate-api/internal/repo"
	"github.com/ozoncp/ocp-certificate-api/internal/tracer"
	desc "github.com/ozoncp/ocp-certificate-api/pkg/ocp-certificate-api"
	"github.com/pressly/goose/v3"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	_ "golang.org/x/tools/godoc/vfs"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"
)

// initDB - init db (postgres)
func initDB() *sqlx.DB {
	dataSourceName := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		cfg.GetConfigInstance().Database.Host,
		cfg.GetConfigInstance().Database.Port,
		cfg.GetConfigInstance().Database.User,
		cfg.GetConfigInstance().Database.Password,
		cfg.GetConfigInstance().Database.Name,
		cfg.GetConfigInstance().Database.SslMode)

	db, err := sqlx.Connect(cfg.GetConfigInstance().Database.Driver, dataSourceName)
	if err != nil {
		log.Error().Err(err).Msgf("failed to create connect to database")
	}

	if err = db.Ping(); err != nil {
		log.Error().Err(err).Msgf("failed to ping to database")
	}

	return db
}

// grpcServer - grpc server
func grpcServer(r repo.Repo, m metrics.Metrics, p broker.Producer, c broker.Consumer) (*grpc.Server, net.Listener) {
	listen, err := net.Listen("tcp", cfg.GetConfigInstance().Grpc.Address)
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}

	gSrv := grpc.NewServer()
	desc.RegisterOcpCertificateApiServer(gSrv, api.NewOcpCertificateAPI(r, m, p, c))

	log.Info().Msg("gRPC server started")
	return gSrv, listen
}

// restServer - rest server for send json
func restServer(ctx context.Context) (*http.Server, error) {
	isReady := &atomic.Value{}
	isReady.Store(false)
	go func() {
		log.Printf("Ready probe is negative by default...")
		time.Sleep(10 * time.Second)
		isReady.Store(true)
		log.Printf("Ready probe is positive.")
	}()

	serMux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := desc.RegisterOcpCertificateApiHandlerFromEndpoint(ctx, serMux, cfg.GetConfigInstance().Grpc.Address, opts)
	if err != nil {
		return nil, err
	}

	mux := http.NewServeMux()
	mux.Handle("/", serMux)
	mux.Handle("/live", http.HandlerFunc(live))
	mux.Handle("/health", http.HandlerFunc(health))
	log.Info().Msg("Liveness started")
	mux.Handle("/ready", ready(isReady))
	log.Info().Msg("Readiness started")

	srv := &http.Server{
		Addr:    cfg.GetConfigInstance().Rest.Address,
		Handler: mux,
	}

	log.Info().Msg("Rest server started")
	return srv, nil
}

// live is a simple HTTP handler function which writes a response.
func live(w http.ResponseWriter, _ *http.Request) {
	body, err := json.Marshal("This is service live!")
	if err != nil {
		log.Error().Err(err).Msgf("Could not encode info data: %v", err)
		http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(body)
	if err != nil {
		log.Error().Err(err).Msgf("Could not write body: %v", err)
		return
	}
}

// health is a liveness probe.
func health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// ready is a readiness probe.
func ready(isReady *atomic.Value) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		if isReady == nil || !isReady.Load().(bool) {
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

// metricsServer - metrics server
func metricsServer() *http.Server {
	mux := http.NewServeMux()
	mux.Handle(cfg.GetConfigInstance().Prometheus.URI, promhttp.Handler())

	srv := &http.Server{
		Addr:    cfg.GetConfigInstance().Prometheus.Port,
		Handler: mux,
	}

	log.Info().Msg("Metrics server started")
	return srv
}

// kafka - message broker
func kafka(r repo.Repo, m metrics.Metrics) (broker.Producer, broker.Consumer) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	syncProducer, err := sarama.NewSyncProducer(cfg.GetConfigInstance().Kafka.Brokers, config)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create Sarama new sync broker")
	}

	prod := broker.NewProducer(syncProducer)
	cons := broker.NewConsumer(r, m)

	log.Info().Msg("Kafka message broker started and init")
	return prod, cons
}

func migrationsCli(db *sqlx.DB) error {
	runCommand := flag.String("migrate", "", "specify the command: [up] or [down]")
	flag.Parse()
	if *runCommand == "" {
		return nil
	}

	rootDir, err := os.Getwd()
	if err != nil {
		return err
	}

	switch *runCommand {
	case "up":
		if err = goose.Up(db.DB, rootDir+"/migrations"); err != nil {
			log.Fatal().Err(err).Msg("failed to up migrate")
		}
	case "down":
		if err = goose.Down(db.DB, rootDir+"/migrations"); err != nil {
			log.Fatal().Err(err).Msg("failed to down migrate")
		}
	default:
		log.Warn().Msgf("available commands [up] or [down], you specified %v", *runCommand)
	}

	return nil
}

func main() {
	// Read config
	err := cfg.ReadConfigYML()
	if err != nil {
		log.Fatal().Msgf("failed read and init configuration file: %v", err)
		return
	}

	// Init channel and register notify
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	var grp errgroup.Group

	// Init tracer
	closer := tracer.InitTracer("ocp-certificate-api")

	// Init DB and after work close
	db := initDB()
	newRepo := repo.NewRepo(db)

	// cli command for up/down migrate
	err = migrationsCli(db)
	if err != nil {
		log.Fatal().Msgf("failed migrate command: %v", err)
	}

	// Metrics server
	m := metrics.NewMetrics()
	mSrv := metricsServer()

	// kafka message broker
	prod, cons := kafka(newRepo, m)

	// Rest server
	rSrv, err := restServer(ctx)
	if err != nil {
		log.Fatal().Msgf("failed start rest server: %v", err)
	}

	// Grpc server
	gSrv, listen := grpcServer(newRepo, m, prod, cons)

	// Rest server running
	grp.Go(func() error {
		return rSrv.ListenAndServe()
	})

	// gRPC server running
	grp.Go(func() error {
		return gSrv.Serve(listen)
	})

	// Metrics register and server running
	grp.Go(func() error {
		return mSrv.ListenAndServe()
	})

	// Signal stopping servers
	osSignal := <-c
	log.Info().Msgf("system syscall:%+v", osSignal)

	if err = mSrv.Shutdown(ctx); err != nil {
		log.Printf("shutdown error: %v\n", err)
	}

	if err = rSrv.Shutdown(ctx); err != nil {
		log.Printf("shutdown error: %v\n", err)
	}

	gSrv.GracefulStop()
	log.Info().Msg("servers stopped")

	closer.Close()
	log.Info().Msg("tracer stopped")

	db.Close()
	log.Info().Msg("db stopped")

	prod.Close()
	log.Info().Msg("broker kafka stopped")

	cons.Close()
	log.Info().Msg("consumer kafka stopped")

	cancel()

	// Handle sync group
	if err = grp.Wait(); err != http.ErrServerClosed {
		log.Fatal().Msgf("server shutdown failed: %v", err)
	}

	log.Info().Msg("services correctly completed its work")
}
