package producer

import (
	"ais_service/internal/generated/grpc/ais_api"
	"context"
	"encoding/json"

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
}

func NewAccountProducer(client Client) AccountProducer {
	return &accountProducer{
		client: client,
	}
}

func (a accountProducer) Produce(ctx context.Context, event AccountEvent) error {
	eventBytes, err := json.Marshal(event)
	if err != nil {
		return status.Error(codes.Internal, "failed to marshal download task created event")
	}

	err = a.client.Produce(ctx, AISAccountTopic, eventBytes)
	if err != nil {
		return status.Error(codes.Internal, "failed to marshal download task created event")
	}
	return nil
}
