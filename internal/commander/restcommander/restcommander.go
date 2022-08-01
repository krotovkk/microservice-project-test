package restcommander

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	gw "gitlab.ozon.dev/krotovkk/homework/pkg/api"
)

func Run() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := gw.RegisterAdminHandlerFromEndpoint(ctx, mux, ":8081", opts)
	if err != nil {
		log.Panic(err)
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Panic(err)
	}
}
