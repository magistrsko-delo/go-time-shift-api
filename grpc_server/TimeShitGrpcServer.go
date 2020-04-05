package grpc_server

import (
	"fmt"
	"go-time-shift-api/Models"
	"go-time-shift-api/grpc_client"
	"golang.org/x/net/context"

	pbTimeShit "go-time-shift-api/proto/timeshift_service_server"
)

type TimeShiftServer struct {
	pbTimeShit.UnimplementedTimeshiftServer

	MediaChunksClientGrpc *grpc_client.MediaChunksClient
	MediaMetadataClientGrpc *grpc_client.MediaMetadataClient
	Env *Models.Env
}

func (server *TimeShiftServer) GetMediaChunkInformation(ctx context.Context, in *pbTimeShit.TimeShiftRequest) (*pbTimeShit.TimeShitResponse, error) {

	mediaMetadata, err := server.MediaMetadataClientGrpc.GetMediaMetadata(in.GetMediaId())

	if err != nil {
		return nil, err
	}

	return &pbTimeShit.TimeShitResponse{
		MediaId:                  mediaMetadata.GetMediaId(),
		Name:                     mediaMetadata.GetName(),
		SiteName:                 mediaMetadata.GetSiteName(),
		Length:                   int64(mediaMetadata.GetLength()),
		Status:                   mediaMetadata.GetStatus(),
		Thumbnail:                mediaMetadata.GetThumbnail(),
		ProjectId:                mediaMetadata.GetProjectId(),
		AwsBucketWholeMedia:      mediaMetadata.GetAwsBucketWholeMedia(),
		AwsStorageNameWholeMedia: mediaMetadata.GetAwsStorageNameWholeMedia(),
		Keywords:                 mediaMetadata.GetKeywords(),
		CreatedAt:                mediaMetadata.GetCreatedAt(),
		UpdatedAt:                mediaMetadata.GetUpdatedAt(),
		MediaUrl:                 server.Env.MediaManagerUrl + "v1/mediaManager/" + mediaMetadata.GetAwsBucketWholeMedia() + "/" + mediaMetadata.GetAwsStorageNameWholeMedia(),
		Data:                     server.getChunkResolutionResponse(in.GetMediaId()),
	}, nil
}

func (server *TimeShiftServer) getChunkResolutionResponse(mediaId int32) [] *pbTimeShit.ChunkResolutionResponse {

	response, err := server.MediaChunksClientGrpc.GetAvailableResolutions()

	if err != nil {
		fmt.Println( "ERR: ", err)
		return [] *pbTimeShit.ChunkResolutionResponse{}
	}

	chunksResolutionRsp := [] *pbTimeShit.ChunkResolutionResponse{}
	for i := 0; i < len(response.GetResolutions()); i++ {
		chunksResolutionRsp = append(chunksResolutionRsp, &pbTimeShit.ChunkResolutionResponse{
			Resolution:           response.GetResolutions()[i],
			Chunks:               server.getChunkResponse(mediaId, response.GetResolutions()[i]),
		})
	}

	return chunksResolutionRsp
}

func (server *TimeShiftServer) getChunkResponse(mediaId int32, resolution string) [] *pbTimeShit.ChunkResponse {

	response, err := server.MediaChunksClientGrpc.GetMediaChunksInfoResolution(mediaId, resolution)

	if err != nil {
		return [] *pbTimeShit.ChunkResponse{}
	}

	chunkRsp := [] *pbTimeShit.ChunkResponse{}

	for i:= 0; i < len(response.GetData()); i++ {
		chunkRsp = append(chunkRsp, &pbTimeShit.ChunkResponse{
			ChunkId:              response.GetData()[i].GetChunk().GetChunkId(),
			Position:             response.GetData()[i].GetPosition(),
			AwsBucketName:        response.GetData()[i].GetChunk().GetAwsBucketName(),
			AwsStorageName:       response.GetData()[i].GetChunk().GetAwsStorageName(),
			Length:               response.GetData()[i].GetChunk().GetLength(),
			CreatedAt:            response.GetData()[i].GetChunk().GetCreatedAt(),
			ChunksUrl:            server.Env.ChunkDownloaderUrl + "v1/chunk/bucket/" +
									response.GetData()[i].GetChunk().GetAwsBucketName() + "/storage/" +
									response.GetData()[i].GetChunk().GetAwsStorageName(),
		})
	}

	return chunkRsp
}

