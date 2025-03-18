package middleware

import (
	"ais_service/internal/utils"
	"context"
	"math/rand"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthInterceptor interface {
	JWTAuthMiddleware(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error)
}

type authInterceptor struct {
	logger *zap.Logger
}

func NewAuthInterceptor(logger *zap.Logger) AuthInterceptor {
	rand.Seed(time.Now().UnixNano())
	return &authInterceptor{
		logger: logger,
	}
}

func (a authInterceptor) JWTAuthMiddleware(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	logger := utils.LoggerWithContext(ctx, a.logger)

	if rand.Intn(100) < 10 {
		logger.Error("Unanthentication request")
		return nil, status.Error(codes.Unauthenticated, "random authentication failure")
	}
	return handler(ctx, req)
}
