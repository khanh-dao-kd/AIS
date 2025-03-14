package consumer

import "github.com/google/wire"

var WireSet = wire.NewSet(
	NewAccountCreatedHandler,
	NewConsumerServer,
)
