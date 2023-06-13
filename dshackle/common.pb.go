// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: common.proto

package dshackle

import (
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

type ChainRef int32

const (
	ChainRef_CHAIN_UNSPECIFIED      ChainRef = 0
	ChainRef_CHAIN_BITCOIN          ChainRef = 1 // CHAIN_GRIN = 2;
	ChainRef_CHAIN_ETHEREUM         ChainRef = 100
	ChainRef_CHAIN_ETHEREUM_CLASSIC ChainRef = 101
	ChainRef_CHAIN_FANTOM           ChainRef = 102 // Fantom, https://fantom.foundation/
	// Sidechains and state channels start with 1_000
	// CHAIN_LIGHTNING = 1001;
	ChainRef_CHAIN_MATIC    ChainRef = 1002 // Matic PoS Ethereum sidechain based on Polygon
	ChainRef_CHAIN_RSK      ChainRef = 1003 // RSK sidechain, https://www.rsk.co/
	ChainRef_CHAIN_ARBITRUM ChainRef = 1004 // Arbitrum (next generation Layer2 technologies for Ethereum), https://arbitrum.io
	ChainRef_CHAIN_OPTIMISM ChainRef = 1005 // Optimism (low-cost and lightning-fast Ethereum L2 blockchain), https://www.optimism.io
	ChainRef_CHAIN_BSC      ChainRef = 1006 // Binance Smart Chain, https://docs.bnbchain.org/docs/learn/intro
	// Testnets start with 10_000
	ChainRef_CHAIN_MORDEN          ChainRef = 10001
	ChainRef_CHAIN_KOVAN           ChainRef = 10002
	ChainRef_CHAIN_TESTNET_BITCOIN ChainRef = 10003
	// CHAIN_FLOONET = 10004;
	ChainRef_CHAIN_GOERLI           ChainRef = 10005
	ChainRef_CHAIN_ROPSTEN          ChainRef = 10006
	ChainRef_CHAIN_RINKEBY          ChainRef = 10007
	ChainRef_CHAIN_SEPOLIA          ChainRef = 10008
	ChainRef_CHAIN_TESTNET_ARBITRUM ChainRef = 10009
	ChainRef_CHAIN_TESTNET_OPTIMISM ChainRef = 10010
)

// Enum value maps for ChainRef.
var (
	ChainRef_name = map[int32]string{
		0:     "CHAIN_UNSPECIFIED",
		1:     "CHAIN_BITCOIN",
		100:   "CHAIN_ETHEREUM",
		101:   "CHAIN_ETHEREUM_CLASSIC",
		102:   "CHAIN_FANTOM",
		1002:  "CHAIN_MATIC",
		1003:  "CHAIN_RSK",
		1004:  "CHAIN_ARBITRUM",
		1005:  "CHAIN_OPTIMISM",
		1006:  "CHAIN_BSC",
		10001: "CHAIN_MORDEN",
		10002: "CHAIN_KOVAN",
		10003: "CHAIN_TESTNET_BITCOIN",
		10005: "CHAIN_GOERLI",
		10006: "CHAIN_ROPSTEN",
		10007: "CHAIN_RINKEBY",
		10008: "CHAIN_SEPOLIA",
		10009: "CHAIN_TESTNET_ARBITRUM",
		10010: "CHAIN_TESTNET_OPTIMISM",
	}
	ChainRef_value = map[string]int32{
		"CHAIN_UNSPECIFIED":      0,
		"CHAIN_BITCOIN":          1,
		"CHAIN_ETHEREUM":         100,
		"CHAIN_ETHEREUM_CLASSIC": 101,
		"CHAIN_FANTOM":           102,
		"CHAIN_MATIC":            1002,
		"CHAIN_RSK":              1003,
		"CHAIN_ARBITRUM":         1004,
		"CHAIN_OPTIMISM":         1005,
		"CHAIN_BSC":              1006,
		"CHAIN_MORDEN":           10001,
		"CHAIN_KOVAN":            10002,
		"CHAIN_TESTNET_BITCOIN":  10003,
		"CHAIN_GOERLI":           10005,
		"CHAIN_ROPSTEN":          10006,
		"CHAIN_RINKEBY":          10007,
		"CHAIN_SEPOLIA":          10008,
		"CHAIN_TESTNET_ARBITRUM": 10009,
		"CHAIN_TESTNET_OPTIMISM": 10010,
	}
)

func (x ChainRef) Enum() *ChainRef {
	p := new(ChainRef)
	*p = x
	return p
}

func (x ChainRef) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChainRef) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[0].Descriptor()
}

func (ChainRef) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[0]
}

func (x ChainRef) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChainRef.Descriptor instead.
func (ChainRef) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

type Chain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type ChainRef `protobuf:"varint,1,opt,name=type,proto3,enum=emerald.ChainRef" json:"type,omitempty"`
}

func (x *Chain) Reset() {
	*x = Chain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chain) ProtoMessage() {}

func (x *Chain) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chain.ProtoReflect.Descriptor instead.
func (*Chain) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *Chain) GetType() ChainRef {
	if x != nil {
		return x.Type
	}
	return ChainRef_CHAIN_UNSPECIFIED
}

type SingleAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *SingleAddress) Reset() {
	*x = SingleAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingleAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleAddress) ProtoMessage() {}

func (x *SingleAddress) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleAddress.ProtoReflect.Descriptor instead.
func (*SingleAddress) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (x *SingleAddress) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type XpubAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// base58 encoded string for xpub, https://en.bitcoin.it/wiki/BIP_0032#Serialization_format
	Xpub        string `protobuf:"bytes,1,opt,name=xpub,proto3" json:"xpub,omitempty"`
	Start       uint64 `protobuf:"varint,3,opt,name=start,proto3" json:"start,omitempty"`
	Limit       uint64 `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	UnusedLimit uint64 `protobuf:"varint,5,opt,name=unused_limit,json=unusedLimit,proto3" json:"unused_limit,omitempty"`
}

func (x *XpubAddress) Reset() {
	*x = XpubAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XpubAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XpubAddress) ProtoMessage() {}

func (x *XpubAddress) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XpubAddress.ProtoReflect.Descriptor instead.
func (*XpubAddress) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

func (x *XpubAddress) GetXpub() string {
	if x != nil {
		return x.Xpub
	}
	return ""
}

func (x *XpubAddress) GetStart() uint64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *XpubAddress) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *XpubAddress) GetUnusedLimit() uint64 {
	if x != nil {
		return x.UnusedLimit
	}
	return 0
}

type MultiAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addresses []*SingleAddress `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
}

func (x *MultiAddress) Reset() {
	*x = MultiAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiAddress) ProtoMessage() {}

func (x *MultiAddress) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiAddress.ProtoReflect.Descriptor instead.
func (*MultiAddress) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{3}
}

func (x *MultiAddress) GetAddresses() []*SingleAddress {
	if x != nil {
		return x.Addresses
	}
	return nil
}

type ReferenceAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Refid uint64 `protobuf:"varint,1,opt,name=refid,proto3" json:"refid,omitempty"`
}

func (x *ReferenceAddress) Reset() {
	*x = ReferenceAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReferenceAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReferenceAddress) ProtoMessage() {}

func (x *ReferenceAddress) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReferenceAddress.ProtoReflect.Descriptor instead.
func (*ReferenceAddress) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{4}
}

func (x *ReferenceAddress) GetRefid() uint64 {
	if x != nil {
		return x.Refid
	}
	return 0
}

type AnyAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to AddrType:
	//
	//	*AnyAddress_AddressSingle
	//	*AnyAddress_AddressMulti
	//	*AnyAddress_AddressXpub
	//	*AnyAddress_AddressRef
	AddrType isAnyAddress_AddrType `protobuf_oneof:"addr_type"`
}

func (x *AnyAddress) Reset() {
	*x = AnyAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnyAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnyAddress) ProtoMessage() {}

func (x *AnyAddress) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnyAddress.ProtoReflect.Descriptor instead.
func (*AnyAddress) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{5}
}

func (m *AnyAddress) GetAddrType() isAnyAddress_AddrType {
	if m != nil {
		return m.AddrType
	}
	return nil
}

func (x *AnyAddress) GetAddressSingle() *SingleAddress {
	if x, ok := x.GetAddrType().(*AnyAddress_AddressSingle); ok {
		return x.AddressSingle
	}
	return nil
}

func (x *AnyAddress) GetAddressMulti() *MultiAddress {
	if x, ok := x.GetAddrType().(*AnyAddress_AddressMulti); ok {
		return x.AddressMulti
	}
	return nil
}

func (x *AnyAddress) GetAddressXpub() *XpubAddress {
	if x, ok := x.GetAddrType().(*AnyAddress_AddressXpub); ok {
		return x.AddressXpub
	}
	return nil
}

func (x *AnyAddress) GetAddressRef() *ReferenceAddress {
	if x, ok := x.GetAddrType().(*AnyAddress_AddressRef); ok {
		return x.AddressRef
	}
	return nil
}

type isAnyAddress_AddrType interface {
	isAnyAddress_AddrType()
}

type AnyAddress_AddressSingle struct {
	AddressSingle *SingleAddress `protobuf:"bytes,1,opt,name=address_single,json=addressSingle,proto3,oneof"`
}

type AnyAddress_AddressMulti struct {
	AddressMulti *MultiAddress `protobuf:"bytes,2,opt,name=address_multi,json=addressMulti,proto3,oneof"`
}

type AnyAddress_AddressXpub struct {
	AddressXpub *XpubAddress `protobuf:"bytes,3,opt,name=address_xpub,json=addressXpub,proto3,oneof"`
}

type AnyAddress_AddressRef struct {
	AddressRef *ReferenceAddress `protobuf:"bytes,4,opt,name=address_ref,json=addressRef,proto3,oneof"`
}

func (*AnyAddress_AddressSingle) isAnyAddress_AddrType() {}

func (*AnyAddress_AddressMulti) isAnyAddress_AddrType() {}

func (*AnyAddress_AddressXpub) isAnyAddress_AddrType() {}

func (*AnyAddress_AddressRef) isAnyAddress_AddrType() {}

type Asset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chain ChainRef `protobuf:"varint,1,opt,name=chain,proto3,enum=emerald.ChainRef" json:"chain,omitempty"`
	Code  string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *Asset) Reset() {
	*x = Asset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Asset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Asset) ProtoMessage() {}

func (x *Asset) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Asset.ProtoReflect.Descriptor instead.
func (*Asset) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{6}
}

func (x *Asset) GetChain() ChainRef {
	if x != nil {
		return x.Chain
	}
	return ChainRef_CHAIN_UNSPECIFIED
}

func (x *Asset) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type BlockInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height    uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	BlockId   string `protobuf:"bytes,2,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	Timestamp uint64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *BlockInfo) Reset() {
	*x = BlockInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockInfo) ProtoMessage() {}

func (x *BlockInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockInfo.ProtoReflect.Descriptor instead.
func (*BlockInfo) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{7}
}

func (x *BlockInfo) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *BlockInfo) GetBlockId() string {
	if x != nil {
		return x.BlockId
	}
	return ""
}

func (x *BlockInfo) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x65, 0x6d, 0x65, 0x72, 0x61, 0x6c, 0x64, 0x22, 0x2e, 0x0a, 0x05, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x12, 0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11,
	0x2e, 0x65, 0x6d, 0x65, 0x72, 0x61, 0x6c, 0x64, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x65,
	0x66, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x29, 0x0a, 0x0d, 0x53, 0x69, 0x6e, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x22, 0x70, 0x0a, 0x0b, 0x58, 0x70, 0x75, 0x62, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x78, 0x70, 0x75, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x78, 0x70, 0x75, 0x62, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x6e, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x75, 0x6e, 0x75, 0x73, 0x65, 0x64, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x22, 0x44, 0x0a, 0x0c, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x34, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6d, 0x65, 0x72, 0x61, 0x6c,
	0x64, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x22, 0x28, 0x0a, 0x10, 0x52, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x72, 0x65, 0x66, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x72,
	0x65, 0x66, 0x69, 0x64, 0x22, 0x91, 0x02, 0x0a, 0x0a, 0x41, 0x6e, 0x79, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x3f, 0x0a, 0x0e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x73,
	0x69, 0x6e, 0x67, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6d,
	0x65, 0x72, 0x61, 0x6c, 0x64, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x48, 0x00, 0x52, 0x0d, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x12, 0x3c, 0x0a, 0x0d, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x6d,
	0x65, 0x72, 0x61, 0x6c, 0x64, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x48, 0x00, 0x52, 0x0c, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4d, 0x75, 0x6c,
	0x74, 0x69, 0x12, 0x39, 0x0a, 0x0c, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x78, 0x70,
	0x75, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x6d, 0x65, 0x72, 0x61,
	0x6c, 0x64, 0x2e, 0x58, 0x70, 0x75, 0x62, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x48, 0x00,
	0x52, 0x0b, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x58, 0x70, 0x75, 0x62, 0x12, 0x3c, 0x0a,
	0x0b, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x6d, 0x65, 0x72, 0x61, 0x6c, 0x64, 0x2e, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x48, 0x00, 0x52,
	0x0a, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x66, 0x42, 0x0b, 0x0a, 0x09, 0x61,
	0x64, 0x64, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x44, 0x0a, 0x05, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x12, 0x27, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x11, 0x2e, 0x65, 0x6d, 0x65, 0x72, 0x61, 0x6c, 0x64, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x52, 0x65, 0x66, 0x52, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x5c,
	0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2a, 0x9c, 0x03, 0x0a,
	0x08, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x66, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x48, 0x41,
	0x49, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x11, 0x0a, 0x0d, 0x43, 0x48, 0x41, 0x49, 0x4e, 0x5f, 0x42, 0x49, 0x54, 0x43, 0x4f, 0x49,
	0x4e, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x48, 0x41, 0x49, 0x4e, 0x5f, 0x45, 0x54, 0x48,
	0x45, 0x52, 0x45, 0x55, 0x4d, 0x10, 0x64, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x48, 0x41, 0x49, 0x4e,
	0x5f, 0x45, 0x54, 0x48, 0x45, 0x52, 0x45, 0x55, 0x4d, 0x5f, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x49,
	0x43, 0x10, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x48, 0x41, 0x49, 0x4e, 0x5f, 0x46, 0x41, 0x4e,
	0x54, 0x4f, 0x4d, 0x10, 0x66, 0x12, 0x10, 0x0a, 0x0b, 0x43, 0x48, 0x41, 0x49, 0x4e, 0x5f, 0x4d,
	0x41, 0x54, 0x49, 0x43, 0x10, 0xea, 0x07, 0x12, 0x0e, 0x0a, 0x09, 0x43, 0x48, 0x41, 0x49, 0x4e,
	0x5f, 0x52, 0x53, 0x4b, 0x10, 0xeb, 0x07, 0x12, 0x13, 0x0a, 0x0e, 0x43, 0x48, 0x41, 0x49, 0x4e,
	0x5f, 0x41, 0x52, 0x42, 0x49, 0x54, 0x52, 0x55, 0x4d, 0x10, 0xec, 0x07, 0x12, 0x13, 0x0a, 0x0e,
	0x43, 0x48, 0x41, 0x49, 0x4e, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4d, 0x49, 0x53, 0x4d, 0x10, 0xed,
	0x07, 0x12, 0x0e, 0x0a, 0x09, 0x43, 0x48, 0x41, 0x49, 0x4e, 0x5f, 0x42, 0x53, 0x43, 0x10, 0xee,
	0x07, 0x12, 0x11, 0x0a, 0x0c, 0x43, 0x48, 0x41, 0x49, 0x4e, 0x5f, 0x4d, 0x4f, 0x52, 0x44, 0x45,
	0x4e, 0x10, 0x91, 0x4e, 0x12, 0x10, 0x0a, 0x0b, 0x43, 0x48, 0x41, 0x49, 0x4e, 0x5f, 0x4b, 0x4f,
	0x56, 0x41, 0x4e, 0x10, 0x92, 0x4e, 0x12, 0x1a, 0x0a, 0x15, 0x43, 0x48, 0x41, 0x49, 0x4e, 0x5f,
	0x54, 0x45, 0x53, 0x54, 0x4e, 0x45, 0x54, 0x5f, 0x42, 0x49, 0x54, 0x43, 0x4f, 0x49, 0x4e, 0x10,
	0x93, 0x4e, 0x12, 0x11, 0x0a, 0x0c, 0x43, 0x48, 0x41, 0x49, 0x4e, 0x5f, 0x47, 0x4f, 0x45, 0x52,
	0x4c, 0x49, 0x10, 0x95, 0x4e, 0x12, 0x12, 0x0a, 0x0d, 0x43, 0x48, 0x41, 0x49, 0x4e, 0x5f, 0x52,
	0x4f, 0x50, 0x53, 0x54, 0x45, 0x4e, 0x10, 0x96, 0x4e, 0x12, 0x12, 0x0a, 0x0d, 0x43, 0x48, 0x41,
	0x49, 0x4e, 0x5f, 0x52, 0x49, 0x4e, 0x4b, 0x45, 0x42, 0x59, 0x10, 0x97, 0x4e, 0x12, 0x12, 0x0a,
	0x0d, 0x43, 0x48, 0x41, 0x49, 0x4e, 0x5f, 0x53, 0x45, 0x50, 0x4f, 0x4c, 0x49, 0x41, 0x10, 0x98,
	0x4e, 0x12, 0x1b, 0x0a, 0x16, 0x43, 0x48, 0x41, 0x49, 0x4e, 0x5f, 0x54, 0x45, 0x53, 0x54, 0x4e,
	0x45, 0x54, 0x5f, 0x41, 0x52, 0x42, 0x49, 0x54, 0x52, 0x55, 0x4d, 0x10, 0x99, 0x4e, 0x12, 0x1b,
	0x0a, 0x16, 0x43, 0x48, 0x41, 0x49, 0x4e, 0x5f, 0x54, 0x45, 0x53, 0x54, 0x4e, 0x45, 0x54, 0x5f,
	0x4f, 0x50, 0x54, 0x49, 0x4d, 0x49, 0x53, 0x4d, 0x10, 0x9a, 0x4e, 0x42, 0x19, 0x0a, 0x17, 0x69,
	0x6f, 0x2e, 0x65, 0x6d, 0x65, 0x72, 0x61, 0x6c, 0x64, 0x70, 0x61, 0x79, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_common_proto_goTypes = []interface{}{
	(ChainRef)(0),            // 0: emerald.ChainRef
	(*Chain)(nil),            // 1: emerald.Chain
	(*SingleAddress)(nil),    // 2: emerald.SingleAddress
	(*XpubAddress)(nil),      // 3: emerald.XpubAddress
	(*MultiAddress)(nil),     // 4: emerald.MultiAddress
	(*ReferenceAddress)(nil), // 5: emerald.ReferenceAddress
	(*AnyAddress)(nil),       // 6: emerald.AnyAddress
	(*Asset)(nil),            // 7: emerald.Asset
	(*BlockInfo)(nil),        // 8: emerald.BlockInfo
}
var file_common_proto_depIdxs = []int32{
	0, // 0: emerald.Chain.type:type_name -> emerald.ChainRef
	2, // 1: emerald.MultiAddress.addresses:type_name -> emerald.SingleAddress
	2, // 2: emerald.AnyAddress.address_single:type_name -> emerald.SingleAddress
	4, // 3: emerald.AnyAddress.address_multi:type_name -> emerald.MultiAddress
	3, // 4: emerald.AnyAddress.address_xpub:type_name -> emerald.XpubAddress
	5, // 5: emerald.AnyAddress.address_ref:type_name -> emerald.ReferenceAddress
	0, // 6: emerald.Asset.chain:type_name -> emerald.ChainRef
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chain); i {
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
		file_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SingleAddress); i {
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
		file_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XpubAddress); i {
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
		file_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiAddress); i {
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
		file_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReferenceAddress); i {
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
		file_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnyAddress); i {
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
		file_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Asset); i {
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
		file_common_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockInfo); i {
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
	file_common_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*AnyAddress_AddressSingle)(nil),
		(*AnyAddress_AddressMulti)(nil),
		(*AnyAddress_AddressXpub)(nil),
		(*AnyAddress_AddressRef)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		EnumInfos:         file_common_proto_enumTypes,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}