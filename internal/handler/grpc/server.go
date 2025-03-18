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
	handler             ais_api.AISServiceServer
	grpcConfig          configs.GRPC
	authInterceptor     middleware.AuthInterceptor
	validateInterceptor middleware.ValidationInterceptor
}

func NewServer(handler ais_api.AISServiceServer, grpcConfig configs.GRPC, authInterceptor middleware.AuthInterceptor, validateInterceptor middleware.ValidationInterceptor) Server {
	return &server{
		handler:             handler,
		grpcConfig:          grpcConfig,
		authInterceptor:     authInterceptor,
		validateInterceptor: validateInterceptor,
	}
}

func (s server) Start(ctx context.Context) error {
	listener, err := net.Listen("tcp", s.grpcConfig.Address)
	if err != nil {
		return err
	}
	defer listener.Close()

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			s.authInterceptor.JWTAuthMiddleware,
			s.validateInterceptor.ValidateRequestMiddleware,
		),
	)
	ais_api.RegisterAISServiceServer(server, s.handler)

	return server.Serve(listener)
}
