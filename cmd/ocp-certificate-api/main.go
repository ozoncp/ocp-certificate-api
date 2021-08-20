package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	api "github.com/ozoncp/ocp-certificate-api/internal/api"
	"github.com/ozoncp/ocp-certificate-api/internal/config"
	"github.com/ozoncp/ocp-certificate-api/internal/repo"
	desc "github.com/ozoncp/ocp-certificate-api/pkg/ocp-certificate-api"
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
		config.GetConfigInstance().Database.Host,
		config.GetConfigInstance().Database.Port,
		config.GetConfigInstance().Database.User,
		config.GetConfigInstance().Database.Password,
		config.GetConfigInstance().Database.Name,
		config.GetConfigInstance().Database.SslMode)

	db, err := sqlx.Connect(config.GetConfigInstance().Database.Driver, dataSourceName)
	if err != nil {
		log.Error().Err(err).Msgf("failed to create connect to database")
	}

	if err = db.Ping(); err != nil {
		log.Error().Err(err).Msgf("failed to ping to database")
	}

	return db
}

// buildGrpc - building grpc server
func buildGrpc(db *sqlx.DB) (*grpc.Server, net.Listener) {
	listen, err := net.Listen("tcp", config.GetConfigInstance().Grpc.Address)
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	newRepo := repo.NewRepo(db)
	desc.RegisterOcpCertificateApiServer(grpcServer, api.NewOcpCertificateApi(newRepo))

	return grpcServer, listen
}

// buildRest - building rest server for send json
func buildRest(ctx context.Context) (*http.Server, error) {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := desc.RegisterOcpCertificateApiHandlerFromEndpoint(ctx, mux, config.GetConfigInstance().Grpc.Address, opts)
	if err != nil {
		return nil, err
	}

	srv := &http.Server{
		Addr:    config.GetConfigInstance().Json.Address,
		Handler: mux,
	}

	log.Info().Msg("Rest server started")
	return srv, nil
}

func main() {
	// Read config
	err := config.ReadConfigYML()
	if err != nil {
		log.Fatal().Msgf("failed read and init configuration file: %v", err)
		return
	}

	// Init channel and register notify
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	var grp errgroup.Group

	// Build DB and after work close
	db := buildDB()
	defer db.Close()

	// Build Rest
	restServer, err := buildRest(ctx)
	if err != nil {
		log.Fatal().Msgf("failed start rest server: %v", err)
	}

	// Build Grpc
	grpcServer, listen := buildGrpc(db)

	// Signal stopping servers
	go func() {
		oscall := <-c
		log.Info().Msgf("system call:%+v", oscall)

		if err = restServer.Shutdown(ctx); err != nil {
			log.Printf("shutdown error: %v\n", err)
		}
		grpcServer.GracefulStop()

		log.Info().Msg("servers stopped")
		cancel()
	}()

	// Rest server running
	grp.Go(func() error {
		return restServer.ListenAndServe()
	})

	// gRPC server running
	grp.Go(func() error {
		return grpcServer.Serve(listen)
	})

	// Handle sync group
	if err = grp.Wait(); err != http.ErrServerClosed {
		log.Fatal().Msgf("server shutdown failed: %v", err)
	}

	log.Info().Msg("servers correctly completed its work")
}
