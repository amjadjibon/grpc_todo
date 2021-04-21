package server

import (
	"fmt"
	"github.com/amjadjibon/grpc_todo/pkg/api/customer"
	"github.com/amjadjibon/grpc_todo/pkg/conf"
	"github.com/amjadjibon/grpc_todo/pkg/logger"
	"github.com/amjadjibon/grpc_todo/pkg/postgres"
	"github.com/amjadjibon/grpc_todo/pkg/protocol/rest"
	"github.com/go-pg/pg/v10"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	customer.UnimplementedCustomerServiceServer
	DB *pg.DB
}

func RunServer()  {
	logger.InitLogger()
	log.Info().
		Str("host", conf.Config.Host).
		Int("port", conf.Config.GrpcPort).
		Msg("Initializing gRPC Server")

	address := fmt.Sprintf("%s:%d", conf.Config.Host, conf.Config.GrpcPort)

	db := postgres.Connect()
	defer db.Close()

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal().Err(err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	customer.RegisterCustomerServiceServer(grpcServer, &customer.Service{DB: db})

	// handle signal
	idleChan := make(chan struct{})
	go func() {
		signChan := make(chan os.Signal, 1)
		signal.Notify(signChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
		sig := <-signChan

		log.Info().Str("signal", sig.String()).Msg("shutdown signal received")
		log.Info().Msg("preparing for shutdown")

		//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		//defer cancel()
		grpcServer.GracefulStop()
		// Actual shutdown trigger.
		close(idleChan)
	}()

	go func() {
		log.Info().
			Msg("server started...")
		err = grpcServer.Serve(lis)
		if err != nil {
			log.Fatal().Msg(err.Error())
			os.Exit(1)
		}
	}()

	go func() {
		rest.RunServer()
	}()

	// Blocking until the shutdown is complete
	<-idleChan
	log.Info().Msg("shutdown complete")
}
