// Code generated by protoc-gen-go. DO NOT EDIT.
// source: runtime.proto

package config

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Runtime Configuration for a WASM OCI Image
type Runtime struct {
	// the type of the runtime
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// the version of the ABI of the target runtime
	// this may be different than the version of the runtime itself
	// this is used to ensure compatibility with the runtime
	AbiVersion string `protobuf:"bytes,2,opt,name=abi_version,json=abiVersion,proto3" json:"abi_version,omitempty"`
	// the config for running the module
	// currently, wasme only supports Envoy config
	Config               *EnvoyConfig `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Runtime) Reset()         { *m = Runtime{} }
func (m *Runtime) String() string { return proto.CompactTextString(m) }
func (*Runtime) ProtoMessage()    {}
func (*Runtime) Descriptor() ([]byte, []int) {
	return fileDescriptor_86e2dd377c869464, []int{0}
}

func (m *Runtime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Runtime.Unmarshal(m, b)
}
func (m *Runtime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Runtime.Marshal(b, m, deterministic)
}
func (m *Runtime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Runtime.Merge(m, src)
}
func (m *Runtime) XXX_Size() int {
	return xxx_messageInfo_Runtime.Size(m)
}
func (m *Runtime) XXX_DiscardUnknown() {
	xxx_messageInfo_Runtime.DiscardUnknown(m)
}

var xxx_messageInfo_Runtime proto.InternalMessageInfo

func (m *Runtime) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Runtime) GetAbiVersion() string {
	if m != nil {
		return m.AbiVersion
	}
	return ""
}

func (m *Runtime) GetConfig() *EnvoyConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

// configuration for an Envoy Filter WASM Image
type EnvoyConfig struct {
	// the set of root IDs exposed by the Envoy Filter
	RootIds              []string `protobuf:"bytes,1,rep,name=root_ids,json=rootIds,proto3" json:"root_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnvoyConfig) Reset()         { *m = EnvoyConfig{} }
func (m *EnvoyConfig) String() string { return proto.CompactTextString(m) }
func (*EnvoyConfig) ProtoMessage()    {}
func (*EnvoyConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_86e2dd377c869464, []int{1}
}

func (m *EnvoyConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnvoyConfig.Unmarshal(m, b)
}
func (m *EnvoyConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnvoyConfig.Marshal(b, m, deterministic)
}
func (m *EnvoyConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnvoyConfig.Merge(m, src)
}
func (m *EnvoyConfig) XXX_Size() int {
	return xxx_messageInfo_EnvoyConfig.Size(m)
}
func (m *EnvoyConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_EnvoyConfig.DiscardUnknown(m)
}

var xxx_messageInfo_EnvoyConfig proto.InternalMessageInfo

func (m *EnvoyConfig) GetRootIds() []string {
	if m != nil {
		return m.RootIds
	}
	return nil
}

func init() {
	proto.RegisterType((*Runtime)(nil), "solo.extendenvoy.config.Runtime")
	proto.RegisterType((*EnvoyConfig)(nil), "solo.extendenvoy.config.EnvoyConfig")
}

func init() { proto.RegisterFile("runtime.proto", fileDescriptor_86e2dd377c869464) }

var fileDescriptor_86e2dd377c869464 = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x2a, 0xcd, 0x2b,
	0xc9, 0xcc, 0x4d, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x2f, 0xce, 0xcf, 0xc9, 0xd7,
	0x4b, 0xad, 0x28, 0x49, 0xcd, 0x4b, 0x49, 0xcd, 0x2b, 0xcb, 0xaf, 0xd4, 0x4b, 0xce, 0xcf, 0x4b,
	0xcb, 0x4c, 0x57, 0xaa, 0xe1, 0x62, 0x0f, 0x82, 0xa8, 0x14, 0x12, 0xe2, 0x62, 0x29, 0xa9, 0x2c,
	0x48, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0xe4, 0xb9, 0xb8, 0x13, 0x93,
	0x32, 0xe3, 0xcb, 0x52, 0x8b, 0x8a, 0x33, 0xf3, 0xf3, 0x24, 0x98, 0xc0, 0x52, 0x5c, 0x89, 0x49,
	0x99, 0x61, 0x10, 0x11, 0x21, 0x1b, 0x2e, 0x36, 0x88, 0x49, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0xdc,
	0x46, 0x2a, 0x7a, 0x38, 0x6c, 0xd2, 0x73, 0x05, 0x71, 0x9c, 0xc1, 0xec, 0x20, 0xa8, 0x1e, 0x25,
	0x0d, 0x2e, 0x6e, 0x24, 0x61, 0x21, 0x49, 0x2e, 0x8e, 0xa2, 0xfc, 0xfc, 0x92, 0xf8, 0xcc, 0x94,
	0x62, 0x09, 0x46, 0x05, 0x66, 0x0d, 0xce, 0x20, 0x76, 0x10, 0xdf, 0x33, 0xa5, 0xd8, 0x89, 0x23,
	0x0a, 0xaa, 0x27, 0x89, 0x0d, 0xec, 0x23, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x9a,
	0xb3, 0x85, 0xe2, 0x00, 0x00, 0x00,
}