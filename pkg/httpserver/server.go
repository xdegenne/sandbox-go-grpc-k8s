package server

import (
	"context"
	"example.com/hello/pkg/helloworld"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net/http"
)

type Server struct {
	helloworld.UnimplementedGreeterServer
}

func (s *Server) SayHello(c context.Context, r *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	reply := helloworld.HelloReply{}
	reply.Message = "Hello2 " + r.Name
	return &reply, nil
}

func (s *Server) Start() {
	log.Print("Listenning")

	grpcServer := grpc.NewServer()
	helloworld.RegisterGreeterServer(grpcServer, s)
	reflection.Register(grpcServer)

	srv := &http.Server{
		Addr:    ":5555",
		Handler: h2c.NewHandler(grpcHandlerFunc(grpcServer), &http2.Server{}),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}

func grpcHandlerFunc(grpcServer *grpc.Server) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		grpcServer.ServeHTTP(w, r)
	})
}
