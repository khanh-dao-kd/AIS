package grpc

import (
	"ais_service/internal/generated/grpc/ais_api"
	"ais_service/internal/logic"
	"context"
)

type grpcHandler struct {
	ais_api.UnimplementedAISServiceServer
	accountLogic   logic.AccountLogic
	publisherLogic logic.PublisherLogic
}

func NewGrpcHandler(accountLogic logic.AccountLogic, publisherLogic logic.PublisherLogic) ais_api.AISServiceServer {
	return &grpcHandler{
		accountLogic:   accountLogic,
		publisherLogic: publisherLogic,
	}
}

func (g grpcHandler) GetAisAccountByID(ctx context.Context, request *ais_api.GetAccountStatusRequest) (*ais_api.GetAccountStatusResponse, error) {
	params := logic.GetAisAccountByIDParams{
		Account_id: request.GetAccountId(),
	}
	output, err := g.accountLogic.GetAisAccountByID(ctx, params)
	if err != nil {
		return nil, err
	}
	return &ais_api.GetAccountStatusResponse{
		AccountId:     output.Account_id,
		AccountName:   output.Account_name,
		AccountType:   output.Account_type,
		AccountStatus: output.Account_status,
	}, nil
}

func (g grpcHandler) PublishAisAccount(ctx context.Context, request *ais_api.PublishAisAccountRequest) (*ais_api.PublishAisAccountResponse, error) {
	params := logic.PublishAisAccountParams{
		Account_id:     request.AccountId,
		Account_name:   request.AccountName,
		Account_type:   request.AccountType,
		Account_status: request.AccountStatus,
	}
	_, err := g.publisherLogic.PublishAisAccount(ctx, params)
	if err != nil {
		return nil, err
	}
	return &ais_api.PublishAisAccountResponse{}, nil
}
