// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.20.3
// source: idl/tunnel.proto

package protogen

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Package struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload []byte `protobuf:"bytes,1,opt,name=Payload,proto3" json:"Payload,omitempty"`
}

func (x *Package) Reset() {
	*x = Package{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_tunnel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Package) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Package) ProtoMessage() {}

func (x *Package) ProtoReflect() protoreflect.Message {
	mi := &file_idl_tunnel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Package.ProtoReflect.Descriptor instead.
func (*Package) Descriptor() ([]byte, []int) {
	return file_idl_tunnel_proto_rawDescGZIP(), []int{0}
}

func (x *Package) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_idl_tunnel_proto protoreflect.FileDescriptor

var file_idl_tunnel_proto_rawDesc = []byte{
	0x0a, 0x10, 0x69, 0x64, 0x6c, 0x2f, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x23, 0x0a, 0x07, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x32, 0x6f, 0x0a, 0x0c, 0x54, 0x75, 0x6e, 0x6e, 0x65,
	0x6c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x12, 0x08, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x1a, 0x08, 0x2e, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0a, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x12, 0x08, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x1a, 0x08,
	0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x08, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x1a, 0x08, 0x2e, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_idl_tunnel_proto_rawDescOnce sync.Once
	file_idl_tunnel_proto_rawDescData = file_idl_tunnel_proto_rawDesc
)

func file_idl_tunnel_proto_rawDescGZIP() []byte {
	file_idl_tunnel_proto_rawDescOnce.Do(func() {
		file_idl_tunnel_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_tunnel_proto_rawDescData)
	})
	return file_idl_tunnel_proto_rawDescData
}

var file_idl_tunnel_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_idl_tunnel_proto_goTypes = []interface{}{
	(*Package)(nil), // 0: Package
}
var file_idl_tunnel_proto_depIdxs = []int32{
	0, // 0: TunnelServer.Connect:input_type -> Package
	0, // 1: TunnelServer.Disconnect:input_type -> Package
	0, // 2: TunnelServer.Data:input_type -> Package
	0, // 3: TunnelServer.Connect:output_type -> Package
	0, // 4: TunnelServer.Disconnect:output_type -> Package
	0, // 5: TunnelServer.Data:output_type -> Package
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_idl_tunnel_proto_init() }
func file_idl_tunnel_proto_init() {
	if File_idl_tunnel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idl_tunnel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Package); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_idl_tunnel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idl_tunnel_proto_goTypes,
		DependencyIndexes: file_idl_tunnel_proto_depIdxs,
		MessageInfos:      file_idl_tunnel_proto_msgTypes,
	}.Build()
	File_idl_tunnel_proto = out.File
	file_idl_tunnel_proto_rawDesc = nil
	file_idl_tunnel_proto_goTypes = nil
	file_idl_tunnel_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TunnelServerClient is the client API for TunnelServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TunnelServerClient interface {
	Connect(ctx context.Context, in *Package, opts ...grpc.CallOption) (*Package, error)
	Disconnect(ctx context.Context, in *Package, opts ...grpc.CallOption) (*Package, error)
	Data(ctx context.Context, opts ...grpc.CallOption) (TunnelServer_DataClient, error)
}

type tunnelServerClient struct {
	cc grpc.ClientConnInterface
}

func NewTunnelServerClient(cc grpc.ClientConnInterface) TunnelServerClient {
	return &tunnelServerClient{cc}
}

func (c *tunnelServerClient) Connect(ctx context.Context, in *Package, opts ...grpc.CallOption) (*Package, error) {
	out := new(Package)
	err := c.cc.Invoke(ctx, "/TunnelServer/Connect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tunnelServerClient) Disconnect(ctx context.Context, in *Package, opts ...grpc.CallOption) (*Package, error) {
	out := new(Package)
	err := c.cc.Invoke(ctx, "/TunnelServer/Disconnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tunnelServerClient) Data(ctx context.Context, opts ...grpc.CallOption) (TunnelServer_DataClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TunnelServer_serviceDesc.Streams[0], "/TunnelServer/Data", opts...)
	if err != nil {
		return nil, err
	}
	x := &tunnelServerDataClient{stream}
	return x, nil
}

type TunnelServer_DataClient interface {
	Send(*Package) error
	Recv() (*Package, error)
	grpc.ClientStream
}

type tunnelServerDataClient struct {
	grpc.ClientStream
}

func (x *tunnelServerDataClient) Send(m *Package) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tunnelServerDataClient) Recv() (*Package, error) {
	m := new(Package)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TunnelServerServer is the server API for TunnelServer service.
type TunnelServerServer interface {
	Connect(context.Context, *Package) (*Package, error)
	Disconnect(context.Context, *Package) (*Package, error)
	Data(TunnelServer_DataServer) error
}

// UnimplementedTunnelServerServer can be embedded to have forward compatible implementations.
type UnimplementedTunnelServerServer struct {
}

func (*UnimplementedTunnelServerServer) Connect(context.Context, *Package) (*Package, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (*UnimplementedTunnelServerServer) Disconnect(context.Context, *Package) (*Package, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disconnect not implemented")
}
func (*UnimplementedTunnelServerServer) Data(TunnelServer_DataServer) error {
	return status.Errorf(codes.Unimplemented, "method Data not implemented")
}

func RegisterTunnelServerServer(s *grpc.Server, srv TunnelServerServer) {
	s.RegisterService(&_TunnelServer_serviceDesc, srv)
}

func _TunnelServer_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Package)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TunnelServerServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TunnelServer/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TunnelServerServer).Connect(ctx, req.(*Package))
	}
	return interceptor(ctx, in, info, handler)
}

func _TunnelServer_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Package)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TunnelServerServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TunnelServer/Disconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TunnelServerServer).Disconnect(ctx, req.(*Package))
	}
	return interceptor(ctx, in, info, handler)
}

func _TunnelServer_Data_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TunnelServerServer).Data(&tunnelServerDataServer{stream})
}

type TunnelServer_DataServer interface {
	Send(*Package) error
	Recv() (*Package, error)
	grpc.ServerStream
}

type tunnelServerDataServer struct {
	grpc.ServerStream
}

func (x *tunnelServerDataServer) Send(m *Package) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tunnelServerDataServer) Recv() (*Package, error) {
	m := new(Package)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TunnelServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "TunnelServer",
	HandlerType: (*TunnelServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connect",
			Handler:    _TunnelServer_Connect_Handler,
		},
		{
			MethodName: "Disconnect",
			Handler:    _TunnelServer_Disconnect_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Data",
			Handler:       _TunnelServer_Data_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "idl/tunnel.proto",
}