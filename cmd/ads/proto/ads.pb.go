// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ads.proto

package ads

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Advertisement struct {
	AdvertisementId                    int64    `protobuf:"varint,1,opt,name=advertisement_id,json=advertisementId,proto3" json:"advertisement_id,omitempty"`
	AdvertisementNo                    int64    `protobuf:"varint,2,opt,name=advertisement_no,json=advertisementNo,proto3" json:"advertisement_no,omitempty"`
	AdvertisementApplicationPositionId int64    `protobuf:"varint,3,opt,name=advertisement_application_position_id,json=advertisementApplicationPositionId,proto3" json:"advertisement_application_position_id,omitempty"`
	AdvertisementIndex                 int64    `protobuf:"varint,4,opt,name=advertisement_index,json=advertisementIndex,proto3" json:"advertisement_index,omitempty"`
	LinkType                           int64    `protobuf:"varint,5,opt,name=link_type,json=linkType,proto3" json:"link_type,omitempty"`
	AdvertisementImgUrl                string   `protobuf:"bytes,6,opt,name=advertisement_img_url,json=advertisementImgUrl,proto3" json:"advertisement_img_url,omitempty"`
	AdvertisementLinkUrl               string   `protobuf:"bytes,7,opt,name=advertisement_link_url,json=advertisementLinkUrl,proto3" json:"advertisement_link_url,omitempty"`
	ApplicationUrl                     string   `protobuf:"bytes,8,opt,name=application_url,json=applicationUrl,proto3" json:"application_url,omitempty"`
	Status                             int64    `protobuf:"varint,9,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt                          string   `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral               struct{} `json:"-"`
	XXX_unrecognized                   []byte   `json:"-"`
	XXX_sizecache                      int32    `json:"-"`
}

func (m *Advertisement) Reset()         { *m = Advertisement{} }
func (m *Advertisement) String() string { return proto.CompactTextString(m) }
func (*Advertisement) ProtoMessage()    {}
func (*Advertisement) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e37b36e10f7b333, []int{0}
}

func (m *Advertisement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Advertisement.Unmarshal(m, b)
}
func (m *Advertisement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Advertisement.Marshal(b, m, deterministic)
}
func (m *Advertisement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Advertisement.Merge(m, src)
}
func (m *Advertisement) XXX_Size() int {
	return xxx_messageInfo_Advertisement.Size(m)
}
func (m *Advertisement) XXX_DiscardUnknown() {
	xxx_messageInfo_Advertisement.DiscardUnknown(m)
}

var xxx_messageInfo_Advertisement proto.InternalMessageInfo

func (m *Advertisement) GetAdvertisementId() int64 {
	if m != nil {
		return m.AdvertisementId
	}
	return 0
}

func (m *Advertisement) GetAdvertisementNo() int64 {
	if m != nil {
		return m.AdvertisementNo
	}
	return 0
}

func (m *Advertisement) GetAdvertisementApplicationPositionId() int64 {
	if m != nil {
		return m.AdvertisementApplicationPositionId
	}
	return 0
}

func (m *Advertisement) GetAdvertisementIndex() int64 {
	if m != nil {
		return m.AdvertisementIndex
	}
	return 0
}

func (m *Advertisement) GetLinkType() int64 {
	if m != nil {
		return m.LinkType
	}
	return 0
}

func (m *Advertisement) GetAdvertisementImgUrl() string {
	if m != nil {
		return m.AdvertisementImgUrl
	}
	return ""
}

func (m *Advertisement) GetAdvertisementLinkUrl() string {
	if m != nil {
		return m.AdvertisementLinkUrl
	}
	return ""
}

func (m *Advertisement) GetApplicationUrl() string {
	if m != nil {
		return m.ApplicationUrl
	}
	return ""
}

func (m *Advertisement) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Advertisement) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

type FindRequest struct {
	PositionId           int64    `protobuf:"varint,1,opt,name=position_id,json=positionId,proto3" json:"position_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindRequest) Reset()         { *m = FindRequest{} }
func (m *FindRequest) String() string { return proto.CompactTextString(m) }
func (*FindRequest) ProtoMessage()    {}
func (*FindRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e37b36e10f7b333, []int{1}
}

func (m *FindRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindRequest.Unmarshal(m, b)
}
func (m *FindRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindRequest.Marshal(b, m, deterministic)
}
func (m *FindRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindRequest.Merge(m, src)
}
func (m *FindRequest) XXX_Size() int {
	return xxx_messageInfo_FindRequest.Size(m)
}
func (m *FindRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindRequest proto.InternalMessageInfo

func (m *FindRequest) GetPositionId() int64 {
	if m != nil {
		return m.PositionId
	}
	return 0
}

type FindResponse struct {
	Result               []*Advertisement `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *FindResponse) Reset()         { *m = FindResponse{} }
func (m *FindResponse) String() string { return proto.CompactTextString(m) }
func (*FindResponse) ProtoMessage()    {}
func (*FindResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e37b36e10f7b333, []int{2}
}

func (m *FindResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindResponse.Unmarshal(m, b)
}
func (m *FindResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindResponse.Marshal(b, m, deterministic)
}
func (m *FindResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindResponse.Merge(m, src)
}
func (m *FindResponse) XXX_Size() int {
	return xxx_messageInfo_FindResponse.Size(m)
}
func (m *FindResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindResponse proto.InternalMessageInfo

func (m *FindResponse) GetResult() []*Advertisement {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*Advertisement)(nil), "ads.Advertisement")
	proto.RegisterType((*FindRequest)(nil), "ads.FindRequest")
	proto.RegisterType((*FindResponse)(nil), "ads.FindResponse")
}

func init() {
	proto.RegisterFile("ads.proto", fileDescriptor_6e37b36e10f7b333)
}

var fileDescriptor_6e37b36e10f7b333 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xd1, 0x4b, 0xeb, 0x30,
	0x18, 0xc5, 0xe9, 0xed, 0x6e, 0xef, 0xfa, 0xed, 0xde, 0xbb, 0x19, 0x75, 0x04, 0x45, 0x1c, 0x05,
	0x71, 0x2a, 0x4c, 0x98, 0xbe, 0xe8, 0x5b, 0x5f, 0x84, 0x81, 0x88, 0x56, 0x7d, 0x2e, 0x71, 0x09,
	0x23, 0xac, 0x4b, 0x63, 0x92, 0x0e, 0xf7, 0x5f, 0xfa, 0x27, 0x49, 0xb2, 0xa2, 0xcd, 0xf4, 0xad,
	0x3d, 0xf9, 0x9d, 0x93, 0xd3, 0xaf, 0x1f, 0xc4, 0x84, 0xea, 0x91, 0x54, 0xa5, 0x29, 0x51, 0x48,
	0xa8, 0x4e, 0xde, 0x43, 0xf8, 0x97, 0xd2, 0x25, 0x53, 0x86, 0x6b, 0xb6, 0x60, 0xc2, 0xa0, 0x13,
	0xe8, 0x91, 0xa6, 0x90, 0x73, 0x8a, 0x83, 0x41, 0x30, 0x0c, 0xb3, 0xae, 0xa7, 0x4f, 0xe8, 0x77,
	0x54, 0x94, 0xf8, 0xd7, 0x0f, 0xe8, 0x5d, 0x89, 0x1e, 0xe0, 0xc8, 0x47, 0x89, 0x94, 0x05, 0x9f,
	0x12, 0xc3, 0x4b, 0x91, 0xcb, 0x52, 0x73, 0xf7, 0xc0, 0x29, 0x0e, 0x9d, 0x3f, 0xf1, 0xe0, 0xf4,
	0x8b, 0xbd, 0xaf, 0xd1, 0x09, 0x45, 0xe7, 0xb0, 0xbd, 0x51, 0x54, 0x50, 0xf6, 0x86, 0x5b, 0x2e,
	0x00, 0xf9, 0x5d, 0xed, 0x09, 0xda, 0x87, 0xb8, 0xe0, 0x62, 0x9e, 0x9b, 0x95, 0x64, 0xf8, 0xb7,
	0xc3, 0xda, 0x56, 0x78, 0x5a, 0x49, 0x86, 0xc6, 0xb0, 0xbb, 0x91, 0xb6, 0x98, 0xe5, 0x95, 0x2a,
	0x70, 0x34, 0x08, 0x86, 0x71, 0xe6, 0x5f, 0x35, 0x59, 0xcc, 0x9e, 0x55, 0x81, 0x2e, 0xa1, 0xef,
	0x7b, 0x5c, 0xbc, 0x35, 0xfd, 0x71, 0xa6, 0x1d, 0xef, 0xf4, 0x96, 0x8b, 0xb9, 0x75, 0x1d, 0x43,
	0xb7, 0xf9, 0xf1, 0x16, 0x6f, 0x3b, 0xfc, 0x7f, 0x43, 0xb6, 0x60, 0x1f, 0x22, 0x6d, 0x88, 0xa9,
	0x34, 0x8e, 0x5d, 0xd9, 0xfa, 0x0d, 0x1d, 0x00, 0x4c, 0x15, 0x23, 0x86, 0xd1, 0x9c, 0x18, 0x0c,
	0xce, 0x1b, 0xd7, 0x4a, 0x6a, 0x92, 0x11, 0x74, 0x6e, 0xb8, 0xa0, 0x19, 0x7b, 0xad, 0x98, 0x36,
	0xe8, 0x10, 0x3a, 0xcd, 0xf9, 0xae, 0x7f, 0x25, 0xc8, 0xcf, 0x39, 0x26, 0xd7, 0xf0, 0x77, 0xcd,
	0x6b, 0x59, 0x0a, 0xcd, 0xd0, 0x29, 0x44, 0x8a, 0xe9, 0xaa, 0x30, 0x38, 0x18, 0x84, 0xc3, 0xce,
	0x18, 0x8d, 0xec, 0xce, 0x78, 0x4b, 0x92, 0xd5, 0xc4, 0xf8, 0x0a, 0x20, 0xa5, 0xfa, 0x91, 0xa9,
	0x25, 0x9f, 0x32, 0x74, 0x06, 0x2d, 0x9b, 0x84, 0x7a, 0xce, 0xd1, 0x28, 0xb1, 0xb7, 0xd5, 0x50,
	0xd6, 0xd7, 0xbc, 0x44, 0x6e, 0x0b, 0x2f, 0x3e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x73, 0x23,
	0x28, 0x92, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AdsServiceClient is the client API for AdsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdsServiceClient interface {
	Find(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindResponse, error)
}

type adsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdsServiceClient(cc grpc.ClientConnInterface) AdsServiceClient {
	return &adsServiceClient{cc}
}

func (c *adsServiceClient) Find(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindResponse, error) {
	out := new(FindResponse)
	err := c.cc.Invoke(ctx, "/ads.AdsService/Find", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdsServiceServer is the server API for AdsService service.
type AdsServiceServer interface {
	Find(context.Context, *FindRequest) (*FindResponse, error)
}

// UnimplementedAdsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAdsServiceServer struct {
}

func (*UnimplementedAdsServiceServer) Find(ctx context.Context, req *FindRequest) (*FindResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}

func RegisterAdsServiceServer(s *grpc.Server, srv AdsServiceServer) {
	s.RegisterService(&_AdsService_serviceDesc, srv)
}

func _AdsService_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdsServiceServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ads.AdsService/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdsServiceServer).Find(ctx, req.(*FindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ads.AdsService",
	HandlerType: (*AdsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Find",
			Handler:    _AdsService_Find_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ads.proto",
}
