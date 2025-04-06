// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: baepo/api/v1/machine.proto

package v1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/baepo-cloud/baepo-proto/go/baepo/api/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// MachineServiceName is the fully-qualified name of the MachineService service.
	MachineServiceName = "baepo.api.v1.MachineService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// MachineServiceListProcedure is the fully-qualified name of the MachineService's List RPC.
	MachineServiceListProcedure = "/baepo.api.v1.MachineService/List"
	// MachineServiceCreateProcedure is the fully-qualified name of the MachineService's Create RPC.
	MachineServiceCreateProcedure = "/baepo.api.v1.MachineService/Create"
	// MachineServiceTerminateProcedure is the fully-qualified name of the MachineService's Terminate
	// RPC.
	MachineServiceTerminateProcedure = "/baepo.api.v1.MachineService/Terminate"
)

// MachineServiceClient is a client for the baepo.api.v1.MachineService service.
type MachineServiceClient interface {
	List(context.Context, *connect.Request[v1.MachineListRequest]) (*connect.Response[v1.MachineListResponse], error)
	Create(context.Context, *connect.Request[v1.MachineCreateRequest]) (*connect.Response[v1.MachineCreateResponse], error)
	Terminate(context.Context, *connect.Request[v1.MachineTerminateRequest]) (*connect.Response[v1.MachineTerminateResponse], error)
}

// NewMachineServiceClient constructs a client for the baepo.api.v1.MachineService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewMachineServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) MachineServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	machineServiceMethods := v1.File_baepo_api_v1_machine_proto.Services().ByName("MachineService").Methods()
	return &machineServiceClient{
		list: connect.NewClient[v1.MachineListRequest, v1.MachineListResponse](
			httpClient,
			baseURL+MachineServiceListProcedure,
			connect.WithSchema(machineServiceMethods.ByName("List")),
			connect.WithClientOptions(opts...),
		),
		create: connect.NewClient[v1.MachineCreateRequest, v1.MachineCreateResponse](
			httpClient,
			baseURL+MachineServiceCreateProcedure,
			connect.WithSchema(machineServiceMethods.ByName("Create")),
			connect.WithClientOptions(opts...),
		),
		terminate: connect.NewClient[v1.MachineTerminateRequest, v1.MachineTerminateResponse](
			httpClient,
			baseURL+MachineServiceTerminateProcedure,
			connect.WithSchema(machineServiceMethods.ByName("Terminate")),
			connect.WithClientOptions(opts...),
		),
	}
}

// machineServiceClient implements MachineServiceClient.
type machineServiceClient struct {
	list      *connect.Client[v1.MachineListRequest, v1.MachineListResponse]
	create    *connect.Client[v1.MachineCreateRequest, v1.MachineCreateResponse]
	terminate *connect.Client[v1.MachineTerminateRequest, v1.MachineTerminateResponse]
}

// List calls baepo.api.v1.MachineService.List.
func (c *machineServiceClient) List(ctx context.Context, req *connect.Request[v1.MachineListRequest]) (*connect.Response[v1.MachineListResponse], error) {
	return c.list.CallUnary(ctx, req)
}

// Create calls baepo.api.v1.MachineService.Create.
func (c *machineServiceClient) Create(ctx context.Context, req *connect.Request[v1.MachineCreateRequest]) (*connect.Response[v1.MachineCreateResponse], error) {
	return c.create.CallUnary(ctx, req)
}

// Terminate calls baepo.api.v1.MachineService.Terminate.
func (c *machineServiceClient) Terminate(ctx context.Context, req *connect.Request[v1.MachineTerminateRequest]) (*connect.Response[v1.MachineTerminateResponse], error) {
	return c.terminate.CallUnary(ctx, req)
}

// MachineServiceHandler is an implementation of the baepo.api.v1.MachineService service.
type MachineServiceHandler interface {
	List(context.Context, *connect.Request[v1.MachineListRequest]) (*connect.Response[v1.MachineListResponse], error)
	Create(context.Context, *connect.Request[v1.MachineCreateRequest]) (*connect.Response[v1.MachineCreateResponse], error)
	Terminate(context.Context, *connect.Request[v1.MachineTerminateRequest]) (*connect.Response[v1.MachineTerminateResponse], error)
}

// NewMachineServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewMachineServiceHandler(svc MachineServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	machineServiceMethods := v1.File_baepo_api_v1_machine_proto.Services().ByName("MachineService").Methods()
	machineServiceListHandler := connect.NewUnaryHandler(
		MachineServiceListProcedure,
		svc.List,
		connect.WithSchema(machineServiceMethods.ByName("List")),
		connect.WithHandlerOptions(opts...),
	)
	machineServiceCreateHandler := connect.NewUnaryHandler(
		MachineServiceCreateProcedure,
		svc.Create,
		connect.WithSchema(machineServiceMethods.ByName("Create")),
		connect.WithHandlerOptions(opts...),
	)
	machineServiceTerminateHandler := connect.NewUnaryHandler(
		MachineServiceTerminateProcedure,
		svc.Terminate,
		connect.WithSchema(machineServiceMethods.ByName("Terminate")),
		connect.WithHandlerOptions(opts...),
	)
	return "/baepo.api.v1.MachineService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case MachineServiceListProcedure:
			machineServiceListHandler.ServeHTTP(w, r)
		case MachineServiceCreateProcedure:
			machineServiceCreateHandler.ServeHTTP(w, r)
		case MachineServiceTerminateProcedure:
			machineServiceTerminateHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedMachineServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedMachineServiceHandler struct{}

func (UnimplementedMachineServiceHandler) List(context.Context, *connect.Request[v1.MachineListRequest]) (*connect.Response[v1.MachineListResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("baepo.api.v1.MachineService.List is not implemented"))
}

func (UnimplementedMachineServiceHandler) Create(context.Context, *connect.Request[v1.MachineCreateRequest]) (*connect.Response[v1.MachineCreateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("baepo.api.v1.MachineService.Create is not implemented"))
}

func (UnimplementedMachineServiceHandler) Terminate(context.Context, *connect.Request[v1.MachineTerminateRequest]) (*connect.Response[v1.MachineTerminateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("baepo.api.v1.MachineService.Terminate is not implemented"))
}
