// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/resources/shared_set.proto

package resources // import "google.golang.org/genproto/googleapis/ads/googleads/v0/resources"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import enums "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// SharedSets are used for sharing criterion exclusions across multiple
// campaigns.
type SharedSet struct {
	// The resource name of the shared set.
	// Shared set resource names have the form:
	//
	// `customers/{customer_id}/sharedSets/{shared_set_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of this shared set. Read only.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// The type of this shared set: each shared set holds only a single kind
	// of entity. Required. Immutable.
	Type enums.SharedSetTypeEnum_SharedSetType `protobuf:"varint,3,opt,name=type,proto3,enum=google.ads.googleads.v0.enums.SharedSetTypeEnum_SharedSetType" json:"type,omitempty"`
	// The name of this shared set. Required.
	// Shared Sets must have names that are unique among active shared sets of
	// the same type.
	// The length of this string should be between 1 and 255 UTF-8 bytes,
	// inclusive.
	Name *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// The status of this shared set. Read only.
	Status enums.SharedSetStatusEnum_SharedSetStatus `protobuf:"varint,5,opt,name=status,proto3,enum=google.ads.googleads.v0.enums.SharedSetStatusEnum_SharedSetStatus" json:"status,omitempty"`
	// The number of shared criteria within this shared set. Read only.
	MemberCount *wrappers.Int64Value `protobuf:"bytes,6,opt,name=member_count,json=memberCount,proto3" json:"member_count,omitempty"`
	// The number of campaigns associated with this shared set. Read only.
	ReferenceCount       *wrappers.Int64Value `protobuf:"bytes,7,opt,name=reference_count,json=referenceCount,proto3" json:"reference_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SharedSet) Reset()         { *m = SharedSet{} }
func (m *SharedSet) String() string { return proto.CompactTextString(m) }
func (*SharedSet) ProtoMessage()    {}
func (*SharedSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_shared_set_b52601ff8e18cc2a, []int{0}
}
func (m *SharedSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SharedSet.Unmarshal(m, b)
}
func (m *SharedSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SharedSet.Marshal(b, m, deterministic)
}
func (dst *SharedSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SharedSet.Merge(dst, src)
}
func (m *SharedSet) XXX_Size() int {
	return xxx_messageInfo_SharedSet.Size(m)
}
func (m *SharedSet) XXX_DiscardUnknown() {
	xxx_messageInfo_SharedSet.DiscardUnknown(m)
}

var xxx_messageInfo_SharedSet proto.InternalMessageInfo

func (m *SharedSet) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *SharedSet) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *SharedSet) GetType() enums.SharedSetTypeEnum_SharedSetType {
	if m != nil {
		return m.Type
	}
	return enums.SharedSetTypeEnum_UNSPECIFIED
}

func (m *SharedSet) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *SharedSet) GetStatus() enums.SharedSetStatusEnum_SharedSetStatus {
	if m != nil {
		return m.Status
	}
	return enums.SharedSetStatusEnum_UNSPECIFIED
}

func (m *SharedSet) GetMemberCount() *wrappers.Int64Value {
	if m != nil {
		return m.MemberCount
	}
	return nil
}

func (m *SharedSet) GetReferenceCount() *wrappers.Int64Value {
	if m != nil {
		return m.ReferenceCount
	}
	return nil
}

func init() {
	proto.RegisterType((*SharedSet)(nil), "google.ads.googleads.v0.resources.SharedSet")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/resources/shared_set.proto", fileDescriptor_shared_set_b52601ff8e18cc2a)
}

var fileDescriptor_shared_set_b52601ff8e18cc2a = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xc1, 0x6a, 0xdb, 0x30,
	0x18, 0xc7, 0xb1, 0x93, 0x65, 0x54, 0xed, 0x32, 0xd0, 0xc9, 0x74, 0x63, 0xa4, 0x1b, 0x85, 0xc0,
	0x40, 0x36, 0xe9, 0xb6, 0x83, 0x06, 0x05, 0x67, 0x1b, 0x65, 0x3b, 0x8c, 0xe2, 0x8c, 0x1c, 0x4a,
	0x20, 0x28, 0xd6, 0x57, 0xcf, 0x10, 0x4b, 0x46, 0x92, 0x3b, 0xfa, 0x3a, 0x3b, 0xee, 0x3d, 0x76,
	0xd9, 0xa3, 0xec, 0x15, 0x76, 0x19, 0x96, 0x6c, 0xb1, 0x32, 0xda, 0xf4, 0xf6, 0x59, 0xfe, 0xff,
	0x7f, 0xfe, 0x49, 0xc8, 0x68, 0x56, 0x48, 0x59, 0x6c, 0x21, 0x66, 0x5c, 0xc7, 0x6e, 0x6c, 0xa7,
	0xab, 0x24, 0x56, 0xa0, 0x65, 0xa3, 0x72, 0xd0, 0xb1, 0xfe, 0xca, 0x14, 0xf0, 0xb5, 0x06, 0x43,
	0x6a, 0x25, 0x8d, 0xc4, 0x47, 0x2e, 0x48, 0x18, 0xd7, 0xc4, 0x77, 0xc8, 0x55, 0x42, 0x7c, 0xe7,
	0xf0, 0xf5, 0x6d, 0x58, 0x10, 0x4d, 0xf5, 0x2f, 0x72, 0xad, 0x0d, 0x33, 0x8d, 0x76, 0xe4, 0xc3,
	0x93, 0x7b, 0xd7, 0xcc, 0x75, 0x0d, 0x5d, 0xe9, 0x59, 0x57, 0xb2, 0x4f, 0x9b, 0xe6, 0x32, 0xfe,
	0xa6, 0x58, 0x5d, 0x83, 0xea, 0xa0, 0xcf, 0x7f, 0x0e, 0xd0, 0xde, 0xc2, 0x36, 0x17, 0x60, 0xf0,
	0x0b, 0xf4, 0xa8, 0xd7, 0x5c, 0x0b, 0x56, 0x41, 0x14, 0x4c, 0x82, 0xe9, 0x5e, 0x76, 0xd0, 0x2f,
	0x7e, 0x66, 0x15, 0xe0, 0x97, 0x28, 0x2c, 0x79, 0x14, 0x4e, 0x82, 0xe9, 0xfe, 0xec, 0x49, 0xb7,
	0x47, 0xd2, 0xf3, 0xc9, 0x47, 0x61, 0xde, 0xbc, 0x5a, 0xb2, 0x6d, 0x03, 0x59, 0x58, 0x72, 0x9c,
	0xa1, 0x61, 0x6b, 0x13, 0x0d, 0x26, 0xc1, 0x74, 0x3c, 0x3b, 0x25, 0xb7, 0x9d, 0x8e, 0xdd, 0x03,
	0xf1, 0x26, 0x5f, 0xae, 0x6b, 0xf8, 0x20, 0x9a, 0xea, 0xe6, 0x4a, 0x66, 0x59, 0x38, 0x41, 0x43,
	0x2b, 0x37, 0xb4, 0x0a, 0x4f, 0xff, 0x53, 0x58, 0x18, 0x55, 0x8a, 0xc2, 0x39, 0xd8, 0x24, 0xbe,
	0x40, 0x23, 0x77, 0x94, 0xd1, 0x03, 0xeb, 0x31, 0xbf, 0xaf, 0xc7, 0xc2, 0xb6, 0x6e, 0x9a, 0xb8,
	0xb5, 0xac, 0x23, 0xe2, 0x53, 0x74, 0x50, 0x41, 0xb5, 0x01, 0xb5, 0xce, 0x65, 0x23, 0x4c, 0x34,
	0xda, 0x7d, 0x30, 0xfb, 0xae, 0xf0, 0xae, 0xcd, 0xe3, 0xf7, 0xe8, 0xb1, 0x82, 0x4b, 0x50, 0x20,
	0x72, 0xe8, 0x10, 0x0f, 0x77, 0x23, 0xc6, 0xbe, 0x63, 0x29, 0xf3, 0x3f, 0x01, 0x3a, 0xce, 0x65,
	0x45, 0x76, 0xde, 0xbe, 0xf9, 0xd8, 0x6f, 0xe4, 0xbc, 0xe5, 0x9e, 0x07, 0x17, 0x9f, 0xba, 0x52,
	0x21, 0xb7, 0x4c, 0x14, 0x44, 0xaa, 0x22, 0x2e, 0x40, 0xd8, 0xaf, 0xf6, 0x17, 0xad, 0x2e, 0xf5,
	0x1d, 0x7f, 0xc1, 0x5b, 0x3f, 0x7d, 0x0f, 0x07, 0x67, 0x69, 0xfa, 0x23, 0x3c, 0x3a, 0x73, 0xc8,
	0x94, 0x6b, 0xe2, 0xc6, 0x76, 0x5a, 0x26, 0x24, 0xeb, 0x93, 0xbf, 0xfa, 0xcc, 0x2a, 0xe5, 0x7a,
	0xe5, 0x33, 0xab, 0x65, 0xb2, 0xf2, 0x99, 0xdf, 0xe1, 0xb1, 0x7b, 0x41, 0x69, 0xca, 0x35, 0xa5,
	0x3e, 0x45, 0xe9, 0x32, 0xa1, 0xd4, 0xe7, 0x36, 0x23, 0x2b, 0x7b, 0xf2, 0x37, 0x00, 0x00, 0xff,
	0xff, 0xd7, 0xad, 0xe0, 0xa5, 0xb1, 0x03, 0x00, 0x00,
}