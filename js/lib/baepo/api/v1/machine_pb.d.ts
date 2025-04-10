// @generated by protoc-gen-es v2.2.5
// @generated from file baepo/api/v1/machine.proto (package baepo.api.v1, syntax proto3)
/* eslint-disable */

import type { GenEnum, GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import type { Message } from "@bufbuild/protobuf";
import type { Timestamp } from "@bufbuild/protobuf/wkt";

/**
 * Describes the file baepo/api/v1/machine.proto.
 */
export declare const file_baepo_api_v1_machine: GenFile;

/**
 * @generated from message baepo.api.v1.Machine
 */
export declare type Machine = Message<"baepo.api.v1.Machine"> & {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: optional string name = 2;
   */
  name?: string;

  /**
   * @generated from field: baepo.api.v1.MachineStatus status = 3;
   */
  status: MachineStatus;

  /**
   * @generated from field: optional string node_id = 4;
   */
  nodeId?: string;

  /**
   * @generated from field: baepo.api.v1.MachineSpec spec = 5;
   */
  spec?: MachineSpec;

  /**
   * @generated from field: optional uint32 timeout = 6;
   */
  timeout?: number;

  /**
   * @generated from field: optional google.protobuf.Timestamp started_at = 7;
   */
  startedAt?: Timestamp;

  /**
   * @generated from field: optional google.protobuf.Timestamp expires_at = 8;
   */
  expiresAt?: Timestamp;

  /**
   * @generated from field: optional google.protobuf.Timestamp terminated_at = 9;
   */
  terminatedAt?: Timestamp;

  /**
   * @generated from field: optional string termination_cause = 10;
   */
  terminationCause?: string;

  /**
   * @generated from field: optional string termination_details = 11;
   */
  terminationDetails?: string;

  /**
   * @generated from field: map<string, string> metadata = 12;
   */
  metadata: { [key: string]: string };

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 13;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp updated_at = 14;
   */
  updatedAt?: Timestamp;
};

/**
 * Describes the message baepo.api.v1.Machine.
 * Use `create(MachineSchema)` to create a new message.
 */
export declare const MachineSchema: GenMessage<Machine>;

/**
 * @generated from message baepo.api.v1.MachineSpec
 */
export declare type MachineSpec = Message<"baepo.api.v1.MachineSpec"> & {
  /**
   * @generated from field: uint32 v_cpus = 1;
   */
  vCpus: number;

  /**
   * @generated from field: uint64 memory_mb = 2;
   */
  memoryMb: bigint;

  /**
   * @generated from field: map<string, string> env = 3;
   */
  env: { [key: string]: string };

  /**
   * @generated from field: string image = 4;
   */
  image: string;

  /**
   * @generated from field: repeated string command = 5;
   */
  command: string[];
};

/**
 * Describes the message baepo.api.v1.MachineSpec.
 * Use `create(MachineSpecSchema)` to create a new message.
 */
export declare const MachineSpecSchema: GenMessage<MachineSpec>;

/**
 * @generated from message baepo.api.v1.MachineListRequest
 */
export declare type MachineListRequest = Message<"baepo.api.v1.MachineListRequest"> & {
  /**
   * @generated from field: string workspace_id = 1;
   */
  workspaceId: string;
};

/**
 * Describes the message baepo.api.v1.MachineListRequest.
 * Use `create(MachineListRequestSchema)` to create a new message.
 */
export declare const MachineListRequestSchema: GenMessage<MachineListRequest>;

/**
 * @generated from message baepo.api.v1.MachineListResponse
 */
export declare type MachineListResponse = Message<"baepo.api.v1.MachineListResponse"> & {
  /**
   * @generated from field: repeated baepo.api.v1.Machine machines = 1;
   */
  machines: Machine[];
};

/**
 * Describes the message baepo.api.v1.MachineListResponse.
 * Use `create(MachineListResponseSchema)` to create a new message.
 */
export declare const MachineListResponseSchema: GenMessage<MachineListResponse>;

/**
 * @generated from message baepo.api.v1.MachineFindByIdRequest
 */
export declare type MachineFindByIdRequest = Message<"baepo.api.v1.MachineFindByIdRequest"> & {
  /**
   * @generated from field: string workspace_id = 1;
   */
  workspaceId: string;

  /**
   * @generated from field: string machine_id = 2;
   */
  machineId: string;
};

/**
 * Describes the message baepo.api.v1.MachineFindByIdRequest.
 * Use `create(MachineFindByIdRequestSchema)` to create a new message.
 */
export declare const MachineFindByIdRequestSchema: GenMessage<MachineFindByIdRequest>;

/**
 * @generated from message baepo.api.v1.MachineFindByIdResponse
 */
export declare type MachineFindByIdResponse = Message<"baepo.api.v1.MachineFindByIdResponse"> & {
  /**
   * @generated from field: baepo.api.v1.Machine machine = 1;
   */
  machine?: Machine;
};

/**
 * Describes the message baepo.api.v1.MachineFindByIdResponse.
 * Use `create(MachineFindByIdResponseSchema)` to create a new message.
 */
export declare const MachineFindByIdResponseSchema: GenMessage<MachineFindByIdResponse>;

/**
 * @generated from message baepo.api.v1.MachineCreateRequest
 */
export declare type MachineCreateRequest = Message<"baepo.api.v1.MachineCreateRequest"> & {
  /**
   * @generated from field: string workspace_id = 1;
   */
  workspaceId: string;

  /**
   * @generated from field: optional string name = 2;
   */
  name?: string;

  /**
   * @generated from field: optional uint32 timeout = 3;
   */
  timeout?: number;

  /**
   * @generated from field: baepo.api.v1.MachineSpec spec = 4;
   */
  spec?: MachineSpec;

  /**
   * @generated from field: map<string, string> metadata = 5;
   */
  metadata: { [key: string]: string };
};

/**
 * Describes the message baepo.api.v1.MachineCreateRequest.
 * Use `create(MachineCreateRequestSchema)` to create a new message.
 */
export declare const MachineCreateRequestSchema: GenMessage<MachineCreateRequest>;

/**
 * @generated from message baepo.api.v1.MachineCreateResponse
 */
export declare type MachineCreateResponse = Message<"baepo.api.v1.MachineCreateResponse"> & {
  /**
   * @generated from field: baepo.api.v1.Machine machine = 1;
   */
  machine?: Machine;
};

/**
 * Describes the message baepo.api.v1.MachineCreateResponse.
 * Use `create(MachineCreateResponseSchema)` to create a new message.
 */
export declare const MachineCreateResponseSchema: GenMessage<MachineCreateResponse>;

/**
 * @generated from message baepo.api.v1.MachineTerminateRequest
 */
export declare type MachineTerminateRequest = Message<"baepo.api.v1.MachineTerminateRequest"> & {
  /**
   * @generated from field: string workspace_id = 1;
   */
  workspaceId: string;

  /**
   * @generated from field: string machine_id = 2;
   */
  machineId: string;
};

/**
 * Describes the message baepo.api.v1.MachineTerminateRequest.
 * Use `create(MachineTerminateRequestSchema)` to create a new message.
 */
export declare const MachineTerminateRequestSchema: GenMessage<MachineTerminateRequest>;

/**
 * @generated from message baepo.api.v1.MachineTerminateResponse
 */
export declare type MachineTerminateResponse = Message<"baepo.api.v1.MachineTerminateResponse"> & {
  /**
   * @generated from field: baepo.api.v1.Machine machine = 1;
   */
  machine?: Machine;
};

/**
 * Describes the message baepo.api.v1.MachineTerminateResponse.
 * Use `create(MachineTerminateResponseSchema)` to create a new message.
 */
export declare const MachineTerminateResponseSchema: GenMessage<MachineTerminateResponse>;

/**
 * @generated from enum baepo.api.v1.MachineStatus
 */
export enum MachineStatus {
  /**
   * @generated from enum value: MachineStatus_Unknown = 0;
   */
  MachineStatus_Unknown = 0,

  /**
   * @generated from enum value: MachineStatus_Scheduling = 1;
   */
  MachineStatus_Scheduling = 1,

  /**
   * @generated from enum value: MachineStatus_Starting = 2;
   */
  MachineStatus_Starting = 2,

  /**
   * @generated from enum value: MachineStatus_Running = 3;
   */
  MachineStatus_Running = 3,

  /**
   * @generated from enum value: MachineStatus_Terminating = 4;
   */
  MachineStatus_Terminating = 4,

  /**
   * @generated from enum value: MachineStatus_Terminated = 5;
   */
  MachineStatus_Terminated = 5,
}

/**
 * Describes the enum baepo.api.v1.MachineStatus.
 */
export declare const MachineStatusSchema: GenEnum<MachineStatus>;

/**
 * @generated from service baepo.api.v1.MachineService
 */
export declare const MachineService: GenService<{
  /**
   * @generated from rpc baepo.api.v1.MachineService.List
   */
  list: {
    methodKind: "unary";
    input: typeof MachineListRequestSchema;
    output: typeof MachineListResponseSchema;
  },
  /**
   * @generated from rpc baepo.api.v1.MachineService.FindById
   */
  findById: {
    methodKind: "unary";
    input: typeof MachineFindByIdRequestSchema;
    output: typeof MachineFindByIdResponseSchema;
  },
  /**
   * @generated from rpc baepo.api.v1.MachineService.Create
   */
  create: {
    methodKind: "unary";
    input: typeof MachineCreateRequestSchema;
    output: typeof MachineCreateResponseSchema;
  },
  /**
   * @generated from rpc baepo.api.v1.MachineService.Terminate
   */
  terminate: {
    methodKind: "unary";
    input: typeof MachineTerminateRequestSchema;
    output: typeof MachineTerminateResponseSchema;
  },
}>;

