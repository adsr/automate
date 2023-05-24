// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.0
// source: external/cfgmgmt/response/stats.proto

package response

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

type RunsCounts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Total count of run reports that have landed in Automate for the node.
	Total int32 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	// Total count of successful run reports that have landed in Automate for the node.
	Success int32 `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	// Total count of failed run reports that have landed in Automate for the node.
	Failure int32 `protobuf:"varint,3,opt,name=failure,proto3" json:"failure,omitempty"`
}

func (x *RunsCounts) Reset() {
	*x = RunsCounts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_cfgmgmt_response_stats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunsCounts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunsCounts) ProtoMessage() {}

func (x *RunsCounts) ProtoReflect() protoreflect.Message {
	mi := &file_external_cfgmgmt_response_stats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunsCounts.ProtoReflect.Descriptor instead.
func (*RunsCounts) Descriptor() ([]byte, []int) {
	return file_external_cfgmgmt_response_stats_proto_rawDescGZIP(), []int{0}
}

func (x *RunsCounts) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *RunsCounts) GetSuccess() int32 {
	if x != nil {
		return x.Success
	}
	return 0
}

func (x *RunsCounts) GetFailure() int32 {
	if x != nil {
		return x.Failure
	}
	return 0
}

type NodesCounts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Total count of nodes that have reported in to Automate.
	Total int32 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	// Total count of nodes that have reported in to Automate whose last run was successful.
	Success int32 `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	// Total count of nodes that have reported in to Automate whose last run was failed.
	Failure int32 `protobuf:"varint,3,opt,name=failure,proto3" json:"failure,omitempty"`
	// Total count of nodes that have been labeled as 'missing' as determined by node lifecycle settings.
	Missing int32 `protobuf:"varint,4,opt,name=missing,proto3" json:"missing,omitempty"`
}

func (x *NodesCounts) Reset() {
	*x = NodesCounts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_cfgmgmt_response_stats_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodesCounts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodesCounts) ProtoMessage() {}

func (x *NodesCounts) ProtoReflect() protoreflect.Message {
	mi := &file_external_cfgmgmt_response_stats_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodesCounts.ProtoReflect.Descriptor instead.
func (*NodesCounts) Descriptor() ([]byte, []int) {
	return file_external_cfgmgmt_response_stats_proto_rawDescGZIP(), []int{1}
}

func (x *NodesCounts) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *NodesCounts) GetSuccess() int32 {
	if x != nil {
		return x.Success
	}
	return 0
}

func (x *NodesCounts) GetFailure() int32 {
	if x != nil {
		return x.Failure
	}
	return 0
}

func (x *NodesCounts) GetMissing() int32 {
	if x != nil {
		return x.Missing
	}
	return 0
}

type CheckInCountsTimeSeries struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of daily checkin counts
	Counts []*CheckInCounts `protobuf:"bytes,1,rep,name=counts,proto3" json:"counts,omitempty"`
}

func (x *CheckInCountsTimeSeries) Reset() {
	*x = CheckInCountsTimeSeries{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_cfgmgmt_response_stats_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckInCountsTimeSeries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckInCountsTimeSeries) ProtoMessage() {}

func (x *CheckInCountsTimeSeries) ProtoReflect() protoreflect.Message {
	mi := &file_external_cfgmgmt_response_stats_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckInCountsTimeSeries.ProtoReflect.Descriptor instead.
func (*CheckInCountsTimeSeries) Descriptor() ([]byte, []int) {
	return file_external_cfgmgmt_response_stats_proto_rawDescGZIP(), []int{2}
}

func (x *CheckInCountsTimeSeries) GetCounts() []*CheckInCounts {
	if x != nil {
		return x.Counts
	}
	return nil
}

type CheckInCounts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start string `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End   string `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
	Count int32  `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	Total int32  `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *CheckInCounts) Reset() {
	*x = CheckInCounts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_cfgmgmt_response_stats_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckInCounts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckInCounts) ProtoMessage() {}

func (x *CheckInCounts) ProtoReflect() protoreflect.Message {
	mi := &file_external_cfgmgmt_response_stats_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckInCounts.ProtoReflect.Descriptor instead.
func (*CheckInCounts) Descriptor() ([]byte, []int) {
	return file_external_cfgmgmt_response_stats_proto_rawDescGZIP(), []int{3}
}

func (x *CheckInCounts) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *CheckInCounts) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

func (x *CheckInCounts) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *CheckInCounts) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_external_cfgmgmt_response_stats_proto protoreflect.FileDescriptor

var file_external_cfgmgmt_response_stats_proto_rawDesc = []byte{
	0x0a, 0x25, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x66, 0x67, 0x6d, 0x67,
	0x6d, 0x74, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75,
	0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x66, 0x67, 0x6d, 0x67,
	0x6d, 0x74, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x56, 0x0a, 0x0a, 0x52,
	0x75, 0x6e, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x61, 0x69,
	0x6c, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x66, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x22, 0x71, 0x0a, 0x0b, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x22, 0x64, 0x0a, 0x17, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49,
	0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65,
	0x73, 0x12, 0x49, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x31, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x66, 0x67, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x52, 0x06, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22, 0x63, 0x0a, 0x0d,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x66, 0x67, 0x6d, 0x67,
	0x6d, 0x74, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_external_cfgmgmt_response_stats_proto_rawDescOnce sync.Once
	file_external_cfgmgmt_response_stats_proto_rawDescData = file_external_cfgmgmt_response_stats_proto_rawDesc
)

func file_external_cfgmgmt_response_stats_proto_rawDescGZIP() []byte {
	file_external_cfgmgmt_response_stats_proto_rawDescOnce.Do(func() {
		file_external_cfgmgmt_response_stats_proto_rawDescData = protoimpl.X.CompressGZIP(file_external_cfgmgmt_response_stats_proto_rawDescData)
	})
	return file_external_cfgmgmt_response_stats_proto_rawDescData
}

var file_external_cfgmgmt_response_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_external_cfgmgmt_response_stats_proto_goTypes = []interface{}{
	(*RunsCounts)(nil),              // 0: chef.automate.api.cfgmgmt.response.RunsCounts
	(*NodesCounts)(nil),             // 1: chef.automate.api.cfgmgmt.response.NodesCounts
	(*CheckInCountsTimeSeries)(nil), // 2: chef.automate.api.cfgmgmt.response.CheckInCountsTimeSeries
	(*CheckInCounts)(nil),           // 3: chef.automate.api.cfgmgmt.response.CheckInCounts
}
var file_external_cfgmgmt_response_stats_proto_depIdxs = []int32{
	3, // 0: chef.automate.api.cfgmgmt.response.CheckInCountsTimeSeries.counts:type_name -> chef.automate.api.cfgmgmt.response.CheckInCounts
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_external_cfgmgmt_response_stats_proto_init() }
func file_external_cfgmgmt_response_stats_proto_init() {
	if File_external_cfgmgmt_response_stats_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_external_cfgmgmt_response_stats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunsCounts); i {
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
		file_external_cfgmgmt_response_stats_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodesCounts); i {
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
		file_external_cfgmgmt_response_stats_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckInCountsTimeSeries); i {
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
		file_external_cfgmgmt_response_stats_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckInCounts); i {
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
			RawDescriptor: file_external_cfgmgmt_response_stats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_external_cfgmgmt_response_stats_proto_goTypes,
		DependencyIndexes: file_external_cfgmgmt_response_stats_proto_depIdxs,
		MessageInfos:      file_external_cfgmgmt_response_stats_proto_msgTypes,
	}.Build()
	File_external_cfgmgmt_response_stats_proto = out.File
	file_external_cfgmgmt_response_stats_proto_rawDesc = nil
	file_external_cfgmgmt_response_stats_proto_goTypes = nil
	file_external_cfgmgmt_response_stats_proto_depIdxs = nil
}
