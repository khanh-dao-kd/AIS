package middleware

import (
	"context"

	pb "ais_service/internal/generated/grpc/ais_api"
	"ais_service/internal/utils"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ValidationInterceptor interface {
	ValidateRequestMiddleware(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error)
}

type validationInterceptor struct {
	logger *zap.Logger
}

func NewValidationInterceptor(logger *zap.Logger) ValidationInterceptor {
	return &validationInterceptor{
		logger: logger,
	}
}

func (v validationInterceptor) ValidateRequestMiddleware(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	logger := utils.LoggerWithContext(ctx, v.logger)

	switch r := req.(type) {
	case *pb.PublishAisAccountRequest:
		// Validate account type
		if r.AccountType != pb.AccountType_CASA &&
			r.AccountType != pb.AccountType_GA &&
			r.AccountType != pb.AccountType_VAN {
			logger.Error("Invalid account_type", zap.String("value", r.AccountType.String()))
			return nil, status.Errorf(codes.InvalidArgument, "invalid account_type: %v", r.AccountType)
		}

		// Validate account status
		if r.AccountStatus != pb.Status_active &&
			r.AccountStatus != pb.Status_inactive &&
			r.AccountStatus != pb.Status_closed {
			logger.Error("Invalid account_status", zap.String("value", r.AccountStatus.String()))
			return nil, status.Errorf(codes.InvalidArgument, "invalid account_status: %v", r.AccountStatus)
		}
	case *pb.GetAccountStatusRequest:
		// No validation needed for this request
	default:
		// If the request is not recognized, just continue
	}

	// Proceed with the request
	return handler(ctx, req)
}
