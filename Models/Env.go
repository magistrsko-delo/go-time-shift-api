package Models

import (
	"fmt"
	"os"
)

var envStruct *Env

type Env struct {
	MediaMetadataGrpcServer string
	MediaMetadataGrpcPort string
	ChunkMetadataGrpcServer string
	ChunkMetadataGrpcPort string

	Env string
}

func InitEnv()  {
	envStruct = &Env{
		Env: 			  			os.Getenv("ENV"),
		MediaMetadataGrpcServer: 	os.Getenv("MEDIA_METADATA_GRPC_SERVER"),
		MediaMetadataGrpcPort:   	os.Getenv("MEDIA_METADATA_GRPC_PORT"),
		ChunkMetadataGrpcServer:  	os.Getenv("CHUNK_METADATA_GRPC_SERVER"),
		ChunkMetadataGrpcPort:		os.Getenv("CHUNK_METADATA_GRPC_PORT"),
	}
	fmt.Println(envStruct)
}

func GetEnvStruct() *Env  {
	return  envStruct
}