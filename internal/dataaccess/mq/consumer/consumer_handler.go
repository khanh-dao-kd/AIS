package consumer

import (
	"context"
	"os"
)

type HandlerFunc func(ctx context.Context, topicName string, payload []byte) error

type consumerHandler struct {
	handlerFunc       HandlerFunc
	exitSignalChannel chan os.Signal
}

func NewConsumerHandler(handlerFunc HandlerFunc, exitSignalChannel chan os.Signal) *consumerHandler {
	return &consumerHandler{
		handlerFunc:       handlerFunc,
		exitSignalChannel: exitSignalChannel,
	}
}
