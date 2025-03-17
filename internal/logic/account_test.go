package logic_test

import (
	"ais_service/internal/dataaccess/database"
	"ais_service/internal/logic"
	"ais_service/mocks"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestAddAisMessage_AccountExists(t *testing.T) {
	mockAccessor := new(mocks.AisAccountDataAccessor)
	logger := zap.NewNop()
	logicService := logic.NewAccountLogic(mockAccessor, logger)
	ctx := context.Background()
	params := logic.AddAisMessaggParams{Account_id: 1, Account_name: "Test", Account_type: 0, Account_status: 1}

	mockAccessor.On("GetAisAccountByIDForUpdate", ctx, params.Account_id).Return(database.AisAccount{}, nil)
	mockAccessor.On("UpdateAisAccount", ctx, mock.Anything).Return(logic.UpdateAccountStatusOutput{Account_id: 1}, nil)

	output, err := logicService.AddAisMessage(ctx, params)

	assert.NoError(t, err)
	assert.Equal(t, params.Account_id, output.Account_id)
	mockAccessor.AssertExpectations(t)
}

func TestAddAisMessage_AccountNotFound(t *testing.T) {
	mockAccessor := new(mocks.AisAccountDataAccessor)
	logger := zap.NewNop()
	logicService := logic.NewAccountLogic(mockAccessor, logger)
	ctx := context.Background()
	params := logic.AddAisMessaggParams{Account_id: 2, Account_name: "Test2", Account_type: 1, Account_status: 0}

	mockAccessor.On("GetAisAccountByIDForUpdate", ctx, params.Account_id).Return(database.AisAccount{}, status.Error(codes.NotFound, "account not found"))
	mockAccessor.On("CreateAisAccount", ctx, mock.Anything).Return(uint64(2), nil)

	output, err := logicService.AddAisMessage(ctx, params)

	assert.NoError(t, err)
	assert.Equal(t, params.Account_id, output.Account_id)
	mockAccessor.AssertExpectations(t)
}

func TestCreateAisAccount(t *testing.T) {
	mockAccessor := new(mocks.AisAccountDataAccessor)
	logicService := logic.NewAccountLogic(mockAccessor, &zap.Logger{})
	ctx := context.Background()
	params := logic.CreateAisAccountParams{Account_id: 3, Account_name: "Test3", Account_type: 1, Account_status: 1}

	mockAccessor.On("CreateAisAccount", ctx, mock.Anything).Return(uint64(3), nil)

	output, err := logicService.CreateAisAccount(ctx, params)

	assert.NoError(t, err)
	assert.Equal(t, params.Account_id, output.Account_id)
	mockAccessor.AssertExpectations(t)
}

func TestGetAisAccountByID(t *testing.T) {
	mockAccessor := new(mocks.AisAccountDataAccessor)
	logicService := logic.NewAccountLogic(mockAccessor, &zap.Logger{})
	ctx := context.Background()
	params := logic.GetAisAccountByIDParams{Account_id: 4}
	expectedAccount := database.AisAccount{Account_id: 4, Account_name: "Test4", Account_type: 2, Account_status: 1}

	mockAccessor.On("GetAisAccountByID", ctx, params.Account_id).Return(expectedAccount, nil)

	output, err := logicService.GetAisAccountByID(ctx, params)

	assert.NoError(t, err)
	assert.Equal(t, expectedAccount.Account_id, output.Account_id)
	mockAccessor.AssertExpectations(t)
}

func TestUpdateAisAccount(t *testing.T) {
	mockAccessor := new(mocks.AisAccountDataAccessor)
	logicService := logic.NewAccountLogic(mockAccessor, &zap.Logger{})
	ctx := context.Background()
	params := logic.UpdateAccountStatusParams{Account_id: 5, Account_name: "Test5", Account_type: 1, Account_status: 0}

	mockAccessor.On("UpdateAisAccount", ctx, mock.Anything).Return(nil)

	output, err := logicService.UpdateAisAccount(ctx, params)

	assert.NoError(t, err)
	assert.Equal(t, params.Account_id, output.Account_id)
	mockAccessor.AssertExpectations(t)
}
