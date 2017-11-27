// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metadata.proto

package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"
import appscode_dtypes "github.com/appscode/api/dtypes"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ServerInfoResponse struct {
	Provider  string `protobuf:"bytes,1,opt,name=provider" json:"provider,omitempty"`
	ServerUrl string `protobuf:"bytes,2,opt,name=server_url,json=serverUrl" json:"server_url,omitempty"`
	CaCert    string `protobuf:"bytes,3,opt,name=ca_cert,json=caCert" json:"ca_cert,omitempty"`
}

func (m *ServerInfoResponse) Reset()                    { *m = ServerInfoResponse{} }
func (m *ServerInfoResponse) String() string            { return proto.CompactTextString(m) }
func (*ServerInfoResponse) ProtoMessage()               {}
func (*ServerInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *ServerInfoResponse) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *ServerInfoResponse) GetServerUrl() string {
	if m != nil {
		return m.ServerUrl
	}
	return ""
}

func (m *ServerInfoResponse) GetCaCert() string {
	if m != nil {
		return m.CaCert
	}
	return ""
}

func init() {
	proto.RegisterType((*ServerInfoResponse)(nil), "appscode.ci.v1beta1.ServerInfoResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Metadata service

type MetadataClient interface {
	ServerInfo(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*ServerInfoResponse, error)
}

type metadataClient struct {
	cc *grpc.ClientConn
}

func NewMetadataClient(cc *grpc.ClientConn) MetadataClient {
	return &metadataClient{cc}
}

func (c *metadataClient) ServerInfo(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*ServerInfoResponse, error) {
	out := new(ServerInfoResponse)
	err := grpc.Invoke(ctx, "/appscode.ci.v1beta1.Metadata/ServerInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Metadata service

type MetadataServer interface {
	ServerInfo(context.Context, *appscode_dtypes.VoidRequest) (*ServerInfoResponse, error)
}

func RegisterMetadataServer(s *grpc.Server, srv MetadataServer) {
	s.RegisterService(&_Metadata_serviceDesc, srv)
}

func _Metadata_ServerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(appscode_dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServer).ServerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.ci.v1beta1.Metadata/ServerInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServer).ServerInfo(ctx, req.(*appscode_dtypes.VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Metadata_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.ci.v1beta1.Metadata",
	HandlerType: (*MetadataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ServerInfo",
			Handler:    _Metadata_ServerInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metadata.proto",
}

func init() { proto.RegisterFile("metadata.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x4b, 0x4a, 0x03, 0x31,
	0x18, 0xc7, 0x49, 0x85, 0x3e, 0x02, 0xba, 0x88, 0x0b, 0x87, 0xb1, 0x3e, 0xe8, 0x46, 0x37, 0x26,
	0xd4, 0x22, 0xae, 0xad, 0x2b, 0x17, 0x42, 0xa9, 0xe8, 0xc2, 0x4d, 0x49, 0x33, 0x5f, 0x6b, 0x64,
	0x9a, 0x2f, 0x26, 0xe9, 0x80, 0xdb, 0x5e, 0x41, 0xf0, 0x02, 0x9e, 0xc1, 0x93, 0x78, 0x05, 0x0f,
	0x22, 0xce, 0xa3, 0xb5, 0xa8, 0x9b, 0x40, 0xf2, 0xcb, 0x9f, 0xff, 0x83, 0x6e, 0xcd, 0x20, 0xc8,
	0x44, 0x06, 0xc9, 0xad, 0xc3, 0x80, 0x6c, 0x5b, 0x5a, 0xeb, 0x15, 0x26, 0xc0, 0x95, 0xe6, 0x59,
	0x77, 0x0c, 0x41, 0x76, 0xe3, 0xf6, 0x14, 0x71, 0x9a, 0x82, 0x90, 0x56, 0x0b, 0x69, 0x0c, 0x06,
	0x19, 0x34, 0x1a, 0x5f, 0x48, 0xe2, 0xfd, 0x4a, 0xf2, 0x0f, 0x3f, 0x58, 0xe3, 0x49, 0x78, 0xb6,
	0xe0, 0x45, 0x7e, 0x16, 0x1f, 0x3a, 0x0f, 0x94, 0xdd, 0x80, 0xcb, 0xc0, 0x5d, 0x99, 0x09, 0x0e,
	0xc1, 0x5b, 0x34, 0x1e, 0x58, 0x4c, 0x9b, 0xd6, 0x61, 0xa6, 0x13, 0x70, 0x11, 0x39, 0x24, 0xc7,
	0xad, 0xe1, 0xf2, 0xce, 0xf6, 0x28, 0xf5, 0xb9, 0x62, 0x34, 0x77, 0x69, 0x54, 0xcb, 0x69, 0xab,
	0x78, 0xb9, 0x75, 0x29, 0xdb, 0xa1, 0x0d, 0x25, 0x47, 0x0a, 0x5c, 0x88, 0x36, 0x72, 0x56, 0x57,
	0xf2, 0x12, 0x5c, 0x38, 0x7d, 0x23, 0xb4, 0x79, 0x5d, 0x16, 0x66, 0xaf, 0x84, 0xd2, 0x95, 0x2f,
	0x6b, 0xf3, 0x65, 0xf5, 0x22, 0x23, 0xbf, 0x43, 0x9d, 0x0c, 0xe1, 0x69, 0x0e, 0x3e, 0xc4, 0x47,
	0xfc, 0x8f, 0x61, 0xf8, 0xef, 0xd8, 0x9d, 0x8b, 0xc5, 0x7b, 0x54, 0x6b, 0x92, 0xc5, 0xc7, 0xe7,
	0x4b, 0xed, 0x8c, 0xf5, 0xc4, 0x68, 0xad, 0xbd, 0xd2, 0xa2, 0xd4, 0x8a, 0x6a, 0x78, 0x51, 0xe4,
	0x3e, 0xd1, 0x66, 0x82, 0xe2, 0xd1, 0xa3, 0xe9, 0x9f, 0xd3, 0x5d, 0x85, 0xb3, 0x95, 0xa1, 0xb4,
	0xfa, 0x87, 0x69, 0x7f, 0xb3, 0x6a, 0x30, 0xf8, 0x5e, 0x6f, 0x40, 0xee, 0x1b, 0x25, 0x19, 0xd7,
	0xf3, 0x3d, 0x7b, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc7, 0xec, 0x3a, 0x0d, 0xd5, 0x01, 0x00,
	0x00,
}