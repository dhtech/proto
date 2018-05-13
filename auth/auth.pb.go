// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package auth

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type SshCertificateRequest struct {
	PublicKey            string   `protobuf:"bytes,1,opt,name=public_key,json=publicKey" json:"public_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SshCertificateRequest) Reset()         { *m = SshCertificateRequest{} }
func (m *SshCertificateRequest) String() string { return proto.CompactTextString(m) }
func (*SshCertificateRequest) ProtoMessage()    {}
func (*SshCertificateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_a4024bdf0d23e69e, []int{0}
}
func (m *SshCertificateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SshCertificateRequest.Unmarshal(m, b)
}
func (m *SshCertificateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SshCertificateRequest.Marshal(b, m, deterministic)
}
func (dst *SshCertificateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SshCertificateRequest.Merge(dst, src)
}
func (m *SshCertificateRequest) XXX_Size() int {
	return xxx_messageInfo_SshCertificateRequest.Size(m)
}
func (m *SshCertificateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SshCertificateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SshCertificateRequest proto.InternalMessageInfo

func (m *SshCertificateRequest) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

type SshCertificate struct {
	Certificate          string   `protobuf:"bytes,1,opt,name=certificate" json:"certificate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SshCertificate) Reset()         { *m = SshCertificate{} }
func (m *SshCertificate) String() string { return proto.CompactTextString(m) }
func (*SshCertificate) ProtoMessage()    {}
func (*SshCertificate) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_a4024bdf0d23e69e, []int{1}
}
func (m *SshCertificate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SshCertificate.Unmarshal(m, b)
}
func (m *SshCertificate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SshCertificate.Marshal(b, m, deterministic)
}
func (dst *SshCertificate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SshCertificate.Merge(dst, src)
}
func (m *SshCertificate) XXX_Size() int {
	return xxx_messageInfo_SshCertificate.Size(m)
}
func (m *SshCertificate) XXX_DiscardUnknown() {
	xxx_messageInfo_SshCertificate.DiscardUnknown(m)
}

var xxx_messageInfo_SshCertificate proto.InternalMessageInfo

func (m *SshCertificate) GetCertificate() string {
	if m != nil {
		return m.Certificate
	}
	return ""
}

type ClientValidation struct {
	// To make it harder to simply send a URL to another user and lure them into
	// logging in with the goal of stealing their session, we do a validation
	// by querying localhost:1215. If we get back the value presented in this
	// field the request is considered to be genuine.
	Ident                string   `protobuf:"bytes,1,opt,name=ident" json:"ident,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientValidation) Reset()         { *m = ClientValidation{} }
func (m *ClientValidation) String() string { return proto.CompactTextString(m) }
func (*ClientValidation) ProtoMessage()    {}
func (*ClientValidation) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_a4024bdf0d23e69e, []int{2}
}
func (m *ClientValidation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientValidation.Unmarshal(m, b)
}
func (m *ClientValidation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientValidation.Marshal(b, m, deterministic)
}
func (dst *ClientValidation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientValidation.Merge(dst, src)
}
func (m *ClientValidation) XXX_Size() int {
	return xxx_messageInfo_ClientValidation.Size(m)
}
func (m *ClientValidation) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientValidation.DiscardUnknown(m)
}

var xxx_messageInfo_ClientValidation proto.InternalMessageInfo

func (m *ClientValidation) GetIdent() string {
	if m != nil {
		return m.Ident
	}
	return ""
}

type VaultTokenRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VaultTokenRequest) Reset()         { *m = VaultTokenRequest{} }
func (m *VaultTokenRequest) String() string { return proto.CompactTextString(m) }
func (*VaultTokenRequest) ProtoMessage()    {}
func (*VaultTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_a4024bdf0d23e69e, []int{3}
}
func (m *VaultTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VaultTokenRequest.Unmarshal(m, b)
}
func (m *VaultTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VaultTokenRequest.Marshal(b, m, deterministic)
}
func (dst *VaultTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VaultTokenRequest.Merge(dst, src)
}
func (m *VaultTokenRequest) XXX_Size() int {
	return xxx_messageInfo_VaultTokenRequest.Size(m)
}
func (m *VaultTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VaultTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VaultTokenRequest proto.InternalMessageInfo

type VaultToken struct {
	Token                string   `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VaultToken) Reset()         { *m = VaultToken{} }
func (m *VaultToken) String() string { return proto.CompactTextString(m) }
func (*VaultToken) ProtoMessage()    {}
func (*VaultToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_a4024bdf0d23e69e, []int{4}
}
func (m *VaultToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VaultToken.Unmarshal(m, b)
}
func (m *VaultToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VaultToken.Marshal(b, m, deterministic)
}
func (dst *VaultToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VaultToken.Merge(dst, src)
}
func (m *VaultToken) XXX_Size() int {
	return xxx_messageInfo_VaultToken.Size(m)
}
func (m *VaultToken) XXX_DiscardUnknown() {
	xxx_messageInfo_VaultToken.DiscardUnknown(m)
}

var xxx_messageInfo_VaultToken proto.InternalMessageInfo

func (m *VaultToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type UserCredentialRequest struct {
	ClientValidation      *ClientValidation      `protobuf:"bytes,1,opt,name=client_validation,json=clientValidation" json:"client_validation,omitempty"`
	SshCertificateRequest *SshCertificateRequest `protobuf:"bytes,2,opt,name=ssh_certificate_request,json=sshCertificateRequest" json:"ssh_certificate_request,omitempty"`
	VaultTokenRequest     *VaultTokenRequest     `protobuf:"bytes,3,opt,name=vault_token_request,json=vaultTokenRequest" json:"vault_token_request,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}               `json:"-"`
	XXX_unrecognized      []byte                 `json:"-"`
	XXX_sizecache         int32                  `json:"-"`
}

func (m *UserCredentialRequest) Reset()         { *m = UserCredentialRequest{} }
func (m *UserCredentialRequest) String() string { return proto.CompactTextString(m) }
func (*UserCredentialRequest) ProtoMessage()    {}
func (*UserCredentialRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_a4024bdf0d23e69e, []int{5}
}
func (m *UserCredentialRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserCredentialRequest.Unmarshal(m, b)
}
func (m *UserCredentialRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserCredentialRequest.Marshal(b, m, deterministic)
}
func (dst *UserCredentialRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserCredentialRequest.Merge(dst, src)
}
func (m *UserCredentialRequest) XXX_Size() int {
	return xxx_messageInfo_UserCredentialRequest.Size(m)
}
func (m *UserCredentialRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserCredentialRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserCredentialRequest proto.InternalMessageInfo

func (m *UserCredentialRequest) GetClientValidation() *ClientValidation {
	if m != nil {
		return m.ClientValidation
	}
	return nil
}

func (m *UserCredentialRequest) GetSshCertificateRequest() *SshCertificateRequest {
	if m != nil {
		return m.SshCertificateRequest
	}
	return nil
}

func (m *UserCredentialRequest) GetVaultTokenRequest() *VaultTokenRequest {
	if m != nil {
		return m.VaultTokenRequest
	}
	return nil
}

type UserAction struct {
	// Interactive URL needs to be visited and acted on
	Url                  string   `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserAction) Reset()         { *m = UserAction{} }
func (m *UserAction) String() string { return proto.CompactTextString(m) }
func (*UserAction) ProtoMessage()    {}
func (*UserAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_a4024bdf0d23e69e, []int{6}
}
func (m *UserAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserAction.Unmarshal(m, b)
}
func (m *UserAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserAction.Marshal(b, m, deterministic)
}
func (dst *UserAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAction.Merge(dst, src)
}
func (m *UserAction) XXX_Size() int {
	return xxx_messageInfo_UserAction.Size(m)
}
func (m *UserAction) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAction.DiscardUnknown(m)
}

var xxx_messageInfo_UserAction proto.InternalMessageInfo

func (m *UserAction) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type CredentialResponse struct {
	// If set, the user needs to do something
	RequiredAction       *UserAction     `protobuf:"bytes,1,opt,name=required_action,json=requiredAction" json:"required_action,omitempty"`
	SshCertificate       *SshCertificate `protobuf:"bytes,2,opt,name=ssh_certificate,json=sshCertificate" json:"ssh_certificate,omitempty"`
	VaultToken           *VaultToken     `protobuf:"bytes,3,opt,name=vault_token,json=vaultToken" json:"vault_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CredentialResponse) Reset()         { *m = CredentialResponse{} }
func (m *CredentialResponse) String() string { return proto.CompactTextString(m) }
func (*CredentialResponse) ProtoMessage()    {}
func (*CredentialResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_a4024bdf0d23e69e, []int{7}
}
func (m *CredentialResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CredentialResponse.Unmarshal(m, b)
}
func (m *CredentialResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CredentialResponse.Marshal(b, m, deterministic)
}
func (dst *CredentialResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CredentialResponse.Merge(dst, src)
}
func (m *CredentialResponse) XXX_Size() int {
	return xxx_messageInfo_CredentialResponse.Size(m)
}
func (m *CredentialResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CredentialResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CredentialResponse proto.InternalMessageInfo

func (m *CredentialResponse) GetRequiredAction() *UserAction {
	if m != nil {
		return m.RequiredAction
	}
	return nil
}

func (m *CredentialResponse) GetSshCertificate() *SshCertificate {
	if m != nil {
		return m.SshCertificate
	}
	return nil
}

func (m *CredentialResponse) GetVaultToken() *VaultToken {
	if m != nil {
		return m.VaultToken
	}
	return nil
}

func init() {
	proto.RegisterType((*SshCertificateRequest)(nil), "auth.SshCertificateRequest")
	proto.RegisterType((*SshCertificate)(nil), "auth.SshCertificate")
	proto.RegisterType((*ClientValidation)(nil), "auth.ClientValidation")
	proto.RegisterType((*VaultTokenRequest)(nil), "auth.VaultTokenRequest")
	proto.RegisterType((*VaultToken)(nil), "auth.VaultToken")
	proto.RegisterType((*UserCredentialRequest)(nil), "auth.UserCredentialRequest")
	proto.RegisterType((*UserAction)(nil), "auth.UserAction")
	proto.RegisterType((*CredentialResponse)(nil), "auth.CredentialResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AuthenticationService service

type AuthenticationServiceClient interface {
	RequestUserCredential(ctx context.Context, in *UserCredentialRequest, opts ...grpc.CallOption) (AuthenticationService_RequestUserCredentialClient, error)
}

type authenticationServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthenticationServiceClient(cc *grpc.ClientConn) AuthenticationServiceClient {
	return &authenticationServiceClient{cc}
}

func (c *authenticationServiceClient) RequestUserCredential(ctx context.Context, in *UserCredentialRequest, opts ...grpc.CallOption) (AuthenticationService_RequestUserCredentialClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_AuthenticationService_serviceDesc.Streams[0], c.cc, "/auth.AuthenticationService/RequestUserCredential", opts...)
	if err != nil {
		return nil, err
	}
	x := &authenticationServiceRequestUserCredentialClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AuthenticationService_RequestUserCredentialClient interface {
	Recv() (*CredentialResponse, error)
	grpc.ClientStream
}

type authenticationServiceRequestUserCredentialClient struct {
	grpc.ClientStream
}

func (x *authenticationServiceRequestUserCredentialClient) Recv() (*CredentialResponse, error) {
	m := new(CredentialResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for AuthenticationService service

type AuthenticationServiceServer interface {
	RequestUserCredential(*UserCredentialRequest, AuthenticationService_RequestUserCredentialServer) error
}

func RegisterAuthenticationServiceServer(s *grpc.Server, srv AuthenticationServiceServer) {
	s.RegisterService(&_AuthenticationService_serviceDesc, srv)
}

func _AuthenticationService_RequestUserCredential_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UserCredentialRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AuthenticationServiceServer).RequestUserCredential(m, &authenticationServiceRequestUserCredentialServer{stream})
}

type AuthenticationService_RequestUserCredentialServer interface {
	Send(*CredentialResponse) error
	grpc.ServerStream
}

type authenticationServiceRequestUserCredentialServer struct {
	grpc.ServerStream
}

func (x *authenticationServiceRequestUserCredentialServer) Send(m *CredentialResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _AuthenticationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthenticationService",
	HandlerType: (*AuthenticationServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RequestUserCredential",
			Handler:       _AuthenticationService_RequestUserCredential_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "auth.proto",
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_auth_a4024bdf0d23e69e) }

var fileDescriptor_auth_a4024bdf0d23e69e = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0xc1, 0x6e, 0xe2, 0x30,
	0x10, 0x86, 0x37, 0xcb, 0xee, 0x4a, 0x0c, 0x12, 0x04, 0x43, 0x16, 0xb4, 0x68, 0x2b, 0xe4, 0x13,
	0x27, 0xd4, 0x52, 0xa9, 0x52, 0x0f, 0x3d, 0xa0, 0x1c, 0x7a, 0xe8, 0x2d, 0xb4, 0x5c, 0xa3, 0x60,
	0xa6, 0x8a, 0x45, 0x94, 0x50, 0xdb, 0x89, 0xc4, 0xe3, 0xf5, 0xb5, 0x7a, 0xaa, 0xec, 0x18, 0x12,
	0x08, 0xb7, 0x99, 0x91, 0xe7, 0x9f, 0xf9, 0xbf, 0x49, 0x00, 0xa2, 0x5c, 0xc5, 0xf3, 0xbd, 0xc8,
	0x54, 0x46, 0x7e, 0xe9, 0x98, 0x3e, 0x80, 0xb7, 0x92, 0xb1, 0x8f, 0x42, 0xf1, 0x77, 0xce, 0x22,
	0x85, 0x01, 0x7e, 0xe4, 0x28, 0x15, 0xf9, 0x0f, 0xb0, 0xcf, 0x37, 0x09, 0x67, 0xe1, 0x0e, 0x0f,
	0x63, 0x67, 0xea, 0xcc, 0xda, 0x41, 0xbb, 0xac, 0xbc, 0xe0, 0x81, 0x2e, 0xa0, 0x7b, 0xde, 0x47,
	0xa6, 0xd0, 0x61, 0x55, 0x6a, 0x3b, 0xea, 0x25, 0x3a, 0x03, 0xd7, 0x4f, 0x38, 0xa6, 0x6a, 0x1d,
	0x25, 0x7c, 0x1b, 0x29, 0x9e, 0xa5, 0x64, 0x08, 0xbf, 0xf9, 0x16, 0x53, 0x65, 0xdf, 0x97, 0x09,
	0x1d, 0x40, 0x7f, 0x1d, 0xe5, 0x89, 0x7a, 0xcd, 0x76, 0x98, 0xda, 0x8d, 0x28, 0x05, 0xa8, 0x8a,
	0xba, 0x51, 0xe9, 0xe0, 0xd8, 0x68, 0x12, 0xfa, 0xe5, 0x80, 0xf7, 0x26, 0x51, 0xf8, 0x02, 0xb5,
	0x10, 0x8f, 0x92, 0xa3, 0x1f, 0x1f, 0xfa, 0xcc, 0x0c, 0x0f, 0x8b, 0xd3, 0x74, 0xd3, 0xdb, 0x59,
	0xfc, 0x9d, 0x1b, 0x2c, 0x97, 0xbb, 0x05, 0x2e, 0xbb, 0xdc, 0x76, 0x05, 0x23, 0x29, 0xe3, 0xb0,
	0x66, 0x2a, 0x14, 0xa5, 0xfe, 0xf8, 0xa7, 0x91, 0x9a, 0x94, 0x52, 0x57, 0x91, 0x06, 0x9e, 0xbc,
	0x4a, 0xfa, 0x19, 0x06, 0x85, 0xf6, 0x15, 0x1a, 0x0b, 0x27, 0xc1, 0x96, 0x11, 0x1c, 0x95, 0x82,
	0x0d, 0x1a, 0x41, 0xbf, 0x68, 0x00, 0xba, 0x01, 0xd0, 0xde, 0x97, 0xcc, 0xec, 0xea, 0x42, 0x2b,
	0x17, 0x89, 0xc5, 0xa3, 0x43, 0xfa, 0xe9, 0x00, 0xa9, 0x83, 0x91, 0xfb, 0x2c, 0x95, 0x48, 0x1e,
	0xa1, 0xa7, 0x67, 0x72, 0x81, 0xdb, 0x30, 0x62, 0x35, 0x2e, 0x6e, 0x39, 0xbb, 0xd2, 0x0c, 0xba,
	0xc7, 0x87, 0x76, 0xc6, 0x13, 0xf4, 0x2e, 0x78, 0x58, 0x0e, 0xc3, 0xab, 0x1c, 0xba, 0xe7, 0x00,
	0xc8, 0x1d, 0x74, 0x6a, 0xce, 0xad, 0x63, 0xb7, 0xe1, 0x18, 0x2a, 0xab, 0x8b, 0x1d, 0x78, 0xcb,
	0x5c, 0xc5, 0xda, 0x02, 0x33, 0x37, 0x59, 0xa1, 0x28, 0x38, 0x43, 0x12, 0x80, 0x67, 0x39, 0x9c,
	0xdf, 0x9f, 0x4c, 0x2a, 0x17, 0x8d, 0xaf, 0xe2, 0xdf, 0xd8, 0x9e, 0xbe, 0x41, 0x85, 0xfe, 0xb8,
	0x75, 0x36, 0x7f, 0xcc, 0x9f, 0x72, 0xff, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x8d, 0xae, 0x47, 0x03,
	0x37, 0x03, 0x00, 0x00,
}
