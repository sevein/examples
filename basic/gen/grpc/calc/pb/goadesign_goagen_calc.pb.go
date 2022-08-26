// Code generated with goa v3.8.4, DO NOT EDIT.
//
// calc protocol buffer definition
//
// Command:
// $ goa gen goa.design/examples/basic/design -o basic

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.2
// source: goadesign_goagen_calc.proto

package calcpb

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

type MultiplyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Left operand
	A int32 `protobuf:"zigzag32,1,opt,name=a,proto3" json:"a,omitempty"`
	// Right operand
	B int32 `protobuf:"zigzag32,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *MultiplyRequest) Reset() {
	*x = MultiplyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_calc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiplyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiplyRequest) ProtoMessage() {}

func (x *MultiplyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_calc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiplyRequest.ProtoReflect.Descriptor instead.
func (*MultiplyRequest) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_calc_proto_rawDescGZIP(), []int{0}
}

func (x *MultiplyRequest) GetA() int32 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *MultiplyRequest) GetB() int32 {
	if x != nil {
		return x.B
	}
	return 0
}

type MultiplyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field int32 `protobuf:"zigzag32,1,opt,name=field,proto3" json:"field,omitempty"`
}

func (x *MultiplyResponse) Reset() {
	*x = MultiplyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_calc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiplyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiplyResponse) ProtoMessage() {}

func (x *MultiplyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_calc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiplyResponse.ProtoReflect.Descriptor instead.
func (*MultiplyResponse) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_calc_proto_rawDescGZIP(), []int{1}
}

func (x *MultiplyResponse) GetField() int32 {
	if x != nil {
		return x.Field
	}
	return 0
}

var File_goadesign_goagen_calc_proto protoreflect.FileDescriptor

var file_goadesign_goagen_calc_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x67, 0x6f, 0x61, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x67, 0x6f, 0x61, 0x67,
	0x65, 0x6e, 0x5f, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63,
	0x61, 0x6c, 0x63, 0x22, 0x2d, 0x0a, 0x0f, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x11, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x11, 0x52,
	0x01, 0x62, 0x22, 0x28, 0x0a, 0x10, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x11, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x41, 0x0a, 0x04,
	0x43, 0x61, 0x6c, 0x63, 0x12, 0x39, 0x0a, 0x08, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x79,
	0x12, 0x15, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x09, 0x5a, 0x07, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_goadesign_goagen_calc_proto_rawDescOnce sync.Once
	file_goadesign_goagen_calc_proto_rawDescData = file_goadesign_goagen_calc_proto_rawDesc
)

func file_goadesign_goagen_calc_proto_rawDescGZIP() []byte {
	file_goadesign_goagen_calc_proto_rawDescOnce.Do(func() {
		file_goadesign_goagen_calc_proto_rawDescData = protoimpl.X.CompressGZIP(file_goadesign_goagen_calc_proto_rawDescData)
	})
	return file_goadesign_goagen_calc_proto_rawDescData
}

var file_goadesign_goagen_calc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_goadesign_goagen_calc_proto_goTypes = []interface{}{
	(*MultiplyRequest)(nil),  // 0: calc.MultiplyRequest
	(*MultiplyResponse)(nil), // 1: calc.MultiplyResponse
}
var file_goadesign_goagen_calc_proto_depIdxs = []int32{
	0, // 0: calc.Calc.Multiply:input_type -> calc.MultiplyRequest
	1, // 1: calc.Calc.Multiply:output_type -> calc.MultiplyResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_goadesign_goagen_calc_proto_init() }
func file_goadesign_goagen_calc_proto_init() {
	if File_goadesign_goagen_calc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_goadesign_goagen_calc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiplyRequest); i {
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
		file_goadesign_goagen_calc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiplyResponse); i {
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
			RawDescriptor: file_goadesign_goagen_calc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goadesign_goagen_calc_proto_goTypes,
		DependencyIndexes: file_goadesign_goagen_calc_proto_depIdxs,
		MessageInfos:      file_goadesign_goagen_calc_proto_msgTypes,
	}.Build()
	File_goadesign_goagen_calc_proto = out.File
	file_goadesign_goagen_calc_proto_rawDesc = nil
	file_goadesign_goagen_calc_proto_goTypes = nil
	file_goadesign_goagen_calc_proto_depIdxs = nil
}
