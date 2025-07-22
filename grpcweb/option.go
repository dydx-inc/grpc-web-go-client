package grpcweb

import (
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/encoding"
	"google.golang.org/grpc/encoding/proto"
	"google.golang.org/grpc/metadata"

	"github.com/ktr0731/grpc-web-go-client/grpcweb/transport"
)

var (
	defaultDialOptions = dialOptions{}
	defaultCallOptions = callOptions{
		codec: encoding.GetCodec(proto.Name),
	}
)

type dialOptions struct {
	defaultCallOptions   []CallOption
	insecure             bool
	transportCredentials credentials.TransportCredentials
	transport            transport.UnaryTransportFactory
}

type DialOption func(*dialOptions)

func WithDefaultCallOptions(opts ...CallOption) DialOption {
	return func(opt *dialOptions) {
		opt.defaultCallOptions = opts
	}
}

func WithInsecure() DialOption {
	return func(opt *dialOptions) {
		opt.insecure = true
	}
}

func WithTransport(t transport.UnaryTransportFactory) DialOption {
	return func(opt *dialOptions) {
		opt.transport = t
	}
}

func WithTransportCredentials(creds credentials.TransportCredentials) DialOption {
	return func(opt *dialOptions) {
		opt.transportCredentials = creds
	}
}

type callOptions struct {
	codec           encoding.Codec
	header, trailer *metadata.MD
}

type CallOption func(*callOptions)

func CallContentSubtype(contentSubtype string) CallOption {
	return func(opt *callOptions) {
		opt.codec = encoding.GetCodec(contentSubtype)
	}
}

func Header(h *metadata.MD) CallOption {
	return func(opt *callOptions) {
		*h = metadata.New(nil)
		opt.header = h
	}
}

func Trailer(t *metadata.MD) CallOption {
	return func(opt *callOptions) {
		*t = metadata.New(nil)
		opt.trailer = t
	}
}
