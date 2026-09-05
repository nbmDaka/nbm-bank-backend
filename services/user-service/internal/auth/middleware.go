package auth

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)


func UnaryAuthInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {


	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return nil, status.Error(
			codes.Unauthenticated,
			"missing metadata",
		)
	}


	authHeader := md.Get("authorization")


	if len(authHeader)==0 {
		return nil, status.Error(
			codes.Unauthenticated,
			"missing token",
		)
	}


	token := authHeader[0]


	if !strings.HasPrefix(
		token,
		"Bearer ",
	){
		return nil, status.Error(
			codes.Unauthenticated,
			"invalid token format",
		)
	}


	return handler(ctx, req)
}