// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: menu.proto

package protobuf

import (
	protobuf "github.com/Etpmls/Etpmls-Micro/protobuf"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type MenuCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Menu string `protobuf:"bytes,1,opt,name=menu,proto3" json:"menu,omitempty"`
}

func (x *MenuCreate) Reset() {
	*x = MenuCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuCreate) ProtoMessage() {}

func (x *MenuCreate) ProtoReflect() protoreflect.Message {
	mi := &file_menu_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuCreate.ProtoReflect.Descriptor instead.
func (*MenuCreate) Descriptor() ([]byte, []int) {
	return file_menu_proto_rawDescGZIP(), []int{0}
}

func (x *MenuCreate) GetMenu() string {
	if x != nil {
		return x.Menu
	}
	return ""
}

var File_menu_proto protoreflect.FileDescriptor

var file_menu_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x1a, 0x20, 0x45, 0x74, 0x70, 0x6d, 0x6c, 0x73, 0x2f, 0x45,
	0x74, 0x70, 0x6d, 0x6c, 0x73, 0x2d, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x0a, 0x4d, 0x65, 0x6e, 0x75, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6e, 0x75, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6e, 0x75, 0x32, 0xaf, 0x01, 0x0a, 0x04, 0x4d, 0x65, 0x6e,
	0x75, 0x12, 0x50, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x12, 0x2e, 0x65, 0x6d,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x15, 0x2e, 0x65, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x2f, 0x67, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x12, 0x55, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x1a, 0x15, 0x2e, 0x65, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x18, 0x22, 0x13, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6e, 0x75,
	0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_menu_proto_rawDescOnce sync.Once
	file_menu_proto_rawDescData = file_menu_proto_rawDesc
)

func file_menu_proto_rawDescGZIP() []byte {
	file_menu_proto_rawDescOnce.Do(func() {
		file_menu_proto_rawDescData = protoimpl.X.CompressGZIP(file_menu_proto_rawDescData)
	})
	return file_menu_proto_rawDescData
}

var file_menu_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_menu_proto_goTypes = []interface{}{
	(*MenuCreate)(nil),        // 0: protobuf.MenuCreate
	(*protobuf.Empty)(nil),    // 1: em_protobuf.Empty
	(*protobuf.Response)(nil), // 2: em_protobuf.Response
}
var file_menu_proto_depIdxs = []int32{
	1, // 0: protobuf.Menu.GetAll:input_type -> em_protobuf.Empty
	0, // 1: protobuf.Menu.Create:input_type -> protobuf.MenuCreate
	2, // 2: protobuf.Menu.GetAll:output_type -> em_protobuf.Response
	2, // 3: protobuf.Menu.Create:output_type -> em_protobuf.Response
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_menu_proto_init() }
func file_menu_proto_init() {
	if File_menu_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_menu_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuCreate); i {
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
			RawDescriptor: file_menu_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_menu_proto_goTypes,
		DependencyIndexes: file_menu_proto_depIdxs,
		MessageInfos:      file_menu_proto_msgTypes,
	}.Build()
	File_menu_proto = out.File
	file_menu_proto_rawDesc = nil
	file_menu_proto_goTypes = nil
	file_menu_proto_depIdxs = nil
}
