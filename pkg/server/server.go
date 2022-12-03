package server

import (
	"context"
	"example.com/hello/pkg/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type Server struct {
	helloworld.UnimplementedGreeterServer
	Config ServerConfig
}

type ServerConfig struct {
	Address string
}

func (s *Server) SayHello(c context.Context, r *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Printf("Responding to: %s", r.Name)
	reply := helloworld.HelloReply{}
	reply.Message = "Hello " + r.Name
	return &reply, nil
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", s.Config.Address)
	log.Printf("Listenning on %s", s.Config.Address)

	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	helloworld.RegisterGreeterServer(grpcServer, s)
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
