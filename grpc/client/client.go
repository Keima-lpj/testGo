package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	hello "testGo/grpc/proto"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connect server err :%s", err)
	}
	defer conn.Close()

	client := hello.NewHelloClient(conn)

	res, err := client.HelloWorld(context.Background(), &hello.Request{Name: "Nero"})
	if err != nil {
		log.Fatalf("调用rpc服务err! %s", err)
	}

	fmt.Println(res.Res)

}
