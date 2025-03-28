// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/mdb/mongodb/v1/backup.proto

package mongodb // import "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/mongodb/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A MongoDB Backup resource. For more information, see the
// [Developer's Guide](/docs/managed-mongodb/concepts).
type Backup struct {
	// ID of the backup.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the folder that the backup belongs to.
	FolderId string `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Creation timestamp in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format
	// (i.e. when the backup operation was completed).
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// ID of the MongoDB cluster that the backup was created for.
	SourceClusterId string `protobuf:"bytes,4,opt,name=source_cluster_id,json=sourceClusterId,proto3" json:"source_cluster_id,omitempty"`
	// Time when the backup operation was started.
	StartedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Backup) Reset()         { *m = Backup{} }
func (m *Backup) String() string { return proto.CompactTextString(m) }
func (*Backup) ProtoMessage()    {}
func (*Backup) Descriptor() ([]byte, []int) {
	return fileDescriptor_backup_7bed62cf0a79db64, []int{0}
}
func (m *Backup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Backup.Unmarshal(m, b)
}
func (m *Backup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Backup.Marshal(b, m, deterministic)
}
func (dst *Backup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Backup.Merge(dst, src)
}
func (m *Backup) XXX_Size() int {
	return xxx_messageInfo_Backup.Size(m)
}
func (m *Backup) XXX_DiscardUnknown() {
	xxx_messageInfo_Backup.DiscardUnknown(m)
}

var xxx_messageInfo_Backup proto.InternalMessageInfo

func (m *Backup) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Backup) GetFolderId() string {
	if m != nil {
		return m.FolderId
	}
	return ""
}

func (m *Backup) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Backup) GetSourceClusterId() string {
	if m != nil {
		return m.SourceClusterId
	}
	return ""
}

func (m *Backup) GetStartedAt() *timestamp.Timestamp {
	if m != nil {
		return m.StartedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Backup)(nil), "yandex.cloud.mdb.mongodb.v1.Backup")
}

func init() {
	proto.RegisterFile("yandex/cloud/mdb/mongodb/v1/backup.proto", fileDescriptor_backup_7bed62cf0a79db64)
}

var fileDescriptor_backup_7bed62cf0a79db64 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0x69, 0xd5, 0x61, 0x23, 0x28, 0xf6, 0x54, 0xb6, 0x83, 0xc3, 0x53, 0x11, 0x96, 0x30,
	0x3d, 0x89, 0xa7, 0xcd, 0x83, 0xec, 0x3a, 0x3c, 0x79, 0x29, 0x49, 0x5e, 0x16, 0x8b, 0x4d, 0xdf,
	0x68, 0x5f, 0x86, 0xfe, 0xa5, 0xfe, 0x3b, 0x42, 0x92, 0x5d, 0xdd, 0x2d, 0x7c, 0xf9, 0xbd, 0xef,
	0x07, 0x1f, 0xab, 0x7f, 0x64, 0x0f, 0xe6, 0x5b, 0xe8, 0x0e, 0x3d, 0x08, 0x07, 0x4a, 0x38, 0xec,
	0x2d, 0x82, 0x12, 0x87, 0xa5, 0x50, 0x52, 0x7f, 0xf9, 0x3d, 0xdf, 0x0f, 0x48, 0x58, 0xce, 0x22,
	0xc9, 0x03, 0xc9, 0x1d, 0x28, 0x9e, 0x48, 0x7e, 0x58, 0x4e, 0xef, 0x2c, 0xa2, 0xed, 0x8c, 0x08,
	0xa8, 0xf2, 0x3b, 0x41, 0xad, 0x33, 0x23, 0x49, 0x97, 0xae, 0xef, 0x7f, 0x33, 0x36, 0x59, 0x87,
	0xba, 0xf2, 0x9a, 0xe5, 0x2d, 0x54, 0xd9, 0x3c, 0xab, 0x8b, 0x6d, 0xde, 0x42, 0x39, 0x63, 0xc5,
	0x0e, 0x3b, 0x30, 0x43, 0xd3, 0x42, 0x95, 0x87, 0xf8, 0x32, 0x06, 0x1b, 0x28, 0x9f, 0x19, 0xd3,
	0x83, 0x91, 0x64, 0xa0, 0x91, 0x54, 0x9d, 0xcd, 0xb3, 0xfa, 0xea, 0x71, 0xca, 0xa3, 0x8d, 0x1f,
	0x6d, 0xfc, 0xfd, 0x68, 0xdb, 0x16, 0x89, 0x5e, 0x51, 0xf9, 0xc0, 0x6e, 0x47, 0xf4, 0x83, 0x36,
	0x8d, 0xee, 0xfc, 0x48, 0xb1, 0xff, 0x3c, 0xf4, 0xdf, 0xc4, 0x8f, 0xd7, 0x98, 0x47, 0xcd, 0x48,
	0x72, 0x48, 0x9a, 0x8b, 0xd3, 0x9a, 0x44, 0xaf, 0x68, 0xbd, 0xf9, 0x78, 0xb3, 0x2d, 0x7d, 0x7a,
	0xc5, 0x35, 0x3a, 0x11, 0x47, 0x5a, 0xc4, 0x39, 0x2d, 0x2e, 0xac, 0xe9, 0xc3, 0xb9, 0xf8, 0x67,
	0xe7, 0x97, 0xf4, 0x54, 0x93, 0x80, 0x3e, 0xfd, 0x05, 0x00, 0x00, 0xff, 0xff, 0xae, 0x68, 0xff,
	0xeb, 0x95, 0x01, 0x00, 0x00,
}
