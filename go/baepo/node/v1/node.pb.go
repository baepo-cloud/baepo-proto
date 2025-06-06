// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: baepo/node/v1/node.proto

package nodev1pb

import (
	v1 "github.com/baepo-cloud/baepo-proto/go/baepo/core/v1"
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

type Machine struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MachineId     string                 `protobuf:"bytes,1,opt,name=machine_id,json=machineId,proto3" json:"machine_id,omitempty"`
	State         v1.MachineState        `protobuf:"varint,2,opt,name=state,proto3,enum=baepo.core.v1.MachineState" json:"state,omitempty"`
	DesiredState  v1.MachineDesiredState `protobuf:"varint,3,opt,name=desired_state,json=desiredState,proto3,enum=baepo.core.v1.MachineDesiredState" json:"desired_state,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Machine) Reset() {
	*x = Machine{}
	mi := &file_baepo_node_v1_node_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Machine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Machine) ProtoMessage() {}

func (x *Machine) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_node_v1_node_proto_msgTypes[0]
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
	return file_baepo_node_v1_node_proto_rawDescGZIP(), []int{0}
}

func (x *Machine) GetMachineId() string {
	if x != nil {
		return x.MachineId
	}
	return ""
}

func (x *Machine) GetState() v1.MachineState {
	if x != nil {
		return x.State
	}
	return v1.MachineState(0)
}

func (x *Machine) GetDesiredState() v1.MachineDesiredState {
	if x != nil {
		return x.DesiredState
	}
	return v1.MachineDesiredState(0)
}

type NodeListMachinesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NodeListMachinesRequest) Reset() {
	*x = NodeListMachinesRequest{}
	mi := &file_baepo_node_v1_node_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NodeListMachinesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeListMachinesRequest) ProtoMessage() {}

func (x *NodeListMachinesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_node_v1_node_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeListMachinesRequest.ProtoReflect.Descriptor instead.
func (*NodeListMachinesRequest) Descriptor() ([]byte, []int) {
	return file_baepo_node_v1_node_proto_rawDescGZIP(), []int{1}
}

type NodeListMachinesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Machines      []*Machine             `protobuf:"bytes,1,rep,name=machines,proto3" json:"machines,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NodeListMachinesResponse) Reset() {
	*x = NodeListMachinesResponse{}
	mi := &file_baepo_node_v1_node_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NodeListMachinesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeListMachinesResponse) ProtoMessage() {}

func (x *NodeListMachinesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_node_v1_node_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeListMachinesResponse.ProtoReflect.Descriptor instead.
func (*NodeListMachinesResponse) Descriptor() ([]byte, []int) {
	return file_baepo_node_v1_node_proto_rawDescGZIP(), []int{2}
}

func (x *NodeListMachinesResponse) GetMachines() []*Machine {
	if x != nil {
		return x.Machines
	}
	return nil
}

type NodeGetMachineRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MachineId     string                 `protobuf:"bytes,1,opt,name=machine_id,json=machineId,proto3" json:"machine_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NodeGetMachineRequest) Reset() {
	*x = NodeGetMachineRequest{}
	mi := &file_baepo_node_v1_node_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NodeGetMachineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeGetMachineRequest) ProtoMessage() {}

func (x *NodeGetMachineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_node_v1_node_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeGetMachineRequest.ProtoReflect.Descriptor instead.
func (*NodeGetMachineRequest) Descriptor() ([]byte, []int) {
	return file_baepo_node_v1_node_proto_rawDescGZIP(), []int{3}
}

func (x *NodeGetMachineRequest) GetMachineId() string {
	if x != nil {
		return x.MachineId
	}
	return ""
}

type NodeGetMachineResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Machine       *Machine               `protobuf:"bytes,1,opt,name=machine,proto3" json:"machine,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NodeGetMachineResponse) Reset() {
	*x = NodeGetMachineResponse{}
	mi := &file_baepo_node_v1_node_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NodeGetMachineResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeGetMachineResponse) ProtoMessage() {}

func (x *NodeGetMachineResponse) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_node_v1_node_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeGetMachineResponse.ProtoReflect.Descriptor instead.
func (*NodeGetMachineResponse) Descriptor() ([]byte, []int) {
	return file_baepo_node_v1_node_proto_rawDescGZIP(), []int{4}
}

func (x *NodeGetMachineResponse) GetMachine() *Machine {
	if x != nil {
		return x.Machine
	}
	return nil
}

type NodeGetMachineLogsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MachineId     string                 `protobuf:"bytes,1,opt,name=machine_id,json=machineId,proto3" json:"machine_id,omitempty"`
	Follow        bool                   `protobuf:"varint,2,opt,name=follow,proto3" json:"follow,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NodeGetMachineLogsRequest) Reset() {
	*x = NodeGetMachineLogsRequest{}
	mi := &file_baepo_node_v1_node_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NodeGetMachineLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeGetMachineLogsRequest) ProtoMessage() {}

func (x *NodeGetMachineLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_node_v1_node_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeGetMachineLogsRequest.ProtoReflect.Descriptor instead.
func (*NodeGetMachineLogsRequest) Descriptor() ([]byte, []int) {
	return file_baepo_node_v1_node_proto_rawDescGZIP(), []int{5}
}

func (x *NodeGetMachineLogsRequest) GetMachineId() string {
	if x != nil {
		return x.MachineId
	}
	return ""
}

func (x *NodeGetMachineLogsRequest) GetFollow() bool {
	if x != nil {
		return x.Follow
	}
	return false
}

type NodeGetMachineLogsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Content       []byte                 `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NodeGetMachineLogsResponse) Reset() {
	*x = NodeGetMachineLogsResponse{}
	mi := &file_baepo_node_v1_node_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NodeGetMachineLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeGetMachineLogsResponse) ProtoMessage() {}

func (x *NodeGetMachineLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_node_v1_node_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeGetMachineLogsResponse.ProtoReflect.Descriptor instead.
func (*NodeGetMachineLogsResponse) Descriptor() ([]byte, []int) {
	return file_baepo_node_v1_node_proto_rawDescGZIP(), []int{6}
}

func (x *NodeGetMachineLogsResponse) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type NodeGetContainerLogsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MachineId     string                 `protobuf:"bytes,1,opt,name=machine_id,json=machineId,proto3" json:"machine_id,omitempty"`
	Container     *string                `protobuf:"bytes,2,opt,name=container,proto3,oneof" json:"container,omitempty"`
	Follow        bool                   `protobuf:"varint,3,opt,name=follow,proto3" json:"follow,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NodeGetContainerLogsRequest) Reset() {
	*x = NodeGetContainerLogsRequest{}
	mi := &file_baepo_node_v1_node_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NodeGetContainerLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeGetContainerLogsRequest) ProtoMessage() {}

func (x *NodeGetContainerLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_node_v1_node_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeGetContainerLogsRequest.ProtoReflect.Descriptor instead.
func (*NodeGetContainerLogsRequest) Descriptor() ([]byte, []int) {
	return file_baepo_node_v1_node_proto_rawDescGZIP(), []int{7}
}

func (x *NodeGetContainerLogsRequest) GetMachineId() string {
	if x != nil {
		return x.MachineId
	}
	return ""
}

func (x *NodeGetContainerLogsRequest) GetContainer() string {
	if x != nil && x.Container != nil {
		return *x.Container
	}
	return ""
}

func (x *NodeGetContainerLogsRequest) GetFollow() bool {
	if x != nil {
		return x.Follow
	}
	return false
}

type NodeGetContainerLogsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ContainerId   string                 `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	ContainerName string                 `protobuf:"bytes,2,opt,name=container_name,json=containerName,proto3" json:"container_name,omitempty"`
	Error         bool                   `protobuf:"varint,3,opt,name=error,proto3" json:"error,omitempty"`
	Content       []byte                 `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Timestamp     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NodeGetContainerLogsResponse) Reset() {
	*x = NodeGetContainerLogsResponse{}
	mi := &file_baepo_node_v1_node_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NodeGetContainerLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeGetContainerLogsResponse) ProtoMessage() {}

func (x *NodeGetContainerLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_node_v1_node_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeGetContainerLogsResponse.ProtoReflect.Descriptor instead.
func (*NodeGetContainerLogsResponse) Descriptor() ([]byte, []int) {
	return file_baepo_node_v1_node_proto_rawDescGZIP(), []int{8}
}

func (x *NodeGetContainerLogsResponse) GetContainerId() string {
	if x != nil {
		return x.ContainerId
	}
	return ""
}

func (x *NodeGetContainerLogsResponse) GetContainerName() string {
	if x != nil {
		return x.ContainerName
	}
	return ""
}

func (x *NodeGetContainerLogsResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *NodeGetContainerLogsResponse) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *NodeGetContainerLogsResponse) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

var File_baepo_node_v1_node_proto protoreflect.FileDescriptor

const file_baepo_node_v1_node_proto_rawDesc = "" +
	"\n" +
	"\x18baepo/node/v1/node.proto\x12\rbaepo.node.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1bbaepo/core/v1/machine.proto\"\xa4\x01\n" +
	"\aMachine\x12\x1d\n" +
	"\n" +
	"machine_id\x18\x01 \x01(\tR\tmachineId\x121\n" +
	"\x05state\x18\x02 \x01(\x0e2\x1b.baepo.core.v1.MachineStateR\x05state\x12G\n" +
	"\rdesired_state\x18\x03 \x01(\x0e2\".baepo.core.v1.MachineDesiredStateR\fdesiredState\"\x19\n" +
	"\x17NodeListMachinesRequest\"N\n" +
	"\x18NodeListMachinesResponse\x122\n" +
	"\bmachines\x18\x01 \x03(\v2\x16.baepo.node.v1.MachineR\bmachines\"6\n" +
	"\x15NodeGetMachineRequest\x12\x1d\n" +
	"\n" +
	"machine_id\x18\x01 \x01(\tR\tmachineId\"J\n" +
	"\x16NodeGetMachineResponse\x120\n" +
	"\amachine\x18\x01 \x01(\v2\x16.baepo.node.v1.MachineR\amachine\"R\n" +
	"\x19NodeGetMachineLogsRequest\x12\x1d\n" +
	"\n" +
	"machine_id\x18\x01 \x01(\tR\tmachineId\x12\x16\n" +
	"\x06follow\x18\x02 \x01(\bR\x06follow\"6\n" +
	"\x1aNodeGetMachineLogsResponse\x12\x18\n" +
	"\acontent\x18\x01 \x01(\fR\acontent\"\x85\x01\n" +
	"\x1bNodeGetContainerLogsRequest\x12\x1d\n" +
	"\n" +
	"machine_id\x18\x01 \x01(\tR\tmachineId\x12!\n" +
	"\tcontainer\x18\x02 \x01(\tH\x00R\tcontainer\x88\x01\x01\x12\x16\n" +
	"\x06follow\x18\x03 \x01(\bR\x06followB\f\n" +
	"\n" +
	"_container\"\xd2\x01\n" +
	"\x1cNodeGetContainerLogsResponse\x12!\n" +
	"\fcontainer_id\x18\x01 \x01(\tR\vcontainerId\x12%\n" +
	"\x0econtainer_name\x18\x02 \x01(\tR\rcontainerName\x12\x14\n" +
	"\x05error\x18\x03 \x01(\bR\x05error\x12\x18\n" +
	"\acontent\x18\x04 \x01(\fR\acontent\x128\n" +
	"\ttimestamp\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\ttimestamp2\xa1\x03\n" +
	"\vNodeService\x12_\n" +
	"\fListMachines\x12&.baepo.node.v1.NodeListMachinesRequest\x1a'.baepo.node.v1.NodeListMachinesResponse\x12Y\n" +
	"\n" +
	"GetMachine\x12$.baepo.node.v1.NodeGetMachineRequest\x1a%.baepo.node.v1.NodeGetMachineResponse\x12g\n" +
	"\x0eGetMachineLogs\x12(.baepo.node.v1.NodeGetMachineLogsRequest\x1a).baepo.node.v1.NodeGetMachineLogsResponse0\x01\x12m\n" +
	"\x10GetContainerLogs\x12*.baepo.node.v1.NodeGetContainerLogsRequest\x1a+.baepo.node.v1.NodeGetContainerLogsResponse0\x01B>Z<github.com/baepo-cloud/baepo-proto/go/baepo/node/v1;nodev1pbb\x06proto3"

var (
	file_baepo_node_v1_node_proto_rawDescOnce sync.Once
	file_baepo_node_v1_node_proto_rawDescData []byte
)

func file_baepo_node_v1_node_proto_rawDescGZIP() []byte {
	file_baepo_node_v1_node_proto_rawDescOnce.Do(func() {
		file_baepo_node_v1_node_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_baepo_node_v1_node_proto_rawDesc), len(file_baepo_node_v1_node_proto_rawDesc)))
	})
	return file_baepo_node_v1_node_proto_rawDescData
}

var file_baepo_node_v1_node_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_baepo_node_v1_node_proto_goTypes = []any{
	(*Machine)(nil),                      // 0: baepo.node.v1.Machine
	(*NodeListMachinesRequest)(nil),      // 1: baepo.node.v1.NodeListMachinesRequest
	(*NodeListMachinesResponse)(nil),     // 2: baepo.node.v1.NodeListMachinesResponse
	(*NodeGetMachineRequest)(nil),        // 3: baepo.node.v1.NodeGetMachineRequest
	(*NodeGetMachineResponse)(nil),       // 4: baepo.node.v1.NodeGetMachineResponse
	(*NodeGetMachineLogsRequest)(nil),    // 5: baepo.node.v1.NodeGetMachineLogsRequest
	(*NodeGetMachineLogsResponse)(nil),   // 6: baepo.node.v1.NodeGetMachineLogsResponse
	(*NodeGetContainerLogsRequest)(nil),  // 7: baepo.node.v1.NodeGetContainerLogsRequest
	(*NodeGetContainerLogsResponse)(nil), // 8: baepo.node.v1.NodeGetContainerLogsResponse
	(v1.MachineState)(0),                 // 9: baepo.core.v1.MachineState
	(v1.MachineDesiredState)(0),          // 10: baepo.core.v1.MachineDesiredState
	(*timestamppb.Timestamp)(nil),        // 11: google.protobuf.Timestamp
}
var file_baepo_node_v1_node_proto_depIdxs = []int32{
	9,  // 0: baepo.node.v1.Machine.state:type_name -> baepo.core.v1.MachineState
	10, // 1: baepo.node.v1.Machine.desired_state:type_name -> baepo.core.v1.MachineDesiredState
	0,  // 2: baepo.node.v1.NodeListMachinesResponse.machines:type_name -> baepo.node.v1.Machine
	0,  // 3: baepo.node.v1.NodeGetMachineResponse.machine:type_name -> baepo.node.v1.Machine
	11, // 4: baepo.node.v1.NodeGetContainerLogsResponse.timestamp:type_name -> google.protobuf.Timestamp
	1,  // 5: baepo.node.v1.NodeService.ListMachines:input_type -> baepo.node.v1.NodeListMachinesRequest
	3,  // 6: baepo.node.v1.NodeService.GetMachine:input_type -> baepo.node.v1.NodeGetMachineRequest
	5,  // 7: baepo.node.v1.NodeService.GetMachineLogs:input_type -> baepo.node.v1.NodeGetMachineLogsRequest
	7,  // 8: baepo.node.v1.NodeService.GetContainerLogs:input_type -> baepo.node.v1.NodeGetContainerLogsRequest
	2,  // 9: baepo.node.v1.NodeService.ListMachines:output_type -> baepo.node.v1.NodeListMachinesResponse
	4,  // 10: baepo.node.v1.NodeService.GetMachine:output_type -> baepo.node.v1.NodeGetMachineResponse
	6,  // 11: baepo.node.v1.NodeService.GetMachineLogs:output_type -> baepo.node.v1.NodeGetMachineLogsResponse
	8,  // 12: baepo.node.v1.NodeService.GetContainerLogs:output_type -> baepo.node.v1.NodeGetContainerLogsResponse
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_baepo_node_v1_node_proto_init() }
func file_baepo_node_v1_node_proto_init() {
	if File_baepo_node_v1_node_proto != nil {
		return
	}
	file_baepo_node_v1_node_proto_msgTypes[7].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_baepo_node_v1_node_proto_rawDesc), len(file_baepo_node_v1_node_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_baepo_node_v1_node_proto_goTypes,
		DependencyIndexes: file_baepo_node_v1_node_proto_depIdxs,
		MessageInfos:      file_baepo_node_v1_node_proto_msgTypes,
	}.Build()
	File_baepo_node_v1_node_proto = out.File
	file_baepo_node_v1_node_proto_goTypes = nil
	file_baepo_node_v1_node_proto_depIdxs = nil
}
