package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	helloworldpb "github.com/wawa0210/go-grpc-journey/pkg/proto/helloworld"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		log.Fatal(err)
	}
	defer conn.Close()
	c := helloworldpb.NewGreeterClient(conn)
	ctx := context.Background()

	// Contact the server and print out its response.
	r, err := c.GetCityWeather(ctx, &helloworldpb.GetCityWeatherRequest{Name: "shanghai"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("grpc message: %v \n", r.GetName())

	// Contact the server and print out its response.
	res, err := c.SayHello(ctx, &helloworldpb.HelloRequest{Name: "jerry"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("grpc message: %v \n", res.GetMessage())
}
