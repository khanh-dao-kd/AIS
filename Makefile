generate:
	protoc -I=. \
		--go_out=internal/generated api/ais_api.proto \
		--go-grpc_out=internal/generated api/ais_api.proto \
		--grpc-gateway_out=internal/generated api/ais_api.proto \
		--grpc-gateway_opt generate_unbound_methods=true \
		--openapiv2_out . \
		--openapiv2_opt generate_unbound_methods=true \
		--validate_out="lang:go:internal/generated" \
		api/ais_api.proto