// @generated by protoc-gen-es v2.2.5 with parameter "import_extension=js"
// @generated from file baepo/core/v1/machine.proto (package baepo.core.v1, syntax proto3)
/* eslint-disable */

import { enumDesc, fileDesc, messageDesc, tsEnum } from "@bufbuild/protobuf/codegenv1";
import { file_google_protobuf_empty, file_google_protobuf_timestamp } from "@bufbuild/protobuf/wkt";

/**
 * Describes the file baepo/core/v1/machine.proto.
 */
export const file_baepo_core_v1_machine = /*@__PURE__*/
  fileDesc("ChtiYWVwby9jb3JlL3YxL21hY2hpbmUucHJvdG8SDWJhZXBvLmNvcmUudjEiiQEKC01hY2hpbmVTcGVjEgwKBGNwdXMYASABKA0SEQoJbWVtb3J5X21iGAIgASgEEjcKCmNvbnRhaW5lcnMYAyADKAsyIy5iYWVwby5jb3JlLnYxLk1hY2hpbmVDb250YWluZXJTcGVjEhQKB3RpbWVvdXQYBCABKARIAIgBAUIKCghfdGltZW91dCLXAgoUTWFjaGluZUNvbnRhaW5lclNwZWMSDAoEbmFtZRgBIAEoCRINCgVpbWFnZRgCIAEoCRI5CgNlbnYYAyADKAsyLC5iYWVwby5jb3JlLnYxLk1hY2hpbmVDb250YWluZXJTcGVjLkVudkVudHJ5Eg8KB2NvbW1hbmQYBCADKAkSQwoLaGVhbHRoY2hlY2sYBSABKAsyLi5iYWVwby5jb3JlLnYxLk1hY2hpbmVDb250YWluZXJIZWFsdGhjaGVja1NwZWMSGAoLd29ya2luZ19kaXIYBiABKAlIAIgBARI7CgdyZXN0YXJ0GAcgASgLMiouYmFlcG8uY29yZS52MS5NYWNoaW5lQ29udGFpbmVyUmVzdGFydFNwZWMaKgoIRW52RW50cnkSCwoDa2V5GAEgASgJEg0KBXZhbHVlGAIgASgJOgI4AUIOCgxfd29ya2luZ19kaXIiywEKG01hY2hpbmVDb250YWluZXJSZXN0YXJ0U3BlYxJBCgZwb2xpY3kYASABKA4yMS5iYWVwby5jb3JlLnYxLk1hY2hpbmVDb250YWluZXJSZXN0YXJ0U3BlYy5Qb2xpY3kSEwoLbWF4X3JldHJpZXMYAiABKAUiVAoGUG9saWN5EhIKDlBvbGljeV9Vbmtub3duEAASDQoJUG9saWN5X05vEAESFAoQUG9saWN5X09uRmFpbHVyZRACEhEKDVBvbGljeV9BbHdheXMQAyKKAwofTWFjaGluZUNvbnRhaW5lckhlYWx0aGNoZWNrU3BlYxIdChVpbml0aWFsX2RlbGF5X3NlY29uZHMYASABKAUSFgoOcGVyaW9kX3NlY29uZHMYAiABKAUSUgoEaHR0cBgDIAEoCzJCLmJhZXBvLmNvcmUudjEuTWFjaGluZUNvbnRhaW5lckhlYWx0aGNoZWNrU3BlYy5IdHRwSGVhbHRoY2hlY2tTcGVjSAAa0wEKE0h0dHBIZWFsdGhjaGVja1NwZWMSDgoGbWV0aG9kGAEgASgJEgwKBHBhdGgYAiABKAkSDAoEcG9ydBgDIAEoBRJgCgdoZWFkZXJzGAQgAygLMk8uYmFlcG8uY29yZS52MS5NYWNoaW5lQ29udGFpbmVySGVhbHRoY2hlY2tTcGVjLkh0dHBIZWFsdGhjaGVja1NwZWMuSGVhZGVyc0VudHJ5Gi4KDEhlYWRlcnNFbnRyeRILCgNrZXkYASABKAkSDQoFdmFsdWUYAiABKAk6AjgBQgYKBHR5cGUi+QwKDE1hY2hpbmVFdmVudBIQCghldmVudF9pZBgBIAEoCRISCgptYWNoaW5lX2lkGAIgASgJEi0KCXRpbWVzdGFtcBgDIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXASRgoNc3RhdGVfY2hhbmdlZBgEIAEoCzItLmJhZXBvLmNvcmUudjEuTWFjaGluZUV2ZW50LlN0YXRlQ2hhbmdlZEV2ZW50SAASOwoHc3RhcnRlZBgFIAEoCzIoLmJhZXBvLmNvcmUudjEuTWFjaGluZUV2ZW50LlN0YXJ0ZWRFdmVudEgAEkEKCnRlcm1pbmF0ZWQYBiABKAsyKy5iYWVwby5jb3JlLnYxLk1hY2hpbmVFdmVudC5UZXJtaW5hdGVkRXZlbnRIABJVChVkZXNpcmVkX3N0YXRlX2NoYW5nZWQYByABKAsyNC5iYWVwby5jb3JlLnYxLk1hY2hpbmVFdmVudC5EZXNpcmVkU3RhdGVDaGFuZ2VkRXZlbnRIABJYChZyZWNvbmNpbGlhdGlvbl9zdGFydGVkGAggASgLMjYuYmFlcG8uY29yZS52MS5NYWNoaW5lRXZlbnQuUmVjb25jaWxpYXRpb25TdGFydGVkRXZlbnRIABJcChhyZWNvbmNpbGlhdGlvbl9jb21wbGV0ZWQYCSABKAsyOC5iYWVwby5jb3JlLnYxLk1hY2hpbmVFdmVudC5SZWNvbmNpbGlhdGlvbkNvbXBsZXRlZEV2ZW50SAASUwoRaGVhbHRoY2hlY2tfZXZlbnQYCiABKAsyNi5iYWVwby5jb3JlLnYxLk1hY2hpbmVFdmVudC5Db250YWluZXJTdGF0ZUNoYW5nZWRFdmVudEgAGj8KEVN0YXRlQ2hhbmdlZEV2ZW50EioKBXN0YXRlGAEgASgOMhsuYmFlcG8uY29yZS52MS5NYWNoaW5lU3RhdGUaUgoMU3RhcnRlZEV2ZW50EjMKCmV4cGlyZXNfYXQYASABKAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wSACIAQFCDQoLX2V4cGlyZXNfYXQaggEKD1Rlcm1pbmF0ZWRFdmVudBI1CgVjYXVzZRgBIAEoDjImLmJhZXBvLmNvcmUudjEuTWFjaGluZVRlcm1pbmF0aW9uQ2F1c2USIAoTdGVybWluYXRpb25fZGV0YWlscxgCIAEoCUgAiAEBQhYKFF90ZXJtaW5hdGlvbl9kZXRhaWxzGlUKGERlc2lyZWRTdGF0ZUNoYW5nZWRFdmVudBI5Cg1kZXNpcmVkX3N0YXRlGAEgASgOMiIuYmFlcG8uY29yZS52MS5NYWNoaW5lRGVzaXJlZFN0YXRlGlcKGlJlY29uY2lsaWF0aW9uU3RhcnRlZEV2ZW50EjkKDWRlc2lyZWRfc3RhdGUYASABKA4yIi5iYWVwby5jb3JlLnYxLk1hY2hpbmVEZXNpcmVkU3RhdGUadwocUmVjb25jaWxpYXRpb25Db21wbGV0ZWRFdmVudBI5Cg1kZXNpcmVkX3N0YXRlGAEgASgOMiIuYmFlcG8uY29yZS52MS5NYWNoaW5lRGVzaXJlZFN0YXRlEhIKBWVycm9yGAIgASgJSACIAQFCCAoGX2Vycm9yGpsDChpDb250YWluZXJTdGF0ZUNoYW5nZWRFdmVudBIWCg5jb250YWluZXJfbmFtZRgBIAEoCRIzCgVzdGF0ZRgCIAEoDjIkLmJhZXBvLmNvcmUudjEuTWFjaGluZUNvbnRhaW5lclN0YXRlEjMKCnN0YXJ0ZWRfYXQYAyABKAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wSACIAQESMgoJZXhpdGVkX2F0GAQgASgLMhouZ29vZ2xlLnByb3RvYnVmLlRpbWVzdGFtcEgBiAEBEhYKCWV4aXRfY29kZRgFIAEoBUgCiAEBEhcKCmV4aXRfZXJyb3IYBiABKAlIA4gBARIPCgdoZWFsdGh5GAcgASgIEh4KEWhlYWx0aGNoZWNrX2Vycm9yGAggASgJSASIAQESFQoNcmVzdGFydF9jb3VudBgJIAEoBUINCgtfc3RhcnRlZF9hdEIMCgpfZXhpdGVkX2F0QgwKCl9leGl0X2NvZGVCDQoLX2V4aXRfZXJyb3JCFAoSX2hlYWx0aGNoZWNrX2Vycm9yQgcKBWV2ZW50KuUBCgxNYWNoaW5lU3RhdGUSGAoUTWFjaGluZVN0YXRlX1Vua25vd24QABIYChRNYWNoaW5lU3RhdGVfUGVuZGluZxABEhkKFU1hY2hpbmVTdGF0ZV9TdGFydGluZxACEhgKFE1hY2hpbmVTdGF0ZV9SdW5uaW5nEAMSGQoVTWFjaGluZVN0YXRlX0RlZ3JhZGVkEAQSFgoSTWFjaGluZVN0YXRlX0Vycm9yEAUSHAoYTWFjaGluZVN0YXRlX1Rlcm1pbmF0aW5nEAYSGwoXTWFjaGluZVN0YXRlX1Rlcm1pbmF0ZWQQByqcAQoTTWFjaGluZURlc2lyZWRTdGF0ZRIfChtNYWNoaW5lRGVzaXJlZFN0YXRlX1Vua25vd24QABIfChtNYWNoaW5lRGVzaXJlZFN0YXRlX1BlbmRpbmcQARIfChtNYWNoaW5lRGVzaXJlZFN0YXRlX1J1bm5pbmcQAhIiCh5NYWNoaW5lRGVzaXJlZFN0YXRlX1Rlcm1pbmF0ZWQQAyqZAgoXTWFjaGluZVRlcm1pbmF0aW9uQ2F1c2USIwofTWFjaGluZVRlcm1pbmF0aW9uQ2F1c2VfVW5rbm93bhAAEi0KKU1hY2hpbmVUZXJtaW5hdGlvbkNhdXNlX0hlYWx0aGNoZWNrRmFpbGVkEAESLQopTWFjaGluZVRlcm1pbmF0aW9uQ2F1c2VfTWFudWFsbHlSZXF1ZXN0ZWQQAhIpCiVNYWNoaW5lVGVybWluYXRpb25DYXVzZV9JbnRlcm5hbEVycm9yEAMSKwonTWFjaGluZVRlcm1pbmF0aW9uQ2F1c2VfTm9Ob2RlQXZhaWxhYmxlEAQSIwofTWFjaGluZVRlcm1pbmF0aW9uQ2F1c2VfRXhwaXJlZBAFKn8KFU1hY2hpbmVDb250YWluZXJTdGF0ZRIhCh1NYWNoaW5lQ29udGFpbmVyU3RhdGVfVW5rbm93bhAAEiEKHU1hY2hpbmVDb250YWluZXJTdGF0ZV9SdW5uaW5nEAESIAocTWFjaGluZUNvbnRhaW5lclN0YXRlX0V4aXRlZBACQj5aPGdpdGh1Yi5jb20vYmFlcG8tY2xvdWQvYmFlcG8tcHJvdG8vZ28vYmFlcG8vY29yZS92MTtjb3JldjFwYmIGcHJvdG8z", [file_google_protobuf_empty, file_google_protobuf_timestamp]);

/**
 * Describes the message baepo.core.v1.MachineSpec.
 * Use `create(MachineSpecSchema)` to create a new message.
 */
export const MachineSpecSchema = /*@__PURE__*/
  messageDesc(file_baepo_core_v1_machine, 0);

/**
 * Describes the message baepo.core.v1.MachineContainerSpec.
 * Use `create(MachineContainerSpecSchema)` to create a new message.
 */
export const MachineContainerSpecSchema = /*@__PURE__*/
  messageDesc(file_baepo_core_v1_machine, 1);

/**
 * Describes the message baepo.core.v1.MachineContainerRestartSpec.
 * Use `create(MachineContainerRestartSpecSchema)` to create a new message.
 */
export const MachineContainerRestartSpecSchema = /*@__PURE__*/
  messageDesc(file_baepo_core_v1_machine, 2);

/**
 * Describes the enum baepo.core.v1.MachineContainerRestartSpec.Policy.
 */
export const MachineContainerRestartSpec_PolicySchema = /*@__PURE__*/
  enumDesc(file_baepo_core_v1_machine, 2, 0);

/**
 * @generated from enum baepo.core.v1.MachineContainerRestartSpec.Policy
 */
export const MachineContainerRestartSpec_Policy = /*@__PURE__*/
  tsEnum(MachineContainerRestartSpec_PolicySchema);

/**
 * Describes the message baepo.core.v1.MachineContainerHealthcheckSpec.
 * Use `create(MachineContainerHealthcheckSpecSchema)` to create a new message.
 */
export const MachineContainerHealthcheckSpecSchema = /*@__PURE__*/
  messageDesc(file_baepo_core_v1_machine, 3);

/**
 * Describes the message baepo.core.v1.MachineContainerHealthcheckSpec.HttpHealthcheckSpec.
 * Use `create(MachineContainerHealthcheckSpec_HttpHealthcheckSpecSchema)` to create a new message.
 */
export const MachineContainerHealthcheckSpec_HttpHealthcheckSpecSchema = /*@__PURE__*/
  messageDesc(file_baepo_core_v1_machine, 3, 0);

/**
 * Describes the message baepo.core.v1.MachineEvent.
 * Use `create(MachineEventSchema)` to create a new message.
 */
export const MachineEventSchema = /*@__PURE__*/
  messageDesc(file_baepo_core_v1_machine, 4);

/**
 * Describes the message baepo.core.v1.MachineEvent.StateChangedEvent.
 * Use `create(MachineEvent_StateChangedEventSchema)` to create a new message.
 */
export const MachineEvent_StateChangedEventSchema = /*@__PURE__*/
  messageDesc(file_baepo_core_v1_machine, 4, 0);

/**
 * Describes the message baepo.core.v1.MachineEvent.StartedEvent.
 * Use `create(MachineEvent_StartedEventSchema)` to create a new message.
 */
export const MachineEvent_StartedEventSchema = /*@__PURE__*/
  messageDesc(file_baepo_core_v1_machine, 4, 1);

/**
 * Describes the message baepo.core.v1.MachineEvent.TerminatedEvent.
 * Use `create(MachineEvent_TerminatedEventSchema)` to create a new message.
 */
export const MachineEvent_TerminatedEventSchema = /*@__PURE__*/
  messageDesc(file_baepo_core_v1_machine, 4, 2);

/**
 * Describes the message baepo.core.v1.MachineEvent.DesiredStateChangedEvent.
 * Use `create(MachineEvent_DesiredStateChangedEventSchema)` to create a new message.
 */
export const MachineEvent_DesiredStateChangedEventSchema = /*@__PURE__*/
  messageDesc(file_baepo_core_v1_machine, 4, 3);

/**
 * Describes the message baepo.core.v1.MachineEvent.ReconciliationStartedEvent.
 * Use `create(MachineEvent_ReconciliationStartedEventSchema)` to create a new message.
 */
export const MachineEvent_ReconciliationStartedEventSchema = /*@__PURE__*/
  messageDesc(file_baepo_core_v1_machine, 4, 4);

/**
 * Describes the message baepo.core.v1.MachineEvent.ReconciliationCompletedEvent.
 * Use `create(MachineEvent_ReconciliationCompletedEventSchema)` to create a new message.
 */
export const MachineEvent_ReconciliationCompletedEventSchema = /*@__PURE__*/
  messageDesc(file_baepo_core_v1_machine, 4, 5);

/**
 * Describes the message baepo.core.v1.MachineEvent.ContainerStateChangedEvent.
 * Use `create(MachineEvent_ContainerStateChangedEventSchema)` to create a new message.
 */
export const MachineEvent_ContainerStateChangedEventSchema = /*@__PURE__*/
  messageDesc(file_baepo_core_v1_machine, 4, 6);

/**
 * Describes the enum baepo.core.v1.MachineState.
 */
export const MachineStateSchema = /*@__PURE__*/
  enumDesc(file_baepo_core_v1_machine, 0);

/**
 * @generated from enum baepo.core.v1.MachineState
 */
export const MachineState = /*@__PURE__*/
  tsEnum(MachineStateSchema);

/**
 * Describes the enum baepo.core.v1.MachineDesiredState.
 */
export const MachineDesiredStateSchema = /*@__PURE__*/
  enumDesc(file_baepo_core_v1_machine, 1);

/**
 * @generated from enum baepo.core.v1.MachineDesiredState
 */
export const MachineDesiredState = /*@__PURE__*/
  tsEnum(MachineDesiredStateSchema);

/**
 * Describes the enum baepo.core.v1.MachineTerminationCause.
 */
export const MachineTerminationCauseSchema = /*@__PURE__*/
  enumDesc(file_baepo_core_v1_machine, 2);

/**
 * @generated from enum baepo.core.v1.MachineTerminationCause
 */
export const MachineTerminationCause = /*@__PURE__*/
  tsEnum(MachineTerminationCauseSchema);

/**
 * Describes the enum baepo.core.v1.MachineContainerState.
 */
export const MachineContainerStateSchema = /*@__PURE__*/
  enumDesc(file_baepo_core_v1_machine, 3);

/**
 * @generated from enum baepo.core.v1.MachineContainerState
 */
export const MachineContainerState = /*@__PURE__*/
  tsEnum(MachineContainerStateSchema);

