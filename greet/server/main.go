package main

import (
	"log"
	"net"

	pb "github.com/byegates/gRPC-go/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

var addr string = "0.0.0.0:50052"
var tag string = "[Server] "

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("%vFailed to listen on : %v\n", tag, err)
	}

	log.Printf("%vListening on %s\n", tag, addr)

	opts := []grpc.ServerOption{}
	tls := true

	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

		if err != nil {
			log.Fatalf("%vFailed loading certificates : %v & %vj\n", tag, certFile, keyFile)
		}

		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)
	pb.RegisterGreetServiceServer(s, &Server{})
	reflection.Register(s)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("%vFailed to serve: %v\n", tag, err)
	}
}
