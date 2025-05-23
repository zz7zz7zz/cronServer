// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: appreview.proto

// 指定文件生成出来的package

package appreview

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AppReviewRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ver           string                 `protobuf:"bytes,1,opt,name=ver,proto3" json:"ver,omitempty"`
	Pkg           string                 `protobuf:"bytes,2,opt,name=pkg,proto3" json:"pkg,omitempty"`
	Platform      string                 `protobuf:"bytes,3,opt,name=platform,proto3" json:"platform,omitempty"`
	Status        string                 `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AppReviewRequest) Reset() {
	*x = AppReviewRequest{}
	mi := &file_appreview_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AppReviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppReviewRequest) ProtoMessage() {}

func (x *AppReviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_appreview_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppReviewRequest.ProtoReflect.Descriptor instead.
func (*AppReviewRequest) Descriptor() ([]byte, []int) {
	return file_appreview_proto_rawDescGZIP(), []int{0}
}

func (x *AppReviewRequest) GetVer() string {
	if x != nil {
		return x.Ver
	}
	return ""
}

func (x *AppReviewRequest) GetPkg() string {
	if x != nil {
		return x.Pkg
	}
	return ""
}

func (x *AppReviewRequest) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *AppReviewRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type AppReviewResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ver           string                 `protobuf:"bytes,1,opt,name=ver,proto3" json:"ver,omitempty"`
	Pkg           string                 `protobuf:"bytes,2,opt,name=pkg,proto3" json:"pkg,omitempty"`
	Platform      string                 `protobuf:"bytes,3,opt,name=platform,proto3" json:"platform,omitempty"`
	Status        string                 `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Key           string                 `protobuf:"bytes,5,opt,name=key,proto3" json:"key,omitempty"`
	Message       string                 `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AppReviewResponse) Reset() {
	*x = AppReviewResponse{}
	mi := &file_appreview_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AppReviewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppReviewResponse) ProtoMessage() {}

func (x *AppReviewResponse) ProtoReflect() protoreflect.Message {
	mi := &file_appreview_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppReviewResponse.ProtoReflect.Descriptor instead.
func (*AppReviewResponse) Descriptor() ([]byte, []int) {
	return file_appreview_proto_rawDescGZIP(), []int{1}
}

func (x *AppReviewResponse) GetVer() string {
	if x != nil {
		return x.Ver
	}
	return ""
}

func (x *AppReviewResponse) GetPkg() string {
	if x != nil {
		return x.Pkg
	}
	return ""
}

func (x *AppReviewResponse) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *AppReviewResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *AppReviewResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *AppReviewResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_appreview_proto protoreflect.FileDescriptor

const file_appreview_proto_rawDesc = "" +
	"\n" +
	"\x0fappreview.proto\x12\tappreview\"j\n" +
	"\x10AppReviewRequest\x12\x10\n" +
	"\x03ver\x18\x01 \x01(\tR\x03ver\x12\x10\n" +
	"\x03pkg\x18\x02 \x01(\tR\x03pkg\x12\x1a\n" +
	"\bplatform\x18\x03 \x01(\tR\bplatform\x12\x16\n" +
	"\x06status\x18\x04 \x01(\tR\x06status\"\x97\x01\n" +
	"\x11AppReviewResponse\x12\x10\n" +
	"\x03ver\x18\x01 \x01(\tR\x03ver\x12\x10\n" +
	"\x03pkg\x18\x02 \x01(\tR\x03pkg\x12\x1a\n" +
	"\bplatform\x18\x03 \x01(\tR\bplatform\x12\x16\n" +
	"\x06status\x18\x04 \x01(\tR\x06status\x12\x10\n" +
	"\x03key\x18\x05 \x01(\tR\x03key\x12\x18\n" +
	"\amessage\x18\x06 \x01(\tR\amessageB\x12Z\x10../src/appreviewb\x06proto3"

var (
	file_appreview_proto_rawDescOnce sync.Once
	file_appreview_proto_rawDescData []byte
)

func file_appreview_proto_rawDescGZIP() []byte {
	file_appreview_proto_rawDescOnce.Do(func() {
		file_appreview_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_appreview_proto_rawDesc), len(file_appreview_proto_rawDesc)))
	})
	return file_appreview_proto_rawDescData
}

var file_appreview_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_appreview_proto_goTypes = []any{
	(*AppReviewRequest)(nil),  // 0: appreview.AppReviewRequest
	(*AppReviewResponse)(nil), // 1: appreview.AppReviewResponse
}
var file_appreview_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_appreview_proto_init() }
func file_appreview_proto_init() {
	if File_appreview_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_appreview_proto_rawDesc), len(file_appreview_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_appreview_proto_goTypes,
		DependencyIndexes: file_appreview_proto_depIdxs,
		MessageInfos:      file_appreview_proto_msgTypes,
	}.Build()
	File_appreview_proto = out.File
	file_appreview_proto_goTypes = nil
	file_appreview_proto_depIdxs = nil
}
