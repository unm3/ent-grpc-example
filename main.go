package main

import (
	"context"
	"ent-grpc-example/ent"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		"localhost", "5432", "postgres", "postgres", "password"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	ctx := context.Background()
	if err := client.Schema.WriteTo(ctx, os.Stdout); err != nil {
		log.Fatal(err)
	}
}
