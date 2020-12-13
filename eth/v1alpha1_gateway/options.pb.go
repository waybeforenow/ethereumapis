// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: eth/v1alpha1/options.proto

package eth

import (
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var file_eth_v1alpha1_options_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50000,
		Name:          "ethereum.eth.v1alpha1.cast_type",
		Tag:           "bytes,50000,opt,name=cast_type",
		Filename:      "eth/v1alpha1/options.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50001,
		Name:          "ethereum.eth.v1alpha1.ssz_size",
		Tag:           "bytes,50001,opt,name=ssz_size",
		Filename:      "eth/v1alpha1/options.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50002,
		Name:          "ethereum.eth.v1alpha1.ssz_max",
		Tag:           "bytes,50002,opt,name=ssz_max",
		Filename:      "eth/v1alpha1/options.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50003,
		Name:          "ethereum.eth.v1alpha1.spec_name",
		Tag:           "bytes,50003,opt,name=spec_name",
		Filename:      "eth/v1alpha1/options.proto",
	},
}

// Extension fields to descriptor.FieldOptions.
var (
	// optional string cast_type = 50000;
	E_CastType = &file_eth_v1alpha1_options_proto_extTypes[0]
	// optional string ssz_size = 50001;
	E_SszSize = &file_eth_v1alpha1_options_proto_extTypes[1]
	// optional string ssz_max = 50002;
	E_SszMax = &file_eth_v1alpha1_options_proto_extTypes[2]
	// optional string spec_name = 50003;
	E_SpecName = &file_eth_v1alpha1_options_proto_extTypes[3]
)

var File_eth_v1alpha1_options_proto protoreflect.FileDescriptor

var file_eth_v1alpha1_options_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x65, 0x74,
	0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x3c, 0x0a, 0x09, 0x63, 0x61, 0x73, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xd0, 0x86, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x73, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x3a, 0x3a, 0x0a, 0x08, 0x73, 0x73, 0x7a, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x12,
	0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd1,
	0x86, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x73, 0x7a, 0x53, 0x69, 0x7a, 0x65, 0x3a,
	0x38, 0x0a, 0x07, 0x73, 0x73, 0x7a, 0x5f, 0x6d, 0x61, 0x78, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd2, 0x86, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x73, 0x7a, 0x4d, 0x61, 0x78, 0x3a, 0x3c, 0x0a, 0x09, 0x73, 0x70, 0x65,
	0x63, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd3, 0x86, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x70, 0x65, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x95, 0x01, 0x0a, 0x19, 0x6f, 0x72, 0x67, 0x2e,
	0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62,
	0x73, 0x2f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x65,
	0x74, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x65, 0x74, 0x68, 0xaa,
	0x02, 0x15, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x45, 0x74, 0x68, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x15, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65,
	0x75, 0x6d, 0x5c, 0x45, 0x74, 0x68, 0x5c, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_eth_v1alpha1_options_proto_goTypes = []interface{}{
	(*descriptor.FieldOptions)(nil), // 0: google.protobuf.FieldOptions
}
var file_eth_v1alpha1_options_proto_depIdxs = []int32{
	0, // 0: ethereum.eth.v1alpha1.cast_type:extendee -> google.protobuf.FieldOptions
	0, // 1: ethereum.eth.v1alpha1.ssz_size:extendee -> google.protobuf.FieldOptions
	0, // 2: ethereum.eth.v1alpha1.ssz_max:extendee -> google.protobuf.FieldOptions
	0, // 3: ethereum.eth.v1alpha1.spec_name:extendee -> google.protobuf.FieldOptions
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	0, // [0:4] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eth_v1alpha1_options_proto_init() }
func file_eth_v1alpha1_options_proto_init() {
	if File_eth_v1alpha1_options_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eth_v1alpha1_options_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 4,
			NumServices:   0,
		},
		GoTypes:           file_eth_v1alpha1_options_proto_goTypes,
		DependencyIndexes: file_eth_v1alpha1_options_proto_depIdxs,
		ExtensionInfos:    file_eth_v1alpha1_options_proto_extTypes,
	}.Build()
	File_eth_v1alpha1_options_proto = out.File
	file_eth_v1alpha1_options_proto_rawDesc = nil
	file_eth_v1alpha1_options_proto_goTypes = nil
	file_eth_v1alpha1_options_proto_depIdxs = nil
}
