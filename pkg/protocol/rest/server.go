package rest

import (
	"context"
	"fmt"
	"github.com/amjadjibon/grpc_todo/pkg/api/customer"
	"github.com/amjadjibon/grpc_todo/pkg/conf"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// RunServer runs HTTP/REST gateway
func RunServer()  {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	endPoint := fmt.Sprintf(":%d", conf.Config.GrpcPort)
	if err := customer.RegisterCustomerServiceHandlerFromEndpoint(ctx, mux, endPoint, opts); err != nil {
		log.Fatal().Err(err)
	}

	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", conf.Config.HttpPort),
		Handler: mux,
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
		}

		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal().Err(err)
	}
}
