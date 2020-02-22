// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/securitycenter/v1p1beta1/asset.proto

package securitycenter

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Cloud Security Command Center's (Cloud SCC) representation of a Google Cloud
// Platform (GCP) resource.
//
// The Asset is a Cloud SCC resource that captures information about a single
// GCP resource. All modifications to an Asset are only within the context of
// Cloud SCC and don't affect the referenced GCP resource.
type Asset struct {
	// The relative resource name of this asset. See:
	// https://cloud.google.com/apis/design/resource_names#relative_resource_name
	// Example:
	// "organizations/{organization_id}/assets/{asset_id}".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Cloud SCC managed properties. These properties are managed by
	// Cloud SCC and cannot be modified by the user.
	SecurityCenterProperties *Asset_SecurityCenterProperties `protobuf:"bytes,2,opt,name=security_center_properties,json=securityCenterProperties,proto3" json:"security_center_properties,omitempty"`
	// Resource managed properties. These properties are managed and defined by
	// the GCP resource and cannot be modified by the user.
	ResourceProperties map[string]*_struct.Value `protobuf:"bytes,7,rep,name=resource_properties,json=resourceProperties,proto3" json:"resource_properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// User specified security marks. These marks are entirely managed by the user
	// and come from the SecurityMarks resource that belongs to the asset.
	SecurityMarks *SecurityMarks `protobuf:"bytes,8,opt,name=security_marks,json=securityMarks,proto3" json:"security_marks,omitempty"`
	// The time at which the asset was created in Cloud SCC.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,9,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The time at which the asset was last updated, added, or deleted in Cloud
	// SCC.
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,10,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// IAM Policy information associated with the GCP resource described by the
	// Cloud SCC asset. This information is managed and defined by the GCP
	// resource and cannot be modified by the user.
	IamPolicy            *Asset_IamPolicy `protobuf:"bytes,11,opt,name=iam_policy,json=iamPolicy,proto3" json:"iam_policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Asset) Reset()         { *m = Asset{} }
func (m *Asset) String() string { return proto.CompactTextString(m) }
func (*Asset) ProtoMessage()    {}
func (*Asset) Descriptor() ([]byte, []int) {
	return fileDescriptor_97f15e91f3a5341a, []int{0}
}

func (m *Asset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Asset.Unmarshal(m, b)
}
func (m *Asset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Asset.Marshal(b, m, deterministic)
}
func (m *Asset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Asset.Merge(m, src)
}
func (m *Asset) XXX_Size() int {
	return xxx_messageInfo_Asset.Size(m)
}
func (m *Asset) XXX_DiscardUnknown() {
	xxx_messageInfo_Asset.DiscardUnknown(m)
}

var xxx_messageInfo_Asset proto.InternalMessageInfo

func (m *Asset) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Asset) GetSecurityCenterProperties() *Asset_SecurityCenterProperties {
	if m != nil {
		return m.SecurityCenterProperties
	}
	return nil
}

func (m *Asset) GetResourceProperties() map[string]*_struct.Value {
	if m != nil {
		return m.ResourceProperties
	}
	return nil
}

func (m *Asset) GetSecurityMarks() *SecurityMarks {
	if m != nil {
		return m.SecurityMarks
	}
	return nil
}

func (m *Asset) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Asset) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *Asset) GetIamPolicy() *Asset_IamPolicy {
	if m != nil {
		return m.IamPolicy
	}
	return nil
}

// Cloud SCC managed properties. These properties are managed by Cloud SCC and
// cannot be modified by the user.
type Asset_SecurityCenterProperties struct {
	// The full resource name of the GCP resource this asset
	// represents. This field is immutable after create time. See:
	// https://cloud.google.com/apis/design/resource_names#full_resource_name
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The type of the GCP resource. Examples include: APPLICATION,
	// PROJECT, and ORGANIZATION. This is a case insensitive field defined by
	// Cloud SCC and/or the producer of the resource and is immutable
	// after create time.
	ResourceType string `protobuf:"bytes,2,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	// The full resource name of the immediate parent of the resource. See:
	// https://cloud.google.com/apis/design/resource_names#full_resource_name
	ResourceParent string `protobuf:"bytes,3,opt,name=resource_parent,json=resourceParent,proto3" json:"resource_parent,omitempty"`
	// The full resource name of the project the resource belongs to. See:
	// https://cloud.google.com/apis/design/resource_names#full_resource_name
	ResourceProject string `protobuf:"bytes,4,opt,name=resource_project,json=resourceProject,proto3" json:"resource_project,omitempty"`
	// Owners of the Google Cloud resource.
	ResourceOwners []string `protobuf:"bytes,5,rep,name=resource_owners,json=resourceOwners,proto3" json:"resource_owners,omitempty"`
	// The user defined display name for this resource.
	ResourceDisplayName string `protobuf:"bytes,6,opt,name=resource_display_name,json=resourceDisplayName,proto3" json:"resource_display_name,omitempty"`
	// The user defined display name for the parent of this resource.
	ResourceParentDisplayName string `protobuf:"bytes,7,opt,name=resource_parent_display_name,json=resourceParentDisplayName,proto3" json:"resource_parent_display_name,omitempty"`
	// The user defined display name for the project of this resource.
	ResourceProjectDisplayName string   `protobuf:"bytes,8,opt,name=resource_project_display_name,json=resourceProjectDisplayName,proto3" json:"resource_project_display_name,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *Asset_SecurityCenterProperties) Reset()         { *m = Asset_SecurityCenterProperties{} }
func (m *Asset_SecurityCenterProperties) String() string { return proto.CompactTextString(m) }
func (*Asset_SecurityCenterProperties) ProtoMessage()    {}
func (*Asset_SecurityCenterProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_97f15e91f3a5341a, []int{0, 0}
}

func (m *Asset_SecurityCenterProperties) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Asset_SecurityCenterProperties.Unmarshal(m, b)
}
func (m *Asset_SecurityCenterProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Asset_SecurityCenterProperties.Marshal(b, m, deterministic)
}
func (m *Asset_SecurityCenterProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Asset_SecurityCenterProperties.Merge(m, src)
}
func (m *Asset_SecurityCenterProperties) XXX_Size() int {
	return xxx_messageInfo_Asset_SecurityCenterProperties.Size(m)
}
func (m *Asset_SecurityCenterProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_Asset_SecurityCenterProperties.DiscardUnknown(m)
}

var xxx_messageInfo_Asset_SecurityCenterProperties proto.InternalMessageInfo

func (m *Asset_SecurityCenterProperties) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *Asset_SecurityCenterProperties) GetResourceType() string {
	if m != nil {
		return m.ResourceType
	}
	return ""
}

func (m *Asset_SecurityCenterProperties) GetResourceParent() string {
	if m != nil {
		return m.ResourceParent
	}
	return ""
}

func (m *Asset_SecurityCenterProperties) GetResourceProject() string {
	if m != nil {
		return m.ResourceProject
	}
	return ""
}

func (m *Asset_SecurityCenterProperties) GetResourceOwners() []string {
	if m != nil {
		return m.ResourceOwners
	}
	return nil
}

func (m *Asset_SecurityCenterProperties) GetResourceDisplayName() string {
	if m != nil {
		return m.ResourceDisplayName
	}
	return ""
}

func (m *Asset_SecurityCenterProperties) GetResourceParentDisplayName() string {
	if m != nil {
		return m.ResourceParentDisplayName
	}
	return ""
}

func (m *Asset_SecurityCenterProperties) GetResourceProjectDisplayName() string {
	if m != nil {
		return m.ResourceProjectDisplayName
	}
	return ""
}

// IAM Policy information associated with the GCP resource described by the
// Cloud SCC asset. This information is managed and defined by the GCP
// resource and cannot be modified by the user.
type Asset_IamPolicy struct {
	// The JSON representation of the Policy associated with the asset.
	// See https://cloud.google.com/iam/reference/rest/v1p1beta1/Policy for
	// format details.
	PolicyBlob           string   `protobuf:"bytes,1,opt,name=policy_blob,json=policyBlob,proto3" json:"policy_blob,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Asset_IamPolicy) Reset()         { *m = Asset_IamPolicy{} }
func (m *Asset_IamPolicy) String() string { return proto.CompactTextString(m) }
func (*Asset_IamPolicy) ProtoMessage()    {}
func (*Asset_IamPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_97f15e91f3a5341a, []int{0, 1}
}

func (m *Asset_IamPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Asset_IamPolicy.Unmarshal(m, b)
}
func (m *Asset_IamPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Asset_IamPolicy.Marshal(b, m, deterministic)
}
func (m *Asset_IamPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Asset_IamPolicy.Merge(m, src)
}
func (m *Asset_IamPolicy) XXX_Size() int {
	return xxx_messageInfo_Asset_IamPolicy.Size(m)
}
func (m *Asset_IamPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_Asset_IamPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_Asset_IamPolicy proto.InternalMessageInfo

func (m *Asset_IamPolicy) GetPolicyBlob() string {
	if m != nil {
		return m.PolicyBlob
	}
	return ""
}

func init() {
	proto.RegisterType((*Asset)(nil), "google.cloud.securitycenter.v1p1beta1.Asset")
	proto.RegisterMapType((map[string]*_struct.Value)(nil), "google.cloud.securitycenter.v1p1beta1.Asset.ResourcePropertiesEntry")
	proto.RegisterType((*Asset_SecurityCenterProperties)(nil), "google.cloud.securitycenter.v1p1beta1.Asset.SecurityCenterProperties")
	proto.RegisterType((*Asset_IamPolicy)(nil), "google.cloud.securitycenter.v1p1beta1.Asset.IamPolicy")
}

func init() {
	proto.RegisterFile("google/cloud/securitycenter/v1p1beta1/asset.proto", fileDescriptor_97f15e91f3a5341a)
}

var fileDescriptor_97f15e91f3a5341a = []byte{
	// 694 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdb, 0x6e, 0xd3, 0x4c,
	0x10, 0x96, 0x93, 0x9e, 0xb2, 0xf9, 0x7b, 0xd0, 0xfe, 0x02, 0x5c, 0xab, 0xa8, 0x11, 0x55, 0x45,
	0x2a, 0x2a, 0x5b, 0x29, 0x08, 0x21, 0xf7, 0x02, 0xf5, 0x24, 0xc4, 0x05, 0x10, 0xb9, 0x87, 0x0b,
	0x28, 0x8a, 0x36, 0xee, 0x62, 0x99, 0xda, 0x5e, 0x6b, 0x77, 0x5d, 0x64, 0xaa, 0x5e, 0xf1, 0x14,
	0xbc, 0x02, 0x0f, 0xc2, 0x05, 0x8f, 0xd2, 0x77, 0x40, 0x42, 0x9e, 0xb5, 0x5d, 0x3b, 0x55, 0xd4,
	0xf4, 0x2a, 0xde, 0xf9, 0x0e, 0x33, 0xb3, 0x33, 0x59, 0xd4, 0xf3, 0x18, 0xf3, 0x02, 0x6a, 0xb9,
	0x01, 0x4b, 0xce, 0x2c, 0x41, 0xdd, 0x84, 0xfb, 0x32, 0x75, 0x69, 0x24, 0x29, 0xb7, 0x2e, 0x7a,
	0x71, 0x6f, 0x48, 0x25, 0xe9, 0x59, 0x44, 0x08, 0x2a, 0xcd, 0x98, 0x33, 0xc9, 0xf0, 0xba, 0x92,
	0x98, 0x20, 0x31, 0xeb, 0x12, 0xb3, 0x94, 0x18, 0x2b, 0xb9, 0x33, 0x89, 0x7d, 0x8b, 0x44, 0x11,
	0x93, 0x44, 0xfa, 0x2c, 0x12, 0xca, 0xc4, 0x58, 0xae, 0xa0, 0x9c, 0x0a, 0x96, 0x70, 0x97, 0xe6,
	0x90, 0x3d, 0x59, 0x49, 0x05, 0x30, 0x08, 0x09, 0x3f, 0x2f, 0x6c, 0x8b, 0xa4, 0x70, 0x1a, 0x26,
	0x5f, 0x2c, 0x21, 0x79, 0xe2, 0xe6, 0x95, 0x1b, 0xab, 0xa3, 0xa8, 0xf4, 0x43, 0x2a, 0x24, 0x09,
	0x63, 0x45, 0x78, 0xf2, 0xbb, 0x85, 0xa6, 0x77, 0xb2, 0x56, 0x31, 0x46, 0x53, 0x11, 0x09, 0xa9,
	0xae, 0x75, 0xb4, 0x6e, 0xcb, 0x81, 0x6f, 0xfc, 0x43, 0x43, 0x46, 0x99, 0x55, 0xd5, 0x33, 0x88,
	0x39, 0x8b, 0x29, 0x97, 0x3e, 0x15, 0x7a, 0xa3, 0xa3, 0x75, 0xdb, 0x5b, 0x07, 0xe6, 0x44, 0xd7,
	0x63, 0x42, 0x1a, 0xf3, 0x30, 0x87, 0xf7, 0x00, 0xee, 0x97, 0x66, 0x8e, 0x2e, 0xc6, 0x20, 0x38,
	0x41, 0xff, 0x17, 0x17, 0x56, 0xcd, 0x3e, 0xdb, 0x69, 0x76, 0xdb, 0x5b, 0xfb, 0xf7, 0xca, 0xee,
	0xe4, 0x3e, 0x37, 0xee, 0x07, 0x91, 0xe4, 0xa9, 0x83, 0xf9, 0x2d, 0x00, 0x7f, 0x42, 0x0b, 0xf5,
	0x1b, 0xd7, 0xe7, 0xa0, 0xdf, 0x17, 0x13, 0x66, 0x2c, 0x3a, 0x7d, 0x97, 0x69, 0x9d, 0x79, 0x51,
	0x3d, 0xe2, 0x6d, 0xd4, 0x76, 0x39, 0x25, 0x92, 0x0e, 0xb2, 0x89, 0xe8, 0x2d, 0x70, 0x36, 0x0a,
	0xe7, 0x62, 0x5c, 0xe6, 0x51, 0x31, 0x2e, 0x07, 0x29, 0x7a, 0x16, 0xc8, 0xc4, 0x49, 0x7c, 0x56,
	0x8a, 0xd1, 0xdd, 0x62, 0x45, 0x07, 0xf1, 0x31, 0x42, 0x3e, 0x09, 0x07, 0x31, 0x0b, 0x7c, 0x37,
	0xd5, 0xdb, 0xa0, 0x7d, 0x79, 0xaf, 0x4b, 0x7c, 0x4b, 0xc2, 0x3e, 0xa8, 0x9d, 0x96, 0x5f, 0x7c,
	0x1a, 0x3f, 0x9b, 0x48, 0x1f, 0x37, 0x5b, 0xbc, 0x86, 0xe6, 0xcb, 0x09, 0x56, 0x96, 0xec, 0xbf,
	0x22, 0xf8, 0x3e, 0x5b, 0xb6, 0x2a, 0x49, 0xa6, 0x31, 0x85, 0xf5, 0xaa, 0x90, 0x8e, 0xd2, 0x98,
	0xe2, 0xa7, 0x68, 0xf1, 0x66, 0x17, 0x08, 0xa7, 0x91, 0xd4, 0x9b, 0x40, 0x5b, 0x28, 0x27, 0x08,
	0x51, 0xbc, 0x81, 0x96, 0xaa, 0x4b, 0xf3, 0x95, 0xba, 0x52, 0x9f, 0x02, 0xe6, 0x62, 0x65, 0xd6,
	0x59, 0xb8, 0xe6, 0xc9, 0xbe, 0x45, 0x94, 0x0b, 0x7d, 0xba, 0xd3, 0xac, 0x7a, 0x7e, 0x80, 0x28,
	0xde, 0x42, 0x0f, 0x4a, 0xe2, 0x99, 0x2f, 0xe2, 0x80, 0xa4, 0xaa, 0x9d, 0x19, 0x30, 0x2e, 0xb7,
	0x74, 0x5f, 0x61, 0xd0, 0xd5, 0x6b, 0xb4, 0x32, 0x52, 0x70, 0x5d, 0x3a, 0x0b, 0xd2, 0xe5, 0x7a,
	0xf5, 0x55, 0x83, 0x1d, 0xf4, 0x78, 0xb4, 0x91, 0xba, 0xc3, 0x1c, 0x38, 0x18, 0x23, 0x5d, 0x55,
	0x2c, 0x8c, 0x4d, 0xd4, 0x2a, 0x67, 0x86, 0x57, 0x51, 0x5b, 0xcd, 0x7e, 0x30, 0x0c, 0xd8, 0x30,
	0x9f, 0x04, 0x52, 0xa1, 0xdd, 0x80, 0x0d, 0x8d, 0xcf, 0xe8, 0xd1, 0x98, 0xbf, 0x09, 0x5e, 0x42,
	0xcd, 0x73, 0x9a, 0xe6, 0x9a, 0xec, 0x13, 0x6f, 0xa2, 0xe9, 0x0b, 0x12, 0x24, 0x34, 0x7f, 0x0b,
	0x1e, 0xde, 0x5a, 0xc2, 0x93, 0x0c, 0x75, 0x14, 0xc9, 0x6e, 0xbc, 0xd2, 0xec, 0xe3, 0xeb, 0x1d,
	0x07, 0xad, 0x8d, 0xac, 0x98, 0x52, 0x91, 0xd8, 0x17, 0xa6, 0xcb, 0x42, 0x4b, 0xbd, 0x48, 0xcf,
	0x18, 0xf7, 0x48, 0xe4, 0x7f, 0x57, 0xcf, 0xa8, 0x75, 0x59, 0x3d, 0x5e, 0xa9, 0x07, 0x5a, 0x58,
	0x97, 0xf0, 0x7b, 0xb5, 0xfb, 0x57, 0x43, 0x1b, 0x2e, 0x0b, 0x27, 0x5b, 0xe4, 0xbe, 0xf6, 0xf1,
	0x30, 0x27, 0x7a, 0x2c, 0x20, 0x91, 0x67, 0x32, 0xee, 0x59, 0x1e, 0x8d, 0xa0, 0x6c, 0xeb, 0xa6,
	0x9a, 0x3b, 0x9e, 0xe4, 0xed, 0x3a, 0xf0, 0xab, 0xb1, 0xfe, 0x46, 0xb9, 0xee, 0x41, 0xfa, 0xfa,
	0xff, 0xc1, 0x3c, 0xe9, 0xf5, 0x7b, 0xbb, 0x99, 0xec, 0x4f, 0xc1, 0x3b, 0x05, 0xde, 0x69, 0x9d,
	0x77, 0x7a, 0x52, 0xd8, 0x5f, 0x37, 0xba, 0x8a, 0x67, 0xdb, 0x40, 0xb4, 0xed, 0x3a, 0xd3, 0xb6,
	0x33, 0x2a, 0x58, 0x0e, 0x67, 0xa0, 0xf4, 0xe7, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xae, 0x8a,
	0xb3, 0x2f, 0xdf, 0x06, 0x00, 0x00,
}