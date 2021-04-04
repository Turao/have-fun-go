// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: card/grpc/card.proto

package grpc

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

type Card struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OwnerId string `protobuf:"bytes,2,opt,name=ownerId,proto3" json:"ownerId,omitempty"`
}

func (x *Card) Reset() {
	*x = Card{}
	if protoimpl.UnsafeEnabled {
		mi := &file_card_grpc_card_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Card) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Card) ProtoMessage() {}

func (x *Card) ProtoReflect() protoreflect.Message {
	mi := &file_card_grpc_card_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Card.ProtoReflect.Descriptor instead.
func (*Card) Descriptor() ([]byte, []int) {
	return file_card_grpc_card_proto_rawDescGZIP(), []int{0}
}

func (x *Card) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Card) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

type CreateCardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateCardRequest) Reset() {
	*x = CreateCardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_card_grpc_card_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCardRequest) ProtoMessage() {}

func (x *CreateCardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_card_grpc_card_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCardRequest.ProtoReflect.Descriptor instead.
func (*CreateCardRequest) Descriptor() ([]byte, []int) {
	return file_card_grpc_card_proto_rawDescGZIP(), []int{1}
}

type AssignOwnerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CardId  string `protobuf:"bytes,1,opt,name=cardId,proto3" json:"cardId,omitempty"`
	OwnerId string `protobuf:"bytes,2,opt,name=ownerId,proto3" json:"ownerId,omitempty"`
}

func (x *AssignOwnerRequest) Reset() {
	*x = AssignOwnerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_card_grpc_card_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssignOwnerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignOwnerRequest) ProtoMessage() {}

func (x *AssignOwnerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_card_grpc_card_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignOwnerRequest.ProtoReflect.Descriptor instead.
func (*AssignOwnerRequest) Descriptor() ([]byte, []int) {
	return file_card_grpc_card_proto_rawDescGZIP(), []int{2}
}

func (x *AssignOwnerRequest) GetCardId() string {
	if x != nil {
		return x.CardId
	}
	return ""
}

func (x *AssignOwnerRequest) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

type UnassignOwnerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CardId string `protobuf:"bytes,1,opt,name=cardId,proto3" json:"cardId,omitempty"`
}

func (x *UnassignOwnerRequest) Reset() {
	*x = UnassignOwnerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_card_grpc_card_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnassignOwnerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnassignOwnerRequest) ProtoMessage() {}

func (x *UnassignOwnerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_card_grpc_card_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnassignOwnerRequest.ProtoReflect.Descriptor instead.
func (*UnassignOwnerRequest) Descriptor() ([]byte, []int) {
	return file_card_grpc_card_proto_rawDescGZIP(), []int{3}
}

func (x *UnassignOwnerRequest) GetCardId() string {
	if x != nil {
		return x.CardId
	}
	return ""
}

var File_card_grpc_card_proto protoreflect.FileDescriptor

var file_card_grpc_card_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x61, 0x72, 0x64, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x61, 0x72, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x46, 0x0a,
	0x12, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x14, 0x55, 0x6e, 0x61, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x63, 0x61, 0x72, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63,
	0x61, 0x72, 0x64, 0x49, 0x64, 0x32, 0x8a, 0x01, 0x0a, 0x05, 0x43, 0x61, 0x72, 0x64, 0x73, 0x12,
	0x27, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x12, 0x12, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x05, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x12, 0x29, 0x0a, 0x0b, 0x41, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x05, 0x2e, 0x43,
	0x61, 0x72, 0x64, 0x12, 0x2d, 0x0a, 0x0d, 0x55, 0x6e, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x4f,
	0x77, 0x6e, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x55, 0x6e, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x4f,
	0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x05, 0x2e, 0x43, 0x61,
	0x72, 0x64, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x74, 0x75, 0x72, 0x61, 0x6f, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x61, 0x72, 0x64, 0x73, 0x2f,
	0x63, 0x61, 0x72, 0x64, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_card_grpc_card_proto_rawDescOnce sync.Once
	file_card_grpc_card_proto_rawDescData = file_card_grpc_card_proto_rawDesc
)

func file_card_grpc_card_proto_rawDescGZIP() []byte {
	file_card_grpc_card_proto_rawDescOnce.Do(func() {
		file_card_grpc_card_proto_rawDescData = protoimpl.X.CompressGZIP(file_card_grpc_card_proto_rawDescData)
	})
	return file_card_grpc_card_proto_rawDescData
}

var file_card_grpc_card_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_card_grpc_card_proto_goTypes = []interface{}{
	(*Card)(nil),                 // 0: Card
	(*CreateCardRequest)(nil),    // 1: CreateCardRequest
	(*AssignOwnerRequest)(nil),   // 2: AssignOwnerRequest
	(*UnassignOwnerRequest)(nil), // 3: UnassignOwnerRequest
}
var file_card_grpc_card_proto_depIdxs = []int32{
	1, // 0: Cards.CreateCard:input_type -> CreateCardRequest
	2, // 1: Cards.AssignOwner:input_type -> AssignOwnerRequest
	3, // 2: Cards.UnassignOwner:input_type -> UnassignOwnerRequest
	0, // 3: Cards.CreateCard:output_type -> Card
	0, // 4: Cards.AssignOwner:output_type -> Card
	0, // 5: Cards.UnassignOwner:output_type -> Card
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_card_grpc_card_proto_init() }
func file_card_grpc_card_proto_init() {
	if File_card_grpc_card_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_card_grpc_card_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Card); i {
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
		file_card_grpc_card_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCardRequest); i {
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
		file_card_grpc_card_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssignOwnerRequest); i {
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
		file_card_grpc_card_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnassignOwnerRequest); i {
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
			RawDescriptor: file_card_grpc_card_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_card_grpc_card_proto_goTypes,
		DependencyIndexes: file_card_grpc_card_proto_depIdxs,
		MessageInfos:      file_card_grpc_card_proto_msgTypes,
	}.Build()
	File_card_grpc_card_proto = out.File
	file_card_grpc_card_proto_rawDesc = nil
	file_card_grpc_card_proto_goTypes = nil
	file_card_grpc_card_proto_depIdxs = nil
}