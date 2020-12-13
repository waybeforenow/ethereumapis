// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: eth/v1alpha1/node.proto

package eth

import (
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type PeerDirection int32

const (
	PeerDirection_UNKNOWN  PeerDirection = 0
	PeerDirection_INBOUND  PeerDirection = 1
	PeerDirection_OUTBOUND PeerDirection = 2
)

// Enum value maps for PeerDirection.
var (
	PeerDirection_name = map[int32]string{
		0: "UNKNOWN",
		1: "INBOUND",
		2: "OUTBOUND",
	}
	PeerDirection_value = map[string]int32{
		"UNKNOWN":  0,
		"INBOUND":  1,
		"OUTBOUND": 2,
	}
)

func (x PeerDirection) Enum() *PeerDirection {
	p := new(PeerDirection)
	*p = x
	return p
}

func (x PeerDirection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PeerDirection) Descriptor() protoreflect.EnumDescriptor {
	return file_eth_v1alpha1_node_proto_enumTypes[0].Descriptor()
}

func (PeerDirection) Type() protoreflect.EnumType {
	return &file_eth_v1alpha1_node_proto_enumTypes[0]
}

func (x PeerDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PeerDirection.Descriptor instead.
func (PeerDirection) EnumDescriptor() ([]byte, []int) {
	return file_eth_v1alpha1_node_proto_rawDescGZIP(), []int{0}
}

type ConnectionState int32

const (
	ConnectionState_DISCONNECTED  ConnectionState = 0
	ConnectionState_DISCONNECTING ConnectionState = 1
	ConnectionState_CONNECTED     ConnectionState = 2
	ConnectionState_CONNECTING    ConnectionState = 3
)

// Enum value maps for ConnectionState.
var (
	ConnectionState_name = map[int32]string{
		0: "DISCONNECTED",
		1: "DISCONNECTING",
		2: "CONNECTED",
		3: "CONNECTING",
	}
	ConnectionState_value = map[string]int32{
		"DISCONNECTED":  0,
		"DISCONNECTING": 1,
		"CONNECTED":     2,
		"CONNECTING":    3,
	}
)

func (x ConnectionState) Enum() *ConnectionState {
	p := new(ConnectionState)
	*p = x
	return p
}

func (x ConnectionState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConnectionState) Descriptor() protoreflect.EnumDescriptor {
	return file_eth_v1alpha1_node_proto_enumTypes[1].Descriptor()
}

func (ConnectionState) Type() protoreflect.EnumType {
	return &file_eth_v1alpha1_node_proto_enumTypes[1]
}

func (x ConnectionState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConnectionState.Descriptor instead.
func (ConnectionState) EnumDescriptor() ([]byte, []int) {
	return file_eth_v1alpha1_node_proto_rawDescGZIP(), []int{1}
}

type SyncStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Syncing bool `protobuf:"varint,1,opt,name=syncing,proto3" json:"syncing,omitempty"`
}

func (x *SyncStatus) Reset() {
	*x = SyncStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1alpha1_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncStatus) ProtoMessage() {}

func (x *SyncStatus) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1alpha1_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncStatus.ProtoReflect.Descriptor instead.
func (*SyncStatus) Descriptor() ([]byte, []int) {
	return file_eth_v1alpha1_node_proto_rawDescGZIP(), []int{0}
}

func (x *SyncStatus) GetSyncing() bool {
	if x != nil {
		return x.Syncing
	}
	return false
}

type Genesis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GenesisTime            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=genesis_time,json=genesisTime,proto3" json:"genesis_time,omitempty"`
	DepositContractAddress []byte               `protobuf:"bytes,2,opt,name=deposit_contract_address,json=depositContractAddress,proto3" json:"deposit_contract_address,omitempty"`
	GenesisValidatorsRoot  []byte               `protobuf:"bytes,3,opt,name=genesis_validators_root,json=genesisValidatorsRoot,proto3" json:"genesis_validators_root,omitempty" ssz-size:"32"`
}

func (x *Genesis) Reset() {
	*x = Genesis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1alpha1_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Genesis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Genesis) ProtoMessage() {}

func (x *Genesis) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1alpha1_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Genesis.ProtoReflect.Descriptor instead.
func (*Genesis) Descriptor() ([]byte, []int) {
	return file_eth_v1alpha1_node_proto_rawDescGZIP(), []int{1}
}

func (x *Genesis) GetGenesisTime() *timestamp.Timestamp {
	if x != nil {
		return x.GenesisTime
	}
	return nil
}

func (x *Genesis) GetDepositContractAddress() []byte {
	if x != nil {
		return x.DepositContractAddress
	}
	return nil
}

func (x *Genesis) GetGenesisValidatorsRoot() []byte {
	if x != nil {
		return x.GenesisValidatorsRoot
	}
	return nil
}

type Version struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version  string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Metadata string `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *Version) Reset() {
	*x = Version{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1alpha1_node_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1alpha1_node_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version.ProtoReflect.Descriptor instead.
func (*Version) Descriptor() ([]byte, []int) {
	return file_eth_v1alpha1_node_proto_rawDescGZIP(), []int{2}
}

func (x *Version) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Version) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

type ImplementedServices struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Services []string `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *ImplementedServices) Reset() {
	*x = ImplementedServices{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1alpha1_node_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImplementedServices) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImplementedServices) ProtoMessage() {}

func (x *ImplementedServices) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1alpha1_node_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImplementedServices.ProtoReflect.Descriptor instead.
func (*ImplementedServices) Descriptor() ([]byte, []int) {
	return file_eth_v1alpha1_node_proto_rawDescGZIP(), []int{3}
}

func (x *ImplementedServices) GetServices() []string {
	if x != nil {
		return x.Services
	}
	return nil
}

type PeerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PeerId string `protobuf:"bytes,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
}

func (x *PeerRequest) Reset() {
	*x = PeerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1alpha1_node_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerRequest) ProtoMessage() {}

func (x *PeerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1alpha1_node_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerRequest.ProtoReflect.Descriptor instead.
func (*PeerRequest) Descriptor() ([]byte, []int) {
	return file_eth_v1alpha1_node_proto_rawDescGZIP(), []int{4}
}

func (x *PeerRequest) GetPeerId() string {
	if x != nil {
		return x.PeerId
	}
	return ""
}

type Peers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Peers []*Peer `protobuf:"bytes,1,rep,name=peers,proto3" json:"peers,omitempty"`
}

func (x *Peers) Reset() {
	*x = Peers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1alpha1_node_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Peers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Peers) ProtoMessage() {}

func (x *Peers) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1alpha1_node_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Peers.ProtoReflect.Descriptor instead.
func (*Peers) Descriptor() ([]byte, []int) {
	return file_eth_v1alpha1_node_proto_rawDescGZIP(), []int{5}
}

func (x *Peers) GetPeers() []*Peer {
	if x != nil {
		return x.Peers
	}
	return nil
}

type Peer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address         string          `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Direction       PeerDirection   `protobuf:"varint,2,opt,name=direction,proto3,enum=ethereum.eth.v1alpha1.PeerDirection" json:"direction,omitempty"`
	ConnectionState ConnectionState `protobuf:"varint,3,opt,name=connection_state,json=connectionState,proto3,enum=ethereum.eth.v1alpha1.ConnectionState" json:"connection_state,omitempty"`
	PeerId          string          `protobuf:"bytes,4,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	Enr             string          `protobuf:"bytes,5,opt,name=enr,proto3" json:"enr,omitempty"`
}

func (x *Peer) Reset() {
	*x = Peer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1alpha1_node_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Peer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Peer) ProtoMessage() {}

func (x *Peer) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1alpha1_node_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Peer.ProtoReflect.Descriptor instead.
func (*Peer) Descriptor() ([]byte, []int) {
	return file_eth_v1alpha1_node_proto_rawDescGZIP(), []int{6}
}

func (x *Peer) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Peer) GetDirection() PeerDirection {
	if x != nil {
		return x.Direction
	}
	return PeerDirection_UNKNOWN
}

func (x *Peer) GetConnectionState() ConnectionState {
	if x != nil {
		return x.ConnectionState
	}
	return ConnectionState_DISCONNECTED
}

func (x *Peer) GetPeerId() string {
	if x != nil {
		return x.PeerId
	}
	return ""
}

func (x *Peer) GetEnr() string {
	if x != nil {
		return x.Enr
	}
	return ""
}

type HostData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addresses []string `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
	PeerId    string   `protobuf:"bytes,2,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	Enr       string   `protobuf:"bytes,3,opt,name=enr,proto3" json:"enr,omitempty"`
}

func (x *HostData) Reset() {
	*x = HostData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1alpha1_node_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostData) ProtoMessage() {}

func (x *HostData) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1alpha1_node_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostData.ProtoReflect.Descriptor instead.
func (*HostData) Descriptor() ([]byte, []int) {
	return file_eth_v1alpha1_node_proto_rawDescGZIP(), []int{7}
}

func (x *HostData) GetAddresses() []string {
	if x != nil {
		return x.Addresses
	}
	return nil
}

func (x *HostData) GetPeerId() string {
	if x != nil {
		return x.PeerId
	}
	return ""
}

func (x *HostData) GetEnr() string {
	if x != nil {
		return x.Enr
	}
	return ""
}

var File_eth_v1alpha1_node_proto protoreflect.FileDescriptor

var file_eth_v1alpha1_node_proto_rawDesc = []byte{
	0x0a, 0x17, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6e,
	0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x65, 0x74, 0x68, 0x65, 0x72,
	0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x65, 0x74,
	0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x0a, 0x53, 0x79, 0x6e, 0x63,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x79, 0x6e, 0x63, 0x69, 0x6e,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x79, 0x6e, 0x63, 0x69, 0x6e, 0x67,
	0x22, 0xc2, 0x01, 0x0a, 0x07, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x12, 0x3d, 0x0a, 0x0c,
	0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b,
	0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x18, 0x64,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x16, 0x64,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x3e, 0x0a, 0x17, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73,
	0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x5f, 0x72, 0x6f, 0x6f, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x33, 0x32, 0x52, 0x15,
	0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x73, 0x52, 0x6f, 0x6f, 0x74, 0x22, 0x3f, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x31, 0x0a, 0x13, 0x49, 0x6d, 0x70, 0x6c, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x26, 0x0a, 0x0b, 0x50, 0x65, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x65, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x3a, 0x0a, 0x05, 0x50, 0x65, 0x65, 0x72, 0x73, 0x12, 0x31, 0x0a, 0x05, 0x70, 0x65,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x74, 0x68, 0x65,
	0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x22, 0xe2, 0x01,
	0x0a, 0x04, 0x50, 0x65, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x42, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65,
	0x74, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x51, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26,
	0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x65, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65,
	0x6e, 0x72, 0x22, 0x53, 0x0a, 0x08, 0x48, 0x6f, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x17, 0x0a, 0x07,
	0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70,
	0x65, 0x65, 0x72, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x72, 0x2a, 0x37, 0x0a, 0x0d, 0x50, 0x65, 0x65, 0x72, 0x44,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x42, 0x4f, 0x55, 0x4e, 0x44,
	0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x55, 0x54, 0x42, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x02,
	0x2a, 0x55, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43,
	0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x4e, 0x4e,
	0x45, 0x43, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4e, 0x4e,
	0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4f, 0x4e, 0x4e, 0x45,
	0x43, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x32, 0x85, 0x06, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65,
	0x12, 0x6e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x21, 0x2e, 0x65, 0x74, 0x68, 0x65,
	0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x22, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x69, 0x6e, 0x67,
	0x12, 0x68, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1e, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75,
	0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47,
	0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a,
	0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6e, 0x6f,
	0x64, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x12, 0x68, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x1e, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x82, 0x01, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6d, 0x70,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2a, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72,
	0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x49, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x65,
	0x74, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x62, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x48, 0x6f, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1f, 0x2e, 0x65,
	0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x22, 0x1e, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x32, 0x70, 0x12, 0x6b, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x50, 0x65, 0x65, 0x72, 0x12, 0x22, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72,
	0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x65,
	0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x19, 0x12, 0x17, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x12, 0x63, 0x0a, 0x09, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x65, 0x65, 0x72, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x1c, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x73, 0x22, 0x20, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x73, 0x42,
	0x90, 0x01, 0x0a, 0x19, 0x6f, 0x72, 0x67, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d,
	0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x09, 0x4e,
	0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63,
	0x6c, 0x61, 0x62, 0x73, 0x2f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x65,
	0x74, 0x68, 0xaa, 0x02, 0x15, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x45, 0x74,
	0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x15, 0x45, 0x74, 0x68,
	0x65, 0x72, 0x65, 0x75, 0x6d, 0x5c, 0x45, 0x74, 0x68, 0x5c, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eth_v1alpha1_node_proto_rawDescOnce sync.Once
	file_eth_v1alpha1_node_proto_rawDescData = file_eth_v1alpha1_node_proto_rawDesc
)

func file_eth_v1alpha1_node_proto_rawDescGZIP() []byte {
	file_eth_v1alpha1_node_proto_rawDescOnce.Do(func() {
		file_eth_v1alpha1_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_eth_v1alpha1_node_proto_rawDescData)
	})
	return file_eth_v1alpha1_node_proto_rawDescData
}

var file_eth_v1alpha1_node_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_eth_v1alpha1_node_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_eth_v1alpha1_node_proto_goTypes = []interface{}{
	(PeerDirection)(0),          // 0: ethereum.eth.v1alpha1.PeerDirection
	(ConnectionState)(0),        // 1: ethereum.eth.v1alpha1.ConnectionState
	(*SyncStatus)(nil),          // 2: ethereum.eth.v1alpha1.SyncStatus
	(*Genesis)(nil),             // 3: ethereum.eth.v1alpha1.Genesis
	(*Version)(nil),             // 4: ethereum.eth.v1alpha1.Version
	(*ImplementedServices)(nil), // 5: ethereum.eth.v1alpha1.ImplementedServices
	(*PeerRequest)(nil),         // 6: ethereum.eth.v1alpha1.PeerRequest
	(*Peers)(nil),               // 7: ethereum.eth.v1alpha1.Peers
	(*Peer)(nil),                // 8: ethereum.eth.v1alpha1.Peer
	(*HostData)(nil),            // 9: ethereum.eth.v1alpha1.HostData
	(*timestamp.Timestamp)(nil), // 10: google.protobuf.Timestamp
	(*empty.Empty)(nil),         // 11: google.protobuf.Empty
}
var file_eth_v1alpha1_node_proto_depIdxs = []int32{
	10, // 0: ethereum.eth.v1alpha1.Genesis.genesis_time:type_name -> google.protobuf.Timestamp
	8,  // 1: ethereum.eth.v1alpha1.Peers.peers:type_name -> ethereum.eth.v1alpha1.Peer
	0,  // 2: ethereum.eth.v1alpha1.Peer.direction:type_name -> ethereum.eth.v1alpha1.PeerDirection
	1,  // 3: ethereum.eth.v1alpha1.Peer.connection_state:type_name -> ethereum.eth.v1alpha1.ConnectionState
	11, // 4: ethereum.eth.v1alpha1.Node.GetSyncStatus:input_type -> google.protobuf.Empty
	11, // 5: ethereum.eth.v1alpha1.Node.GetGenesis:input_type -> google.protobuf.Empty
	11, // 6: ethereum.eth.v1alpha1.Node.GetVersion:input_type -> google.protobuf.Empty
	11, // 7: ethereum.eth.v1alpha1.Node.ListImplementedServices:input_type -> google.protobuf.Empty
	11, // 8: ethereum.eth.v1alpha1.Node.GetHost:input_type -> google.protobuf.Empty
	6,  // 9: ethereum.eth.v1alpha1.Node.GetPeer:input_type -> ethereum.eth.v1alpha1.PeerRequest
	11, // 10: ethereum.eth.v1alpha1.Node.ListPeers:input_type -> google.protobuf.Empty
	2,  // 11: ethereum.eth.v1alpha1.Node.GetSyncStatus:output_type -> ethereum.eth.v1alpha1.SyncStatus
	3,  // 12: ethereum.eth.v1alpha1.Node.GetGenesis:output_type -> ethereum.eth.v1alpha1.Genesis
	4,  // 13: ethereum.eth.v1alpha1.Node.GetVersion:output_type -> ethereum.eth.v1alpha1.Version
	5,  // 14: ethereum.eth.v1alpha1.Node.ListImplementedServices:output_type -> ethereum.eth.v1alpha1.ImplementedServices
	9,  // 15: ethereum.eth.v1alpha1.Node.GetHost:output_type -> ethereum.eth.v1alpha1.HostData
	8,  // 16: ethereum.eth.v1alpha1.Node.GetPeer:output_type -> ethereum.eth.v1alpha1.Peer
	7,  // 17: ethereum.eth.v1alpha1.Node.ListPeers:output_type -> ethereum.eth.v1alpha1.Peers
	11, // [11:18] is the sub-list for method output_type
	4,  // [4:11] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_eth_v1alpha1_node_proto_init() }
func file_eth_v1alpha1_node_proto_init() {
	if File_eth_v1alpha1_node_proto != nil {
		return
	}
	file_eth_v1alpha1_options_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_eth_v1alpha1_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncStatus); i {
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
		file_eth_v1alpha1_node_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Genesis); i {
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
		file_eth_v1alpha1_node_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Version); i {
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
		file_eth_v1alpha1_node_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImplementedServices); i {
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
		file_eth_v1alpha1_node_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerRequest); i {
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
		file_eth_v1alpha1_node_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Peers); i {
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
		file_eth_v1alpha1_node_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Peer); i {
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
		file_eth_v1alpha1_node_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostData); i {
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
			RawDescriptor: file_eth_v1alpha1_node_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eth_v1alpha1_node_proto_goTypes,
		DependencyIndexes: file_eth_v1alpha1_node_proto_depIdxs,
		EnumInfos:         file_eth_v1alpha1_node_proto_enumTypes,
		MessageInfos:      file_eth_v1alpha1_node_proto_msgTypes,
	}.Build()
	File_eth_v1alpha1_node_proto = out.File
	file_eth_v1alpha1_node_proto_rawDesc = nil
	file_eth_v1alpha1_node_proto_goTypes = nil
	file_eth_v1alpha1_node_proto_depIdxs = nil
}
