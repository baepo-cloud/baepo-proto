// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: baepo/core/v1/container.proto

package corev1pb

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

type ContainerState int32

const (
	ContainerState_ContainerState_Unknown ContainerState = 0
	ContainerState_ContainerState_Running ContainerState = 1
	ContainerState_ContainerState_Exited  ContainerState = 2
)

// Enum value maps for ContainerState.
var (
	ContainerState_name = map[int32]string{
		0: "ContainerState_Unknown",
		1: "ContainerState_Running",
		2: "ContainerState_Exited",
	}
	ContainerState_value = map[string]int32{
		"ContainerState_Unknown": 0,
		"ContainerState_Running": 1,
		"ContainerState_Exited":  2,
	}
)

func (x ContainerState) Enum() *ContainerState {
	p := new(ContainerState)
	*p = x
	return p
}

func (x ContainerState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ContainerState) Descriptor() protoreflect.EnumDescriptor {
	return file_baepo_core_v1_container_proto_enumTypes[0].Descriptor()
}

func (ContainerState) Type() protoreflect.EnumType {
	return &file_baepo_core_v1_container_proto_enumTypes[0]
}

func (x ContainerState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ContainerState.Descriptor instead.
func (ContainerState) EnumDescriptor() ([]byte, []int) {
	return file_baepo_core_v1_container_proto_rawDescGZIP(), []int{0}
}

type ContainerRestartSpec_Policy int32

const (
	ContainerRestartSpec_Policy_Unknown   ContainerRestartSpec_Policy = 0
	ContainerRestartSpec_Policy_No        ContainerRestartSpec_Policy = 1
	ContainerRestartSpec_Policy_OnFailure ContainerRestartSpec_Policy = 2
	ContainerRestartSpec_Policy_Always    ContainerRestartSpec_Policy = 3
)

// Enum value maps for ContainerRestartSpec_Policy.
var (
	ContainerRestartSpec_Policy_name = map[int32]string{
		0: "Policy_Unknown",
		1: "Policy_No",
		2: "Policy_OnFailure",
		3: "Policy_Always",
	}
	ContainerRestartSpec_Policy_value = map[string]int32{
		"Policy_Unknown":   0,
		"Policy_No":        1,
		"Policy_OnFailure": 2,
		"Policy_Always":    3,
	}
)

func (x ContainerRestartSpec_Policy) Enum() *ContainerRestartSpec_Policy {
	p := new(ContainerRestartSpec_Policy)
	*p = x
	return p
}

func (x ContainerRestartSpec_Policy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ContainerRestartSpec_Policy) Descriptor() protoreflect.EnumDescriptor {
	return file_baepo_core_v1_container_proto_enumTypes[1].Descriptor()
}

func (ContainerRestartSpec_Policy) Type() protoreflect.EnumType {
	return &file_baepo_core_v1_container_proto_enumTypes[1]
}

func (x ContainerRestartSpec_Policy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ContainerRestartSpec_Policy.Descriptor instead.
func (ContainerRestartSpec_Policy) EnumDescriptor() ([]byte, []int) {
	return file_baepo_core_v1_container_proto_rawDescGZIP(), []int{1, 0}
}

type ContainerSpec struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	Name          *string                   `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Image         string                    `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	Env           map[string]string         `protobuf:"bytes,3,rep,name=env,proto3" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Command       []string                  `protobuf:"bytes,4,rep,name=command,proto3" json:"command,omitempty"`
	Healthcheck   *ContainerHealthcheckSpec `protobuf:"bytes,5,opt,name=healthcheck,proto3" json:"healthcheck,omitempty"`
	WorkingDir    *string                   `protobuf:"bytes,6,opt,name=working_dir,json=workingDir,proto3,oneof" json:"working_dir,omitempty"`
	Restart       *ContainerRestartSpec     `protobuf:"bytes,7,opt,name=restart,proto3" json:"restart,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ContainerSpec) Reset() {
	*x = ContainerSpec{}
	mi := &file_baepo_core_v1_container_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ContainerSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerSpec) ProtoMessage() {}

func (x *ContainerSpec) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_core_v1_container_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerSpec.ProtoReflect.Descriptor instead.
func (*ContainerSpec) Descriptor() ([]byte, []int) {
	return file_baepo_core_v1_container_proto_rawDescGZIP(), []int{0}
}

func (x *ContainerSpec) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *ContainerSpec) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *ContainerSpec) GetEnv() map[string]string {
	if x != nil {
		return x.Env
	}
	return nil
}

func (x *ContainerSpec) GetCommand() []string {
	if x != nil {
		return x.Command
	}
	return nil
}

func (x *ContainerSpec) GetHealthcheck() *ContainerHealthcheckSpec {
	if x != nil {
		return x.Healthcheck
	}
	return nil
}

func (x *ContainerSpec) GetWorkingDir() string {
	if x != nil && x.WorkingDir != nil {
		return *x.WorkingDir
	}
	return ""
}

func (x *ContainerSpec) GetRestart() *ContainerRestartSpec {
	if x != nil {
		return x.Restart
	}
	return nil
}

type ContainerRestartSpec struct {
	state         protoimpl.MessageState      `protogen:"open.v1"`
	Policy        ContainerRestartSpec_Policy `protobuf:"varint,1,opt,name=policy,proto3,enum=baepo.core.v1.ContainerRestartSpec_Policy" json:"policy,omitempty"`
	MaxRetries    int32                       `protobuf:"varint,2,opt,name=max_retries,json=maxRetries,proto3" json:"max_retries,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ContainerRestartSpec) Reset() {
	*x = ContainerRestartSpec{}
	mi := &file_baepo_core_v1_container_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ContainerRestartSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerRestartSpec) ProtoMessage() {}

func (x *ContainerRestartSpec) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_core_v1_container_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerRestartSpec.ProtoReflect.Descriptor instead.
func (*ContainerRestartSpec) Descriptor() ([]byte, []int) {
	return file_baepo_core_v1_container_proto_rawDescGZIP(), []int{1}
}

func (x *ContainerRestartSpec) GetPolicy() ContainerRestartSpec_Policy {
	if x != nil {
		return x.Policy
	}
	return ContainerRestartSpec_Policy_Unknown
}

func (x *ContainerRestartSpec) GetMaxRetries() int32 {
	if x != nil {
		return x.MaxRetries
	}
	return 0
}

type ContainerHealthcheckSpec struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	InitialDelaySeconds int32                  `protobuf:"varint,1,opt,name=initial_delay_seconds,json=initialDelaySeconds,proto3" json:"initial_delay_seconds,omitempty"`
	PeriodSeconds       int32                  `protobuf:"varint,2,opt,name=period_seconds,json=periodSeconds,proto3" json:"period_seconds,omitempty"`
	// Types that are valid to be assigned to Type:
	//
	//	*ContainerHealthcheckSpec_Http
	Type          isContainerHealthcheckSpec_Type `protobuf_oneof:"type"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ContainerHealthcheckSpec) Reset() {
	*x = ContainerHealthcheckSpec{}
	mi := &file_baepo_core_v1_container_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ContainerHealthcheckSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerHealthcheckSpec) ProtoMessage() {}

func (x *ContainerHealthcheckSpec) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_core_v1_container_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerHealthcheckSpec.ProtoReflect.Descriptor instead.
func (*ContainerHealthcheckSpec) Descriptor() ([]byte, []int) {
	return file_baepo_core_v1_container_proto_rawDescGZIP(), []int{2}
}

func (x *ContainerHealthcheckSpec) GetInitialDelaySeconds() int32 {
	if x != nil {
		return x.InitialDelaySeconds
	}
	return 0
}

func (x *ContainerHealthcheckSpec) GetPeriodSeconds() int32 {
	if x != nil {
		return x.PeriodSeconds
	}
	return 0
}

func (x *ContainerHealthcheckSpec) GetType() isContainerHealthcheckSpec_Type {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *ContainerHealthcheckSpec) GetHttp() *ContainerHealthcheckSpec_HttpHealthcheckSpec {
	if x != nil {
		if x, ok := x.Type.(*ContainerHealthcheckSpec_Http); ok {
			return x.Http
		}
	}
	return nil
}

type isContainerHealthcheckSpec_Type interface {
	isContainerHealthcheckSpec_Type()
}

type ContainerHealthcheckSpec_Http struct {
	Http *ContainerHealthcheckSpec_HttpHealthcheckSpec `protobuf:"bytes,3,opt,name=http,proto3,oneof"`
}

func (*ContainerHealthcheckSpec_Http) isContainerHealthcheckSpec_Type() {}

type ContainerEvent struct {
	state       protoimpl.MessageState `protogen:"open.v1"`
	EventId     string                 `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	ContainerId string                 `protobuf:"bytes,2,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	Timestamp   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Types that are valid to be assigned to Event:
	//
	//	*ContainerEvent_StateChanged
	Event         isContainerEvent_Event `protobuf_oneof:"event"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ContainerEvent) Reset() {
	*x = ContainerEvent{}
	mi := &file_baepo_core_v1_container_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ContainerEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerEvent) ProtoMessage() {}

func (x *ContainerEvent) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_core_v1_container_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerEvent.ProtoReflect.Descriptor instead.
func (*ContainerEvent) Descriptor() ([]byte, []int) {
	return file_baepo_core_v1_container_proto_rawDescGZIP(), []int{3}
}

func (x *ContainerEvent) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *ContainerEvent) GetContainerId() string {
	if x != nil {
		return x.ContainerId
	}
	return ""
}

func (x *ContainerEvent) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *ContainerEvent) GetEvent() isContainerEvent_Event {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *ContainerEvent) GetStateChanged() *ContainerEvent_StateChangedEvent {
	if x != nil {
		if x, ok := x.Event.(*ContainerEvent_StateChanged); ok {
			return x.StateChanged
		}
	}
	return nil
}

type isContainerEvent_Event interface {
	isContainerEvent_Event()
}

type ContainerEvent_StateChanged struct {
	StateChanged *ContainerEvent_StateChangedEvent `protobuf:"bytes,4,opt,name=state_changed,json=stateChanged,proto3,oneof"`
}

func (*ContainerEvent_StateChanged) isContainerEvent_Event() {}

type ContainerHealthcheckSpec_HttpHealthcheckSpec struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Method        string                 `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Path          string                 `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Port          int32                  `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	Headers       map[string]string      `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ContainerHealthcheckSpec_HttpHealthcheckSpec) Reset() {
	*x = ContainerHealthcheckSpec_HttpHealthcheckSpec{}
	mi := &file_baepo_core_v1_container_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ContainerHealthcheckSpec_HttpHealthcheckSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerHealthcheckSpec_HttpHealthcheckSpec) ProtoMessage() {}

func (x *ContainerHealthcheckSpec_HttpHealthcheckSpec) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_core_v1_container_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerHealthcheckSpec_HttpHealthcheckSpec.ProtoReflect.Descriptor instead.
func (*ContainerHealthcheckSpec_HttpHealthcheckSpec) Descriptor() ([]byte, []int) {
	return file_baepo_core_v1_container_proto_rawDescGZIP(), []int{2, 0}
}

func (x *ContainerHealthcheckSpec_HttpHealthcheckSpec) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *ContainerHealthcheckSpec_HttpHealthcheckSpec) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ContainerHealthcheckSpec_HttpHealthcheckSpec) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *ContainerHealthcheckSpec_HttpHealthcheckSpec) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

type ContainerEvent_StateChangedEvent struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	State            ContainerState         `protobuf:"varint,1,opt,name=state,proto3,enum=baepo.core.v1.ContainerState" json:"state,omitempty"`
	StartedAt        *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=started_at,json=startedAt,proto3,oneof" json:"started_at,omitempty"`
	ExitedAt         *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=exited_at,json=exitedAt,proto3,oneof" json:"exited_at,omitempty"`
	ExitCode         *int32                 `protobuf:"varint,4,opt,name=exit_code,json=exitCode,proto3,oneof" json:"exit_code,omitempty"`
	ExitError        *string                `protobuf:"bytes,5,opt,name=exit_error,json=exitError,proto3,oneof" json:"exit_error,omitempty"`
	Healthy          bool                   `protobuf:"varint,6,opt,name=healthy,proto3" json:"healthy,omitempty"`
	HealthcheckError *string                `protobuf:"bytes,7,opt,name=healthcheck_error,json=healthcheckError,proto3,oneof" json:"healthcheck_error,omitempty"`
	RestartCount     int32                  `protobuf:"varint,8,opt,name=restart_count,json=restartCount,proto3" json:"restart_count,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *ContainerEvent_StateChangedEvent) Reset() {
	*x = ContainerEvent_StateChangedEvent{}
	mi := &file_baepo_core_v1_container_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ContainerEvent_StateChangedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerEvent_StateChangedEvent) ProtoMessage() {}

func (x *ContainerEvent_StateChangedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_baepo_core_v1_container_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerEvent_StateChangedEvent.ProtoReflect.Descriptor instead.
func (*ContainerEvent_StateChangedEvent) Descriptor() ([]byte, []int) {
	return file_baepo_core_v1_container_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ContainerEvent_StateChangedEvent) GetState() ContainerState {
	if x != nil {
		return x.State
	}
	return ContainerState_ContainerState_Unknown
}

func (x *ContainerEvent_StateChangedEvent) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *ContainerEvent_StateChangedEvent) GetExitedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExitedAt
	}
	return nil
}

func (x *ContainerEvent_StateChangedEvent) GetExitCode() int32 {
	if x != nil && x.ExitCode != nil {
		return *x.ExitCode
	}
	return 0
}

func (x *ContainerEvent_StateChangedEvent) GetExitError() string {
	if x != nil && x.ExitError != nil {
		return *x.ExitError
	}
	return ""
}

func (x *ContainerEvent_StateChangedEvent) GetHealthy() bool {
	if x != nil {
		return x.Healthy
	}
	return false
}

func (x *ContainerEvent_StateChangedEvent) GetHealthcheckError() string {
	if x != nil && x.HealthcheckError != nil {
		return *x.HealthcheckError
	}
	return ""
}

func (x *ContainerEvent_StateChangedEvent) GetRestartCount() int32 {
	if x != nil {
		return x.RestartCount
	}
	return 0
}

var File_baepo_core_v1_container_proto protoreflect.FileDescriptor

const file_baepo_core_v1_container_proto_rawDesc = "" +
	"\n" +
	"\x1dbaepo/core/v1/container.proto\x12\rbaepo.core.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\x92\x03\n" +
	"\rContainerSpec\x12\x17\n" +
	"\x04name\x18\x01 \x01(\tH\x00R\x04name\x88\x01\x01\x12\x14\n" +
	"\x05image\x18\x02 \x01(\tR\x05image\x127\n" +
	"\x03env\x18\x03 \x03(\v2%.baepo.core.v1.ContainerSpec.EnvEntryR\x03env\x12\x18\n" +
	"\acommand\x18\x04 \x03(\tR\acommand\x12I\n" +
	"\vhealthcheck\x18\x05 \x01(\v2'.baepo.core.v1.ContainerHealthcheckSpecR\vhealthcheck\x12$\n" +
	"\vworking_dir\x18\x06 \x01(\tH\x01R\n" +
	"workingDir\x88\x01\x01\x12=\n" +
	"\arestart\x18\a \x01(\v2#.baepo.core.v1.ContainerRestartSpecR\arestart\x1a6\n" +
	"\bEnvEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01B\a\n" +
	"\x05_nameB\x0e\n" +
	"\f_working_dir\"\xd1\x01\n" +
	"\x14ContainerRestartSpec\x12B\n" +
	"\x06policy\x18\x01 \x01(\x0e2*.baepo.core.v1.ContainerRestartSpec.PolicyR\x06policy\x12\x1f\n" +
	"\vmax_retries\x18\x02 \x01(\x05R\n" +
	"maxRetries\"T\n" +
	"\x06Policy\x12\x12\n" +
	"\x0ePolicy_Unknown\x10\x00\x12\r\n" +
	"\tPolicy_No\x10\x01\x12\x14\n" +
	"\x10Policy_OnFailure\x10\x02\x12\x11\n" +
	"\rPolicy_Always\x10\x03\"\xc8\x03\n" +
	"\x18ContainerHealthcheckSpec\x122\n" +
	"\x15initial_delay_seconds\x18\x01 \x01(\x05R\x13initialDelaySeconds\x12%\n" +
	"\x0eperiod_seconds\x18\x02 \x01(\x05R\rperiodSeconds\x12Q\n" +
	"\x04http\x18\x03 \x01(\v2;.baepo.core.v1.ContainerHealthcheckSpec.HttpHealthcheckSpecH\x00R\x04http\x1a\xf5\x01\n" +
	"\x13HttpHealthcheckSpec\x12\x16\n" +
	"\x06method\x18\x01 \x01(\tR\x06method\x12\x12\n" +
	"\x04path\x18\x02 \x01(\tR\x04path\x12\x12\n" +
	"\x04port\x18\x03 \x01(\x05R\x04port\x12b\n" +
	"\aheaders\x18\x04 \x03(\v2H.baepo.core.v1.ContainerHealthcheckSpec.HttpHealthcheckSpec.HeadersEntryR\aheaders\x1a:\n" +
	"\fHeadersEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01B\x06\n" +
	"\x04type\"\xb9\x05\n" +
	"\x0eContainerEvent\x12\x19\n" +
	"\bevent_id\x18\x01 \x01(\tR\aeventId\x12!\n" +
	"\fcontainer_id\x18\x02 \x01(\tR\vcontainerId\x128\n" +
	"\ttimestamp\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\ttimestamp\x12V\n" +
	"\rstate_changed\x18\x04 \x01(\v2/.baepo.core.v1.ContainerEvent.StateChangedEventH\x00R\fstateChanged\x1a\xcd\x03\n" +
	"\x11StateChangedEvent\x123\n" +
	"\x05state\x18\x01 \x01(\x0e2\x1d.baepo.core.v1.ContainerStateR\x05state\x12>\n" +
	"\n" +
	"started_at\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampH\x00R\tstartedAt\x88\x01\x01\x12<\n" +
	"\texited_at\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampH\x01R\bexitedAt\x88\x01\x01\x12 \n" +
	"\texit_code\x18\x04 \x01(\x05H\x02R\bexitCode\x88\x01\x01\x12\"\n" +
	"\n" +
	"exit_error\x18\x05 \x01(\tH\x03R\texitError\x88\x01\x01\x12\x18\n" +
	"\ahealthy\x18\x06 \x01(\bR\ahealthy\x120\n" +
	"\x11healthcheck_error\x18\a \x01(\tH\x04R\x10healthcheckError\x88\x01\x01\x12#\n" +
	"\rrestart_count\x18\b \x01(\x05R\frestartCountB\r\n" +
	"\v_started_atB\f\n" +
	"\n" +
	"_exited_atB\f\n" +
	"\n" +
	"_exit_codeB\r\n" +
	"\v_exit_errorB\x14\n" +
	"\x12_healthcheck_errorB\a\n" +
	"\x05event*c\n" +
	"\x0eContainerState\x12\x1a\n" +
	"\x16ContainerState_Unknown\x10\x00\x12\x1a\n" +
	"\x16ContainerState_Running\x10\x01\x12\x19\n" +
	"\x15ContainerState_Exited\x10\x02B>Z<github.com/baepo-cloud/baepo-proto/go/baepo/core/v1;corev1pbb\x06proto3"

var (
	file_baepo_core_v1_container_proto_rawDescOnce sync.Once
	file_baepo_core_v1_container_proto_rawDescData []byte
)

func file_baepo_core_v1_container_proto_rawDescGZIP() []byte {
	file_baepo_core_v1_container_proto_rawDescOnce.Do(func() {
		file_baepo_core_v1_container_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_baepo_core_v1_container_proto_rawDesc), len(file_baepo_core_v1_container_proto_rawDesc)))
	})
	return file_baepo_core_v1_container_proto_rawDescData
}

var file_baepo_core_v1_container_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_baepo_core_v1_container_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_baepo_core_v1_container_proto_goTypes = []any{
	(ContainerState)(0),              // 0: baepo.core.v1.ContainerState
	(ContainerRestartSpec_Policy)(0), // 1: baepo.core.v1.ContainerRestartSpec.Policy
	(*ContainerSpec)(nil),            // 2: baepo.core.v1.ContainerSpec
	(*ContainerRestartSpec)(nil),     // 3: baepo.core.v1.ContainerRestartSpec
	(*ContainerHealthcheckSpec)(nil), // 4: baepo.core.v1.ContainerHealthcheckSpec
	(*ContainerEvent)(nil),           // 5: baepo.core.v1.ContainerEvent
	nil,                              // 6: baepo.core.v1.ContainerSpec.EnvEntry
	(*ContainerHealthcheckSpec_HttpHealthcheckSpec)(nil), // 7: baepo.core.v1.ContainerHealthcheckSpec.HttpHealthcheckSpec
	nil,                                      // 8: baepo.core.v1.ContainerHealthcheckSpec.HttpHealthcheckSpec.HeadersEntry
	(*ContainerEvent_StateChangedEvent)(nil), // 9: baepo.core.v1.ContainerEvent.StateChangedEvent
	(*timestamppb.Timestamp)(nil),            // 10: google.protobuf.Timestamp
}
var file_baepo_core_v1_container_proto_depIdxs = []int32{
	6,  // 0: baepo.core.v1.ContainerSpec.env:type_name -> baepo.core.v1.ContainerSpec.EnvEntry
	4,  // 1: baepo.core.v1.ContainerSpec.healthcheck:type_name -> baepo.core.v1.ContainerHealthcheckSpec
	3,  // 2: baepo.core.v1.ContainerSpec.restart:type_name -> baepo.core.v1.ContainerRestartSpec
	1,  // 3: baepo.core.v1.ContainerRestartSpec.policy:type_name -> baepo.core.v1.ContainerRestartSpec.Policy
	7,  // 4: baepo.core.v1.ContainerHealthcheckSpec.http:type_name -> baepo.core.v1.ContainerHealthcheckSpec.HttpHealthcheckSpec
	10, // 5: baepo.core.v1.ContainerEvent.timestamp:type_name -> google.protobuf.Timestamp
	9,  // 6: baepo.core.v1.ContainerEvent.state_changed:type_name -> baepo.core.v1.ContainerEvent.StateChangedEvent
	8,  // 7: baepo.core.v1.ContainerHealthcheckSpec.HttpHealthcheckSpec.headers:type_name -> baepo.core.v1.ContainerHealthcheckSpec.HttpHealthcheckSpec.HeadersEntry
	0,  // 8: baepo.core.v1.ContainerEvent.StateChangedEvent.state:type_name -> baepo.core.v1.ContainerState
	10, // 9: baepo.core.v1.ContainerEvent.StateChangedEvent.started_at:type_name -> google.protobuf.Timestamp
	10, // 10: baepo.core.v1.ContainerEvent.StateChangedEvent.exited_at:type_name -> google.protobuf.Timestamp
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_baepo_core_v1_container_proto_init() }
func file_baepo_core_v1_container_proto_init() {
	if File_baepo_core_v1_container_proto != nil {
		return
	}
	file_baepo_core_v1_container_proto_msgTypes[0].OneofWrappers = []any{}
	file_baepo_core_v1_container_proto_msgTypes[2].OneofWrappers = []any{
		(*ContainerHealthcheckSpec_Http)(nil),
	}
	file_baepo_core_v1_container_proto_msgTypes[3].OneofWrappers = []any{
		(*ContainerEvent_StateChanged)(nil),
	}
	file_baepo_core_v1_container_proto_msgTypes[7].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_baepo_core_v1_container_proto_rawDesc), len(file_baepo_core_v1_container_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_baepo_core_v1_container_proto_goTypes,
		DependencyIndexes: file_baepo_core_v1_container_proto_depIdxs,
		EnumInfos:         file_baepo_core_v1_container_proto_enumTypes,
		MessageInfos:      file_baepo_core_v1_container_proto_msgTypes,
	}.Build()
	File_baepo_core_v1_container_proto = out.File
	file_baepo_core_v1_container_proto_goTypes = nil
	file_baepo_core_v1_container_proto_depIdxs = nil
}
