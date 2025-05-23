// @generated by protoc-gen-es v2.2.5 with parameter "import_extension=js"
// @generated from file baepo/api/v1/node_controller.proto (package baepo.api.v1, syntax proto3)
/* eslint-disable */

import { fileDesc, messageDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import { file_google_protobuf_timestamp } from "@bufbuild/protobuf/wkt";
import { file_baepo_core_v1_machine } from "../../core/v1/machine_pb.js";
import { file_baepo_core_v1_container } from "../../core/v1/container_pb.js";

/**
 * Describes the file baepo/api/v1/node_controller.proto.
 */
export const file_baepo_api_v1_node_controller = /*@__PURE__*/
  fileDesc("CiJiYWVwby9hcGkvdjEvbm9kZV9jb250cm9sbGVyLnByb3RvEgxiYWVwby5hcGkudjEi8gcKGU5vZGVDb250cm9sbGVyU2VydmVyRXZlbnQSZAoWcmVnaXN0cmF0aW9uX2NvbXBsZXRlZBgBIAEoCzJCLmJhZXBvLmFwaS52MS5Ob2RlQ29udHJvbGxlclNlcnZlckV2ZW50LlJlZ2lzdHJhdGlvbkNvbXBsZXRlZEV2ZW50SAASQQoEcGluZxgCIAEoCzIxLmJhZXBvLmFwaS52MS5Ob2RlQ29udHJvbGxlclNlcnZlckV2ZW50LlBpbmdFdmVudEgAEkkKDmNyZWF0ZV9tYWNoaW5lGAMgASgLMi8uYmFlcG8uYXBpLnYxLk5vZGVDb250cm9sbGVyU2VydmVyRXZlbnQuTWFjaGluZUgAEm4KHHVwZGF0ZV9tYWNoaW5lX2Rlc2lyZWRfc3RhdGUYBCABKAsyRi5iYWVwby5hcGkudjEuTm9kZUNvbnRyb2xsZXJTZXJ2ZXJFdmVudC5VcGRhdGVNYWNoaW5lRGVzaXJlZFN0YXRlRXZlbnRIABrOAQoaUmVnaXN0cmF0aW9uQ29tcGxldGVkRXZlbnQSDwoHbm9kZV9pZBgBIAEoCRISCgpub2RlX3Rva2VuGAIgASgJEhYKDmF1dGhvcml0eV9jZXJ0GAMgASgMEhMKC3NlcnZlcl9jZXJ0GAQgASgMEhIKCnNlcnZlcl9rZXkYBSABKAwSSgoRZXhwZWN0ZWRfbWFjaGluZXMYBiADKAsyLy5iYWVwby5hcGkudjEuTm9kZUNvbnRyb2xsZXJTZXJ2ZXJFdmVudC5NYWNoaW5lGgsKCVBpbmdFdmVudBpvCh5VcGRhdGVNYWNoaW5lRGVzaXJlZFN0YXRlRXZlbnQSEgoKbWFjaGluZV9pZBgBIAEoCRI5Cg1kZXNpcmVkX3N0YXRlGAIgASgOMiIuYmFlcG8uY29yZS52MS5NYWNoaW5lRGVzaXJlZFN0YXRlGskBCgdNYWNoaW5lEhIKCm1hY2hpbmVfaWQYASABKAkSOQoNZGVzaXJlZF9zdGF0ZRgCIAEoDjIiLmJhZXBvLmNvcmUudjEuTWFjaGluZURlc2lyZWRTdGF0ZRIoCgRzcGVjGAMgASgLMhouYmFlcG8uY29yZS52MS5NYWNoaW5lU3BlYxJFCgpjb250YWluZXJzGAQgAygLMjEuYmFlcG8uYXBpLnYxLk5vZGVDb250cm9sbGVyU2VydmVyRXZlbnQuQ29udGFpbmVyGk0KCUNvbnRhaW5lchIUCgxjb250YWluZXJfaWQYASABKAkSKgoEc3BlYxgCIAEoCzIcLmJhZXBvLmNvcmUudjEuQ29udGFpbmVyU3BlY0IHCgVldmVudCLbBAoZTm9kZUNvbnRyb2xsZXJDbGllbnRFdmVudBJECghyZWdpc3RlchgBIAEoCzIwLmJhZXBvLmFwaS52MS5Ob2RlQ29udHJvbGxlckNsaWVudEV2ZW50LlJlZ2lzdGVySAASPgoFc3RhdHMYAiABKAsyLS5iYWVwby5hcGkudjEuTm9kZUNvbnRyb2xsZXJDbGllbnRFdmVudC5TdGF0c0gAEi4KB21hY2hpbmUYAyABKAsyGy5iYWVwby5jb3JlLnYxLk1hY2hpbmVFdmVudEgAEjIKCWNvbnRhaW5lchgEIAEoCzIdLmJhZXBvLmNvcmUudjEuQ29udGFpbmVyRXZlbnRIABrhAQoIUmVnaXN0ZXISEgoKY2x1c3Rlcl9pZBgBIAEoCRIXCg9ib290c3RyYXBfdG9rZW4YAiABKAkSFwoKbm9kZV90b2tlbhgDIAEoCUgAiAEBEhIKCmlwX2FkZHJlc3MYBCABKAkSFAoMYXBpX2VuZHBvaW50GAUgASgJEhgKEGdhdGV3YXlfZW5kcG9pbnQYBiABKAkSPAoFc3RhdHMYByABKAsyLS5iYWVwby5hcGkudjEuTm9kZUNvbnRyb2xsZXJDbGllbnRFdmVudC5TdGF0c0INCgtfbm9kZV90b2tlbhpnCgVTdGF0cxIXCg90b3RhbF9tZW1vcnlfbWIYASABKAQSFgoOdXNlZF9tZW1vcnlfbWIYAiABKAQSGgoScmVzZXJ2ZWRfbWVtb3J5X21iGAMgASgEEhEKCWNwdV9jb3VudBgEIAEoDUIHCgVldmVudDJ3ChVOb2RlQ29udHJvbGxlclNlcnZpY2USXgoGRXZlbnRzEicuYmFlcG8uYXBpLnYxLk5vZGVDb250cm9sbGVyQ2xpZW50RXZlbnQaJy5iYWVwby5hcGkudjEuTm9kZUNvbnRyb2xsZXJTZXJ2ZXJFdmVudCgBMAFCPFo6Z2l0aHViLmNvbS9iYWVwby1jbG91ZC9iYWVwby1wcm90by9nby9iYWVwby9hcGkvdjE7YXBpdjFwYmIGcHJvdG8z", [file_google_protobuf_timestamp, file_baepo_core_v1_machine, file_baepo_core_v1_container]);

/**
 * Describes the message baepo.api.v1.NodeControllerServerEvent.
 * Use `create(NodeControllerServerEventSchema)` to create a new message.
 */
export const NodeControllerServerEventSchema = /*@__PURE__*/
  messageDesc(file_baepo_api_v1_node_controller, 0);

/**
 * Describes the message baepo.api.v1.NodeControllerServerEvent.RegistrationCompletedEvent.
 * Use `create(NodeControllerServerEvent_RegistrationCompletedEventSchema)` to create a new message.
 */
export const NodeControllerServerEvent_RegistrationCompletedEventSchema = /*@__PURE__*/
  messageDesc(file_baepo_api_v1_node_controller, 0, 0);

/**
 * Describes the message baepo.api.v1.NodeControllerServerEvent.PingEvent.
 * Use `create(NodeControllerServerEvent_PingEventSchema)` to create a new message.
 */
export const NodeControllerServerEvent_PingEventSchema = /*@__PURE__*/
  messageDesc(file_baepo_api_v1_node_controller, 0, 1);

/**
 * Describes the message baepo.api.v1.NodeControllerServerEvent.UpdateMachineDesiredStateEvent.
 * Use `create(NodeControllerServerEvent_UpdateMachineDesiredStateEventSchema)` to create a new message.
 */
export const NodeControllerServerEvent_UpdateMachineDesiredStateEventSchema = /*@__PURE__*/
  messageDesc(file_baepo_api_v1_node_controller, 0, 2);

/**
 * Describes the message baepo.api.v1.NodeControllerServerEvent.Machine.
 * Use `create(NodeControllerServerEvent_MachineSchema)` to create a new message.
 */
export const NodeControllerServerEvent_MachineSchema = /*@__PURE__*/
  messageDesc(file_baepo_api_v1_node_controller, 0, 3);

/**
 * Describes the message baepo.api.v1.NodeControllerServerEvent.Container.
 * Use `create(NodeControllerServerEvent_ContainerSchema)` to create a new message.
 */
export const NodeControllerServerEvent_ContainerSchema = /*@__PURE__*/
  messageDesc(file_baepo_api_v1_node_controller, 0, 4);

/**
 * Describes the message baepo.api.v1.NodeControllerClientEvent.
 * Use `create(NodeControllerClientEventSchema)` to create a new message.
 */
export const NodeControllerClientEventSchema = /*@__PURE__*/
  messageDesc(file_baepo_api_v1_node_controller, 1);

/**
 * Describes the message baepo.api.v1.NodeControllerClientEvent.Register.
 * Use `create(NodeControllerClientEvent_RegisterSchema)` to create a new message.
 */
export const NodeControllerClientEvent_RegisterSchema = /*@__PURE__*/
  messageDesc(file_baepo_api_v1_node_controller, 1, 0);

/**
 * Describes the message baepo.api.v1.NodeControllerClientEvent.Stats.
 * Use `create(NodeControllerClientEvent_StatsSchema)` to create a new message.
 */
export const NodeControllerClientEvent_StatsSchema = /*@__PURE__*/
  messageDesc(file_baepo_api_v1_node_controller, 1, 1);

/**
 * @generated from service baepo.api.v1.NodeControllerService
 */
export const NodeControllerService = /*@__PURE__*/
  serviceDesc(file_baepo_api_v1_node_controller, 0);

