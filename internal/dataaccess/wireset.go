package dataaccess

import (
	"ais_service/internal/dataaccess/database"
	"ais_service/internal/dataaccess/mq/consumer"
	"ais_service/internal/dataaccess/mq/producer"

	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	database.WireSet,
	producer.WireSet,
	consumer.WireSet,
)
