// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: lotus/evm/erc721.proto

package evm

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

type ERC721 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Symbol  string `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	BaseUri string `protobuf:"bytes,4,opt,name=base_uri,json=baseUri,proto3" json:"base_uri,omitempty"`
}

func (x *ERC721) Reset() {
	*x = ERC721{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lotus_evm_erc721_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ERC721) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ERC721) ProtoMessage() {}

func (x *ERC721) ProtoReflect() protoreflect.Message {
	mi := &file_lotus_evm_erc721_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ERC721.ProtoReflect.Descriptor instead.
func (*ERC721) Descriptor() ([]byte, []int) {
	return file_lotus_evm_erc721_proto_rawDescGZIP(), []int{0}
}

func (x *ERC721) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *ERC721) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ERC721) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *ERC721) GetBaseUri() string {
	if x != nil {
		return x.BaseUri
	}
	return ""
}

var File_lotus_evm_erc721_proto protoreflect.FileDescriptor

var file_lotus_evm_erc721_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6c, 0x6f, 0x74, 0x75, 0x73, 0x2f, 0x65, 0x76, 0x6d, 0x2f, 0x65, 0x72, 0x63, 0x37,
	0x32, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6c, 0x6f, 0x74, 0x75, 0x73, 0x2e,
	0x65, 0x76, 0x6d, 0x1a, 0x11, 0x6c, 0x6f, 0x74, 0x75, 0x73, 0x2f, 0x6c, 0x6f, 0x74, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x71, 0x0a, 0x06, 0x45, 0x52, 0x43, 0x37, 0x32, 0x31,
	0x12, 0x20, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x42, 0x06, 0xd2, 0xab, 0x30, 0x02, 0x08, 0x01, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x19,
	0x0a, 0x08, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x62, 0x61, 0x73, 0x65, 0x55, 0x72, 0x69, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72, 0x79, 0x6c, 0x6f, 0x74, 0x75, 0x73,
	0x2f, 0x67, 0x6f, 0x2d, 0x6c, 0x6f, 0x74, 0x75, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x65, 0x76, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lotus_evm_erc721_proto_rawDescOnce sync.Once
	file_lotus_evm_erc721_proto_rawDescData = file_lotus_evm_erc721_proto_rawDesc
)

func file_lotus_evm_erc721_proto_rawDescGZIP() []byte {
	file_lotus_evm_erc721_proto_rawDescOnce.Do(func() {
		file_lotus_evm_erc721_proto_rawDescData = protoimpl.X.CompressGZIP(file_lotus_evm_erc721_proto_rawDescData)
	})
	return file_lotus_evm_erc721_proto_rawDescData
}

var file_lotus_evm_erc721_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_lotus_evm_erc721_proto_goTypes = []interface{}{
	(*ERC721)(nil), // 0: lotus.evm.ERC721
}
var file_lotus_evm_erc721_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_lotus_evm_erc721_proto_init() }
func file_lotus_evm_erc721_proto_init() {
	if File_lotus_evm_erc721_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lotus_evm_erc721_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ERC721); i {
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
			RawDescriptor: file_lotus_evm_erc721_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lotus_evm_erc721_proto_goTypes,
		DependencyIndexes: file_lotus_evm_erc721_proto_depIdxs,
		MessageInfos:      file_lotus_evm_erc721_proto_msgTypes,
	}.Build()
	File_lotus_evm_erc721_proto = out.File
	file_lotus_evm_erc721_proto_rawDesc = nil
	file_lotus_evm_erc721_proto_goTypes = nil
	file_lotus_evm_erc721_proto_depIdxs = nil
}
