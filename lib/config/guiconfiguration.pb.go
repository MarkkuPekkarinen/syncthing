// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lib/config/guiconfiguration.proto

package config

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/syncthing/syncthing/proto/ext"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type GUIConfiguration struct {
	Enabled                   bool     `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled" xml:"enabled,attr" default:"true"`
	RawAddress                string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address" xml:"address" default:"127.0.0.1:8384"`
	RawUnixSocketPermissions  string   `protobuf:"bytes,3,opt,name=unix_socket_permissions,json=unixSocketPermissions,proto3" json:"unixSocketPermissions" xml:"unixSocketPermissions,omitempty"`
	User                      string   `protobuf:"bytes,4,opt,name=user,proto3" json:"user" xml:"user,omitempty"`
	Password                  string   `protobuf:"bytes,5,opt,name=password,proto3" json:"password" xml:"password,omitempty"`
	AuthMode                  AuthMode `protobuf:"varint,6,opt,name=auth_mode,json=authMode,proto3,enum=config.AuthMode" json:"authMode" xml:"authMode,omitempty"`
	RawUseTLS                 bool     `protobuf:"varint,7,opt,name=use_tls,json=useTls,proto3" json:"useTLS" xml:"tls,attr"`
	APIKey                    string   `protobuf:"bytes,8,opt,name=api_key,json=apiKey,proto3" json:"apiKey" xml:"apikey,omitempty"`
	InsecureAdminAccess       bool     `protobuf:"varint,9,opt,name=insecure_admin_access,json=insecureAdminAccess,proto3" json:"insecureAdminAccess" xml:"insecureAdminAccess,omitempty"`
	Theme                     string   `protobuf:"bytes,10,opt,name=theme,proto3" json:"theme" xml:"theme" default:"default"`
	Debugging                 bool     `protobuf:"varint,11,opt,name=debugging,proto3" json:"debugging" xml:"debugging,attr"`
	InsecureSkipHostCheck     bool     `protobuf:"varint,12,opt,name=insecure_skip_host_check,json=insecureSkipHostCheck,proto3" json:"insecureSkipHostcheck" xml:"insecureSkipHostcheck,omitempty"`
	InsecureAllowFrameLoading bool     `protobuf:"varint,13,opt,name=insecure_allow_frame_loading,json=insecureAllowFrameLoading,proto3" json:"insecureAllowFrameLoading" xml:"insecureAllowFrameLoading,omitempty"`
}

func (m *GUIConfiguration) Reset()         { *m = GUIConfiguration{} }
func (m *GUIConfiguration) String() string { return proto.CompactTextString(m) }
func (*GUIConfiguration) ProtoMessage()    {}
func (*GUIConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a9586d611855d64, []int{0}
}
func (m *GUIConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GUIConfiguration.Unmarshal(m, b)
}
func (m *GUIConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GUIConfiguration.Marshal(b, m, deterministic)
}
func (m *GUIConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GUIConfiguration.Merge(m, src)
}
func (m *GUIConfiguration) XXX_Size() int {
	return xxx_messageInfo_GUIConfiguration.Size(m)
}
func (m *GUIConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_GUIConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_GUIConfiguration proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GUIConfiguration)(nil), "config.GUIConfiguration")
}

func init() { proto.RegisterFile("lib/config/guiconfiguration.proto", fileDescriptor_2a9586d611855d64) }

var fileDescriptor_2a9586d611855d64 = []byte{
	// 834 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x16, 0x5b, 0x47, 0xb2, 0xb6, 0xae, 0x60, 0xb0, 0x4d, 0xcb, 0x04, 0x8d, 0x56, 0x51, 0xd8,
	0xc2, 0x01, 0x02, 0x39, 0x71, 0x5a, 0x24, 0xf0, 0xa1, 0x80, 0x1c, 0x20, 0x3f, 0xb0, 0x0b, 0x04,
	0x74, 0x7d, 0xc9, 0x85, 0x58, 0x91, 0x6b, 0x69, 0x21, 0xfe, 0x95, 0xbb, 0x84, 0xa4, 0x43, 0xfb,
	0x0c, 0x85, 0x7a, 0x2e, 0xd0, 0x67, 0xe8, 0xa5, 0xaf, 0xd0, 0x9b, 0x74, 0x2a, 0x72, 0x5a, 0x20,
	0xd2, 0xa5, 0xe0, 0x91, 0xc7, 0x9c, 0x8a, 0x5d, 0xfe, 0x48, 0x94, 0x95, 0x26, 0xb7, 0x9d, 0x6f,
	0xbe, 0x99, 0x6f, 0x66, 0x38, 0x03, 0x82, 0xdb, 0x0e, 0xe9, 0x1d, 0x5a, 0xbe, 0x77, 0x49, 0xfa,
	0x87, 0xfd, 0x88, 0xa4, 0xaf, 0x28, 0x44, 0x8c, 0xf8, 0x5e, 0x27, 0x08, 0x7d, 0xe6, 0xab, 0xd5,
	0x14, 0xbc, 0x79, 0x63, 0x8d, 0x8a, 0x22, 0x36, 0x70, 0x7d, 0x1b, 0xa7, 0x94, 0x9b, 0x75, 0x3c,
	0x66, 0xe9, 0xb3, 0xfd, 0x7a, 0x0f, 0xec, 0x3f, 0xbb, 0x78, 0xf1, 0x64, 0x3d, 0x91, 0xda, 0x03,
	0x35, 0xec, 0xa1, 0x9e, 0x83, 0x6d, 0x4d, 0x69, 0x29, 0x07, 0xbb, 0x27, 0xcf, 0x63, 0x0e, 0x73,
	0x28, 0xe1, 0xf0, 0xf6, 0xd8, 0x75, 0x8e, 0xdb, 0x99, 0x7d, 0x0f, 0x31, 0x16, 0xb6, 0x5b, 0x36,
	0xbe, 0x44, 0x91, 0xc3, 0x8e, 0xdb, 0x2c, 0x8c, 0x70, 0x3b, 0x9e, 0xe9, 0x7b, 0xeb, 0xfe, 0xb7,
	0x33, 0x7d, 0x47, 0x38, 0x8c, 0x3c, 0x8b, 0xfa, 0x33, 0xa8, 0x21, 0xdb, 0x0e, 0x31, 0xa5, 0xda,
	0x47, 0x2d, 0xe5, 0xa0, 0x7e, 0x62, 0x2d, 0x38, 0x04, 0x06, 0x1a, 0x75, 0x53, 0x54, 0x28, 0x66,
	0x84, 0x84, 0xc3, 0x6f, 0xa4, 0x62, 0x66, 0xaf, 0x89, 0x3d, 0x38, 0x7a, 0xd4, 0xb9, 0xdf, 0xb9,
	0xdf, 0x79, 0x70, 0xfc, 0xf8, 0xe1, 0xe3, 0x6f, 0xdb, 0x6f, 0x67, 0x7a, 0xa3, 0x0c, 0x4d, 0xe7,
	0xfa, 0x5a, 0x52, 0x23, 0x4f, 0xa9, 0xfe, 0xa3, 0x80, 0x2f, 0x23, 0x8f, 0x8c, 0x4d, 0xea, 0x5b,
	0x43, 0xcc, 0xcc, 0x00, 0x87, 0x2e, 0xa1, 0x94, 0xf8, 0x1e, 0xd5, 0x3e, 0x96, 0xf5, 0xfc, 0xae,
	0x2c, 0x38, 0xd4, 0x0c, 0x34, 0xba, 0xf0, 0xc8, 0xf8, 0x5c, 0xb2, 0x5e, 0xae, 0x48, 0x31, 0x87,
	0xd7, 0xa3, 0x6d, 0x8e, 0x84, 0xc3, 0xaf, 0x65, 0xb1, 0x5b, 0xbd, 0xf7, 0x7c, 0x97, 0x30, 0xec,
	0x06, 0x6c, 0x22, 0x46, 0x04, 0xdf, 0xc3, 0x99, 0xce, 0xf5, 0x77, 0x16, 0x60, 0x6c, 0x97, 0x57,
	0x9f, 0x82, 0x9d, 0x88, 0xe2, 0x50, 0xdb, 0x91, 0x4d, 0x1c, 0xc5, 0x1c, 0x4a, 0x3b, 0xe1, 0xf0,
	0xf3, 0xb4, 0x2c, 0x8a, 0xc3, 0x72, 0x15, 0x8d, 0x32, 0x64, 0x48, 0xbe, 0xfa, 0x0a, 0xec, 0x06,
	0x88, 0xd2, 0x91, 0x1f, 0xda, 0xda, 0x35, 0x99, 0xeb, 0xfb, 0x98, 0xc3, 0x02, 0x4b, 0x38, 0xd4,
	0x64, 0xbe, 0x1c, 0x28, 0xe7, 0x54, 0xaf, 0xc2, 0x46, 0x11, 0xab, 0xba, 0xa0, 0x2e, 0x36, 0xd2,
	0x14, 0x2b, 0xa9, 0x55, 0x5b, 0xca, 0x41, 0xe3, 0x68, 0xbf, 0x93, 0xae, 0x6a, 0xa7, 0x1b, 0xb1,
	0xc1, 0x0f, 0xbe, 0x8d, 0x53, 0x39, 0x94, 0x59, 0x85, 0x5c, 0x0e, 0x6c, 0xc8, 0x5d, 0x85, 0x8d,
	0x22, 0x56, 0xc5, 0xa0, 0x16, 0x51, 0x6c, 0x32, 0x87, 0x6a, 0x35, 0xb9, 0xce, 0x67, 0x0b, 0x0e,
	0xeb, 0x62, 0xb0, 0x14, 0xff, 0x78, 0x76, 0x1e, 0x73, 0x58, 0x8d, 0xe4, 0x2b, 0xe1, 0xb0, 0x21,
	0x55, 0x98, 0x43, 0xd3, 0xb5, 0x8e, 0x67, 0xfa, 0x6e, 0x6e, 0x24, 0x33, 0x3d, 0xe3, 0x4d, 0xe7,
	0xfa, 0x2a, 0xdc, 0x90, 0xa0, 0x43, 0x85, 0x0c, 0x0a, 0x88, 0x39, 0xc4, 0x13, 0x6d, 0x57, 0x0e,
	0x4c, 0xc8, 0x54, 0xbb, 0x2f, 0x5f, 0x9c, 0xe2, 0x89, 0xd0, 0x40, 0x01, 0x39, 0xc5, 0x93, 0x84,
	0xc3, 0x2f, 0xd2, 0x4e, 0x02, 0x32, 0xc4, 0x93, 0x72, 0x1f, 0xfb, 0x9b, 0xe0, 0x74, 0xae, 0x67,
	0x19, 0x8c, 0x2c, 0x5e, 0xfd, 0x4d, 0x01, 0xd7, 0x89, 0x47, 0xb1, 0x15, 0x85, 0xd8, 0x44, 0xb6,
	0x4b, 0x3c, 0x13, 0x59, 0x96, 0xb8, 0xa3, 0xba, 0x6c, 0xce, 0x8c, 0x39, 0xfc, 0x2c, 0x27, 0x74,
	0x85, 0xbf, 0x2b, 0xdd, 0x09, 0x87, 0x77, 0xa4, 0xf0, 0x16, 0x5f, 0xb9, 0x8a, 0x5b, 0xff, 0xcb,
	0x30, 0xb6, 0x25, 0x57, 0x4f, 0xc1, 0x35, 0x36, 0xc0, 0x2e, 0xd6, 0x80, 0x6c, 0xfd, 0xbb, 0x98,
	0xc3, 0x14, 0x48, 0x38, 0xbc, 0x95, 0xce, 0x54, 0x58, 0x6b, 0xa7, 0x9b, 0x3d, 0xc4, 0xcd, 0xd6,
	0xb2, 0xb7, 0x91, 0x86, 0xa8, 0x17, 0xa0, 0x6e, 0xe3, 0x5e, 0xd4, 0xef, 0x13, 0xaf, 0xaf, 0x7d,
	0x22, 0xbb, 0x7a, 0x14, 0x73, 0xb8, 0x02, 0x8b, 0x6d, 0x2e, 0x90, 0xe2, 0x73, 0x35, 0xca, 0x90,
	0xb1, 0x0a, 0x52, 0xff, 0x52, 0x80, 0x56, 0x4c, 0x8e, 0x0e, 0x49, 0x60, 0x0e, 0x7c, 0xca, 0x4c,
	0x6b, 0x80, 0xad, 0xa1, 0xb6, 0x27, 0x65, 0x7e, 0x11, 0x77, 0x9d, 0x73, 0xce, 0x87, 0x24, 0x78,
	0xee, 0x53, 0x26, 0x09, 0xc5, 0x5d, 0x6f, 0xf5, 0x6e, 0xdc, 0xf5, 0x7b, 0x38, 0xc9, 0x4c, 0xdf,
	0x2e, 0x62, 0x5c, 0x81, 0x9f, 0x08, 0x58, 0xfd, 0x53, 0x01, 0x5f, 0xad, 0xbe, 0xb9, 0xe3, 0xf8,
	0x23, 0xf3, 0x32, 0x44, 0x2e, 0x36, 0x1d, 0x1f, 0xd9, 0x62, 0x48, 0x9f, 0xca, 0xea, 0x7f, 0x8a,
	0x39, 0xbc, 0x51, 0x7c, 0x1d, 0x41, 0x7b, 0x2a, 0x58, 0x67, 0x29, 0x29, 0xe1, 0xf0, 0x6e, 0x79,
	0x01, 0x36, 0x19, 0xe5, 0x2e, 0xee, 0x7c, 0x00, 0xcf, 0x78, 0xb7, 0xdc, 0xc9, 0xb3, 0xbf, 0xdf,
	0x34, 0x2b, 0xf3, 0x37, 0xcd, 0xca, 0xbf, 0x8b, 0x66, 0xe5, 0xd7, 0x65, 0xb3, 0xf2, 0xc7, 0xb2,
	0xa9, 0xcc, 0x97, 0xcd, 0xca, 0xeb, 0x65, 0xb3, 0xf2, 0xea, 0x6e, 0x9f, 0xb0, 0x41, 0xd4, 0xeb,
	0x58, 0xbe, 0x7b, 0x48, 0x27, 0x9e, 0xc5, 0x06, 0xc4, 0xeb, 0xaf, 0xbd, 0x56, 0x7f, 0xaf, 0x5e,
	0x55, 0xfe, 0xaa, 0x1e, 0xfe, 0x17, 0x00, 0x00, 0xff, 0xff, 0xfa, 0x2b, 0x9a, 0x81, 0xfd, 0x06,
	0x00, 0x00,
}

func (m *GUIConfiguration) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Enabled {
		n += 2
	}
	l = len(m.RawAddress)
	if l > 0 {
		n += 1 + l + sovGuiconfiguration(uint64(l))
	}
	l = len(m.RawUnixSocketPermissions)
	if l > 0 {
		n += 1 + l + sovGuiconfiguration(uint64(l))
	}
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovGuiconfiguration(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovGuiconfiguration(uint64(l))
	}
	if m.AuthMode != 0 {
		n += 1 + sovGuiconfiguration(uint64(m.AuthMode))
	}
	if m.RawUseTLS {
		n += 2
	}
	l = len(m.APIKey)
	if l > 0 {
		n += 1 + l + sovGuiconfiguration(uint64(l))
	}
	if m.InsecureAdminAccess {
		n += 2
	}
	l = len(m.Theme)
	if l > 0 {
		n += 1 + l + sovGuiconfiguration(uint64(l))
	}
	if m.Debugging {
		n += 2
	}
	if m.InsecureSkipHostCheck {
		n += 2
	}
	if m.InsecureAllowFrameLoading {
		n += 2
	}
	return n
}

func sovGuiconfiguration(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGuiconfiguration(x uint64) (n int) {
	return sovGuiconfiguration(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
