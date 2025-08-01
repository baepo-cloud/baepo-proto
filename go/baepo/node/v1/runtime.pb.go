// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: baepo/node/v1/runtime.proto

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

type RuntimeGetStateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Pid           int64                  `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	Running       bool                   `protobuf:"varint,2,opt,name=running,proto3" json:"running,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RuntimeGetStateResponse) Reset() {
	*x = RuntimeGetStateResponse{}
	mi := &file_baepo_node_v1_runtime_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RuntimeGetStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuntimeGetStateResponse) ProtoMessage() {}

func (x *RuntimeGetStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_node_v1_runtime_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuntimeGetStateResponse.ProtoReflect.Descriptor instead.
func (*RuntimeGetStateResponse) Descriptor() ([]byte, []int) {
	return file_baepo_node_v1_runtime_proto_rawDescGZIP(), []int{0}
}

func (x *RuntimeGetStateResponse) GetPid() int64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *RuntimeGetStateResponse) GetRunning() bool {
	if x != nil {
		return x.Running
	}
	return false
}

type RuntimeGetLogsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Follow        bool                   `protobuf:"varint,1,opt,name=follow,proto3" json:"follow,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RuntimeGetLogsRequest) Reset() {
	*x = RuntimeGetLogsRequest{}
	mi := &file_baepo_node_v1_runtime_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RuntimeGetLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuntimeGetLogsRequest) ProtoMessage() {}

func (x *RuntimeGetLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_node_v1_runtime_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuntimeGetLogsRequest.ProtoReflect.Descriptor instead.
func (*RuntimeGetLogsRequest) Descriptor() ([]byte, []int) {
	return file_baepo_node_v1_runtime_proto_rawDescGZIP(), []int{1}
}

func (x *RuntimeGetLogsRequest) GetFollow() bool {
	if x != nil {
		return x.Follow
	}
	return false
}

type RuntimeGetLogsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Content       []byte                 `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Timestamp     *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RuntimeGetLogsResponse) Reset() {
	*x = RuntimeGetLogsResponse{}
	mi := &file_baepo_node_v1_runtime_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RuntimeGetLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuntimeGetLogsResponse) ProtoMessage() {}

func (x *RuntimeGetLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_node_v1_runtime_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuntimeGetLogsResponse.ProtoReflect.Descriptor instead.
func (*RuntimeGetLogsResponse) Descriptor() ([]byte, []int) {
	return file_baepo_node_v1_runtime_proto_rawDescGZIP(), []int{2}
}

func (x *RuntimeGetLogsResponse) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *RuntimeGetLogsResponse) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type RuntimeGetContainerLogsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Container     string                 `protobuf:"bytes,1,opt,name=container,proto3" json:"container,omitempty"`
	Follow        bool                   `protobuf:"varint,2,opt,name=follow,proto3" json:"follow,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RuntimeGetContainerLogsRequest) Reset() {
	*x = RuntimeGetContainerLogsRequest{}
	mi := &file_baepo_node_v1_runtime_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RuntimeGetContainerLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuntimeGetContainerLogsRequest) ProtoMessage() {}

func (x *RuntimeGetContainerLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_node_v1_runtime_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuntimeGetContainerLogsRequest.ProtoReflect.Descriptor instead.
func (*RuntimeGetContainerLogsRequest) Descriptor() ([]byte, []int) {
	return file_baepo_node_v1_runtime_proto_rawDescGZIP(), []int{3}
}

func (x *RuntimeGetContainerLogsRequest) GetContainer() string {
	if x != nil {
		return x.Container
	}
	return ""
}

func (x *RuntimeGetContainerLogsRequest) GetFollow() bool {
	if x != nil {
		return x.Follow
	}
	return false
}

type RuntimeGetContainerLogsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ContainerId   string                 `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	Error         bool                   `protobuf:"varint,2,opt,name=error,proto3" json:"error,omitempty"`
	Content       []byte                 `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Timestamp     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RuntimeGetContainerLogsResponse) Reset() {
	*x = RuntimeGetContainerLogsResponse{}
	mi := &file_baepo_node_v1_runtime_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RuntimeGetContainerLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuntimeGetContainerLogsResponse) ProtoMessage() {}

func (x *RuntimeGetContainerLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_node_v1_runtime_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuntimeGetContainerLogsResponse.ProtoReflect.Descriptor instead.
func (*RuntimeGetContainerLogsResponse) Descriptor() ([]byte, []int) {
	return file_baepo_node_v1_runtime_proto_rawDescGZIP(), []int{4}
}

func (x *RuntimeGetContainerLogsResponse) GetContainerId() string {
	if x != nil {
		return x.ContainerId
	}
	return ""
}

func (x *RuntimeGetContainerLogsResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *RuntimeGetContainerLogsResponse) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *RuntimeGetContainerLogsResponse) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type RuntimeEventsResponse struct {
	state     protoimpl.MessageState `protogen:"open.v1"`
	EventId   string                 `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Types that are valid to be assigned to Event:
	//
	//	*RuntimeEventsResponse_ContainerStateChanged
	//	*RuntimeEventsResponse_Ping
	Event         isRuntimeEventsResponse_Event `protobuf_oneof:"event"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RuntimeEventsResponse) Reset() {
	*x = RuntimeEventsResponse{}
	mi := &file_baepo_node_v1_runtime_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RuntimeEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuntimeEventsResponse) ProtoMessage() {}

func (x *RuntimeEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_node_v1_runtime_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuntimeEventsResponse.ProtoReflect.Descriptor instead.
func (*RuntimeEventsResponse) Descriptor() ([]byte, []int) {
	return file_baepo_node_v1_runtime_proto_rawDescGZIP(), []int{5}
}

func (x *RuntimeEventsResponse) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *RuntimeEventsResponse) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *RuntimeEventsResponse) GetEvent() isRuntimeEventsResponse_Event {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *RuntimeEventsResponse) GetContainerStateChanged() *RuntimeEventsResponse_ContainerStateChangedEvent {
	if x != nil {
		if x, ok := x.Event.(*RuntimeEventsResponse_ContainerStateChanged); ok {
			return x.ContainerStateChanged
		}
	}
	return nil
}

func (x *RuntimeEventsResponse) GetPing() *RuntimeEventsResponse_PingEvent {
	if x != nil {
		if x, ok := x.Event.(*RuntimeEventsResponse_Ping); ok {
			return x.Ping
		}
	}
	return nil
}

type isRuntimeEventsResponse_Event interface {
	isRuntimeEventsResponse_Event()
}

type RuntimeEventsResponse_ContainerStateChanged struct {
	ContainerStateChanged *RuntimeEventsResponse_ContainerStateChangedEvent `protobuf:"bytes,3,opt,name=container_state_changed,json=containerStateChanged,proto3,oneof"`
}

type RuntimeEventsResponse_Ping struct {
	Ping *RuntimeEventsResponse_PingEvent `protobuf:"bytes,4,opt,name=ping,proto3,oneof"`
}

func (*RuntimeEventsResponse_ContainerStateChanged) isRuntimeEventsResponse_Event() {}

func (*RuntimeEventsResponse_Ping) isRuntimeEventsResponse_Event() {}

type RuntimeEventsResponse_ContainerStateChangedEvent struct {
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

func (x *RuntimeEventsResponse_ContainerStateChangedEvent) Reset() {
	*x = RuntimeEventsResponse_ContainerStateChangedEvent{}
	mi := &file_baepo_node_v1_runtime_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RuntimeEventsResponse_ContainerStateChangedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuntimeEventsResponse_ContainerStateChangedEvent) ProtoMessage() {}

func (x *RuntimeEventsResponse_ContainerStateChangedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_node_v1_runtime_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuntimeEventsResponse_ContainerStateChangedEvent.ProtoReflect.Descriptor instead.
func (*RuntimeEventsResponse_ContainerStateChangedEvent) Descriptor() ([]byte, []int) {
	return file_baepo_node_v1_runtime_proto_rawDescGZIP(), []int{5, 0}
}

func (x *RuntimeEventsResponse_ContainerStateChangedEvent) GetContainerId() string {
	if x != nil {
		return x.ContainerId
	}
	return ""
}

func (x *RuntimeEventsResponse_ContainerStateChangedEvent) GetState() v1.ContainerState {
	if x != nil {
		return x.State
	}
	return v1.ContainerState(0)
}

func (x *RuntimeEventsResponse_ContainerStateChangedEvent) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *RuntimeEventsResponse_ContainerStateChangedEvent) GetExitedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExitedAt
	}
	return nil
}

func (x *RuntimeEventsResponse_ContainerStateChangedEvent) GetExitCode() int32 {
	if x != nil && x.ExitCode != nil {
		return *x.ExitCode
	}
	return 0
}

func (x *RuntimeEventsResponse_ContainerStateChangedEvent) GetExitError() string {
	if x != nil && x.ExitError != nil {
		return *x.ExitError
	}
	return ""
}

func (x *RuntimeEventsResponse_ContainerStateChangedEvent) GetHealthy() bool {
	if x != nil {
		return x.Healthy
	}
	return false
}

func (x *RuntimeEventsResponse_ContainerStateChangedEvent) GetHealthcheckError() string {
	if x != nil && x.HealthcheckError != nil {
		return *x.HealthcheckError
	}
	return ""
}

func (x *RuntimeEventsResponse_ContainerStateChangedEvent) GetRestartCount() int32 {
	if x != nil {
		return x.RestartCount
	}
	return 0
}

type RuntimeEventsResponse_PingEvent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RuntimeEventsResponse_PingEvent) Reset() {
	*x = RuntimeEventsResponse_PingEvent{}
	mi := &file_baepo_node_v1_runtime_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RuntimeEventsResponse_PingEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuntimeEventsResponse_PingEvent) ProtoMessage() {}

func (x *RuntimeEventsResponse_PingEvent) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_node_v1_runtime_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuntimeEventsResponse_PingEvent.ProtoReflect.Descriptor instead.
func (*RuntimeEventsResponse_PingEvent) Descriptor() ([]byte, []int) {
	return file_baepo_node_v1_runtime_proto_rawDescGZIP(), []int{5, 1}
}

var File_baepo_node_v1_runtime_proto protoreflect.FileDescriptor

const file_baepo_node_v1_runtime_proto_rawDesc = "" +
	"\n" +
	"\x1bbaepo/node/v1/runtime.proto\x12\rbaepo.node.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1dbaepo/core/v1/container.proto\"E\n" +
	"\x17RuntimeGetStateResponse\x12\x10\n" +
	"\x03pid\x18\x01 \x01(\x03R\x03pid\x12\x18\n" +
	"\arunning\x18\x02 \x01(\bR\arunning\"/\n" +
	"\x15RuntimeGetLogsRequest\x12\x16\n" +
	"\x06follow\x18\x01 \x01(\bR\x06follow\"l\n" +
	"\x16RuntimeGetLogsResponse\x12\x18\n" +
	"\acontent\x18\x01 \x01(\fR\acontent\x128\n" +
	"\ttimestamp\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\ttimestamp\"V\n" +
	"\x1eRuntimeGetContainerLogsRequest\x12\x1c\n" +
	"\tcontainer\x18\x01 \x01(\tR\tcontainer\x12\x16\n" +
	"\x06follow\x18\x02 \x01(\bR\x06follow\"\xae\x01\n" +
	"\x1fRuntimeGetContainerLogsResponse\x12!\n" +
	"\fcontainer_id\x18\x01 \x01(\tR\vcontainerId\x12\x14\n" +
	"\x05error\x18\x02 \x01(\bR\x05error\x12\x18\n" +
	"\acontent\x18\x03 \x01(\fR\acontent\x128\n" +
	"\ttimestamp\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\ttimestamp\"\xbf\x06\n" +
	"\x15RuntimeEventsResponse\x12\x19\n" +
	"\bevent_id\x18\x01 \x01(\tR\aeventId\x128\n" +
	"\ttimestamp\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\ttimestamp\x12y\n" +
	"\x17container_state_changed\x18\x03 \x01(\v2?.baepo.node.v1.RuntimeEventsResponse.ContainerStateChangedEventH\x00R\x15containerStateChanged\x12D\n" +
	"\x04ping\x18\x04 \x01(\v2..baepo.node.v1.RuntimeEventsResponse.PingEventH\x00R\x04ping\x1a\xf9\x03\n" +
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
	"\x05event2\xee\x02\n" +
	"\aRuntime\x12J\n" +
	"\bGetState\x12\x16.google.protobuf.Empty\x1a&.baepo.node.v1.RuntimeGetStateResponse\x12X\n" +
	"\aGetLogs\x12$.baepo.node.v1.RuntimeGetLogsRequest\x1a%.baepo.node.v1.RuntimeGetLogsResponse0\x01\x12s\n" +
	"\x10GetContainerLogs\x12-.baepo.node.v1.RuntimeGetContainerLogsRequest\x1a..baepo.node.v1.RuntimeGetContainerLogsResponse0\x01\x12H\n" +
	"\x06Events\x12\x16.google.protobuf.Empty\x1a$.baepo.node.v1.RuntimeEventsResponse0\x01B>Z<github.com/baepo-cloud/baepo-proto/go/baepo/node/v1;nodev1pbb\x06proto3"

var (
	file_baepo_node_v1_runtime_proto_rawDescOnce sync.Once
	file_baepo_node_v1_runtime_proto_rawDescData []byte
)

func file_baepo_node_v1_runtime_proto_rawDescGZIP() []byte {
	file_baepo_node_v1_runtime_proto_rawDescOnce.Do(func() {
		file_baepo_node_v1_runtime_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_baepo_node_v1_runtime_proto_rawDesc), len(file_baepo_node_v1_runtime_proto_rawDesc)))
	})
	return file_baepo_node_v1_runtime_proto_rawDescData
}

var file_baepo_node_v1_runtime_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_baepo_node_v1_runtime_proto_goTypes = []any{
	(*RuntimeGetStateResponse)(nil),                          // 0: baepo.node.v1.RuntimeGetStateResponse
	(*RuntimeGetLogsRequest)(nil),                            // 1: baepo.node.v1.RuntimeGetLogsRequest
	(*RuntimeGetLogsResponse)(nil),                           // 2: baepo.node.v1.RuntimeGetLogsResponse
	(*RuntimeGetContainerLogsRequest)(nil),                   // 3: baepo.node.v1.RuntimeGetContainerLogsRequest
	(*RuntimeGetContainerLogsResponse)(nil),                  // 4: baepo.node.v1.RuntimeGetContainerLogsResponse
	(*RuntimeEventsResponse)(nil),                            // 5: baepo.node.v1.RuntimeEventsResponse
	(*RuntimeEventsResponse_ContainerStateChangedEvent)(nil), // 6: baepo.node.v1.RuntimeEventsResponse.ContainerStateChangedEvent
	(*RuntimeEventsResponse_PingEvent)(nil),                  // 7: baepo.node.v1.RuntimeEventsResponse.PingEvent
	(*timestamppb.Timestamp)(nil),                            // 8: google.protobuf.Timestamp
	(v1.ContainerState)(0),                                   // 9: baepo.core.v1.ContainerState
	(*emptypb.Empty)(nil),                                    // 10: google.protobuf.Empty
}
var file_baepo_node_v1_runtime_proto_depIdxs = []int32{
	8,  // 0: baepo.node.v1.RuntimeGetLogsResponse.timestamp:type_name -> google.protobuf.Timestamp
	8,  // 1: baepo.node.v1.RuntimeGetContainerLogsResponse.timestamp:type_name -> google.protobuf.Timestamp
	8,  // 2: baepo.node.v1.RuntimeEventsResponse.timestamp:type_name -> google.protobuf.Timestamp
	6,  // 3: baepo.node.v1.RuntimeEventsResponse.container_state_changed:type_name -> baepo.node.v1.RuntimeEventsResponse.ContainerStateChangedEvent
	7,  // 4: baepo.node.v1.RuntimeEventsResponse.ping:type_name -> baepo.node.v1.RuntimeEventsResponse.PingEvent
	9,  // 5: baepo.node.v1.RuntimeEventsResponse.ContainerStateChangedEvent.state:type_name -> baepo.core.v1.ContainerState
	8,  // 6: baepo.node.v1.RuntimeEventsResponse.ContainerStateChangedEvent.started_at:type_name -> google.protobuf.Timestamp
	8,  // 7: baepo.node.v1.RuntimeEventsResponse.ContainerStateChangedEvent.exited_at:type_name -> google.protobuf.Timestamp
	10, // 8: baepo.node.v1.Runtime.GetState:input_type -> google.protobuf.Empty
	1,  // 9: baepo.node.v1.Runtime.GetLogs:input_type -> baepo.node.v1.RuntimeGetLogsRequest
	3,  // 10: baepo.node.v1.Runtime.GetContainerLogs:input_type -> baepo.node.v1.RuntimeGetContainerLogsRequest
	10, // 11: baepo.node.v1.Runtime.Events:input_type -> google.protobuf.Empty
	0,  // 12: baepo.node.v1.Runtime.GetState:output_type -> baepo.node.v1.RuntimeGetStateResponse
	2,  // 13: baepo.node.v1.Runtime.GetLogs:output_type -> baepo.node.v1.RuntimeGetLogsResponse
	4,  // 14: baepo.node.v1.Runtime.GetContainerLogs:output_type -> baepo.node.v1.RuntimeGetContainerLogsResponse
	5,  // 15: baepo.node.v1.Runtime.Events:output_type -> baepo.node.v1.RuntimeEventsResponse
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_baepo_node_v1_runtime_proto_init() }
func file_baepo_node_v1_runtime_proto_init() {
	if File_baepo_node_v1_runtime_proto != nil {
		return
	}
	file_baepo_node_v1_runtime_proto_msgTypes[5].OneofWrappers = []any{
		(*RuntimeEventsResponse_ContainerStateChanged)(nil),
		(*RuntimeEventsResponse_Ping)(nil),
	}
	file_baepo_node_v1_runtime_proto_msgTypes[6].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_baepo_node_v1_runtime_proto_rawDesc), len(file_baepo_node_v1_runtime_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_baepo_node_v1_runtime_proto_goTypes,
		DependencyIndexes: file_baepo_node_v1_runtime_proto_depIdxs,
		MessageInfos:      file_baepo_node_v1_runtime_proto_msgTypes,
	}.Build()
	File_baepo_node_v1_runtime_proto = out.File
	file_baepo_node_v1_runtime_proto_goTypes = nil
	file_baepo_node_v1_runtime_proto_depIdxs = nil
}
