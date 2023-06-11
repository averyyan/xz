package xz_error

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ErrorDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Status  string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ErrorDetail) Reset() {
	*x = ErrorDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_boot_error_proto_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorDetail) ProtoMessage() {}

func (x *ErrorDetail) ProtoReflect() protoreflect.Message {
	mi := &file_boot_error_proto_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorDetail.ProtoReflect.Descriptor instead.
func (*ErrorDetail) Descriptor() ([]byte, []int) {
	return file_boot_error_proto_error_proto_rawDescGZIP(), []int{0}
}

func (x *ErrorDetail) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ErrorDetail) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ErrorDetail) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_boot_error_proto_error_proto protoreflect.FileDescriptor

var file_boot_error_proto_error_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x62, 0x6f, 0x6f, 0x74, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x72, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x22, 0x53, 0x0a, 0x0b, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x10,
	0x5a, 0x0e, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6b, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_boot_error_proto_error_proto_rawDescOnce sync.Once
	file_boot_error_proto_error_proto_rawDescData = file_boot_error_proto_error_proto_rawDesc
)

func file_boot_error_proto_error_proto_rawDescGZIP() []byte {
	file_boot_error_proto_error_proto_rawDescOnce.Do(func() {
		file_boot_error_proto_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_boot_error_proto_error_proto_rawDescData)
	})
	return file_boot_error_proto_error_proto_rawDescData
}

var file_boot_error_proto_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_boot_error_proto_error_proto_goTypes = []interface{}{
	(*ErrorDetail)(nil), // 0: rk.api.v1.ErrorDetail
}
var file_boot_error_proto_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_boot_error_proto_error_proto_init() }
func file_boot_error_proto_error_proto_init() {
	if File_boot_error_proto_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_boot_error_proto_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorDetail); i {
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
			RawDescriptor: file_boot_error_proto_error_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_boot_error_proto_error_proto_goTypes,
		DependencyIndexes: file_boot_error_proto_error_proto_depIdxs,
		MessageInfos:      file_boot_error_proto_error_proto_msgTypes,
	}.Build()
	File_boot_error_proto_error_proto = out.File
	file_boot_error_proto_error_proto_rawDesc = nil
	file_boot_error_proto_error_proto_goTypes = nil
	file_boot_error_proto_error_proto_depIdxs = nil
}
