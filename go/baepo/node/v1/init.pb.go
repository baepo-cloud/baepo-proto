// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: baepo/node/v1/init.proto

package nodev1pb

import (
	v1 "github.com/baepo-cloud/baepo-proto/go/baepo/core/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type InitGetLogsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Container     *string                `protobuf:"bytes,1,opt,name=container,proto3,oneof" json:"container,omitempty"`
	Follow        bool                   `protobuf:"varint,2,opt,name=follow,proto3" json:"follow,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InitGetLogsRequest) Reset() {
	*x = InitGetLogsRequest{}
	mi := &file_baepo_node_v1_init_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InitGetLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitGetLogsRequest) ProtoMessage() {}

func (x *InitGetLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_node_v1_init_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitGetLogsRequest.ProtoReflect.Descriptor instead.
func (*InitGetLogsRequest) Descriptor() ([]byte, []int) {
	return file_baepo_node_v1_init_proto_rawDescGZIP(), []int{0}
}

func (x *InitGetLogsRequest) GetContainer() string {
	if x != nil && x.Container != nil {
		return *x.Container
	}
	return ""
}

func (x *InitGetLogsRequest) GetFollow() bool {
	if x != nil {
		return x.Follow
	}
	return false
}

type InitGetLogsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ContainerId   string                 `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	ContainerName string                 `protobuf:"bytes,2,opt,name=container_name,json=containerName,proto3" json:"container_name,omitempty"`
	Error         bool                   `protobuf:"varint,3,opt,name=error,proto3" json:"error,omitempty"`
	Content       []byte                 `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Timestamp     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InitGetLogsResponse) Reset() {
	*x = InitGetLogsResponse{}
	mi := &file_baepo_node_v1_init_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InitGetLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitGetLogsResponse) ProtoMessage() {}

func (x *InitGetLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_node_v1_init_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitGetLogsResponse.ProtoReflect.Descriptor instead.
func (*InitGetLogsResponse) Descriptor() ([]byte, []int) {
	return file_baepo_node_v1_init_proto_rawDescGZIP(), []int{1}
}

func (x *InitGetLogsResponse) GetContainerId() string {
	if x != nil {
		return x.ContainerId
	}
	return ""
}

func (x *InitGetLogsResponse) GetContainerName() string {
	if x != nil {
		return x.ContainerName
	}
	return ""
}

func (x *InitGetLogsResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *InitGetLogsResponse) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *InitGetLogsResponse) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type InitEventsResponse struct {
	state     protoimpl.MessageState `protogen:"open.v1"`
	EventId   string                 `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Types that are valid to be assigned to Event:
	//
	//	*InitEventsResponse_ContainerStateChanged
	//	*InitEventsResponse_Ping
	Event         isInitEventsResponse_Event `protobuf_oneof:"event"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InitEventsResponse) Reset() {
	*x = InitEventsResponse{}
	mi := &file_baepo_node_v1_init_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InitEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitEventsResponse) ProtoMessage() {}

func (x *InitEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_node_v1_init_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitEventsResponse.ProtoReflect.Descriptor instead.
func (*InitEventsResponse) Descriptor() ([]byte, []int) {
	return file_baepo_node_v1_init_proto_rawDescGZIP(), []int{2}
}

func (x *InitEventsResponse) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *InitEventsResponse) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *InitEventsResponse) GetEvent() isInitEventsResponse_Event {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *InitEventsResponse) GetContainerStateChanged() *InitEventsResponse_ContainerStateChangedEvent {
	if x != nil {
		if x, ok := x.Event.(*InitEventsResponse_ContainerStateChanged); ok {
			return x.ContainerStateChanged
		}
	}
	return nil
}

func (x *InitEventsResponse) GetPing() *InitEventsResponse_PingEvent {
	if x != nil {
		if x, ok := x.Event.(*InitEventsResponse_Ping); ok {
			return x.Ping
		}
	}
	return nil
}

type isInitEventsResponse_Event interface {
	isInitEventsResponse_Event()
}

type InitEventsResponse_ContainerStateChanged struct {
	ContainerStateChanged *InitEventsResponse_ContainerStateChangedEvent `protobuf:"bytes,3,opt,name=container_state_changed,json=containerStateChanged,proto3,oneof"`
}

type InitEventsResponse_Ping struct {
	Ping *InitEventsResponse_PingEvent `protobuf:"bytes,4,opt,name=ping,proto3,oneof"`
}

func (*InitEventsResponse_ContainerStateChanged) isInitEventsResponse_Event() {}

func (*InitEventsResponse_Ping) isInitEventsResponse_Event() {}

type InitEventsResponse_ContainerStateChangedEvent struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	ContainerId      string                 `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	State            v1.ContainerState      `protobuf:"varint,2,opt,name=state,proto3,enum=baepo.core.v1.ContainerState" json:"state,omitempty"`
	StartedAt        *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=started_at,json=startedAt,proto3,oneof" json:"started_at,omitempty"`
	ExitedAt         *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=exited_at,json=exitedAt,proto3,oneof" json:"exited_at,omitempty"`
	ExitCode         *int32                 `protobuf:"varint,5,opt,name=exit_code,json=exitCode,proto3,oneof" json:"exit_code,omitempty"`
	ExitError        *string                `protobuf:"bytes,6,opt,name=exit_error,json=exitError,proto3,oneof" json:"exit_error,omitempty"`
	Healthy          bool                   `protobuf:"varint,7,opt,name=healthy,proto3" json:"healthy,omitempty"`
	HealthcheckError *string                `protobuf:"bytes,8,opt,name=healthcheck_error,json=healthcheckError,proto3,oneof" json:"healthcheck_error,omitempty"`
	RestartCount     int32                  `protobuf:"varint,9,opt,name=restart_count,json=restartCount,proto3" json:"restart_count,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *InitEventsResponse_ContainerStateChangedEvent) Reset() {
	*x = InitEventsResponse_ContainerStateChangedEvent{}
	mi := &file_baepo_node_v1_init_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InitEventsResponse_ContainerStateChangedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitEventsResponse_ContainerStateChangedEvent) ProtoMessage() {}

func (x *InitEventsResponse_ContainerStateChangedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_node_v1_init_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitEventsResponse_ContainerStateChangedEvent.ProtoReflect.Descriptor instead.
func (*InitEventsResponse_ContainerStateChangedEvent) Descriptor() ([]byte, []int) {
	return file_baepo_node_v1_init_proto_rawDescGZIP(), []int{2, 0}
}

func (x *InitEventsResponse_ContainerStateChangedEvent) GetContainerId() string {
	if x != nil {
		return x.ContainerId
	}
	return ""
}

func (x *InitEventsResponse_ContainerStateChangedEvent) GetState() v1.ContainerState {
	if x != nil {
		return x.State
	}
	return v1.ContainerState(0)
}

func (x *InitEventsResponse_ContainerStateChangedEvent) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *InitEventsResponse_ContainerStateChangedEvent) GetExitedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExitedAt
	}
	return nil
}

func (x *InitEventsResponse_ContainerStateChangedEvent) GetExitCode() int32 {
	if x != nil && x.ExitCode != nil {
		return *x.ExitCode
	}
	return 0
}

func (x *InitEventsResponse_ContainerStateChangedEvent) GetExitError() string {
	if x != nil && x.ExitError != nil {
		return *x.ExitError
	}
	return ""
}

func (x *InitEventsResponse_ContainerStateChangedEvent) GetHealthy() bool {
	if x != nil {
		return x.Healthy
	}
	return false
}

func (x *InitEventsResponse_ContainerStateChangedEvent) GetHealthcheckError() string {
	if x != nil && x.HealthcheckError != nil {
		return *x.HealthcheckError
	}
	return ""
}

func (x *InitEventsResponse_ContainerStateChangedEvent) GetRestartCount() int32 {
	if x != nil {
		return x.RestartCount
	}
	return 0
}

type InitEventsResponse_PingEvent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InitEventsResponse_PingEvent) Reset() {
	*x = InitEventsResponse_PingEvent{}
	mi := &file_baepo_node_v1_init_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InitEventsResponse_PingEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitEventsResponse_PingEvent) ProtoMessage() {}

func (x *InitEventsResponse_PingEvent) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_node_v1_init_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitEventsResponse_PingEvent.ProtoReflect.Descriptor instead.
func (*InitEventsResponse_PingEvent) Descriptor() ([]byte, []int) {
	return file_baepo_node_v1_init_proto_rawDescGZIP(), []int{2, 1}
}

var File_baepo_node_v1_init_proto protoreflect.FileDescriptor

const file_baepo_node_v1_init_proto_rawDesc = "" +
	"\n" +
	"\x18baepo/node/v1/init.proto\x12\rbaepo.node.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1dbaepo/core/v1/container.proto\"]\n" +
	"\x12InitGetLogsRequest\x12!\n" +
	"\tcontainer\x18\x01 \x01(\tH\x00R\tcontainer\x88\x01\x01\x12\x16\n" +
	"\x06follow\x18\x02 \x01(\bR\x06followB\f\n" +
	"\n" +
	"_container\"\xc9\x01\n" +
	"\x13InitGetLogsResponse\x12!\n" +
	"\fcontainer_id\x18\x01 \x01(\tR\vcontainerId\x12%\n" +
	"\x0econtainer_name\x18\x02 \x01(\tR\rcontainerName\x12\x14\n" +
	"\x05error\x18\x03 \x01(\bR\x05error\x12\x18\n" +
	"\acontent\x18\x04 \x01(\fR\acontent\x128\n" +
	"\ttimestamp\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\ttimestamp\"\xb6\x06\n" +
	"\x12InitEventsResponse\x12\x19\n" +
	"\bevent_id\x18\x01 \x01(\tR\aeventId\x128\n" +
	"\ttimestamp\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\ttimestamp\x12v\n" +
	"\x17container_state_changed\x18\x03 \x01(\v2<.baepo.node.v1.InitEventsResponse.ContainerStateChangedEventH\x00R\x15containerStateChanged\x12A\n" +
	"\x04ping\x18\x04 \x01(\v2+.baepo.node.v1.InitEventsResponse.PingEventH\x00R\x04ping\x1a\xf9\x03\n" +
	"\x1aContainerStateChangedEvent\x12!\n" +
	"\fcontainer_id\x18\x01 \x01(\tR\vcontainerId\x123\n" +
	"\x05state\x18\x02 \x01(\x0e2\x1d.baepo.core.v1.ContainerStateR\x05state\x12>\n" +
	"\n" +
	"started_at\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampH\x00R\tstartedAt\x88\x01\x01\x12<\n" +
	"\texited_at\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampH\x01R\bexitedAt\x88\x01\x01\x12 \n" +
	"\texit_code\x18\x05 \x01(\x05H\x02R\bexitCode\x88\x01\x01\x12\"\n" +
	"\n" +
	"exit_error\x18\x06 \x01(\tH\x03R\texitError\x88\x01\x01\x12\x18\n" +
	"\ahealthy\x18\a \x01(\bR\ahealthy\x120\n" +
	"\x11healthcheck_error\x18\b \x01(\tH\x04R\x10healthcheckError\x88\x01\x01\x12#\n" +
	"\rrestart_count\x18\t \x01(\x05R\frestartCountB\r\n" +
	"\v_started_atB\f\n" +
	"\n" +
	"_exited_atB\f\n" +
	"\n" +
	"_exit_codeB\r\n" +
	"\v_exit_errorB\x14\n" +
	"\x12_healthcheck_error\x1a\v\n" +
	"\tPingEventB\a\n" +
	"\x05event2\xa1\x01\n" +
	"\x04Init\x12R\n" +
	"\aGetLogs\x12!.baepo.node.v1.InitGetLogsRequest\x1a\".baepo.node.v1.InitGetLogsResponse0\x01\x12E\n" +
	"\x06Events\x12\x16.google.protobuf.Empty\x1a!.baepo.node.v1.InitEventsResponse0\x01B>Z<github.com/baepo-cloud/baepo-proto/go/baepo/node/v1;nodev1pbb\x06proto3"

var (
	file_baepo_node_v1_init_proto_rawDescOnce sync.Once
	file_baepo_node_v1_init_proto_rawDescData []byte
)

func file_baepo_node_v1_init_proto_rawDescGZIP() []byte {
	file_baepo_node_v1_init_proto_rawDescOnce.Do(func() {
		file_baepo_node_v1_init_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_baepo_node_v1_init_proto_rawDesc), len(file_baepo_node_v1_init_proto_rawDesc)))
	})
	return file_baepo_node_v1_init_proto_rawDescData
}

var file_baepo_node_v1_init_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_baepo_node_v1_init_proto_goTypes = []any{
	(*InitGetLogsRequest)(nil),                            // 0: baepo.node.v1.InitGetLogsRequest
	(*InitGetLogsResponse)(nil),                           // 1: baepo.node.v1.InitGetLogsResponse
	(*InitEventsResponse)(nil),                            // 2: baepo.node.v1.InitEventsResponse
	(*InitEventsResponse_ContainerStateChangedEvent)(nil), // 3: baepo.node.v1.InitEventsResponse.ContainerStateChangedEvent
	(*InitEventsResponse_PingEvent)(nil),                  // 4: baepo.node.v1.InitEventsResponse.PingEvent
	(*timestamppb.Timestamp)(nil),                         // 5: google.protobuf.Timestamp
	(v1.ContainerState)(0),                                // 6: baepo.core.v1.ContainerState
	(*emptypb.Empty)(nil),                                 // 7: google.protobuf.Empty
}
var file_baepo_node_v1_init_proto_depIdxs = []int32{
	5, // 0: baepo.node.v1.InitGetLogsResponse.timestamp:type_name -> google.protobuf.Timestamp
	5, // 1: baepo.node.v1.InitEventsResponse.timestamp:type_name -> google.protobuf.Timestamp
	3, // 2: baepo.node.v1.InitEventsResponse.container_state_changed:type_name -> baepo.node.v1.InitEventsResponse.ContainerStateChangedEvent
	4, // 3: baepo.node.v1.InitEventsResponse.ping:type_name -> baepo.node.v1.InitEventsResponse.PingEvent
	6, // 4: baepo.node.v1.InitEventsResponse.ContainerStateChangedEvent.state:type_name -> baepo.core.v1.ContainerState
	5, // 5: baepo.node.v1.InitEventsResponse.ContainerStateChangedEvent.started_at:type_name -> google.protobuf.Timestamp
	5, // 6: baepo.node.v1.InitEventsResponse.ContainerStateChangedEvent.exited_at:type_name -> google.protobuf.Timestamp
	0, // 7: baepo.node.v1.Init.GetLogs:input_type -> baepo.node.v1.InitGetLogsRequest
	7, // 8: baepo.node.v1.Init.Events:input_type -> google.protobuf.Empty
	1, // 9: baepo.node.v1.Init.GetLogs:output_type -> baepo.node.v1.InitGetLogsResponse
	2, // 10: baepo.node.v1.Init.Events:output_type -> baepo.node.v1.InitEventsResponse
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_baepo_node_v1_init_proto_init() }
func file_baepo_node_v1_init_proto_init() {
	if File_baepo_node_v1_init_proto != nil {
		return
	}
	file_baepo_node_v1_init_proto_msgTypes[0].OneofWrappers = []any{}
	file_baepo_node_v1_init_proto_msgTypes[2].OneofWrappers = []any{
		(*InitEventsResponse_ContainerStateChanged)(nil),
		(*InitEventsResponse_Ping)(nil),
	}
	file_baepo_node_v1_init_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_baepo_node_v1_init_proto_rawDesc), len(file_baepo_node_v1_init_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_baepo_node_v1_init_proto_goTypes,
		DependencyIndexes: file_baepo_node_v1_init_proto_depIdxs,
		MessageInfos:      file_baepo_node_v1_init_proto_msgTypes,
	}.Build()
	File_baepo_node_v1_init_proto = out.File
	file_baepo_node_v1_init_proto_goTypes = nil
	file_baepo_node_v1_init_proto_depIdxs = nil
}
