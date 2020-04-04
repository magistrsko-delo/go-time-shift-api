package grpc_server

import (
	pbTimeShit "go-time-shift-api/proto/timeshift_service_server"
	"golang.org/x/net/context"
	"log"
)

type TimeShiftServer struct {
	pbTimeShit.UnimplementedTimeshiftServer
}

func (s *TimeShiftServer) GetMediaChunkInformation(ctx context.Context, in *pbTimeShit.TimeShiftRequest) (*pbTimeShit.TimeShitResponse, error) {
	log.Printf("Received: %v", in.GetMediaId())

	return &pbTimeShit.TimeShitResponse{
		MediaId:                  0,
		Name:                     "",
		SiteName:                 "",
		Length:                   0,
		Status:                   0,
		Thumbnail:                "",
		ProjectId:                0,
		AwsBucketWholeMedia:      "",
		AwsStorageNameWholeMedia: "",
		Keywords:                 []string{"test1", "test2"},
		CreatedAt:                0,
		UpdatedAt:                0,
		MediaUrl:                 "",
		Data:                     [] *pbTimeShit.ChunkResolutionResponse{
			&pbTimeShit.ChunkResolutionResponse{
				Resolution:           "1920x1080",
				Chunks:               [] *pbTimeShit.ChunkResponse {
					&pbTimeShit.ChunkResponse{
						ChunkId:              1,
						Position:             0,
						AwsBucketName:        "test2",
						AwsStorageName:       "",
						Length:               0,
						CreatedAt:            0,
						ChunksUrl:            "http://localhost",
					},
					&pbTimeShit.ChunkResponse{
						ChunkId:              2,
						Position:             1,
						AwsBucketName:        "",
						AwsStorageName:       "",
						Length:               0,
						CreatedAt:            0,
						ChunksUrl:            "",
					},
				},
			},
			&pbTimeShit.ChunkResolutionResponse{
				Resolution:           "1280:x720",
				Chunks:               [] *pbTimeShit.ChunkResponse{},
			},
		},
	}, nil
}

