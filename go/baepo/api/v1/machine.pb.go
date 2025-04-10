// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: baepo/api/v1/machine.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type MachineStatus int32

const (
	MachineStatus_MachineStatus_Unknown     MachineStatus = 0
	MachineStatus_MachineStatus_Scheduling  MachineStatus = 1
	MachineStatus_MachineStatus_Starting    MachineStatus = 2
	MachineStatus_MachineStatus_Running     MachineStatus = 3
	MachineStatus_MachineStatus_Terminating MachineStatus = 4
	MachineStatus_MachineStatus_Terminated  MachineStatus = 5
)

// Enum value maps for MachineStatus.
var (
	MachineStatus_name = map[int32]string{
		0: "MachineStatus_Unknown",
		1: "MachineStatus_Scheduling",
		2: "MachineStatus_Starting",
		3: "MachineStatus_Running",
		4: "MachineStatus_Terminating",
		5: "MachineStatus_Terminated",
	}
	MachineStatus_value = map[string]int32{
		"MachineStatus_Unknown":     0,
		"MachineStatus_Scheduling":  1,
		"MachineStatus_Starting":    2,
		"MachineStatus_Running":     3,
		"MachineStatus_Terminating": 4,
		"MachineStatus_Terminated":  5,
	}
)

func (x MachineStatus) Enum() *MachineStatus {
	p := new(MachineStatus)
	*p = x
	return p
}

func (x MachineStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MachineStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_baepo_api_v1_machine_proto_enumTypes[0].Descriptor()
}

func (MachineStatus) Type() protoreflect.EnumType {
	return &file_baepo_api_v1_machine_proto_enumTypes[0]
}

func (x MachineStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MachineStatus.Descriptor instead.
func (MachineStatus) EnumDescriptor() ([]byte, []int) {
	return file_baepo_api_v1_machine_proto_rawDescGZIP(), []int{0}
}

type Machine struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	Id                 string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name               *string                `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Status             MachineStatus          `protobuf:"varint,3,opt,name=status,proto3,enum=baepo.api.v1.MachineStatus" json:"status,omitempty"`
	NodeId             *string                `protobuf:"bytes,4,opt,name=node_id,json=nodeId,proto3,oneof" json:"node_id,omitempty"`
	Spec               *MachineSpec           `protobuf:"bytes,5,opt,name=spec,proto3" json:"spec,omitempty"`
	Timeout            *uint32                `protobuf:"varint,6,opt,name=timeout,proto3,oneof" json:"timeout,omitempty"`
	StartedAt          *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=started_at,json=startedAt,proto3,oneof" json:"started_at,omitempty"`
	ExpiresAt          *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=expires_at,json=expiresAt,proto3,oneof" json:"expires_at,omitempty"`
	TerminatedAt       *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=terminated_at,json=terminatedAt,proto3,oneof" json:"terminated_at,omitempty"`
	TerminationCause   *string                `protobuf:"bytes,10,opt,name=termination_cause,json=terminationCause,proto3,oneof" json:"termination_cause,omitempty"`
	TerminationDetails *string                `protobuf:"bytes,11,opt,name=termination_details,json=terminationDetails,proto3,oneof" json:"termination_details,omitempty"`
	Metadata           map[string]string      `protobuf:"bytes,12,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	CreatedAt          *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt          *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	WorkspaceId        string                 `protobuf:"bytes,15,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *Machine) Reset() {
	*x = Machine{}
	mi := &file_baepo_api_v1_machine_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Machine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Machine) ProtoMessage() {}

func (x *Machine) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_api_v1_machine_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Machine.ProtoReflect.Descriptor instead.
func (*Machine) Descriptor() ([]byte, []int) {
	return file_baepo_api_v1_machine_proto_rawDescGZIP(), []int{0}
}

func (x *Machine) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Machine) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Machine) GetStatus() MachineStatus {
	if x != nil {
		return x.Status
	}
	return MachineStatus_MachineStatus_Unknown
}

func (x *Machine) GetNodeId() string {
	if x != nil && x.NodeId != nil {
		return *x.NodeId
	}
	return ""
}

func (x *Machine) GetSpec() *MachineSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Machine) GetTimeout() uint32 {
	if x != nil && x.Timeout != nil {
		return *x.Timeout
	}
	return 0
}

func (x *Machine) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *Machine) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

func (x *Machine) GetTerminatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.TerminatedAt
	}
	return nil
}

func (x *Machine) GetTerminationCause() string {
	if x != nil && x.TerminationCause != nil {
		return *x.TerminationCause
	}
	return ""
}

func (x *Machine) GetTerminationDetails() string {
	if x != nil && x.TerminationDetails != nil {
		return *x.TerminationDetails
	}
	return ""
}

func (x *Machine) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Machine) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Machine) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Machine) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

type MachineSpec struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	VCpus         uint32                 `protobuf:"varint,1,opt,name=v_cpus,json=vCpus,proto3" json:"v_cpus,omitempty"`
	MemoryMb      uint64                 `protobuf:"varint,2,opt,name=memory_mb,json=memoryMb,proto3" json:"memory_mb,omitempty"`
	Env           map[string]string      `protobuf:"bytes,3,rep,name=env,proto3" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Image         string                 `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	Command       []string               `protobuf:"bytes,5,rep,name=command,proto3" json:"command,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MachineSpec) Reset() {
	*x = MachineSpec{}
	mi := &file_baepo_api_v1_machine_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MachineSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineSpec) ProtoMessage() {}

func (x *MachineSpec) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_api_v1_machine_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineSpec.ProtoReflect.Descriptor instead.
func (*MachineSpec) Descriptor() ([]byte, []int) {
	return file_baepo_api_v1_machine_proto_rawDescGZIP(), []int{1}
}

func (x *MachineSpec) GetVCpus() uint32 {
	if x != nil {
		return x.VCpus
	}
	return 0
}

func (x *MachineSpec) GetMemoryMb() uint64 {
	if x != nil {
		return x.MemoryMb
	}
	return 0
}

func (x *MachineSpec) GetEnv() map[string]string {
	if x != nil {
		return x.Env
	}
	return nil
}

func (x *MachineSpec) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *MachineSpec) GetCommand() []string {
	if x != nil {
		return x.Command
	}
	return nil
}

type MachineListRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	WorkspaceId   string                 `protobuf:"bytes,1,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MachineListRequest) Reset() {
	*x = MachineListRequest{}
	mi := &file_baepo_api_v1_machine_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MachineListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineListRequest) ProtoMessage() {}

func (x *MachineListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_api_v1_machine_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineListRequest.ProtoReflect.Descriptor instead.
func (*MachineListRequest) Descriptor() ([]byte, []int) {
	return file_baepo_api_v1_machine_proto_rawDescGZIP(), []int{2}
}

func (x *MachineListRequest) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

type MachineListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Machines      []*Machine             `protobuf:"bytes,1,rep,name=machines,proto3" json:"machines,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MachineListResponse) Reset() {
	*x = MachineListResponse{}
	mi := &file_baepo_api_v1_machine_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MachineListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineListResponse) ProtoMessage() {}

func (x *MachineListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_api_v1_machine_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineListResponse.ProtoReflect.Descriptor instead.
func (*MachineListResponse) Descriptor() ([]byte, []int) {
	return file_baepo_api_v1_machine_proto_rawDescGZIP(), []int{3}
}

func (x *MachineListResponse) GetMachines() []*Machine {
	if x != nil {
		return x.Machines
	}
	return nil
}

type MachineFindByIdRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MachineId     string                 `protobuf:"bytes,1,opt,name=machine_id,json=machineId,proto3" json:"machine_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MachineFindByIdRequest) Reset() {
	*x = MachineFindByIdRequest{}
	mi := &file_baepo_api_v1_machine_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MachineFindByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineFindByIdRequest) ProtoMessage() {}

func (x *MachineFindByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_api_v1_machine_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineFindByIdRequest.ProtoReflect.Descriptor instead.
func (*MachineFindByIdRequest) Descriptor() ([]byte, []int) {
	return file_baepo_api_v1_machine_proto_rawDescGZIP(), []int{4}
}

func (x *MachineFindByIdRequest) GetMachineId() string {
	if x != nil {
		return x.MachineId
	}
	return ""
}

type MachineFindByIdResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Machine       *Machine               `protobuf:"bytes,1,opt,name=machine,proto3" json:"machine,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MachineFindByIdResponse) Reset() {
	*x = MachineFindByIdResponse{}
	mi := &file_baepo_api_v1_machine_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MachineFindByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineFindByIdResponse) ProtoMessage() {}

func (x *MachineFindByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_api_v1_machine_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineFindByIdResponse.ProtoReflect.Descriptor instead.
func (*MachineFindByIdResponse) Descriptor() ([]byte, []int) {
	return file_baepo_api_v1_machine_proto_rawDescGZIP(), []int{5}
}

func (x *MachineFindByIdResponse) GetMachine() *Machine {
	if x != nil {
		return x.Machine
	}
	return nil
}

type MachineCreateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	WorkspaceId   string                 `protobuf:"bytes,1,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	Name          *string                `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Timeout       *uint32                `protobuf:"varint,3,opt,name=timeout,proto3,oneof" json:"timeout,omitempty"`
	Spec          *MachineSpec           `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	Metadata      map[string]string      `protobuf:"bytes,5,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MachineCreateRequest) Reset() {
	*x = MachineCreateRequest{}
	mi := &file_baepo_api_v1_machine_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MachineCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineCreateRequest) ProtoMessage() {}

func (x *MachineCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_api_v1_machine_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineCreateRequest.ProtoReflect.Descriptor instead.
func (*MachineCreateRequest) Descriptor() ([]byte, []int) {
	return file_baepo_api_v1_machine_proto_rawDescGZIP(), []int{6}
}

func (x *MachineCreateRequest) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *MachineCreateRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *MachineCreateRequest) GetTimeout() uint32 {
	if x != nil && x.Timeout != nil {
		return *x.Timeout
	}
	return 0
}

func (x *MachineCreateRequest) GetSpec() *MachineSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *MachineCreateRequest) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type MachineCreateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Machine       *Machine               `protobuf:"bytes,1,opt,name=machine,proto3" json:"machine,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MachineCreateResponse) Reset() {
	*x = MachineCreateResponse{}
	mi := &file_baepo_api_v1_machine_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MachineCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineCreateResponse) ProtoMessage() {}

func (x *MachineCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_api_v1_machine_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineCreateResponse.ProtoReflect.Descriptor instead.
func (*MachineCreateResponse) Descriptor() ([]byte, []int) {
	return file_baepo_api_v1_machine_proto_rawDescGZIP(), []int{7}
}

func (x *MachineCreateResponse) GetMachine() *Machine {
	if x != nil {
		return x.Machine
	}
	return nil
}

type MachineTerminateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MachineId     string                 `protobuf:"bytes,1,opt,name=machine_id,json=machineId,proto3" json:"machine_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MachineTerminateRequest) Reset() {
	*x = MachineTerminateRequest{}
	mi := &file_baepo_api_v1_machine_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MachineTerminateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineTerminateRequest) ProtoMessage() {}

func (x *MachineTerminateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_api_v1_machine_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineTerminateRequest.ProtoReflect.Descriptor instead.
func (*MachineTerminateRequest) Descriptor() ([]byte, []int) {
	return file_baepo_api_v1_machine_proto_rawDescGZIP(), []int{8}
}

func (x *MachineTerminateRequest) GetMachineId() string {
	if x != nil {
		return x.MachineId
	}
	return ""
}

type MachineTerminateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Machine       *Machine               `protobuf:"bytes,1,opt,name=machine,proto3" json:"machine,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MachineTerminateResponse) Reset() {
	*x = MachineTerminateResponse{}
	mi := &file_baepo_api_v1_machine_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MachineTerminateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineTerminateResponse) ProtoMessage() {}

func (x *MachineTerminateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_api_v1_machine_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineTerminateResponse.ProtoReflect.Descriptor instead.
func (*MachineTerminateResponse) Descriptor() ([]byte, []int) {
	return file_baepo_api_v1_machine_proto_rawDescGZIP(), []int{9}
}

func (x *MachineTerminateResponse) GetMachine() *Machine {
	if x != nil {
		return x.Machine
	}
	return nil
}

var File_baepo_api_v1_machine_proto protoreflect.FileDescriptor

const file_baepo_api_v1_machine_proto_rawDesc = "" +
	"\n" +
	"\x1abaepo/api/v1/machine.proto\x12\fbaepo.api.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\x97\a\n" +
	"\aMachine\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x17\n" +
	"\x04name\x18\x02 \x01(\tH\x00R\x04name\x88\x01\x01\x123\n" +
	"\x06status\x18\x03 \x01(\x0e2\x1b.baepo.api.v1.MachineStatusR\x06status\x12\x1c\n" +
	"\anode_id\x18\x04 \x01(\tH\x01R\x06nodeId\x88\x01\x01\x12-\n" +
	"\x04spec\x18\x05 \x01(\v2\x19.baepo.api.v1.MachineSpecR\x04spec\x12\x1d\n" +
	"\atimeout\x18\x06 \x01(\rH\x02R\atimeout\x88\x01\x01\x12>\n" +
	"\n" +
	"started_at\x18\a \x01(\v2\x1a.google.protobuf.TimestampH\x03R\tstartedAt\x88\x01\x01\x12>\n" +
	"\n" +
	"expires_at\x18\b \x01(\v2\x1a.google.protobuf.TimestampH\x04R\texpiresAt\x88\x01\x01\x12D\n" +
	"\rterminated_at\x18\t \x01(\v2\x1a.google.protobuf.TimestampH\x05R\fterminatedAt\x88\x01\x01\x120\n" +
	"\x11termination_cause\x18\n" +
	" \x01(\tH\x06R\x10terminationCause\x88\x01\x01\x124\n" +
	"\x13termination_details\x18\v \x01(\tH\aR\x12terminationDetails\x88\x01\x01\x12?\n" +
	"\bmetadata\x18\f \x03(\v2#.baepo.api.v1.Machine.MetadataEntryR\bmetadata\x129\n" +
	"\n" +
	"created_at\x18\r \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\x0e \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\x12!\n" +
	"\fworkspace_id\x18\x0f \x01(\tR\vworkspaceId\x1a;\n" +
	"\rMetadataEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01B\a\n" +
	"\x05_nameB\n" +
	"\n" +
	"\b_node_idB\n" +
	"\n" +
	"\b_timeoutB\r\n" +
	"\v_started_atB\r\n" +
	"\v_expires_atB\x10\n" +
	"\x0e_terminated_atB\x14\n" +
	"\x12_termination_causeB\x16\n" +
	"\x14_termination_details\"\xdf\x01\n" +
	"\vMachineSpec\x12\x15\n" +
	"\x06v_cpus\x18\x01 \x01(\rR\x05vCpus\x12\x1b\n" +
	"\tmemory_mb\x18\x02 \x01(\x04R\bmemoryMb\x124\n" +
	"\x03env\x18\x03 \x03(\v2\".baepo.api.v1.MachineSpec.EnvEntryR\x03env\x12\x14\n" +
	"\x05image\x18\x04 \x01(\tR\x05image\x12\x18\n" +
	"\acommand\x18\x05 \x03(\tR\acommand\x1a6\n" +
	"\bEnvEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"7\n" +
	"\x12MachineListRequest\x12!\n" +
	"\fworkspace_id\x18\x01 \x01(\tR\vworkspaceId\"H\n" +
	"\x13MachineListResponse\x121\n" +
	"\bmachines\x18\x01 \x03(\v2\x15.baepo.api.v1.MachineR\bmachines\"7\n" +
	"\x16MachineFindByIdRequest\x12\x1d\n" +
	"\n" +
	"machine_id\x18\x01 \x01(\tR\tmachineId\"J\n" +
	"\x17MachineFindByIdResponse\x12/\n" +
	"\amachine\x18\x01 \x01(\v2\x15.baepo.api.v1.MachineR\amachine\"\xc0\x02\n" +
	"\x14MachineCreateRequest\x12!\n" +
	"\fworkspace_id\x18\x01 \x01(\tR\vworkspaceId\x12\x17\n" +
	"\x04name\x18\x02 \x01(\tH\x00R\x04name\x88\x01\x01\x12\x1d\n" +
	"\atimeout\x18\x03 \x01(\rH\x01R\atimeout\x88\x01\x01\x12-\n" +
	"\x04spec\x18\x04 \x01(\v2\x19.baepo.api.v1.MachineSpecR\x04spec\x12L\n" +
	"\bmetadata\x18\x05 \x03(\v20.baepo.api.v1.MachineCreateRequest.MetadataEntryR\bmetadata\x1a;\n" +
	"\rMetadataEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01B\a\n" +
	"\x05_nameB\n" +
	"\n" +
	"\b_timeout\"H\n" +
	"\x15MachineCreateResponse\x12/\n" +
	"\amachine\x18\x01 \x01(\v2\x15.baepo.api.v1.MachineR\amachine\"8\n" +
	"\x17MachineTerminateRequest\x12\x1d\n" +
	"\n" +
	"machine_id\x18\x01 \x01(\tR\tmachineId\"K\n" +
	"\x18MachineTerminateResponse\x12/\n" +
	"\amachine\x18\x01 \x01(\v2\x15.baepo.api.v1.MachineR\amachine*\xbc\x01\n" +
	"\rMachineStatus\x12\x19\n" +
	"\x15MachineStatus_Unknown\x10\x00\x12\x1c\n" +
	"\x18MachineStatus_Scheduling\x10\x01\x12\x1a\n" +
	"\x16MachineStatus_Starting\x10\x02\x12\x19\n" +
	"\x15MachineStatus_Running\x10\x03\x12\x1d\n" +
	"\x19MachineStatus_Terminating\x10\x04\x12\x1c\n" +
	"\x18MachineStatus_Terminated\x10\x052\xe5\x02\n" +
	"\x0eMachineService\x12K\n" +
	"\x04List\x12 .baepo.api.v1.MachineListRequest\x1a!.baepo.api.v1.MachineListResponse\x12W\n" +
	"\bFindById\x12$.baepo.api.v1.MachineFindByIdRequest\x1a%.baepo.api.v1.MachineFindByIdResponse\x12Q\n" +
	"\x06Create\x12\".baepo.api.v1.MachineCreateRequest\x1a#.baepo.api.v1.MachineCreateResponse\x12Z\n" +
	"\tTerminate\x12%.baepo.api.v1.MachineTerminateRequest\x1a&.baepo.api.v1.MachineTerminateResponseB4Z2github.com/baepo-cloud/baepo-proto/go/baepo/api/v1b\x06proto3"

var (
	file_baepo_api_v1_machine_proto_rawDescOnce sync.Once
	file_baepo_api_v1_machine_proto_rawDescData []byte
)

func file_baepo_api_v1_machine_proto_rawDescGZIP() []byte {
	file_baepo_api_v1_machine_proto_rawDescOnce.Do(func() {
		file_baepo_api_v1_machine_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_baepo_api_v1_machine_proto_rawDesc), len(file_baepo_api_v1_machine_proto_rawDesc)))
	})
	return file_baepo_api_v1_machine_proto_rawDescData
}

var file_baepo_api_v1_machine_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_baepo_api_v1_machine_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_baepo_api_v1_machine_proto_goTypes = []any{
	(MachineStatus)(0),               // 0: baepo.api.v1.MachineStatus
	(*Machine)(nil),                  // 1: baepo.api.v1.Machine
	(*MachineSpec)(nil),              // 2: baepo.api.v1.MachineSpec
	(*MachineListRequest)(nil),       // 3: baepo.api.v1.MachineListRequest
	(*MachineListResponse)(nil),      // 4: baepo.api.v1.MachineListResponse
	(*MachineFindByIdRequest)(nil),   // 5: baepo.api.v1.MachineFindByIdRequest
	(*MachineFindByIdResponse)(nil),  // 6: baepo.api.v1.MachineFindByIdResponse
	(*MachineCreateRequest)(nil),     // 7: baepo.api.v1.MachineCreateRequest
	(*MachineCreateResponse)(nil),    // 8: baepo.api.v1.MachineCreateResponse
	(*MachineTerminateRequest)(nil),  // 9: baepo.api.v1.MachineTerminateRequest
	(*MachineTerminateResponse)(nil), // 10: baepo.api.v1.MachineTerminateResponse
	nil,                              // 11: baepo.api.v1.Machine.MetadataEntry
	nil,                              // 12: baepo.api.v1.MachineSpec.EnvEntry
	nil,                              // 13: baepo.api.v1.MachineCreateRequest.MetadataEntry
	(*timestamppb.Timestamp)(nil),    // 14: google.protobuf.Timestamp
}
var file_baepo_api_v1_machine_proto_depIdxs = []int32{
	0,  // 0: baepo.api.v1.Machine.status:type_name -> baepo.api.v1.MachineStatus
	2,  // 1: baepo.api.v1.Machine.spec:type_name -> baepo.api.v1.MachineSpec
	14, // 2: baepo.api.v1.Machine.started_at:type_name -> google.protobuf.Timestamp
	14, // 3: baepo.api.v1.Machine.expires_at:type_name -> google.protobuf.Timestamp
	14, // 4: baepo.api.v1.Machine.terminated_at:type_name -> google.protobuf.Timestamp
	11, // 5: baepo.api.v1.Machine.metadata:type_name -> baepo.api.v1.Machine.MetadataEntry
	14, // 6: baepo.api.v1.Machine.created_at:type_name -> google.protobuf.Timestamp
	14, // 7: baepo.api.v1.Machine.updated_at:type_name -> google.protobuf.Timestamp
	12, // 8: baepo.api.v1.MachineSpec.env:type_name -> baepo.api.v1.MachineSpec.EnvEntry
	1,  // 9: baepo.api.v1.MachineListResponse.machines:type_name -> baepo.api.v1.Machine
	1,  // 10: baepo.api.v1.MachineFindByIdResponse.machine:type_name -> baepo.api.v1.Machine
	2,  // 11: baepo.api.v1.MachineCreateRequest.spec:type_name -> baepo.api.v1.MachineSpec
	13, // 12: baepo.api.v1.MachineCreateRequest.metadata:type_name -> baepo.api.v1.MachineCreateRequest.MetadataEntry
	1,  // 13: baepo.api.v1.MachineCreateResponse.machine:type_name -> baepo.api.v1.Machine
	1,  // 14: baepo.api.v1.MachineTerminateResponse.machine:type_name -> baepo.api.v1.Machine
	3,  // 15: baepo.api.v1.MachineService.List:input_type -> baepo.api.v1.MachineListRequest
	5,  // 16: baepo.api.v1.MachineService.FindById:input_type -> baepo.api.v1.MachineFindByIdRequest
	7,  // 17: baepo.api.v1.MachineService.Create:input_type -> baepo.api.v1.MachineCreateRequest
	9,  // 18: baepo.api.v1.MachineService.Terminate:input_type -> baepo.api.v1.MachineTerminateRequest
	4,  // 19: baepo.api.v1.MachineService.List:output_type -> baepo.api.v1.MachineListResponse
	6,  // 20: baepo.api.v1.MachineService.FindById:output_type -> baepo.api.v1.MachineFindByIdResponse
	8,  // 21: baepo.api.v1.MachineService.Create:output_type -> baepo.api.v1.MachineCreateResponse
	10, // 22: baepo.api.v1.MachineService.Terminate:output_type -> baepo.api.v1.MachineTerminateResponse
	19, // [19:23] is the sub-list for method output_type
	15, // [15:19] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_baepo_api_v1_machine_proto_init() }
func file_baepo_api_v1_machine_proto_init() {
	if File_baepo_api_v1_machine_proto != nil {
		return
	}
	file_baepo_api_v1_machine_proto_msgTypes[0].OneofWrappers = []any{}
	file_baepo_api_v1_machine_proto_msgTypes[6].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_baepo_api_v1_machine_proto_rawDesc), len(file_baepo_api_v1_machine_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_baepo_api_v1_machine_proto_goTypes,
		DependencyIndexes: file_baepo_api_v1_machine_proto_depIdxs,
		EnumInfos:         file_baepo_api_v1_machine_proto_enumTypes,
		MessageInfos:      file_baepo_api_v1_machine_proto_msgTypes,
	}.Build()
	File_baepo_api_v1_machine_proto = out.File
	file_baepo_api_v1_machine_proto_goTypes = nil
	file_baepo_api_v1_machine_proto_depIdxs = nil
}
