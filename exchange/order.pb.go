// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: lotus/exchange/order.proto

package exchange

import (
	_ "github.com/trylotus/go-lotus-proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OrderType int32

const (
	OrderType_ORDER_TYPE_UNSPECIFIED OrderType = 0
	OrderType_ORDER_TYPE_LIMIT       OrderType = 1
	OrderType_ORDER_TYPE_MARKET      OrderType = 2
)

// Enum value maps for OrderType.
var (
	OrderType_name = map[int32]string{
		0: "ORDER_TYPE_UNSPECIFIED",
		1: "ORDER_TYPE_LIMIT",
		2: "ORDER_TYPE_MARKET",
	}
	OrderType_value = map[string]int32{
		"ORDER_TYPE_UNSPECIFIED": 0,
		"ORDER_TYPE_LIMIT":       1,
		"ORDER_TYPE_MARKET":      2,
	}
)

func (x OrderType) Enum() *OrderType {
	p := new(OrderType)
	*p = x
	return p
}

func (x OrderType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderType) Descriptor() protoreflect.EnumDescriptor {
	return file_lotus_exchange_order_proto_enumTypes[0].Descriptor()
}

func (OrderType) Type() protoreflect.EnumType {
	return &file_lotus_exchange_order_proto_enumTypes[0]
}

func (x OrderType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderType.Descriptor instead.
func (OrderType) EnumDescriptor() ([]byte, []int) {
	return file_lotus_exchange_order_proto_rawDescGZIP(), []int{0}
}

type OrderStatus int32

const (
	OrderStatus_ORDER_STATUS_UNSPECIFIED OrderStatus = 0
	OrderStatus_ORDER_STATUS_OPEN        OrderStatus = 1
	OrderStatus_ORDER_STATUS_FILLED      OrderStatus = 2
	OrderStatus_ORDER_STATUS_CANCELLED   OrderStatus = 3
)

// Enum value maps for OrderStatus.
var (
	OrderStatus_name = map[int32]string{
		0: "ORDER_STATUS_UNSPECIFIED",
		1: "ORDER_STATUS_OPEN",
		2: "ORDER_STATUS_FILLED",
		3: "ORDER_STATUS_CANCELLED",
	}
	OrderStatus_value = map[string]int32{
		"ORDER_STATUS_UNSPECIFIED": 0,
		"ORDER_STATUS_OPEN":        1,
		"ORDER_STATUS_FILLED":      2,
		"ORDER_STATUS_CANCELLED":   3,
	}
)

func (x OrderStatus) Enum() *OrderStatus {
	p := new(OrderStatus)
	*p = x
	return p
}

func (x OrderStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_lotus_exchange_order_proto_enumTypes[1].Descriptor()
}

func (OrderStatus) Type() protoreflect.EnumType {
	return &file_lotus_exchange_order_proto_enumTypes[1]
}

func (x OrderStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderStatus.Descriptor instead.
func (OrderStatus) EnumDescriptor() ([]byte, []int) {
	return file_lotus_exchange_order_proto_rawDescGZIP(), []int{1}
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp    *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	MakerAddress []byte                 `protobuf:"bytes,2,opt,name=maker_address,json=makerAddress,proto3" json:"maker_address,omitempty"`
	OrderId      string                 `protobuf:"bytes,3,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	TokenAddress []byte                 `protobuf:"bytes,4,opt,name=token_address,json=tokenAddress,proto3" json:"token_address,omitempty"`
	Amount       string                 `protobuf:"bytes,5,opt,name=amount,proto3" json:"amount,omitempty"`
	Price        string                 `protobuf:"bytes,6,opt,name=price,proto3" json:"price,omitempty"`
	OrderType    OrderType              `protobuf:"varint,7,opt,name=order_type,json=orderType,proto3,enum=lotus.exchange.OrderType" json:"order_type,omitempty"`
	OrderStatus  OrderStatus            `protobuf:"varint,8,opt,name=order_status,json=orderStatus,proto3,enum=lotus.exchange.OrderStatus" json:"order_status,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lotus_exchange_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_lotus_exchange_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_lotus_exchange_order_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Order) GetMakerAddress() []byte {
	if x != nil {
		return x.MakerAddress
	}
	return nil
}

func (x *Order) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *Order) GetTokenAddress() []byte {
	if x != nil {
		return x.TokenAddress
	}
	return nil
}

func (x *Order) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *Order) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *Order) GetOrderType() OrderType {
	if x != nil {
		return x.OrderType
	}
	return OrderType_ORDER_TYPE_UNSPECIFIED
}

func (x *Order) GetOrderStatus() OrderStatus {
	if x != nil {
		return x.OrderStatus
	}
	return OrderStatus_ORDER_STATUS_UNSPECIFIED
}

var File_lotus_exchange_order_proto protoreflect.FileDescriptor

var file_lotus_exchange_order_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6c, 0x6f, 0x74, 0x75, 0x73, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6c, 0x6f,
	0x74, 0x75, 0x73, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6c,
	0x6f, 0x74, 0x75, 0x73, 0x2f, 0x6c, 0x6f, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xea, 0x02, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x2b, 0x0a, 0x0d, 0x6d, 0x61, 0x6b, 0x65, 0x72, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0xd2, 0xab, 0x30,
	0x02, 0x08, 0x01, 0x52, 0x0c, 0x6d, 0x61, 0x6b, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x0d,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x42, 0x06, 0xd2, 0xab, 0x30, 0x02, 0x08, 0x01, 0x52, 0x0c, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xda, 0xab, 0x30, 0x00, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xda, 0xab, 0x30, 0x00, 0x52, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x6c, 0x6f, 0x74, 0x75, 0x73, 0x2e,
	0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3e, 0x0a,
	0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x6c, 0x6f, 0x74, 0x75, 0x73, 0x2e, 0x65, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x54, 0x0a,
	0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x4f, 0x52,
	0x44, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11,
	0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x41, 0x52, 0x4b, 0x45,
	0x54, 0x10, 0x02, 0x2a, 0x77, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1c, 0x0a, 0x18, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x15, 0x0a, 0x11, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x4f, 0x52, 0x44, 0x45, 0x52,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x49, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x02,
	0x12, 0x1a, 0x0a, 0x16, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x42, 0x2d, 0x5a, 0x2b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72, 0x79, 0x6c, 0x6f,
	0x74, 0x75, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x6c, 0x6f, 0x74, 0x75, 0x73, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_lotus_exchange_order_proto_rawDescOnce sync.Once
	file_lotus_exchange_order_proto_rawDescData = file_lotus_exchange_order_proto_rawDesc
)

func file_lotus_exchange_order_proto_rawDescGZIP() []byte {
	file_lotus_exchange_order_proto_rawDescOnce.Do(func() {
		file_lotus_exchange_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_lotus_exchange_order_proto_rawDescData)
	})
	return file_lotus_exchange_order_proto_rawDescData
}

var file_lotus_exchange_order_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_lotus_exchange_order_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_lotus_exchange_order_proto_goTypes = []interface{}{
	(OrderType)(0),                // 0: lotus.exchange.OrderType
	(OrderStatus)(0),              // 1: lotus.exchange.OrderStatus
	(*Order)(nil),                 // 2: lotus.exchange.Order
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_lotus_exchange_order_proto_depIdxs = []int32{
	3, // 0: lotus.exchange.Order.timestamp:type_name -> google.protobuf.Timestamp
	0, // 1: lotus.exchange.Order.order_type:type_name -> lotus.exchange.OrderType
	1, // 2: lotus.exchange.Order.order_status:type_name -> lotus.exchange.OrderStatus
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_lotus_exchange_order_proto_init() }
func file_lotus_exchange_order_proto_init() {
	if File_lotus_exchange_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lotus_exchange_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
			RawDescriptor: file_lotus_exchange_order_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lotus_exchange_order_proto_goTypes,
		DependencyIndexes: file_lotus_exchange_order_proto_depIdxs,
		EnumInfos:         file_lotus_exchange_order_proto_enumTypes,
		MessageInfos:      file_lotus_exchange_order_proto_msgTypes,
	}.Build()
	File_lotus_exchange_order_proto = out.File
	file_lotus_exchange_order_proto_rawDesc = nil
	file_lotus_exchange_order_proto_goTypes = nil
	file_lotus_exchange_order_proto_depIdxs = nil
}
