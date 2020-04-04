package main

import (
	"go-time-shift-api/grpc_server"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "go-time-shift-api/proto/timeshift_service_server"
)

func init()  {
}


func main()  {
	lis, err := net.Listen("tcp", ":9004")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterTimeshiftServer(s, &grpc_server.TimeShiftServer{})
	log.Println("gRPC server running on port: " + "9004")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
