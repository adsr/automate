// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.0
// source: external/infra_proxy/request/cookbooks.proto

package request

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

type Cookbooks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the Org.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// ID of the Server
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
}

func (x *Cookbooks) Reset() {
	*x = Cookbooks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_request_cookbooks_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cookbooks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cookbooks) ProtoMessage() {}

func (x *Cookbooks) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_request_cookbooks_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cookbooks.ProtoReflect.Descriptor instead.
func (*Cookbooks) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_request_cookbooks_proto_rawDescGZIP(), []int{0}
}

func (x *Cookbooks) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *Cookbooks) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

type CookbookVersions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the Org.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// ID of the Server.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	// Name of the cookbook.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CookbookVersions) Reset() {
	*x = CookbookVersions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_request_cookbooks_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CookbookVersions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CookbookVersions) ProtoMessage() {}

func (x *CookbookVersions) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_request_cookbooks_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CookbookVersions.ProtoReflect.Descriptor instead.
func (*CookbookVersions) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_request_cookbooks_proto_rawDescGZIP(), []int{1}
}

func (x *CookbookVersions) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *CookbookVersions) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *CookbookVersions) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Cookbook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the Org.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// ID of the Server.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	// Name of the cookbook.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Version of the cookbook.
	Version string `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Cookbook) Reset() {
	*x = Cookbook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_request_cookbooks_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cookbook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cookbook) ProtoMessage() {}

func (x *Cookbook) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_request_cookbooks_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cookbook.ProtoReflect.Descriptor instead.
func (*Cookbook) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_request_cookbooks_proto_rawDescGZIP(), []int{2}
}

func (x *Cookbook) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *Cookbook) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *Cookbook) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cookbook) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type CookbookFileContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the org.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// ID of the server.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	// Name of the cookbook.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Version of the cookbook.
	Version string `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	// Cookbook data file URL.
	Url string `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *CookbookFileContent) Reset() {
	*x = CookbookFileContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_request_cookbooks_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CookbookFileContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CookbookFileContent) ProtoMessage() {}

func (x *CookbookFileContent) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_request_cookbooks_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CookbookFileContent.ProtoReflect.Descriptor instead.
func (*CookbookFileContent) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_request_cookbooks_proto_rawDescGZIP(), []int{3}
}

func (x *CookbookFileContent) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *CookbookFileContent) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *CookbookFileContent) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CookbookFileContent) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *CookbookFileContent) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_external_infra_proxy_request_cookbooks_proto protoreflect.FileDescriptor

var file_external_infra_proxy_request_cookbooks_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x63,
	0x6f, 0x6f, 0x6b, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25,
	0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3f, 0x0a, 0x09, 0x43, 0x6f, 0x6f, 0x6b, 0x62, 0x6f, 0x6f,
	0x6b, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x22, 0x5a, 0x0a, 0x10, 0x43, 0x6f, 0x6f, 0x6b, 0x62, 0x6f,
	0x6f, 0x6b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72,
	0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x6c, 0x0a, 0x08, 0x43, 0x6f, 0x6f, 0x6b, 0x62, 0x6f, 0x6f, 0x6b, 0x12, 0x15,
	0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x22, 0x89, 0x01, 0x0a, 0x13, 0x43, 0x6f, 0x6f, 0x6b, 0x62, 0x6f, 0x6f, 0x6b, 0x46, 0x69, 0x6c,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x42, 0x3b, 0x5a, 0x39,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f,
	0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_external_infra_proxy_request_cookbooks_proto_rawDescOnce sync.Once
	file_external_infra_proxy_request_cookbooks_proto_rawDescData = file_external_infra_proxy_request_cookbooks_proto_rawDesc
)

func file_external_infra_proxy_request_cookbooks_proto_rawDescGZIP() []byte {
	file_external_infra_proxy_request_cookbooks_proto_rawDescOnce.Do(func() {
		file_external_infra_proxy_request_cookbooks_proto_rawDescData = protoimpl.X.CompressGZIP(file_external_infra_proxy_request_cookbooks_proto_rawDescData)
	})
	return file_external_infra_proxy_request_cookbooks_proto_rawDescData
}

var file_external_infra_proxy_request_cookbooks_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_external_infra_proxy_request_cookbooks_proto_goTypes = []interface{}{
	(*Cookbooks)(nil),           // 0: chef.automate.api.infra_proxy.request.Cookbooks
	(*CookbookVersions)(nil),    // 1: chef.automate.api.infra_proxy.request.CookbookVersions
	(*Cookbook)(nil),            // 2: chef.automate.api.infra_proxy.request.Cookbook
	(*CookbookFileContent)(nil), // 3: chef.automate.api.infra_proxy.request.CookbookFileContent
}
var file_external_infra_proxy_request_cookbooks_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_external_infra_proxy_request_cookbooks_proto_init() }
func file_external_infra_proxy_request_cookbooks_proto_init() {
	if File_external_infra_proxy_request_cookbooks_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_external_infra_proxy_request_cookbooks_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cookbooks); i {
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
		file_external_infra_proxy_request_cookbooks_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CookbookVersions); i {
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
		file_external_infra_proxy_request_cookbooks_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cookbook); i {
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
		file_external_infra_proxy_request_cookbooks_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CookbookFileContent); i {
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
			RawDescriptor: file_external_infra_proxy_request_cookbooks_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_external_infra_proxy_request_cookbooks_proto_goTypes,
		DependencyIndexes: file_external_infra_proxy_request_cookbooks_proto_depIdxs,
		MessageInfos:      file_external_infra_proxy_request_cookbooks_proto_msgTypes,
	}.Build()
	File_external_infra_proxy_request_cookbooks_proto = out.File
	file_external_infra_proxy_request_cookbooks_proto_rawDesc = nil
	file_external_infra_proxy_request_cookbooks_proto_goTypes = nil
	file_external_infra_proxy_request_cookbooks_proto_depIdxs = nil
}
