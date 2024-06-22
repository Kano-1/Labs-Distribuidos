// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0--rc1
// source: Proto/communication.proto

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_communication_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_communication_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_Proto_communication_proto_rawDescGZIP(), []int{0}
}

type ServerAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address int32 `protobuf:"varint,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *ServerAddress) Reset() {
	*x = ServerAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_communication_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerAddress) ProtoMessage() {}

func (x *ServerAddress) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_communication_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerAddress.ProtoReflect.Descriptor instead.
func (*ServerAddress) Descriptor() ([]byte, []int) {
	return file_Proto_communication_proto_rawDescGZIP(), []int{1}
}

func (x *ServerAddress) GetAddress() int32 {
	if x != nil {
		return x.Address
	}
	return 0
}

type ActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action string `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	Sector string `protobuf:"bytes,2,opt,name=sector,proto3" json:"sector,omitempty"`
	Base   string `protobuf:"bytes,3,opt,name=base,proto3" json:"base,omitempty"`
	Value  string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ActionRequest) Reset() {
	*x = ActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_communication_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionRequest) ProtoMessage() {}

func (x *ActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_communication_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionRequest.ProtoReflect.Descriptor instead.
func (*ActionRequest) Descriptor() ([]byte, []int) {
	return file_Proto_communication_proto_rawDescGZIP(), []int{2}
}

func (x *ActionRequest) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *ActionRequest) GetSector() string {
	if x != nil {
		return x.Sector
	}
	return ""
}

func (x *ActionRequest) GetBase() string {
	if x != nil {
		return x.Base
	}
	return ""
}

func (x *ActionRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type ClockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clock []int32 `protobuf:"varint,1,rep,packed,name=clock,proto3" json:"clock,omitempty"`
}

func (x *ClockResponse) Reset() {
	*x = ClockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_communication_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClockResponse) ProtoMessage() {}

func (x *ClockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_communication_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClockResponse.ProtoReflect.Descriptor instead.
func (*ClockResponse) Descriptor() ([]byte, []int) {
	return file_Proto_communication_proto_rawDescGZIP(), []int{3}
}

func (x *ClockResponse) GetClock() []int32 {
	if x != nil {
		return x.Clock
	}
	return nil
}

type ReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sector string `protobuf:"bytes,1,opt,name=sector,proto3" json:"sector,omitempty"`
	Base   string `protobuf:"bytes,2,opt,name=base,proto3" json:"base,omitempty"`
}

func (x *ReadRequest) Reset() {
	*x = ReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_communication_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest) ProtoMessage() {}

func (x *ReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_communication_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRequest.ProtoReflect.Descriptor instead.
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return file_Proto_communication_proto_rawDescGZIP(), []int{4}
}

func (x *ReadRequest) GetSector() string {
	if x != nil {
		return x.Sector
	}
	return ""
}

func (x *ReadRequest) GetBase() string {
	if x != nil {
		return x.Base
	}
	return ""
}

type EnemiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enemies int32   `protobuf:"varint,1,opt,name=enemies,proto3" json:"enemies,omitempty"`
	Clock   []int32 `protobuf:"varint,2,rep,packed,name=clock,proto3" json:"clock,omitempty"`
}

func (x *EnemiesResponse) Reset() {
	*x = EnemiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_communication_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnemiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnemiesResponse) ProtoMessage() {}

func (x *EnemiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_communication_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnemiesResponse.ProtoReflect.Descriptor instead.
func (*EnemiesResponse) Descriptor() ([]byte, []int) {
	return file_Proto_communication_proto_rawDescGZIP(), []int{5}
}

func (x *EnemiesResponse) GetEnemies() int32 {
	if x != nil {
		return x.Enemies
	}
	return 0
}

func (x *EnemiesResponse) GetClock() []int32 {
	if x != nil {
		return x.Clock
	}
	return nil
}

type VectorClocks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sector string  `protobuf:"bytes,1,opt,name=sector,proto3" json:"sector,omitempty"`
	Clock  []int32 `protobuf:"varint,2,rep,packed,name=clock,proto3" json:"clock,omitempty"`
}

func (x *VectorClocks) Reset() {
	*x = VectorClocks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_communication_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VectorClocks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VectorClocks) ProtoMessage() {}

func (x *VectorClocks) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_communication_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VectorClocks.ProtoReflect.Descriptor instead.
func (*VectorClocks) Descriptor() ([]byte, []int) {
	return file_Proto_communication_proto_rawDescGZIP(), []int{6}
}

func (x *VectorClocks) GetSector() string {
	if x != nil {
		return x.Sector
	}
	return ""
}

func (x *VectorClocks) GetClock() []int32 {
	if x != nil {
		return x.Clock
	}
	return nil
}

type PropagationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Log    []string        `protobuf:"bytes,1,rep,name=log,proto3" json:"log,omitempty"`
	Clocks []*VectorClocks `protobuf:"bytes,2,rep,name=clocks,proto3" json:"clocks,omitempty"`
}

func (x *PropagationRequest) Reset() {
	*x = PropagationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_communication_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PropagationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropagationRequest) ProtoMessage() {}

func (x *PropagationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_communication_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropagationRequest.ProtoReflect.Descriptor instead.
func (*PropagationRequest) Descriptor() ([]byte, []int) {
	return file_Proto_communication_proto_rawDescGZIP(), []int{7}
}

func (x *PropagationRequest) GetLog() []string {
	if x != nil {
		return x.Log
	}
	return nil
}

func (x *PropagationRequest) GetClocks() []*VectorClocks {
	if x != nil {
		return x.Clocks
	}
	return nil
}

type PropagationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *PropagationResponse) Reset() {
	*x = PropagationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_communication_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PropagationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropagationResponse) ProtoMessage() {}

func (x *PropagationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_communication_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropagationResponse.ProtoReflect.Descriptor instead.
func (*PropagationResponse) Descriptor() ([]byte, []int) {
	return file_Proto_communication_proto_rawDescGZIP(), []int{8}
}

func (x *PropagationResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_Proto_communication_proto protoreflect.FileDescriptor

var file_Proto_communication_proto_rawDesc = []byte{
	0x0a, 0x19, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x29, 0x0a, 0x0d, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x69, 0x0a, 0x0d, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x25, 0x0a, 0x0d, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x05, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x39, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62,
	0x61, 0x73, 0x65, 0x22, 0x41, 0x0a, 0x0f, 0x45, 0x6e, 0x65, 0x6d, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x65, 0x6d, 0x69, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x6e, 0x65, 0x6d, 0x69, 0x65, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x3c, 0x0a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05, 0x63,
	0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x53, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x70, 0x61, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f,
	0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x12, 0x2b, 0x0a, 0x06,
	0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b,
	0x73, 0x52, 0x06, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x22, 0x2f, 0x0a, 0x13, 0x50, 0x72, 0x6f,
	0x70, 0x61, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x40, 0x0a, 0x06, 0x42, 0x72,
	0x6f, 0x6b, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x0c, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x32, 0xc5, 0x01, 0x0a,
	0x07, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x12, 0x37, 0x0a, 0x09, 0x57, 0x72, 0x69, 0x74,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x36, 0x0a, 0x08, 0x52, 0x65, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6e, 0x65, 0x6d, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x10, 0x50, 0x72, 0x6f,
	0x70, 0x61, 0x67, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x12, 0x19, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x61, 0x67, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x72, 0x6f, 0x70, 0x61, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x4c, 0x61, 0x62, 0x35, 0x2f, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Proto_communication_proto_rawDescOnce sync.Once
	file_Proto_communication_proto_rawDescData = file_Proto_communication_proto_rawDesc
)

func file_Proto_communication_proto_rawDescGZIP() []byte {
	file_Proto_communication_proto_rawDescOnce.Do(func() {
		file_Proto_communication_proto_rawDescData = protoimpl.X.CompressGZIP(file_Proto_communication_proto_rawDescData)
	})
	return file_Proto_communication_proto_rawDescData
}

var file_Proto_communication_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_Proto_communication_proto_goTypes = []interface{}{
	(*Empty)(nil),               // 0: Proto.Empty
	(*ServerAddress)(nil),       // 1: Proto.ServerAddress
	(*ActionRequest)(nil),       // 2: Proto.ActionRequest
	(*ClockResponse)(nil),       // 3: Proto.ClockResponse
	(*ReadRequest)(nil),         // 4: Proto.ReadRequest
	(*EnemiesResponse)(nil),     // 5: Proto.EnemiesResponse
	(*VectorClocks)(nil),        // 6: Proto.VectorClocks
	(*PropagationRequest)(nil),  // 7: Proto.PropagationRequest
	(*PropagationResponse)(nil), // 8: Proto.PropagationResponse
}
var file_Proto_communication_proto_depIdxs = []int32{
	6, // 0: Proto.PropagationRequest.clocks:type_name -> Proto.VectorClocks
	0, // 1: Proto.Broker.GetServerAddress:input_type -> Proto.Empty
	2, // 2: Proto.Servers.WriteInfo:input_type -> Proto.ActionRequest
	4, // 3: Proto.Servers.ReadInfo:input_type -> Proto.ReadRequest
	7, // 4: Proto.Servers.PropagateChanges:input_type -> Proto.PropagationRequest
	1, // 5: Proto.Broker.GetServerAddress:output_type -> Proto.ServerAddress
	3, // 6: Proto.Servers.WriteInfo:output_type -> Proto.ClockResponse
	5, // 7: Proto.Servers.ReadInfo:output_type -> Proto.EnemiesResponse
	8, // 8: Proto.Servers.PropagateChanges:output_type -> Proto.PropagationResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Proto_communication_proto_init() }
func file_Proto_communication_proto_init() {
	if File_Proto_communication_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Proto_communication_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_Proto_communication_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerAddress); i {
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
		file_Proto_communication_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionRequest); i {
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
		file_Proto_communication_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClockResponse); i {
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
		file_Proto_communication_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRequest); i {
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
		file_Proto_communication_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnemiesResponse); i {
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
		file_Proto_communication_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VectorClocks); i {
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
		file_Proto_communication_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropagationRequest); i {
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
		file_Proto_communication_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropagationResponse); i {
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
			RawDescriptor: file_Proto_communication_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_Proto_communication_proto_goTypes,
		DependencyIndexes: file_Proto_communication_proto_depIdxs,
		MessageInfos:      file_Proto_communication_proto_msgTypes,
	}.Build()
	File_Proto_communication_proto = out.File
	file_Proto_communication_proto_rawDesc = nil
	file_Proto_communication_proto_goTypes = nil
	file_Proto_communication_proto_depIdxs = nil
}
