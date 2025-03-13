// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: api/ais_api.proto

package ais_api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Status int32

const (
	Status_active   Status = 0
	Status_inactive Status = 1
	Status_closed   Status = 2
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "active",
		1: "inactive",
		2: "closed",
	}
	Status_value = map[string]int32{
		"active":   0,
		"inactive": 1,
		"closed":   2,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_api_ais_api_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_api_ais_api_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_api_ais_api_proto_rawDescGZIP(), []int{0}
}

type AccountType int32

const (
	AccountType_CASA AccountType = 0
	AccountType_GA   AccountType = 1
	AccountType_VAN  AccountType = 2
)

// Enum value maps for AccountType.
var (
	AccountType_name = map[int32]string{
		0: "CASA",
		1: "GA",
		2: "VAN",
	}
	AccountType_value = map[string]int32{
		"CASA": 0,
		"GA":   1,
		"VAN":  2,
	}
)

func (x AccountType) Enum() *AccountType {
	p := new(AccountType)
	*p = x
	return p
}

func (x AccountType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AccountType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_ais_api_proto_enumTypes[1].Descriptor()
}

func (AccountType) Type() protoreflect.EnumType {
	return &file_api_ais_api_proto_enumTypes[1]
}

func (x AccountType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AccountType.Descriptor instead.
func (AccountType) EnumDescriptor() ([]byte, []int) {
	return file_api_ais_api_proto_rawDescGZIP(), []int{1}
}

//	enum Source {
//	    MBBanking = 0;
//	    GA = 1;
//	    VAN = 2;
//	}
type GetAccountStatusRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccountId     uint64                 `protobuf:"varint,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAccountStatusRequest) Reset() {
	*x = GetAccountStatusRequest{}
	mi := &file_api_ais_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAccountStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountStatusRequest) ProtoMessage() {}

func (x *GetAccountStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_ais_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountStatusRequest.ProtoReflect.Descriptor instead.
func (*GetAccountStatusRequest) Descriptor() ([]byte, []int) {
	return file_api_ais_api_proto_rawDescGZIP(), []int{0}
}

func (x *GetAccountStatusRequest) GetAccountId() uint64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

type GetAccountStatusResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccountId     uint64                 `protobuf:"varint,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	AccountName   string                 `protobuf:"bytes,2,opt,name=account_name,json=accountName,proto3" json:"account_name,omitempty"`
	AccountType   AccountType            `protobuf:"varint,3,opt,name=account_type,json=accountType,proto3,enum=ais_api.AccountType" json:"account_type,omitempty"`
	AccountStatus Status                 `protobuf:"varint,4,opt,name=account_status,json=accountStatus,proto3,enum=ais_api.Status" json:"account_status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAccountStatusResponse) Reset() {
	*x = GetAccountStatusResponse{}
	mi := &file_api_ais_api_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAccountStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountStatusResponse) ProtoMessage() {}

func (x *GetAccountStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_ais_api_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountStatusResponse.ProtoReflect.Descriptor instead.
func (*GetAccountStatusResponse) Descriptor() ([]byte, []int) {
	return file_api_ais_api_proto_rawDescGZIP(), []int{1}
}

func (x *GetAccountStatusResponse) GetAccountId() uint64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *GetAccountStatusResponse) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *GetAccountStatusResponse) GetAccountType() AccountType {
	if x != nil {
		return x.AccountType
	}
	return AccountType_CASA
}

func (x *GetAccountStatusResponse) GetAccountStatus() Status {
	if x != nil {
		return x.AccountStatus
	}
	return Status_active
}

var File_api_ais_api_proto protoreflect.FileDescriptor

var file_api_ais_api_proto_rawDesc = string([]byte{
	0x0a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x69, 0x73, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x69, 0x73, 0x5f, 0x61, 0x70, 0x69, 0x22, 0x38, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0xcd, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x61, 0x69,
	0x73, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x36,
	0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x61, 0x69, 0x73, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x2e, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0a, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x69, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x63, 0x6c,
	0x6f, 0x73, 0x65, 0x64, 0x10, 0x02, 0x2a, 0x28, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x41, 0x53, 0x41, 0x10, 0x00, 0x12,
	0x06, 0x0a, 0x02, 0x47, 0x41, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x56, 0x41, 0x4e, 0x10, 0x02,
	0x32, 0x68, 0x0a, 0x0a, 0x41, 0x49, 0x53, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x69, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42,
	0x79, 0x49, 0x44, 0x12, 0x20, 0x2e, 0x61, 0x69, 0x73, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x69, 0x73, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x61, 0x69, 0x73, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_api_ais_api_proto_rawDescOnce sync.Once
	file_api_ais_api_proto_rawDescData []byte
)

func file_api_ais_api_proto_rawDescGZIP() []byte {
	file_api_ais_api_proto_rawDescOnce.Do(func() {
		file_api_ais_api_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_ais_api_proto_rawDesc), len(file_api_ais_api_proto_rawDesc)))
	})
	return file_api_ais_api_proto_rawDescData
}

var file_api_ais_api_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_ais_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_ais_api_proto_goTypes = []any{
	(Status)(0),                      // 0: ais_api.Status
	(AccountType)(0),                 // 1: ais_api.AccountType
	(*GetAccountStatusRequest)(nil),  // 2: ais_api.GetAccountStatusRequest
	(*GetAccountStatusResponse)(nil), // 3: ais_api.GetAccountStatusResponse
}
var file_api_ais_api_proto_depIdxs = []int32{
	1, // 0: ais_api.GetAccountStatusResponse.account_type:type_name -> ais_api.AccountType
	0, // 1: ais_api.GetAccountStatusResponse.account_status:type_name -> ais_api.Status
	2, // 2: ais_api.AISService.GetAisAccountByID:input_type -> ais_api.GetAccountStatusRequest
	3, // 3: ais_api.AISService.GetAisAccountByID:output_type -> ais_api.GetAccountStatusResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_ais_api_proto_init() }
func file_api_ais_api_proto_init() {
	if File_api_ais_api_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_ais_api_proto_rawDesc), len(file_api_ais_api_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_ais_api_proto_goTypes,
		DependencyIndexes: file_api_ais_api_proto_depIdxs,
		EnumInfos:         file_api_ais_api_proto_enumTypes,
		MessageInfos:      file_api_ais_api_proto_msgTypes,
	}.Build()
	File_api_ais_api_proto = out.File
	file_api_ais_api_proto_goTypes = nil
	file_api_ais_api_proto_depIdxs = nil
}
