// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: qr.proto

package services

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

type QrRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message  string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Filename string `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
}

func (x *QrRequest) Reset() {
	*x = QrRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qr_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QrRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QrRequest) ProtoMessage() {}

func (x *QrRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qr_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QrRequest.ProtoReflect.Descriptor instead.
func (*QrRequest) Descriptor() ([]byte, []int) {
	return file_qr_proto_rawDescGZIP(), []int{0}
}

func (x *QrRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QrRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

type QrResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message  string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Filename string `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
}

func (x *QrResponse) Reset() {
	*x = QrResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qr_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QrResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QrResponse) ProtoMessage() {}

func (x *QrResponse) ProtoReflect() protoreflect.Message {
	mi := &file_qr_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QrResponse.ProtoReflect.Descriptor instead.
func (*QrResponse) Descriptor() ([]byte, []int) {
	return file_qr_proto_rawDescGZIP(), []int{1}
}

func (x *QrResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QrResponse) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

var File_qr_proto protoreflect.FileDescriptor

var file_qr_proto_rawDesc = []byte{
	0x0a, 0x08, 0x71, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x22, 0x41, 0x0a, 0x09, 0x51, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x42, 0x0a, 0x0a, 0x51, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x39, 0x0a, 0x02, 0x51,
	0x72, 0x12, 0x33, 0x0a, 0x02, 0x51, 0x72, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x51, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x51, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_qr_proto_rawDescOnce sync.Once
	file_qr_proto_rawDescData = file_qr_proto_rawDesc
)

func file_qr_proto_rawDescGZIP() []byte {
	file_qr_proto_rawDescOnce.Do(func() {
		file_qr_proto_rawDescData = protoimpl.X.CompressGZIP(file_qr_proto_rawDescData)
	})
	return file_qr_proto_rawDescData
}

var file_qr_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_qr_proto_goTypes = []interface{}{
	(*QrRequest)(nil),  // 0: services.QrRequest
	(*QrResponse)(nil), // 1: services.QrResponse
}
var file_qr_proto_depIdxs = []int32{
	0, // 0: services.Qr.Qr:input_type -> services.QrRequest
	1, // 1: services.Qr.Qr:output_type -> services.QrResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_qr_proto_init() }
func file_qr_proto_init() {
	if File_qr_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_qr_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QrRequest); i {
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
		file_qr_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QrResponse); i {
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
			RawDescriptor: file_qr_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_qr_proto_goTypes,
		DependencyIndexes: file_qr_proto_depIdxs,
		MessageInfos:      file_qr_proto_msgTypes,
	}.Build()
	File_qr_proto = out.File
	file_qr_proto_rawDesc = nil
	file_qr_proto_goTypes = nil
	file_qr_proto_depIdxs = nil
}