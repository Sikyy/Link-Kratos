// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             v3.21.12
// source: link/v1/link.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationLinkInterceptor = "/link.v1.Link/Interceptor"
const OperationLinkLogin = "/link.v1.Link/Login"
const OperationLinkRealTimeTraffic = "/link.v1.Link/RealTimeTraffic"
const OperationLinkRegister = "/link.v1.Link/Register"
const OperationLinkStatistics = "/link.v1.Link/Statistics"

type LinkHTTPServer interface {
	// Interceptor拦截流量信息，并返回
	Interceptor(context.Context, *InterceptorRequest) (*InterceptorReply, error)
	// Login登陆
	Login(context.Context, *LoginRequest) (*UserReply, error)
	// RealTimeTraffic获取实时流量信息，并返回
	RealTimeTraffic(context.Context, *RealTimeTrafficRequest) (*RealTimeTrafficReply, error)
	// Register注册
	Register(context.Context, *RegisterRequest) (*UserReply, error)
	// Statistics统计流量信息，并返回
	Statistics(context.Context, *StatisticsRequest) (*StatisticsReply, error)
}

func RegisterLinkHTTPServer(s *http.Server, srv LinkHTTPServer) {
	r := s.Route("/")
	r.POST("/api/users", _Link_Register0_HTTP_Handler(srv))
	r.POST("/api/users/login", _Link_Login0_HTTP_Handler(srv))
	r.GET("/api/users/interceptor", _Link_Interceptor0_HTTP_Handler(srv))
	r.GET("/api/users/statistics", _Link_Statistics0_HTTP_Handler(srv))
	r.GET("/api/users/realtimetraffic", _Link_RealTimeTraffic0_HTTP_Handler(srv))
}

func _Link_Register0_HTTP_Handler(srv LinkHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLinkRegister)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserReply)
		return ctx.Result(200, reply)
	}
}

func _Link_Login0_HTTP_Handler(srv LinkHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLinkLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserReply)
		return ctx.Result(200, reply)
	}
}

func _Link_Interceptor0_HTTP_Handler(srv LinkHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in InterceptorRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLinkInterceptor)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Interceptor(ctx, req.(*InterceptorRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*InterceptorReply)
		return ctx.Result(200, reply)
	}
}

func _Link_Statistics0_HTTP_Handler(srv LinkHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in StatisticsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLinkStatistics)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Statistics(ctx, req.(*StatisticsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*StatisticsReply)
		return ctx.Result(200, reply)
	}
}

func _Link_RealTimeTraffic0_HTTP_Handler(srv LinkHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RealTimeTrafficRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLinkRealTimeTraffic)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RealTimeTraffic(ctx, req.(*RealTimeTrafficRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RealTimeTrafficReply)
		return ctx.Result(200, reply)
	}
}

type LinkHTTPClient interface {
	Interceptor(ctx context.Context, req *InterceptorRequest, opts ...http.CallOption) (rsp *InterceptorReply, err error)
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *UserReply, err error)
	RealTimeTraffic(ctx context.Context, req *RealTimeTrafficRequest, opts ...http.CallOption) (rsp *RealTimeTrafficReply, err error)
	Register(ctx context.Context, req *RegisterRequest, opts ...http.CallOption) (rsp *UserReply, err error)
	Statistics(ctx context.Context, req *StatisticsRequest, opts ...http.CallOption) (rsp *StatisticsReply, err error)
}

type LinkHTTPClientImpl struct {
	cc *http.Client
}

func NewLinkHTTPClient(client *http.Client) LinkHTTPClient {
	return &LinkHTTPClientImpl{client}
}

func (c *LinkHTTPClientImpl) Interceptor(ctx context.Context, in *InterceptorRequest, opts ...http.CallOption) (*InterceptorReply, error) {
	var out InterceptorReply
	pattern := "/api/users/interceptor"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationLinkInterceptor))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LinkHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*UserReply, error) {
	var out UserReply
	pattern := "/api/users/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationLinkLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LinkHTTPClientImpl) RealTimeTraffic(ctx context.Context, in *RealTimeTrafficRequest, opts ...http.CallOption) (*RealTimeTrafficReply, error) {
	var out RealTimeTrafficReply
	pattern := "/api/users/realtimetraffic"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationLinkRealTimeTraffic))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LinkHTTPClientImpl) Register(ctx context.Context, in *RegisterRequest, opts ...http.CallOption) (*UserReply, error) {
	var out UserReply
	pattern := "/api/users"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationLinkRegister))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LinkHTTPClientImpl) Statistics(ctx context.Context, in *StatisticsRequest, opts ...http.CallOption) (*StatisticsReply, error) {
	var out StatisticsReply
	pattern := "/api/users/statistics"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationLinkStatistics))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
