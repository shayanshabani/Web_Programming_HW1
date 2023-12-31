// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: AuthService.proto

package proto

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

type ReqPq_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nonce     string `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"` // random string sent by client with size 20
	MessageId int32  `protobuf:"varint,2,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
}

func (x *ReqPq_Request) Reset() {
	*x = ReqPq_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AuthService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqPq_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqPq_Request) ProtoMessage() {}

func (x *ReqPq_Request) ProtoReflect() protoreflect.Message {
	mi := &file_AuthService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqPq_Request.ProtoReflect.Descriptor instead.
func (*ReqPq_Request) Descriptor() ([]byte, []int) {
	return file_AuthService_proto_rawDescGZIP(), []int{0}
}

func (x *ReqPq_Request) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *ReqPq_Request) GetMessageId() int32 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

type ReqPq_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nonce       string `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`                                // the same string that client sent with size 20
	ServerNonce string `protobuf:"bytes,2,opt,name=server_nonce,json=serverNonce,proto3" json:"server_nonce,omitempty"` // another random string with size 20, that server sends to client
	MessageId   int32  `protobuf:"varint,3,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	P           int32  `protobuf:"varint,4,opt,name=p,proto3" json:"p,omitempty"` // a prime number like 23
	G           int32  `protobuf:"varint,5,opt,name=g,proto3" json:"g,omitempty"` // a generating number like 5
}

func (x *ReqPq_Response) Reset() {
	*x = ReqPq_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AuthService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqPq_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqPq_Response) ProtoMessage() {}

func (x *ReqPq_Response) ProtoReflect() protoreflect.Message {
	mi := &file_AuthService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqPq_Response.ProtoReflect.Descriptor instead.
func (*ReqPq_Response) Descriptor() ([]byte, []int) {
	return file_AuthService_proto_rawDescGZIP(), []int{1}
}

func (x *ReqPq_Response) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *ReqPq_Response) GetServerNonce() string {
	if x != nil {
		return x.ServerNonce
	}
	return ""
}

func (x *ReqPq_Response) GetMessageId() int32 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

func (x *ReqPq_Response) GetP() int32 {
	if x != nil {
		return x.P
	}
	return 0
}

func (x *ReqPq_Response) GetG() int32 {
	if x != nil {
		return x.G
	}
	return 0
}

type ReqDh_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nonce       string `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`                                // random string generated by client in req_pq
	ServerNonce string `protobuf:"bytes,2,opt,name=server_nonce,json=serverNonce,proto3" json:"server_nonce,omitempty"` // random string generated by server in req_pq
	MessageId   int32  `protobuf:"varint,3,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`      // a random even integer, greater than of that in req_pq
	A           int32  `protobuf:"varint,4,opt,name=A,proto3" json:"A,omitempty"`                                       // public key generated by client
}

func (x *ReqDh_Request) Reset() {
	*x = ReqDh_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AuthService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqDh_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqDh_Request) ProtoMessage() {}

func (x *ReqDh_Request) ProtoReflect() protoreflect.Message {
	mi := &file_AuthService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqDh_Request.ProtoReflect.Descriptor instead.
func (*ReqDh_Request) Descriptor() ([]byte, []int) {
	return file_AuthService_proto_rawDescGZIP(), []int{2}
}

func (x *ReqDh_Request) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *ReqDh_Request) GetServerNonce() string {
	if x != nil {
		return x.ServerNonce
	}
	return ""
}

func (x *ReqDh_Request) GetMessageId() int32 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

func (x *ReqDh_Request) GetA() int32 {
	if x != nil {
		return x.A
	}
	return 0
}

type ReqDh_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nonce       string `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`                                // the same string that client sent
	ServerNonce string `protobuf:"bytes,2,opt,name=server_nonce,json=serverNonce,proto3" json:"server_nonce,omitempty"` // another random string of size 20
	MessageId   int32  `protobuf:"varint,3,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`      // random odd integer
	B           int32  `protobuf:"varint,4,opt,name=B,proto3" json:"B,omitempty"`                                       // public key generated by server
}

func (x *ReqDh_Response) Reset() {
	*x = ReqDh_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AuthService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqDh_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqDh_Response) ProtoMessage() {}

func (x *ReqDh_Response) ProtoReflect() protoreflect.Message {
	mi := &file_AuthService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqDh_Response.ProtoReflect.Descriptor instead.
func (*ReqDh_Response) Descriptor() ([]byte, []int) {
	return file_AuthService_proto_rawDescGZIP(), []int{3}
}

func (x *ReqDh_Response) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *ReqDh_Response) GetServerNonce() string {
	if x != nil {
		return x.ServerNonce
	}
	return ""
}

func (x *ReqDh_Response) GetMessageId() int32 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

func (x *ReqDh_Response) GetB() int32 {
	if x != nil {
		return x.B
	}
	return 0
}

type Seek struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthKey int32 `protobuf:"varint,1,opt,name=auth_key,json=authKey,proto3" json:"auth_key,omitempty"` // the key that the user suppose to have
}

func (x *Seek) Reset() {
	*x = Seek{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AuthService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Seek) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Seek) ProtoMessage() {}

func (x *Seek) ProtoReflect() protoreflect.Message {
	mi := &file_AuthService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Seek.ProtoReflect.Descriptor instead.
func (*Seek) Descriptor() ([]byte, []int) {
	return file_AuthService_proto_rawDescGZIP(), []int{4}
}

func (x *Seek) GetAuthKey() int32 {
	if x != nil {
		return x.AuthKey
	}
	return 0
}

type Confirm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valid bool `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"` // if the auth key is valid or not
}

func (x *Confirm) Reset() {
	*x = Confirm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AuthService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Confirm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Confirm) ProtoMessage() {}

func (x *Confirm) ProtoReflect() protoreflect.Message {
	mi := &file_AuthService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Confirm.ProtoReflect.Descriptor instead.
func (*Confirm) Descriptor() ([]byte, []int) {
	return file_AuthService_proto_rawDescGZIP(), []int{5}
}

func (x *Confirm) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

var File_AuthService_proto protoreflect.FileDescriptor

var file_AuthService_proto_rawDesc = []byte{
	0x0a, 0x11, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x0e, 0x72, 0x65,
	0x71, 0x5f, 0x70, 0x71, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e,
	0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49,
	0x64, 0x22, 0x85, 0x01, 0x0a, 0x0f, 0x72, 0x65, 0x71, 0x5f, 0x70, 0x71, 0x5f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x0c, 0x0a,
	0x01, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x70, 0x12, 0x0c, 0x0a, 0x01, 0x67,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x67, 0x22, 0x76, 0x0a, 0x0e, 0x72, 0x65, 0x71,
	0x5f, 0x64, 0x68, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e,
	0x6f, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x6f, 0x6e, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e,
	0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x49, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x41, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01,
	0x41, 0x22, 0x77, 0x0a, 0x0f, 0x72, 0x65, 0x71, 0x5f, 0x64, 0x68, 0x5f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x0c, 0x0a, 0x01,
	0x42, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x42, 0x22, 0x21, 0x0a, 0x04, 0x53, 0x65,
	0x65, 0x6b, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x22, 0x1f, 0x0a,
	0x07, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x32, 0xba,
	0x01, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x5f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x39, 0x0a, 0x06, 0x72, 0x65, 0x71, 0x5f, 0x70, 0x71, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x72, 0x65, 0x71, 0x5f, 0x70, 0x71, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x71, 0x5f, 0x70, 0x71, 0x5f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0d, 0x72, 0x65,
	0x71, 0x5f, 0x44, 0x48, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x15, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x71, 0x5f, 0x64, 0x68, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x71, 0x5f, 0x64,
	0x68, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x0c,
	0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x65, 0x6b, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x77,
	0x65, 0x62, 0x5f, 0x68, 0x77, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AuthService_proto_rawDescOnce sync.Once
	file_AuthService_proto_rawDescData = file_AuthService_proto_rawDesc
)

func file_AuthService_proto_rawDescGZIP() []byte {
	file_AuthService_proto_rawDescOnce.Do(func() {
		file_AuthService_proto_rawDescData = protoimpl.X.CompressGZIP(file_AuthService_proto_rawDescData)
	})
	return file_AuthService_proto_rawDescData
}

var file_AuthService_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_AuthService_proto_goTypes = []interface{}{
	(*ReqPq_Request)(nil),  // 0: proto.req_pq_Request
	(*ReqPq_Response)(nil), // 1: proto.req_pq_Response
	(*ReqDh_Request)(nil),  // 2: proto.req_dh_Request
	(*ReqDh_Response)(nil), // 3: proto.req_dh_Response
	(*Seek)(nil),           // 4: proto.Seek
	(*Confirm)(nil),        // 5: proto.Confirm
}
var file_AuthService_proto_depIdxs = []int32{
	0, // 0: proto.Auth_Service.req_pq:input_type -> proto.req_pq_Request
	2, // 1: proto.Auth_Service.req_DH_params:input_type -> proto.req_dh_Request
	4, // 2: proto.Auth_Service.Authenticate:input_type -> proto.Seek
	1, // 3: proto.Auth_Service.req_pq:output_type -> proto.req_pq_Response
	3, // 4: proto.Auth_Service.req_DH_params:output_type -> proto.req_dh_Response
	5, // 5: proto.Auth_Service.Authenticate:output_type -> proto.Confirm
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_AuthService_proto_init() }
func file_AuthService_proto_init() {
	if File_AuthService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_AuthService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqPq_Request); i {
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
		file_AuthService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqPq_Response); i {
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
		file_AuthService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqDh_Request); i {
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
		file_AuthService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqDh_Response); i {
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
		file_AuthService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Seek); i {
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
		file_AuthService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Confirm); i {
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
			RawDescriptor: file_AuthService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_AuthService_proto_goTypes,
		DependencyIndexes: file_AuthService_proto_depIdxs,
		MessageInfos:      file_AuthService_proto_msgTypes,
	}.Build()
	File_AuthService_proto = out.File
	file_AuthService_proto_rawDesc = nil
	file_AuthService_proto_goTypes = nil
	file_AuthService_proto_depIdxs = nil
}
