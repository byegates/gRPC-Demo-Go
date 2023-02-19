package main

import (
	"log"
	"net"

	pb "github.com/byegates/gRPC-go/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var addr string = "0.0.0.0:50053"
var tag string = "[Server] "

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("%vFailed to listen on : %v\n", tag, err)
	}

	log.Printf("%vListening on %s\n", tag, addr)

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})
	reflection.Register(s)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("%vFailed to serve: %v\n", tag, err)
	}
}
