package grpc

import (
	"ais_service/internal/configs"
	"ais_service/internal/generated/grpc/ais_api"
	"ais_service/internal/handler/grpc/middleware"
	"context"
	"net"

	"google.golang.org/grpc"
)

type Server interface {
	Start(ctx context.Context) error
}

type server struct {
	handler         ais_api.AISServiceServer
	grpcConfig      configs.GRPC
	authInterceptor middleware.AuthInterceptor
}

func NewServer(handler ais_api.AISServiceServer, grpcConfig configs.GRPC, authInterceptor middleware.AuthInterceptor) Server {
	return &server{
		handler:         handler,
		grpcConfig:      grpcConfig,
		authInterceptor: authInterceptor,
	}
}

func (s server) Start(ctx context.Context) error {
	listener, err := net.Listen("tcp", s.grpcConfig.Address)
	if err != nil {
		return err
	}
	defer listener.Close()

	server := grpc.NewServer(
		grpc.UnaryInterceptor(s.authInterceptor.JWTAuthMiddleware),
	)
	ais_api.RegisterAISServiceServer(server, s.handler)

	return server.Serve(listener)
}
