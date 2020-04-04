// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/media_chunks/mediachunksmetadata_service.proto

package mediachunksmetadata_service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MediaChunkInfoResponseRepeated struct {
	Data                 []*MediaChunkInfoResponse `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *MediaChunkInfoResponseRepeated) Reset()         { *m = MediaChunkInfoResponseRepeated{} }
func (m *MediaChunkInfoResponseRepeated) String() string { return proto.CompactTextString(m) }
func (*MediaChunkInfoResponseRepeated) ProtoMessage()    {}
func (*MediaChunkInfoResponseRepeated) Descriptor() ([]byte, []int) {
	return fileDescriptor_7241b774279ee94f, []int{0}
}

func (m *MediaChunkInfoResponseRepeated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MediaChunkInfoResponseRepeated.Unmarshal(m, b)
}
func (m *MediaChunkInfoResponseRepeated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MediaChunkInfoResponseRepeated.Marshal(b, m, deterministic)
}
func (m *MediaChunkInfoResponseRepeated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MediaChunkInfoResponseRepeated.Merge(m, src)
}
func (m *MediaChunkInfoResponseRepeated) XXX_Size() int {
	return xxx_messageInfo_MediaChunkInfoResponseRepeated.Size(m)
}
func (m *MediaChunkInfoResponseRepeated) XXX_DiscardUnknown() {
	xxx_messageInfo_MediaChunkInfoResponseRepeated.DiscardUnknown(m)
}

var xxx_messageInfo_MediaChunkInfoResponseRepeated proto.InternalMessageInfo

func (m *MediaChunkInfoResponseRepeated) GetData() []*MediaChunkInfoResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

type MediaChunkInfoResponse struct {
	MediaId              int32      `protobuf:"varint,1,opt,name=mediaId,proto3" json:"mediaId,omitempty"`
	Resolution           string     `protobuf:"bytes,2,opt,name=resolution,proto3" json:"resolution,omitempty"`
	Position             int32      `protobuf:"varint,3,opt,name=position,proto3" json:"position,omitempty"`
	Id                   int32      `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	Chunk                *ChunkInfo `protobuf:"bytes,5,opt,name=chunk,proto3" json:"chunk,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *MediaChunkInfoResponse) Reset()         { *m = MediaChunkInfoResponse{} }
func (m *MediaChunkInfoResponse) String() string { return proto.CompactTextString(m) }
func (*MediaChunkInfoResponse) ProtoMessage()    {}
func (*MediaChunkInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7241b774279ee94f, []int{1}
}

func (m *MediaChunkInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MediaChunkInfoResponse.Unmarshal(m, b)
}
func (m *MediaChunkInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MediaChunkInfoResponse.Marshal(b, m, deterministic)
}
func (m *MediaChunkInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MediaChunkInfoResponse.Merge(m, src)
}
func (m *MediaChunkInfoResponse) XXX_Size() int {
	return xxx_messageInfo_MediaChunkInfoResponse.Size(m)
}
func (m *MediaChunkInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MediaChunkInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MediaChunkInfoResponse proto.InternalMessageInfo

func (m *MediaChunkInfoResponse) GetMediaId() int32 {
	if m != nil {
		return m.MediaId
	}
	return 0
}

func (m *MediaChunkInfoResponse) GetResolution() string {
	if m != nil {
		return m.Resolution
	}
	return ""
}

func (m *MediaChunkInfoResponse) GetPosition() int32 {
	if m != nil {
		return m.Position
	}
	return 0
}

func (m *MediaChunkInfoResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MediaChunkInfoResponse) GetChunk() *ChunkInfo {
	if m != nil {
		return m.Chunk
	}
	return nil
}

type ResolutionResponse struct {
	Resolutions          []string `protobuf:"bytes,1,rep,name=resolutions,proto3" json:"resolutions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResolutionResponse) Reset()         { *m = ResolutionResponse{} }
func (m *ResolutionResponse) String() string { return proto.CompactTextString(m) }
func (*ResolutionResponse) ProtoMessage()    {}
func (*ResolutionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7241b774279ee94f, []int{2}
}

func (m *ResolutionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResolutionResponse.Unmarshal(m, b)
}
func (m *ResolutionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResolutionResponse.Marshal(b, m, deterministic)
}
func (m *ResolutionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResolutionResponse.Merge(m, src)
}
func (m *ResolutionResponse) XXX_Size() int {
	return xxx_messageInfo_ResolutionResponse.Size(m)
}
func (m *ResolutionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResolutionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResolutionResponse proto.InternalMessageInfo

func (m *ResolutionResponse) GetResolutions() []string {
	if m != nil {
		return m.Resolutions
	}
	return nil
}

type ChunkInfo struct {
	AwsBucketName        string   `protobuf:"bytes,1,opt,name=awsBucketName,proto3" json:"awsBucketName,omitempty"`
	AwsStorageName       string   `protobuf:"bytes,2,opt,name=awsStorageName,proto3" json:"awsStorageName,omitempty"`
	Length               float64  `protobuf:"fixed64,3,opt,name=length,proto3" json:"length,omitempty"`
	ChunkId              int32    `protobuf:"varint,4,opt,name=chunkId,proto3" json:"chunkId,omitempty"`
	CreatedAt            int64    `protobuf:"varint,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChunkInfo) Reset()         { *m = ChunkInfo{} }
func (m *ChunkInfo) String() string { return proto.CompactTextString(m) }
func (*ChunkInfo) ProtoMessage()    {}
func (*ChunkInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7241b774279ee94f, []int{3}
}

func (m *ChunkInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChunkInfo.Unmarshal(m, b)
}
func (m *ChunkInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChunkInfo.Marshal(b, m, deterministic)
}
func (m *ChunkInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChunkInfo.Merge(m, src)
}
func (m *ChunkInfo) XXX_Size() int {
	return xxx_messageInfo_ChunkInfo.Size(m)
}
func (m *ChunkInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ChunkInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ChunkInfo proto.InternalMessageInfo

func (m *ChunkInfo) GetAwsBucketName() string {
	if m != nil {
		return m.AwsBucketName
	}
	return ""
}

func (m *ChunkInfo) GetAwsStorageName() string {
	if m != nil {
		return m.AwsStorageName
	}
	return ""
}

func (m *ChunkInfo) GetLength() float64 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *ChunkInfo) GetChunkId() int32 {
	if m != nil {
		return m.ChunkId
	}
	return 0
}

func (m *ChunkInfo) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

type NewMediaChunkRequest struct {
	AwsBucketName        string   `protobuf:"bytes,1,opt,name=awsBucketName,proto3" json:"awsBucketName,omitempty"`
	AwsStorageName       string   `protobuf:"bytes,2,opt,name=awsStorageName,proto3" json:"awsStorageName,omitempty"`
	Length               float64  `protobuf:"fixed64,3,opt,name=length,proto3" json:"length,omitempty"`
	MediaId              int32    `protobuf:"varint,4,opt,name=mediaId,proto3" json:"mediaId,omitempty"`
	Resolution           string   `protobuf:"bytes,5,opt,name=resolution,proto3" json:"resolution,omitempty"`
	Position             int32    `protobuf:"varint,6,opt,name=position,proto3" json:"position,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewMediaChunkRequest) Reset()         { *m = NewMediaChunkRequest{} }
func (m *NewMediaChunkRequest) String() string { return proto.CompactTextString(m) }
func (*NewMediaChunkRequest) ProtoMessage()    {}
func (*NewMediaChunkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7241b774279ee94f, []int{4}
}

func (m *NewMediaChunkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewMediaChunkRequest.Unmarshal(m, b)
}
func (m *NewMediaChunkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewMediaChunkRequest.Marshal(b, m, deterministic)
}
func (m *NewMediaChunkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewMediaChunkRequest.Merge(m, src)
}
func (m *NewMediaChunkRequest) XXX_Size() int {
	return xxx_messageInfo_NewMediaChunkRequest.Size(m)
}
func (m *NewMediaChunkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewMediaChunkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewMediaChunkRequest proto.InternalMessageInfo

func (m *NewMediaChunkRequest) GetAwsBucketName() string {
	if m != nil {
		return m.AwsBucketName
	}
	return ""
}

func (m *NewMediaChunkRequest) GetAwsStorageName() string {
	if m != nil {
		return m.AwsStorageName
	}
	return ""
}

func (m *NewMediaChunkRequest) GetLength() float64 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *NewMediaChunkRequest) GetMediaId() int32 {
	if m != nil {
		return m.MediaId
	}
	return 0
}

func (m *NewMediaChunkRequest) GetResolution() string {
	if m != nil {
		return m.Resolution
	}
	return ""
}

func (m *NewMediaChunkRequest) GetPosition() int32 {
	if m != nil {
		return m.Position
	}
	return 0
}

type MediaChunkResolutionRequest struct {
	Resolution           string   `protobuf:"bytes,1,opt,name=resolution,proto3" json:"resolution,omitempty"`
	MediaId              int32    `protobuf:"varint,2,opt,name=mediaId,proto3" json:"mediaId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MediaChunkResolutionRequest) Reset()         { *m = MediaChunkResolutionRequest{} }
func (m *MediaChunkResolutionRequest) String() string { return proto.CompactTextString(m) }
func (*MediaChunkResolutionRequest) ProtoMessage()    {}
func (*MediaChunkResolutionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7241b774279ee94f, []int{5}
}

func (m *MediaChunkResolutionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MediaChunkResolutionRequest.Unmarshal(m, b)
}
func (m *MediaChunkResolutionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MediaChunkResolutionRequest.Marshal(b, m, deterministic)
}
func (m *MediaChunkResolutionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MediaChunkResolutionRequest.Merge(m, src)
}
func (m *MediaChunkResolutionRequest) XXX_Size() int {
	return xxx_messageInfo_MediaChunkResolutionRequest.Size(m)
}
func (m *MediaChunkResolutionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MediaChunkResolutionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MediaChunkResolutionRequest proto.InternalMessageInfo

func (m *MediaChunkResolutionRequest) GetResolution() string {
	if m != nil {
		return m.Resolution
	}
	return ""
}

func (m *MediaChunkResolutionRequest) GetMediaId() int32 {
	if m != nil {
		return m.MediaId
	}
	return 0
}

func init() {
	proto.RegisterType((*MediaChunkInfoResponseRepeated)(nil), "MediaChunkInfoResponseRepeated")
	proto.RegisterType((*MediaChunkInfoResponse)(nil), "MediaChunkInfoResponse")
	proto.RegisterType((*ResolutionResponse)(nil), "ResolutionResponse")
	proto.RegisterType((*ChunkInfo)(nil), "ChunkInfo")
	proto.RegisterType((*NewMediaChunkRequest)(nil), "NewMediaChunkRequest")
	proto.RegisterType((*MediaChunkResolutionRequest)(nil), "MediaChunkResolutionRequest")
}

func init() {
	proto.RegisterFile("proto/media_chunks/mediachunksmetadata_service.proto", fileDescriptor_7241b774279ee94f)
}

var fileDescriptor_7241b774279ee94f = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xed, 0xe6, 0x0b, 0x3c, 0x51, 0x7a, 0x58, 0x20, 0xb5, 0xd2, 0xaa, 0x58, 0x16, 0x42, 0x91,
	0x90, 0x36, 0x52, 0x40, 0xdc, 0x5b, 0x84, 0x44, 0x0e, 0xe9, 0x61, 0x39, 0x70, 0xe0, 0x50, 0x6d,
	0xec, 0xa9, 0x6b, 0xd5, 0xf1, 0x1a, 0xef, 0xba, 0x11, 0x47, 0xfe, 0x09, 0x57, 0xfe, 0x0e, 0xbf,
	0x08, 0x79, 0x5c, 0xc7, 0x0e, 0x2d, 0x11, 0xa7, 0xde, 0x3c, 0x33, 0xcf, 0xcf, 0xef, 0xcd, 0xf3,
	0xc0, 0xbb, 0x2c, 0xd7, 0x56, 0xcf, 0xd6, 0x18, 0xc6, 0xea, 0x32, 0xb8, 0x2e, 0xd2, 0x1b, 0x53,
	0x15, 0xd5, 0xf3, 0x1a, 0xad, 0x0a, 0x95, 0x55, 0x97, 0x06, 0xf3, 0xdb, 0x38, 0x40, 0x41, 0xf0,
	0xc9, 0x71, 0xa4, 0x75, 0x94, 0xe0, 0x8c, 0xaa, 0x55, 0x71, 0x35, 0xc3, 0x75, 0x66, 0xbf, 0x57,
	0x43, 0x7f, 0x09, 0xa7, 0xcb, 0x92, 0xe1, 0x43, 0xc9, 0xb0, 0x48, 0xaf, 0xb4, 0x44, 0x93, 0xe9,
	0xd4, 0xa0, 0xc4, 0x0c, 0x95, 0xc5, 0x90, 0xbf, 0x81, 0x5e, 0x49, 0xea, 0x32, 0xaf, 0x3b, 0x1d,
	0xce, 0x8f, 0xc4, 0x3f, 0xe0, 0x04, 0xf2, 0x7f, 0x32, 0x18, 0x3f, 0x0c, 0xe0, 0x2e, 0x3c, 0x21,
	0xad, 0x8b, 0xd0, 0x65, 0x1e, 0x9b, 0xf6, 0x65, 0x5d, 0xf2, 0x53, 0x80, 0x1c, 0x8d, 0x4e, 0x0a,
	0x1b, 0xeb, 0xd4, 0xed, 0x78, 0x6c, 0xea, 0xc8, 0x56, 0x87, 0x4f, 0xe0, 0x69, 0xa6, 0x4d, 0x4c,
	0xd3, 0x2e, 0xbd, 0xba, 0xad, 0xf9, 0x21, 0x74, 0xe2, 0xd0, 0xed, 0x51, 0xb7, 0x13, 0x87, 0xdc,
	0x83, 0x3e, 0x2d, 0xc3, 0xed, 0x7b, 0x6c, 0x3a, 0x9c, 0x83, 0x68, 0x84, 0x54, 0x03, 0xff, 0x3d,
	0x70, 0xb9, 0xe5, 0xde, 0xaa, 0xf3, 0x60, 0xd8, 0x7c, 0xd1, 0x90, 0x59, 0x47, 0xb6, 0x5b, 0xfe,
	0x2f, 0x06, 0xce, 0x96, 0x8c, 0xbf, 0x82, 0x91, 0xda, 0x98, 0xf3, 0x22, 0xb8, 0x41, 0x7b, 0xa1,
	0xd6, 0x48, 0x9e, 0x1c, 0xb9, 0xdb, 0xe4, 0xaf, 0xe1, 0x50, 0x6d, 0xcc, 0x67, 0xab, 0x73, 0x15,
	0x21, 0xc1, 0x2a, 0x77, 0x7f, 0x75, 0xf9, 0x18, 0x06, 0x09, 0xa6, 0x91, 0xbd, 0x26, 0x7f, 0x4c,
	0xde, 0x55, 0xe5, 0xce, 0x48, 0xf4, 0xa2, 0xb6, 0x58, 0x97, 0xfc, 0x04, 0x9c, 0x20, 0xa7, 0x80,
	0xce, 0x2c, 0x79, 0xed, 0xca, 0xa6, 0xe1, 0xff, 0x66, 0xf0, 0xfc, 0x02, 0x37, 0x4d, 0x12, 0x12,
	0xbf, 0x15, 0x68, 0xec, 0xe3, 0xc9, 0xae, 0xa3, 0xee, 0xed, 0x8b, 0xba, 0xbf, 0x37, 0xea, 0xc1,
	0x6e, 0xd4, 0xfe, 0x17, 0x38, 0x6e, 0x1b, 0x6a, 0x22, 0xac, 0xac, 0xed, 0x52, 0xb3, 0x7b, 0xd4,
	0x2d, 0x51, 0x9d, 0x1d, 0x51, 0xf3, 0x1f, 0x1d, 0x18, 0x11, 0xf3, 0xf2, 0xee, 0x80, 0xf8, 0x02,
	0x46, 0x69, 0x7b, 0x7d, 0xfc, 0x85, 0x78, 0x68, 0x9d, 0x93, 0x97, 0x62, 0xff, 0xf1, 0xf8, 0x07,
	0xfc, 0x2b, 0xb8, 0x11, 0xda, 0x06, 0x66, 0x1a, 0xe5, 0xfc, 0x44, 0xec, 0x31, 0xf4, 0x3f, 0xe4,
	0x9f, 0xe0, 0x28, 0x42, 0x7b, 0x76, 0xab, 0xe2, 0x44, 0xad, 0x12, 0x6c, 0x38, 0x0c, 0x1f, 0x8b,
	0xea, 0xec, 0x45, 0x7d, 0xf6, 0xe2, 0x63, 0x79, 0xf6, 0x93, 0x67, 0xe2, 0xfe, 0xdf, 0xef, 0x1f,
	0x9c, 0x0f, 0xa0, 0x17, 0xe5, 0x59, 0xb0, 0x1a, 0x10, 0xfc, 0xed, 0x9f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x40, 0xdf, 0x4d, 0x95, 0x6b, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MediaMetadataClient is the client API for MediaMetadata service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MediaMetadataClient interface {
	NewMediaChunk(ctx context.Context, in *NewMediaChunkRequest, opts ...grpc.CallOption) (*MediaChunkInfoResponseRepeated, error)
	GetMediaChunksResolution(ctx context.Context, in *MediaChunkResolutionRequest, opts ...grpc.CallOption) (*MediaChunkInfoResponseRepeated, error)
	GetAvailableResolutions(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ResolutionResponse, error)
}

type mediaMetadataClient struct {
	cc grpc.ClientConnInterface
}

func NewMediaMetadataClient(cc grpc.ClientConnInterface) MediaMetadataClient {
	return &mediaMetadataClient{cc}
}

func (c *mediaMetadataClient) NewMediaChunk(ctx context.Context, in *NewMediaChunkRequest, opts ...grpc.CallOption) (*MediaChunkInfoResponseRepeated, error) {
	out := new(MediaChunkInfoResponseRepeated)
	err := c.cc.Invoke(ctx, "/MediaMetadata/newMediaChunk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaMetadataClient) GetMediaChunksResolution(ctx context.Context, in *MediaChunkResolutionRequest, opts ...grpc.CallOption) (*MediaChunkInfoResponseRepeated, error) {
	out := new(MediaChunkInfoResponseRepeated)
	err := c.cc.Invoke(ctx, "/MediaMetadata/getMediaChunksResolution", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaMetadataClient) GetAvailableResolutions(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ResolutionResponse, error) {
	out := new(ResolutionResponse)
	err := c.cc.Invoke(ctx, "/MediaMetadata/getAvailableResolutions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MediaMetadataServer is the server API for MediaMetadata service.
type MediaMetadataServer interface {
	NewMediaChunk(context.Context, *NewMediaChunkRequest) (*MediaChunkInfoResponseRepeated, error)
	GetMediaChunksResolution(context.Context, *MediaChunkResolutionRequest) (*MediaChunkInfoResponseRepeated, error)
	GetAvailableResolutions(context.Context, *empty.Empty) (*ResolutionResponse, error)
}

// UnimplementedMediaMetadataServer can be embedded to have forward compatible implementations.
type UnimplementedMediaMetadataServer struct {
}

func (*UnimplementedMediaMetadataServer) NewMediaChunk(ctx context.Context, req *NewMediaChunkRequest) (*MediaChunkInfoResponseRepeated, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewMediaChunk not implemented")
}
func (*UnimplementedMediaMetadataServer) GetMediaChunksResolution(ctx context.Context, req *MediaChunkResolutionRequest) (*MediaChunkInfoResponseRepeated, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMediaChunksResolution not implemented")
}
func (*UnimplementedMediaMetadataServer) GetAvailableResolutions(ctx context.Context, req *empty.Empty) (*ResolutionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailableResolutions not implemented")
}

func RegisterMediaMetadataServer(s *grpc.Server, srv MediaMetadataServer) {
	s.RegisterService(&_MediaMetadata_serviceDesc, srv)
}

func _MediaMetadata_NewMediaChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewMediaChunkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaMetadataServer).NewMediaChunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MediaMetadata/NewMediaChunk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaMetadataServer).NewMediaChunk(ctx, req.(*NewMediaChunkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaMetadata_GetMediaChunksResolution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MediaChunkResolutionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaMetadataServer).GetMediaChunksResolution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MediaMetadata/GetMediaChunksResolution",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaMetadataServer).GetMediaChunksResolution(ctx, req.(*MediaChunkResolutionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaMetadata_GetAvailableResolutions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaMetadataServer).GetAvailableResolutions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MediaMetadata/GetAvailableResolutions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaMetadataServer).GetAvailableResolutions(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _MediaMetadata_serviceDesc = grpc.ServiceDesc{
	ServiceName: "MediaMetadata",
	HandlerType: (*MediaMetadataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "newMediaChunk",
			Handler:    _MediaMetadata_NewMediaChunk_Handler,
		},
		{
			MethodName: "getMediaChunksResolution",
			Handler:    _MediaMetadata_GetMediaChunksResolution_Handler,
		},
		{
			MethodName: "getAvailableResolutions",
			Handler:    _MediaMetadata_GetAvailableResolutions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/media_chunks/mediachunksmetadata_service.proto",
}
