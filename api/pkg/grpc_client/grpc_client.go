package grpc_client

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
)

const (
	backoffLinear = 100 * time.Millisecond
)

func NewGRPCClientServiceConn(ctx context.Context, target string) (*grpc.ClientConn, error) {
	opts := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(backoffLinear)),
		grpc_retry.WithCodes(codes.NotFound, codes.Aborted),
	}

	clientGRPCConn, err := grpc.DialContext(
		ctx,
		target,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(opts...)),
	)
	if err != nil {
		return nil, err
	}

	return clientGRPCConn, nil
}
