package helloworld

import (
	context "context"

	helloworldpb "github.com/wawa0210/go-grpc-journey/pkg/proto/helloworld"
)

type Server struct {
	// Embed the unimplemented server
	helloworldpb.UnimplementedGreeterServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) SayHello(ctx context.Context, in *helloworldpb.HelloRequest) (*helloworldpb.HelloReply, error) {
	return &helloworldpb.HelloReply{Message: in.Name + " world"}, nil
}

func (s *Server) GetCityWeather(_ context.Context, in *helloworldpb.GetCityWeatherRequest) (*helloworldpb.CityWeatherResponse, error) {
	response := &helloworldpb.CityWeatherResponse{
		Name: in.Name,
	}

	if in.Name == "shanghai" {
		response.Temperature = "10"
	} else if in.Name == "beijing" {
		response.Temperature = "-10"
	} else {
		response.Temperature = "UNKOWN"
	}
	return response, nil
}
