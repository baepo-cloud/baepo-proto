// @generated by protoc-gen-es v2.2.5 with parameter "import_extension=js"
// @generated from file baepo/api/v1/node_controller.proto (package baepo.api.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import type { Message } from "@bufbuild/protobuf";
import type { MachineDesiredState, MachineEvent, MachineSpec } from "../../core/v1/machine_pb.js";
import type { ContainerEvent, ContainerSpec } from "../../core/v1/container_pb.js";

/**
 * Describes the file baepo/api/v1/node_controller.proto.
 */
export declare const file_baepo_api_v1_node_controller: GenFile;

/**
 * @generated from message baepo.api.v1.NodeControllerServerEvent
 */
export declare type NodeControllerServerEvent = Message<"baepo.api.v1.NodeControllerServerEvent"> & {
  /**
   * @generated from oneof baepo.api.v1.NodeControllerServerEvent.event
   */
  event: {
    /**
     * @generated from field: baepo.api.v1.NodeControllerServerEvent.RegistrationCompletedEvent registration_completed = 1;
     */
    value: NodeControllerServerEvent_RegistrationCompletedEvent;
    case: "registrationCompleted";
  } | {
    /**
     * @generated from field: baepo.api.v1.NodeControllerServerEvent.PingEvent ping = 2;
     */
    value: NodeControllerServerEvent_PingEvent;
    case: "ping";
  } | {
    /**
     * @generated from field: baepo.api.v1.NodeControllerServerEvent.Machine create_machine = 3;
     */
    value: NodeControllerServerEvent_Machine;
    case: "createMachine";
  } | {
    /**
     * @generated from field: baepo.api.v1.NodeControllerServerEvent.UpdateMachineDesiredStateEvent update_machine_desired_state = 4;
     */
    value: NodeControllerServerEvent_UpdateMachineDesiredStateEvent;
    case: "updateMachineDesiredState";
  } | { case: undefined; value?: undefined };
};

/**
 * Describes the message baepo.api.v1.NodeControllerServerEvent.
 * Use `create(NodeControllerServerEventSchema)` to create a new message.
 */
export declare const NodeControllerServerEventSchema: GenMessage<NodeControllerServerEvent>;

/**
 * @generated from message baepo.api.v1.NodeControllerServerEvent.RegistrationCompletedEvent
 */
export declare type NodeControllerServerEvent_RegistrationCompletedEvent = Message<"baepo.api.v1.NodeControllerServerEvent.RegistrationCompletedEvent"> & {
  /**
   * @generated from field: string node_id = 1;
   */
  nodeId: string;

  /**
   * @generated from field: string node_token = 2;
   */
  nodeToken: string;

  /**
   * @generated from field: bytes authority_cert = 3;
   */
  authorityCert: Uint8Array;

  /**
   * @generated from field: bytes server_cert = 4;
   */
  serverCert: Uint8Array;

  /**
   * @generated from field: bytes server_key = 5;
   */
  serverKey: Uint8Array;

  /**
   * @generated from field: repeated baepo.api.v1.NodeControllerServerEvent.Machine expected_machines = 6;
   */
  expectedMachines: NodeControllerServerEvent_Machine[];
};

/**
 * Describes the message baepo.api.v1.NodeControllerServerEvent.RegistrationCompletedEvent.
 * Use `create(NodeControllerServerEvent_RegistrationCompletedEventSchema)` to create a new message.
 */
export declare const NodeControllerServerEvent_RegistrationCompletedEventSchema: GenMessage<NodeControllerServerEvent_RegistrationCompletedEvent>;

/**
 * @generated from message baepo.api.v1.NodeControllerServerEvent.PingEvent
 */
export declare type NodeControllerServerEvent_PingEvent = Message<"baepo.api.v1.NodeControllerServerEvent.PingEvent"> & {
};

/**
 * Describes the message baepo.api.v1.NodeControllerServerEvent.PingEvent.
 * Use `create(NodeControllerServerEvent_PingEventSchema)` to create a new message.
 */
export declare const NodeControllerServerEvent_PingEventSchema: GenMessage<NodeControllerServerEvent_PingEvent>;

/**
 * @generated from message baepo.api.v1.NodeControllerServerEvent.UpdateMachineDesiredStateEvent
 */
export declare type NodeControllerServerEvent_UpdateMachineDesiredStateEvent = Message<"baepo.api.v1.NodeControllerServerEvent.UpdateMachineDesiredStateEvent"> & {
  /**
   * @generated from field: string machine_id = 1;
   */
  machineId: string;

  /**
   * @generated from field: baepo.core.v1.MachineDesiredState desired_state = 2;
   */
  desiredState: MachineDesiredState;
};

/**
 * Describes the message baepo.api.v1.NodeControllerServerEvent.UpdateMachineDesiredStateEvent.
 * Use `create(NodeControllerServerEvent_UpdateMachineDesiredStateEventSchema)` to create a new message.
 */
export declare const NodeControllerServerEvent_UpdateMachineDesiredStateEventSchema: GenMessage<NodeControllerServerEvent_UpdateMachineDesiredStateEvent>;

/**
 * @generated from message baepo.api.v1.NodeControllerServerEvent.Machine
 */
export declare type NodeControllerServerEvent_Machine = Message<"baepo.api.v1.NodeControllerServerEvent.Machine"> & {
  /**
   * @generated from field: string machine_id = 1;
   */
  machineId: string;

  /**
   * @generated from field: baepo.core.v1.MachineDesiredState desired_state = 2;
   */
  desiredState: MachineDesiredState;

  /**
   * @generated from field: baepo.core.v1.MachineSpec spec = 3;
   */
  spec?: MachineSpec;

  /**
   * @generated from field: repeated baepo.api.v1.NodeControllerServerEvent.Container containers = 4;
   */
  containers: NodeControllerServerEvent_Container[];
};

/**
 * Describes the message baepo.api.v1.NodeControllerServerEvent.Machine.
 * Use `create(NodeControllerServerEvent_MachineSchema)` to create a new message.
 */
export declare const NodeControllerServerEvent_MachineSchema: GenMessage<NodeControllerServerEvent_Machine>;

/**
 * @generated from message baepo.api.v1.NodeControllerServerEvent.Container
 */
export declare type NodeControllerServerEvent_Container = Message<"baepo.api.v1.NodeControllerServerEvent.Container"> & {
  /**
   * @generated from field: string container_id = 1;
   */
  containerId: string;

  /**
   * @generated from field: baepo.core.v1.ContainerSpec spec = 2;
   */
  spec?: ContainerSpec;
};

/**
 * Describes the message baepo.api.v1.NodeControllerServerEvent.Container.
 * Use `create(NodeControllerServerEvent_ContainerSchema)` to create a new message.
 */
export declare const NodeControllerServerEvent_ContainerSchema: GenMessage<NodeControllerServerEvent_Container>;

/**
 * @generated from message baepo.api.v1.NodeControllerClientEvent
 */
export declare type NodeControllerClientEvent = Message<"baepo.api.v1.NodeControllerClientEvent"> & {
  /**
   * @generated from oneof baepo.api.v1.NodeControllerClientEvent.event
   */
  event: {
    /**
     * @generated from field: baepo.api.v1.NodeControllerClientEvent.Register register = 1;
     */
    value: NodeControllerClientEvent_Register;
    case: "register";
  } | {
    /**
     * @generated from field: baepo.api.v1.NodeControllerClientEvent.Stats stats = 2;
     */
    value: NodeControllerClientEvent_Stats;
    case: "stats";
  } | {
    /**
     * @generated from field: baepo.core.v1.MachineEvent machine = 3;
     */
    value: MachineEvent;
    case: "machine";
  } | {
    /**
     * @generated from field: baepo.core.v1.ContainerEvent container = 4;
     */
    value: ContainerEvent;
    case: "container";
  } | { case: undefined; value?: undefined };
};

/**
 * Describes the message baepo.api.v1.NodeControllerClientEvent.
 * Use `create(NodeControllerClientEventSchema)` to create a new message.
 */
export declare const NodeControllerClientEventSchema: GenMessage<NodeControllerClientEvent>;

/**
 * @generated from message baepo.api.v1.NodeControllerClientEvent.Register
 */
export declare type NodeControllerClientEvent_Register = Message<"baepo.api.v1.NodeControllerClientEvent.Register"> & {
  /**
   * @generated from field: string cluster_id = 1;
   */
  clusterId: string;

  /**
   * @generated from field: string bootstrap_token = 2;
   */
  bootstrapToken: string;

  /**
   * @generated from field: optional string node_token = 3;
   */
  nodeToken?: string;

  /**
   * @generated from field: string ip_address = 4;
   */
  ipAddress: string;

  /**
   * @generated from field: string api_endpoint = 5;
   */
  apiEndpoint: string;

  /**
   * @generated from field: string gateway_endpoint = 6;
   */
  gatewayEndpoint: string;

  /**
   * @generated from field: baepo.api.v1.NodeControllerClientEvent.Stats stats = 7;
   */
  stats?: NodeControllerClientEvent_Stats;
};

/**
 * Describes the message baepo.api.v1.NodeControllerClientEvent.Register.
 * Use `create(NodeControllerClientEvent_RegisterSchema)` to create a new message.
 */
export declare const NodeControllerClientEvent_RegisterSchema: GenMessage<NodeControllerClientEvent_Register>;

/**
 * @generated from message baepo.api.v1.NodeControllerClientEvent.Stats
 */
export declare type NodeControllerClientEvent_Stats = Message<"baepo.api.v1.NodeControllerClientEvent.Stats"> & {
  /**
   * @generated from field: uint64 total_memory_mb = 1;
   */
  totalMemoryMb: bigint;

  /**
   * @generated from field: uint64 used_memory_mb = 2;
   */
  usedMemoryMb: bigint;

  /**
   * @generated from field: uint64 reserved_memory_mb = 3;
   */
  reservedMemoryMb: bigint;

  /**
   * @generated from field: uint32 cpu_count = 4;
   */
  cpuCount: number;
};

/**
 * Describes the message baepo.api.v1.NodeControllerClientEvent.Stats.
 * Use `create(NodeControllerClientEvent_StatsSchema)` to create a new message.
 */
export declare const NodeControllerClientEvent_StatsSchema: GenMessage<NodeControllerClientEvent_Stats>;

/**
 * @generated from service baepo.api.v1.NodeControllerService
 */
export declare const NodeControllerService: GenService<{
  /**
   * @generated from rpc baepo.api.v1.NodeControllerService.Events
   */
  events: {
    methodKind: "bidi_streaming";
    input: typeof NodeControllerClientEventSchema;
    output: typeof NodeControllerServerEventSchema;
  },
}>;

