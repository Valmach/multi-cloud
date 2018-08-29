// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: s3.proto

/*
Package s3 is a generated protocol buffer package.

It is generated from these files:
	s3.proto

It has these top-level messages:
	GetObjectRequest
	GetObjectResponse
*/
package s3

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for S3 service

type S3Service interface {
	GetObject(ctx context.Context, in *GetObjectRequest, opts ...client.CallOption) (*GetObjectResponse, error)
}

type s3Service struct {
	c    client.Client
	name string
}

func NewS3Service(name string, c client.Client) S3Service {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "s3"
	}
	return &s3Service{
		c:    c,
		name: name,
	}
}

func (c *s3Service) GetObject(ctx context.Context, in *GetObjectRequest, opts ...client.CallOption) (*GetObjectResponse, error) {
	req := c.c.NewRequest(c.name, "S3.GetObject", in)
	out := new(GetObjectResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for S3 service

type S3Handler interface {
	GetObject(context.Context, *GetObjectRequest, *GetObjectResponse) error
}

func RegisterS3Handler(s server.Server, hdlr S3Handler, opts ...server.HandlerOption) error {
	type s3 interface {
		GetObject(ctx context.Context, in *GetObjectRequest, out *GetObjectResponse) error
	}
	type S3 struct {
		s3
	}
	h := &s3Handler{hdlr}
	return s.Handle(s.NewHandler(&S3{h}, opts...))
}

type s3Handler struct {
	S3Handler
}

func (h *s3Handler) GetObject(ctx context.Context, in *GetObjectRequest, out *GetObjectResponse) error {
	return h.S3Handler.GetObject(ctx, in, out)
}
