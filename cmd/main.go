package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"strings"

	mmux "github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/soheilhy/cmux"
	helloworld "github.com/wawa0210/go-grpc-journey/pkg/helloworld"
	helloworldpb "github.com/wawa0210/go-grpc-journey/pkg/proto/helloworld"
	"google.golang.org/grpc"
)

func runServer(ctx context.Context, endpoint string) error {
	// Create the main listener.
	l, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.Fatal(err)
	}
	// Create a cmux.
	m := cmux.New(l)

	// Match connections in order:
	// First grpc, then HTTP, and otherwise Go RPC/TCP.
	grpcL := m.Match(cmux.HTTP2())
	httpL := m.Match(cmux.HTTP1Fast())

	// Create your protocol servers.
	grpcS := grpc.NewServer()
	helloworldpb.RegisterGreeterServer(grpcS, &helloworld.Server{})

	mux := runtime.NewServeMux()
	helloworldpb.RegisterGreeterHandlerServer(ctx, mux, helloworld.NewServer())
	r := mmux.NewRouter()
	r.PathPrefix("/ui").Handler(http.StripPrefix("/ui", http.FileServer(http.Dir("assets/swagger"))))
	r.HandleFunc("/v1/citys/{name}", mux.ServeHTTP)
	httpS := &http.Server{
		Handler: r,
	}

	// Use the muxed listeners for your servers.
	go grpcS.Serve(grpcL)
	go httpS.Serve(httpL)

	log.Printf("grpc-gateway service listing %v", endpoint)
	// Start serving!
	if err := m.Serve(); !strings.Contains(err.Error(), "use of closed network connection") {
		return err
	}

	return nil
}

func main() {
	err := runServer(context.TODO(), ":8080")
	if err != nil {
		panic(err)
	}
}
