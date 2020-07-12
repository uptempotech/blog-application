package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"uptempo.tech/blog-application/global"
	"uptempo.tech/blog-application/proto"
)

func main() {
	global.EmptyDB()

	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err.Error())
	}
	client := proto.NewAuthServiceClient(conn)
	_, err = client.Signup(context.Background(), &proto.SignupRequest{Username: "Carl", Email: "carl@gmail.com", Password: "examplestring"})
	if err != nil {
		log.Fatal(err.Error())
	}
}
