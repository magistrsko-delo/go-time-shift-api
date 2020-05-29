package Models

import (
	"fmt"
	"os"
)

var envStruct *Env

/*
CHUNK_DOWNLOADER_HEALTH="http://localhost:8888/"
MEDIA_METADATA_URL="http://localhost:8001/"
CHUNK_METADATA_URL="http://localhost:8004/"
SEQUENCE_SERVICE_URL="http://localhost:8009/"
*/

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
	ChunksDownloaderHealth string
	MediaMetadataUrl string
	ChunkMetadataUrl string
	SeqeunceServiceUrl string
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
		ChunksDownloaderHealth: 	os.Getenv("CHUNK_DOWNLOADER_HEALTH"),
		MediaMetadataUrl: 			os.Getenv("MEDIA_METADATA_URL"),
		ChunkMetadataUrl: 			os.Getenv("CHUNK_METADATA_URL"),
		SeqeunceServiceUrl: 		os.Getenv("SEQUENCE_SERVICE_URL"),
	}
	fmt.Println(envStruct)
}

func GetEnvStruct() *Env  {
	return  envStruct
}