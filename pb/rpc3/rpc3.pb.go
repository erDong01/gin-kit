// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: rpc3.proto

package rpc3

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RpcHead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"` //token
	SocketId  uint32 `protobuf:"varint,2,opt,name=SocketId,proto3" json:"SocketId,omitempty"`
	ActorName string `protobuf:"bytes,7,opt,name=ActorName,proto3" json:"ActorName,omitempty"`
	Reply     string `protobuf:"bytes,8,opt,name=Reply,proto3" json:"Reply,omitempty"` //call sessionid
	Code      int32  `protobuf:"varint,9,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg       string `protobuf:"bytes,10,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *RpcHead) Reset() {
	*x = RpcHead{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc3_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcHead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcHead) ProtoMessage() {}

func (x *RpcHead) ProtoReflect() protoreflect.Message {
	mi := &file_rpc3_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcHead.ProtoReflect.Descriptor instead.
func (*RpcHead) Descriptor() ([]byte, []int) {
	return file_rpc3_proto_rawDescGZIP(), []int{0}
}

func (x *RpcHead) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RpcHead) GetSocketId() uint32 {
	if x != nil {
		return x.SocketId
	}
	return 0
}

func (x *RpcHead) GetActorName() string {
	if x != nil {
		return x.ActorName
	}
	return ""
}

func (x *RpcHead) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

func (x *RpcHead) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *RpcHead) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type RpcPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FuncName string   `protobuf:"bytes,1,opt,name=FuncName,proto3" json:"FuncName,omitempty"`
	ArgLen   int32    `protobuf:"varint,2,opt,name=ArgLen,proto3" json:"ArgLen,omitempty"`
	RpcHead  *RpcHead `protobuf:"bytes,3,opt,name=RpcHead,proto3" json:"RpcHead,omitempty"`
	RpcBody  []byte   `protobuf:"bytes,4,opt,name=RpcBody,proto3" json:"RpcBody,omitempty"`
}

func (x *RpcPacket) Reset() {
	*x = RpcPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc3_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcPacket) ProtoMessage() {}

func (x *RpcPacket) ProtoReflect() protoreflect.Message {
	mi := &file_rpc3_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcPacket.ProtoReflect.Descriptor instead.
func (*RpcPacket) Descriptor() ([]byte, []int) {
	return file_rpc3_proto_rawDescGZIP(), []int{1}
}

func (x *RpcPacket) GetFuncName() string {
	if x != nil {
		return x.FuncName
	}
	return ""
}

func (x *RpcPacket) GetArgLen() int32 {
	if x != nil {
		return x.ArgLen
	}
	return 0
}

func (x *RpcPacket) GetRpcHead() *RpcHead {
	if x != nil {
		return x.RpcHead
	}
	return nil
}

func (x *RpcPacket) GetRpcBody() []byte {
	if x != nil {
		return x.RpcBody
	}
	return nil
}

//原始包
type Packet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`      //socketid
	Reply string `protobuf:"bytes,2,opt,name=Reply,proto3" json:"Reply,omitempty"` //call sessionid
	Buff  []byte `protobuf:"bytes,3,opt,name=Buff,proto3" json:"Buff,omitempty"`   //buff
}

func (x *Packet) Reset() {
	*x = Packet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc3_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Packet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Packet) ProtoMessage() {}

func (x *Packet) ProtoReflect() protoreflect.Message {
	mi := &file_rpc3_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Packet.ProtoReflect.Descriptor instead.
func (*Packet) Descriptor() ([]byte, []int) {
	return file_rpc3_proto_rawDescGZIP(), []int{2}
}

func (x *Packet) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Packet) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

func (x *Packet) GetBuff() []byte {
	if x != nil {
		return x.Buff
	}
	return nil
}

var File_rpc3_proto protoreflect.FileDescriptor

var file_rpc3_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x70, 0x63, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x70,
	0x63, 0x22, 0x8f, 0x01, 0x0a, 0x07, 0x52, 0x70, 0x63, 0x48, 0x65, 0x61, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x63, 0x74,
	0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x41, 0x63,
	0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x4d, 0x73, 0x67, 0x22, 0x81, 0x01, 0x0a, 0x09, 0x52, 0x70, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x75, 0x6e, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x75, 0x6e, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x41, 0x72, 0x67, 0x4c, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x41,
	0x72, 0x67, 0x4c, 0x65, 0x6e, 0x12, 0x26, 0x0a, 0x07, 0x52, 0x70, 0x63, 0x48, 0x65, 0x61, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x70, 0x63,
	0x48, 0x65, 0x61, 0x64, 0x52, 0x07, 0x52, 0x70, 0x63, 0x48, 0x65, 0x61, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x52, 0x70, 0x63, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x52, 0x70, 0x63, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x42, 0x0a, 0x06, 0x50, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x42, 0x75, 0x66, 0x66, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x42, 0x75, 0x66, 0x66, 0x42, 0x09, 0x5a, 0x07, 0x2e,
	0x2e, 0x2f, 0x72, 0x70, 0x63, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc3_proto_rawDescOnce sync.Once
	file_rpc3_proto_rawDescData = file_rpc3_proto_rawDesc
)

func file_rpc3_proto_rawDescGZIP() []byte {
	file_rpc3_proto_rawDescOnce.Do(func() {
		file_rpc3_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc3_proto_rawDescData)
	})
	return file_rpc3_proto_rawDescData
}

var file_rpc3_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_rpc3_proto_goTypes = []interface{}{
	(*RpcHead)(nil),   // 0: rpc.RpcHead
	(*RpcPacket)(nil), // 1: rpc.RpcPacket
	(*Packet)(nil),    // 2: rpc.Packet
}
var file_rpc3_proto_depIdxs = []int32{
	0, // 0: rpc.RpcPacket.RpcHead:type_name -> rpc.RpcHead
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc3_proto_init() }
func file_rpc3_proto_init() {
	if File_rpc3_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc3_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpcHead); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc3_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpcPacket); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc3_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Packet); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc3_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc3_proto_goTypes,
		DependencyIndexes: file_rpc3_proto_depIdxs,
		MessageInfos:      file_rpc3_proto_msgTypes,
	}.Build()
	File_rpc3_proto = out.File
	file_rpc3_proto_rawDesc = nil
	file_rpc3_proto_goTypes = nil
	file_rpc3_proto_depIdxs = nil
}
