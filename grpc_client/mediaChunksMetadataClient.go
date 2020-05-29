package grpc_client

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"go-time-shift-api/Models"
	"google.golang.org/grpc"
	"log"

	pbMediaChunks "go-time-shift-api/proto/media_chunks"

	"golang.org/x/net/context"
	"github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	ot "github.com/opentracing/opentracing-go"
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

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithStreamInterceptor(
		grpc_opentracing.StreamClientInterceptor(
			grpc_opentracing.WithTracer(ot.GlobalTracer()))))
	opts = append(opts, grpc.WithUnaryInterceptor(
		grpc_opentracing.UnaryClientInterceptor(
			grpc_opentracing.WithTracer(ot.GlobalTracer()))))
	opts = append(opts, grpc.WithBlock())
	opts = append(opts, grpc.WithInsecure())

	env := Models.GetEnvStruct()
	fmt.Println("CONNECTING chunks metadata")

	conn, err := grpc.DialContext(context.Background(), env.ChunkMetadataGrpcServer + ":" + env.ChunkMetadataGrpcPort, opts...) // grpc.Dial(env.ChunkMetadataGrpcServer + ":" + env.ChunkMetadataGrpcPort, grpc.WithInsecure(), grpc.WithBlock())
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