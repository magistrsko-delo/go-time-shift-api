package Models

import (
	"fmt"
	"os"
)

var envStruct *Env

type Env struct {
	Port string
	Url string
	MediaManagerUrl string
	ChunkDownloaderUrl string
	MediaMetadataGrpcServer string
	MediaMetadataGrpcPort string
	ChunkMetadataGrpcServer string
	ChunkMetadataGrpcPort string
	SequenceServiceServer string
	SequenceServicePort string
	Env string
}

func InitEnv()  {
	envStruct = &Env{
		Env: 			  			os.Getenv("ENV"),
		MediaMetadataGrpcServer: 	os.Getenv("MEDIA_METADATA_GRPC_SERVER"),
		MediaMetadataGrpcPort:   	os.Getenv("MEDIA_METADATA_GRPC_PORT"),
		ChunkMetadataGrpcServer:  	os.Getenv("CHUNK_METADATA_GRPC_SERVER"),
		ChunkMetadataGrpcPort:		os.Getenv("CHUNK_METADATA_GRPC_PORT"),
		MediaManagerUrl: 			os.Getenv("MEDIA_MANAGER_URL"),
		ChunkDownloaderUrl:			os.Getenv("CHUNK_DOWNLOADER_URL"),
		Url:						os.Getenv("URL"),
		Port:						os.Getenv("PORT"),
		SequenceServiceServer:		os.Getenv("SEQUENCE_SERVICE_GRPC_SERVER"),
		SequenceServicePort:		os.Getenv("SEQUENCE_SERVICE_GRPC_PORT"),
	}
	fmt.Println(envStruct)
}

func GetEnvStruct() *Env  {
	return  envStruct
}