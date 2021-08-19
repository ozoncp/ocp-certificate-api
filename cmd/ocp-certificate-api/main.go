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

func runGrpc() error {
	listen, err := net.Listen("tcp", config.Get.Grpc.Address)
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}

	dataSourceName := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		config.Get.Database.Host,
		config.Get.Database.Port,
		config.Get.Database.User,
		config.Get.Database.Password,
		config.Get.Database.Name,
		config.Get.Database.SslMode)

	db, err := sqlx.Connect(config.Get.Database.Driver, dataSourceName)
	if err != nil {
		log.Error().Err(err).Msgf("failed to create connect to database")
		return err
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Error().Err(err).Msgf("failed to ping to database")
		return err
	}

	grpcServer := grpc.NewServer()
	newRepo := repo.NewRepo(db)
	desc.RegisterOcpCertificateApiServer(grpcServer, api.NewOcpCertificateApi(newRepo))

	if err = grpcServer.Serve(listen); err != nil {
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

	err := desc.RegisterOcpCertificateApiHandlerFromEndpoint(ctx, mux, config.Get.Grpc.Address, opts)
	if err != nil {
		panic(err)
	}

	err = http.ListenAndServe(config.Get.Json.Address, mux)
	if err != nil {
		panic(err)
	}
}

func main() {
	err := config.Init()
	if err != nil {
		log.Fatal().Msgf("failed read and init configuration file: %v", err)
		return
	}

	go runJson()

	if err = runGrpc(); err != nil {
		log.Fatal().Msgf("failed to start gRPC server: %v", err)
	}

	log.Info().Msgf("run success")
}
