package middleware

import "github.com/google/wire"

var WireSet = wire.NewSet(
	NewAuthInterceptor,
	NewValidationInterceptor,
)
