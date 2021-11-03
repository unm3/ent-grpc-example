package main

import (
	"context"
	"log"

	"ent-grpc-example/ent/proto/entpb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed connect to server: %s", err)
	}
	defer conn.Close()

	client := entpb.NewUserServiceClient(conn)

	// Create a new User
	ctx := context.Background()
	created, err := client.Create(ctx, &entpb.CreateUserRequest{
		User: &entpb.User{
			Name:  "hoge",
			Email: "hoge@exmaple.com",
		},
	})
	if err != nil {
		log.Fatalf("failed to create user: %s", err)
	}

	log.Printf("created an user:ID: %d, Name: %s, Email: %s", created.Id, created.Name, created.Email)

	// Get an User
	got, err := client.Get(ctx, &entpb.GetUserRequest{
		Id: created.Id,
	})
	if err != nil {
		log.Fatalf("failed to get an user: %s", err)
	}

	log.Printf("got an user:ID: %d, Name: %s, Email: %s", got.Id, got.Name, got.Email)

	// Update an User
	updated, err := client.Update(ctx, &entpb.UpdateUserRequest{
		User: &entpb.User{
			Id:    got.Id,
			Name:  "fuga",
			Email: "fuga@example.com",
		},
	})
	if err != nil {
		log.Fatalf("failed to update an user: %s", err)
	}

	log.Printf("updated an user: ID: %d Name: %s, Email: %s", updated.Id, updated.Name, updated.Email)

	/*Delete an User
	_, err = client.Delete(ctx, &entpb.DeleteUserRequest{
		Id: updated.Id,
	})
	if err != nil {
		log.Fatalf("failed to delete an user: %s", err)
	}

	log.Printf("deleted an user: ID: %d Name: %s Email: %s", updated.Id, updated.Name, updated.Email)
	*/
}
