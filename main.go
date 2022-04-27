package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"test/test-grpc/server/grpcimplementation"
)

var GRPCInstance grpcimplementation.GrpcServiceServer

func main() {
	startServer()
}

func startServer() {
	lis, err := net.Listen("tcp", "172.25.0.100:4424")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	grpcimplementation.RegisterGrpcServiceServer(server, &grpcimplementation.ServerGrpc{})

	err = server.Serve(lis)
	if err != nil {
		log.Fatalln(err)
	}
}
