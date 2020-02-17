// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: credentials.proto

package credentials

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	uuid "github.com/koverto/uuid"
	grpc "google.golang.org/grpc"
	io "io"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type CredentialType int32

const (
	CredentialType_NONE     CredentialType = 0
	CredentialType_PASSWORD CredentialType = 1
)

var CredentialType_name = map[int32]string{
	0: "NONE",
	1: "PASSWORD",
}

var CredentialType_value = map[string]int32{
	"NONE":     0,
	"PASSWORD": 1,
}

func (x CredentialType) String() string {
	return proto.EnumName(CredentialType_name, int32(x))
}

func (CredentialType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9f10b41ff9e8e07a, []int{0}
}

type Credential struct {
	Id             *uuid.UUID     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"_id"`
	UserID         *uuid.UUID     `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
	CredentialType CredentialType `protobuf:"varint,3,opt,name=credentialType,proto3,enum=credentials.CredentialType" json:"credentialType,omitempty"`
	Credential     []byte         `protobuf:"bytes,4,opt,name=credential,proto3" json:"credential,omitempty"`
}

func (m *Credential) Reset()         { *m = Credential{} }
func (m *Credential) String() string { return proto.CompactTextString(m) }
func (*Credential) ProtoMessage()    {}
func (*Credential) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f10b41ff9e8e07a, []int{0}
}
func (m *Credential) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Credential) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Credential.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Credential) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Credential.Merge(m, src)
}
func (m *Credential) XXX_Size() int {
	return m.Size()
}
func (m *Credential) XXX_DiscardUnknown() {
	xxx_messageInfo_Credential.DiscardUnknown(m)
}

var xxx_messageInfo_Credential proto.InternalMessageInfo

func (m *Credential) GetId() *uuid.UUID {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Credential) GetUserID() *uuid.UUID {
	if m != nil {
		return m.UserID
	}
	return nil
}

func (m *Credential) GetCredentialType() CredentialType {
	if m != nil {
		return m.CredentialType
	}
	return CredentialType_NONE
}

func (m *Credential) GetCredential() []byte {
	if m != nil {
		return m.Credential
	}
	return nil
}

type CredentialResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (m *CredentialResponse) Reset()         { *m = CredentialResponse{} }
func (m *CredentialResponse) String() string { return proto.CompactTextString(m) }
func (*CredentialResponse) ProtoMessage()    {}
func (*CredentialResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f10b41ff9e8e07a, []int{1}
}
func (m *CredentialResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CredentialResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CredentialResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CredentialResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CredentialResponse.Merge(m, src)
}
func (m *CredentialResponse) XXX_Size() int {
	return m.Size()
}
func (m *CredentialResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CredentialResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CredentialResponse proto.InternalMessageInfo

func (m *CredentialResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type CredentialUpdate struct {
	Current *Credential `protobuf:"bytes,1,opt,name=current,proto3" json:"current,omitempty"`
	New     *Credential `protobuf:"bytes,2,opt,name=new,proto3" json:"new,omitempty"`
}

func (m *CredentialUpdate) Reset()         { *m = CredentialUpdate{} }
func (m *CredentialUpdate) String() string { return proto.CompactTextString(m) }
func (*CredentialUpdate) ProtoMessage()    {}
func (*CredentialUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f10b41ff9e8e07a, []int{2}
}
func (m *CredentialUpdate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CredentialUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CredentialUpdate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CredentialUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CredentialUpdate.Merge(m, src)
}
func (m *CredentialUpdate) XXX_Size() int {
	return m.Size()
}
func (m *CredentialUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_CredentialUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_CredentialUpdate proto.InternalMessageInfo

func (m *CredentialUpdate) GetCurrent() *Credential {
	if m != nil {
		return m.Current
	}
	return nil
}

func (m *CredentialUpdate) GetNew() *Credential {
	if m != nil {
		return m.New
	}
	return nil
}

func init() {
	proto.RegisterEnum("credentials.CredentialType", CredentialType_name, CredentialType_value)
	proto.RegisterType((*Credential)(nil), "credentials.Credential")
	proto.RegisterType((*CredentialResponse)(nil), "credentials.CredentialResponse")
	proto.RegisterType((*CredentialUpdate)(nil), "credentials.CredentialUpdate")
}

func init() { proto.RegisterFile("credentials.proto", fileDescriptor_9f10b41ff9e8e07a) }

var fileDescriptor_9f10b41ff9e8e07a = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xb1, 0xee, 0xd2, 0x50,
	0x14, 0xc6, 0x7b, 0x81, 0x94, 0xe6, 0x40, 0x1a, 0xbc, 0x8b, 0x0d, 0xc6, 0x42, 0xee, 0x54, 0x19,
	0x4a, 0x00, 0x27, 0x27, 0xa5, 0xd5, 0x04, 0x07, 0x30, 0x17, 0xd1, 0xd1, 0xd0, 0xf6, 0x8a, 0x8d,
	0xd8, 0xdb, 0xf4, 0xb6, 0x18, 0xdf, 0xc2, 0x57, 0x72, 0x73, 0x64, 0x74, 0x32, 0x06, 0x06, 0x77,
	0x9f, 0xc0, 0x70, 0x29, 0xb6, 0x90, 0x34, 0x31, 0xf9, 0x2f, 0x4d, 0xfb, 0x9d, 0xdf, 0xf9, 0xce,
	0xf9, 0x4e, 0x0a, 0xf7, 0xfc, 0x84, 0x05, 0x2c, 0x4a, 0xc3, 0xf5, 0x56, 0xd8, 0x71, 0xc2, 0x53,
	0x8e, 0x5b, 0x25, 0xa9, 0xfb, 0x78, 0x13, 0xa6, 0x1f, 0x32, 0xcf, 0xf6, 0xf9, 0xa7, 0xe1, 0x86,
	0x6f, 0xf8, 0x50, 0x32, 0x5e, 0xf6, 0xfe, 0xe9, 0x6e, 0x64, 0x4f, 0xec, 0x91, 0x14, 0xa5, 0x26,
	0xdf, 0xce, 0x16, 0x5d, 0x52, 0xea, 0xfa, 0xc8, 0x77, 0x2c, 0x49, 0xf9, 0x30, 0xcb, 0xc2, 0x40,
	0x3e, 0xce, 0x0c, 0xf9, 0x86, 0x00, 0x9c, 0x7f, 0x93, 0xf0, 0x00, 0x6a, 0x61, 0x60, 0xa0, 0x3e,
	0xb2, 0x5a, 0x63, 0xb0, 0x25, 0xb7, 0x5a, 0xcd, 0xdc, 0xa9, 0xfe, 0xe7, 0x67, 0x0f, 0x3c, 0xc1,
	0xa3, 0x27, 0xe4, 0x5d, 0x18, 0x10, 0x5a, 0x0b, 0x03, 0x4c, 0x40, 0xcd, 0x04, 0x4b, 0x66, 0xae,
	0x51, 0xbb, 0xe5, 0x69, 0x5e, 0xc1, 0x0e, 0xe8, 0x45, 0x8e, 0xd7, 0x5f, 0x62, 0x66, 0xd4, 0xfb,
	0xc8, 0xd2, 0xc7, 0x0f, 0xec, 0x72, 0x62, 0xe7, 0x0a, 0xa1, 0x37, 0x2d, 0xd8, 0x04, 0x28, 0x14,
	0xa3, 0xd1, 0x47, 0x56, 0x9b, 0x96, 0x14, 0x62, 0x03, 0x2e, 0x1c, 0x28, 0x13, 0x31, 0x8f, 0x04,
	0xc3, 0x06, 0x34, 0x45, 0xe6, 0xfb, 0x4c, 0x08, 0x99, 0x47, 0xa3, 0x97, 0x4f, 0x12, 0x43, 0xa7,
	0xe0, 0x57, 0x71, 0xb0, 0x4e, 0x19, 0x1e, 0x41, 0xd3, 0xcf, 0x92, 0x84, 0x45, 0x69, 0x9e, 0xfe,
	0x7e, 0xc5, 0x86, 0xf4, 0xc2, 0xe1, 0x47, 0x50, 0x8f, 0xd8, 0xe7, 0x3c, 0x7c, 0x25, 0x7e, 0x62,
	0x06, 0x16, 0xe8, 0xd7, 0x19, 0xb1, 0x06, 0x8d, 0xf9, 0x62, 0xfe, 0xbc, 0xa3, 0xe0, 0x36, 0x68,
	0xaf, 0x9e, 0x2d, 0x97, 0x6f, 0x17, 0xd4, 0xed, 0xa0, 0xf1, 0x6f, 0x04, 0xad, 0x02, 0x15, 0xd8,
	0x05, 0xd5, 0x49, 0xd8, 0x69, 0xc3, 0xaa, 0x09, 0xdd, 0x5e, 0xd5, 0xe8, 0xfc, 0x12, 0x44, 0xc1,
	0x2f, 0x40, 0x7b, 0xb3, 0xde, 0x86, 0xc1, 0x5d, 0x7d, 0x5e, 0x82, 0x9a, 0xdf, 0xeb, 0x61, 0x05,
	0x7c, 0x2e, 0xff, 0x87, 0xd7, 0xd4, 0xf8, 0x7e, 0x30, 0xd1, 0xfe, 0x60, 0xa2, 0x5f, 0x07, 0x13,
	0x7d, 0x3d, 0x9a, 0xca, 0xfe, 0x68, 0x2a, 0x3f, 0x8e, 0xa6, 0xe2, 0xa9, 0xf2, 0xd7, 0x9c, 0xfc,
	0x0d, 0x00, 0x00, 0xff, 0xff, 0xa6, 0xbc, 0xab, 0xf4, 0x16, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CredentialsClient is the client API for Credentials service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CredentialsClient interface {
	Create(ctx context.Context, in *Credential, opts ...grpc.CallOption) (*CredentialResponse, error)
	Validate(ctx context.Context, in *Credential, opts ...grpc.CallOption) (*CredentialResponse, error)
	Update(ctx context.Context, in *CredentialUpdate, opts ...grpc.CallOption) (*CredentialResponse, error)
}

type credentialsClient struct {
	cc *grpc.ClientConn
}

func NewCredentialsClient(cc *grpc.ClientConn) CredentialsClient {
	return &credentialsClient{cc}
}

func (c *credentialsClient) Create(ctx context.Context, in *Credential, opts ...grpc.CallOption) (*CredentialResponse, error) {
	out := new(CredentialResponse)
	err := c.cc.Invoke(ctx, "/credentials.Credentials/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *credentialsClient) Validate(ctx context.Context, in *Credential, opts ...grpc.CallOption) (*CredentialResponse, error) {
	out := new(CredentialResponse)
	err := c.cc.Invoke(ctx, "/credentials.Credentials/Validate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *credentialsClient) Update(ctx context.Context, in *CredentialUpdate, opts ...grpc.CallOption) (*CredentialResponse, error) {
	out := new(CredentialResponse)
	err := c.cc.Invoke(ctx, "/credentials.Credentials/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CredentialsServer is the server API for Credentials service.
type CredentialsServer interface {
	Create(context.Context, *Credential) (*CredentialResponse, error)
	Validate(context.Context, *Credential) (*CredentialResponse, error)
	Update(context.Context, *CredentialUpdate) (*CredentialResponse, error)
}

func RegisterCredentialsServer(s *grpc.Server, srv CredentialsServer) {
	s.RegisterService(&_Credentials_serviceDesc, srv)
}

func _Credentials_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Credential)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/credentials.Credentials/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialsServer).Create(ctx, req.(*Credential))
	}
	return interceptor(ctx, in, info, handler)
}

func _Credentials_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Credential)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialsServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/credentials.Credentials/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialsServer).Validate(ctx, req.(*Credential))
	}
	return interceptor(ctx, in, info, handler)
}

func _Credentials_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CredentialUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialsServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/credentials.Credentials/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialsServer).Update(ctx, req.(*CredentialUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

var _Credentials_serviceDesc = grpc.ServiceDesc{
	ServiceName: "credentials.Credentials",
	HandlerType: (*CredentialsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Credentials_Create_Handler,
		},
		{
			MethodName: "Validate",
			Handler:    _Credentials_Validate_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Credentials_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "credentials.proto",
}

func (m *Credential) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Credential) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCredentials(dAtA, i, uint64(m.Id.Size()))
		n1, err := m.Id.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.UserID != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCredentials(dAtA, i, uint64(m.UserID.Size()))
		n2, err := m.UserID.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.CredentialType != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintCredentials(dAtA, i, uint64(m.CredentialType))
	}
	if len(m.Credential) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintCredentials(dAtA, i, uint64(len(m.Credential)))
		i += copy(dAtA[i:], m.Credential)
	}
	return i, nil
}

func (m *CredentialResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CredentialResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Success {
		dAtA[i] = 0x8
		i++
		if m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *CredentialUpdate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CredentialUpdate) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Current != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCredentials(dAtA, i, uint64(m.Current.Size()))
		n3, err := m.Current.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.New != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCredentials(dAtA, i, uint64(m.New.Size()))
		n4, err := m.New.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func encodeVarintCredentials(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Credential) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != nil {
		l = m.Id.Size()
		n += 1 + l + sovCredentials(uint64(l))
	}
	if m.UserID != nil {
		l = m.UserID.Size()
		n += 1 + l + sovCredentials(uint64(l))
	}
	if m.CredentialType != 0 {
		n += 1 + sovCredentials(uint64(m.CredentialType))
	}
	l = len(m.Credential)
	if l > 0 {
		n += 1 + l + sovCredentials(uint64(l))
	}
	return n
}

func (m *CredentialResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Success {
		n += 2
	}
	return n
}

func (m *CredentialUpdate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Current != nil {
		l = m.Current.Size()
		n += 1 + l + sovCredentials(uint64(l))
	}
	if m.New != nil {
		l = m.New.Size()
		n += 1 + l + sovCredentials(uint64(l))
	}
	return n
}

func sovCredentials(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCredentials(x uint64) (n int) {
	return sovCredentials(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Credential) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCredentials
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Credential: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Credential: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredentials
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCredentials
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Id == nil {
				m.Id = &uuid.UUID{}
			}
			if err := m.Id.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredentials
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCredentials
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.UserID == nil {
				m.UserID = &uuid.UUID{}
			}
			if err := m.UserID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CredentialType", wireType)
			}
			m.CredentialType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredentials
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CredentialType |= (CredentialType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Credential", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredentials
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCredentials
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Credential = append(m.Credential[:0], dAtA[iNdEx:postIndex]...)
			if m.Credential == nil {
				m.Credential = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCredentials(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCredentials
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CredentialResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCredentials
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CredentialResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CredentialResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredentials
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Success = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipCredentials(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCredentials
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CredentialUpdate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCredentials
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CredentialUpdate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CredentialUpdate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Current", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredentials
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCredentials
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Current == nil {
				m.Current = &Credential{}
			}
			if err := m.Current.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field New", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredentials
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCredentials
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.New == nil {
				m.New = &Credential{}
			}
			if err := m.New.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCredentials(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCredentials
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCredentials(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCredentials
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCredentials
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCredentials
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthCredentials
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCredentials
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCredentials(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCredentials = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCredentials   = fmt.Errorf("proto: integer overflow")
)