package main

import (
	"context"
	"log"

	"github.com/uptempotech/blog-application/global"
	"github.com/uptempotech/blog-application/proto"
	"google.golang.org/grpc"
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
