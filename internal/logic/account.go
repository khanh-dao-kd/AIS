package logic

import (
	"ais_service/internal/dataaccess/database"
	"ais_service/internal/utils"
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AccountLogic interface {
	AddAisMessage(ctx context.Context, params AddAisMessaggParams) (AddAisMessageOutput, error)
	CreateAisAccount(ctx context.Context, params CreateAisAccountParams) (CreateAisAccountOutput, error)
	GetAisAccountByID(ctx context.Context, params GetAisAccountByIDParams) (GetAisAccountByIDOutput, error)
	UpdateAisAccount(ctx context.Context, params UpdateAccountStatusParams) (UpdateAccountStatusOutput, error)
}

type accountLogic struct {
	accountDataAccessor database.AisAccountDataAccessor
	logger              *zap.Logger
}

func NewAccountLogic(accountDataAccessor database.AisAccountDataAccessor, logger *zap.Logger) AccountLogic {
	return &accountLogic{
		accountDataAccessor: accountDataAccessor,
		logger:              logger,
	}
}

func (a accountLogic) AddAisMessage(ctx context.Context, params AddAisMessaggParams) (AddAisMessageOutput, error) {
	logger := utils.LoggerWithContext(ctx, a.logger)

	_, err := a.accountDataAccessor.GetAisAccountByIDForUpdate(ctx, params.Account_id)
	if err == nil {
		logger.Warn("Account_id is existed, updating the account",
			zap.Uint64("account_id", params.Account_id),
		)

		output, update_err := a.UpdateAisAccount(ctx, UpdateAccountStatusParams(params))
		if update_err != nil {
			return AddAisMessageOutput{}, update_err
		}
		return AddAisMessageOutput(output), nil
	}
	if status.Code(err) == codes.NotFound {
		logger.Warn("Account_id is not existed, create the account",
			zap.Uint64("account_id", params.Account_id),
		)

		output, create_err := a.CreateAisAccount(ctx, CreateAisAccountParams(params))
		if create_err != nil {
			return AddAisMessageOutput{}, create_err
		}
		return AddAisMessageOutput(output), nil
	}
	return AddAisMessageOutput{}, err
}

func (a accountLogic) CreateAisAccount(ctx context.Context, params CreateAisAccountParams) (CreateAisAccountOutput, error) {
	account_id, err := a.accountDataAccessor.CreateAisAccount(
		ctx,
		database.AisAccount{
			Account_id:     params.Account_id,
			Account_name:   params.Account_name,
			Account_type:   params.Account_type,
			Account_status: params.Account_status,
		},
	)
	if err != nil {
		return CreateAisAccountOutput{}, err
	}
	return CreateAisAccountOutput{Account_id: account_id}, nil
}
func (a accountLogic) GetAisAccountByID(ctx context.Context, params GetAisAccountByIDParams) (GetAisAccountByIDOutput, error) {
	account, err := a.accountDataAccessor.GetAisAccountByID(ctx, params.Account_id)
	if err != nil {
		return GetAisAccountByIDOutput{}, err
	}
	return GetAisAccountByIDOutput{
		Account_id:     account.Account_id,
		Account_name:   account.Account_name,
		Account_type:   account.Account_type,
		Account_status: account.Account_status,
	}, nil
}
func (a accountLogic) UpdateAisAccount(ctx context.Context, params UpdateAccountStatusParams) (UpdateAccountStatusOutput, error) {
	err := a.accountDataAccessor.UpdateAisAccount(
		ctx,
		database.AisAccount{
			Account_id:     params.Account_id,
			Account_name:   params.Account_name,
			Account_type:   params.Account_type,
			Account_status: params.Account_status,
		},
	)
	if err != nil {
		return UpdateAccountStatusOutput{}, err
	}
	return UpdateAccountStatusOutput{params.Account_id}, nil
}
