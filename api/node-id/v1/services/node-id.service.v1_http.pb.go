// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.3.1

package servicev1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	api "github.com/ikaiguang/go-snowflake-node-id/api"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationSrvWorkerExtendNodeId = "/node.id.api.servicev1.SrvWorker/ExtendNodeId"
const OperationSrvWorkerGetNodeId = "/node.id.api.servicev1.SrvWorker/GetNodeId"

type SrvWorkerHTTPServer interface {
	ExtendNodeId(context.Context, *api.ExtendNodeIdReq) (*api.Result, error)
	GetNodeId(context.Context, *api.GetNodeIdReq) (*api.SnowflakeWorkerNode, error)
}

func RegisterSrvWorkerHTTPServer(s *http.Server, srv SrvWorkerHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/node-id/get-node-id", _SrvWorker_GetNodeId0_HTTP_Handler(srv))
	r.POST("/api/v1/node-id/extend-node-id", _SrvWorker_ExtendNodeId0_HTTP_Handler(srv))
}

func _SrvWorker_GetNodeId0_HTTP_Handler(srv SrvWorkerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in api.GetNodeIdReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSrvWorkerGetNodeId)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetNodeId(ctx, req.(*api.GetNodeIdReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*api.SnowflakeWorkerNode)
		return ctx.Result(200, reply)
	}
}

func _SrvWorker_ExtendNodeId0_HTTP_Handler(srv SrvWorkerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in api.ExtendNodeIdReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSrvWorkerExtendNodeId)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ExtendNodeId(ctx, req.(*api.ExtendNodeIdReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*api.Result)
		return ctx.Result(200, reply)
	}
}

type SrvWorkerHTTPClient interface {
	ExtendNodeId(ctx context.Context, req *api.ExtendNodeIdReq, opts ...http.CallOption) (rsp *api.Result, err error)
	GetNodeId(ctx context.Context, req *api.GetNodeIdReq, opts ...http.CallOption) (rsp *api.SnowflakeWorkerNode, err error)
}

type SrvWorkerHTTPClientImpl struct {
	cc *http.Client
}

func NewSrvWorkerHTTPClient(client *http.Client) SrvWorkerHTTPClient {
	return &SrvWorkerHTTPClientImpl{client}
}

func (c *SrvWorkerHTTPClientImpl) ExtendNodeId(ctx context.Context, in *api.ExtendNodeIdReq, opts ...http.CallOption) (*api.Result, error) {
	var out api.Result
	pattern := "/api/v1/node-id/extend-node-id"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSrvWorkerExtendNodeId))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SrvWorkerHTTPClientImpl) GetNodeId(ctx context.Context, in *api.GetNodeIdReq, opts ...http.CallOption) (*api.SnowflakeWorkerNode, error) {
	var out api.SnowflakeWorkerNode
	pattern := "/api/v1/node-id/get-node-id"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSrvWorkerGetNodeId))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}