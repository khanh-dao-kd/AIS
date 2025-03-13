package handler

import (
	"ais_service/internal/handler/grpc"
	"ais_service/internal/handler/http"

	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	grpc.WireSet,
	http.WireSet,
)
