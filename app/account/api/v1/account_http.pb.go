// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.0
// - protoc             v3.17.3
// source: app/account/api/v1/account.proto

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

const OperationAccountCreateAccount = "/app.account.api.v1.Account/CreateAccount"

type AccountHTTPServer interface {
	CreateAccount(context.Context, *CreateAccountRequest) (*CommonReply, error)
}

func RegisterAccountHTTPServer(s *http.Server, srv AccountHTTPServer) {
	r := s.Route("/")
	r.POST("/account/v1/create", _Account_CreateAccount0_HTTP_Handler(srv))
}

func _Account_CreateAccount0_HTTP_Handler(srv AccountHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateAccountRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAccountCreateAccount)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateAccount(ctx, req.(*CreateAccountRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CommonReply)
		return ctx.Result(200, reply)
	}
}

type AccountHTTPClient interface {
	CreateAccount(ctx context.Context, req *CreateAccountRequest, opts ...http.CallOption) (rsp *CommonReply, err error)
}

type AccountHTTPClientImpl struct {
	cc *http.Client
}

func NewAccountHTTPClient(client *http.Client) AccountHTTPClient {
	return &AccountHTTPClientImpl{client}
}

func (c *AccountHTTPClientImpl) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...http.CallOption) (*CommonReply, error) {
	var out CommonReply
	pattern := "/account/v1/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAccountCreateAccount))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
