package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jmoiron/sqlx"
	"github.com/ozoncp/ocp-certificate-api/internal/config"
	"github.com/ozoncp/ocp-certificate-api/internal/repo"
	"google.golang.org/grpc"
	"net"
	"net/http"

	_ "github.com/lib/pq"
	api "github.com/ozoncp/ocp-certificate-api/internal/api"
	desc "github.com/ozoncp/ocp-certificate-api/pkg/ocp-certificate-api"
	log "github.com/rs/zerolog/log"
)

var (
	cfg *config.Config
	err error
)

const (
	configYML = "config.yml"
)

func runGrpc() error {
	listen, err := net.Listen("tcp", cfg.Grpc.Address)
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}

	dataSourceName := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SslMode)

	db, err := sqlx.Connect(cfg.Database.Driver, dataSourceName)
	if err != nil {
		log.Error().Err(err).Msgf("failed to create connect to database")
		return err
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Error().Err(err).Msgf("failed to ping to database")
		return err
	}

	grpcServer := grpc.NewServer()
	newRepo := repo.NewRepo(db)
	desc.RegisterOcpCertificateApiServer(grpcServer, api.NewOcpCertificateApi(newRepo))

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatal().Msgf("failed to serve: %v", err)
	}

	return nil
}

func runJson() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := desc.RegisterOcpCertificateApiHandlerFromEndpoint(ctx, mux, cfg.Grpc.Address, opts)
	if err != nil {
		panic(err)
	}

	err = http.ListenAndServe(cfg.Json.Address, mux)
	if err != nil {
		panic(err)
	}
}

func main() {
	cfg, err = config.Read(configYML)
	if err != nil {
		log.Fatal().Msgf("failed read configuration file: %v", err)
		return
	}

	go runJson()

	if err := runGrpc(); err != nil {
		log.Fatal().Msgf("failed to start gRPC server: %v", err)
	}

	fmt.Println("run success")
}
