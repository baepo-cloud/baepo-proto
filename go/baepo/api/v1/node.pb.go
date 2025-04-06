// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: baepo/api/v1/node.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
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

type NodeConnectServerEvent struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Event:
	//
	//	*NodeConnectServerEvent_Register
	//	*NodeConnectServerEvent_Ping
	Event         isNodeConnectServerEvent_Event `protobuf_oneof:"event"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NodeConnectServerEvent) Reset() {
	*x = NodeConnectServerEvent{}
	mi := &file_baepo_api_v1_node_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NodeConnectServerEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeConnectServerEvent) ProtoMessage() {}

func (x *NodeConnectServerEvent) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_api_v1_node_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeConnectServerEvent.ProtoReflect.Descriptor instead.
func (*NodeConnectServerEvent) Descriptor() ([]byte, []int) {
	return file_baepo_api_v1_node_proto_rawDescGZIP(), []int{0}
}

func (x *NodeConnectServerEvent) GetEvent() isNodeConnectServerEvent_Event {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *NodeConnectServerEvent) GetRegister() *NodeConnectServerEvent_RegisterResponse {
	if x != nil {
		if x, ok := x.Event.(*NodeConnectServerEvent_Register); ok {
			return x.Register
		}
	}
	return nil
}

func (x *NodeConnectServerEvent) GetPing() *NodeConnectServerEvent_PingEvent {
	if x != nil {
		if x, ok := x.Event.(*NodeConnectServerEvent_Ping); ok {
			return x.Ping
		}
	}
	return nil
}

type isNodeConnectServerEvent_Event interface {
	isNodeConnectServerEvent_Event()
}

type NodeConnectServerEvent_Register struct {
	Register *NodeConnectServerEvent_RegisterResponse `protobuf:"bytes,1,opt,name=register,proto3,oneof"`
}

type NodeConnectServerEvent_Ping struct {
	Ping *NodeConnectServerEvent_PingEvent `protobuf:"bytes,2,opt,name=ping,proto3,oneof"`
}

func (*NodeConnectServerEvent_Register) isNodeConnectServerEvent_Event() {}

func (*NodeConnectServerEvent_Ping) isNodeConnectServerEvent_Event() {}

type NodeConnectClientEvent struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Event:
	//
	//	*NodeConnectClientEvent_Register
	//	*NodeConnectClientEvent_Stats
	Event         isNodeConnectClientEvent_Event `protobuf_oneof:"event"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NodeConnectClientEvent) Reset() {
	*x = NodeConnectClientEvent{}
	mi := &file_baepo_api_v1_node_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NodeConnectClientEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeConnectClientEvent) ProtoMessage() {}

func (x *NodeConnectClientEvent) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_api_v1_node_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeConnectClientEvent.ProtoReflect.Descriptor instead.
func (*NodeConnectClientEvent) Descriptor() ([]byte, []int) {
	return file_baepo_api_v1_node_proto_rawDescGZIP(), []int{1}
}

func (x *NodeConnectClientEvent) GetEvent() isNodeConnectClientEvent_Event {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *NodeConnectClientEvent) GetRegister() *NodeConnectClientEvent_RegisterRequest {
	if x != nil {
		if x, ok := x.Event.(*NodeConnectClientEvent_Register); ok {
			return x.Register
		}
	}
	return nil
}

func (x *NodeConnectClientEvent) GetStats() *NodeConnectClientEvent_StatsEvent {
	if x != nil {
		if x, ok := x.Event.(*NodeConnectClientEvent_Stats); ok {
			return x.Stats
		}
	}
	return nil
}

type isNodeConnectClientEvent_Event interface {
	isNodeConnectClientEvent_Event()
}

type NodeConnectClientEvent_Register struct {
	Register *NodeConnectClientEvent_RegisterRequest `protobuf:"bytes,1,opt,name=register,proto3,oneof"`
}

type NodeConnectClientEvent_Stats struct {
	Stats *NodeConnectClientEvent_StatsEvent `protobuf:"bytes,2,opt,name=stats,proto3,oneof"`
}

func (*NodeConnectClientEvent_Register) isNodeConnectClientEvent_Event() {}

func (*NodeConnectClientEvent_Stats) isNodeConnectClientEvent_Event() {}

type NodeConnectServerEvent_RegisterResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	NodeId        string                 `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	NodeToken     string                 `protobuf:"bytes,2,opt,name=node_token,json=nodeToken,proto3" json:"node_token,omitempty"`
	AuthorityCert []byte                 `protobuf:"bytes,3,opt,name=authority_cert,json=authorityCert,proto3" json:"authority_cert,omitempty"`
	ServerCert    []byte                 `protobuf:"bytes,4,opt,name=server_cert,json=serverCert,proto3" json:"server_cert,omitempty"`
	ServerKey     []byte                 `protobuf:"bytes,5,opt,name=server_key,json=serverKey,proto3" json:"server_key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NodeConnectServerEvent_RegisterResponse) Reset() {
	*x = NodeConnectServerEvent_RegisterResponse{}
	mi := &file_baepo_api_v1_node_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NodeConnectServerEvent_RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeConnectServerEvent_RegisterResponse) ProtoMessage() {}

func (x *NodeConnectServerEvent_RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_api_v1_node_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeConnectServerEvent_RegisterResponse.ProtoReflect.Descriptor instead.
func (*NodeConnectServerEvent_RegisterResponse) Descriptor() ([]byte, []int) {
	return file_baepo_api_v1_node_proto_rawDescGZIP(), []int{0, 0}
}

func (x *NodeConnectServerEvent_RegisterResponse) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *NodeConnectServerEvent_RegisterResponse) GetNodeToken() string {
	if x != nil {
		return x.NodeToken
	}
	return ""
}

func (x *NodeConnectServerEvent_RegisterResponse) GetAuthorityCert() []byte {
	if x != nil {
		return x.AuthorityCert
	}
	return nil
}

func (x *NodeConnectServerEvent_RegisterResponse) GetServerCert() []byte {
	if x != nil {
		return x.ServerCert
	}
	return nil
}

func (x *NodeConnectServerEvent_RegisterResponse) GetServerKey() []byte {
	if x != nil {
		return x.ServerKey
	}
	return nil
}

type NodeConnectServerEvent_PingEvent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NodeConnectServerEvent_PingEvent) Reset() {
	*x = NodeConnectServerEvent_PingEvent{}
	mi := &file_baepo_api_v1_node_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NodeConnectServerEvent_PingEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeConnectServerEvent_PingEvent) ProtoMessage() {}

func (x *NodeConnectServerEvent_PingEvent) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_api_v1_node_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeConnectServerEvent_PingEvent.ProtoReflect.Descriptor instead.
func (*NodeConnectServerEvent_PingEvent) Descriptor() ([]byte, []int) {
	return file_baepo_api_v1_node_proto_rawDescGZIP(), []int{0, 1}
}

type NodeConnectClientEvent_RegisterRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	ClusterId       string                 `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	BootstrapToken  string                 `protobuf:"bytes,2,opt,name=bootstrap_token,json=bootstrapToken,proto3" json:"bootstrap_token,omitempty"`
	NodeToken       *string                `protobuf:"bytes,3,opt,name=node_token,json=nodeToken,proto3,oneof" json:"node_token,omitempty"`
	IpAddress       string                 `protobuf:"bytes,4,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	ApiEndpoint     string                 `protobuf:"bytes,5,opt,name=api_endpoint,json=apiEndpoint,proto3" json:"api_endpoint,omitempty"`
	GatewayEndpoint string                 `protobuf:"bytes,6,opt,name=gateway_endpoint,json=gatewayEndpoint,proto3" json:"gateway_endpoint,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *NodeConnectClientEvent_RegisterRequest) Reset() {
	*x = NodeConnectClientEvent_RegisterRequest{}
	mi := &file_baepo_api_v1_node_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NodeConnectClientEvent_RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeConnectClientEvent_RegisterRequest) ProtoMessage() {}

func (x *NodeConnectClientEvent_RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_api_v1_node_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeConnectClientEvent_RegisterRequest.ProtoReflect.Descriptor instead.
func (*NodeConnectClientEvent_RegisterRequest) Descriptor() ([]byte, []int) {
	return file_baepo_api_v1_node_proto_rawDescGZIP(), []int{1, 0}
}

func (x *NodeConnectClientEvent_RegisterRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *NodeConnectClientEvent_RegisterRequest) GetBootstrapToken() string {
	if x != nil {
		return x.BootstrapToken
	}
	return ""
}

func (x *NodeConnectClientEvent_RegisterRequest) GetNodeToken() string {
	if x != nil && x.NodeToken != nil {
		return *x.NodeToken
	}
	return ""
}

func (x *NodeConnectClientEvent_RegisterRequest) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

func (x *NodeConnectClientEvent_RegisterRequest) GetApiEndpoint() string {
	if x != nil {
		return x.ApiEndpoint
	}
	return ""
}

func (x *NodeConnectClientEvent_RegisterRequest) GetGatewayEndpoint() string {
	if x != nil {
		return x.GatewayEndpoint
	}
	return ""
}

type NodeConnectClientEvent_StatsEvent struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	TotalMemory       uint64                 `protobuf:"varint,1,opt,name=total_memory,json=totalMemory,proto3" json:"total_memory,omitempty"`
	UsedMemory        uint64                 `protobuf:"varint,2,opt,name=used_memory,json=usedMemory,proto3" json:"used_memory,omitempty"`
	CpuCount          uint32                 `protobuf:"varint,3,opt,name=cpu_count,json=cpuCount,proto3" json:"cpu_count,omitempty"`
	RunningMachineIds []string               `protobuf:"bytes,4,rep,name=running_machine_ids,json=runningMachineIds,proto3" json:"running_machine_ids,omitempty"`
	ReservedMemory    uint64                 `protobuf:"varint,5,opt,name=reserved_memory,json=reservedMemory,proto3" json:"reserved_memory,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *NodeConnectClientEvent_StatsEvent) Reset() {
	*x = NodeConnectClientEvent_StatsEvent{}
	mi := &file_baepo_api_v1_node_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NodeConnectClientEvent_StatsEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeConnectClientEvent_StatsEvent) ProtoMessage() {}

func (x *NodeConnectClientEvent_StatsEvent) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_api_v1_node_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeConnectClientEvent_StatsEvent.ProtoReflect.Descriptor instead.
func (*NodeConnectClientEvent_StatsEvent) Descriptor() ([]byte, []int) {
	return file_baepo_api_v1_node_proto_rawDescGZIP(), []int{1, 1}
}

func (x *NodeConnectClientEvent_StatsEvent) GetTotalMemory() uint64 {
	if x != nil {
		return x.TotalMemory
	}
	return 0
}

func (x *NodeConnectClientEvent_StatsEvent) GetUsedMemory() uint64 {
	if x != nil {
		return x.UsedMemory
	}
	return 0
}

func (x *NodeConnectClientEvent_StatsEvent) GetCpuCount() uint32 {
	if x != nil {
		return x.CpuCount
	}
	return 0
}

func (x *NodeConnectClientEvent_StatsEvent) GetRunningMachineIds() []string {
	if x != nil {
		return x.RunningMachineIds
	}
	return nil
}

func (x *NodeConnectClientEvent_StatsEvent) GetReservedMemory() uint64 {
	if x != nil {
		return x.ReservedMemory
	}
	return 0
}

var File_baepo_api_v1_node_proto protoreflect.FileDescriptor

const file_baepo_api_v1_node_proto_rawDesc = "" +
	"\n" +
	"\x17baepo/api/v1/node.proto\x12\fbaepo.api.v1\x1a\x1bgoogle/protobuf/empty.proto\"\xfd\x02\n" +
	"\x16NodeConnectServerEvent\x12S\n" +
	"\bregister\x18\x01 \x01(\v25.baepo.api.v1.NodeConnectServerEvent.RegisterResponseH\x00R\bregister\x12D\n" +
	"\x04ping\x18\x02 \x01(\v2..baepo.api.v1.NodeConnectServerEvent.PingEventH\x00R\x04ping\x1a\xb1\x01\n" +
	"\x10RegisterResponse\x12\x17\n" +
	"\anode_id\x18\x01 \x01(\tR\x06nodeId\x12\x1d\n" +
	"\n" +
	"node_token\x18\x02 \x01(\tR\tnodeToken\x12%\n" +
	"\x0eauthority_cert\x18\x03 \x01(\fR\rauthorityCert\x12\x1f\n" +
	"\vserver_cert\x18\x04 \x01(\fR\n" +
	"serverCert\x12\x1d\n" +
	"\n" +
	"server_key\x18\x05 \x01(\fR\tserverKey\x1a\v\n" +
	"\tPingEventB\a\n" +
	"\x05event\"\x83\x05\n" +
	"\x16NodeConnectClientEvent\x12R\n" +
	"\bregister\x18\x01 \x01(\v24.baepo.api.v1.NodeConnectClientEvent.RegisterRequestH\x00R\bregister\x12G\n" +
	"\x05stats\x18\x02 \x01(\v2/.baepo.api.v1.NodeConnectClientEvent.StatsEventH\x00R\x05stats\x1a\xf9\x01\n" +
	"\x0fRegisterRequest\x12\x1d\n" +
	"\n" +
	"cluster_id\x18\x01 \x01(\tR\tclusterId\x12'\n" +
	"\x0fbootstrap_token\x18\x02 \x01(\tR\x0ebootstrapToken\x12\"\n" +
	"\n" +
	"node_token\x18\x03 \x01(\tH\x00R\tnodeToken\x88\x01\x01\x12\x1d\n" +
	"\n" +
	"ip_address\x18\x04 \x01(\tR\tipAddress\x12!\n" +
	"\fapi_endpoint\x18\x05 \x01(\tR\vapiEndpoint\x12)\n" +
	"\x10gateway_endpoint\x18\x06 \x01(\tR\x0fgatewayEndpointB\r\n" +
	"\v_node_token\x1a\xc6\x01\n" +
	"\n" +
	"StatsEvent\x12!\n" +
	"\ftotal_memory\x18\x01 \x01(\x04R\vtotalMemory\x12\x1f\n" +
	"\vused_memory\x18\x02 \x01(\x04R\n" +
	"usedMemory\x12\x1b\n" +
	"\tcpu_count\x18\x03 \x01(\rR\bcpuCount\x12.\n" +
	"\x13running_machine_ids\x18\x04 \x03(\tR\x11runningMachineIds\x12'\n" +
	"\x0freserved_memory\x18\x05 \x01(\x04R\x0ereservedMemoryB\a\n" +
	"\x05event2h\n" +
	"\vNodeService\x12Y\n" +
	"\aConnect\x12$.baepo.api.v1.NodeConnectClientEvent\x1a$.baepo.api.v1.NodeConnectServerEvent(\x010\x01B4Z2github.com/baepo-cloud/baepo-proto/go/baepo/api/v1b\x06proto3"

var (
	file_baepo_api_v1_node_proto_rawDescOnce sync.Once
	file_baepo_api_v1_node_proto_rawDescData []byte
)

func file_baepo_api_v1_node_proto_rawDescGZIP() []byte {
	file_baepo_api_v1_node_proto_rawDescOnce.Do(func() {
		file_baepo_api_v1_node_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_baepo_api_v1_node_proto_rawDesc), len(file_baepo_api_v1_node_proto_rawDesc)))
	})
	return file_baepo_api_v1_node_proto_rawDescData
}

var file_baepo_api_v1_node_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_baepo_api_v1_node_proto_goTypes = []any{
	(*NodeConnectServerEvent)(nil),                  // 0: baepo.api.v1.NodeConnectServerEvent
	(*NodeConnectClientEvent)(nil),                  // 1: baepo.api.v1.NodeConnectClientEvent
	(*NodeConnectServerEvent_RegisterResponse)(nil), // 2: baepo.api.v1.NodeConnectServerEvent.RegisterResponse
	(*NodeConnectServerEvent_PingEvent)(nil),        // 3: baepo.api.v1.NodeConnectServerEvent.PingEvent
	(*NodeConnectClientEvent_RegisterRequest)(nil),  // 4: baepo.api.v1.NodeConnectClientEvent.RegisterRequest
	(*NodeConnectClientEvent_StatsEvent)(nil),       // 5: baepo.api.v1.NodeConnectClientEvent.StatsEvent
}
var file_baepo_api_v1_node_proto_depIdxs = []int32{
	2, // 0: baepo.api.v1.NodeConnectServerEvent.register:type_name -> baepo.api.v1.NodeConnectServerEvent.RegisterResponse
	3, // 1: baepo.api.v1.NodeConnectServerEvent.ping:type_name -> baepo.api.v1.NodeConnectServerEvent.PingEvent
	4, // 2: baepo.api.v1.NodeConnectClientEvent.register:type_name -> baepo.api.v1.NodeConnectClientEvent.RegisterRequest
	5, // 3: baepo.api.v1.NodeConnectClientEvent.stats:type_name -> baepo.api.v1.NodeConnectClientEvent.StatsEvent
	1, // 4: baepo.api.v1.NodeService.Connect:input_type -> baepo.api.v1.NodeConnectClientEvent
	0, // 5: baepo.api.v1.NodeService.Connect:output_type -> baepo.api.v1.NodeConnectServerEvent
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_baepo_api_v1_node_proto_init() }
func file_baepo_api_v1_node_proto_init() {
	if File_baepo_api_v1_node_proto != nil {
		return
	}
	file_baepo_api_v1_node_proto_msgTypes[0].OneofWrappers = []any{
		(*NodeConnectServerEvent_Register)(nil),
		(*NodeConnectServerEvent_Ping)(nil),
	}
	file_baepo_api_v1_node_proto_msgTypes[1].OneofWrappers = []any{
		(*NodeConnectClientEvent_Register)(nil),
		(*NodeConnectClientEvent_Stats)(nil),
	}
	file_baepo_api_v1_node_proto_msgTypes[4].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_baepo_api_v1_node_proto_rawDesc), len(file_baepo_api_v1_node_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_baepo_api_v1_node_proto_goTypes,
		DependencyIndexes: file_baepo_api_v1_node_proto_depIdxs,
		MessageInfos:      file_baepo_api_v1_node_proto_msgTypes,
	}.Build()
	File_baepo_api_v1_node_proto = out.File
	file_baepo_api_v1_node_proto_goTypes = nil
	file_baepo_api_v1_node_proto_depIdxs = nil
}
