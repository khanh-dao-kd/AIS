package http

import (
	"ais_service/internal/configs"
	"ais_service/internal/generated/grpc/ais_api"
	"context"
	"net/http"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type Server interface {
	Start(ctx context.Context) error
}

type server struct {
	httpConfig configs.HTTP
	grpcConfig configs.GRPC
}

func NewServer(httpConfig configs.HTTP, grpcConfig configs.GRPC) Server {
	return &server{
		httpConfig: httpConfig,
		grpcConfig: grpcConfig,
	}
}

func (s server) getGrpcGatewayHandler(ctx context.Context) (http.Handler, error) {
	grpcMux := runtime.NewServeMux()
	err := ais_api.RegisterAISServiceHandlerFromEndpoint(
		ctx,
		grpcMux,
		s.grpcConfig.Address,
		[]grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		})
	if err != nil {
		return nil, err
	}
	return grpcMux, nil
}

func (s server) Start(ctx context.Context) error {
	grpcGatewayHandler, err := s.getGrpcGatewayHandler(ctx)
	if err != nil {
		return err
	}
	httpServer := http.Server{
		Addr:              s.httpConfig.Address,
		ReadHeaderTimeout: time.Minute,
		Handler:           grpcGatewayHandler,
	}

	return httpServer.ListenAndServe()
}
