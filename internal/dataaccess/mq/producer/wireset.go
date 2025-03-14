package producer

import (
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	NewAccountProducer,
	NewClient,
)
