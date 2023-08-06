// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: types/api_chat.proto

package api

import (
	common "github.com/cd-home/Hissssss/api/pb/common"
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

type SendMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From int64           `protobuf:"varint,1,opt,name=from,proto3" json:"from,omitempty"`                      // 发送者
	To   int64           `protobuf:"varint,2,opt,name=to,proto3" json:"to,omitempty"`                          // 接收者
	Room int64           `protobuf:"varint,3,opt,name=room,proto3" json:"room,omitempty"`                      // 房间接收
	Body string          `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`                       // 消息体
	Type common.PushType `protobuf:"varint,5,opt,name=type,proto3,enum=common.PushType" json:"type,omitempty"` // 发送类型
}

func (x *SendMessageRequest) Reset() {
	*x = SendMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_api_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageRequest) ProtoMessage() {}

func (x *SendMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_types_api_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageRequest.ProtoReflect.Descriptor instead.
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return file_types_api_chat_proto_rawDescGZIP(), []int{0}
}

func (x *SendMessageRequest) GetFrom() int64 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *SendMessageRequest) GetTo() int64 {
	if x != nil {
		return x.To
	}
	return 0
}

func (x *SendMessageRequest) GetRoom() int64 {
	if x != nil {
		return x.Room
	}
	return 0
}

func (x *SendMessageRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *SendMessageRequest) GetType() common.PushType {
	if x != nil {
		return x.Type
	}
	return common.PushType(0)
}

type SendMessageReplyAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  int64     `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg   string    `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	MsgId int64     `protobuf:"varint,3,opt,name=msgId,proto3" json:"msgId,omitempty"`          // 消息id 返回双方的对话的全局消息id
	Op    common.OP `protobuf:"varint,4,opt,name=op,proto3,enum=common.OP" json:"op,omitempty"` // 消息动作 确认动作, 服务器确认已经收到消息, 正在处理中
}

func (x *SendMessageReplyAck) Reset() {
	*x = SendMessageReplyAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_api_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageReplyAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageReplyAck) ProtoMessage() {}

func (x *SendMessageReplyAck) ProtoReflect() protoreflect.Message {
	mi := &file_types_api_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageReplyAck.ProtoReflect.Descriptor instead.
func (*SendMessageReplyAck) Descriptor() ([]byte, []int) {
	return file_types_api_chat_proto_rawDescGZIP(), []int{1}
}

func (x *SendMessageReplyAck) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SendMessageReplyAck) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SendMessageReplyAck) GetMsgId() int64 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

func (x *SendMessageReplyAck) GetOp() common.OP {
	if x != nil {
		return x.Op
	}
	return common.OP(0)
}

// AckSingleMsgReqeust 收到消息的确认请求
type AckSingleMsgReqeust struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From  int64     `protobuf:"varint,1,opt,name=from,proto3" json:"from,omitempty"`
	To    int64     `protobuf:"varint,2,opt,name=to,proto3" json:"to,omitempty"`
	MsgId int64     `protobuf:"varint,3,opt,name=msgId,proto3" json:"msgId,omitempty"`
	Op    common.OP `protobuf:"varint,4,opt,name=op,proto3,enum=common.OP" json:"op,omitempty"`
}

func (x *AckSingleMsgReqeust) Reset() {
	*x = AckSingleMsgReqeust{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_api_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckSingleMsgReqeust) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckSingleMsgReqeust) ProtoMessage() {}

func (x *AckSingleMsgReqeust) ProtoReflect() protoreflect.Message {
	mi := &file_types_api_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckSingleMsgReqeust.ProtoReflect.Descriptor instead.
func (*AckSingleMsgReqeust) Descriptor() ([]byte, []int) {
	return file_types_api_chat_proto_rawDescGZIP(), []int{2}
}

func (x *AckSingleMsgReqeust) GetFrom() int64 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *AckSingleMsgReqeust) GetTo() int64 {
	if x != nil {
		return x.To
	}
	return 0
}

func (x *AckSingleMsgReqeust) GetMsgId() int64 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

func (x *AckSingleMsgReqeust) GetOp() common.OP {
	if x != nil {
		return x.Op
	}
	return common.OP(0)
}

// AckSingleMsgReplyAck 服务端确认收到
type AckSingleMsgReplyAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From  int64     `protobuf:"varint,1,opt,name=from,proto3" json:"from,omitempty"`
	To    int64     `protobuf:"varint,2,opt,name=to,proto3" json:"to,omitempty"`
	MsgId int64     `protobuf:"varint,3,opt,name=msgId,proto3" json:"msgId,omitempty"`
	Op    common.OP `protobuf:"varint,4,opt,name=op,proto3,enum=common.OP" json:"op,omitempty"`
}

func (x *AckSingleMsgReplyAck) Reset() {
	*x = AckSingleMsgReplyAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_api_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckSingleMsgReplyAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckSingleMsgReplyAck) ProtoMessage() {}

func (x *AckSingleMsgReplyAck) ProtoReflect() protoreflect.Message {
	mi := &file_types_api_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckSingleMsgReplyAck.ProtoReflect.Descriptor instead.
func (*AckSingleMsgReplyAck) Descriptor() ([]byte, []int) {
	return file_types_api_chat_proto_rawDescGZIP(), []int{3}
}

func (x *AckSingleMsgReplyAck) GetFrom() int64 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *AckSingleMsgReplyAck) GetTo() int64 {
	if x != nil {
		return x.To
	}
	return 0
}

func (x *AckSingleMsgReplyAck) GetMsgId() int64 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

func (x *AckSingleMsgReplyAck) GetOp() common.OP {
	if x != nil {
		return x.Op
	}
	return common.OP(0)
}

var File_types_api_chat_proto protoreflect.FileDescriptor

var file_types_api_chat_proto_rawDesc = []byte{
	0x0a, 0x14, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x0c, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x01, 0x0a, 0x12, 0x53, 0x65,
	0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x24, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x22, 0x6d, 0x0a, 0x13, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x41, 0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12,
	0x14, 0x0a, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x6d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x50, 0x52, 0x02, 0x6f,
	0x70, 0x22, 0x6b, 0x0a, 0x13, 0x41, 0x63, 0x6b, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x4d, 0x73,
	0x67, 0x52, 0x65, 0x71, 0x65, 0x75, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02,
	0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05,
	0x6d, 0x73, 0x67, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6d, 0x73, 0x67,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x50, 0x52, 0x02, 0x6f, 0x70, 0x22, 0x6c,
	0x0a, 0x14, 0x41, 0x63, 0x6b, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x4d, 0x73, 0x67, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x41, 0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x73,
	0x67, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x50, 0x52, 0x02, 0x6f, 0x70, 0x42, 0x0b, 0x5a, 0x09,
	0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_types_api_chat_proto_rawDescOnce sync.Once
	file_types_api_chat_proto_rawDescData = file_types_api_chat_proto_rawDesc
)

func file_types_api_chat_proto_rawDescGZIP() []byte {
	file_types_api_chat_proto_rawDescOnce.Do(func() {
		file_types_api_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_api_chat_proto_rawDescData)
	})
	return file_types_api_chat_proto_rawDescData
}

var file_types_api_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_types_api_chat_proto_goTypes = []interface{}{
	(*SendMessageRequest)(nil),   // 0: api.SendMessageRequest
	(*SendMessageReplyAck)(nil),  // 1: api.SendMessageReplyAck
	(*AckSingleMsgReqeust)(nil),  // 2: api.AckSingleMsgReqeust
	(*AckSingleMsgReplyAck)(nil), // 3: api.AckSingleMsgReplyAck
	(common.PushType)(0),         // 4: common.PushType
	(common.OP)(0),               // 5: common.OP
}
var file_types_api_chat_proto_depIdxs = []int32{
	4, // 0: api.SendMessageRequest.type:type_name -> common.PushType
	5, // 1: api.SendMessageReplyAck.op:type_name -> common.OP
	5, // 2: api.AckSingleMsgReqeust.op:type_name -> common.OP
	5, // 3: api.AckSingleMsgReplyAck.op:type_name -> common.OP
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_types_api_chat_proto_init() }
func file_types_api_chat_proto_init() {
	if File_types_api_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_types_api_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageRequest); i {
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
		file_types_api_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageReplyAck); i {
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
		file_types_api_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AckSingleMsgReqeust); i {
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
		file_types_api_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AckSingleMsgReplyAck); i {
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
			RawDescriptor: file_types_api_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_api_chat_proto_goTypes,
		DependencyIndexes: file_types_api_chat_proto_depIdxs,
		MessageInfos:      file_types_api_chat_proto_msgTypes,
	}.Build()
	File_types_api_chat_proto = out.File
	file_types_api_chat_proto_rawDesc = nil
	file_types_api_chat_proto_goTypes = nil
	file_types_api_chat_proto_depIdxs = nil
}