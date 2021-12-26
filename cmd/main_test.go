package main

import (
	"context"
	"log"
	"testing"
	"time"

	helloworldpb "github.com/wawa0210/go-grpc-journey/pkg/proto/helloworld"
	"google.golang.org/grpc"
)

func Test_Main(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	go runServer(ctx, ":8080")

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		log.Fatal(err)
	}
	defer conn.Close()
	c := helloworldpb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	r, err := c.GetCityWeather(ctx, &helloworldpb.GetCityWeatherRequest{Name: "shanghai"})
	if err != nil {
		log.Fatal(err)
	}

	expectMsg := "shanghai"
	gotMsg := r.GetName()

	if gotMsg != expectMsg {
		t.Fatalf("Expected error: %s, got %s", gotMsg, expectMsg)
	}

	// Contact the server and print out its response.
	res, err := c.SayHello(ctx, &helloworldpb.HelloRequest{Name: "jerry"})
	if err != nil {
		t.Fatal(err)
	}

	expectMsg = "jerry world"
	gotMsg = res.GetMessage()

	if gotMsg != expectMsg {
		t.Fatalf("Expected error: %s, got %s", gotMsg, expectMsg)
	}

	// time.Sleep(1000 * time.Minute)
	//test swagger api sdk

	// create the transport
	// transport := httptransport.New("localhost:8080", "", nil)

	// // create the API client, with the transport
	// client := apiclient.New(transport, strfmt.Default)

	// response, err := client.Greeter.GreeterSayHello(&greeter.GreeterSayHelloParams{Body: &models.HelloworldHelloRequest{Name: "jerry"}, Context: ctx})
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// expectMsg = "jerry world"
	// gotMsg = response.Payload.Message

	// if gotMsg != expectMsg {
	// 	t.Fatalf("Expected error: %s, got %s", gotMsg, expectMsg)
	// }
}
