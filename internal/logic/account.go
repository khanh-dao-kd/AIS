package logic

import (
	"ais_service/internal/dataaccess/database"
	"context"
)

type AccountLogic interface {
	CreateAisAccount(ctx context.Context, params CreateAisAccountParams) (CreateAisAccountOutput, error)
	GetAisAccountByID(ctx context.Context, params GetAisAccountByIDParams) (GetAisAccountByIDOutput, error)
	UpdateAisAccount(ctx context.Context, params UpdateAccountStatusParams) error
}

type accountLogic struct {
	accountDataAccessor database.AisAccountDataAccessor
}

func NewAccountLogic(accountDataAccessor database.AisAccountDataAccessor) AccountLogic {
	return &accountLogic{
		accountDataAccessor: accountDataAccessor,
	}
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
func (a accountLogic) UpdateAisAccount(ctx context.Context, params UpdateAccountStatusParams) error {
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
		return err
	}
	return nil
}
