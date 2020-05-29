package grpc_client

import (
	"context"
	"fmt"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	ot "github.com/opentracing/opentracing-go"
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
	fmt.Println("CONNECTING Media metadata")
	conn, err := grpc.DialContext(context.Background(), env.MediaMetadataGrpcServer + ":" + env.MediaMetadataGrpcPort, opts...) // grpc.Dial(env.MediaMetadataGrpcServer + ":" + env.MediaMetadataGrpcPort, grpc.WithInsecure(), grpc.WithBlock())
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