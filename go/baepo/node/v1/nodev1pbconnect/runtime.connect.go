// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: baepo/node/v1/runtime.proto

package nodev1pbconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/baepo-cloud/baepo-proto/go/baepo/node/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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
	// RuntimeName is the fully-qualified name of the Runtime service.
	RuntimeName = "baepo.node.v1.Runtime"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// RuntimeGetStateProcedure is the fully-qualified name of the Runtime's GetState RPC.
	RuntimeGetStateProcedure = "/baepo.node.v1.Runtime/GetState"
	// RuntimeGetLogsProcedure is the fully-qualified name of the Runtime's GetLogs RPC.
	RuntimeGetLogsProcedure = "/baepo.node.v1.Runtime/GetLogs"
	// RuntimeGetContainerLogsProcedure is the fully-qualified name of the Runtime's GetContainerLogs
	// RPC.
	RuntimeGetContainerLogsProcedure = "/baepo.node.v1.Runtime/GetContainerLogs"
	// RuntimeEventsProcedure is the fully-qualified name of the Runtime's Events RPC.
	RuntimeEventsProcedure = "/baepo.node.v1.Runtime/Events"
)

// RuntimeClient is a client for the baepo.node.v1.Runtime service.
type RuntimeClient interface {
	GetState(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[v1.RuntimeGetStateResponse], error)
	GetLogs(context.Context, *connect.Request[v1.RuntimeGetLogsRequest]) (*connect.ServerStreamForClient[v1.RuntimeGetLogsResponse], error)
	GetContainerLogs(context.Context, *connect.Request[v1.RuntimeGetContainerLogsRequest]) (*connect.ServerStreamForClient[v1.RuntimeGetContainerLogsResponse], error)
	Events(context.Context, *connect.Request[emptypb.Empty]) (*connect.ServerStreamForClient[v1.RuntimeEventsResponse], error)
}

// NewRuntimeClient constructs a client for the baepo.node.v1.Runtime service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewRuntimeClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) RuntimeClient {
	baseURL = strings.TrimRight(baseURL, "/")
	runtimeMethods := v1.File_baepo_node_v1_runtime_proto.Services().ByName("Runtime").Methods()
	return &runtimeClient{
		getState: connect.NewClient[emptypb.Empty, v1.RuntimeGetStateResponse](
			httpClient,
			baseURL+RuntimeGetStateProcedure,
			connect.WithSchema(runtimeMethods.ByName("GetState")),
			connect.WithClientOptions(opts...),
		),
		getLogs: connect.NewClient[v1.RuntimeGetLogsRequest, v1.RuntimeGetLogsResponse](
			httpClient,
			baseURL+RuntimeGetLogsProcedure,
			connect.WithSchema(runtimeMethods.ByName("GetLogs")),
			connect.WithClientOptions(opts...),
		),
		getContainerLogs: connect.NewClient[v1.RuntimeGetContainerLogsRequest, v1.RuntimeGetContainerLogsResponse](
			httpClient,
			baseURL+RuntimeGetContainerLogsProcedure,
			connect.WithSchema(runtimeMethods.ByName("GetContainerLogs")),
			connect.WithClientOptions(opts...),
		),
		events: connect.NewClient[emptypb.Empty, v1.RuntimeEventsResponse](
			httpClient,
			baseURL+RuntimeEventsProcedure,
			connect.WithSchema(runtimeMethods.ByName("Events")),
			connect.WithClientOptions(opts...),
		),
	}
}

// runtimeClient implements RuntimeClient.
type runtimeClient struct {
	getState         *connect.Client[emptypb.Empty, v1.RuntimeGetStateResponse]
	getLogs          *connect.Client[v1.RuntimeGetLogsRequest, v1.RuntimeGetLogsResponse]
	getContainerLogs *connect.Client[v1.RuntimeGetContainerLogsRequest, v1.RuntimeGetContainerLogsResponse]
	events           *connect.Client[emptypb.Empty, v1.RuntimeEventsResponse]
}

// GetState calls baepo.node.v1.Runtime.GetState.
func (c *runtimeClient) GetState(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[v1.RuntimeGetStateResponse], error) {
	return c.getState.CallUnary(ctx, req)
}

// GetLogs calls baepo.node.v1.Runtime.GetLogs.
func (c *runtimeClient) GetLogs(ctx context.Context, req *connect.Request[v1.RuntimeGetLogsRequest]) (*connect.ServerStreamForClient[v1.RuntimeGetLogsResponse], error) {
	return c.getLogs.CallServerStream(ctx, req)
}

// GetContainerLogs calls baepo.node.v1.Runtime.GetContainerLogs.
func (c *runtimeClient) GetContainerLogs(ctx context.Context, req *connect.Request[v1.RuntimeGetContainerLogsRequest]) (*connect.ServerStreamForClient[v1.RuntimeGetContainerLogsResponse], error) {
	return c.getContainerLogs.CallServerStream(ctx, req)
}

// Events calls baepo.node.v1.Runtime.Events.
func (c *runtimeClient) Events(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.ServerStreamForClient[v1.RuntimeEventsResponse], error) {
	return c.events.CallServerStream(ctx, req)
}

// RuntimeHandler is an implementation of the baepo.node.v1.Runtime service.
type RuntimeHandler interface {
	GetState(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[v1.RuntimeGetStateResponse], error)
	GetLogs(context.Context, *connect.Request[v1.RuntimeGetLogsRequest], *connect.ServerStream[v1.RuntimeGetLogsResponse]) error
	GetContainerLogs(context.Context, *connect.Request[v1.RuntimeGetContainerLogsRequest], *connect.ServerStream[v1.RuntimeGetContainerLogsResponse]) error
	Events(context.Context, *connect.Request[emptypb.Empty], *connect.ServerStream[v1.RuntimeEventsResponse]) error
}

// NewRuntimeHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewRuntimeHandler(svc RuntimeHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	runtimeMethods := v1.File_baepo_node_v1_runtime_proto.Services().ByName("Runtime").Methods()
	runtimeGetStateHandler := connect.NewUnaryHandler(
		RuntimeGetStateProcedure,
		svc.GetState,
		connect.WithSchema(runtimeMethods.ByName("GetState")),
		connect.WithHandlerOptions(opts...),
	)
	runtimeGetLogsHandler := connect.NewServerStreamHandler(
		RuntimeGetLogsProcedure,
		svc.GetLogs,
		connect.WithSchema(runtimeMethods.ByName("GetLogs")),
		connect.WithHandlerOptions(opts...),
	)
	runtimeGetContainerLogsHandler := connect.NewServerStreamHandler(
		RuntimeGetContainerLogsProcedure,
		svc.GetContainerLogs,
		connect.WithSchema(runtimeMethods.ByName("GetContainerLogs")),
		connect.WithHandlerOptions(opts...),
	)
	runtimeEventsHandler := connect.NewServerStreamHandler(
		RuntimeEventsProcedure,
		svc.Events,
		connect.WithSchema(runtimeMethods.ByName("Events")),
		connect.WithHandlerOptions(opts...),
	)
	return "/baepo.node.v1.Runtime/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case RuntimeGetStateProcedure:
			runtimeGetStateHandler.ServeHTTP(w, r)
		case RuntimeGetLogsProcedure:
			runtimeGetLogsHandler.ServeHTTP(w, r)
		case RuntimeGetContainerLogsProcedure:
			runtimeGetContainerLogsHandler.ServeHTTP(w, r)
		case RuntimeEventsProcedure:
			runtimeEventsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedRuntimeHandler returns CodeUnimplemented from all methods.
type UnimplementedRuntimeHandler struct{}

func (UnimplementedRuntimeHandler) GetState(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[v1.RuntimeGetStateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("baepo.node.v1.Runtime.GetState is not implemented"))
}

func (UnimplementedRuntimeHandler) GetLogs(context.Context, *connect.Request[v1.RuntimeGetLogsRequest], *connect.ServerStream[v1.RuntimeGetLogsResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("baepo.node.v1.Runtime.GetLogs is not implemented"))
}

func (UnimplementedRuntimeHandler) GetContainerLogs(context.Context, *connect.Request[v1.RuntimeGetContainerLogsRequest], *connect.ServerStream[v1.RuntimeGetContainerLogsResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("baepo.node.v1.Runtime.GetContainerLogs is not implemented"))
}

func (UnimplementedRuntimeHandler) Events(context.Context, *connect.Request[emptypb.Empty], *connect.ServerStream[v1.RuntimeEventsResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("baepo.node.v1.Runtime.Events is not implemented"))
}
