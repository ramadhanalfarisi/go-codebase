package grpc

import (
	"context"
	"log"
	"net"
	"os"
	"syscall"
	"time"
	"os/signal"

	"github.com/ramadhanalfarisi/go-codebase/config"
	gpc "google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type Grpc struct {
	App *gpc.Server
}

func NewGrpc() *Grpc {
	// ── Interceptors (middleware) ─────────────────────────────────────────────
	grpcServer := gpc.NewServer(
		gpc.ChainUnaryInterceptor(
			loggingInterceptor,
			recoveryInterceptor,
		),
		gpc.ChainStreamInterceptor(
			streamLoggingInterceptor,
		),
	)

	// Health check service — used by Docker healthcheck and load balancers
	healthSvc := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthSvc)

	// Reflection — allows tools like grpcurl to discover services without .proto
	reflection.Register(grpcServer)

	return &Grpc{
		App: grpcServer,
	}
}

func (g *Grpc) Run() {
	lis, err := net.Listen("tcp", config.PORT_GRPC)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// ── Graceful shutdown ─────────────────────────────────────────────────────
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
 
	go func() {
		log.Printf("gRPC server listening on %s", config.PORT_GRPC)
		if err := g.App.Serve(lis); err != nil {
			log.Fatalf("server error: %v", err)
		}
	}()
 
	<-ctx.Done()
	log.Println("Shutting down gRPC server...")
	g.App.GracefulStop()
	log.Println("Server stopped")
}

// ── Interceptors ─────────────────────────────────────────────────────────────

// loggingInterceptor logs every unary RPC call with duration.
func loggingInterceptor(
	ctx context.Context,
	req any,
	info *gpc.UnaryServerInfo,
	handler gpc.UnaryHandler,
) (any, error) {
	start := time.Now()
	resp, err := handler(ctx, req)
	log.Printf("[gRPC] %s | %v | err=%v", info.FullMethod, time.Since(start), err)
	return resp, err
}

// recoveryInterceptor catches panics and returns an INTERNAL gRPC error.
func recoveryInterceptor(
	ctx context.Context,
	req any,
	info *gpc.UnaryServerInfo,
	handler gpc.UnaryHandler,
) (resp any, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("[gRPC] panic recovered in %s: %v", info.FullMethod, r)
			err = gpc.Errorf(gpc.Code(nil), "internal server error") //nolint
		}
	}()
	return handler(ctx, req)
}

// streamLoggingInterceptor logs streaming RPC calls.
func streamLoggingInterceptor(
	srv any,
	ss gpc.ServerStream,
	info *gpc.StreamServerInfo,
	handler gpc.StreamHandler,
) error {
	start := time.Now()
	err := handler(srv, ss)
	log.Printf("[gRPC stream] %s | %v | err=%v", info.FullMethod, time.Since(start), err)
	return err
}
