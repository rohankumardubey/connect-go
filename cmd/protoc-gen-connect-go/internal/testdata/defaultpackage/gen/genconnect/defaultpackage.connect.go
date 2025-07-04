// Copyright 2021-2025 The Connect Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: defaultpackage.proto

package genconnect

import (
	connect "connectrpc.com/connect"
	gen "connectrpc.com/connect/cmd/protoc-gen-connect-go/internal/testdata/defaultpackage/gen"
	context "context"
	errors "errors"
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
	// TestServiceName is the fully-qualified name of the TestService service.
	TestServiceName = "connect.test.default_package.TestService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// TestServiceMethodProcedure is the fully-qualified name of the TestService's Method RPC.
	TestServiceMethodProcedure = "/connect.test.default_package.TestService/Method"
)

// TestServiceClient is a client for the connect.test.default_package.TestService service.
type TestServiceClient interface {
	Method(context.Context, *connect.Request[gen.Request]) (*connect.Response[gen.Response], error)
}

// NewTestServiceClient constructs a client for the connect.test.default_package.TestService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTestServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) TestServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	testServiceMethods := gen.File_defaultpackage_proto.Services().ByName("TestService").Methods()
	return &testServiceClient{
		method: connect.NewClient[gen.Request, gen.Response](
			httpClient,
			baseURL+TestServiceMethodProcedure,
			connect.WithSchema(testServiceMethods.ByName("Method")),
			connect.WithClientOptions(opts...),
		),
	}
}

// testServiceClient implements TestServiceClient.
type testServiceClient struct {
	method *connect.Client[gen.Request, gen.Response]
}

// Method calls connect.test.default_package.TestService.Method.
func (c *testServiceClient) Method(ctx context.Context, req *connect.Request[gen.Request]) (*connect.Response[gen.Response], error) {
	return c.method.CallUnary(ctx, req)
}

// TestServiceHandler is an implementation of the connect.test.default_package.TestService service.
type TestServiceHandler interface {
	Method(context.Context, *connect.Request[gen.Request]) (*connect.Response[gen.Response], error)
}

// NewTestServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTestServiceHandler(svc TestServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	testServiceMethods := gen.File_defaultpackage_proto.Services().ByName("TestService").Methods()
	testServiceMethodHandler := connect.NewUnaryHandler(
		TestServiceMethodProcedure,
		svc.Method,
		connect.WithSchema(testServiceMethods.ByName("Method")),
		connect.WithHandlerOptions(opts...),
	)
	return "/connect.test.default_package.TestService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case TestServiceMethodProcedure:
			testServiceMethodHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedTestServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTestServiceHandler struct{}

func (UnimplementedTestServiceHandler) Method(context.Context, *connect.Request[gen.Request]) (*connect.Response[gen.Response], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("connect.test.default_package.TestService.Method is not implemented"))
}
