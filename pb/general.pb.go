// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.19.5
// source: proto/general.proto

package pb

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

type CmdId int32

const (
	CmdId_None CmdId = 0
	// 错误提示
	CmdId_Error CmdId = 1
	// 心跳
	CmdId_HeartBeat CmdId = 2
	// ======================= 背包相关 ========================
	// 背包更变（后端调用）
	CmdId_BagChange CmdId = 100
)

// Enum value maps for CmdId.
var (
	CmdId_name = map[int32]string{
		0:   "None",
		1:   "Error",
		2:   "HeartBeat",
		100: "BagChange",
	}
	CmdId_value = map[string]int32{
		"None":      0,
		"Error":     1,
		"HeartBeat": 2,
		"BagChange": 100,
	}
)

func (x CmdId) Enum() *CmdId {
	p := new(CmdId)
	*p = x
	return p
}

func (x CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_general_proto_enumTypes[0].Descriptor()
}

func (CmdId) Type() protoreflect.EnumType {
	return &file_proto_general_proto_enumTypes[0]
}

func (x CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CmdId.Descriptor instead.
func (CmdId) EnumDescriptor() ([]byte, []int) {
	return file_proto_general_proto_rawDescGZIP(), []int{0}
}

type Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cmd  CmdId  `protobuf:"varint,1,opt,name=cmd,proto3,enum=proto.CmdId" json:"cmd,omitempty"`
	Body []byte `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Time int64  `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *Msg) Reset() {
	*x = Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_general_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msg) ProtoMessage() {}

func (x *Msg) ProtoReflect() protoreflect.Message {
	mi := &file_proto_general_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Msg.ProtoReflect.Descriptor instead.
func (*Msg) Descriptor() ([]byte, []int) {
	return file_proto_general_proto_rawDescGZIP(), []int{0}
}

func (x *Msg) GetCmd() CmdId {
	if x != nil {
		return x.Cmd
	}
	return CmdId_None
}

func (x *Msg) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *Msg) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

var File_proto_general_proto protoreflect.FileDescriptor

var file_proto_general_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x03,
	0x4d, 0x73, 0x67, 0x12, 0x1e, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x52, 0x03,
	0x63, 0x6d, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x2a, 0x3a, 0x0a, 0x05, 0x43,
	0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x09,
	0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x48, 0x65, 0x61,
	0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x61, 0x67, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x10, 0x64, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_general_proto_rawDescOnce sync.Once
	file_proto_general_proto_rawDescData = file_proto_general_proto_rawDesc
)

func file_proto_general_proto_rawDescGZIP() []byte {
	file_proto_general_proto_rawDescOnce.Do(func() {
		file_proto_general_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_general_proto_rawDescData)
	})
	return file_proto_general_proto_rawDescData
}

var file_proto_general_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_general_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_general_proto_goTypes = []interface{}{
	(CmdId)(0),  // 0: proto.CmdId
	(*Msg)(nil), // 1: proto.Msg
}
var file_proto_general_proto_depIdxs = []int32{
	0, // 0: proto.Msg.cmd:type_name -> proto.CmdId
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_general_proto_init() }
func file_proto_general_proto_init() {
	if File_proto_general_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_general_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Msg); i {
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
			RawDescriptor: file_proto_general_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_general_proto_goTypes,
		DependencyIndexes: file_proto_general_proto_depIdxs,
		EnumInfos:         file_proto_general_proto_enumTypes,
		MessageInfos:      file_proto_general_proto_msgTypes,
	}.Build()
	File_proto_general_proto = out.File
	file_proto_general_proto_rawDesc = nil
	file_proto_general_proto_goTypes = nil
	file_proto_general_proto_depIdxs = nil
}
