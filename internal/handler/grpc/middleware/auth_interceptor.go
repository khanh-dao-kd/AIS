package middleware

import (
	"context"

	"google.golang.org/grpc"
)

type AuthInterceptor interface {
	JWTAuthMiddleware(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error)
}

type authInterceptor struct{}

func NewAuthInterceptor() AuthInterceptor {
	return &authInterceptor{}
}

func (a authInterceptor) JWTAuthMiddleware(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	// md, ok := metadata.FromIncomingContext(ctx)
	// if !ok {
	// 	return nil, status.Error(codes.Unauthenticated, "metadata is not provided")
	// }

	// // extract token from authorization header
	// token := md["authorization"]
	// if len(token) == 0 {
	// 	return nil, status.Error(codes.Unauthenticated, "authorization token is not provided")
	// }

	// // validate token and retrieve the userID
	// userID, err := a.authService.ValidateToken(token[0])
	// if err != nil {
	// 	return nil, status.Error(codes.Unauthenticated, "invalid token: %v", err)
	// }

	// // add our user ID to the context, so we can use it in our RPC handler
	// ctx = context.WithValue(ctx, "user_id", userID)

	// call our handler
	return handler(ctx, req)
}
