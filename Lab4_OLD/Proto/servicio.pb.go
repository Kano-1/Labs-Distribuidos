// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.27.0--rc1
// source: servicio.proto

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

type ComunicarDecision struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mercenario string   `protobuf:"bytes,1,opt,name=mercenario,proto3" json:"mercenario,omitempty"`
	Decisiones []string `protobuf:"bytes,2,rep,name=decisiones,proto3" json:"decisiones,omitempty"`
	Piso       int32    `protobuf:"varint,3,opt,name=piso,proto3" json:"piso,omitempty"`
}

func (x *ComunicarDecision) Reset() {
	*x = ComunicarDecision{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servicio_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComunicarDecision) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComunicarDecision) ProtoMessage() {}

func (x *ComunicarDecision) ProtoReflect() protoreflect.Message {
	mi := &file_servicio_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComunicarDecision.ProtoReflect.Descriptor instead.
func (*ComunicarDecision) Descriptor() ([]byte, []int) {
	return file_servicio_proto_rawDescGZIP(), []int{0}
}

func (x *ComunicarDecision) GetMercenario() string {
	if x != nil {
		return x.Mercenario
	}
	return ""
}

func (x *ComunicarDecision) GetDecisiones() []string {
	if x != nil {
		return x.Decisiones
	}
	return nil
}

func (x *ComunicarDecision) GetPiso() int32 {
	if x != nil {
		return x.Piso
	}
	return 0
}

type RespuestaExito struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exito bool `protobuf:"varint,1,opt,name=exito,proto3" json:"exito,omitempty"`
}

func (x *RespuestaExito) Reset() {
	*x = RespuestaExito{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servicio_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespuestaExito) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespuestaExito) ProtoMessage() {}

func (x *RespuestaExito) ProtoReflect() protoreflect.Message {
	mi := &file_servicio_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespuestaExito.ProtoReflect.Descriptor instead.
func (*RespuestaExito) Descriptor() ([]byte, []int) {
	return file_servicio_proto_rawDescGZIP(), []int{1}
}

func (x *RespuestaExito) GetExito() bool {
	if x != nil {
		return x.Exito
	}
	return false
}

type SolicitarMonto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mercenario string `protobuf:"bytes,1,opt,name=mercenario,proto3" json:"mercenario,omitempty"`
}

func (x *SolicitarMonto) Reset() {
	*x = SolicitarMonto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servicio_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SolicitarMonto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolicitarMonto) ProtoMessage() {}

func (x *SolicitarMonto) ProtoReflect() protoreflect.Message {
	mi := &file_servicio_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolicitarMonto.ProtoReflect.Descriptor instead.
func (*SolicitarMonto) Descriptor() ([]byte, []int) {
	return file_servicio_proto_rawDescGZIP(), []int{2}
}

func (x *SolicitarMonto) GetMercenario() string {
	if x != nil {
		return x.Mercenario
	}
	return ""
}

type EntregarMonto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Monto int32 `protobuf:"varint,1,opt,name=monto,proto3" json:"monto,omitempty"`
}

func (x *EntregarMonto) Reset() {
	*x = EntregarMonto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servicio_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntregarMonto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntregarMonto) ProtoMessage() {}

func (x *EntregarMonto) ProtoReflect() protoreflect.Message {
	mi := &file_servicio_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntregarMonto.ProtoReflect.Descriptor instead.
func (*EntregarMonto) Descriptor() ([]byte, []int) {
	return file_servicio_proto_rawDescGZIP(), []int{3}
}

func (x *EntregarMonto) GetMonto() int32 {
	if x != nil {
		return x.Monto
	}
	return 0
}

type EnviarEstado struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mercenario string `protobuf:"bytes,1,opt,name=mercenario,proto3" json:"mercenario,omitempty"`
}

func (x *EnviarEstado) Reset() {
	*x = EnviarEstado{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servicio_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnviarEstado) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnviarEstado) ProtoMessage() {}

func (x *EnviarEstado) ProtoReflect() protoreflect.Message {
	mi := &file_servicio_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnviarEstado.ProtoReflect.Descriptor instead.
func (*EnviarEstado) Descriptor() ([]byte, []int) {
	return file_servicio_proto_rawDescGZIP(), []int{4}
}

func (x *EnviarEstado) GetMercenario() string {
	if x != nil {
		return x.Mercenario
	}
	return ""
}

type ConfirmarInicio struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inicio bool `protobuf:"varint,1,opt,name=inicio,proto3" json:"inicio,omitempty"`
}

func (x *ConfirmarInicio) Reset() {
	*x = ConfirmarInicio{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servicio_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmarInicio) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmarInicio) ProtoMessage() {}

func (x *ConfirmarInicio) ProtoReflect() protoreflect.Message {
	mi := &file_servicio_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmarInicio.ProtoReflect.Descriptor instead.
func (*ConfirmarInicio) Descriptor() ([]byte, []int) {
	return file_servicio_proto_rawDescGZIP(), []int{5}
}

func (x *ConfirmarInicio) GetInicio() bool {
	if x != nil {
		return x.Inicio
	}
	return false
}

var File_servicio_proto protoreflect.FileDescriptor

var file_servicio_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x75, 0x6e,
	0x69, 0x63, 0x61, 0x72, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a,
	0x6d, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x12, 0x1e, 0x0a, 0x0a,
	0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0a, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x69, 0x73, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x69, 0x73, 0x6f,
	0x22, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x45, 0x78, 0x69,
	0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x78, 0x69, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x65, 0x78, 0x69, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x0e, 0x73, 0x6f, 0x6c, 0x69,
	0x63, 0x69, 0x74, 0x61, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x6d, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x22, 0x25, 0x0a, 0x0d, 0x65, 0x6e,
	0x74, 0x72, 0x65, 0x67, 0x61, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x6d,
	0x6f, 0x6e, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74,
	0x6f, 0x22, 0x2e, 0x0a, 0x0c, 0x65, 0x6e, 0x76, 0x69, 0x61, 0x72, 0x45, 0x73, 0x74, 0x61, 0x64,
	0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69,
	0x6f, 0x22, 0x29, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x72, 0x49, 0x6e,
	0x69, 0x63, 0x69, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6e, 0x69, 0x63, 0x69, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x6e, 0x69, 0x63, 0x69, 0x6f, 0x32, 0xd9, 0x01, 0x0a,
	0x13, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x69, 0x6f, 0x4d, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x61,
	0x72, 0x69, 0x6f, 0x73, 0x12, 0x43, 0x0a, 0x10, 0x65, 0x6e, 0x76, 0x69, 0x61, 0x72, 0x44, 0x65,
	0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x63, 0x6f, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x72, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x1a, 0x15, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x75,
	0x65, 0x73, 0x74, 0x61, 0x45, 0x78, 0x69, 0x74, 0x6f, 0x12, 0x40, 0x0a, 0x11, 0x76, 0x65, 0x72,
	0x4d, 0x6f, 0x6e, 0x74, 0x6f, 0x41, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x64, 0x6f, 0x12, 0x15,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x61, 0x72,
	0x4d, 0x6f, 0x6e, 0x74, 0x6f, 0x1a, 0x14, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x6e,
	0x74, 0x72, 0x65, 0x67, 0x61, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x6f, 0x12, 0x3b, 0x0a, 0x0c, 0x63,
	0x6f, 0x6d, 0x65, 0x6e, 0x7a, 0x61, 0x72, 0x50, 0x69, 0x73, 0x6f, 0x12, 0x13, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x61, 0x72, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f,
	0x1a, 0x16, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x61, 0x72, 0x49, 0x6e, 0x69, 0x63, 0x69, 0x6f, 0x42, 0x0c, 0x5a, 0x0a, 0x4c, 0x61, 0x62, 0x34,
	0x2f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_servicio_proto_rawDescOnce sync.Once
	file_servicio_proto_rawDescData = file_servicio_proto_rawDesc
)

func file_servicio_proto_rawDescGZIP() []byte {
	file_servicio_proto_rawDescOnce.Do(func() {
		file_servicio_proto_rawDescData = protoimpl.X.CompressGZIP(file_servicio_proto_rawDescData)
	})
	return file_servicio_proto_rawDescData
}

var file_servicio_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_servicio_proto_goTypes = []interface{}{
	(*ComunicarDecision)(nil), // 0: Proto.comunicarDecision
	(*RespuestaExito)(nil),    // 1: Proto.respuestaExito
	(*SolicitarMonto)(nil),    // 2: Proto.solicitarMonto
	(*EntregarMonto)(nil),     // 3: Proto.entregarMonto
	(*EnviarEstado)(nil),      // 4: Proto.enviarEstado
	(*ConfirmarInicio)(nil),   // 5: Proto.confirmarInicio
}
var file_servicio_proto_depIdxs = []int32{
	0, // 0: Proto.servicioMercenarios.enviarDecisiones:input_type -> Proto.comunicarDecision
	2, // 1: Proto.servicioMercenarios.verMontoAcumulado:input_type -> Proto.solicitarMonto
	4, // 2: Proto.servicioMercenarios.comenzarPiso:input_type -> Proto.enviarEstado
	1, // 3: Proto.servicioMercenarios.enviarDecisiones:output_type -> Proto.respuestaExito
	3, // 4: Proto.servicioMercenarios.verMontoAcumulado:output_type -> Proto.entregarMonto
	5, // 5: Proto.servicioMercenarios.comenzarPiso:output_type -> Proto.confirmarInicio
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_servicio_proto_init() }
func file_servicio_proto_init() {
	if File_servicio_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_servicio_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComunicarDecision); i {
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
		file_servicio_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespuestaExito); i {
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
		file_servicio_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SolicitarMonto); i {
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
		file_servicio_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntregarMonto); i {
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
		file_servicio_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnviarEstado); i {
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
		file_servicio_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmarInicio); i {
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
			RawDescriptor: file_servicio_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_servicio_proto_goTypes,
		DependencyIndexes: file_servicio_proto_depIdxs,
		MessageInfos:      file_servicio_proto_msgTypes,
	}.Build()
	File_servicio_proto = out.File
	file_servicio_proto_rawDesc = nil
	file_servicio_proto_goTypes = nil
	file_servicio_proto_depIdxs = nil
}