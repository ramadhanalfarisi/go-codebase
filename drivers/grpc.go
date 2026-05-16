package drivers

import (
	"context"
	"log"
	"time"

	"github.com/ramadhanalfarisi/go-codebase/config"
	upb "github.com/ramadhanalfarisi/go-codebase/services/user/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClient struct {
	UserClient upb.UserControllerClient
	conn       *grpc.ClientConn
}

func NewGrpcClient() (*GrpcClient, func() error) {
	conn, err := grpc.NewClient(
		config.GRPC_SERVER,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(clientLoggingInterceptor),
	)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	client := upb.NewUserControllerClient(conn)

	grpcClient := &GrpcClient{
		UserClient: client,
		conn:       conn, // ← store conn in struct
	}

	return grpcClient, conn.Close // ← return Close as cleanup func
}

// clientLoggingInterceptor logs outgoing RPC calls on the client side.
func clientLoggingInterceptor(
	ctx context.Context,
	method string,
	req, reply any,
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	log.Printf("[client] %s | %v | err=%v", method, time.Since(start), err)
	return err
}
