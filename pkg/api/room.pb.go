// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: room.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type RoomConnect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId string `protobuf:"bytes,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	// string roomName = 4;
	UserId string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Name   string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RoomConnect) Reset() {
	*x = RoomConnect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomConnect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomConnect) ProtoMessage() {}

func (x *RoomConnect) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomConnect.ProtoReflect.Descriptor instead.
func (*RoomConnect) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{0}
}

func (x *RoomConnect) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *RoomConnect) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RoomConnect) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Room struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	MemberCount       int32  `protobuf:"varint,2,opt,name=memberCount,proto3" json:"memberCount,omitempty"`
	ActiveMemberCount int32  `protobuf:"varint,3,opt,name=activeMemberCount,proto3" json:"activeMemberCount,omitempty"` // string name = 2;
}

func (x *Room) Reset() {
	*x = Room{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Room) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Room) ProtoMessage() {}

func (x *Room) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Room.ProtoReflect.Descriptor instead.
func (*Room) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{1}
}

func (x *Room) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Room) GetMemberCount() int32 {
	if x != nil {
		return x.MemberCount
	}
	return 0
}

func (x *Room) GetActiveMemberCount() int32 {
	if x != nil {
		return x.ActiveMemberCount
	}
	return 0
}

type RoomMember struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Present bool   `protobuf:"varint,3,opt,name=present,proto3" json:"present,omitempty"`
}

func (x *RoomMember) Reset() {
	*x = RoomMember{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomMember) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomMember) ProtoMessage() {}

func (x *RoomMember) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomMember.ProtoReflect.Descriptor instead.
func (*RoomMember) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{2}
}

func (x *RoomMember) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RoomMember) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RoomMember) GetPresent() bool {
	if x != nil {
		return x.Present
	}
	return false
}

var File_room_proto protoreflect.FileDescriptor

var file_room_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x72, 0x6f,
	0x6f, 0x6d, 0x22, 0x51, 0x0a, 0x0b, 0x52, 0x6f, 0x6f, 0x6d, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x66, 0x0a, 0x04, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x2c, 0x0a, 0x11, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x52, 0x0a,
	0x0a, 0x52, 0x6f, 0x6f, 0x6d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x65, 0x73, 0x65,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e,
	0x74, 0x32, 0x75, 0x0a, 0x0b, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x32, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x54, 0x6f, 0x52, 0x6f, 0x6f,
	0x6d, 0x12, 0x11, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x1a, 0x0a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x52, 0x6f, 0x6f, 0x6d,
	0x22, 0x00, 0x30, 0x01, 0x12, 0x32, 0x0a, 0x07, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12,
	0x11, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x1a, 0x10, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x22, 0x00, 0x30, 0x01, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_room_proto_rawDescOnce sync.Once
	file_room_proto_rawDescData = file_room_proto_rawDesc
)

func file_room_proto_rawDescGZIP() []byte {
	file_room_proto_rawDescOnce.Do(func() {
		file_room_proto_rawDescData = protoimpl.X.CompressGZIP(file_room_proto_rawDescData)
	})
	return file_room_proto_rawDescData
}

var file_room_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_room_proto_goTypes = []interface{}{
	(*RoomConnect)(nil), // 0: room.RoomConnect
	(*Room)(nil),        // 1: room.Room
	(*RoomMember)(nil),  // 2: room.RoomMember
}
var file_room_proto_depIdxs = []int32{
	0, // 0: room.RoomService.ConnectToRoom:input_type -> room.RoomConnect
	0, // 1: room.RoomService.Members:input_type -> room.RoomConnect
	1, // 2: room.RoomService.ConnectToRoom:output_type -> room.Room
	2, // 3: room.RoomService.Members:output_type -> room.RoomMember
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_room_proto_init() }
func file_room_proto_init() {
	if File_room_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_room_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomConnect); i {
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
		file_room_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Room); i {
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
		file_room_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomMember); i {
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
			RawDescriptor: file_room_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_room_proto_goTypes,
		DependencyIndexes: file_room_proto_depIdxs,
		MessageInfos:      file_room_proto_msgTypes,
	}.Build()
	File_room_proto = out.File
	file_room_proto_rawDesc = nil
	file_room_proto_goTypes = nil
	file_room_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RoomServiceClient is the client API for RoomService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RoomServiceClient interface {
	ConnectToRoom(ctx context.Context, in *RoomConnect, opts ...grpc.CallOption) (RoomService_ConnectToRoomClient, error)
	Members(ctx context.Context, in *RoomConnect, opts ...grpc.CallOption) (RoomService_MembersClient, error)
}

type roomServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRoomServiceClient(cc grpc.ClientConnInterface) RoomServiceClient {
	return &roomServiceClient{cc}
}

func (c *roomServiceClient) ConnectToRoom(ctx context.Context, in *RoomConnect, opts ...grpc.CallOption) (RoomService_ConnectToRoomClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RoomService_serviceDesc.Streams[0], "/room.RoomService/ConnectToRoom", opts...)
	if err != nil {
		return nil, err
	}
	x := &roomServiceConnectToRoomClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RoomService_ConnectToRoomClient interface {
	Recv() (*Room, error)
	grpc.ClientStream
}

type roomServiceConnectToRoomClient struct {
	grpc.ClientStream
}

func (x *roomServiceConnectToRoomClient) Recv() (*Room, error) {
	m := new(Room)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *roomServiceClient) Members(ctx context.Context, in *RoomConnect, opts ...grpc.CallOption) (RoomService_MembersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RoomService_serviceDesc.Streams[1], "/room.RoomService/Members", opts...)
	if err != nil {
		return nil, err
	}
	x := &roomServiceMembersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RoomService_MembersClient interface {
	Recv() (*RoomMember, error)
	grpc.ClientStream
}

type roomServiceMembersClient struct {
	grpc.ClientStream
}

func (x *roomServiceMembersClient) Recv() (*RoomMember, error) {
	m := new(RoomMember)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RoomServiceServer is the server API for RoomService service.
type RoomServiceServer interface {
	ConnectToRoom(*RoomConnect, RoomService_ConnectToRoomServer) error
	Members(*RoomConnect, RoomService_MembersServer) error
}

// UnimplementedRoomServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRoomServiceServer struct {
}

func (*UnimplementedRoomServiceServer) ConnectToRoom(*RoomConnect, RoomService_ConnectToRoomServer) error {
	return status.Errorf(codes.Unimplemented, "method ConnectToRoom not implemented")
}
func (*UnimplementedRoomServiceServer) Members(*RoomConnect, RoomService_MembersServer) error {
	return status.Errorf(codes.Unimplemented, "method Members not implemented")
}

func RegisterRoomServiceServer(s *grpc.Server, srv RoomServiceServer) {
	s.RegisterService(&_RoomService_serviceDesc, srv)
}

func _RoomService_ConnectToRoom_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RoomConnect)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RoomServiceServer).ConnectToRoom(m, &roomServiceConnectToRoomServer{stream})
}

type RoomService_ConnectToRoomServer interface {
	Send(*Room) error
	grpc.ServerStream
}

type roomServiceConnectToRoomServer struct {
	grpc.ServerStream
}

func (x *roomServiceConnectToRoomServer) Send(m *Room) error {
	return x.ServerStream.SendMsg(m)
}

func _RoomService_Members_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RoomConnect)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RoomServiceServer).Members(m, &roomServiceMembersServer{stream})
}

type RoomService_MembersServer interface {
	Send(*RoomMember) error
	grpc.ServerStream
}

type roomServiceMembersServer struct {
	grpc.ServerStream
}

func (x *roomServiceMembersServer) Send(m *RoomMember) error {
	return x.ServerStream.SendMsg(m)
}

var _RoomService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "room.RoomService",
	HandlerType: (*RoomServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ConnectToRoom",
			Handler:       _RoomService_ConnectToRoom_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Members",
			Handler:       _RoomService_Members_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "room.proto",
}
