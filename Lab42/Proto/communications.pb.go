// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.12.4
// source: communications.proto

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

type FloorInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Floor int32 `protobuf:"varint,1,opt,name=floor,proto3" json:"floor,omitempty"`
}

func (x *FloorInfo) Reset() {
	*x = FloorInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communications_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FloorInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FloorInfo) ProtoMessage() {}

func (x *FloorInfo) ProtoReflect() protoreflect.Message {
	mi := &file_communications_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FloorInfo.ProtoReflect.Descriptor instead.
func (*FloorInfo) Descriptor() ([]byte, []int) {
	return file_communications_proto_rawDescGZIP(), []int{0}
}

func (x *FloorInfo) GetFloor() int32 {
	if x != nil {
		return x.Floor
	}
	return 0
}

type Move struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action       string `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	LeadsToDeath bool   `protobuf:"varint,2,opt,name=leads_to_death,json=leadsToDeath,proto3" json:"leads_to_death,omitempty"`
}

func (x *Move) Reset() {
	*x = Move{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communications_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Move) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Move) ProtoMessage() {}

func (x *Move) ProtoReflect() protoreflect.Message {
	mi := &file_communications_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Move.ProtoReflect.Descriptor instead.
func (*Move) Descriptor() ([]byte, []int) {
	return file_communications_proto_rawDescGZIP(), []int{1}
}

func (x *Move) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *Move) GetLeadsToDeath() bool {
	if x != nil {
		return x.LeadsToDeath
	}
	return false
}

type MercenaryInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *MercenaryInfo) Reset() {
	*x = MercenaryInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communications_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MercenaryInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MercenaryInfo) ProtoMessage() {}

func (x *MercenaryInfo) ProtoReflect() protoreflect.Message {
	mi := &file_communications_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MercenaryInfo.ProtoReflect.Descriptor instead.
func (*MercenaryInfo) Descriptor() ([]byte, []int) {
	return file_communications_proto_rawDescGZIP(), []int{2}
}

func (x *MercenaryInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MercenaryInfo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type DeathInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeathInfo) Reset() {
	*x = DeathInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communications_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeathInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeathInfo) ProtoMessage() {}

func (x *DeathInfo) ProtoReflect() protoreflect.Message {
	mi := &file_communications_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeathInfo.ProtoReflect.Descriptor instead.
func (*DeathInfo) Descriptor() ([]byte, []int) {
	return file_communications_proto_rawDescGZIP(), []int{3}
}

func (x *DeathInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type MoveInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Move *Move  `protobuf:"bytes,2,opt,name=move,proto3" json:"move,omitempty"`
}

func (x *MoveInfo) Reset() {
	*x = MoveInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communications_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoveInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveInfo) ProtoMessage() {}

func (x *MoveInfo) ProtoReflect() protoreflect.Message {
	mi := &file_communications_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveInfo.ProtoReflect.Descriptor instead.
func (*MoveInfo) Descriptor() ([]byte, []int) {
	return file_communications_proto_rawDescGZIP(), []int{4}
}

func (x *MoveInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MoveInfo) GetMove() *Move {
	if x != nil {
		return x.Move
	}
	return nil
}

type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mercenaries []*MercenaryStatus `protobuf:"bytes,1,rep,name=mercenaries,proto3" json:"mercenaries,omitempty"`
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communications_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_communications_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_communications_proto_rawDescGZIP(), []int{5}
}

func (x *StatusResponse) GetMercenaries() []*MercenaryStatus {
	if x != nil {
		return x.Mercenaries
	}
	return nil
}

type MercenaryStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Alive   bool   `protobuf:"varint,2,opt,name=alive,proto3" json:"alive,omitempty"`
	Ready   bool   `protobuf:"varint,3,opt,name=ready,proto3" json:"ready,omitempty"`
	Address string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *MercenaryStatus) Reset() {
	*x = MercenaryStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communications_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MercenaryStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MercenaryStatus) ProtoMessage() {}

func (x *MercenaryStatus) ProtoReflect() protoreflect.Message {
	mi := &file_communications_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MercenaryStatus.ProtoReflect.Descriptor instead.
func (*MercenaryStatus) Descriptor() ([]byte, []int) {
	return file_communications_proto_rawDescGZIP(), []int{6}
}

func (x *MercenaryStatus) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MercenaryStatus) GetAlive() bool {
	if x != nil {
		return x.Alive
	}
	return false
}

func (x *MercenaryStatus) GetReady() bool {
	if x != nil {
		return x.Ready
	}
	return false
}

func (x *MercenaryStatus) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type AccumulatedAmountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount int32 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *AccumulatedAmountResponse) Reset() {
	*x = AccumulatedAmountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communications_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccumulatedAmountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccumulatedAmountResponse) ProtoMessage() {}

func (x *AccumulatedAmountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_communications_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccumulatedAmountResponse.ProtoReflect.Descriptor instead.
func (*AccumulatedAmountResponse) Descriptor() ([]byte, []int) {
	return file_communications_proto_rawDescGZIP(), []int{7}
}

func (x *AccumulatedAmountResponse) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communications_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_communications_proto_msgTypes[8]
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
	return file_communications_proto_rawDescGZIP(), []int{8}
}

var File_communications_proto protoreflect.FileDescriptor

var file_communications_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a,
	0x09, 0x46, 0x6c, 0x6f, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c,
	0x6f, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x66, 0x6c, 0x6f, 0x6f, 0x72,
	0x22, 0x44, 0x0a, 0x04, 0x4d, 0x6f, 0x76, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x24, 0x0a, 0x0e, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x64, 0x65, 0x61,
	0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x54,
	0x6f, 0x44, 0x65, 0x61, 0x74, 0x68, 0x22, 0x3d, 0x0a, 0x0d, 0x4d, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x61, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x1f, 0x0a, 0x09, 0x44, 0x65, 0x61, 0x74, 0x68, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3f, 0x0a, 0x08, 0x4d, 0x6f, 0x76, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x76,
	0x65, 0x52, 0x04, 0x6d, 0x6f, 0x76, 0x65, 0x22, 0x4a, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x6d, 0x65, 0x72,
	0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x79,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0b, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x61, 0x72,
	0x69, 0x65, 0x73, 0x22, 0x6b, 0x0a, 0x0f, 0x4d, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x79,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c,
	0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x76, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x61, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x72, 0x65, 0x61, 0x64, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x22, 0x33, 0x0a, 0x19, 0x41, 0x63, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x62,
	0x0a, 0x09, 0x4d, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x2c, 0x0a, 0x0a, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x46, 0x6c, 0x6f, 0x6f, 0x72, 0x12, 0x10, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x46, 0x6c, 0x6f, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0c, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x27, 0x0a, 0x0a, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x4d, 0x6f, 0x76, 0x65, 0x12, 0x0b, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4d, 0x6f, 0x76, 0x65, 0x1a, 0x0c, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x32, 0x9a, 0x02, 0x0a, 0x08, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12,
	0x32, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x79, 0x12,
	0x14, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x61, 0x72,
	0x79, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0c, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x2d, 0x0a, 0x0b, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x65, 0x61,
	0x74, 0x68, 0x12, 0x10, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x61, 0x74, 0x68,
	0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0c, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x2b, 0x0a, 0x0a, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x6f, 0x76, 0x65,
	0x12, 0x0f, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x1a, 0x0c, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x32, 0x0a, 0x0b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0c,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x18, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x63,
	0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x0c, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x20, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0x52, 0x0a, 0x08, 0x44, 0x6f, 0x73, 0x68, 0x42, 0x61, 0x6e, 0x6b, 0x12, 0x46, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x41, 0x63, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x0c, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x20, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x75, 0x6d, 0x75,
	0x6c, 0x61, 0x74, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_communications_proto_rawDescOnce sync.Once
	file_communications_proto_rawDescData = file_communications_proto_rawDesc
)

func file_communications_proto_rawDescGZIP() []byte {
	file_communications_proto_rawDescOnce.Do(func() {
		file_communications_proto_rawDescData = protoimpl.X.CompressGZIP(file_communications_proto_rawDescData)
	})
	return file_communications_proto_rawDescData
}

var file_communications_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_communications_proto_goTypes = []interface{}{
	(*FloorInfo)(nil),                 // 0: Proto.FloorInfo
	(*Move)(nil),                      // 1: Proto.Move
	(*MercenaryInfo)(nil),             // 2: Proto.MercenaryInfo
	(*DeathInfo)(nil),                 // 3: Proto.DeathInfo
	(*MoveInfo)(nil),                  // 4: Proto.MoveInfo
	(*StatusResponse)(nil),            // 5: Proto.StatusResponse
	(*MercenaryStatus)(nil),           // 6: Proto.MercenaryStatus
	(*AccumulatedAmountResponse)(nil), // 7: Proto.AccumulatedAmountResponse
	(*Empty)(nil),                     // 8: Proto.Empty
}
var file_communications_proto_depIdxs = []int32{
	1,  // 0: Proto.MoveInfo.move:type_name -> Proto.Move
	6,  // 1: Proto.StatusResponse.mercenaries:type_name -> Proto.MercenaryStatus
	0,  // 2: Proto.Mercenary.StartFloor:input_type -> Proto.FloorInfo
	1,  // 3: Proto.Mercenary.ReportMove:input_type -> Proto.Move
	2,  // 4: Proto.Director.AddMercenary:input_type -> Proto.MercenaryInfo
	3,  // 5: Proto.Director.ReportDeath:input_type -> Proto.DeathInfo
	4,  // 6: Proto.Director.ReportMove:input_type -> Proto.MoveInfo
	8,  // 7: Proto.Director.CheckStatus:input_type -> Proto.Empty
	8,  // 8: Proto.Director.RequestAccumulatedAmount:input_type -> Proto.Empty
	8,  // 9: Proto.DoshBank.GetAccumulatedAmount:input_type -> Proto.Empty
	8,  // 10: Proto.Mercenary.StartFloor:output_type -> Proto.Empty
	8,  // 11: Proto.Mercenary.ReportMove:output_type -> Proto.Empty
	8,  // 12: Proto.Director.AddMercenary:output_type -> Proto.Empty
	8,  // 13: Proto.Director.ReportDeath:output_type -> Proto.Empty
	8,  // 14: Proto.Director.ReportMove:output_type -> Proto.Empty
	5,  // 15: Proto.Director.CheckStatus:output_type -> Proto.StatusResponse
	7,  // 16: Proto.Director.RequestAccumulatedAmount:output_type -> Proto.AccumulatedAmountResponse
	7,  // 17: Proto.DoshBank.GetAccumulatedAmount:output_type -> Proto.AccumulatedAmountResponse
	10, // [10:18] is the sub-list for method output_type
	2,  // [2:10] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_communications_proto_init() }
func file_communications_proto_init() {
	if File_communications_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_communications_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FloorInfo); i {
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
		file_communications_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Move); i {
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
		file_communications_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MercenaryInfo); i {
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
		file_communications_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeathInfo); i {
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
		file_communications_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MoveInfo); i {
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
		file_communications_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResponse); i {
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
		file_communications_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MercenaryStatus); i {
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
		file_communications_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccumulatedAmountResponse); i {
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
		file_communications_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_communications_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_communications_proto_goTypes,
		DependencyIndexes: file_communications_proto_depIdxs,
		MessageInfos:      file_communications_proto_msgTypes,
	}.Build()
	File_communications_proto = out.File
	file_communications_proto_rawDesc = nil
	file_communications_proto_goTypes = nil
	file_communications_proto_depIdxs = nil
}