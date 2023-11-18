// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: division.proto

package division

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

type DivisionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num1 float64 `protobuf:"fixed64,1,opt,name=num1,proto3" json:"num1,omitempty"`
	Num2 float64 `protobuf:"fixed64,2,opt,name=num2,proto3" json:"num2,omitempty"`
}

func (x *DivisionRequest) Reset() {
	*x = DivisionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_division_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DivisionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DivisionRequest) ProtoMessage() {}

func (x *DivisionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_division_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DivisionRequest.ProtoReflect.Descriptor instead.
func (*DivisionRequest) Descriptor() ([]byte, []int) {
	return file_division_proto_rawDescGZIP(), []int{0}
}

func (x *DivisionRequest) GetNum1() float64 {
	if x != nil {
		return x.Num1
	}
	return 0
}

func (x *DivisionRequest) GetNum2() float64 {
	if x != nil {
		return x.Num2
	}
	return 0
}

type DivisionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result float64 `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *DivisionResponse) Reset() {
	*x = DivisionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_division_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DivisionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DivisionResponse) ProtoMessage() {}

func (x *DivisionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_division_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DivisionResponse.ProtoReflect.Descriptor instead.
func (*DivisionResponse) Descriptor() ([]byte, []int) {
	return file_division_proto_rawDescGZIP(), []int{1}
}

func (x *DivisionResponse) GetResult() float64 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_division_proto protoreflect.FileDescriptor

var file_division_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x64, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x64, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x39, 0x0a, 0x0f, 0x44, 0x69,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x75, 0x6d, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x6e, 0x75, 0x6d,
	0x31, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x75, 0x6d, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x04, 0x6e, 0x75, 0x6d, 0x32, 0x22, 0x2a, 0x0a, 0x10, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x32, 0x52, 0x0a, 0x0f, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x44, 0x69, 0x76, 0x69, 0x64, 0x65, 0x12, 0x19,
	0x2e, 0x64, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x64, 0x69, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x52, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x69, 0x72, 0x67, 0x65, 0x6e, 0x75, 0x69, 0x6e, 0x65, 0x2f,
	0x6f, 0x76, 0x65, 0x72, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x65, 0x72, 0x65, 0x64, 0x2d, 0x63,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x64, 0x69, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x3b, 0x64, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_division_proto_rawDescOnce sync.Once
	file_division_proto_rawDescData = file_division_proto_rawDesc
)

func file_division_proto_rawDescGZIP() []byte {
	file_division_proto_rawDescOnce.Do(func() {
		file_division_proto_rawDescData = protoimpl.X.CompressGZIP(file_division_proto_rawDescData)
	})
	return file_division_proto_rawDescData
}

var file_division_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_division_proto_goTypes = []interface{}{
	(*DivisionRequest)(nil),  // 0: division.DivisionRequest
	(*DivisionResponse)(nil), // 1: division.DivisionResponse
}
var file_division_proto_depIdxs = []int32{
	0, // 0: division.DivisionService.Divide:input_type -> division.DivisionRequest
	1, // 1: division.DivisionService.Divide:output_type -> division.DivisionResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_division_proto_init() }
func file_division_proto_init() {
	if File_division_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_division_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DivisionRequest); i {
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
		file_division_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DivisionResponse); i {
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
			RawDescriptor: file_division_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_division_proto_goTypes,
		DependencyIndexes: file_division_proto_depIdxs,
		MessageInfos:      file_division_proto_msgTypes,
	}.Build()
	File_division_proto = out.File
	file_division_proto_rawDesc = nil
	file_division_proto_goTypes = nil
	file_division_proto_depIdxs = nil
}