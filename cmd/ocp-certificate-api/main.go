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
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func runGrpc() error {
	listen, err := net.Listen("tcp", config.GetConfigInstance().Grpc.Address)
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}

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
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		oscall := <-c
		log.Info().Msgf("system call:%+v", oscall)
		cancel()
	}()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := desc.RegisterOcpCertificateApiHandlerFromEndpoint(ctx, mux, config.GetConfigInstance().Grpc.Address, opts)
	if err != nil {
		panic(err)
	}

	srv := &http.Server{
		Addr:    config.GetConfigInstance().Json.Address,
		Handler: mux,
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Msgf("listen:%+s\n", err)
		}
	}()

	log.Info().Msg("server started")
	<-ctx.Done()
	log.Info().Msg("server stopped")

	if err = srv.Shutdown(ctx); err != nil {
		log.Fatal().Msgf("server shutdown failed:%+s", err)
	}

	log.Fatal().Msg("server correctly completed its work")
	if err == http.ErrServerClosed {
		err = nil
	}
}

func main() {
	err := config.ReadConfigYML()
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
