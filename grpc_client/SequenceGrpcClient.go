package grpc_client

import (
	"context"
	"fmt"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	ot "github.com/opentracing/opentracing-go"
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
	fmt.Println("CONNECTING sequence client")
	conn, err := grpc.DialContext(context.Background(), env.SequenceServiceServer + ":" + env.SequenceServicePort, opts...) // grpc.Dial(env.SequenceServiceServer + ":" + env.SequenceServicePort, grpc.WithInsecure(), grpc.WithBlock())
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