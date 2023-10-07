// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.11.2
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

type Signal int32

const (
	Signal_UNKNOWN Signal = 0
	Signal_CLOSE   Signal = 1
	Signal_OPEN    Signal = 2
)

// Enum value maps for Signal.
var (
	Signal_name = map[int32]string{
		0: "UNKNOWN",
		1: "CLOSE",
		2: "OPEN",
	}
	Signal_value = map[string]int32{
		"UNKNOWN": 0,
		"CLOSE":   1,
		"OPEN":    2,
	}
)

func (x Signal) Enum() *Signal {
	p := new(Signal)
	*p = x
	return p
}

func (x Signal) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Signal) Descriptor() protoreflect.EnumDescriptor {
	return file_idl_tunnel_proto_enumTypes[0].Descriptor()
}

func (Signal) Type() protoreflect.EnumType {
	return &file_idl_tunnel_proto_enumTypes[0]
}

func (x Signal) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Signal.Descriptor instead.
func (Signal) EnumDescriptor() ([]byte, []int) {
	return file_idl_tunnel_proto_rawDescGZIP(), []int{0}
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seq      int64  `protobuf:"varint,1,opt,name=Seq,proto3" json:"Seq,omitempty"`
	Payload  []byte `protobuf:"bytes,2,opt,name=Payload,proto3" json:"Payload,omitempty"`
	Signal   Signal `protobuf:"varint,3,opt,name=Signal,proto3,enum=Signal" json:"Signal,omitempty"`
	ClientId string `protobuf:"bytes,4,opt,name=ClientId,proto3" json:"ClientId,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_tunnel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_idl_tunnel_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetSeq() int64 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *Request) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *Request) GetSignal() Signal {
	if x != nil {
		return x.Signal
	}
	return Signal_UNKNOWN
}

func (x *Request) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seq     int64  `protobuf:"varint,1,opt,name=Seq,proto3" json:"Seq,omitempty"`
	Payload []byte `protobuf:"bytes,2,opt,name=Payload,proto3" json:"Payload,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_tunnel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_idl_tunnel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_idl_tunnel_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetSeq() int64 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *Response) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type Tunnel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Port     int64  `protobuf:"varint,2,opt,name=Port,proto3" json:"Port,omitempty"`
	ClientId string `protobuf:"bytes,3,opt,name=ClientId,proto3" json:"ClientId,omitempty"`
}

func (x *Tunnel) Reset() {
	*x = Tunnel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_tunnel_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tunnel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tunnel) ProtoMessage() {}

func (x *Tunnel) ProtoReflect() protoreflect.Message {
	mi := &file_idl_tunnel_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tunnel.ProtoReflect.Descriptor instead.
func (*Tunnel) Descriptor() ([]byte, []int) {
	return file_idl_tunnel_proto_rawDescGZIP(), []int{2}
}

func (x *Tunnel) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Tunnel) GetPort() int64 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Tunnel) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

var File_idl_tunnel_proto protoreflect.FileDescriptor

var file_idl_tunnel_proto_rawDesc = []byte{
	0x0a, 0x10, 0x69, 0x64, 0x6c, 0x2f, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x72, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x53, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x53, 0x65, 0x71, 0x12,
	0x18, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1f, 0x0a, 0x06, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x6c, 0x52, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x36, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x53, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x48,
	0x0a, 0x06, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x2a, 0x2a, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x6c, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12,
	0x09, 0x0a, 0x05, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4f, 0x50,
	0x45, 0x4e, 0x10, 0x02, 0x32, 0x4d, 0x0a, 0x0c, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x08, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x1c, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x07, 0x2e, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x1a, 0x07, 0x2e, 0x54, 0x75, 0x6e,
	0x6e, 0x65, 0x6c, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_idl_tunnel_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_idl_tunnel_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_idl_tunnel_proto_goTypes = []interface{}{
	(Signal)(0),      // 0: Signal
	(*Request)(nil),  // 1: Request
	(*Response)(nil), // 2: Response
	(*Tunnel)(nil),   // 3: Tunnel
}
var file_idl_tunnel_proto_depIdxs = []int32{
	0, // 0: Request.Signal:type_name -> Signal
	1, // 1: TunnelServer.Call:input_type -> Request
	3, // 2: TunnelServer.Register:input_type -> Tunnel
	2, // 3: TunnelServer.Call:output_type -> Response
	3, // 4: TunnelServer.Register:output_type -> Tunnel
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_idl_tunnel_proto_init() }
func file_idl_tunnel_proto_init() {
	if File_idl_tunnel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idl_tunnel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_idl_tunnel_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_idl_tunnel_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tunnel); i {
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
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idl_tunnel_proto_goTypes,
		DependencyIndexes: file_idl_tunnel_proto_depIdxs,
		EnumInfos:         file_idl_tunnel_proto_enumTypes,
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
	Call(ctx context.Context, opts ...grpc.CallOption) (TunnelServer_CallClient, error)
	Register(ctx context.Context, in *Tunnel, opts ...grpc.CallOption) (*Tunnel, error)
}

type tunnelServerClient struct {
	cc grpc.ClientConnInterface
}

func NewTunnelServerClient(cc grpc.ClientConnInterface) TunnelServerClient {
	return &tunnelServerClient{cc}
}

func (c *tunnelServerClient) Call(ctx context.Context, opts ...grpc.CallOption) (TunnelServer_CallClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TunnelServer_serviceDesc.Streams[0], "/TunnelServer/Call", opts...)
	if err != nil {
		return nil, err
	}
	x := &tunnelServerCallClient{stream}
	return x, nil
}

type TunnelServer_CallClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type tunnelServerCallClient struct {
	grpc.ClientStream
}

func (x *tunnelServerCallClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tunnelServerCallClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tunnelServerClient) Register(ctx context.Context, in *Tunnel, opts ...grpc.CallOption) (*Tunnel, error) {
	out := new(Tunnel)
	err := c.cc.Invoke(ctx, "/TunnelServer/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TunnelServerServer is the server API for TunnelServer service.
type TunnelServerServer interface {
	Call(TunnelServer_CallServer) error
	Register(context.Context, *Tunnel) (*Tunnel, error)
}

// UnimplementedTunnelServerServer can be embedded to have forward compatible implementations.
type UnimplementedTunnelServerServer struct {
}

func (*UnimplementedTunnelServerServer) Call(TunnelServer_CallServer) error {
	return status.Errorf(codes.Unimplemented, "method Call not implemented")
}
func (*UnimplementedTunnelServerServer) Register(context.Context, *Tunnel) (*Tunnel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}

func RegisterTunnelServerServer(s *grpc.Server, srv TunnelServerServer) {
	s.RegisterService(&_TunnelServer_serviceDesc, srv)
}

func _TunnelServer_Call_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TunnelServerServer).Call(&tunnelServerCallServer{stream})
}

type TunnelServer_CallServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type tunnelServerCallServer struct {
	grpc.ServerStream
}

func (x *tunnelServerCallServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tunnelServerCallServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TunnelServer_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Tunnel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TunnelServerServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TunnelServer/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TunnelServerServer).Register(ctx, req.(*Tunnel))
	}
	return interceptor(ctx, in, info, handler)
}

var _TunnelServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "TunnelServer",
	HandlerType: (*TunnelServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _TunnelServer_Register_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Call",
			Handler:       _TunnelServer_Call_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "idl/tunnel.proto",
}
