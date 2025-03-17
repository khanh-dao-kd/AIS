package producer

import (
	"ais_service/internal/generated/grpc/ais_api"
	"ais_service/internal/utils"
	"context"
	"encoding/json"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	AISAccountTopic = "ais_account_topic"
)

type AccountEvent struct {
	Account_id     uint64
	Account_name   string
	Account_type   ais_api.AccountType
	Account_status ais_api.Status
}

type AccountProducer interface {
	Produce(ctx context.Context, event AccountEvent) error
}

type accountProducer struct {
	client Client
	logger *zap.Logger
}

func NewAccountProducer(client Client, logger *zap.Logger) AccountProducer {
	return &accountProducer{
		client: client,
		logger: logger,
	}
}

func (a accountProducer) Produce(ctx context.Context, event AccountEvent) error {
	logger := utils.LoggerWithContext(ctx, a.logger)

	eventBytes, err := json.Marshal(event)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to marshal event")
		return status.Error(codes.Internal, "failed to marshal event")
	}

	err = a.client.Produce(ctx, AISAccountTopic, eventBytes)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to produce event")
		return status.Error(codes.Internal, "failed to produce event")
	}
	return nil
}
