// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shortener.proto

/*
Package go_micro_shortener is a generated protocol buffer package.

It is generated from these files:
	shortener.proto

It has these top-level messages:
	UrlRequest
	UrlResponse
	ShortUrlResponse
	TextRequest
	TextResponse
*/
package go_micro_shortener

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type UrlRequest struct {
	Url string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
}

func (m *UrlRequest) Reset()                    { *m = UrlRequest{} }
func (m *UrlRequest) String() string            { return proto.CompactTextString(m) }
func (*UrlRequest) ProtoMessage()               {}
func (*UrlRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UrlRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type UrlResponse struct {
	OperationResponse string `protobuf:"bytes,2,opt,name=operationResponse" json:"operationResponse,omitempty"`
}

func (m *UrlResponse) Reset()                    { *m = UrlResponse{} }
func (m *UrlResponse) String() string            { return proto.CompactTextString(m) }
func (*UrlResponse) ProtoMessage()               {}
func (*UrlResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UrlResponse) GetOperationResponse() string {
	if m != nil {
		return m.OperationResponse
	}
	return ""
}

type ShortUrlResponse struct {
	ShortUrl string `protobuf:"bytes,3,opt,name=shortUrl" json:"shortUrl,omitempty"`
}

func (m *ShortUrlResponse) Reset()                    { *m = ShortUrlResponse{} }
func (m *ShortUrlResponse) String() string            { return proto.CompactTextString(m) }
func (*ShortUrlResponse) ProtoMessage()               {}
func (*ShortUrlResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ShortUrlResponse) GetShortUrl() string {
	if m != nil {
		return m.ShortUrl
	}
	return ""
}

type TextRequest struct {
	Text string `protobuf:"bytes,4,opt,name=text" json:"text,omitempty"`
}

func (m *TextRequest) Reset()                    { *m = TextRequest{} }
func (m *TextRequest) String() string            { return proto.CompactTextString(m) }
func (*TextRequest) ProtoMessage()               {}
func (*TextRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TextRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type TextResponse struct {
	TextWithShort string `protobuf:"bytes,5,opt,name=textWithShort" json:"textWithShort,omitempty"`
}

func (m *TextResponse) Reset()                    { *m = TextResponse{} }
func (m *TextResponse) String() string            { return proto.CompactTextString(m) }
func (*TextResponse) ProtoMessage()               {}
func (*TextResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *TextResponse) GetTextWithShort() string {
	if m != nil {
		return m.TextWithShort
	}
	return ""
}

func init() {
	proto.RegisterType((*UrlRequest)(nil), "go.micro.shortener.UrlRequest")
	proto.RegisterType((*UrlResponse)(nil), "go.micro.shortener.UrlResponse")
	proto.RegisterType((*ShortUrlResponse)(nil), "go.micro.shortener.ShortUrlResponse")
	proto.RegisterType((*TextRequest)(nil), "go.micro.shortener.TextRequest")
	proto.RegisterType((*TextResponse)(nil), "go.micro.shortener.TextResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Shortener service

type ShortenerClient interface {
	AddShort(ctx context.Context, in *UrlRequest, opts ...client.CallOption) (*UrlResponse, error)
	GetShort(ctx context.Context, in *UrlRequest, opts ...client.CallOption) (*ShortUrlResponse, error)
	ReplaceAll(ctx context.Context, in *TextRequest, opts ...client.CallOption) (*TextResponse, error)
}

type shortenerClient struct {
	c           client.Client
	serviceName string
}

func NewShortenerClient(serviceName string, c client.Client) ShortenerClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.shortener"
	}
	return &shortenerClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *shortenerClient) AddShort(ctx context.Context, in *UrlRequest, opts ...client.CallOption) (*UrlResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Shortener.AddShort", in)
	out := new(UrlResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenerClient) GetShort(ctx context.Context, in *UrlRequest, opts ...client.CallOption) (*ShortUrlResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Shortener.GetShort", in)
	out := new(ShortUrlResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenerClient) ReplaceAll(ctx context.Context, in *TextRequest, opts ...client.CallOption) (*TextResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Shortener.ReplaceAll", in)
	out := new(TextResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Shortener service

type ShortenerHandler interface {
	AddShort(context.Context, *UrlRequest, *UrlResponse) error
	GetShort(context.Context, *UrlRequest, *ShortUrlResponse) error
	ReplaceAll(context.Context, *TextRequest, *TextResponse) error
}

func RegisterShortenerHandler(s server.Server, hdlr ShortenerHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Shortener{hdlr}, opts...))
}

type Shortener struct {
	ShortenerHandler
}

func (h *Shortener) AddShort(ctx context.Context, in *UrlRequest, out *UrlResponse) error {
	return h.ShortenerHandler.AddShort(ctx, in, out)
}

func (h *Shortener) GetShort(ctx context.Context, in *UrlRequest, out *ShortUrlResponse) error {
	return h.ShortenerHandler.GetShort(ctx, in, out)
}

func (h *Shortener) ReplaceAll(ctx context.Context, in *TextRequest, out *TextResponse) error {
	return h.ShortenerHandler.ReplaceAll(ctx, in, out)
}

func init() { proto.RegisterFile("shortener.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x3d, 0x4f, 0x02, 0x41,
	0x10, 0xe5, 0x4b, 0x73, 0x0c, 0x1a, 0x71, 0xaa, 0xcd, 0x15, 0x88, 0x1b, 0x0a, 0x0b, 0xb3, 0x85,
	0xda, 0x59, 0x51, 0x59, 0x59, 0x78, 0x68, 0xac, 0x11, 0x26, 0x72, 0xc9, 0x7a, 0x7b, 0xee, 0x0e,
	0x09, 0xad, 0xff, 0xdc, 0xdc, 0x78, 0x8b, 0x28, 0x90, 0xd0, 0xcd, 0xbc, 0xf7, 0xf6, 0xcd, 0xcc,
	0xcb, 0xc2, 0x59, 0x58, 0x38, 0xcf, 0x54, 0x90, 0x37, 0xa5, 0x77, 0xec, 0x10, 0xdf, 0x9d, 0xf9,
	0xc8, 0x67, 0xde, 0x99, 0x35, 0xa3, 0x07, 0x00, 0x2f, 0xde, 0x66, 0xf4, 0xb9, 0xa4, 0xc0, 0xd8,
	0x87, 0xf6, 0xd2, 0x5b, 0xd5, 0x1c, 0x36, 0xaf, 0xba, 0x59, 0x55, 0xea, 0x7b, 0xe8, 0x09, 0x1f,
	0x4a, 0x57, 0x04, 0xc2, 0x6b, 0x38, 0x77, 0x25, 0xf9, 0x29, 0xe7, 0xae, 0x88, 0xa0, 0x6a, 0x89,
	0x7c, 0x9b, 0xd0, 0x06, 0xfa, 0x93, 0x6a, 0xd2, 0xa6, 0x43, 0x0a, 0x49, 0xa8, 0x31, 0xd5, 0x96,
	0x87, 0xeb, 0x5e, 0x5f, 0x42, 0xef, 0x99, 0x56, 0x1c, 0xb7, 0x41, 0xe8, 0x30, 0xad, 0x58, 0x75,
	0x44, 0x26, 0xb5, 0xbe, 0x83, 0x93, 0x1f, 0x49, 0x6d, 0x37, 0x82, 0xd3, 0x0a, 0x7f, 0xcd, 0x79,
	0x21, 0xa3, 0xd4, 0x91, 0x88, 0xff, 0x82, 0x37, 0x5f, 0x2d, 0xe8, 0x4e, 0xe2, 0xcd, 0xf8, 0x08,
	0xc9, 0x78, 0x3e, 0x97, 0x1e, 0x07, 0x66, 0x3b, 0x14, 0xf3, 0x9b, 0x48, 0x7a, 0xb1, 0x97, 0xaf,
	0x6f, 0x6c, 0x60, 0x06, 0xc9, 0x03, 0xf1, 0x61, 0x76, 0xa3, 0x5d, 0xfc, 0xff, 0x8c, 0x74, 0x03,
	0x9f, 0x00, 0x32, 0x2a, 0xed, 0x74, 0x46, 0x63, 0x6b, 0x71, 0xe7, 0x12, 0x1b, 0x49, 0xa5, 0xc3,
	0xfd, 0x82, 0x68, 0xf9, 0x76, 0x2c, 0x9f, 0xe0, 0xf6, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x98, 0xaa,
	0x40, 0xd0, 0x17, 0x02, 0x00, 0x00,
}
