// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: GRPCcom.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdProducto  string `protobuf:"bytes,1,opt,name=IdProducto,proto3" json:"IdProducto,omitempty"`
	Producto    string `protobuf:"bytes,2,opt,name=Producto,proto3" json:"Producto,omitempty"`
	Valor       int32  `protobuf:"varint,3,opt,name=Valor,proto3" json:"Valor,omitempty"`
	Origen      string `protobuf:"bytes,4,opt,name=Origen,proto3" json:"Origen,omitempty"`
	Destino     string `protobuf:"bytes,5,opt,name=Destino,proto3" json:"Destino,omitempty"`
	Prioritario string `protobuf:"bytes,6,opt,name=Prioritario,proto3" json:"Prioritario,omitempty"`
	Convenio    string `protobuf:"bytes,7,opt,name=Convenio,proto3" json:"Convenio,omitempty"`
}

func (x *ClientRequest) Reset() {
	*x = ClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GRPCcom_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientRequest) ProtoMessage() {}

func (x *ClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_GRPCcom_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientRequest.ProtoReflect.Descriptor instead.
func (*ClientRequest) Descriptor() ([]byte, []int) {
	return file_GRPCcom_proto_rawDescGZIP(), []int{0}
}

func (x *ClientRequest) GetIdProducto() string {
	if x != nil {
		return x.IdProducto
	}
	return ""
}

func (x *ClientRequest) GetProducto() string {
	if x != nil {
		return x.Producto
	}
	return ""
}

func (x *ClientRequest) GetValor() int32 {
	if x != nil {
		return x.Valor
	}
	return 0
}

func (x *ClientRequest) GetOrigen() string {
	if x != nil {
		return x.Origen
	}
	return ""
}

func (x *ClientRequest) GetDestino() string {
	if x != nil {
		return x.Destino
	}
	return ""
}

func (x *ClientRequest) GetPrioritario() string {
	if x != nil {
		return x.Prioritario
	}
	return ""
}

func (x *ClientRequest) GetConvenio() string {
	if x != nil {
		return x.Convenio
	}
	return ""
}

type TrackCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=Code,proto3" json:"Code,omitempty"`
}

func (x *TrackCode) Reset() {
	*x = TrackCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GRPCcom_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrackCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackCode) ProtoMessage() {}

func (x *TrackCode) ProtoReflect() protoreflect.Message {
	mi := &file_GRPCcom_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackCode.ProtoReflect.Descriptor instead.
func (*TrackCode) Descriptor() ([]byte, []int) {
	return file_GRPCcom_proto_rawDescGZIP(), []int{1}
}

func (x *TrackCode) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type OrderStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Estado   string `protobuf:"bytes,1,opt,name=Estado,proto3" json:"Estado,omitempty"`
	Intentos int32  `protobuf:"varint,2,opt,name=Intentos,proto3" json:"Intentos,omitempty"`
}

func (x *OrderStatus) Reset() {
	*x = OrderStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GRPCcom_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderStatus) ProtoMessage() {}

func (x *OrderStatus) ProtoReflect() protoreflect.Message {
	mi := &file_GRPCcom_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderStatus.ProtoReflect.Descriptor instead.
func (*OrderStatus) Descriptor() ([]byte, []int) {
	return file_GRPCcom_proto_rawDescGZIP(), []int{2}
}

func (x *OrderStatus) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

func (x *OrderStatus) GetIntentos() int32 {
	if x != nil {
		return x.Intentos
	}
	return 0
}

type RespuestaCamion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Idcamion string `protobuf:"bytes,1,opt,name=Idcamion,proto3" json:"Idcamion,omitempty"`
}

func (x *RespuestaCamion) Reset() {
	*x = RespuestaCamion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GRPCcom_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespuestaCamion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespuestaCamion) ProtoMessage() {}

func (x *RespuestaCamion) ProtoReflect() protoreflect.Message {
	mi := &file_GRPCcom_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespuestaCamion.ProtoReflect.Descriptor instead.
func (*RespuestaCamion) Descriptor() ([]byte, []int) {
	return file_GRPCcom_proto_rawDescGZIP(), []int{3}
}

func (x *RespuestaCamion) GetIdcamion() string {
	if x != nil {
		return x.Idcamion
	}
	return ""
}

//
//message ConsultaCamion{
//
//string IdSeguimiento = 1;
//string IdCamion = 2;
//
//}
type Package struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdPaquete   string `protobuf:"bytes,1,opt,name=IdPaquete,proto3" json:"IdPaquete,omitempty"`
	Seguimiento string `protobuf:"bytes,2,opt,name=Seguimiento,proto3" json:"Seguimiento,omitempty"`
	Tipo        string `protobuf:"bytes,3,opt,name=Tipo,proto3" json:"Tipo,omitempty"`
	Valor       int32  `protobuf:"varint,4,opt,name=Valor,proto3" json:"Valor,omitempty"`
	Intentos    int32  `protobuf:"varint,5,opt,name=Intentos,proto3" json:"Intentos,omitempty"`
	Estado      string `protobuf:"bytes,6,opt,name=Estado,proto3" json:"Estado,omitempty"`
	Origen      string `protobuf:"bytes,7,opt,name=Origen,proto3" json:"Origen,omitempty"`
	Destino     string `protobuf:"bytes,8,opt,name=Destino,proto3" json:"Destino,omitempty"`
}

func (x *Package) Reset() {
	*x = Package{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GRPCcom_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Package) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Package) ProtoMessage() {}

func (x *Package) ProtoReflect() protoreflect.Message {
	mi := &file_GRPCcom_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Package.ProtoReflect.Descriptor instead.
func (*Package) Descriptor() ([]byte, []int) {
	return file_GRPCcom_proto_rawDescGZIP(), []int{4}
}

func (x *Package) GetIdPaquete() string {
	if x != nil {
		return x.IdPaquete
	}
	return ""
}

func (x *Package) GetSeguimiento() string {
	if x != nil {
		return x.Seguimiento
	}
	return ""
}

func (x *Package) GetTipo() string {
	if x != nil {
		return x.Tipo
	}
	return ""
}

func (x *Package) GetValor() int32 {
	if x != nil {
		return x.Valor
	}
	return 0
}

func (x *Package) GetIntentos() int32 {
	if x != nil {
		return x.Intentos
	}
	return 0
}

func (x *Package) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

func (x *Package) GetOrigen() string {
	if x != nil {
		return x.Origen
	}
	return ""
}

func (x *Package) GetDestino() string {
	if x != nil {
		return x.Destino
	}
	return ""
}

type PaqueteEntregado struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdPaquete   string `protobuf:"bytes,1,opt,name=IdPaquete,proto3" json:"IdPaquete,omitempty"`
	Tipo        string `protobuf:"bytes,3,opt,name=Tipo,proto3" json:"Tipo,omitempty"`
	Valor       int32  `protobuf:"varint,4,opt,name=Valor,proto3" json:"Valor,omitempty"`
	Intentos    int32  `protobuf:"varint,5,opt,name=Intentos,proto3" json:"Intentos,omitempty"`
	Timestamp   string `protobuf:"bytes,7,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	IdCamion    string `protobuf:"bytes,8,opt,name=IdCamion,proto3" json:"IdCamion,omitempty"`
	Estado      string `protobuf:"bytes,9,opt,name=estado,proto3" json:"estado,omitempty"`
	Seguimiento string `protobuf:"bytes,10,opt,name=seguimiento,proto3" json:"seguimiento,omitempty"`
}

func (x *PaqueteEntregado) Reset() {
	*x = PaqueteEntregado{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GRPCcom_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaqueteEntregado) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaqueteEntregado) ProtoMessage() {}

func (x *PaqueteEntregado) ProtoReflect() protoreflect.Message {
	mi := &file_GRPCcom_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaqueteEntregado.ProtoReflect.Descriptor instead.
func (*PaqueteEntregado) Descriptor() ([]byte, []int) {
	return file_GRPCcom_proto_rawDescGZIP(), []int{5}
}

func (x *PaqueteEntregado) GetIdPaquete() string {
	if x != nil {
		return x.IdPaquete
	}
	return ""
}

func (x *PaqueteEntregado) GetTipo() string {
	if x != nil {
		return x.Tipo
	}
	return ""
}

func (x *PaqueteEntregado) GetValor() int32 {
	if x != nil {
		return x.Valor
	}
	return 0
}

func (x *PaqueteEntregado) GetIntentos() int32 {
	if x != nil {
		return x.Intentos
	}
	return 0
}

func (x *PaqueteEntregado) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *PaqueteEntregado) GetIdCamion() string {
	if x != nil {
		return x.IdCamion
	}
	return ""
}

func (x *PaqueteEntregado) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

func (x *PaqueteEntregado) GetSeguimiento() string {
	if x != nil {
		return x.Seguimiento
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GRPCcom_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_GRPCcom_proto_msgTypes[6]
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
	return file_GRPCcom_proto_rawDescGZIP(), []int{6}
}

var File_GRPCcom_proto protoreflect.FileDescriptor

var file_GRPCcom_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x47, 0x52, 0x50, 0x43, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x01, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x64, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x49, 0x64,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x72,
	0x69, 0x67, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4f, 0x72, 0x69, 0x67,
	0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f, 0x12, 0x20, 0x0a, 0x0b,
	0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x61, 0x72, 0x69, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x61, 0x72, 0x69, 0x6f, 0x12, 0x1a,
	0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x6e, 0x69, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x6e, 0x69, 0x6f, 0x22, 0x1f, 0x0a, 0x09, 0x54, 0x72,
	0x61, 0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x41, 0x0a, 0x0b, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x45, 0x73,
	0x74, 0x61, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x45, 0x73, 0x74, 0x61,
	0x64, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x6f, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x6f, 0x73, 0x22, 0x2d,
	0x0a, 0x0f, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x43, 0x61, 0x6d, 0x69, 0x6f,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x64, 0x63, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x49, 0x64, 0x63, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x22, 0xd9, 0x01,
	0x0a, 0x07, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x64, 0x50,
	0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x49, 0x64,
	0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x65, 0x67, 0x75, 0x69,
	0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x53, 0x65,
	0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x69, 0x70,
	0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x69, 0x70, 0x6f, 0x12, 0x14, 0x0a,
	0x05, 0x56, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x56, 0x61,
	0x6c, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x6f, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x6f, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x72, 0x69, 0x67, 0x65,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4f, 0x72, 0x69, 0x67, 0x65, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f, 0x22, 0xea, 0x01, 0x0a, 0x10, 0x50, 0x61,
	0x71, 0x75, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x65, 0x67, 0x61, 0x64, 0x6f, 0x12, 0x1c,
	0x0a, 0x09, 0x49, 0x64, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x49, 0x64, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x54, 0x69, 0x70, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x69, 0x70, 0x6f,
	0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x56, 0x61, 0x6c, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x6f, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x6f, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x1a, 0x0a, 0x08, 0x49, 0x64, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x49, 0x64, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x73,
	0x74, 0x61, 0x64, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65,
	0x6e, 0x74, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x67, 0x75, 0x69,
	0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32,
	0xb5, 0x01, 0x0a, 0x10, 0x4c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x61, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x6e, 0x61, 0x72, 0x12,
	0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72,
	0x61, 0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x34, 0x0a, 0x0c, 0x45, 0x73, 0x74, 0x61, 0x64,
	0x6f, 0x50, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x38, 0x0a,
	0x0f, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x65, 0x67, 0x61,
	0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65,
	0x45, 0x6e, 0x74, 0x72, 0x65, 0x67, 0x61, 0x64, 0x6f, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xb0, 0x01, 0x0a, 0x0f, 0x43, 0x61, 0x6d, 0x69,
	0x6f, 0x6e, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x09, 0x41,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x72, 0x43, 0x31, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e,
	0x12, 0x33, 0x0a, 0x09, 0x41, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x72, 0x43, 0x32, 0x12, 0x0e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x1a, 0x16, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x43,
	0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x09, 0x41, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x72,
	0x43, 0x33, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x75,
	0x65, 0x73, 0x74, 0x61, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_GRPCcom_proto_rawDescOnce sync.Once
	file_GRPCcom_proto_rawDescData = file_GRPCcom_proto_rawDesc
)

func file_GRPCcom_proto_rawDescGZIP() []byte {
	file_GRPCcom_proto_rawDescOnce.Do(func() {
		file_GRPCcom_proto_rawDescData = protoimpl.X.CompressGZIP(file_GRPCcom_proto_rawDescData)
	})
	return file_GRPCcom_proto_rawDescData
}

var file_GRPCcom_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_GRPCcom_proto_goTypes = []interface{}{
	(*ClientRequest)(nil),    // 0: proto.ClientRequest
	(*TrackCode)(nil),        // 1: proto.TrackCode
	(*OrderStatus)(nil),      // 2: proto.OrderStatus
	(*RespuestaCamion)(nil),  // 3: proto.RespuestaCamion
	(*Package)(nil),          // 4: proto.Package
	(*PaqueteEntregado)(nil), // 5: proto.PaqueteEntregado
	(*Empty)(nil),            // 6: proto.Empty
}
var file_GRPCcom_proto_depIdxs = []int32{
	0, // 0: proto.LogisticaService.Ordenar:input_type -> proto.ClientRequest
	1, // 1: proto.LogisticaService.EstadoPedido:input_type -> proto.TrackCode
	5, // 2: proto.LogisticaService.InformarEntrega:input_type -> proto.PaqueteEntregado
	4, // 3: proto.CamionesService.AsignarC1:input_type -> proto.Package
	4, // 4: proto.CamionesService.AsignarC2:input_type -> proto.Package
	4, // 5: proto.CamionesService.AsignarC3:input_type -> proto.Package
	1, // 6: proto.LogisticaService.Ordenar:output_type -> proto.TrackCode
	2, // 7: proto.LogisticaService.EstadoPedido:output_type -> proto.OrderStatus
	6, // 8: proto.LogisticaService.InformarEntrega:output_type -> proto.Empty
	3, // 9: proto.CamionesService.AsignarC1:output_type -> proto.RespuestaCamion
	3, // 10: proto.CamionesService.AsignarC2:output_type -> proto.RespuestaCamion
	3, // 11: proto.CamionesService.AsignarC3:output_type -> proto.RespuestaCamion
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GRPCcom_proto_init() }
func file_GRPCcom_proto_init() {
	if File_GRPCcom_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GRPCcom_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientRequest); i {
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
		file_GRPCcom_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrackCode); i {
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
		file_GRPCcom_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderStatus); i {
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
		file_GRPCcom_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespuestaCamion); i {
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
		file_GRPCcom_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Package); i {
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
		file_GRPCcom_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaqueteEntregado); i {
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
		file_GRPCcom_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_GRPCcom_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_GRPCcom_proto_goTypes,
		DependencyIndexes: file_GRPCcom_proto_depIdxs,
		MessageInfos:      file_GRPCcom_proto_msgTypes,
	}.Build()
	File_GRPCcom_proto = out.File
	file_GRPCcom_proto_rawDesc = nil
	file_GRPCcom_proto_goTypes = nil
	file_GRPCcom_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LogisticaServiceClient is the client API for LogisticaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogisticaServiceClient interface {
	Ordenar(ctx context.Context, in *ClientRequest, opts ...grpc.CallOption) (*TrackCode, error)
	EstadoPedido(ctx context.Context, in *TrackCode, opts ...grpc.CallOption) (*OrderStatus, error)
	InformarEntrega(ctx context.Context, in *PaqueteEntregado, opts ...grpc.CallOption) (*Empty, error)
}

type logisticaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLogisticaServiceClient(cc grpc.ClientConnInterface) LogisticaServiceClient {
	return &logisticaServiceClient{cc}
}

func (c *logisticaServiceClient) Ordenar(ctx context.Context, in *ClientRequest, opts ...grpc.CallOption) (*TrackCode, error) {
	out := new(TrackCode)
	err := c.cc.Invoke(ctx, "/proto.LogisticaService/Ordenar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logisticaServiceClient) EstadoPedido(ctx context.Context, in *TrackCode, opts ...grpc.CallOption) (*OrderStatus, error) {
	out := new(OrderStatus)
	err := c.cc.Invoke(ctx, "/proto.LogisticaService/EstadoPedido", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logisticaServiceClient) InformarEntrega(ctx context.Context, in *PaqueteEntregado, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.LogisticaService/InformarEntrega", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogisticaServiceServer is the server API for LogisticaService service.
type LogisticaServiceServer interface {
	Ordenar(context.Context, *ClientRequest) (*TrackCode, error)
	EstadoPedido(context.Context, *TrackCode) (*OrderStatus, error)
	InformarEntrega(context.Context, *PaqueteEntregado) (*Empty, error)
}

// UnimplementedLogisticaServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLogisticaServiceServer struct {
}

func (*UnimplementedLogisticaServiceServer) Ordenar(context.Context, *ClientRequest) (*TrackCode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ordenar not implemented")
}
func (*UnimplementedLogisticaServiceServer) EstadoPedido(context.Context, *TrackCode) (*OrderStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EstadoPedido not implemented")
}
func (*UnimplementedLogisticaServiceServer) InformarEntrega(context.Context, *PaqueteEntregado) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InformarEntrega not implemented")
}

func RegisterLogisticaServiceServer(s *grpc.Server, srv LogisticaServiceServer) {
	s.RegisterService(&_LogisticaService_serviceDesc, srv)
}

func _LogisticaService_Ordenar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogisticaServiceServer).Ordenar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LogisticaService/Ordenar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogisticaServiceServer).Ordenar(ctx, req.(*ClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogisticaService_EstadoPedido_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrackCode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogisticaServiceServer).EstadoPedido(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LogisticaService/EstadoPedido",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogisticaServiceServer).EstadoPedido(ctx, req.(*TrackCode))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogisticaService_InformarEntrega_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaqueteEntregado)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogisticaServiceServer).InformarEntrega(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LogisticaService/InformarEntrega",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogisticaServiceServer).InformarEntrega(ctx, req.(*PaqueteEntregado))
	}
	return interceptor(ctx, in, info, handler)
}

var _LogisticaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.LogisticaService",
	HandlerType: (*LogisticaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ordenar",
			Handler:    _LogisticaService_Ordenar_Handler,
		},
		{
			MethodName: "EstadoPedido",
			Handler:    _LogisticaService_EstadoPedido_Handler,
		},
		{
			MethodName: "InformarEntrega",
			Handler:    _LogisticaService_InformarEntrega_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "GRPCcom.proto",
}

// CamionesServiceClient is the client API for CamionesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CamionesServiceClient interface {
	AsignarC1(ctx context.Context, in *Package, opts ...grpc.CallOption) (*RespuestaCamion, error)
	AsignarC2(ctx context.Context, in *Package, opts ...grpc.CallOption) (*RespuestaCamion, error)
	AsignarC3(ctx context.Context, in *Package, opts ...grpc.CallOption) (*RespuestaCamion, error)
}

type camionesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCamionesServiceClient(cc grpc.ClientConnInterface) CamionesServiceClient {
	return &camionesServiceClient{cc}
}

func (c *camionesServiceClient) AsignarC1(ctx context.Context, in *Package, opts ...grpc.CallOption) (*RespuestaCamion, error) {
	out := new(RespuestaCamion)
	err := c.cc.Invoke(ctx, "/proto.CamionesService/AsignarC1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *camionesServiceClient) AsignarC2(ctx context.Context, in *Package, opts ...grpc.CallOption) (*RespuestaCamion, error) {
	out := new(RespuestaCamion)
	err := c.cc.Invoke(ctx, "/proto.CamionesService/AsignarC2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *camionesServiceClient) AsignarC3(ctx context.Context, in *Package, opts ...grpc.CallOption) (*RespuestaCamion, error) {
	out := new(RespuestaCamion)
	err := c.cc.Invoke(ctx, "/proto.CamionesService/AsignarC3", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CamionesServiceServer is the server API for CamionesService service.
type CamionesServiceServer interface {
	AsignarC1(context.Context, *Package) (*RespuestaCamion, error)
	AsignarC2(context.Context, *Package) (*RespuestaCamion, error)
	AsignarC3(context.Context, *Package) (*RespuestaCamion, error)
}

// UnimplementedCamionesServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCamionesServiceServer struct {
}

func (*UnimplementedCamionesServiceServer) AsignarC1(context.Context, *Package) (*RespuestaCamion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AsignarC1 not implemented")
}
func (*UnimplementedCamionesServiceServer) AsignarC2(context.Context, *Package) (*RespuestaCamion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AsignarC2 not implemented")
}
func (*UnimplementedCamionesServiceServer) AsignarC3(context.Context, *Package) (*RespuestaCamion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AsignarC3 not implemented")
}

func RegisterCamionesServiceServer(s *grpc.Server, srv CamionesServiceServer) {
	s.RegisterService(&_CamionesService_serviceDesc, srv)
}

func _CamionesService_AsignarC1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Package)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CamionesServiceServer).AsignarC1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CamionesService/AsignarC1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CamionesServiceServer).AsignarC1(ctx, req.(*Package))
	}
	return interceptor(ctx, in, info, handler)
}

func _CamionesService_AsignarC2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Package)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CamionesServiceServer).AsignarC2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CamionesService/AsignarC2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CamionesServiceServer).AsignarC2(ctx, req.(*Package))
	}
	return interceptor(ctx, in, info, handler)
}

func _CamionesService_AsignarC3_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Package)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CamionesServiceServer).AsignarC3(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CamionesService/AsignarC3",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CamionesServiceServer).AsignarC3(ctx, req.(*Package))
	}
	return interceptor(ctx, in, info, handler)
}

var _CamionesService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CamionesService",
	HandlerType: (*CamionesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AsignarC1",
			Handler:    _CamionesService_AsignarC1_Handler,
		},
		{
			MethodName: "AsignarC2",
			Handler:    _CamionesService_AsignarC2_Handler,
		},
		{
			MethodName: "AsignarC3",
			Handler:    _CamionesService_AsignarC3_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "GRPCcom.proto",
}
