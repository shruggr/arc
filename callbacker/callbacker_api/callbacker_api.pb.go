// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: callbacker/callbacker_api/callbacker_api.proto

package callbacker_api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// swagger:model RegisterCallbackResponse
type RegisterCallbackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *RegisterCallbackResponse) Reset() {
	*x = RegisterCallbackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_callbacker_callbacker_api_callbacker_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterCallbackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterCallbackResponse) ProtoMessage() {}

func (x *RegisterCallbackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_callbacker_callbacker_api_callbacker_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterCallbackResponse.ProtoReflect.Descriptor instead.
func (*RegisterCallbackResponse) Descriptor() ([]byte, []int) {
	return file_callbacker_callbacker_api_callbacker_api_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterCallbackResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// swagger:model HealthResponse
type HealthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok        bool                   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Details   string                 `protobuf:"bytes,2,opt,name=details,proto3" json:"details,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *HealthResponse) Reset() {
	*x = HealthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_callbacker_callbacker_api_callbacker_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthResponse) ProtoMessage() {}

func (x *HealthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_callbacker_callbacker_api_callbacker_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthResponse.ProtoReflect.Descriptor instead.
func (*HealthResponse) Descriptor() ([]byte, []int) {
	return file_callbacker_callbacker_api_callbacker_api_proto_rawDescGZIP(), []int{1}
}

func (x *HealthResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *HealthResponse) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

func (x *HealthResponse) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

// swagger:model Callback
type Callback struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash        []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Url         string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Token       string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	Status      int32  `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	BlockHash   []byte `protobuf:"bytes,5,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	BlockHeight uint64 `protobuf:"varint,6,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
}

func (x *Callback) Reset() {
	*x = Callback{}
	if protoimpl.UnsafeEnabled {
		mi := &file_callbacker_callbacker_api_callbacker_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Callback) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Callback) ProtoMessage() {}

func (x *Callback) ProtoReflect() protoreflect.Message {
	mi := &file_callbacker_callbacker_api_callbacker_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Callback.ProtoReflect.Descriptor instead.
func (*Callback) Descriptor() ([]byte, []int) {
	return file_callbacker_callbacker_api_callbacker_api_proto_rawDescGZIP(), []int{2}
}

func (x *Callback) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *Callback) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Callback) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Callback) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Callback) GetBlockHash() []byte {
	if x != nil {
		return x.BlockHash
	}
	return nil
}

func (x *Callback) GetBlockHeight() uint64 {
	if x != nil {
		return x.BlockHeight
	}
	return 0
}

var File_callbacker_callbacker_api_callbacker_api_proto protoreflect.FileDescriptor

var file_callbacker_callbacker_api_callbacker_api_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2f, 0x63, 0x61, 0x6c,
	0x6c, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0e, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c,
	0x0a, 0x18, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x74, 0x0a, 0x0e,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x18,
	0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x22, 0xa0, 0x01, 0x0a, 0x08, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12,
	0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68,
	0x61, 0x73, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x32, 0xad, 0x01, 0x0a, 0x0d, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x72, 0x41, 0x50, 0x49, 0x12, 0x42, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1e, 0x2e, 0x63, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x10, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12,
	0x18, 0x2e, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x1a, 0x28, 0x2e, 0x63, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x2e, 0x3b, 0x63, 0x61, 0x6c, 0x6c, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_callbacker_callbacker_api_callbacker_api_proto_rawDescOnce sync.Once
	file_callbacker_callbacker_api_callbacker_api_proto_rawDescData = file_callbacker_callbacker_api_callbacker_api_proto_rawDesc
)

func file_callbacker_callbacker_api_callbacker_api_proto_rawDescGZIP() []byte {
	file_callbacker_callbacker_api_callbacker_api_proto_rawDescOnce.Do(func() {
		file_callbacker_callbacker_api_callbacker_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_callbacker_callbacker_api_callbacker_api_proto_rawDescData)
	})
	return file_callbacker_callbacker_api_callbacker_api_proto_rawDescData
}

var file_callbacker_callbacker_api_callbacker_api_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_callbacker_callbacker_api_callbacker_api_proto_goTypes = []interface{}{
	(*RegisterCallbackResponse)(nil), // 0: callbacker_api.RegisterCallbackResponse
	(*HealthResponse)(nil),           // 1: callbacker_api.HealthResponse
	(*Callback)(nil),                 // 2: callbacker_api.Callback
	(*timestamppb.Timestamp)(nil),    // 3: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),            // 4: google.protobuf.Empty
}
var file_callbacker_callbacker_api_callbacker_api_proto_depIdxs = []int32{
	3, // 0: callbacker_api.HealthResponse.timestamp:type_name -> google.protobuf.Timestamp
	4, // 1: callbacker_api.CallbackerAPI.Health:input_type -> google.protobuf.Empty
	2, // 2: callbacker_api.CallbackerAPI.RegisterCallback:input_type -> callbacker_api.Callback
	1, // 3: callbacker_api.CallbackerAPI.Health:output_type -> callbacker_api.HealthResponse
	0, // 4: callbacker_api.CallbackerAPI.RegisterCallback:output_type -> callbacker_api.RegisterCallbackResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_callbacker_callbacker_api_callbacker_api_proto_init() }
func file_callbacker_callbacker_api_callbacker_api_proto_init() {
	if File_callbacker_callbacker_api_callbacker_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_callbacker_callbacker_api_callbacker_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterCallbackResponse); i {
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
		file_callbacker_callbacker_api_callbacker_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthResponse); i {
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
		file_callbacker_callbacker_api_callbacker_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Callback); i {
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
			RawDescriptor: file_callbacker_callbacker_api_callbacker_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_callbacker_callbacker_api_callbacker_api_proto_goTypes,
		DependencyIndexes: file_callbacker_callbacker_api_callbacker_api_proto_depIdxs,
		MessageInfos:      file_callbacker_callbacker_api_callbacker_api_proto_msgTypes,
	}.Build()
	File_callbacker_callbacker_api_callbacker_api_proto = out.File
	file_callbacker_callbacker_api_callbacker_api_proto_rawDesc = nil
	file_callbacker_callbacker_api_callbacker_api_proto_goTypes = nil
	file_callbacker_callbacker_api_callbacker_api_proto_depIdxs = nil
}
