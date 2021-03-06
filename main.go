package main

import (
	"github.com/joho/godotenv"
	"go-time-shift-api/Models"
	"go-time-shift-api/grpc_client"
	"go-time-shift-api/grpc_server"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"

	pb "go-time-shift-api/proto/timeshift_service_server"
)

func init()  {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	Models.InitEnv()
}


func main()  {
	var wg sync.WaitGroup

	wg.Add(1)
	go gRPCServer()

	wg.Wait()
}

func gRPCServer()  {
	lis, err := net.Listen("tcp", ":" + Models.GetEnvStruct().Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterTimeshiftServer(s, &grpc_server.TimeShiftServer{
		UnimplementedTimeshiftServer: 	pb.UnimplementedTimeshiftServer{},
		MediaChunksClientGrpc:        	grpc_client.InitChunkMetadataClient(),
		MediaMetadataClientGrpc:      	grpc_client.InitMediaMetadataGrpcClient(),
		SequenceServiceClientGrpc: 		grpc_client.InitSequenceServiceMetadata(),
		Env:							Models.GetEnvStruct(),
	})
	log.Println("gRPC server running on: " + ":" + Models.GetEnvStruct().Port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}