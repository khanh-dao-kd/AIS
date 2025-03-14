package logic

import (
	"ais_service/internal/dataaccess/mq/producer"
	"context"
)

type PublisherLogic interface {
	PublishAisAccount(ctx context.Context, params PublishAisAccountParams) (PublishAisAccountOutput, error)
}

type publisherLogic struct {
	accountProducer producer.AccountProducer
}

func NewPublisher(accountProducer producer.AccountProducer) PublisherLogic {
	return &publisherLogic{
		accountProducer: accountProducer,
	}
}

func (p publisherLogic) PublishAisAccount(ctx context.Context, params PublishAisAccountParams) (PublishAisAccountOutput, error) {
	event := producer.AccountEvent{
		Account_id:     params.Account_id,
		Account_name:   params.Account_name,
		Account_type:   params.Account_type,
		Account_status: params.Account_status,
	}
	err := p.accountProducer.Produce(ctx, event)
	if err != nil {
		return PublishAisAccountOutput{}, err
	}
	return PublishAisAccountOutput{}, nil
}
