package dataaccess

import (
	"ais_service/internal/dataaccess/database"

	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	database.WireSet,
)
