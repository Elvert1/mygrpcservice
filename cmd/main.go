package main

import (
	"log"
	pb "mygprcservice/proto"
	"mygrpcservice/handler"
	"net"

	"golang.org/x/net/netutil"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	lis = netutil.LimitListener(lis, 110)

	grpcServer := grpc.NewServer()
	pb.RegisterFileServiceServer(grpcServer, &handler.Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
