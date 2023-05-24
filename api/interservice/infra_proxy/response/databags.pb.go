// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.0
// source: interservice/infra_proxy/response/databags.proto

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

type DataBags struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Data bags list.
	DataBags []*DataBagListItem `protobuf:"bytes,1,rep,name=data_bags,json=dataBags,proto3" json:"data_bags,omitempty" toml:"data_bags,omitempty" mapstructure:"data_bags,omitempty"`
}

func (x *DataBags) Reset() {
	*x = DataBags{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_response_databags_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataBags) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataBags) ProtoMessage() {}

func (x *DataBags) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_response_databags_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataBags.ProtoReflect.Descriptor instead.
func (*DataBags) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_response_databags_proto_rawDescGZIP(), []int{0}
}

func (x *DataBags) GetDataBags() []*DataBagListItem {
	if x != nil {
		return x.DataBags
	}
	return nil
}

type DataBagItems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Data bag name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	// Data bags item list.
	Items []*DataBagListItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty" toml:"items,omitempty" mapstructure:"items,omitempty"`
	// Starting page for the results.
	Page int32 `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty" toml:"page,omitempty" mapstructure:"page,omitempty"`
	// Total number of records.
	Total int32 `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty" toml:"total,omitempty" mapstructure:"total,omitempty"`
}

func (x *DataBagItems) Reset() {
	*x = DataBagItems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_response_databags_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataBagItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataBagItems) ProtoMessage() {}

func (x *DataBagItems) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_response_databags_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataBagItems.ProtoReflect.Descriptor instead.
func (*DataBagItems) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_response_databags_proto_rawDescGZIP(), []int{1}
}

func (x *DataBagItems) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DataBagItems) GetItems() []*DataBagListItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *DataBagItems) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *DataBagItems) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type DataBagListItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Data bag item name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
}

func (x *DataBagListItem) Reset() {
	*x = DataBagListItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_response_databags_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataBagListItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataBagListItem) ProtoMessage() {}

func (x *DataBagListItem) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_response_databags_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataBagListItem.ProtoReflect.Descriptor instead.
func (*DataBagListItem) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_response_databags_proto_rawDescGZIP(), []int{2}
}

func (x *DataBagListItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DataBagItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Data bag name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	// Data bag item ID.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
	// Stringified json of data bag item.
	Data string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty" toml:"data,omitempty" mapstructure:"data,omitempty"`
}

func (x *DataBagItem) Reset() {
	*x = DataBagItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_response_databags_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataBagItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataBagItem) ProtoMessage() {}

func (x *DataBagItem) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_response_databags_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataBagItem.ProtoReflect.Descriptor instead.
func (*DataBagItem) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_response_databags_proto_rawDescGZIP(), []int{3}
}

func (x *DataBagItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DataBagItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DataBagItem) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type DataBag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Data bag name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
}

func (x *DataBag) Reset() {
	*x = DataBag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_response_databags_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataBag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataBag) ProtoMessage() {}

func (x *DataBag) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_response_databags_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataBag.ProtoReflect.Descriptor instead.
func (*DataBag) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_response_databags_proto_rawDescGZIP(), []int{4}
}

func (x *DataBag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateDataBag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Data bag name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
}

func (x *CreateDataBag) Reset() {
	*x = CreateDataBag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_response_databags_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDataBag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDataBag) ProtoMessage() {}

func (x *CreateDataBag) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_response_databags_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDataBag.ProtoReflect.Descriptor instead.
func (*CreateDataBag) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_response_databags_proto_rawDescGZIP(), []int{5}
}

func (x *CreateDataBag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateDataBagItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Data bag name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	// Data bag item ID.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
}

func (x *CreateDataBagItem) Reset() {
	*x = CreateDataBagItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_response_databags_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDataBagItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDataBagItem) ProtoMessage() {}

func (x *CreateDataBagItem) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_response_databags_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDataBagItem.ProtoReflect.Descriptor instead.
func (*CreateDataBagItem) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_response_databags_proto_rawDescGZIP(), []int{6}
}

func (x *CreateDataBagItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateDataBagItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateDataBagItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Data bag name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	// Data bag item ID.
	ItemId string `protobuf:"bytes,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty" toml:"item_id,omitempty" mapstructure:"item_id,omitempty"`
}

func (x *UpdateDataBagItem) Reset() {
	*x = UpdateDataBagItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_response_databags_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDataBagItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDataBagItem) ProtoMessage() {}

func (x *UpdateDataBagItem) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_response_databags_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDataBagItem.ProtoReflect.Descriptor instead.
func (*UpdateDataBagItem) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_response_databags_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateDataBagItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateDataBagItem) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

var File_interservice_infra_proxy_response_databags_proto protoreflect.FileDescriptor

var file_interservice_infra_proxy_response_databags_proto_rawDesc = []byte{
	0x0a, 0x30, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x29, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x63, 0x0a,
	0x08, 0x44, 0x61, 0x74, 0x61, 0x42, 0x61, 0x67, 0x73, 0x12, 0x57, 0x0a, 0x09, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x62, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x63,
	0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x42, 0x61, 0x67,
	0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x42, 0x61,
	0x67, 0x73, 0x22, 0x9e, 0x01, 0x0a, 0x0c, 0x44, 0x61, 0x74, 0x61, 0x42, 0x61, 0x67, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x50, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75,
	0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x42, 0x61, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x22, 0x25, 0x0a, 0x0f, 0x44, 0x61, 0x74, 0x61, 0x42, 0x61, 0x67, 0x4c, 0x69,
	0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x45, 0x0a, 0x0b, 0x44, 0x61,
	0x74, 0x61, 0x42, 0x61, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x1d, 0x0a, 0x07, 0x44, 0x61, 0x74, 0x61, 0x42, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x23, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x61,
	0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x37, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x42, 0x61, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x40,
	0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x61, 0x67, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64,
	0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_interservice_infra_proxy_response_databags_proto_rawDescOnce sync.Once
	file_interservice_infra_proxy_response_databags_proto_rawDescData = file_interservice_infra_proxy_response_databags_proto_rawDesc
)

func file_interservice_infra_proxy_response_databags_proto_rawDescGZIP() []byte {
	file_interservice_infra_proxy_response_databags_proto_rawDescOnce.Do(func() {
		file_interservice_infra_proxy_response_databags_proto_rawDescData = protoimpl.X.CompressGZIP(file_interservice_infra_proxy_response_databags_proto_rawDescData)
	})
	return file_interservice_infra_proxy_response_databags_proto_rawDescData
}

var file_interservice_infra_proxy_response_databags_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_interservice_infra_proxy_response_databags_proto_goTypes = []interface{}{
	(*DataBags)(nil),          // 0: chef.automate.domain.infra_proxy.response.DataBags
	(*DataBagItems)(nil),      // 1: chef.automate.domain.infra_proxy.response.DataBagItems
	(*DataBagListItem)(nil),   // 2: chef.automate.domain.infra_proxy.response.DataBagListItem
	(*DataBagItem)(nil),       // 3: chef.automate.domain.infra_proxy.response.DataBagItem
	(*DataBag)(nil),           // 4: chef.automate.domain.infra_proxy.response.DataBag
	(*CreateDataBag)(nil),     // 5: chef.automate.domain.infra_proxy.response.CreateDataBag
	(*CreateDataBagItem)(nil), // 6: chef.automate.domain.infra_proxy.response.CreateDataBagItem
	(*UpdateDataBagItem)(nil), // 7: chef.automate.domain.infra_proxy.response.UpdateDataBagItem
}
var file_interservice_infra_proxy_response_databags_proto_depIdxs = []int32{
	2, // 0: chef.automate.domain.infra_proxy.response.DataBags.data_bags:type_name -> chef.automate.domain.infra_proxy.response.DataBagListItem
	2, // 1: chef.automate.domain.infra_proxy.response.DataBagItems.items:type_name -> chef.automate.domain.infra_proxy.response.DataBagListItem
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_interservice_infra_proxy_response_databags_proto_init() }
func file_interservice_infra_proxy_response_databags_proto_init() {
	if File_interservice_infra_proxy_response_databags_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_interservice_infra_proxy_response_databags_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataBags); i {
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
		file_interservice_infra_proxy_response_databags_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataBagItems); i {
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
		file_interservice_infra_proxy_response_databags_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataBagListItem); i {
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
		file_interservice_infra_proxy_response_databags_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataBagItem); i {
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
		file_interservice_infra_proxy_response_databags_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataBag); i {
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
		file_interservice_infra_proxy_response_databags_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDataBag); i {
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
		file_interservice_infra_proxy_response_databags_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDataBagItem); i {
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
		file_interservice_infra_proxy_response_databags_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDataBagItem); i {
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
			RawDescriptor: file_interservice_infra_proxy_response_databags_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_interservice_infra_proxy_response_databags_proto_goTypes,
		DependencyIndexes: file_interservice_infra_proxy_response_databags_proto_depIdxs,
		MessageInfos:      file_interservice_infra_proxy_response_databags_proto_msgTypes,
	}.Build()
	File_interservice_infra_proxy_response_databags_proto = out.File
	file_interservice_infra_proxy_response_databags_proto_rawDesc = nil
	file_interservice_infra_proxy_response_databags_proto_goTypes = nil
	file_interservice_infra_proxy_response_databags_proto_depIdxs = nil
}
