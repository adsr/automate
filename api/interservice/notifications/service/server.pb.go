// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.0
// source: interservice/notifications/service/server.proto

package service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_interservice_notifications_service_server_proto protoreflect.FileDescriptor

var file_interservice_notifications_service_server_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x2a, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65,
	0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0xa8, 0x08, 0x0a, 0x0d, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x71, 0x0a, 0x06, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x31, 0x2e, 0x63, 0x68,
	0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x34,
	0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x78, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x52, 0x75, 0x6c, 0x65, 0x12,
	0x30, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x75, 0x6c,
	0x65, 0x1a, 0x3b, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52,
	0x75, 0x6c, 0x65, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x88,
	0x01, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x3a, 0x2e,
	0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x1a, 0x3e, 0x2e, 0x63, 0x68, 0x65, 0x66,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7e, 0x0a, 0x0a, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x30, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x1a, 0x3e, 0x2e, 0x63, 0x68, 0x65, 0x66,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x82, 0x01, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x3a, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x1a, 0x3b, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52,
	0x75, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7c,
	0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x31, 0x2e, 0x63, 0x68,
	0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x3c,
	0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x75, 0x6c, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x96, 0x01, 0x0a,
	0x0f, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b,
	0x12, 0x40, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65,
	0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x52,
	0x4c, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x41, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61,
	0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x55, 0x52, 0x4c, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x82, 0x01, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x3a, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e,
	0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75,
	0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_interservice_notifications_service_server_proto_goTypes = []interface{}{
	(*Event)(nil),                 // 0: chef.automate.domain.notifications.service.Event
	(*Rule)(nil),                  // 1: chef.automate.domain.notifications.service.Rule
	(*RuleIdentifier)(nil),        // 2: chef.automate.domain.notifications.service.RuleIdentifier
	(*Empty)(nil),                 // 3: chef.automate.domain.notifications.service.Empty
	(*URLValidationRequest)(nil),  // 4: chef.automate.domain.notifications.service.URLValidationRequest
	(*VersionRequest)(nil),        // 5: chef.automate.domain.notifications.service.VersionRequest
	(*Response)(nil),              // 6: chef.automate.domain.notifications.service.Response
	(*RuleAddResponse)(nil),       // 7: chef.automate.domain.notifications.service.RuleAddResponse
	(*RuleDeleteResponse)(nil),    // 8: chef.automate.domain.notifications.service.RuleDeleteResponse
	(*RuleUpdateResponse)(nil),    // 9: chef.automate.domain.notifications.service.RuleUpdateResponse
	(*RuleGetResponse)(nil),       // 10: chef.automate.domain.notifications.service.RuleGetResponse
	(*RuleListResponse)(nil),      // 11: chef.automate.domain.notifications.service.RuleListResponse
	(*URLValidationResponse)(nil), // 12: chef.automate.domain.notifications.service.URLValidationResponse
	(*VersionResponse)(nil),       // 13: chef.automate.domain.notifications.service.VersionResponse
}
var file_interservice_notifications_service_server_proto_depIdxs = []int32{
	0,  // 0: chef.automate.domain.notifications.service.Notifications.Notify:input_type -> chef.automate.domain.notifications.service.Event
	1,  // 1: chef.automate.domain.notifications.service.Notifications.AddRule:input_type -> chef.automate.domain.notifications.service.Rule
	2,  // 2: chef.automate.domain.notifications.service.Notifications.DeleteRule:input_type -> chef.automate.domain.notifications.service.RuleIdentifier
	1,  // 3: chef.automate.domain.notifications.service.Notifications.UpdateRule:input_type -> chef.automate.domain.notifications.service.Rule
	2,  // 4: chef.automate.domain.notifications.service.Notifications.GetRule:input_type -> chef.automate.domain.notifications.service.RuleIdentifier
	3,  // 5: chef.automate.domain.notifications.service.Notifications.ListRules:input_type -> chef.automate.domain.notifications.service.Empty
	4,  // 6: chef.automate.domain.notifications.service.Notifications.ValidateWebhook:input_type -> chef.automate.domain.notifications.service.URLValidationRequest
	5,  // 7: chef.automate.domain.notifications.service.Notifications.Version:input_type -> chef.automate.domain.notifications.service.VersionRequest
	6,  // 8: chef.automate.domain.notifications.service.Notifications.Notify:output_type -> chef.automate.domain.notifications.service.Response
	7,  // 9: chef.automate.domain.notifications.service.Notifications.AddRule:output_type -> chef.automate.domain.notifications.service.RuleAddResponse
	8,  // 10: chef.automate.domain.notifications.service.Notifications.DeleteRule:output_type -> chef.automate.domain.notifications.service.RuleDeleteResponse
	9,  // 11: chef.automate.domain.notifications.service.Notifications.UpdateRule:output_type -> chef.automate.domain.notifications.service.RuleUpdateResponse
	10, // 12: chef.automate.domain.notifications.service.Notifications.GetRule:output_type -> chef.automate.domain.notifications.service.RuleGetResponse
	11, // 13: chef.automate.domain.notifications.service.Notifications.ListRules:output_type -> chef.automate.domain.notifications.service.RuleListResponse
	12, // 14: chef.automate.domain.notifications.service.Notifications.ValidateWebhook:output_type -> chef.automate.domain.notifications.service.URLValidationResponse
	13, // 15: chef.automate.domain.notifications.service.Notifications.Version:output_type -> chef.automate.domain.notifications.service.VersionResponse
	8,  // [8:16] is the sub-list for method output_type
	0,  // [0:8] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_interservice_notifications_service_server_proto_init() }
func file_interservice_notifications_service_server_proto_init() {
	if File_interservice_notifications_service_server_proto != nil {
		return
	}
	file_interservice_notifications_service_events_proto_init()
	file_interservice_notifications_service_rules_proto_init()
	file_interservice_notifications_service_health_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_interservice_notifications_service_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_interservice_notifications_service_server_proto_goTypes,
		DependencyIndexes: file_interservice_notifications_service_server_proto_depIdxs,
	}.Build()
	File_interservice_notifications_service_server_proto = out.File
	file_interservice_notifications_service_server_proto_rawDesc = nil
	file_interservice_notifications_service_server_proto_goTypes = nil
	file_interservice_notifications_service_server_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NotificationsClient is the client API for Notifications service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NotificationsClient interface {
	// Publish a notification
	Notify(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Response, error)
	// Manage notification alerting rules
	AddRule(ctx context.Context, in *Rule, opts ...grpc.CallOption) (*RuleAddResponse, error)
	DeleteRule(ctx context.Context, in *RuleIdentifier, opts ...grpc.CallOption) (*RuleDeleteResponse, error)
	UpdateRule(ctx context.Context, in *Rule, opts ...grpc.CallOption) (*RuleUpdateResponse, error)
	GetRule(ctx context.Context, in *RuleIdentifier, opts ...grpc.CallOption) (*RuleGetResponse, error)
	ListRules(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RuleListResponse, error)
	ValidateWebhook(ctx context.Context, in *URLValidationRequest, opts ...grpc.CallOption) (*URLValidationResponse, error)
	// Health checks and metadata
	Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
}

type notificationsClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationsClient(cc grpc.ClientConnInterface) NotificationsClient {
	return &notificationsClient{cc}
}

func (c *notificationsClient) Notify(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.notifications.service.Notifications/Notify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsClient) AddRule(ctx context.Context, in *Rule, opts ...grpc.CallOption) (*RuleAddResponse, error) {
	out := new(RuleAddResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.notifications.service.Notifications/AddRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsClient) DeleteRule(ctx context.Context, in *RuleIdentifier, opts ...grpc.CallOption) (*RuleDeleteResponse, error) {
	out := new(RuleDeleteResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.notifications.service.Notifications/DeleteRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsClient) UpdateRule(ctx context.Context, in *Rule, opts ...grpc.CallOption) (*RuleUpdateResponse, error) {
	out := new(RuleUpdateResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.notifications.service.Notifications/UpdateRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsClient) GetRule(ctx context.Context, in *RuleIdentifier, opts ...grpc.CallOption) (*RuleGetResponse, error) {
	out := new(RuleGetResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.notifications.service.Notifications/GetRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsClient) ListRules(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RuleListResponse, error) {
	out := new(RuleListResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.notifications.service.Notifications/ListRules", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsClient) ValidateWebhook(ctx context.Context, in *URLValidationRequest, opts ...grpc.CallOption) (*URLValidationResponse, error) {
	out := new(URLValidationResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.notifications.service.Notifications/ValidateWebhook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsClient) Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.notifications.service.Notifications/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationsServer is the server API for Notifications service.
type NotificationsServer interface {
	// Publish a notification
	Notify(context.Context, *Event) (*Response, error)
	// Manage notification alerting rules
	AddRule(context.Context, *Rule) (*RuleAddResponse, error)
	DeleteRule(context.Context, *RuleIdentifier) (*RuleDeleteResponse, error)
	UpdateRule(context.Context, *Rule) (*RuleUpdateResponse, error)
	GetRule(context.Context, *RuleIdentifier) (*RuleGetResponse, error)
	ListRules(context.Context, *Empty) (*RuleListResponse, error)
	ValidateWebhook(context.Context, *URLValidationRequest) (*URLValidationResponse, error)
	// Health checks and metadata
	Version(context.Context, *VersionRequest) (*VersionResponse, error)
}

// UnimplementedNotificationsServer can be embedded to have forward compatible implementations.
type UnimplementedNotificationsServer struct {
}

func (*UnimplementedNotificationsServer) Notify(context.Context, *Event) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Notify not implemented")
}
func (*UnimplementedNotificationsServer) AddRule(context.Context, *Rule) (*RuleAddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRule not implemented")
}
func (*UnimplementedNotificationsServer) DeleteRule(context.Context, *RuleIdentifier) (*RuleDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRule not implemented")
}
func (*UnimplementedNotificationsServer) UpdateRule(context.Context, *Rule) (*RuleUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRule not implemented")
}
func (*UnimplementedNotificationsServer) GetRule(context.Context, *RuleIdentifier) (*RuleGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRule not implemented")
}
func (*UnimplementedNotificationsServer) ListRules(context.Context, *Empty) (*RuleListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRules not implemented")
}
func (*UnimplementedNotificationsServer) ValidateWebhook(context.Context, *URLValidationRequest) (*URLValidationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateWebhook not implemented")
}
func (*UnimplementedNotificationsServer) Version(context.Context, *VersionRequest) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}

func RegisterNotificationsServer(s *grpc.Server, srv NotificationsServer) {
	s.RegisterService(&_Notifications_serviceDesc, srv)
}

func _Notifications_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.notifications.service.Notifications/Notify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).Notify(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifications_AddRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Rule)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).AddRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.notifications.service.Notifications/AddRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).AddRule(ctx, req.(*Rule))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifications_DeleteRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RuleIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).DeleteRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.notifications.service.Notifications/DeleteRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).DeleteRule(ctx, req.(*RuleIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifications_UpdateRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Rule)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).UpdateRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.notifications.service.Notifications/UpdateRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).UpdateRule(ctx, req.(*Rule))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifications_GetRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RuleIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).GetRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.notifications.service.Notifications/GetRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).GetRule(ctx, req.(*RuleIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifications_ListRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).ListRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.notifications.service.Notifications/ListRules",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).ListRules(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifications_ValidateWebhook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(URLValidationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).ValidateWebhook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.notifications.service.Notifications/ValidateWebhook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).ValidateWebhook(ctx, req.(*URLValidationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifications_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.notifications.service.Notifications/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).Version(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Notifications_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.domain.notifications.service.Notifications",
	HandlerType: (*NotificationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Notify",
			Handler:    _Notifications_Notify_Handler,
		},
		{
			MethodName: "AddRule",
			Handler:    _Notifications_AddRule_Handler,
		},
		{
			MethodName: "DeleteRule",
			Handler:    _Notifications_DeleteRule_Handler,
		},
		{
			MethodName: "UpdateRule",
			Handler:    _Notifications_UpdateRule_Handler,
		},
		{
			MethodName: "GetRule",
			Handler:    _Notifications_GetRule_Handler,
		},
		{
			MethodName: "ListRules",
			Handler:    _Notifications_ListRules_Handler,
		},
		{
			MethodName: "ValidateWebhook",
			Handler:    _Notifications_ValidateWebhook_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _Notifications_Version_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interservice/notifications/service/server.proto",
}
