// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.3
// source: idl/pir.proto

package pir

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

// protoc -I. --go_out=:D:/gopath/src ./idl/*.proto
type PirQueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhoneNumber uint64 `protobuf:"varint,1,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"` // 注册用户名，最长32个字符
}

func (x *PirQueryRequest) Reset() {
	*x = PirQueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_pir_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PirQueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PirQueryRequest) ProtoMessage() {}

func (x *PirQueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_pir_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PirQueryRequest.ProtoReflect.Descriptor instead.
func (*PirQueryRequest) Descriptor() ([]byte, []int) {
	return file_idl_pir_proto_rawDescGZIP(), []int{0}
}

func (x *PirQueryRequest) GetPhoneNumber() uint64 {
	if x != nil {
		return x.PhoneNumber
	}
	return 0
}

type PirQueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32   `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`   // 状态码，0-成功，其他值-失败
	StatusMsg  *string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3,oneof" json:"status_msg,omitempty"` // 返回状态描述
}

func (x *PirQueryResponse) Reset() {
	*x = PirQueryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_pir_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PirQueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PirQueryResponse) ProtoMessage() {}

func (x *PirQueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_pir_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PirQueryResponse.ProtoReflect.Descriptor instead.
func (*PirQueryResponse) Descriptor() ([]byte, []int) {
	return file_idl_pir_proto_rawDescGZIP(), []int{1}
}

func (x *PirQueryResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *PirQueryResponse) GetStatusMsg() string {
	if x != nil && x.StatusMsg != nil {
		return *x.StatusMsg
	}
	return ""
}

type PirRefreshRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OptionMsg *string `protobuf:"bytes,1,opt,name=option_msg,json=optionMsg,proto3,oneof" json:"option_msg,omitempty"` // 可选信息
}

func (x *PirRefreshRequest) Reset() {
	*x = PirRefreshRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_pir_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PirRefreshRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PirRefreshRequest) ProtoMessage() {}

func (x *PirRefreshRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_pir_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PirRefreshRequest.ProtoReflect.Descriptor instead.
func (*PirRefreshRequest) Descriptor() ([]byte, []int) {
	return file_idl_pir_proto_rawDescGZIP(), []int{2}
}

func (x *PirRefreshRequest) GetOptionMsg() string {
	if x != nil && x.OptionMsg != nil {
		return *x.OptionMsg
	}
	return ""
}

type PirRefreshResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32     `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Data       []*Matrix `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`                                  // “提示”
	StatusMsg  *string   `protobuf:"bytes,3,opt,name=status_msg,json=statusMsg,proto3,oneof" json:"status_msg,omitempty"` // 返回状态描述
}

func (x *PirRefreshResponse) Reset() {
	*x = PirRefreshResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_pir_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PirRefreshResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PirRefreshResponse) ProtoMessage() {}

func (x *PirRefreshResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_pir_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PirRefreshResponse.ProtoReflect.Descriptor instead.
func (*PirRefreshResponse) Descriptor() ([]byte, []int) {
	return file_idl_pir_proto_rawDescGZIP(), []int{3}
}

func (x *PirRefreshResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *PirRefreshResponse) GetData() []*Matrix {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PirRefreshResponse) GetStatusMsg() string {
	if x != nil && x.StatusMsg != nil {
		return *x.StatusMsg
	}
	return ""
}

type Matrix struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rows uint64   `protobuf:"varint,1,opt,name=rows,proto3" json:"rows,omitempty"`
	Cols uint64   `protobuf:"varint,2,opt,name=cols,proto3" json:"cols,omitempty"`
	Data []uint64 `protobuf:"varint,3,rep,packed,name=data,proto3" json:"data,omitempty"`
}

func (x *Matrix) Reset() {
	*x = Matrix{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_pir_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Matrix) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Matrix) ProtoMessage() {}

func (x *Matrix) ProtoReflect() protoreflect.Message {
	mi := &file_idl_pir_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Matrix.ProtoReflect.Descriptor instead.
func (*Matrix) Descriptor() ([]byte, []int) {
	return file_idl_pir_proto_rawDescGZIP(), []int{4}
}

func (x *Matrix) GetRows() uint64 {
	if x != nil {
		return x.Rows
	}
	return 0
}

func (x *Matrix) GetCols() uint64 {
	if x != nil {
		return x.Cols
	}
	return 0
}

func (x *Matrix) GetData() []uint64 {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_idl_pir_proto protoreflect.FileDescriptor

var file_idl_pir_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x69, 0x64, 0x6c, 0x2f, 0x70, 0x69, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x70, 0x69, 0x72, 0x22, 0x36, 0x0a, 0x11, 0x70, 0x69, 0x72, 0x5f, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x68, 0x0a, 0x12,
	0x70, 0x69, 0x72, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x4d, 0x73, 0x67, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x22, 0x48, 0x0a, 0x13, 0x70, 0x69, 0x72, 0x5f, 0x72, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a,
	0x0a, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x09, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x73, 0x67, 0x88, 0x01,
	0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x73, 0x67,
	0x22, 0x8b, 0x01, 0x0a, 0x14, 0x70, 0x69, 0x72, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x69, 0x72, 0x2e, 0x4d,
	0x61, 0x74, 0x72, 0x69, 0x78, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x88, 0x01, 0x01, 0x42,
	0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x22, 0x44,
	0x0a, 0x06, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x63, 0x6f, 0x6c, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x04, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x32, 0x88, 0x01, 0x0a, 0x06, 0x50, 0x69, 0x72, 0x53, 0x72, 0x76, 0x12,
	0x3c, 0x0a, 0x07, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x69, 0x12, 0x16, 0x2e, 0x70, 0x69, 0x72,
	0x2e, 0x70, 0x69, 0x72, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x69, 0x72, 0x2e, 0x70, 0x69, 0x72, 0x5f, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a,
	0x07, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x12, 0x18, 0x2e, 0x70, 0x69, 0x72, 0x2e, 0x70,
	0x69, 0x72, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x69, 0x72, 0x2e, 0x70, 0x69, 0x72, 0x5f, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x66,
	0x2d, 0x30, 0x31, 0x31, 0x31, 0x30, 0x31, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x69,
	0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x69, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_idl_pir_proto_rawDescOnce sync.Once
	file_idl_pir_proto_rawDescData = file_idl_pir_proto_rawDesc
)

func file_idl_pir_proto_rawDescGZIP() []byte {
	file_idl_pir_proto_rawDescOnce.Do(func() {
		file_idl_pir_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_pir_proto_rawDescData)
	})
	return file_idl_pir_proto_rawDescData
}

var file_idl_pir_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_idl_pir_proto_goTypes = []interface{}{
	(*PirQueryRequest)(nil),    // 0: pir.pir_query_request
	(*PirQueryResponse)(nil),   // 1: pir.pir_query_response
	(*PirRefreshRequest)(nil),  // 2: pir.pir_refresh_request
	(*PirRefreshResponse)(nil), // 3: pir.pir_refresh_response
	(*Matrix)(nil),             // 4: pir.Matrix
}
var file_idl_pir_proto_depIdxs = []int32{
	4, // 0: pir.pir_refresh_response.data:type_name -> pir.Matrix
	0, // 1: pir.PirSrv.QueryPi:input_type -> pir.pir_query_request
	2, // 2: pir.PirSrv.Refresh:input_type -> pir.pir_refresh_request
	1, // 3: pir.PirSrv.QueryPi:output_type -> pir.pir_query_response
	3, // 4: pir.PirSrv.Refresh:output_type -> pir.pir_refresh_response
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_idl_pir_proto_init() }
func file_idl_pir_proto_init() {
	if File_idl_pir_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idl_pir_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PirQueryRequest); i {
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
		file_idl_pir_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PirQueryResponse); i {
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
		file_idl_pir_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PirRefreshRequest); i {
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
		file_idl_pir_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PirRefreshResponse); i {
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
		file_idl_pir_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Matrix); i {
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
	file_idl_pir_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_idl_pir_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_idl_pir_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_idl_pir_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idl_pir_proto_goTypes,
		DependencyIndexes: file_idl_pir_proto_depIdxs,
		MessageInfos:      file_idl_pir_proto_msgTypes,
	}.Build()
	File_idl_pir_proto = out.File
	file_idl_pir_proto_rawDesc = nil
	file_idl_pir_proto_goTypes = nil
	file_idl_pir_proto_depIdxs = nil
}
