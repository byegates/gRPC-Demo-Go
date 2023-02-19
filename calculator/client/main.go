package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/byegates/gRPC-go/calculator/proto"
)

var addr string = "localhost:50053"
var tag string = "[Client] "

func main() {
	con, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("%vFailed to connect: %v\n", tag, err)
	}

	defer con.Close()

	c := pb.NewCalculatorServiceClient(con)

	doSum(c)
	doPrimes(c)
	doAvg(c)
	doMax(c)
	doSqrt(c, 100)
	doSqrt(c, -10)
	doSqrt(c, 10)
	//...
}
