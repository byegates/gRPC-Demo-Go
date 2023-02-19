package main

import (
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/byegates/gRPC-go/greet/proto"
)

var addr string = "localhost:50052"
var tag string = "[Client] "
var reqs = []*pb.GreetRequest{
	{FirstName: "邱晨"},
	{FirstName: "颜怡颜悦"},
	{FirstName: "鸟鸟"},
	{FirstName: "肖骁"},
	{FirstName: "艾力"},
}

func main() {
	tls := true // change to false if needed
	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("%vError while loading CA trust certificate: %v\n", tag, certFile)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	con, err := grpc.Dial(addr, opts...)

	if err != nil {
		log.Fatalf("%vFailed to connect: %v\n", tag, err)
	}

	defer con.Close()

	c := pb.NewGreetServiceClient(con)

	doGreetWithDeadline(c, 2*time.Second)
	doGreetWithDeadline(c, 4*time.Second)
	doGreet(c)
	doGreetStream(c)
	doLongGreet(c)
	doGreetEveryone(c)
	//...
}
