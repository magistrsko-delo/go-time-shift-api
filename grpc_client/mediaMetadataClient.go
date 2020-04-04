package grpc_client

import (
	"context"
	"fmt"
	"go-time-shift-api/Models"
	"google.golang.org/grpc"
	pbMediaMetadata "go-time-shift-api/proto/media_metadata"
	"log"
)

type MediaMetadataClient struct {
	Conn *grpc.ClientConn
	client pbMediaMetadata.MediaMetadataClient
}

func (mediaMetadataClient *MediaMetadataClient) GetMediaMetadata(mediaId int32) (*pbMediaMetadata.MediaMetadataResponse, error)  {

	response, err := mediaMetadataClient.client.GetMediaMetadata(context.Background(), &pbMediaMetadata.GetMediaMetadataRequest{
		MediaId:	mediaId,
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}



func InitMediaMetadataGrpcClient() *MediaMetadataClient  {
	env := Models.GetEnvStruct()
	fmt.Println("CONNECTING Media metadata")
	conn, err := grpc.Dial(env.MediaMetadataGrpcServer + ":" + env.MediaMetadataGrpcPort, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	fmt.Println("END CONNECTION mediaMetadata")

	client := pbMediaMetadata.NewMediaMetadataClient(conn)
	return &MediaMetadataClient{
		Conn:    conn,
		client:  client,
	}
}