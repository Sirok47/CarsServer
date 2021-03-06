// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: protocol/cars.proto

package tutorialpb

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

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_cars_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_cars_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_protocol_cars_proto_rawDescGZIP(), []int{0}
}

func (x *Token) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type Userdata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nick     string `protobuf:"bytes,1,opt,name=Nick,proto3" json:"Nick,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *Userdata) Reset() {
	*x = Userdata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_cars_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Userdata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Userdata) ProtoMessage() {}

func (x *Userdata) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_cars_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Userdata.ProtoReflect.Descriptor instead.
func (*Userdata) Descriptor() ([]byte, []int) {
	return file_protocol_cars_proto_rawDescGZIP(), []int{1}
}

func (x *Userdata) GetNick() string {
	if x != nil {
		return x.Nick
	}
	return ""
}

func (x *Userdata) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Carparams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarBrand  string `protobuf:"bytes,1,opt,name=CarBrand,proto3" json:"CarBrand,omitempty"`
	CarNumber int32  `protobuf:"varint,2,opt,name=CarNumber,proto3" json:"CarNumber,omitempty"`
	Mileage   int32  `protobuf:"varint,3,opt,name=Mileage,proto3" json:"Mileage,omitempty"`
	CarType   string `protobuf:"bytes,4,opt,name=CarType,proto3" json:"CarType,omitempty"`
}

func (x *Carparams) Reset() {
	*x = Carparams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_cars_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Carparams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Carparams) ProtoMessage() {}

func (x *Carparams) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_cars_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Carparams.ProtoReflect.Descriptor instead.
func (*Carparams) Descriptor() ([]byte, []int) {
	return file_protocol_cars_proto_rawDescGZIP(), []int{2}
}

func (x *Carparams) GetCarBrand() string {
	if x != nil {
		return x.CarBrand
	}
	return ""
}

func (x *Carparams) GetCarNumber() int32 {
	if x != nil {
		return x.CarNumber
	}
	return 0
}

func (x *Carparams) GetMileage() int32 {
	if x != nil {
		return x.Mileage
	}
	return 0
}

func (x *Carparams) GetCarType() string {
	if x != nil {
		return x.CarType
	}
	return ""
}

type Errmsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Errmsg) Reset() {
	*x = Errmsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_cars_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Errmsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Errmsg) ProtoMessage() {}

func (x *Errmsg) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_cars_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Errmsg.ProtoReflect.Descriptor instead.
func (*Errmsg) Descriptor() ([]byte, []int) {
	return file_protocol_cars_proto_rawDescGZIP(), []int{3}
}

func (x *Errmsg) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_protocol_cars_proto protoreflect.FileDescriptor

var file_protocol_cars_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x61, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22,
	0x1d, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x69,
	0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x69, 0x63, 0x6b, 0x12, 0x1a,
	0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x79, 0x0a, 0x09, 0x63, 0x61,
	0x72, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x61, 0x72, 0x42, 0x72,
	0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x61, 0x72, 0x42, 0x72,
	0x61, 0x6e, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x61, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x43, 0x61, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x69, 0x6c, 0x65, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x4d, 0x69, 0x6c, 0x65, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43,
	0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x61,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x22, 0x1e, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xb4, 0x02, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x73, 0x12, 0x2e,
	0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x0f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x00, 0x12, 0x30,
	0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x10, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x22, 0x00,
	0x12, 0x31, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x63, 0x61, 0x72, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a,
	0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x65, 0x72, 0x72, 0x6d, 0x73,
	0x67, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x63, 0x61, 0x72, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a,
	0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x63, 0x61, 0x72, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x63, 0x61, 0x72, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x06, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x63,
	0x61, 0x72, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x22, 0x00, 0x42, 0x3c, 0x5a, 0x3a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x67, 0x6f, 0x2f,
	0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_protocol_cars_proto_rawDescOnce sync.Once
	file_protocol_cars_proto_rawDescData = file_protocol_cars_proto_rawDesc
)

func file_protocol_cars_proto_rawDescGZIP() []byte {
	file_protocol_cars_proto_rawDescOnce.Do(func() {
		file_protocol_cars_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_cars_proto_rawDescData)
	})
	return file_protocol_cars_proto_rawDescData
}

var file_protocol_cars_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protocol_cars_proto_goTypes = []interface{}{
	(*Token)(nil),     // 0: protocol.token
	(*Userdata)(nil),  // 1: protocol.userdata
	(*Carparams)(nil), // 2: protocol.carparams
	(*Errmsg)(nil),    // 3: protocol.errmsg
}
var file_protocol_cars_proto_depIdxs = []int32{
	1, // 0: protocol.Cars.LogIn:input_type -> protocol.userdata
	1, // 1: protocol.Cars.SignUp:input_type -> protocol.userdata
	2, // 2: protocol.Cars.Create:input_type -> protocol.carparams
	2, // 3: protocol.Cars.Get:input_type -> protocol.carparams
	2, // 4: protocol.Cars.Delete:input_type -> protocol.carparams
	2, // 5: protocol.Cars.Update:input_type -> protocol.carparams
	0, // 6: protocol.Cars.LogIn:output_type -> protocol.token
	3, // 7: protocol.Cars.SignUp:output_type -> protocol.errmsg
	3, // 8: protocol.Cars.Create:output_type -> protocol.errmsg
	2, // 9: protocol.Cars.Get:output_type -> protocol.carparams
	3, // 10: protocol.Cars.Delete:output_type -> protocol.errmsg
	3, // 11: protocol.Cars.Update:output_type -> protocol.errmsg
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protocol_cars_proto_init() }
func file_protocol_cars_proto_init() {
	if File_protocol_cars_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocol_cars_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_protocol_cars_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Userdata); i {
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
		file_protocol_cars_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Carparams); i {
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
		file_protocol_cars_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Errmsg); i {
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
			RawDescriptor: file_protocol_cars_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protocol_cars_proto_goTypes,
		DependencyIndexes: file_protocol_cars_proto_depIdxs,
		MessageInfos:      file_protocol_cars_proto_msgTypes,
	}.Build()
	File_protocol_cars_proto = out.File
	file_protocol_cars_proto_rawDesc = nil
	file_protocol_cars_proto_goTypes = nil
	file_protocol_cars_proto_depIdxs = nil
}
