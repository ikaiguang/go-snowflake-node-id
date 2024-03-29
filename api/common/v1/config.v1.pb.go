// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: api/common/v1/config.v1.proto

package commonv1

import (
	v1 "github.com/ikaiguang/go-srv-kit/api/conf/v1"
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

// Bootstrap 配置引导
type Bootstrap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// app application
	App *App `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
	// log 日志
	Log *v1.Log `protobuf:"bytes,2,opt,name=log,proto3" json:"log,omitempty"`
	// data 数据
	Data *v1.Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	// server 服务
	Server *v1.Server `protobuf:"bytes,4,opt,name=server,proto3" json:"server,omitempty"`
	// base 基础配置
	Base *v1.Base `protobuf:"bytes,5,opt,name=base,proto3" json:"base,omitempty"`
	// Business 业务配置
	Business *v1.Business `protobuf:"bytes,6,opt,name=business,proto3" json:"business,omitempty"`
}

func (x *Bootstrap) Reset() {
	*x = Bootstrap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_common_v1_config_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bootstrap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bootstrap) ProtoMessage() {}

func (x *Bootstrap) ProtoReflect() protoreflect.Message {
	mi := &file_api_common_v1_config_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bootstrap.ProtoReflect.Descriptor instead.
func (*Bootstrap) Descriptor() ([]byte, []int) {
	return file_api_common_v1_config_v1_proto_rawDescGZIP(), []int{0}
}

func (x *Bootstrap) GetApp() *App {
	if x != nil {
		return x.App
	}
	return nil
}

func (x *Bootstrap) GetLog() *v1.Log {
	if x != nil {
		return x.Log
	}
	return nil
}

func (x *Bootstrap) GetData() *v1.Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Bootstrap) GetServer() *v1.Server {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *Bootstrap) GetBase() *v1.Base {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *Bootstrap) GetBusiness() *v1.Business {
	if x != nil {
		return x.Business
	}
	return nil
}

// App application
type App struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// belong_to 属于哪个项目
	BelongTo string `protobuf:"bytes,1,opt,name=belong_to,json=belongTo,proto3" json:"belong_to,omitempty"`
	// name app名字
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// version app版本
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// env app 环境
	Env string `protobuf:"bytes,4,opt,name=env,proto3" json:"env,omitempty"`
	// env_branch 环境分支
	EnvBranch string `protobuf:"bytes,5,opt,name=env_branch,json=envBranch,proto3" json:"env_branch,omitempty"`
	// endpoints app站点
	Endpoints []string `protobuf:"bytes,6,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	// metadata 元数据
	Metadata map[string]string `protobuf:"bytes,7,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *App) Reset() {
	*x = App{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_common_v1_config_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *App) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*App) ProtoMessage() {}

func (x *App) ProtoReflect() protoreflect.Message {
	mi := &file_api_common_v1_config_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use App.ProtoReflect.Descriptor instead.
func (*App) Descriptor() ([]byte, []int) {
	return file_api_common_v1_config_v1_proto_rawDescGZIP(), []int{1}
}

func (x *App) GetBelongTo() string {
	if x != nil {
		return x.BelongTo
	}
	return ""
}

func (x *App) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *App) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *App) GetEnv() string {
	if x != nil {
		return x.Env
	}
	return ""
}

func (x *App) GetEnvBranch() string {
	if x != nil {
		return x.EnvBranch
	}
	return ""
}

func (x *App) GetEndpoints() []string {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

func (x *App) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_api_common_v1_config_v1_proto protoreflect.FileDescriptor

var file_api_common_v1_config_v1_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1e, 0x73, 0x6e, 0x6f, 0x77, 0x66, 0x6c, 0x61, 0x6b, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x76, 0x31, 0x1a,
	0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6b, 0x61, 0x69,
	0x67, 0x75, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x72, 0x76, 0x2d, 0x6b, 0x69, 0x74,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x02, 0x0a, 0x09, 0x42,
	0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x12, 0x35, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x6e, 0x6f, 0x77, 0x66, 0x6c, 0x61, 0x6b,
	0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x03, 0x61, 0x70, 0x70, 0x12,
	0x25, 0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6b,
	0x69, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x76, 0x31, 0x2e, 0x4c, 0x6f,
	0x67, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6b, 0x69, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x2e, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x6b, 0x69, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x12, 0x28, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x6b, 0x69, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x76, 0x31, 0x2e,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x62, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6b,
	0x69, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x76, 0x31, 0x2e, 0x42, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x08, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x22, 0xab, 0x02, 0x0a, 0x03, 0x41, 0x70, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x65, 0x6c, 0x6f,
	0x6e, 0x67, 0x5f, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x65, 0x6c,
	0x6f, 0x6e, 0x67, 0x54, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x65, 0x6e, 0x76, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6e, 0x76, 0x5f, 0x62, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6e, 0x76, 0x42, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x12, 0x4d, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x73, 0x6e, 0x6f, 0x77, 0x66, 0x6c, 0x61, 0x6b, 0x65,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x81,
	0x01, 0x0a, 0x1e, 0x73, 0x6e, 0x6f, 0x77, 0x66, 0x6c, 0x61, 0x6b, 0x65, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x76,
	0x31, 0x42, 0x1b, 0x53, 0x6e, 0x6f, 0x77, 0x66, 0x6c, 0x61, 0x6b, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x41, 0x70, 0x69, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x56, 0x31, 0x50, 0x01,
	0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6b, 0x61,
	0x69, 0x67, 0x75, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x6e, 0x6f, 0x77, 0x66, 0x6c,
	0x61, 0x6b, 0x65, 0x2d, 0x6e, 0x6f, 0x64, 0x65, 0x2d, 0x69, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_common_v1_config_v1_proto_rawDescOnce sync.Once
	file_api_common_v1_config_v1_proto_rawDescData = file_api_common_v1_config_v1_proto_rawDesc
)

func file_api_common_v1_config_v1_proto_rawDescGZIP() []byte {
	file_api_common_v1_config_v1_proto_rawDescOnce.Do(func() {
		file_api_common_v1_config_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_common_v1_config_v1_proto_rawDescData)
	})
	return file_api_common_v1_config_v1_proto_rawDescData
}

var file_api_common_v1_config_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_common_v1_config_v1_proto_goTypes = []interface{}{
	(*Bootstrap)(nil),   // 0: snowflake.service.api.commonv1.Bootstrap
	(*App)(nil),         // 1: snowflake.service.api.commonv1.App
	nil,                 // 2: snowflake.service.api.commonv1.App.MetadataEntry
	(*v1.Log)(nil),      // 3: kit.api.confv1.Log
	(*v1.Data)(nil),     // 4: kit.api.confv1.Data
	(*v1.Server)(nil),   // 5: kit.api.confv1.Server
	(*v1.Base)(nil),     // 6: kit.api.confv1.Base
	(*v1.Business)(nil), // 7: kit.api.confv1.Business
}
var file_api_common_v1_config_v1_proto_depIdxs = []int32{
	1, // 0: snowflake.service.api.commonv1.Bootstrap.app:type_name -> snowflake.service.api.commonv1.App
	3, // 1: snowflake.service.api.commonv1.Bootstrap.log:type_name -> kit.api.confv1.Log
	4, // 2: snowflake.service.api.commonv1.Bootstrap.data:type_name -> kit.api.confv1.Data
	5, // 3: snowflake.service.api.commonv1.Bootstrap.server:type_name -> kit.api.confv1.Server
	6, // 4: snowflake.service.api.commonv1.Bootstrap.base:type_name -> kit.api.confv1.Base
	7, // 5: snowflake.service.api.commonv1.Bootstrap.business:type_name -> kit.api.confv1.Business
	2, // 6: snowflake.service.api.commonv1.App.metadata:type_name -> snowflake.service.api.commonv1.App.MetadataEntry
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_api_common_v1_config_v1_proto_init() }
func file_api_common_v1_config_v1_proto_init() {
	if File_api_common_v1_config_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_common_v1_config_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bootstrap); i {
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
		file_api_common_v1_config_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*App); i {
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
			RawDescriptor: file_api_common_v1_config_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_common_v1_config_v1_proto_goTypes,
		DependencyIndexes: file_api_common_v1_config_v1_proto_depIdxs,
		MessageInfos:      file_api_common_v1_config_v1_proto_msgTypes,
	}.Build()
	File_api_common_v1_config_v1_proto = out.File
	file_api_common_v1_config_v1_proto_rawDesc = nil
	file_api_common_v1_config_v1_proto_goTypes = nil
	file_api_common_v1_config_v1_proto_depIdxs = nil
}
