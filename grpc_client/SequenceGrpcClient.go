package grpc_client

import (
	"context"
	"fmt"
	"go-time-shift-api/Models"
	pbSequenceService "go-time-shift-api/proto/sequence_service"
	"google.golang.org/grpc"
	"log"
)

type SequenceServiceClient struct {
	Conn *grpc.ClientConn
	client pbSequenceService.SequenceMetadataClient
}

func (sequenceServiceClient *SequenceServiceClient) GetSequenceMedia(sequenceId int32) (*pbSequenceService.SequenceMediaResponse, error)  {
	response, err := sequenceServiceClient.client.GetSequenceMedia(context.Background(), &pbSequenceService.SequenceIdRequest{
		SequenceId:           sequenceId,
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func InitSequenceServiceMetadata() *SequenceServiceClient  {
	env := Models.GetEnvStruct()
	fmt.Println("CONNECTING sequence client")
	conn, err := grpc.Dial(env.SequenceServiceServer + ":" + env.SequenceServicePort, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	fmt.Println("END CONNECTION sequence client")

	client := pbSequenceService.NewSequenceMetadataClient(conn)
	return &SequenceServiceClient{
		Conn:    conn,
		client:  client,
	}
}