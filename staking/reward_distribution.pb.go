// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: lotus/staking/reward_distribution.proto

package staking

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

type RewardDistribution struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp          *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	DistributionId     string                 `protobuf:"bytes,2,opt,name=distribution_id,json=distributionId,proto3" json:"distribution_id,omitempty"`
	RecipientAddresses [][]byte               `protobuf:"bytes,3,rep,name=recipient_addresses,json=recipientAddresses,proto3" json:"recipient_addresses,omitempty"`
	RewardAmounts      []string               `protobuf:"bytes,4,rep,name=reward_amounts,json=rewardAmounts,proto3" json:"reward_amounts,omitempty"`
}

func (x *RewardDistribution) Reset() {
	*x = RewardDistribution{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lotus_staking_reward_distribution_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RewardDistribution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RewardDistribution) ProtoMessage() {}

func (x *RewardDistribution) ProtoReflect() protoreflect.Message {
	mi := &file_lotus_staking_reward_distribution_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RewardDistribution.ProtoReflect.Descriptor instead.
func (*RewardDistribution) Descriptor() ([]byte, []int) {
	return file_lotus_staking_reward_distribution_proto_rawDescGZIP(), []int{0}
}

func (x *RewardDistribution) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *RewardDistribution) GetDistributionId() string {
	if x != nil {
		return x.DistributionId
	}
	return ""
}

func (x *RewardDistribution) GetRecipientAddresses() [][]byte {
	if x != nil {
		return x.RecipientAddresses
	}
	return nil
}

func (x *RewardDistribution) GetRewardAmounts() []string {
	if x != nil {
		return x.RewardAmounts
	}
	return nil
}

var File_lotus_staking_reward_distribution_proto protoreflect.FileDescriptor

var file_lotus_staking_reward_distribution_proto_rawDesc = []byte{
	0x0a, 0x27, 0x6c, 0x6f, 0x74, 0x75, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2f,
	0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6c, 0x6f, 0x74, 0x75, 0x73,
	0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6c, 0x6f, 0x74, 0x75, 0x73,
	0x2f, 0x6c, 0x6f, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x01, 0x0a,
	0x12, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x27, 0x0a,
	0x0f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x13, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0c, 0x42, 0x06, 0xd2, 0xab, 0x30, 0x02, 0x08, 0x01, 0x52, 0x12, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12,
	0x2b, 0x0a, 0x0e, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x42, 0x04, 0xda, 0xab, 0x30, 0x00, 0x52, 0x0d, 0x72,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x42, 0x2c, 0x5a, 0x2a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72, 0x79, 0x6c, 0x6f,
	0x74, 0x75, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x6c, 0x6f, 0x74, 0x75, 0x73, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_lotus_staking_reward_distribution_proto_rawDescOnce sync.Once
	file_lotus_staking_reward_distribution_proto_rawDescData = file_lotus_staking_reward_distribution_proto_rawDesc
)

func file_lotus_staking_reward_distribution_proto_rawDescGZIP() []byte {
	file_lotus_staking_reward_distribution_proto_rawDescOnce.Do(func() {
		file_lotus_staking_reward_distribution_proto_rawDescData = protoimpl.X.CompressGZIP(file_lotus_staking_reward_distribution_proto_rawDescData)
	})
	return file_lotus_staking_reward_distribution_proto_rawDescData
}

var file_lotus_staking_reward_distribution_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_lotus_staking_reward_distribution_proto_goTypes = []interface{}{
	(*RewardDistribution)(nil),    // 0: lotus.staking.RewardDistribution
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_lotus_staking_reward_distribution_proto_depIdxs = []int32{
	1, // 0: lotus.staking.RewardDistribution.timestamp:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_lotus_staking_reward_distribution_proto_init() }
func file_lotus_staking_reward_distribution_proto_init() {
	if File_lotus_staking_reward_distribution_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lotus_staking_reward_distribution_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RewardDistribution); i {
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
			RawDescriptor: file_lotus_staking_reward_distribution_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lotus_staking_reward_distribution_proto_goTypes,
		DependencyIndexes: file_lotus_staking_reward_distribution_proto_depIdxs,
		MessageInfos:      file_lotus_staking_reward_distribution_proto_msgTypes,
	}.Build()
	File_lotus_staking_reward_distribution_proto = out.File
	file_lotus_staking_reward_distribution_proto_rawDesc = nil
	file_lotus_staking_reward_distribution_proto_goTypes = nil
	file_lotus_staking_reward_distribution_proto_depIdxs = nil
}