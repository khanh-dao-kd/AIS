package grpc

import (
	"ais_service/internal/handler/grpc/middleware"

	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	middleware.WireSet,
	NewGrpcHandler,
	NewServer,
)
