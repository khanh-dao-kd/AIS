package grpc_test

import (
	"ais_service/internal/generated/grpc/ais_api"
	"ais_service/internal/handler/grpc"
	"ais_service/internal/logic"
	"ais_service/mocks"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test GetAisAccountByID - Success
func TestGetAisAccountByID_Success(t *testing.T) {
	mockAccountLogic := new(mocks.AccountLogic)
	handler := grpc.NewGrpcHandler(mockAccountLogic, nil) // Pass mock dependencies

	ctx := context.Background()
	req := &ais_api.GetAccountStatusRequest{AccountId: 1}

	mockAccountLogic.On("GetAisAccountByID", ctx, logic.GetAisAccountByIDParams{Account_id: 1}).
		Return(logic.GetAisAccountByIDOutput{
			Account_id:     1,
			Account_name:   "Test Account",
			Account_type:   1,
			Account_status: 0,
		}, nil)

	resp, err := handler.GetAisAccountByID(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, req.AccountId, resp.AccountId)
	assert.Equal(t, "Test Account", resp.AccountName)
	mockAccountLogic.AssertExpectations(t)
}

// Test GetAisAccountByID - Account Not Found
func TestGetAisAccountByID_NotFound(t *testing.T) {
	mockAccountLogic := new(mocks.AccountLogic)
	handler := grpc.NewGrpcHandler(mockAccountLogic, nil)

	ctx := context.Background()
	req := &ais_api.GetAccountStatusRequest{AccountId: 2}

	mockAccountLogic.On("GetAisAccountByID", ctx, logic.GetAisAccountByIDParams{Account_id: 2}).
		Return(logic.GetAisAccountByIDOutput{}, errors.New("account not found"))

	resp, err := handler.GetAisAccountByID(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	mockAccountLogic.AssertExpectations(t)
}

// Test PublishAisAccount - Success
func TestPublishAisAccount_Success(t *testing.T) {
	mockPublisherLogic := new(mocks.PublisherLogic)
	handler := grpc.NewGrpcHandler(nil, mockPublisherLogic)

	ctx := context.Background()
	req := &ais_api.PublishAisAccountRequest{
		AccountId:     3,
		AccountName:   "Publish Test",
		AccountType:   2,
		AccountStatus: 1,
	}

	mockPublisherLogic.On("PublishAisAccount", ctx, logic.PublishAisAccountParams{
		Account_id:     req.AccountId,
		Account_name:   req.AccountName,
		Account_type:   req.AccountType,
		Account_status: req.AccountStatus,
	}).Return(logic.PublishAisAccountOutput{}, nil)

	resp, err := handler.PublishAisAccount(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	mockPublisherLogic.AssertExpectations(t)
}

// Test PublishAisAccount - Failure
func TestPublishAisAccount_Failure(t *testing.T) {
	mockPublisherLogic := new(mocks.PublisherLogic)
	handler := grpc.NewGrpcHandler(nil, mockPublisherLogic)

	ctx := context.Background()
	req := &ais_api.PublishAisAccountRequest{
		AccountId:     4,
		AccountName:   "Fail Test",
		AccountType:   2,
		AccountStatus: 1,
	}

	mockPublisherLogic.On("PublishAisAccount", ctx, logic.PublishAisAccountParams{
		Account_id:     req.AccountId,
		Account_name:   req.AccountName,
		Account_type:   req.AccountType,
		Account_status: req.AccountStatus,
	}).Return(logic.PublishAisAccountOutput{}, errors.New("publish failed"))

	resp, err := handler.PublishAisAccount(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	mockPublisherLogic.AssertExpectations(t)
}
