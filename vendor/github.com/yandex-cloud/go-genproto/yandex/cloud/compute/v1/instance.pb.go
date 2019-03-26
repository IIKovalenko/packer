// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/compute/v1/instance.proto

package compute // import "github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"

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

type IpVersion int32

const (
	IpVersion_IP_VERSION_UNSPECIFIED IpVersion = 0
	// IPv4 address, for example 192.0.2.235.
	IpVersion_IPV4 IpVersion = 1
	// IPv6 address: not available yet.
	IpVersion_IPV6 IpVersion = 2
)

var IpVersion_name = map[int32]string{
	0: "IP_VERSION_UNSPECIFIED",
	1: "IPV4",
	2: "IPV6",
}
var IpVersion_value = map[string]int32{
	"IP_VERSION_UNSPECIFIED": 0,
	"IPV4": 1,
	"IPV6": 2,
}

func (x IpVersion) String() string {
	return proto.EnumName(IpVersion_name, int32(x))
}
func (IpVersion) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_instance_9ed228c16f2b2625, []int{0}
}

type Instance_Status int32

const (
	Instance_STATUS_UNSPECIFIED Instance_Status = 0
	// Instance is waiting for resources to be allocated.
	Instance_PROVISIONING Instance_Status = 1
	// Instance is running normally.
	Instance_RUNNING Instance_Status = 2
	// Instance is being stopped.
	Instance_STOPPING Instance_Status = 3
	// Instance stopped.
	Instance_STOPPED Instance_Status = 4
	// Instance is being started.
	Instance_STARTING Instance_Status = 5
	// Instance is being restarted.
	Instance_RESTARTING Instance_Status = 6
	// Instance is being updated.
	Instance_UPDATING Instance_Status = 7
	// Instance encountered a problem and cannot operate.
	Instance_ERROR Instance_Status = 8
	// Instance crashed and will be restarted automatically.
	Instance_CRASHED Instance_Status = 9
	// Instance is being deleted.
	Instance_DELETING Instance_Status = 10
)

var Instance_Status_name = map[int32]string{
	0:  "STATUS_UNSPECIFIED",
	1:  "PROVISIONING",
	2:  "RUNNING",
	3:  "STOPPING",
	4:  "STOPPED",
	5:  "STARTING",
	6:  "RESTARTING",
	7:  "UPDATING",
	8:  "ERROR",
	9:  "CRASHED",
	10: "DELETING",
}
var Instance_Status_value = map[string]int32{
	"STATUS_UNSPECIFIED": 0,
	"PROVISIONING":       1,
	"RUNNING":            2,
	"STOPPING":           3,
	"STOPPED":            4,
	"STARTING":           5,
	"RESTARTING":         6,
	"UPDATING":           7,
	"ERROR":              8,
	"CRASHED":            9,
	"DELETING":           10,
}

func (x Instance_Status) String() string {
	return proto.EnumName(Instance_Status_name, int32(x))
}
func (Instance_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_instance_9ed228c16f2b2625, []int{0, 0}
}

type AttachedDisk_Mode int32

const (
	AttachedDisk_MODE_UNSPECIFIED AttachedDisk_Mode = 0
	// Read-only access.
	AttachedDisk_READ_ONLY AttachedDisk_Mode = 1
	// Read/Write access.
	AttachedDisk_READ_WRITE AttachedDisk_Mode = 2
)

var AttachedDisk_Mode_name = map[int32]string{
	0: "MODE_UNSPECIFIED",
	1: "READ_ONLY",
	2: "READ_WRITE",
}
var AttachedDisk_Mode_value = map[string]int32{
	"MODE_UNSPECIFIED": 0,
	"READ_ONLY":        1,
	"READ_WRITE":       2,
}

func (x AttachedDisk_Mode) String() string {
	return proto.EnumName(AttachedDisk_Mode_name, int32(x))
}
func (AttachedDisk_Mode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_instance_9ed228c16f2b2625, []int{2, 0}
}

// An Instance resource. For more information, see [Instances](/docs/compute/concepts/vm).
type Instance struct {
	// ID of the instance.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the folder that the instance belongs to.
	FolderId string `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Creation timestamp in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Name of the instance. 1-63 characters long.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the instance. 0-256 characters long.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Resource labels as `` key:value `` pairs. Maximum of 64 per resource.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// ID of the availability zone where the instance resides.
	ZoneId string `protobuf:"bytes,7,opt,name=zone_id,json=zoneId,proto3" json:"zone_id,omitempty"`
	// ID of the hardware platform configuration for the instance.
	PlatformId string `protobuf:"bytes,8,opt,name=platform_id,json=platformId,proto3" json:"platform_id,omitempty"`
	// Computing resources of the instance such as the amount of memory and number of cores.
	Resources *Resources `protobuf:"bytes,9,opt,name=resources,proto3" json:"resources,omitempty"`
	// Status of the instance.
	Status Instance_Status `protobuf:"varint,10,opt,name=status,proto3,enum=yandex.cloud.compute.v1.Instance_Status" json:"status,omitempty"`
	// The metadata key/value pairs assigned to this instance. This includes custom metadata and predefined keys.
	//
	// For example, you may use the metadata in order to provide your public SSH key to the instance.
	// For more information, see [Metadata](/docs/compute/concepts/vm-metadata).
	Metadata map[string]string `protobuf:"bytes,11,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Boot disk that is attached to the instance.
	BootDisk *AttachedDisk `protobuf:"bytes,12,opt,name=boot_disk,json=bootDisk,proto3" json:"boot_disk,omitempty"`
	// Array of secondary disks that are attached to the instance.
	SecondaryDisks []*AttachedDisk `protobuf:"bytes,13,rep,name=secondary_disks,json=secondaryDisks,proto3" json:"secondary_disks,omitempty"`
	// Array of network interfaces that are attached to the instance.
	NetworkInterfaces []*NetworkInterface `protobuf:"bytes,14,rep,name=network_interfaces,json=networkInterfaces,proto3" json:"network_interfaces,omitempty"`
	// A domain name of the instance. FQDN is defined by the server
	// in the format `<hostname>.<region_id>.internal` when the instance is created.
	// If the hostname were not specified when the instance was created, FQDN would be `<id>.auto.internal`.
	Fqdn string `protobuf:"bytes,16,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
	// Scheduling policy configuration.
	SchedulingPolicy     *SchedulingPolicy `protobuf:"bytes,17,opt,name=scheduling_policy,json=schedulingPolicy,proto3" json:"scheduling_policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Instance) Reset()         { *m = Instance{} }
func (m *Instance) String() string { return proto.CompactTextString(m) }
func (*Instance) ProtoMessage()    {}
func (*Instance) Descriptor() ([]byte, []int) {
	return fileDescriptor_instance_9ed228c16f2b2625, []int{0}
}
func (m *Instance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Instance.Unmarshal(m, b)
}
func (m *Instance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Instance.Marshal(b, m, deterministic)
}
func (dst *Instance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Instance.Merge(dst, src)
}
func (m *Instance) XXX_Size() int {
	return xxx_messageInfo_Instance.Size(m)
}
func (m *Instance) XXX_DiscardUnknown() {
	xxx_messageInfo_Instance.DiscardUnknown(m)
}

var xxx_messageInfo_Instance proto.InternalMessageInfo

func (m *Instance) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Instance) GetFolderId() string {
	if m != nil {
		return m.FolderId
	}
	return ""
}

func (m *Instance) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Instance) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Instance) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Instance) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Instance) GetZoneId() string {
	if m != nil {
		return m.ZoneId
	}
	return ""
}

func (m *Instance) GetPlatformId() string {
	if m != nil {
		return m.PlatformId
	}
	return ""
}

func (m *Instance) GetResources() *Resources {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *Instance) GetStatus() Instance_Status {
	if m != nil {
		return m.Status
	}
	return Instance_STATUS_UNSPECIFIED
}

func (m *Instance) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Instance) GetBootDisk() *AttachedDisk {
	if m != nil {
		return m.BootDisk
	}
	return nil
}

func (m *Instance) GetSecondaryDisks() []*AttachedDisk {
	if m != nil {
		return m.SecondaryDisks
	}
	return nil
}

func (m *Instance) GetNetworkInterfaces() []*NetworkInterface {
	if m != nil {
		return m.NetworkInterfaces
	}
	return nil
}

func (m *Instance) GetFqdn() string {
	if m != nil {
		return m.Fqdn
	}
	return ""
}

func (m *Instance) GetSchedulingPolicy() *SchedulingPolicy {
	if m != nil {
		return m.SchedulingPolicy
	}
	return nil
}

type Resources struct {
	// The amount of memory available to the instance, specified in bytes.
	Memory int64 `protobuf:"varint,1,opt,name=memory,proto3" json:"memory,omitempty"`
	// The number of cores available to the instance.
	Cores int64 `protobuf:"varint,2,opt,name=cores,proto3" json:"cores,omitempty"`
	// Baseline level of CPU performance with the ability to burst performance above that baseline level.
	// This field sets baseline performance for each core.
	CoreFraction         int64    `protobuf:"varint,3,opt,name=core_fraction,json=coreFraction,proto3" json:"core_fraction,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resources) Reset()         { *m = Resources{} }
func (m *Resources) String() string { return proto.CompactTextString(m) }
func (*Resources) ProtoMessage()    {}
func (*Resources) Descriptor() ([]byte, []int) {
	return fileDescriptor_instance_9ed228c16f2b2625, []int{1}
}
func (m *Resources) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resources.Unmarshal(m, b)
}
func (m *Resources) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resources.Marshal(b, m, deterministic)
}
func (dst *Resources) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resources.Merge(dst, src)
}
func (m *Resources) XXX_Size() int {
	return xxx_messageInfo_Resources.Size(m)
}
func (m *Resources) XXX_DiscardUnknown() {
	xxx_messageInfo_Resources.DiscardUnknown(m)
}

var xxx_messageInfo_Resources proto.InternalMessageInfo

func (m *Resources) GetMemory() int64 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func (m *Resources) GetCores() int64 {
	if m != nil {
		return m.Cores
	}
	return 0
}

func (m *Resources) GetCoreFraction() int64 {
	if m != nil {
		return m.CoreFraction
	}
	return 0
}

type AttachedDisk struct {
	// Access mode to the Disk resource.
	Mode AttachedDisk_Mode `protobuf:"varint,1,opt,name=mode,proto3,enum=yandex.cloud.compute.v1.AttachedDisk_Mode" json:"mode,omitempty"`
	// Serial number that is reflected into the /dev/disk/by-id/ tree
	// of a Linux operating system running within the instance.
	//
	// This value can be used to reference the device for mounting, resizing, and so on, from within the instance.
	DeviceName string `protobuf:"bytes,2,opt,name=device_name,json=deviceName,proto3" json:"device_name,omitempty"`
	// Specifies whether the disk will be auto-deleted when the instance is deleted.
	AutoDelete bool `protobuf:"varint,3,opt,name=auto_delete,json=autoDelete,proto3" json:"auto_delete,omitempty"`
	// ID of the disk that is attached to the instance.
	DiskId               string   `protobuf:"bytes,4,opt,name=disk_id,json=diskId,proto3" json:"disk_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AttachedDisk) Reset()         { *m = AttachedDisk{} }
func (m *AttachedDisk) String() string { return proto.CompactTextString(m) }
func (*AttachedDisk) ProtoMessage()    {}
func (*AttachedDisk) Descriptor() ([]byte, []int) {
	return fileDescriptor_instance_9ed228c16f2b2625, []int{2}
}
func (m *AttachedDisk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttachedDisk.Unmarshal(m, b)
}
func (m *AttachedDisk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttachedDisk.Marshal(b, m, deterministic)
}
func (dst *AttachedDisk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttachedDisk.Merge(dst, src)
}
func (m *AttachedDisk) XXX_Size() int {
	return xxx_messageInfo_AttachedDisk.Size(m)
}
func (m *AttachedDisk) XXX_DiscardUnknown() {
	xxx_messageInfo_AttachedDisk.DiscardUnknown(m)
}

var xxx_messageInfo_AttachedDisk proto.InternalMessageInfo

func (m *AttachedDisk) GetMode() AttachedDisk_Mode {
	if m != nil {
		return m.Mode
	}
	return AttachedDisk_MODE_UNSPECIFIED
}

func (m *AttachedDisk) GetDeviceName() string {
	if m != nil {
		return m.DeviceName
	}
	return ""
}

func (m *AttachedDisk) GetAutoDelete() bool {
	if m != nil {
		return m.AutoDelete
	}
	return false
}

func (m *AttachedDisk) GetDiskId() string {
	if m != nil {
		return m.DiskId
	}
	return ""
}

type NetworkInterface struct {
	// The index of the network interface, generated by the server, 0,1,2... etc.
	// Currently only one network interface is supported per instance.
	Index string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	// MAC address that is assigned to the network interface.
	MacAddress string `protobuf:"bytes,2,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	// ID of the subnet.
	SubnetId string `protobuf:"bytes,3,opt,name=subnet_id,json=subnetId,proto3" json:"subnet_id,omitempty"`
	// Primary IPv4 address that is assigned to the instance for this network interface.
	PrimaryV4Address *PrimaryAddress `protobuf:"bytes,4,opt,name=primary_v4_address,json=primaryV4Address,proto3" json:"primary_v4_address,omitempty"`
	// Primary IPv6 address that is assigned to the instance for this network interface. IPv6 not available yet.
	PrimaryV6Address     *PrimaryAddress `protobuf:"bytes,5,opt,name=primary_v6_address,json=primaryV6Address,proto3" json:"primary_v6_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *NetworkInterface) Reset()         { *m = NetworkInterface{} }
func (m *NetworkInterface) String() string { return proto.CompactTextString(m) }
func (*NetworkInterface) ProtoMessage()    {}
func (*NetworkInterface) Descriptor() ([]byte, []int) {
	return fileDescriptor_instance_9ed228c16f2b2625, []int{3}
}
func (m *NetworkInterface) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkInterface.Unmarshal(m, b)
}
func (m *NetworkInterface) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkInterface.Marshal(b, m, deterministic)
}
func (dst *NetworkInterface) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkInterface.Merge(dst, src)
}
func (m *NetworkInterface) XXX_Size() int {
	return xxx_messageInfo_NetworkInterface.Size(m)
}
func (m *NetworkInterface) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkInterface.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkInterface proto.InternalMessageInfo

func (m *NetworkInterface) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *NetworkInterface) GetMacAddress() string {
	if m != nil {
		return m.MacAddress
	}
	return ""
}

func (m *NetworkInterface) GetSubnetId() string {
	if m != nil {
		return m.SubnetId
	}
	return ""
}

func (m *NetworkInterface) GetPrimaryV4Address() *PrimaryAddress {
	if m != nil {
		return m.PrimaryV4Address
	}
	return nil
}

func (m *NetworkInterface) GetPrimaryV6Address() *PrimaryAddress {
	if m != nil {
		return m.PrimaryV6Address
	}
	return nil
}

type PrimaryAddress struct {
	// An IPv4 internal network address that is assigned to the instance for this network interface.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// One-to-one NAT configuration. If missing, NAT has not been set up.
	OneToOneNat          *OneToOneNat `protobuf:"bytes,2,opt,name=one_to_one_nat,json=oneToOneNat,proto3" json:"one_to_one_nat,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PrimaryAddress) Reset()         { *m = PrimaryAddress{} }
func (m *PrimaryAddress) String() string { return proto.CompactTextString(m) }
func (*PrimaryAddress) ProtoMessage()    {}
func (*PrimaryAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_instance_9ed228c16f2b2625, []int{4}
}
func (m *PrimaryAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimaryAddress.Unmarshal(m, b)
}
func (m *PrimaryAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimaryAddress.Marshal(b, m, deterministic)
}
func (dst *PrimaryAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimaryAddress.Merge(dst, src)
}
func (m *PrimaryAddress) XXX_Size() int {
	return xxx_messageInfo_PrimaryAddress.Size(m)
}
func (m *PrimaryAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimaryAddress.DiscardUnknown(m)
}

var xxx_messageInfo_PrimaryAddress proto.InternalMessageInfo

func (m *PrimaryAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *PrimaryAddress) GetOneToOneNat() *OneToOneNat {
	if m != nil {
		return m.OneToOneNat
	}
	return nil
}

type OneToOneNat struct {
	// An external IP address associated with this instance.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// IP version for the external IP address.
	IpVersion            IpVersion `protobuf:"varint,2,opt,name=ip_version,json=ipVersion,proto3,enum=yandex.cloud.compute.v1.IpVersion" json:"ip_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *OneToOneNat) Reset()         { *m = OneToOneNat{} }
func (m *OneToOneNat) String() string { return proto.CompactTextString(m) }
func (*OneToOneNat) ProtoMessage()    {}
func (*OneToOneNat) Descriptor() ([]byte, []int) {
	return fileDescriptor_instance_9ed228c16f2b2625, []int{5}
}
func (m *OneToOneNat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OneToOneNat.Unmarshal(m, b)
}
func (m *OneToOneNat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OneToOneNat.Marshal(b, m, deterministic)
}
func (dst *OneToOneNat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OneToOneNat.Merge(dst, src)
}
func (m *OneToOneNat) XXX_Size() int {
	return xxx_messageInfo_OneToOneNat.Size(m)
}
func (m *OneToOneNat) XXX_DiscardUnknown() {
	xxx_messageInfo_OneToOneNat.DiscardUnknown(m)
}

var xxx_messageInfo_OneToOneNat proto.InternalMessageInfo

func (m *OneToOneNat) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *OneToOneNat) GetIpVersion() IpVersion {
	if m != nil {
		return m.IpVersion
	}
	return IpVersion_IP_VERSION_UNSPECIFIED
}

type SchedulingPolicy struct {
	// Set if instance is preemptible.
	Preemptible          bool     `protobuf:"varint,1,opt,name=preemptible,proto3" json:"preemptible,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SchedulingPolicy) Reset()         { *m = SchedulingPolicy{} }
func (m *SchedulingPolicy) String() string { return proto.CompactTextString(m) }
func (*SchedulingPolicy) ProtoMessage()    {}
func (*SchedulingPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_instance_9ed228c16f2b2625, []int{6}
}
func (m *SchedulingPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SchedulingPolicy.Unmarshal(m, b)
}
func (m *SchedulingPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SchedulingPolicy.Marshal(b, m, deterministic)
}
func (dst *SchedulingPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SchedulingPolicy.Merge(dst, src)
}
func (m *SchedulingPolicy) XXX_Size() int {
	return xxx_messageInfo_SchedulingPolicy.Size(m)
}
func (m *SchedulingPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_SchedulingPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_SchedulingPolicy proto.InternalMessageInfo

func (m *SchedulingPolicy) GetPreemptible() bool {
	if m != nil {
		return m.Preemptible
	}
	return false
}

func init() {
	proto.RegisterType((*Instance)(nil), "yandex.cloud.compute.v1.Instance")
	proto.RegisterMapType((map[string]string)(nil), "yandex.cloud.compute.v1.Instance.LabelsEntry")
	proto.RegisterMapType((map[string]string)(nil), "yandex.cloud.compute.v1.Instance.MetadataEntry")
	proto.RegisterType((*Resources)(nil), "yandex.cloud.compute.v1.Resources")
	proto.RegisterType((*AttachedDisk)(nil), "yandex.cloud.compute.v1.AttachedDisk")
	proto.RegisterType((*NetworkInterface)(nil), "yandex.cloud.compute.v1.NetworkInterface")
	proto.RegisterType((*PrimaryAddress)(nil), "yandex.cloud.compute.v1.PrimaryAddress")
	proto.RegisterType((*OneToOneNat)(nil), "yandex.cloud.compute.v1.OneToOneNat")
	proto.RegisterType((*SchedulingPolicy)(nil), "yandex.cloud.compute.v1.SchedulingPolicy")
	proto.RegisterEnum("yandex.cloud.compute.v1.IpVersion", IpVersion_name, IpVersion_value)
	proto.RegisterEnum("yandex.cloud.compute.v1.Instance_Status", Instance_Status_name, Instance_Status_value)
	proto.RegisterEnum("yandex.cloud.compute.v1.AttachedDisk_Mode", AttachedDisk_Mode_name, AttachedDisk_Mode_value)
}

func init() {
	proto.RegisterFile("yandex/cloud/compute/v1/instance.proto", fileDescriptor_instance_9ed228c16f2b2625)
}

var fileDescriptor_instance_9ed228c16f2b2625 = []byte{
	// 1058 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0x6d, 0x6f, 0xe3, 0x44,
	0x10, 0x26, 0x2f, 0x4d, 0xe3, 0x49, 0x1b, 0xdc, 0xd5, 0xa9, 0x67, 0x95, 0x0f, 0xad, 0xc2, 0x5b,
	0x39, 0xa9, 0x8e, 0xae, 0x54, 0x15, 0x47, 0x25, 0x74, 0xb9, 0xc6, 0x07, 0x16, 0x6d, 0x12, 0x6d,
	0xd2, 0xf0, 0xf2, 0x01, 0x6b, 0xe3, 0xdd, 0xe4, 0x4c, 0x6d, 0xaf, 0xb1, 0xd7, 0x81, 0xf2, 0x3b,
	0xf8, 0x19, 0xfc, 0x2e, 0xfe, 0x04, 0x5f, 0xd0, 0xee, 0xda, 0xb9, 0x5c, 0xa4, 0x70, 0x07, 0x9f,
	0xb2, 0xf3, 0xcc, 0xcc, 0x33, 0x2f, 0x99, 0x9d, 0x35, 0x7c, 0xf2, 0x40, 0x62, 0xca, 0x7e, 0xeb,
	0xfa, 0x21, 0xcf, 0x69, 0xd7, 0xe7, 0x51, 0x92, 0x0b, 0xd6, 0x5d, 0x3e, 0xed, 0x06, 0x71, 0x26,
	0x48, 0xec, 0x33, 0x3b, 0x49, 0xb9, 0xe0, 0xe8, 0xb1, 0xb6, 0xb3, 0x95, 0x9d, 0x5d, 0xd8, 0xd9,
	0xcb, 0xa7, 0x47, 0xc7, 0x0b, 0xce, 0x17, 0x21, 0xeb, 0x2a, 0xb3, 0x59, 0x3e, 0xef, 0x8a, 0x20,
	0x62, 0x99, 0x20, 0x51, 0xa2, 0x3d, 0x3b, 0x7f, 0x37, 0xa1, 0xe9, 0x16, 0x64, 0xa8, 0x0d, 0xd5,
	0x80, 0x5a, 0x95, 0x93, 0xca, 0xa9, 0x81, 0xab, 0x01, 0x45, 0x1f, 0x80, 0x31, 0xe7, 0x21, 0x65,
	0xa9, 0x17, 0x50, 0xab, 0xaa, 0xe0, 0xa6, 0x06, 0x5c, 0x8a, 0x9e, 0x01, 0xf8, 0x29, 0x23, 0x82,
	0x51, 0x8f, 0x08, 0xab, 0x76, 0x52, 0x39, 0x6d, 0x9d, 0x1f, 0xd9, 0x3a, 0x9e, 0x5d, 0xc6, 0xb3,
	0x27, 0x65, 0x3c, 0x6c, 0x14, 0xd6, 0x3d, 0x81, 0x10, 0xd4, 0x63, 0x12, 0x31, 0xab, 0xae, 0x28,
	0xd5, 0x19, 0x9d, 0x40, 0x8b, 0xb2, 0xcc, 0x4f, 0x83, 0x44, 0x04, 0x3c, 0xb6, 0x76, 0x94, 0x6a,
	0x1d, 0x42, 0x0e, 0x34, 0x42, 0x32, 0x63, 0x61, 0x66, 0x35, 0x4e, 0x6a, 0xa7, 0xad, 0xf3, 0x33,
	0x7b, 0x4b, 0xd5, 0x76, 0x59, 0x90, 0x7d, 0xa3, 0xec, 0x9d, 0x58, 0xa4, 0x0f, 0xb8, 0x70, 0x46,
	0x8f, 0x61, 0xf7, 0x77, 0x1e, 0x33, 0x59, 0xd2, 0xae, 0x0a, 0xd2, 0x90, 0xa2, 0x4b, 0xd1, 0x31,
	0xb4, 0x92, 0x90, 0x88, 0x39, 0x4f, 0x23, 0xa9, 0x6c, 0x2a, 0x25, 0x94, 0x90, 0x4b, 0xd1, 0x73,
	0x30, 0x52, 0x96, 0xf1, 0x3c, 0xf5, 0x59, 0x66, 0x19, 0xaa, 0xe0, 0xce, 0xd6, 0x1c, 0x70, 0x69,
	0x89, 0x5f, 0x3b, 0xa1, 0xe7, 0xd0, 0xc8, 0x04, 0x11, 0x79, 0x66, 0xc1, 0x49, 0xe5, 0xb4, 0x7d,
	0x7e, 0xfa, 0xf6, 0x12, 0xc6, 0xca, 0x1e, 0x17, 0x7e, 0xe8, 0x5b, 0x68, 0x46, 0x4c, 0x10, 0x4a,
	0x04, 0xb1, 0x5a, 0xaa, 0x0d, 0xdd, 0xb7, 0x73, 0xdc, 0x16, 0x1e, 0xba, 0x11, 0x2b, 0x02, 0xf4,
	0x02, 0x8c, 0x19, 0xe7, 0xc2, 0xa3, 0x41, 0x76, 0x6f, 0xed, 0xa9, 0x82, 0x3e, 0xde, 0xca, 0xd6,
	0x13, 0x82, 0xf8, 0xaf, 0x18, 0xed, 0x07, 0xd9, 0x3d, 0x6e, 0x4a, 0x3f, 0x79, 0x42, 0x03, 0x78,
	0x3f, 0x63, 0x3e, 0x8f, 0x29, 0x49, 0x1f, 0x14, 0x51, 0x66, 0xed, 0xab, 0xbc, 0xde, 0x91, 0xa9,
	0xbd, 0xf2, 0x96, 0x62, 0x86, 0xbe, 0x07, 0x14, 0x33, 0xf1, 0x2b, 0x4f, 0xef, 0xbd, 0x20, 0x16,
	0x2c, 0x9d, 0x13, 0xd9, 0xed, 0xb6, 0xa2, 0xfc, 0x6c, 0x2b, 0xe5, 0x40, 0xbb, 0xb8, 0xa5, 0x07,
	0x3e, 0x88, 0x37, 0x90, 0x4c, 0x4e, 0xdd, 0xfc, 0x17, 0x1a, 0x5b, 0xa6, 0x9e, 0x3a, 0x79, 0x46,
	0x53, 0x38, 0xc8, 0x64, 0x2a, 0x79, 0x18, 0xc4, 0x0b, 0x2f, 0xe1, 0x61, 0xe0, 0x3f, 0x58, 0x07,
	0xaa, 0x13, 0xdb, 0x83, 0x8d, 0x57, 0x1e, 0x23, 0xe5, 0x80, 0xcd, 0x6c, 0x03, 0x39, 0x7a, 0x06,
	0xad, 0xb5, 0xd9, 0x43, 0x26, 0xd4, 0xee, 0xd9, 0x43, 0x71, 0xb3, 0xe4, 0x11, 0x3d, 0x82, 0x9d,
	0x25, 0x09, 0x73, 0x56, 0x5c, 0x2b, 0x2d, 0x7c, 0x59, 0xfd, 0xa2, 0x72, 0x74, 0x05, 0xfb, 0x6f,
	0xfc, 0x5f, 0xff, 0xc5, 0xb9, 0xf3, 0x67, 0x05, 0x1a, 0x7a, 0x62, 0xd0, 0x21, 0xa0, 0xf1, 0xa4,
	0x37, 0xb9, 0x1b, 0x7b, 0x77, 0x83, 0xf1, 0xc8, 0xb9, 0x76, 0x5f, 0xba, 0x4e, 0xdf, 0x7c, 0x0f,
	0x99, 0xb0, 0x37, 0xc2, 0xc3, 0xa9, 0x3b, 0x76, 0x87, 0x03, 0x77, 0xf0, 0xb5, 0x59, 0x41, 0x2d,
	0xd8, 0xc5, 0x77, 0x03, 0x25, 0x54, 0xd1, 0x1e, 0x34, 0xc7, 0x93, 0xe1, 0x68, 0x24, 0xa5, 0x9a,
	0x54, 0x29, 0xc9, 0xe9, 0x9b, 0x75, 0xad, 0xea, 0xe1, 0x89, 0x54, 0xed, 0xa0, 0x36, 0x00, 0x76,
	0x56, 0x72, 0x43, 0x6a, 0xef, 0x46, 0xfd, 0x9e, 0x92, 0x76, 0x91, 0x01, 0x3b, 0x0e, 0xc6, 0x43,
	0x6c, 0x36, 0x25, 0xc7, 0x35, 0xee, 0x8d, 0xbf, 0x71, 0xfa, 0xa6, 0x21, 0xad, 0xfa, 0xce, 0x8d,
	0xa3, 0xac, 0xa0, 0xf3, 0x13, 0x18, 0xab, 0x7b, 0x82, 0x0e, 0xa1, 0x11, 0xb1, 0x88, 0xa7, 0xba,
	0xd4, 0x1a, 0x2e, 0x24, 0x59, 0xad, 0xcf, 0x53, 0x96, 0xa9, 0x6a, 0x6b, 0x58, 0x0b, 0xe8, 0x43,
	0xd8, 0x97, 0x07, 0x6f, 0x9e, 0x12, 0x5f, 0x6d, 0x8c, 0x9a, 0xd2, 0xee, 0x49, 0xf0, 0x65, 0x81,
	0x75, 0xfe, 0xaa, 0xc0, 0xde, 0xfa, 0xb4, 0xa1, 0xaf, 0xa0, 0x1e, 0x71, 0xca, 0x54, 0x84, 0xf6,
	0xf9, 0x93, 0x77, 0x1a, 0x51, 0xfb, 0x96, 0x53, 0x86, 0x95, 0x9f, 0xdc, 0x11, 0x94, 0x2d, 0x03,
	0x9f, 0x79, 0x6a, 0x81, 0xe9, 0xfe, 0x83, 0x86, 0x06, 0x72, 0x8d, 0x1d, 0x43, 0x8b, 0xe4, 0x82,
	0x7b, 0x94, 0x85, 0x4c, 0x30, 0x95, 0x54, 0x13, 0x83, 0x84, 0xfa, 0x0a, 0x91, 0xeb, 0x47, 0xde,
	0x12, 0xb9, 0x61, 0xf4, 0xfa, 0x6b, 0x48, 0xd1, 0xa5, 0x9d, 0x2b, 0xa8, 0xcb, 0x40, 0xe8, 0x11,
	0x98, 0xb7, 0xc3, 0xbe, 0xb3, 0xf1, 0xaf, 0xed, 0x83, 0x81, 0x9d, 0x5e, 0xdf, 0x1b, 0x0e, 0x6e,
	0x7e, 0x30, 0x2b, 0xba, 0xf9, 0xbd, 0xbe, 0xf7, 0x1d, 0x76, 0x27, 0x8e, 0x59, 0xed, 0xfc, 0x51,
	0x05, 0x73, 0xf3, 0x0e, 0xc8, 0xc6, 0x05, 0xb2, 0xbc, 0x62, 0x74, 0xb4, 0x20, 0x33, 0x8c, 0x88,
	0xef, 0x11, 0x4a, 0x53, 0x96, 0x65, 0x65, 0x09, 0x11, 0xf1, 0x7b, 0x1a, 0x91, 0x5b, 0x3f, 0xcb,
	0x67, 0x31, 0x13, 0x32, 0xc7, 0x9a, 0xde, 0xfa, 0x1a, 0x70, 0x29, 0xba, 0x03, 0x94, 0xa4, 0x41,
	0x24, 0x2f, 0xfb, 0xf2, 0x62, 0x45, 0x52, 0x57, 0x37, 0xe6, 0xd3, 0xad, 0xed, 0x1c, 0x69, 0x97,
	0x22, 0x02, 0x36, 0x0b, 0x8a, 0xe9, 0x45, 0x19, 0x73, 0x9d, 0xf6, 0x72, 0x45, 0xbb, 0xf3, 0x3f,
	0x69, 0x2f, 0x0b, 0xa4, 0x93, 0x43, 0xfb, 0x4d, 0x1b, 0x64, 0xc1, 0x6e, 0xc9, 0xae, 0xbb, 0x52,
	0x8a, 0xc8, 0x85, 0xb6, 0x7c, 0x16, 0x04, 0xf7, 0xe4, 0x4f, 0x4c, 0x84, 0x6a, 0x4d, 0xeb, 0xfc,
	0xa3, 0xad, 0xe1, 0x87, 0x31, 0x9b, 0xf0, 0x61, 0xcc, 0x06, 0x44, 0xe0, 0x16, 0x7f, 0x2d, 0x74,
	0x7e, 0x86, 0xd6, 0x9a, 0xee, 0x5f, 0x62, 0xf6, 0x00, 0x82, 0xc4, 0x5b, 0xb2, 0x34, 0x93, 0x13,
	0x5c, 0x55, 0x43, 0xb9, 0xfd, 0x49, 0x71, 0x93, 0xa9, 0xb6, 0xc4, 0x46, 0x50, 0x1e, 0x3b, 0x17,
	0x60, 0x6e, 0xee, 0x23, 0xf9, 0x96, 0x26, 0x29, 0x63, 0x51, 0x22, 0x82, 0x59, 0xa8, 0x87, 0xbd,
	0x89, 0xd7, 0xa1, 0x27, 0x57, 0x60, 0xac, 0xd8, 0xd0, 0x11, 0x1c, 0xba, 0x23, 0x6f, 0xea, 0x60,
	0xb9, 0x12, 0x36, 0xe6, 0xae, 0x09, 0x75, 0x77, 0x34, 0xbd, 0x30, 0x2b, 0xc5, 0xe9, 0xd2, 0xac,
	0xbe, 0x70, 0x7e, 0xbc, 0x5e, 0x04, 0xe2, 0x55, 0x3e, 0x93, 0xc9, 0x75, 0x75, 0xb6, 0x67, 0xfa,
	0x13, 0x65, 0xc1, 0xcf, 0x16, 0x2c, 0x56, 0xaf, 0x7f, 0x77, 0xcb, 0xb7, 0xcb, 0x55, 0x71, 0x9c,
	0x35, 0x94, 0xd9, 0xe7, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x78, 0x8d, 0xb6, 0x7a, 0xe5, 0x08,
	0x00, 0x00,
}