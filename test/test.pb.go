// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1-devel
// 	protoc        v4.22.2
// source: test/test.proto

package test

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TestItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestField string `protobuf:"bytes,1,opt,name=testField,proto3" json:"testField,omitempty"`
}

func (x *TestItem) Reset() {
	*x = TestItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestItem) ProtoMessage() {}

func (x *TestItem) ProtoReflect() protoreflect.Message {
	mi := &file_test_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestItem.ProtoReflect.Descriptor instead.
func (*TestItem) Descriptor() ([]byte, []int) {
	return file_test_test_proto_rawDescGZIP(), []int{0}
}

func (x *TestItem) GetTestField() string {
	if x != nil {
		return x.TestField
	}
	return ""
}

type Test struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FieldA   string      `protobuf:"bytes,1,opt,name=fieldA,proto3" json:"fieldA,omitempty"`
	TestItem []*TestItem `protobuf:"bytes,2,rep,name=testItem,proto3" json:"testItem,omitempty"`
}

func (x *Test) Reset() {
	*x = Test{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test) ProtoMessage() {}

func (x *Test) ProtoReflect() protoreflect.Message {
	mi := &file_test_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test.ProtoReflect.Descriptor instead.
func (*Test) Descriptor() ([]byte, []int) {
	return file_test_test_proto_rawDescGZIP(), []int{1}
}

func (x *Test) GetFieldA() string {
	if x != nil {
		return x.FieldA
	}
	return ""
}

func (x *Test) GetTestItem() []*TestItem {
	if x != nil {
		return x.TestItem
	}
	return nil
}

var file_test_test_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         10000,
		Name:          "validate",
		Tag:           "bytes,10000,opt,name=validate",
		Filename:      "test/test.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional string validate = 10000;
	E_Validate = &file_test_test_proto_extTypes[0]
)

var File_test_test_proto protoreflect.FileDescriptor

var file_test_test_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x08, 0x54, 0x65, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x41, 0x0a, 0x09, 0x74, 0x65, 0x73, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x23, 0x82, 0xf1, 0x04, 0x1f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
	0x7c, 0x6d, 0x69, 0x6e, 0x5f, 0x6c, 0x65, 0x6e, 0x3a, 0x31, 0x30, 0x7c, 0x6d, 0x61, 0x78, 0x5f,
	0x6c, 0x65, 0x6e, 0x3a, 0x31, 0x30, 0x30, 0x52, 0x09, 0x74, 0x65, 0x73, 0x74, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x22, 0x78, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x06, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x41, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x23, 0x82, 0xf1, 0x04, 0x1f,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x7c, 0x6d, 0x69, 0x6e, 0x5f, 0x6c, 0x65, 0x6e,
	0x3a, 0x31, 0x30, 0x7c, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x65, 0x6e, 0x3a, 0x31, 0x30, 0x30, 0x52,
	0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x41, 0x12, 0x33, 0x0a, 0x08, 0x74, 0x65, 0x73, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x42, 0x0c, 0x82, 0xf1, 0x04, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x52, 0x08, 0x74, 0x65, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x3a, 0x3a, 0x0a, 0x08,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x90, 0x4e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x74, 0x65, 0x73, 0x74,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_test_proto_rawDescOnce sync.Once
	file_test_test_proto_rawDescData = file_test_test_proto_rawDesc
)

func file_test_test_proto_rawDescGZIP() []byte {
	file_test_test_proto_rawDescOnce.Do(func() {
		file_test_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_test_proto_rawDescData)
	})
	return file_test_test_proto_rawDescData
}

var file_test_test_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_test_test_proto_goTypes = []interface{}{
	(*TestItem)(nil),                  // 0: TestItem
	(*Test)(nil),                      // 1: Test
	(*descriptorpb.FieldOptions)(nil), // 2: google.protobuf.FieldOptions
}
var file_test_test_proto_depIdxs = []int32{
	0, // 0: Test.testItem:type_name -> TestItem
	2, // 1: validate:extendee -> google.protobuf.FieldOptions
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	1, // [1:2] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_test_test_proto_init() }
func file_test_test_proto_init() {
	if File_test_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestItem); i {
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
		file_test_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test); i {
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
			RawDescriptor: file_test_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_test_test_proto_goTypes,
		DependencyIndexes: file_test_test_proto_depIdxs,
		MessageInfos:      file_test_test_proto_msgTypes,
		ExtensionInfos:    file_test_test_proto_extTypes,
	}.Build()
	File_test_test_proto = out.File
	file_test_test_proto_rawDesc = nil
	file_test_test_proto_goTypes = nil
	file_test_test_proto_depIdxs = nil
}
