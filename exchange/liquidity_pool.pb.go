// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: lotus/exchange/liquidity_pool.proto

package exchange

import (
	_ "github.com/trylotus/go-lotus-proto"
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

type LiquidityPool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PoolId         string   `protobuf:"bytes,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	TokenAddresses [][]byte `protobuf:"bytes,2,rep,name=token_addresses,json=tokenAddresses,proto3" json:"token_addresses,omitempty"`
	TokenReserves  [][]byte `protobuf:"bytes,3,rep,name=token_reserves,json=tokenReserves,proto3" json:"token_reserves,omitempty"`
	LpTokenAddress []byte   `protobuf:"bytes,4,opt,name=lp_token_address,json=lpTokenAddress,proto3" json:"lp_token_address,omitempty"`
	LpTokenSupply  string   `protobuf:"bytes,5,opt,name=lp_token_supply,json=lpTokenSupply,proto3" json:"lp_token_supply,omitempty"`
	PoolFee        string   `protobuf:"bytes,6,opt,name=pool_fee,json=poolFee,proto3" json:"pool_fee,omitempty"`
}

func (x *LiquidityPool) Reset() {
	*x = LiquidityPool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lotus_exchange_liquidity_pool_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LiquidityPool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LiquidityPool) ProtoMessage() {}

func (x *LiquidityPool) ProtoReflect() protoreflect.Message {
	mi := &file_lotus_exchange_liquidity_pool_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LiquidityPool.ProtoReflect.Descriptor instead.
func (*LiquidityPool) Descriptor() ([]byte, []int) {
	return file_lotus_exchange_liquidity_pool_proto_rawDescGZIP(), []int{0}
}

func (x *LiquidityPool) GetPoolId() string {
	if x != nil {
		return x.PoolId
	}
	return ""
}

func (x *LiquidityPool) GetTokenAddresses() [][]byte {
	if x != nil {
		return x.TokenAddresses
	}
	return nil
}

func (x *LiquidityPool) GetTokenReserves() [][]byte {
	if x != nil {
		return x.TokenReserves
	}
	return nil
}

func (x *LiquidityPool) GetLpTokenAddress() []byte {
	if x != nil {
		return x.LpTokenAddress
	}
	return nil
}

func (x *LiquidityPool) GetLpTokenSupply() string {
	if x != nil {
		return x.LpTokenSupply
	}
	return ""
}

func (x *LiquidityPool) GetPoolFee() string {
	if x != nil {
		return x.PoolFee
	}
	return ""
}

var File_lotus_exchange_liquidity_pool_proto protoreflect.FileDescriptor

var file_lotus_exchange_liquidity_pool_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6c, 0x6f, 0x74, 0x75, 0x73, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x2f, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6c, 0x6f, 0x74, 0x75, 0x73, 0x2e, 0x65, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x1a, 0x11, 0x6c, 0x6f, 0x74, 0x75, 0x73, 0x2f, 0x6c, 0x6f, 0x74,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x02, 0x0a, 0x0d, 0x4c, 0x69, 0x71,
	0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f,
	0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x6f,
	0x6c, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x0f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x42, 0x06, 0xd2, 0xab,
	0x30, 0x02, 0x08, 0x01, 0x52, 0x0e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x0e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x72, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0c, 0x42, 0x04, 0xda, 0xab,
	0x30, 0x00, 0x52, 0x0d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x73, 0x12, 0x30, 0x0a, 0x10, 0x6c, 0x70, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0xd2, 0xab, 0x30,
	0x02, 0x08, 0x01, 0x52, 0x0e, 0x6c, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x2c, 0x0a, 0x0f, 0x6c, 0x70, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f,
	0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xda, 0xab,
	0x30, 0x00, 0x52, 0x0d, 0x6c, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x75, 0x70, 0x70, 0x6c,
	0x79, 0x12, 0x1f, 0x0a, 0x08, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x04, 0xda, 0xab, 0x30, 0x00, 0x52, 0x07, 0x70, 0x6f, 0x6f, 0x6c, 0x46,
	0x65, 0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x74, 0x72, 0x79, 0x6c, 0x6f, 0x74, 0x75, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x6c, 0x6f, 0x74,
	0x75, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lotus_exchange_liquidity_pool_proto_rawDescOnce sync.Once
	file_lotus_exchange_liquidity_pool_proto_rawDescData = file_lotus_exchange_liquidity_pool_proto_rawDesc
)

func file_lotus_exchange_liquidity_pool_proto_rawDescGZIP() []byte {
	file_lotus_exchange_liquidity_pool_proto_rawDescOnce.Do(func() {
		file_lotus_exchange_liquidity_pool_proto_rawDescData = protoimpl.X.CompressGZIP(file_lotus_exchange_liquidity_pool_proto_rawDescData)
	})
	return file_lotus_exchange_liquidity_pool_proto_rawDescData
}

var file_lotus_exchange_liquidity_pool_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_lotus_exchange_liquidity_pool_proto_goTypes = []interface{}{
	(*LiquidityPool)(nil), // 0: lotus.exchange.LiquidityPool
}
var file_lotus_exchange_liquidity_pool_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_lotus_exchange_liquidity_pool_proto_init() }
func file_lotus_exchange_liquidity_pool_proto_init() {
	if File_lotus_exchange_liquidity_pool_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lotus_exchange_liquidity_pool_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LiquidityPool); i {
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
			RawDescriptor: file_lotus_exchange_liquidity_pool_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lotus_exchange_liquidity_pool_proto_goTypes,
		DependencyIndexes: file_lotus_exchange_liquidity_pool_proto_depIdxs,
		MessageInfos:      file_lotus_exchange_liquidity_pool_proto_msgTypes,
	}.Build()
	File_lotus_exchange_liquidity_pool_proto = out.File
	file_lotus_exchange_liquidity_pool_proto_rawDesc = nil
	file_lotus_exchange_liquidity_pool_proto_goTypes = nil
	file_lotus_exchange_liquidity_pool_proto_depIdxs = nil
}