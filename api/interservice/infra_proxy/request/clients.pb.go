// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.0
// source: interservice/infra_proxy/request/clients.proto

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

type Clients struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty" toml:"org_id,omitempty" mapstructure:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
	// Clients search query.
	SearchQuery *SearchQuery `protobuf:"bytes,3,opt,name=search_query,json=searchQuery,proto3" json:"search_query,omitempty" toml:"search_query,omitempty" mapstructure:"search_query,omitempty"`
}

func (x *Clients) Reset() {
	*x = Clients{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_request_clients_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Clients) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Clients) ProtoMessage() {}

func (x *Clients) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_request_clients_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Clients.ProtoReflect.Descriptor instead.
func (*Clients) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_request_clients_proto_rawDescGZIP(), []int{0}
}

func (x *Clients) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *Clients) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *Clients) GetSearchQuery() *SearchQuery {
	if x != nil {
		return x.SearchQuery
	}
	return nil
}

type Client struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty" toml:"org_id,omitempty" mapstructure:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
	// Client name.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
}

func (x *Client) Reset() {
	*x = Client{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_request_clients_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Client) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Client) ProtoMessage() {}

func (x *Client) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_request_clients_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Client.ProtoReflect.Descriptor instead.
func (*Client) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_request_clients_proto_rawDescGZIP(), []int{1}
}

func (x *Client) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *Client) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *Client) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateClient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty" toml:"org_id,omitempty" mapstructure:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
	// Client name.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	// Boolean indicates client type is validator or not.
	Validator bool `protobuf:"varint,4,opt,name=validator,proto3" json:"validator,omitempty" toml:"validator,omitempty" mapstructure:"validator,omitempty"`
	// Boolean indicates whether it's required to create a key or not.
	CreateKey bool `protobuf:"varint,5,opt,name=create_key,json=createKey,proto3" json:"create_key,omitempty" toml:"create_key,omitempty" mapstructure:"create_key,omitempty"`
}

func (x *CreateClient) Reset() {
	*x = CreateClient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_request_clients_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClient) ProtoMessage() {}

func (x *CreateClient) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_request_clients_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClient.ProtoReflect.Descriptor instead.
func (*CreateClient) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_request_clients_proto_rawDescGZIP(), []int{2}
}

func (x *CreateClient) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *CreateClient) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *CreateClient) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateClient) GetValidator() bool {
	if x != nil {
		return x.Validator
	}
	return false
}

func (x *CreateClient) GetCreateKey() bool {
	if x != nil {
		return x.CreateKey
	}
	return false
}

type ClientKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty" toml:"org_id,omitempty" mapstructure:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
	// Client name.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	// Client key name.
	Key string `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty" toml:"key,omitempty" mapstructure:"key,omitempty"`
}

func (x *ClientKey) Reset() {
	*x = ClientKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_request_clients_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientKey) ProtoMessage() {}

func (x *ClientKey) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_request_clients_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientKey.ProtoReflect.Descriptor instead.
func (*ClientKey) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_request_clients_proto_rawDescGZIP(), []int{3}
}

func (x *ClientKey) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *ClientKey) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *ClientKey) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ClientKey) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

var File_interservice_infra_proxy_request_clients_proto protoreflect.FileDescriptor

var file_interservice_infra_proxy_request_clients_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x28, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x01, 0x0a, 0x07, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x58, 0x0a, 0x0c, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x35, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x0b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x22, 0x50, 0x0a, 0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x0a,
	0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f,
	0x72, 0x67, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x93, 0x01, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x22, 0x65, 0x0a, 0x09, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_interservice_infra_proxy_request_clients_proto_rawDescOnce sync.Once
	file_interservice_infra_proxy_request_clients_proto_rawDescData = file_interservice_infra_proxy_request_clients_proto_rawDesc
)

func file_interservice_infra_proxy_request_clients_proto_rawDescGZIP() []byte {
	file_interservice_infra_proxy_request_clients_proto_rawDescOnce.Do(func() {
		file_interservice_infra_proxy_request_clients_proto_rawDescData = protoimpl.X.CompressGZIP(file_interservice_infra_proxy_request_clients_proto_rawDescData)
	})
	return file_interservice_infra_proxy_request_clients_proto_rawDescData
}

var file_interservice_infra_proxy_request_clients_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_interservice_infra_proxy_request_clients_proto_goTypes = []interface{}{
	(*Clients)(nil),      // 0: chef.automate.domain.infra_proxy.request.Clients
	(*Client)(nil),       // 1: chef.automate.domain.infra_proxy.request.Client
	(*CreateClient)(nil), // 2: chef.automate.domain.infra_proxy.request.CreateClient
	(*ClientKey)(nil),    // 3: chef.automate.domain.infra_proxy.request.ClientKey
	(*SearchQuery)(nil),  // 4: chef.automate.domain.infra_proxy.request.SearchQuery
}
var file_interservice_infra_proxy_request_clients_proto_depIdxs = []int32{
	4, // 0: chef.automate.domain.infra_proxy.request.Clients.search_query:type_name -> chef.automate.domain.infra_proxy.request.SearchQuery
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_interservice_infra_proxy_request_clients_proto_init() }
func file_interservice_infra_proxy_request_clients_proto_init() {
	if File_interservice_infra_proxy_request_clients_proto != nil {
		return
	}
	file_interservice_infra_proxy_request_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_interservice_infra_proxy_request_clients_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Clients); i {
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
		file_interservice_infra_proxy_request_clients_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Client); i {
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
		file_interservice_infra_proxy_request_clients_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateClient); i {
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
		file_interservice_infra_proxy_request_clients_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientKey); i {
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
			RawDescriptor: file_interservice_infra_proxy_request_clients_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_interservice_infra_proxy_request_clients_proto_goTypes,
		DependencyIndexes: file_interservice_infra_proxy_request_clients_proto_depIdxs,
		MessageInfos:      file_interservice_infra_proxy_request_clients_proto_msgTypes,
	}.Build()
	File_interservice_infra_proxy_request_clients_proto = out.File
	file_interservice_infra_proxy_request_clients_proto_rawDesc = nil
	file_interservice_infra_proxy_request_clients_proto_goTypes = nil
	file_interservice_infra_proxy_request_clients_proto_depIdxs = nil
}
