// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/http/squash/v3/squash.proto

package envoy_extensions_filters_http_squash_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	_struct "github.com/golang/protobuf/ptypes/struct"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Squash struct {
	Cluster              string             `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	AttachmentTemplate   *_struct.Struct    `protobuf:"bytes,2,opt,name=attachment_template,json=attachmentTemplate,proto3" json:"attachment_template,omitempty"`
	RequestTimeout       *duration.Duration `protobuf:"bytes,3,opt,name=request_timeout,json=requestTimeout,proto3" json:"request_timeout,omitempty"`
	AttachmentTimeout    *duration.Duration `protobuf:"bytes,4,opt,name=attachment_timeout,json=attachmentTimeout,proto3" json:"attachment_timeout,omitempty"`
	AttachmentPollPeriod *duration.Duration `protobuf:"bytes,5,opt,name=attachment_poll_period,json=attachmentPollPeriod,proto3" json:"attachment_poll_period,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Squash) Reset()         { *m = Squash{} }
func (m *Squash) String() string { return proto.CompactTextString(m) }
func (*Squash) ProtoMessage()    {}
func (*Squash) Descriptor() ([]byte, []int) {
	return fileDescriptor_2dba825384a5c607, []int{0}
}

func (m *Squash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Squash.Unmarshal(m, b)
}
func (m *Squash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Squash.Marshal(b, m, deterministic)
}
func (m *Squash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Squash.Merge(m, src)
}
func (m *Squash) XXX_Size() int {
	return xxx_messageInfo_Squash.Size(m)
}
func (m *Squash) XXX_DiscardUnknown() {
	xxx_messageInfo_Squash.DiscardUnknown(m)
}

var xxx_messageInfo_Squash proto.InternalMessageInfo

func (m *Squash) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *Squash) GetAttachmentTemplate() *_struct.Struct {
	if m != nil {
		return m.AttachmentTemplate
	}
	return nil
}

func (m *Squash) GetRequestTimeout() *duration.Duration {
	if m != nil {
		return m.RequestTimeout
	}
	return nil
}

func (m *Squash) GetAttachmentTimeout() *duration.Duration {
	if m != nil {
		return m.AttachmentTimeout
	}
	return nil
}

func (m *Squash) GetAttachmentPollPeriod() *duration.Duration {
	if m != nil {
		return m.AttachmentPollPeriod
	}
	return nil
}

func init() {
	proto.RegisterType((*Squash)(nil), "envoy.extensions.filters.http.squash.v3.Squash")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/http/squash/v3/squash.proto", fileDescriptor_2dba825384a5c607)
}

var fileDescriptor_2dba825384a5c607 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x3f, 0x0f, 0x12, 0x31,
	0x18, 0xc6, 0x73, 0x07, 0x42, 0x2c, 0x89, 0x7f, 0x4e, 0x23, 0x27, 0x51, 0x02, 0x2e, 0x62, 0x4c,
	0x5a, 0x03, 0xba, 0x38, 0x5e, 0x1c, 0x98, 0xf4, 0x02, 0xec, 0xa4, 0xdc, 0x95, 0xa3, 0x49, 0x69,
	0x8f, 0xf6, 0xed, 0x05, 0x36, 0x27, 0xe3, 0x67, 0xf0, 0xa3, 0xb8, 0x9b, 0xb8, 0xfa, 0x75, 0x9c,
	0x0c, 0xbd, 0x12, 0x50, 0x06, 0xdc, 0x9a, 0x3e, 0xfd, 0xfd, 0xf2, 0xbc, 0xcd, 0x8b, 0xde, 0x32,
	0x59, 0xa9, 0x03, 0x61, 0x7b, 0x60, 0xd2, 0x70, 0x25, 0x0d, 0x59, 0x73, 0x01, 0x4c, 0x1b, 0xb2,
	0x01, 0x28, 0x89, 0xd9, 0x59, 0x6a, 0x36, 0xa4, 0x9a, 0xf8, 0x13, 0x2e, 0xb5, 0x02, 0x15, 0xbd,
	0x74, 0x14, 0x3e, 0x53, 0xd8, 0x53, 0xf8, 0x48, 0x61, 0xff, 0xb6, 0x9a, 0xf4, 0xfa, 0x85, 0x52,
	0x85, 0x60, 0xc4, 0x61, 0x2b, 0xbb, 0x26, 0xb9, 0xd5, 0x14, 0xb8, 0x92, 0xb5, 0xa8, 0xf7, 0xec,
	0xdf, 0xdc, 0x80, 0xb6, 0x19, 0xf8, 0xf4, 0xb9, 0xcd, 0x4b, 0x4a, 0xa8, 0x94, 0x0a, 0x1c, 0x64,
	0x88, 0x01, 0x0a, 0xd6, 0xf8, 0x78, 0x78, 0x15, 0x57, 0x4c, 0x1f, 0xeb, 0x70, 0x59, 0xf8, 0x27,
	0xdd, 0x8a, 0x0a, 0x9e, 0x53, 0x60, 0xe4, 0x74, 0xa8, 0x83, 0x17, 0x5f, 0x1a, 0xa8, 0x35, 0x77,
	0x35, 0xa3, 0x21, 0x6a, 0x67, 0xc2, 0x1a, 0x60, 0x3a, 0x0e, 0x06, 0xc1, 0xe8, 0x6e, 0xd2, 0xfe,
	0x9d, 0x34, 0x75, 0x38, 0x08, 0x66, 0xa7, 0xfb, 0x68, 0x8a, 0x1e, 0x51, 0x00, 0x9a, 0x6d, 0xb6,
	0x4c, 0xc2, 0x12, 0xd8, 0xb6, 0x14, 0x14, 0x58, 0x1c, 0x0e, 0x82, 0x51, 0x67, 0xdc, 0xc5, 0xf5,
	0x10, 0xf8, 0x34, 0x04, 0x9e, 0xbb, 0x21, 0x66, 0xd1, 0x99, 0x59, 0x78, 0x24, 0x4a, 0xd0, 0x7d,
	0xcd, 0x76, 0x96, 0x19, 0x58, 0x02, 0xdf, 0x32, 0x65, 0x21, 0x6e, 0x38, 0xcb, 0xd3, 0x2b, 0xcb,
	0x07, 0xff, 0x55, 0xb3, 0x7b, 0x9e, 0x58, 0xd4, 0x40, 0x34, 0x45, 0xd1, 0x65, 0x1b, 0xaf, 0x69,
	0xde, 0xd2, 0x3c, 0xbc, 0xa8, 0xe3, 0x4d, 0x9f, 0xd0, 0x93, 0x0b, 0x53, 0xa9, 0x84, 0x58, 0x96,
	0x4c, 0x73, 0x95, 0xc7, 0x77, 0x6e, 0xd9, 0x1e, 0x9f, 0xc1, 0x54, 0x09, 0x91, 0x3a, 0xec, 0xfd,
	0x9b, 0x6f, 0x3f, 0xbe, 0xf6, 0x5f, 0xa3, 0x57, 0xf5, 0x7e, 0x64, 0x4a, 0xae, 0x79, 0xe1, 0x77,
	0xe3, 0xef, 0xd5, 0x18, 0xe3, 0xfa, 0xf7, 0x93, 0x8f, 0xdf, 0x3f, 0xff, 0xfc, 0xd5, 0x0a, 0x1f,
	0x84, 0xe8, 0x1d, 0x57, 0xd8, 0x71, 0xa5, 0x56, 0xfb, 0x03, 0xfe, 0xcf, 0x15, 0x4b, 0x3a, 0xb5,
	0x28, 0x3d, 0x36, 0x4c, 0x83, 0x55, 0xcb, 0x55, 0x9d, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0xb8,
	0xbf, 0x42, 0xbe, 0xd9, 0x02, 0x00, 0x00,
}
