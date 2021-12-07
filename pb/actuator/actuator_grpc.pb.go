// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package actuator

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ActuatorServiceClient is the client API for ActuatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActuatorServiceClient interface {
	GetHealth(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Health, error)
	GetInfo(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Info, error)
}

type actuatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewActuatorServiceClient(cc grpc.ClientConnInterface) ActuatorServiceClient {
	return &actuatorServiceClient{cc}
}

func (c *actuatorServiceClient) GetHealth(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Health, error) {
	out := new(Health)
	err := c.cc.Invoke(ctx, "/actuator.ActuatorService/GetHealth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actuatorServiceClient) GetInfo(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Info, error) {
	out := new(Info)
	err := c.cc.Invoke(ctx, "/actuator.ActuatorService/GetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActuatorServiceServer is the server API for ActuatorService service.
// All implementations must embed UnimplementedActuatorServiceServer
// for forward compatibility
type ActuatorServiceServer interface {
	GetHealth(context.Context, *empty.Empty) (*Health, error)
	GetInfo(context.Context, *empty.Empty) (*Info, error)
	mustEmbedUnimplementedActuatorServiceServer()
}

// UnimplementedActuatorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedActuatorServiceServer struct {
}

func (UnimplementedActuatorServiceServer) GetHealth(context.Context, *empty.Empty) (*Health, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHealth not implemented")
}
func (UnimplementedActuatorServiceServer) GetInfo(context.Context, *empty.Empty) (*Info, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfo not implemented")
}
func (UnimplementedActuatorServiceServer) mustEmbedUnimplementedActuatorServiceServer() {}

// UnsafeActuatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActuatorServiceServer will
// result in compilation errors.
type UnsafeActuatorServiceServer interface {
	mustEmbedUnimplementedActuatorServiceServer()
}

func RegisterActuatorServiceServer(s *grpc.Server, srv ActuatorServiceServer) {
	s.RegisterService(&_ActuatorService_serviceDesc, srv)
}

func _ActuatorService_GetHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActuatorServiceServer).GetHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/actuator.ActuatorService/GetHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActuatorServiceServer).GetHealth(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActuatorService_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActuatorServiceServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/actuator.ActuatorService/GetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActuatorServiceServer).GetInfo(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ActuatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "actuator.ActuatorService",
	HandlerType: (*ActuatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHealth",
			Handler:    _ActuatorService_GetHealth_Handler,
		},
		{
			MethodName: "GetInfo",
			Handler:    _ActuatorService_GetInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "actuator.proto",
}