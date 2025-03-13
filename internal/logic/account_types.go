package logic

import "ais_service/internal/generated/grpc/ais_api"

type CreateAisAccountParams struct {
	Account_id     uint64
	Account_name   string
	Account_type   ais_api.AccountType
	Account_status ais_api.Status
}

type CreateAisAccountOutput struct {
	Account_id uint64
}

type GetAisAccountByIDParams struct {
	Account_id uint64
}

type GetAisAccountByIDOutput struct {
	Account_id     uint64
	Account_name   string
	Account_type   ais_api.AccountType
	Account_status ais_api.Status
}

type UpdateAccountStatusParams struct {
	Account_id     uint64
	Account_name   string
	Account_type   ais_api.AccountType
	Account_status ais_api.Status
}
