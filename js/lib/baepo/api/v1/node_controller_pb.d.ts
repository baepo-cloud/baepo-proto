// @generated by protoc-gen-es v2.2.5
// @generated from file baepo/api/v1/node_controller.proto (package baepo.api.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file baepo/api/v1/node_controller.proto.
 */
export declare const file_baepo_api_v1_node_controller: GenFile;

/**
 * @generated from message baepo.api.v1.NodeControllerConnectServerEvent
 */
export declare type NodeControllerConnectServerEvent = Message<"baepo.api.v1.NodeControllerConnectServerEvent"> & {
  /**
   * @generated from oneof baepo.api.v1.NodeControllerConnectServerEvent.event
   */
  event: {
    /**
     * @generated from field: baepo.api.v1.NodeControllerConnectServerEvent.RegisterResponse register = 1;
     */
    value: NodeControllerConnectServerEvent_RegisterResponse;
    case: "register";
  } | {
    /**
     * @generated from field: baepo.api.v1.NodeControllerConnectServerEvent.PingEvent ping = 2;
     */
    value: NodeControllerConnectServerEvent_PingEvent;
    case: "ping";
  } | { case: undefined; value?: undefined };
};

/**
 * Describes the message baepo.api.v1.NodeControllerConnectServerEvent.
 * Use `create(NodeControllerConnectServerEventSchema)` to create a new message.
 */
export declare const NodeControllerConnectServerEventSchema: GenMessage<NodeControllerConnectServerEvent>;

/**
 * @generated from message baepo.api.v1.NodeControllerConnectServerEvent.RegisterResponse
 */
export declare type NodeControllerConnectServerEvent_RegisterResponse = Message<"baepo.api.v1.NodeControllerConnectServerEvent.RegisterResponse"> & {
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
};

/**
 * Describes the message baepo.api.v1.NodeControllerConnectServerEvent.RegisterResponse.
 * Use `create(NodeControllerConnectServerEvent_RegisterResponseSchema)` to create a new message.
 */
export declare const NodeControllerConnectServerEvent_RegisterResponseSchema: GenMessage<NodeControllerConnectServerEvent_RegisterResponse>;

/**
 * @generated from message baepo.api.v1.NodeControllerConnectServerEvent.PingEvent
 */
export declare type NodeControllerConnectServerEvent_PingEvent = Message<"baepo.api.v1.NodeControllerConnectServerEvent.PingEvent"> & {
};

/**
 * Describes the message baepo.api.v1.NodeControllerConnectServerEvent.PingEvent.
 * Use `create(NodeControllerConnectServerEvent_PingEventSchema)` to create a new message.
 */
export declare const NodeControllerConnectServerEvent_PingEventSchema: GenMessage<NodeControllerConnectServerEvent_PingEvent>;

/**
 * @generated from message baepo.api.v1.NodeControllerConnectClientEvent
 */
export declare type NodeControllerConnectClientEvent = Message<"baepo.api.v1.NodeControllerConnectClientEvent"> & {
  /**
   * @generated from oneof baepo.api.v1.NodeControllerConnectClientEvent.event
   */
  event: {
    /**
     * @generated from field: baepo.api.v1.NodeControllerConnectClientEvent.RegisterRequest register = 1;
     */
    value: NodeControllerConnectClientEvent_RegisterRequest;
    case: "register";
  } | {
    /**
     * @generated from field: baepo.api.v1.NodeControllerConnectClientEvent.StatsEvent stats = 2;
     */
    value: NodeControllerConnectClientEvent_StatsEvent;
    case: "stats";
  } | { case: undefined; value?: undefined };
};

/**
 * Describes the message baepo.api.v1.NodeControllerConnectClientEvent.
 * Use `create(NodeControllerConnectClientEventSchema)` to create a new message.
 */
export declare const NodeControllerConnectClientEventSchema: GenMessage<NodeControllerConnectClientEvent>;

/**
 * @generated from message baepo.api.v1.NodeControllerConnectClientEvent.RegisterRequest
 */
export declare type NodeControllerConnectClientEvent_RegisterRequest = Message<"baepo.api.v1.NodeControllerConnectClientEvent.RegisterRequest"> & {
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
};

/**
 * Describes the message baepo.api.v1.NodeControllerConnectClientEvent.RegisterRequest.
 * Use `create(NodeControllerConnectClientEvent_RegisterRequestSchema)` to create a new message.
 */
export declare const NodeControllerConnectClientEvent_RegisterRequestSchema: GenMessage<NodeControllerConnectClientEvent_RegisterRequest>;

/**
 * @generated from message baepo.api.v1.NodeControllerConnectClientEvent.StatsEvent
 */
export declare type NodeControllerConnectClientEvent_StatsEvent = Message<"baepo.api.v1.NodeControllerConnectClientEvent.StatsEvent"> & {
  /**
   * @generated from field: uint64 total_memory = 1;
   */
  totalMemory: bigint;

  /**
   * @generated from field: uint64 used_memory = 2;
   */
  usedMemory: bigint;

  /**
   * @generated from field: uint32 cpu_count = 3;
   */
  cpuCount: number;

  /**
   * @generated from field: repeated string running_machine_ids = 4;
   */
  runningMachineIds: string[];

  /**
   * @generated from field: uint64 reserved_memory = 5;
   */
  reservedMemory: bigint;
};

/**
 * Describes the message baepo.api.v1.NodeControllerConnectClientEvent.StatsEvent.
 * Use `create(NodeControllerConnectClientEvent_StatsEventSchema)` to create a new message.
 */
export declare const NodeControllerConnectClientEvent_StatsEventSchema: GenMessage<NodeControllerConnectClientEvent_StatsEvent>;

/**
 * @generated from service baepo.api.v1.NodeControllerService
 */
export declare const NodeControllerService: GenService<{
  /**
   * @generated from rpc baepo.api.v1.NodeControllerService.Connect
   */
  connect: {
    methodKind: "bidi_streaming";
    input: typeof NodeControllerConnectClientEventSchema;
    output: typeof NodeControllerConnectServerEventSchema;
  },
}>;

