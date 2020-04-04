package main

import (
	"github.com/joho/godotenv"
	"go-time-shift-api/Models"
	"go-time-shift-api/grpc_client"
	"go-time-shift-api/grpc_server"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "go-time-shift-api/proto/timeshift_service_server"
)

func init()  {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	Models.InitEnv()
}


func main()  {

	// TODO pack into the request
	mediaMetadataClient := grpc_client.InitMediaMetadataGrpcClient()
	rsp, err := mediaMetadataClient.GetMediaMetadata(3)
	if err != nil {
		log.Println(err)
	}
	log.Println(rsp)

	mediaChunksClient := grpc_client.InitChunkMetadataClient()
	res, err := mediaChunksClient.GetMediaChunksInfoResolution(2, "1920x1080")

	if err != nil {
		log.Println(err)
	}
	log.Println(res.GetData())


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
