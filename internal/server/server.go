package server

import (
	"ais_service/internal/handler/consumer"
	"ais_service/internal/handler/grpc"
	"ais_service/internal/handler/http"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type StandaloneServer struct {
	grpcServer grpc.Server
	httpServer http.Server
	consumer   consumer.ConsumerServer
}

func NewStandaloneServer(grpcServer grpc.Server, httpServer http.Server, consumer consumer.ConsumerServer) *StandaloneServer {
	return &StandaloneServer{
		grpcServer: grpcServer,
		httpServer: httpServer,
		consumer:   consumer,
	}
}

func (s StandaloneServer) Start() error {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		s.grpcServer.Start(ctx)
	}()

	go func() {
		s.httpServer.Start(ctx)
	}()

	go func() {
		s.consumer.Start(context.Background())
	}()
	// Listen for OS termination signals (Ctrl+C, SIGTERM)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Block until a termination signal is received
	sig := <-sigChan
	fmt.Printf("\nReceived signal: %v. Shutting down...\n", sig)

	// Cancel context to notify goroutines to stop
	cancel()

	return nil
}
