package main

import (
	"context"
	"log"
	"net/http"

	pb "demo/shared/proto"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	err := pb.RegisterTodoServiceHandlerFromEndpoint(ctx, mux, "server:50051", opts)
	if err != nil {
		log.Fatalf("failed to start gateway: %v", err)
	}

	log.Println("HTTP gateway started on :8080")
	http.ListenAndServe(":8080", mux)
}
