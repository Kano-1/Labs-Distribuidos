// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.27.0--rc1
// source: Proto/municion.proto

package Proto

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

type SolicitarMunicion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	AT int32 `protobuf:"varint,2,opt,name=AT,proto3" json:"AT,omitempty"`
	MP int32 `protobuf:"varint,3,opt,name=MP,proto3" json:"MP,omitempty"`
}

func (x *SolicitarMunicion) Reset() {
	*x = SolicitarMunicion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_municion_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SolicitarMunicion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolicitarMunicion) ProtoMessage() {}

func (x *SolicitarMunicion) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_municion_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolicitarMunicion.ProtoReflect.Descriptor instead.
func (*SolicitarMunicion) Descriptor() ([]byte, []int) {
	return file_Proto_municion_proto_rawDescGZIP(), []int{0}
}

func (x *SolicitarMunicion) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *SolicitarMunicion) GetAT() int32 {
	if x != nil {
		return x.AT
	}
	return 0
}

func (x *SolicitarMunicion) GetMP() int32 {
	if x != nil {
		return x.MP
	}
	return 0
}

type RespuestaMunicion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RespuestaMunicion) Reset() {
	*x = RespuestaMunicion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_municion_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespuestaMunicion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespuestaMunicion) ProtoMessage() {}

func (x *RespuestaMunicion) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_municion_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespuestaMunicion.ProtoReflect.Descriptor instead.
func (*RespuestaMunicion) Descriptor() ([]byte, []int) {
	return file_Proto_municion_proto_rawDescGZIP(), []int{1}
}

func (x *RespuestaMunicion) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_Proto_municion_proto protoreflect.FileDescriptor

var file_Proto_municion_proto_rawDesc = []byte{
	0x0a, 0x14, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a,
	0x11, 0x73, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x61, 0x72, 0x4d, 0x75, 0x6e, 0x69, 0x63, 0x69,
	0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x41, 0x54, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x41, 0x54, 0x12, 0x0e, 0x0a, 0x02, 0x4d, 0x50, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x4d, 0x50, 0x22, 0x2d, 0x0a, 0x11, 0x72, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x4d,
	0x75, 0x6e, 0x69, 0x63, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x32, 0x54, 0x0a, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x69, 0x6f, 0x4d, 0x75, 0x6e,
	0x69, 0x63, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x0a, 0x53, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74,
	0x61, 0x72, 0x4d, 0x12, 0x18, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x69,
	0x63, 0x69, 0x74, 0x61, 0x72, 0x4d, 0x75, 0x6e, 0x69, 0x63, 0x69, 0x6f, 0x6e, 0x1a, 0x18, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x4d,
	0x75, 0x6e, 0x69, 0x63, 0x69, 0x6f, 0x6e, 0x42, 0x0c, 0x5a, 0x0a, 0x4c, 0x61, 0x62, 0x33, 0x2f,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Proto_municion_proto_rawDescOnce sync.Once
	file_Proto_municion_proto_rawDescData = file_Proto_municion_proto_rawDesc
)

func file_Proto_municion_proto_rawDescGZIP() []byte {
	file_Proto_municion_proto_rawDescOnce.Do(func() {
		file_Proto_municion_proto_rawDescData = protoimpl.X.CompressGZIP(file_Proto_municion_proto_rawDescData)
	})
	return file_Proto_municion_proto_rawDescData
}

var file_Proto_municion_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Proto_municion_proto_goTypes = []interface{}{
	(*SolicitarMunicion)(nil), // 0: Proto.solicitarMunicion
	(*RespuestaMunicion)(nil), // 1: Proto.respuestaMunicion
}
var file_Proto_municion_proto_depIdxs = []int32{
	0, // 0: Proto.servicioMunicion.SolicitarM:input_type -> Proto.solicitarMunicion
	1, // 1: Proto.servicioMunicion.SolicitarM:output_type -> Proto.respuestaMunicion
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Proto_municion_proto_init() }
func file_Proto_municion_proto_init() {
	if File_Proto_municion_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Proto_municion_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SolicitarMunicion); i {
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
		file_Proto_municion_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespuestaMunicion); i {
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
			RawDescriptor: file_Proto_municion_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Proto_municion_proto_goTypes,
		DependencyIndexes: file_Proto_municion_proto_depIdxs,
		MessageInfos:      file_Proto_municion_proto_msgTypes,
	}.Build()
	File_Proto_municion_proto = out.File
	file_Proto_municion_proto_rawDesc = nil
	file_Proto_municion_proto_goTypes = nil
	file_Proto_municion_proto_depIdxs = nil
}