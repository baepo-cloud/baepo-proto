// @generated by protoc-gen-es v2.2.5
// @generated from file baepo/api/v1/machine.proto (package baepo.api.v1, syntax proto3)
/* eslint-disable */

import { enumDesc, fileDesc, messageDesc, serviceDesc, tsEnum } from "@bufbuild/protobuf/codegenv1";
import { file_google_protobuf_empty, file_google_protobuf_timestamp } from "@bufbuild/protobuf/wkt";

/**
 * Describes the file baepo/api/v1/machine.proto.
 */
export const file_baepo_api_v1_machine = /*@__PURE__*/
  fileDesc("ChpiYWVwby9hcGkvdjEvbWFjaGluZS5wcm90bxIMYmFlcG8uYXBpLnYxIusFCgdNYWNoaW5lEgoKAmlkGAEgASgJEhEKBG5hbWUYAiABKAlIAIgBARIrCgZzdGF0dXMYAyABKA4yGy5iYWVwby5hcGkudjEuTWFjaGluZVN0YXR1cxIUCgdub2RlX2lkGAQgASgJSAGIAQESJwoEc3BlYxgFIAEoCzIZLmJhZXBvLmFwaS52MS5NYWNoaW5lU3BlYxIUCgd0aW1lb3V0GAYgASgNSAKIAQESMwoKc3RhcnRlZF9hdBgHIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXBIA4gBARIzCgpleHBpcmVzX2F0GAggASgLMhouZ29vZ2xlLnByb3RvYnVmLlRpbWVzdGFtcEgEiAEBEjYKDXRlcm1pbmF0ZWRfYXQYCSABKAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wSAWIAQESHgoRdGVybWluYXRpb25fY2F1c2UYCiABKAlIBogBARIgChN0ZXJtaW5hdGlvbl9kZXRhaWxzGAsgASgJSAeIAQESNQoIbWV0YWRhdGEYDCADKAsyIy5iYWVwby5hcGkudjEuTWFjaGluZS5NZXRhZGF0YUVudHJ5Ei4KCmNyZWF0ZWRfYXQYDSABKAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wEi4KCnVwZGF0ZWRfYXQYDiABKAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wEhQKDHdvcmtzcGFjZV9pZBgPIAEoCRovCg1NZXRhZGF0YUVudHJ5EgsKA2tleRgBIAEoCRINCgV2YWx1ZRgCIAEoCToCOAFCBwoFX25hbWVCCgoIX25vZGVfaWRCCgoIX3RpbWVvdXRCDQoLX3N0YXJ0ZWRfYXRCDQoLX2V4cGlyZXNfYXRCEAoOX3Rlcm1pbmF0ZWRfYXRCFAoSX3Rlcm1pbmF0aW9uX2NhdXNlQhYKFF90ZXJtaW5hdGlvbl9kZXRhaWxzIq0BCgtNYWNoaW5lU3BlYxIOCgZ2X2NwdXMYASABKA0SEQoJbWVtb3J5X21iGAIgASgEEi8KA2VudhgDIAMoCzIiLmJhZXBvLmFwaS52MS5NYWNoaW5lU3BlYy5FbnZFbnRyeRINCgVpbWFnZRgEIAEoCRIPCgdjb21tYW5kGAUgAygJGioKCEVudkVudHJ5EgsKA2tleRgBIAEoCRINCgV2YWx1ZRgCIAEoCToCOAEiKgoSTWFjaGluZUxpc3RSZXF1ZXN0EhQKDHdvcmtzcGFjZV9pZBgBIAEoCSI+ChNNYWNoaW5lTGlzdFJlc3BvbnNlEicKCG1hY2hpbmVzGAEgAygLMhUuYmFlcG8uYXBpLnYxLk1hY2hpbmUiLAoWTWFjaGluZUZpbmRCeUlkUmVxdWVzdBISCgptYWNoaW5lX2lkGAEgASgJIkEKF01hY2hpbmVGaW5kQnlJZFJlc3BvbnNlEiYKB21hY2hpbmUYASABKAsyFS5iYWVwby5hcGkudjEuTWFjaGluZSKIAgoUTWFjaGluZUNyZWF0ZVJlcXVlc3QSFAoMd29ya3NwYWNlX2lkGAEgASgJEhEKBG5hbWUYAiABKAlIAIgBARIUCgd0aW1lb3V0GAMgASgNSAGIAQESJwoEc3BlYxgEIAEoCzIZLmJhZXBvLmFwaS52MS5NYWNoaW5lU3BlYxJCCghtZXRhZGF0YRgFIAMoCzIwLmJhZXBvLmFwaS52MS5NYWNoaW5lQ3JlYXRlUmVxdWVzdC5NZXRhZGF0YUVudHJ5Gi8KDU1ldGFkYXRhRW50cnkSCwoDa2V5GAEgASgJEg0KBXZhbHVlGAIgASgJOgI4AUIHCgVfbmFtZUIKCghfdGltZW91dCI/ChVNYWNoaW5lQ3JlYXRlUmVzcG9uc2USJgoHbWFjaGluZRgBIAEoCzIVLmJhZXBvLmFwaS52MS5NYWNoaW5lIi0KF01hY2hpbmVUZXJtaW5hdGVSZXF1ZXN0EhIKCm1hY2hpbmVfaWQYASABKAkiQgoYTWFjaGluZVRlcm1pbmF0ZVJlc3BvbnNlEiYKB21hY2hpbmUYASABKAsyFS5iYWVwby5hcGkudjEuTWFjaGluZSq8AQoNTWFjaGluZVN0YXR1cxIZChVNYWNoaW5lU3RhdHVzX1Vua25vd24QABIcChhNYWNoaW5lU3RhdHVzX1NjaGVkdWxpbmcQARIaChZNYWNoaW5lU3RhdHVzX1N0YXJ0aW5nEAISGQoVTWFjaGluZVN0YXR1c19SdW5uaW5nEAMSHQoZTWFjaGluZVN0YXR1c19UZXJtaW5hdGluZxAEEhwKGE1hY2hpbmVTdGF0dXNfVGVybWluYXRlZBAFMuUCCg5NYWNoaW5lU2VydmljZRJLCgRMaXN0EiAuYmFlcG8uYXBpLnYxLk1hY2hpbmVMaXN0UmVxdWVzdBohLmJhZXBvLmFwaS52MS5NYWNoaW5lTGlzdFJlc3BvbnNlElcKCEZpbmRCeUlkEiQuYmFlcG8uYXBpLnYxLk1hY2hpbmVGaW5kQnlJZFJlcXVlc3QaJS5iYWVwby5hcGkudjEuTWFjaGluZUZpbmRCeUlkUmVzcG9uc2USUQoGQ3JlYXRlEiIuYmFlcG8uYXBpLnYxLk1hY2hpbmVDcmVhdGVSZXF1ZXN0GiMuYmFlcG8uYXBpLnYxLk1hY2hpbmVDcmVhdGVSZXNwb25zZRJaCglUZXJtaW5hdGUSJS5iYWVwby5hcGkudjEuTWFjaGluZVRlcm1pbmF0ZVJlcXVlc3QaJi5iYWVwby5hcGkudjEuTWFjaGluZVRlcm1pbmF0ZVJlc3BvbnNlQjRaMmdpdGh1Yi5jb20vYmFlcG8tY2xvdWQvYmFlcG8tcHJvdG8vZ28vYmFlcG8vYXBpL3YxYgZwcm90bzM", [file_google_protobuf_empty, file_google_protobuf_timestamp]);

/**
 * Describes the message baepo.api.v1.Machine.
 * Use `create(MachineSchema)` to create a new message.
 */
export const MachineSchema = /*@__PURE__*/
  messageDesc(file_baepo_api_v1_machine, 0);

/**
 * Describes the message baepo.api.v1.MachineSpec.
 * Use `create(MachineSpecSchema)` to create a new message.
 */
export const MachineSpecSchema = /*@__PURE__*/
  messageDesc(file_baepo_api_v1_machine, 1);

/**
 * Describes the message baepo.api.v1.MachineListRequest.
 * Use `create(MachineListRequestSchema)` to create a new message.
 */
export const MachineListRequestSchema = /*@__PURE__*/
  messageDesc(file_baepo_api_v1_machine, 2);

/**
 * Describes the message baepo.api.v1.MachineListResponse.
 * Use `create(MachineListResponseSchema)` to create a new message.
 */
export const MachineListResponseSchema = /*@__PURE__*/
  messageDesc(file_baepo_api_v1_machine, 3);

/**
 * Describes the message baepo.api.v1.MachineFindByIdRequest.
 * Use `create(MachineFindByIdRequestSchema)` to create a new message.
 */
export const MachineFindByIdRequestSchema = /*@__PURE__*/
  messageDesc(file_baepo_api_v1_machine, 4);

/**
 * Describes the message baepo.api.v1.MachineFindByIdResponse.
 * Use `create(MachineFindByIdResponseSchema)` to create a new message.
 */
export const MachineFindByIdResponseSchema = /*@__PURE__*/
  messageDesc(file_baepo_api_v1_machine, 5);

/**
 * Describes the message baepo.api.v1.MachineCreateRequest.
 * Use `create(MachineCreateRequestSchema)` to create a new message.
 */
export const MachineCreateRequestSchema = /*@__PURE__*/
  messageDesc(file_baepo_api_v1_machine, 6);

/**
 * Describes the message baepo.api.v1.MachineCreateResponse.
 * Use `create(MachineCreateResponseSchema)` to create a new message.
 */
export const MachineCreateResponseSchema = /*@__PURE__*/
  messageDesc(file_baepo_api_v1_machine, 7);

/**
 * Describes the message baepo.api.v1.MachineTerminateRequest.
 * Use `create(MachineTerminateRequestSchema)` to create a new message.
 */
export const MachineTerminateRequestSchema = /*@__PURE__*/
  messageDesc(file_baepo_api_v1_machine, 8);

/**
 * Describes the message baepo.api.v1.MachineTerminateResponse.
 * Use `create(MachineTerminateResponseSchema)` to create a new message.
 */
export const MachineTerminateResponseSchema = /*@__PURE__*/
  messageDesc(file_baepo_api_v1_machine, 9);

/**
 * Describes the enum baepo.api.v1.MachineStatus.
 */
export const MachineStatusSchema = /*@__PURE__*/
  enumDesc(file_baepo_api_v1_machine, 0);

/**
 * @generated from enum baepo.api.v1.MachineStatus
 */
export const MachineStatus = /*@__PURE__*/
  tsEnum(MachineStatusSchema);

/**
 * @generated from service baepo.api.v1.MachineService
 */
export const MachineService = /*@__PURE__*/
  serviceDesc(file_baepo_api_v1_machine, 0);

