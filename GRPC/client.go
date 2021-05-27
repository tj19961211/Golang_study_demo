package main

import (
	"context"
	"log"

	"study_gRPC/chat"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := &chat.Message{
		Body: "Hello from the Client!!",
	}

	response, err := c.SayHello(context.Background(), message)
	if err != nil {
		log.Fatal("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from Server: %s", response.Body)
}
