// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: api/snowflake-service/v1/errors/snowflake.error.v1.proto

package snowflakeerrorv1

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

// ERROR ...
type ERROR int32

const (
	// UNSPECIFIED 未指定
	ERROR_UNSPECIFIED ERROR = 0
	// CANNOT_FOUNT_USABLE_ID 未找到可用的节点ID
	ERROR_CANNOT_FOUNT_USABLE_ID ERROR = 1
	// CANNOT_FOUNT_EXTEND_ID 未找到续期的节点ID
	ERROR_CANNOT_FOUNT_EXTEND_ID ERROR = 2
	// WORKER_INSTANCE_ID_EMPTY 实例ID不能为空
	ERROR_WORKER_INSTANCE_ID_EMPTY ERROR = 3
)

// Enum value maps for ERROR.
var (
	ERROR_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "CANNOT_FOUNT_USABLE_ID",
		2: "CANNOT_FOUNT_EXTEND_ID",
		3: "WORKER_INSTANCE_ID_EMPTY",
	}
	ERROR_value = map[string]int32{
		"UNSPECIFIED":              0,
		"CANNOT_FOUNT_USABLE_ID":   1,
		"CANNOT_FOUNT_EXTEND_ID":   2,
		"WORKER_INSTANCE_ID_EMPTY": 3,
	}
)

func (x ERROR) Enum() *ERROR {
	p := new(ERROR)
	*p = x
	return p
}

func (x ERROR) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ERROR) Descriptor() protoreflect.EnumDescriptor {
	return file_api_snowflake_service_v1_errors_snowflake_error_v1_proto_enumTypes[0].Descriptor()
}

func (ERROR) Type() protoreflect.EnumType {
	return &file_api_snowflake_service_v1_errors_snowflake_error_v1_proto_enumTypes[0]
}

func (x ERROR) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ERROR.Descriptor instead.
func (ERROR) EnumDescriptor() ([]byte, []int) {
	return file_api_snowflake_service_v1_errors_snowflake_error_v1_proto_rawDescGZIP(), []int{0}
}

var File_api_snowflake_service_v1_errors_snowflake_error_v1_proto protoreflect.FileDescriptor

var file_api_snowflake_service_v1_errors_snowflake_error_v1_proto_rawDesc = []byte{
	0x0a, 0x38, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x6e, 0x6f, 0x77, 0x66, 0x6c, 0x61, 0x6b, 0x65, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x73, 0x6e, 0x6f, 0x77, 0x66, 0x6c, 0x61, 0x6b, 0x65, 0x2e, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x73, 0x6e, 0x6f, 0x77,
	0x66, 0x6c, 0x61, 0x6b, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x6e, 0x6f, 0x77, 0x66, 0x6c, 0x61, 0x6b, 0x65, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x76, 0x31, 0x2a, 0x6e, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x12, 0x0f, 0x0a, 0x0b, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16,
	0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x55, 0x53, 0x41,
	0x42, 0x4c, 0x45, 0x5f, 0x49, 0x44, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x41, 0x4e, 0x4e,
	0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x45, 0x58, 0x54, 0x45, 0x4e, 0x44, 0x5f,
	0x49, 0x44, 0x10, 0x02, 0x12, 0x1c, 0x0a, 0x18, 0x57, 0x4f, 0x52, 0x4b, 0x45, 0x52, 0x5f, 0x49,
	0x4e, 0x53, 0x54, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x49, 0x44, 0x5f, 0x45, 0x4d, 0x50, 0x54, 0x59,
	0x10, 0x03, 0x42, 0xab, 0x01, 0x0a, 0x26, 0x73, 0x6e, 0x6f, 0x77, 0x66, 0x6c, 0x61, 0x6b, 0x65,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x6e, 0x6f,
	0x77, 0x66, 0x6c, 0x61, 0x6b, 0x65, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x76, 0x31, 0x42, 0x23, 0x53,
	0x6e, 0x6f, 0x77, 0x66, 0x6c, 0x61, 0x6b, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41,
	0x70, 0x69, 0x53, 0x6e, 0x6f, 0x77, 0x66, 0x6c, 0x61, 0x6b, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x56, 0x31, 0x50, 0x01, 0x5a, 0x5a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x69, 0x6b, 0x61, 0x69, 0x67, 0x75, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x6e,
	0x6f, 0x77, 0x66, 0x6c, 0x61, 0x6b, 0x65, 0x2d, 0x6e, 0x6f, 0x64, 0x65, 0x2d, 0x69, 0x64, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x73, 0x6e, 0x6f, 0x77, 0x66, 0x6c, 0x61, 0x6b, 0x65, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b,
	0x73, 0x6e, 0x6f, 0x77, 0x66, 0x6c, 0x61, 0x6b, 0x65, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_snowflake_service_v1_errors_snowflake_error_v1_proto_rawDescOnce sync.Once
	file_api_snowflake_service_v1_errors_snowflake_error_v1_proto_rawDescData = file_api_snowflake_service_v1_errors_snowflake_error_v1_proto_rawDesc
)

func file_api_snowflake_service_v1_errors_snowflake_error_v1_proto_rawDescGZIP() []byte {
	file_api_snowflake_service_v1_errors_snowflake_error_v1_proto_rawDescOnce.Do(func() {
		file_api_snowflake_service_v1_errors_snowflake_error_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_snowflake_service_v1_errors_snowflake_error_v1_proto_rawDescData)
	})
	return file_api_snowflake_service_v1_errors_snowflake_error_v1_proto_rawDescData
}

var file_api_snowflake_service_v1_errors_snowflake_error_v1_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_snowflake_service_v1_errors_snowflake_error_v1_proto_goTypes = []interface{}{
	(ERROR)(0), // 0: snowflake.service.api.snowflakeerrorv1.ERROR
}
var file_api_snowflake_service_v1_errors_snowflake_error_v1_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_snowflake_service_v1_errors_snowflake_error_v1_proto_init() }
func file_api_snowflake_service_v1_errors_snowflake_error_v1_proto_init() {
	if File_api_snowflake_service_v1_errors_snowflake_error_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_snowflake_service_v1_errors_snowflake_error_v1_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_snowflake_service_v1_errors_snowflake_error_v1_proto_goTypes,
		DependencyIndexes: file_api_snowflake_service_v1_errors_snowflake_error_v1_proto_depIdxs,
		EnumInfos:         file_api_snowflake_service_v1_errors_snowflake_error_v1_proto_enumTypes,
	}.Build()
	File_api_snowflake_service_v1_errors_snowflake_error_v1_proto = out.File
	file_api_snowflake_service_v1_errors_snowflake_error_v1_proto_rawDesc = nil
	file_api_snowflake_service_v1_errors_snowflake_error_v1_proto_goTypes = nil
	file_api_snowflake_service_v1_errors_snowflake_error_v1_proto_depIdxs = nil
}
