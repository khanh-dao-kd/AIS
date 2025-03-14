package consumer

import (
	"ais_service/internal/dataaccess/mq/producer"
	"ais_service/internal/logic"
	"context"
)

type AccountCreatedHandler interface {
	Handle(ctx context.Context, event producer.AccountEvent) error
}

type accountCreatedHandler struct {
	accountLogic logic.AccountLogic
}

func NewAccountCreatedHandler(accountLogic logic.AccountLogic) AccountCreatedHandler {
	return &accountCreatedHandler{
		accountLogic: accountLogic,
	}
}

func (a accountCreatedHandler) Handle(ctx context.Context, event producer.AccountEvent) error {
	_, err := a.accountLogic.CreateAisAccount(
		ctx,
		logic.CreateAisAccountParams{
			Account_id:     event.Account_id,
			Account_name:   event.Account_name,
			Account_type:   event.Account_type,
			Account_status: event.Account_status,
		},
	)
	if err != nil {
		return err
	}
	return nil
}
