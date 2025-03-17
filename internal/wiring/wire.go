//go:build wireinject

package wiring

import (
	"ais_service/internal/configs"
	"ais_service/internal/dataaccess"
	"ais_service/internal/handler"
	"ais_service/internal/logic"
	"ais_service/internal/server"
	"ais_service/internal/utils"

	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	configs.WireSet,
	utils.WireSet,
	dataaccess.WireSet,
	logic.WireSet,
	handler.WireSet,
	server.WireSet,
)

func InitializeServer(configFilePath configs.ConfigFilePath) (*server.StandaloneServer, func(), error) {
	wire.Build(WireSet)
	return nil, nil, nil
}
