// Code generated with goa v3.7.6, DO NOT EDIT.
//
// secured_service protocol buffer definition
//
// Command:
// $ goa gen goa.design/examples/security/multiauth/design -o security/multiauth

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: goadesign_goagen_secured_service.proto

package secured_servicepb

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

type SigninRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SigninRequest) Reset() {
	*x = SigninRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_secured_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SigninRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SigninRequest) ProtoMessage() {}

func (x *SigninRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_secured_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SigninRequest.ProtoReflect.Descriptor instead.
func (*SigninRequest) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_secured_service_proto_rawDescGZIP(), []int{0}
}

type SigninResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// JWT token
	Jwt string `protobuf:"bytes,1,opt,name=jwt,proto3" json:"jwt,omitempty"`
	// API Key
	ApiKey string `protobuf:"bytes,2,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	// OAuth2 token
	OauthToken string `protobuf:"bytes,3,opt,name=oauth_token,json=oauthToken,proto3" json:"oauth_token,omitempty"`
}

func (x *SigninResponse) Reset() {
	*x = SigninResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_secured_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SigninResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SigninResponse) ProtoMessage() {}

func (x *SigninResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_secured_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SigninResponse.ProtoReflect.Descriptor instead.
func (*SigninResponse) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_secured_service_proto_rawDescGZIP(), []int{1}
}

func (x *SigninResponse) GetJwt() string {
	if x != nil {
		return x.Jwt
	}
	return ""
}

func (x *SigninResponse) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

func (x *SigninResponse) GetOauthToken() string {
	if x != nil {
		return x.OauthToken
	}
	return ""
}

type SecureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether to force auth failure even with a valid JWT
	Fail bool `protobuf:"varint,1,opt,name=fail,proto3" json:"fail,omitempty"`
}

func (x *SecureRequest) Reset() {
	*x = SecureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_secured_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecureRequest) ProtoMessage() {}

func (x *SecureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_secured_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecureRequest.ProtoReflect.Descriptor instead.
func (*SecureRequest) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_secured_service_proto_rawDescGZIP(), []int{2}
}

func (x *SecureRequest) GetFail() bool {
	if x != nil {
		return x.Fail
	}
	return false
}

type SecureResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
}

func (x *SecureResponse) Reset() {
	*x = SecureResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_secured_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecureResponse) ProtoMessage() {}

func (x *SecureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_secured_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecureResponse.ProtoReflect.Descriptor instead.
func (*SecureResponse) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_secured_service_proto_rawDescGZIP(), []int{3}
}

func (x *SecureResponse) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

type DoublySecureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// API key
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *DoublySecureRequest) Reset() {
	*x = DoublySecureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_secured_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoublySecureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoublySecureRequest) ProtoMessage() {}

func (x *DoublySecureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_secured_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoublySecureRequest.ProtoReflect.Descriptor instead.
func (*DoublySecureRequest) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_secured_service_proto_rawDescGZIP(), []int{4}
}

func (x *DoublySecureRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type DoublySecureResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
}

func (x *DoublySecureResponse) Reset() {
	*x = DoublySecureResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_secured_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoublySecureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoublySecureResponse) ProtoMessage() {}

func (x *DoublySecureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_secured_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoublySecureResponse.ProtoReflect.Descriptor instead.
func (*DoublySecureResponse) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_secured_service_proto_rawDescGZIP(), []int{5}
}

func (x *DoublySecureResponse) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

type AlsoDoublySecureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Username used to perform signin
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// Password used to perform signin
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// API key
	Key string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *AlsoDoublySecureRequest) Reset() {
	*x = AlsoDoublySecureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_secured_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlsoDoublySecureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlsoDoublySecureRequest) ProtoMessage() {}

func (x *AlsoDoublySecureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_secured_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlsoDoublySecureRequest.ProtoReflect.Descriptor instead.
func (*AlsoDoublySecureRequest) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_secured_service_proto_rawDescGZIP(), []int{6}
}

func (x *AlsoDoublySecureRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AlsoDoublySecureRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AlsoDoublySecureRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type AlsoDoublySecureResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
}

func (x *AlsoDoublySecureResponse) Reset() {
	*x = AlsoDoublySecureResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_secured_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlsoDoublySecureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlsoDoublySecureResponse) ProtoMessage() {}

func (x *AlsoDoublySecureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_secured_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlsoDoublySecureResponse.ProtoReflect.Descriptor instead.
func (*AlsoDoublySecureResponse) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_secured_service_proto_rawDescGZIP(), []int{7}
}

func (x *AlsoDoublySecureResponse) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

var File_goadesign_goagen_secured_service_proto protoreflect.FileDescriptor

var file_goadesign_goagen_secured_service_proto_rawDesc = []byte{
	0x0a, 0x26, 0x67, 0x6f, 0x61, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x67, 0x6f, 0x61, 0x67,
	0x65, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65,
	0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x53, 0x69, 0x67,
	0x6e, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x5c, 0x0a, 0x0e, 0x53, 0x69,
	0x67, 0x6e, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6a, 0x77, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6a, 0x77, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x61,
	0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x23, 0x0a, 0x0d, 0x53, 0x65, 0x63, 0x75,
	0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x61, 0x69,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x66, 0x61, 0x69, 0x6c, 0x22, 0x26, 0x0a,
	0x0e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x27, 0x0a, 0x13, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x79, 0x53,
	0x65, 0x63, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x2c,
	0x0a, 0x14, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x79, 0x53, 0x65, 0x63, 0x75, 0x72, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x63, 0x0a, 0x17,
	0x41, 0x6c, 0x73, 0x6f, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x79, 0x53, 0x65, 0x63, 0x75, 0x72, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x22, 0x30, 0x0a, 0x18, 0x41, 0x6c, 0x73, 0x6f, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x79, 0x53,
	0x65, 0x63, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x32, 0xec, 0x02, 0x0a, 0x0e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x65, 0x64, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e,
	0x12, 0x1e, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x49, 0x0a, 0x06, 0x53, 0x65, 0x63, 0x75, 0x72, 0x65, 0x12, 0x1e, 0x2e, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65,
	0x63, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65,
	0x63, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x0c,
	0x44, 0x6f, 0x75, 0x62, 0x6c, 0x79, 0x53, 0x65, 0x63, 0x75, 0x72, 0x65, 0x12, 0x24, 0x2e, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44,
	0x6f, 0x75, 0x62, 0x6c, 0x79, 0x53, 0x65, 0x63, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x79, 0x53, 0x65, 0x63, 0x75, 0x72,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x67, 0x0a, 0x10, 0x41, 0x6c, 0x73,
	0x6f, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x79, 0x53, 0x65, 0x63, 0x75, 0x72, 0x65, 0x12, 0x28, 0x2e,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x41, 0x6c, 0x73, 0x6f, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x79, 0x53, 0x65, 0x63, 0x75, 0x72, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65,
	0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x6c, 0x73, 0x6f, 0x44, 0x6f,
	0x75, 0x62, 0x6c, 0x79, 0x53, 0x65, 0x63, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x14, 0x5a, 0x12, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x64, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_goadesign_goagen_secured_service_proto_rawDescOnce sync.Once
	file_goadesign_goagen_secured_service_proto_rawDescData = file_goadesign_goagen_secured_service_proto_rawDesc
)

func file_goadesign_goagen_secured_service_proto_rawDescGZIP() []byte {
	file_goadesign_goagen_secured_service_proto_rawDescOnce.Do(func() {
		file_goadesign_goagen_secured_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_goadesign_goagen_secured_service_proto_rawDescData)
	})
	return file_goadesign_goagen_secured_service_proto_rawDescData
}

var file_goadesign_goagen_secured_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_goadesign_goagen_secured_service_proto_goTypes = []interface{}{
	(*SigninRequest)(nil),            // 0: secured_service.SigninRequest
	(*SigninResponse)(nil),           // 1: secured_service.SigninResponse
	(*SecureRequest)(nil),            // 2: secured_service.SecureRequest
	(*SecureResponse)(nil),           // 3: secured_service.SecureResponse
	(*DoublySecureRequest)(nil),      // 4: secured_service.DoublySecureRequest
	(*DoublySecureResponse)(nil),     // 5: secured_service.DoublySecureResponse
	(*AlsoDoublySecureRequest)(nil),  // 6: secured_service.AlsoDoublySecureRequest
	(*AlsoDoublySecureResponse)(nil), // 7: secured_service.AlsoDoublySecureResponse
}
var file_goadesign_goagen_secured_service_proto_depIdxs = []int32{
	0, // 0: secured_service.SecuredService.Signin:input_type -> secured_service.SigninRequest
	2, // 1: secured_service.SecuredService.Secure:input_type -> secured_service.SecureRequest
	4, // 2: secured_service.SecuredService.DoublySecure:input_type -> secured_service.DoublySecureRequest
	6, // 3: secured_service.SecuredService.AlsoDoublySecure:input_type -> secured_service.AlsoDoublySecureRequest
	1, // 4: secured_service.SecuredService.Signin:output_type -> secured_service.SigninResponse
	3, // 5: secured_service.SecuredService.Secure:output_type -> secured_service.SecureResponse
	5, // 6: secured_service.SecuredService.DoublySecure:output_type -> secured_service.DoublySecureResponse
	7, // 7: secured_service.SecuredService.AlsoDoublySecure:output_type -> secured_service.AlsoDoublySecureResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_goadesign_goagen_secured_service_proto_init() }
func file_goadesign_goagen_secured_service_proto_init() {
	if File_goadesign_goagen_secured_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_goadesign_goagen_secured_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SigninRequest); i {
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
		file_goadesign_goagen_secured_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SigninResponse); i {
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
		file_goadesign_goagen_secured_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecureRequest); i {
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
		file_goadesign_goagen_secured_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecureResponse); i {
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
		file_goadesign_goagen_secured_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoublySecureRequest); i {
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
		file_goadesign_goagen_secured_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoublySecureResponse); i {
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
		file_goadesign_goagen_secured_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlsoDoublySecureRequest); i {
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
		file_goadesign_goagen_secured_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlsoDoublySecureResponse); i {
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
			RawDescriptor: file_goadesign_goagen_secured_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goadesign_goagen_secured_service_proto_goTypes,
		DependencyIndexes: file_goadesign_goagen_secured_service_proto_depIdxs,
		MessageInfos:      file_goadesign_goagen_secured_service_proto_msgTypes,
	}.Build()
	File_goadesign_goagen_secured_service_proto = out.File
	file_goadesign_goagen_secured_service_proto_rawDesc = nil
	file_goadesign_goagen_secured_service_proto_goTypes = nil
	file_goadesign_goagen_secured_service_proto_depIdxs = nil
}
