package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	api "github.com/ozoncp/ocp-certificate-api/internal/api"
	cfg "github.com/ozoncp/ocp-certificate-api/internal/config"
	"github.com/ozoncp/ocp-certificate-api/internal/metrics"
	"github.com/ozoncp/ocp-certificate-api/internal/producer"
	"github.com/ozoncp/ocp-certificate-api/internal/repo"
	"github.com/ozoncp/ocp-certificate-api/internal/tracer"
	desc "github.com/ozoncp/ocp-certificate-api/pkg/ocp-certificate-api"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// buildDB - init db (postgres)
func buildDB() *sqlx.DB {
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
func grpcServer(db *sqlx.DB, prod producer.Producer) (*grpc.Server, net.Listener) {
	listen, err := net.Listen("tcp", cfg.GetConfigInstance().Grpc.Address)
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	newRepo := repo.NewRepo(db)
	desc.RegisterOcpCertificateApiServer(grpcServer, api.NewOcpCertificateApi(newRepo, prod, cfg.GetConfigInstance().BatchSize))

	log.Info().Msg("gRPC server started")
	return grpcServer, listen
}

// restServer - rest server for send json
func restServer(ctx context.Context) (*http.Server, error) {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := desc.RegisterOcpCertificateApiHandlerFromEndpoint(ctx, mux, cfg.GetConfigInstance().Grpc.Address, opts)
	if err != nil {
		return nil, err
	}

	srv := &http.Server{
		Addr:    cfg.GetConfigInstance().Rest.Address,
		Handler: mux,
	}

	log.Info().Msg("Rest server started")
	return srv, nil
}

// metricsServer - metrics server
func metricsServer() (*http.Server, error) {
	mux := http.NewServeMux()
	mux.Handle(cfg.GetConfigInstance().Prometheus.Uri, promhttp.Handler())

	srv := &http.Server{
		Addr:    cfg.GetConfigInstance().Prometheus.Port,
		Handler: mux,
	}

	log.Info().Msg("Metrics server started")
	return srv, nil
}

// kafka - message broker
func kafka(ctx context.Context) producer.Producer {
	prod := producer.NewProducer()
	prod.Init(ctx)

	log.Info().Msg("Kafka message broker started and init")
	return prod
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
	defer closer.Close()

	// Build DB and after work close
	db := buildDB()
	defer db.Close()

	// Metrics server
	metricsServer, err := metricsServer()
	if err != nil {
		log.Fatal().Msgf("failed start metrics server: %v", err)
	}

	// Rest server
	restServer, err := restServer(ctx)
	if err != nil {
		log.Fatal().Msgf("failed start rest server: %v", err)
	}

	// kafka message broker
	prod := kafka(ctx)

	// Grpc server
	grpcServer, listen := grpcServer(db, prod)

	// Rest server running
	grp.Go(func() error {
		return restServer.ListenAndServe()
	})

	// gRPC server running
	grp.Go(func() error {
		return grpcServer.Serve(listen)
	})

	// Metrics register and server running
	grp.Go(func() error {
		metrics.RegisterMetrics()
		return metricsServer.ListenAndServe()
	})

	// Signal stopping servers
	osSignal := <-c
	log.Info().Msgf("system syscall:%+v", osSignal)

	if err = metricsServer.Shutdown(ctx); err != nil {
		log.Printf("shutdown error: %v\n", err)
	}

	if err = restServer.Shutdown(ctx); err != nil {
		log.Printf("shutdown error: %v\n", err)
	}

	grpcServer.GracefulStop()

	log.Info().Msg("servers stopped")
	cancel()

	// Handle sync group
	if err = grp.Wait(); err != http.ErrServerClosed {
		log.Fatal().Msgf("server shutdown failed: %v", err)
	}

	log.Info().Msg("servers correctly completed its work")
}
