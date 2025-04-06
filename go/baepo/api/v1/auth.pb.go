// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: baepo/api/v1/auth.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuthRegisterRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	FirstName     string                 `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName      string                 `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthRegisterRequest) Reset() {
	*x = AuthRegisterRequest{}
	mi := &file_baepo_api_v1_auth_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthRegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRegisterRequest) ProtoMessage() {}

func (x *AuthRegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_api_v1_auth_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRegisterRequest.ProtoReflect.Descriptor instead.
func (*AuthRegisterRequest) Descriptor() ([]byte, []int) {
	return file_baepo_api_v1_auth_proto_rawDescGZIP(), []int{0}
}

func (x *AuthRegisterRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AuthRegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AuthRegisterRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *AuthRegisterRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

type AuthRegisterResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	SecretKey     string                 `protobuf:"bytes,2,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthRegisterResponse) Reset() {
	*x = AuthRegisterResponse{}
	mi := &file_baepo_api_v1_auth_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthRegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRegisterResponse) ProtoMessage() {}

func (x *AuthRegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_api_v1_auth_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRegisterResponse.ProtoReflect.Descriptor instead.
func (*AuthRegisterResponse) Descriptor() ([]byte, []int) {
	return file_baepo_api_v1_auth_proto_rawDescGZIP(), []int{1}
}

func (x *AuthRegisterResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AuthRegisterResponse) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

type AuthLoginRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthLoginRequest) Reset() {
	*x = AuthLoginRequest{}
	mi := &file_baepo_api_v1_auth_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthLoginRequest) ProtoMessage() {}

func (x *AuthLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_api_v1_auth_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthLoginRequest.ProtoReflect.Descriptor instead.
func (*AuthLoginRequest) Descriptor() ([]byte, []int) {
	return file_baepo_api_v1_auth_proto_rawDescGZIP(), []int{2}
}

func (x *AuthLoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AuthLoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type AuthLoginResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	SecretKey     string                 `protobuf:"bytes,2,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthLoginResponse) Reset() {
	*x = AuthLoginResponse{}
	mi := &file_baepo_api_v1_auth_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthLoginResponse) ProtoMessage() {}

func (x *AuthLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_api_v1_auth_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthLoginResponse.ProtoReflect.Descriptor instead.
func (*AuthLoginResponse) Descriptor() ([]byte, []int) {
	return file_baepo_api_v1_auth_proto_rawDescGZIP(), []int{3}
}

func (x *AuthLoginResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AuthLoginResponse) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

var File_baepo_api_v1_auth_proto protoreflect.FileDescriptor

const file_baepo_api_v1_auth_proto_rawDesc = "" +
	"\n" +
	"\x17baepo/api/v1/auth.proto\x12\fbaepo.api.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\x83\x01\n" +
	"\x13AuthRegisterRequest\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\x12\x1d\n" +
	"\n" +
	"first_name\x18\x03 \x01(\tR\tfirstName\x12\x1b\n" +
	"\tlast_name\x18\x04 \x01(\tR\blastName\"N\n" +
	"\x14AuthRegisterResponse\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x1d\n" +
	"\n" +
	"secret_key\x18\x02 \x01(\tR\tsecretKey\"D\n" +
	"\x10AuthLoginRequest\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\"K\n" +
	"\x11AuthLoginResponse\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x1d\n" +
	"\n" +
	"secret_key\x18\x02 \x01(\tR\tsecretKey2\xaa\x01\n" +
	"\vAuthService\x12Q\n" +
	"\bRegister\x12!.baepo.api.v1.AuthRegisterRequest\x1a\".baepo.api.v1.AuthRegisterResponse\x12H\n" +
	"\x05Login\x12\x1e.baepo.api.v1.AuthLoginRequest\x1a\x1f.baepo.api.v1.AuthLoginResponseB4Z2github.com/baepo-cloud/baepo-proto/go/baepo/api/v1b\x06proto3"

var (
	file_baepo_api_v1_auth_proto_rawDescOnce sync.Once
	file_baepo_api_v1_auth_proto_rawDescData []byte
)

func file_baepo_api_v1_auth_proto_rawDescGZIP() []byte {
	file_baepo_api_v1_auth_proto_rawDescOnce.Do(func() {
		file_baepo_api_v1_auth_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_baepo_api_v1_auth_proto_rawDesc), len(file_baepo_api_v1_auth_proto_rawDesc)))
	})
	return file_baepo_api_v1_auth_proto_rawDescData
}

var file_baepo_api_v1_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_baepo_api_v1_auth_proto_goTypes = []any{
	(*AuthRegisterRequest)(nil),  // 0: baepo.api.v1.AuthRegisterRequest
	(*AuthRegisterResponse)(nil), // 1: baepo.api.v1.AuthRegisterResponse
	(*AuthLoginRequest)(nil),     // 2: baepo.api.v1.AuthLoginRequest
	(*AuthLoginResponse)(nil),    // 3: baepo.api.v1.AuthLoginResponse
}
var file_baepo_api_v1_auth_proto_depIdxs = []int32{
	0, // 0: baepo.api.v1.AuthService.Register:input_type -> baepo.api.v1.AuthRegisterRequest
	2, // 1: baepo.api.v1.AuthService.Login:input_type -> baepo.api.v1.AuthLoginRequest
	1, // 2: baepo.api.v1.AuthService.Register:output_type -> baepo.api.v1.AuthRegisterResponse
	3, // 3: baepo.api.v1.AuthService.Login:output_type -> baepo.api.v1.AuthLoginResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_baepo_api_v1_auth_proto_init() }
func file_baepo_api_v1_auth_proto_init() {
	if File_baepo_api_v1_auth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_baepo_api_v1_auth_proto_rawDesc), len(file_baepo_api_v1_auth_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_baepo_api_v1_auth_proto_goTypes,
		DependencyIndexes: file_baepo_api_v1_auth_proto_depIdxs,
		MessageInfos:      file_baepo_api_v1_auth_proto_msgTypes,
	}.Build()
	File_baepo_api_v1_auth_proto = out.File
	file_baepo_api_v1_auth_proto_goTypes = nil
	file_baepo_api_v1_auth_proto_depIdxs = nil
}
