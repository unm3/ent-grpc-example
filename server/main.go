package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"ent-grpc-example/ent"
	"ent-grpc-example/ent/proto/entpb"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	log.Print("server is starting...")
	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		"localhost", "5432", "postgres", "postgres", "password"))
	if err != nil {
		log.Fatalf("failed to connect to db: %s", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed to create schema: %s", err)
	}

	svc := entpb.NewUserService(client)

	server := grpc.NewServer()

	entpb.RegisterUserServiceServer(server, svc)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
