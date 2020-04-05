package grpc_client

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"go-time-shift-api/Models"
	"google.golang.org/grpc"
	"log"

	pbMediaChunks "go-time-shift-api/proto/media_chunks"
)

type MediaChunksClient struct {
	Conn *grpc.ClientConn
	client pbMediaChunks.MediaMetadataClient
}

func (mediaChunksClient *MediaChunksClient) GetMediaChunksInfoResolution(mediaId int32, resolution string) (*pbMediaChunks.MediaChunkInfoResponseRepeated, error)  {
	response, err := mediaChunksClient.client.GetMediaChunksResolution(context.Background(), &pbMediaChunks.MediaChunkResolutionRequest{
		Resolution:           resolution,
		MediaId:              mediaId,
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (mediaChunksClient *MediaChunksClient) GetAvailableResolutions() (*pbMediaChunks.ResolutionResponse, error) {
	response, err := mediaChunksClient.client.GetAvailableResolutions(context.Background(),  &empty.Empty{
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	});

	if err != nil {
		return nil, err
	}

	return response, nil
}

func InitChunkMetadataClient() *MediaChunksClient  {
	env := Models.GetEnvStruct()
	fmt.Println("CONNECTING chunks metadata")

	conn, err := grpc.Dial(env.ChunkMetadataGrpcServer + ":" + env.ChunkMetadataGrpcPort, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	fmt.Println("END CONNECTION chunk metadata")

	client := pbMediaChunks.NewMediaMetadataClient(conn)
	return &MediaChunksClient{
		Conn:    conn,
		client:  client,
	}

}