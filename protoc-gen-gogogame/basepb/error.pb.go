// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/error.proto

package basepb

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	github_com_mailru_easyjson_buffer "github.com/mailru/easyjson/buffer"
	github_com_mailru_easyjson_jwriter "github.com/mailru/easyjson/jwriter"
	github_com_ravinggo_tools_jsonany "github.com/ravinggo/tools/jsonany"
	google_golang_org_protobuf_reflect_protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	io "io"
	math "math"
	math_bits "math/bits"
	strconv "strconv"
	unsafe "unsafe"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ErrorType int32

const (
	ErrorType_ETNo     ErrorType = 0
	ErrorType_ETNormal ErrorType = 1
	// the client can handle the following error types uniformly
	// ETProtocol,ETPanic,ETDataBase can be unified as an internal server error
	// ETNoAuth not login, you can treat it as disconnected and then jump to the login interface
	ErrorType_ETProtocol ErrorType = 2
	ErrorType_ETPanic    ErrorType = 3
	ErrorType_ETDataBase ErrorType = 4
	ErrorType_ETNoAuth   ErrorType = 5
)

var ErrorType_name = map[int32]string{
	0: "ETNo",
	1: "ETNormal",
	2: "ETProtocol",
	3: "ETPanic",
	4: "ETDataBase",
	5: "ETNoAuth",
}

var ErrorType_value = map[string]int32{
	"ETNo":       0,
	"ETNormal":   1,
	"ETProtocol": 2,
	"ETPanic":    3,
	"ETDataBase": 4,
	"ETNoAuth":   5,
}

func (ErrorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_faef5ce421884d5a, []int{0}
}

// just for error define and code generate
type ErrorInfo struct {
	TipName   string `protobuf:"bytes,1,opt,name=tip_name,json=tipName,proto3" json:"tip_name,omitempty"`
	TipDesc   string `protobuf:"bytes,2,opt,name=tip_desc,json=tipDesc,proto3" json:"tip_desc,omitempty"`
	ErrorCode int64  `protobuf:"varint,3,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
}

func (m *ErrorInfo) Reset()      { *m = ErrorInfo{} }
func (*ErrorInfo) ProtoMessage() {}
func (*ErrorInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_faef5ce421884d5a, []int{0}
}
func (m *ErrorInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ErrorInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ErrorInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ErrorInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorInfo.Merge(m, src)
}
func (m *ErrorInfo) XXX_Size() int {
	return m.Size()
}
func (m *ErrorInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorInfo proto.InternalMessageInfo

func (m *ErrorInfo) GetTipName() string {
	if m != nil {
		return m.TipName
	}
	return ""
}

func (m *ErrorInfo) GetTipDesc() string {
	if m != nil {
		return m.TipDesc
	}
	return ""
}

func (m *ErrorInfo) GetErrorCode() int64 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (*ErrorInfo) XXX_MessageName() string {
	return "basepb.ErrorInfo"
}

type ErrorMessage struct {
	ErrCode ErrorType `protobuf:"varint,1,opt,name=err_code,json=errCode,proto3,enum=basepb.ErrorType" json:"err_code,omitempty"`
	// every error unique flag on same ErrorType
	ErrMsg string `protobuf:"bytes,2,opt,name=err_msg,json=errMsg,proto3" json:"err_msg,omitempty"`
	// server error detail, only for log, empty for client
	ErrInternalInfo string `protobuf:"bytes,3,opt,name=err_internal_info,json=errInternalInfo,proto3" json:"err_internal_info,omitempty"`
	// server error stack stace, only for log, empty for client
	StackStace []uint64 `protobuf:"varint,4,rep,packed,name=stack_stace,json=stackStace,proto3" json:"stack_stace,omitempty"`
	// extra for error. For example, if there is an insufficient item error, this field can be the specific item ID.
	ErrExtraInfo string `protobuf:"bytes,5,opt,name=err_extra_info,json=errExtraInfo,proto3" json:"err_extra_info,omitempty"`
}

func (m *ErrorMessage) Reset()      { *m = ErrorMessage{} }
func (*ErrorMessage) ProtoMessage() {}
func (*ErrorMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_faef5ce421884d5a, []int{1}
}
func (m *ErrorMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ErrorMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ErrorMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ErrorMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorMessage.Merge(m, src)
}
func (m *ErrorMessage) XXX_Size() int {
	return m.Size()
}
func (m *ErrorMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorMessage proto.InternalMessageInfo

func (m *ErrorMessage) GetErrCode() ErrorType {
	if m != nil {
		return m.ErrCode
	}
	return ErrorType_ETNo
}

func (m *ErrorMessage) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *ErrorMessage) GetErrInternalInfo() string {
	if m != nil {
		return m.ErrInternalInfo
	}
	return ""
}

func (m *ErrorMessage) GetStackStace() []uint64 {
	if m != nil {
		return m.StackStace
	}
	return nil
}

func (m *ErrorMessage) GetErrExtraInfo() string {
	if m != nil {
		return m.ErrExtraInfo
	}
	return ""
}

func (*ErrorMessage) XXX_MessageName() string {
	return "basepb.ErrorMessage"
}

var E_ErrorInfo = &proto.ExtensionDesc{
	ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
	ExtensionType: (*ErrorInfo)(nil),
	Field:         50001,
	Name:          "basepb.error_info",
	Tag:           "bytes,50001,opt,name=error_info",
	Filename:      "proto/error.proto",
}

func init() {
	proto.RegisterEnum("basepb.ErrorType", ErrorType_name, ErrorType_value)
	proto.RegisterType((*ErrorInfo)(nil), "basepb.ErrorInfo")
	proto.RegisterType((*ErrorMessage)(nil), "basepb.ErrorMessage")
	proto.RegisterExtension(E_ErrorInfo)
}

func init() { proto.RegisterFile("proto/error.proto", fileDescriptor_faef5ce421884d5a) }

var fileDescriptor_faef5ce421884d5a = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x51, 0x31, 0x8f, 0xd3, 0x4c,
	0x10, 0xf5, 0x5e, 0x72, 0x89, 0xb3, 0x89, 0xf2, 0x39, 0xdb, 0x7c, 0x06, 0x89, 0x25, 0x9c, 0x28,
	0xa2, 0x13, 0x72, 0x24, 0x90, 0x28, 0xe8, 0x38, 0xce, 0xc5, 0x15, 0x77, 0x9c, 0x4c, 0xa0, 0xa0,
	0xc0, 0x6c, 0x9c, 0x89, 0xb1, 0xb0, 0xbd, 0xd6, 0xee, 0x46, 0x82, 0x3f, 0x40, 0xcd, 0xcf, 0xe0,
	0xa7, 0xa4, 0x3c, 0xba, 0x2b, 0xc1, 0x69, 0x28, 0xf9, 0x09, 0x68, 0xd6, 0x36, 0x05, 0xcd, 0x6a,
	0xf4, 0x66, 0xe6, 0xbd, 0xb7, 0x6f, 0xe8, 0xac, 0x52, 0xd2, 0xc8, 0x25, 0x28, 0x25, 0x55, 0x60,
	0x6b, 0x36, 0x58, 0x0b, 0x0d, 0xd5, 0xfa, 0xee, 0x3c, 0x95, 0x32, 0xcd, 0x61, 0x69, 0xd1, 0xf5,
	0x6e, 0xbb, 0xdc, 0x80, 0x4e, 0x54, 0x56, 0x99, 0x6e, 0xf2, 0xe4, 0x3d, 0x1d, 0x85, 0xb8, 0x78,
	0x51, 0x6e, 0x25, 0xbb, 0x43, 0x5d, 0x93, 0x55, 0x71, 0x29, 0x0a, 0xf0, 0xc9, 0x9c, 0x2c, 0x46,
	0xd1, 0xd0, 0x64, 0xd5, 0x95, 0x28, 0xa0, 0x6b, 0xe1, 0xbe, 0x7f, 0xf4, 0xb7, 0x75, 0x0e, 0x3a,
	0x61, 0xf7, 0x28, 0xb5, 0xda, 0x71, 0x22, 0x37, 0xe0, 0xf7, 0xe6, 0x64, 0xd1, 0x8b, 0x46, 0x16,
	0x79, 0x21, 0x37, 0x70, 0xb2, 0x27, 0x74, 0x62, 0x25, 0x2e, 0x41, 0x6b, 0x91, 0x02, 0x7b, 0x44,
	0x5d, 0x50, 0xed, 0x34, 0xaa, 0x4c, 0x1f, 0xcf, 0x82, 0xc6, 0x6f, 0x60, 0xe7, 0x56, 0x9f, 0x2b,
	0x88, 0x86, 0xa0, 0xec, 0x3a, 0xfb, 0x9f, 0x62, 0x19, 0x17, 0x3a, 0x6d, 0x75, 0x07, 0xa0, 0xd4,
	0xa5, 0x4e, 0xd9, 0x29, 0x9d, 0x61, 0x23, 0x2b, 0x0d, 0xa8, 0x52, 0xe4, 0x71, 0x56, 0x6e, 0xa5,
	0x55, 0x1f, 0x45, 0xff, 0x81, 0x52, 0x17, 0x2d, 0x6e, 0x3f, 0x76, 0x9f, 0x8e, 0xb5, 0x11, 0xc9,
	0xc7, 0x18, 0x5f, 0xf0, 0xfb, 0xf3, 0xde, 0xa2, 0x1f, 0x51, 0x0b, 0xbd, 0x42, 0x84, 0x3d, 0xa4,
	0x53, 0x24, 0x83, 0x4f, 0x46, 0x89, 0x86, 0xe9, 0xd8, 0x32, 0x4d, 0x40, 0xa9, 0x10, 0x41, 0xa4,
	0x39, 0x7d, 0xd7, 0x86, 0x85, 0x0e, 0x99, 0x4b, 0xfb, 0xe1, 0xea, 0x4a, 0x7a, 0x0e, 0x9b, 0x50,
	0x17, 0x2b, 0x55, 0x88, 0xdc, 0x23, 0x6c, 0x4a, 0x69, 0xb8, 0xba, 0xc6, 0x70, 0x13, 0x99, 0x7b,
	0x47, 0x6c, 0x4c, 0x87, 0xe1, 0xea, 0x5a, 0x94, 0x59, 0xe2, 0xf5, 0x9a, 0xe6, 0xb9, 0x30, 0xe2,
	0x4c, 0x68, 0xf0, 0xfa, 0xdd, 0xea, 0xf3, 0x9d, 0xf9, 0xe0, 0x1d, 0x3f, 0x7b, 0xdd, 0x25, 0x89,
	0x0e, 0xd8, 0x83, 0xa0, 0xb9, 0x5e, 0xd0, 0x5d, 0x2f, 0x08, 0xcb, 0x5d, 0xf1, 0x46, 0xe4, 0x3b,
	0x78, 0x59, 0x99, 0x4c, 0x96, 0xda, 0xff, 0xfe, 0x05, 0x3f, 0x3c, 0xfe, 0x27, 0x40, 0xf4, 0xda,
	0x5e, 0x00, 0xcb, 0xb3, 0xa7, 0xb7, 0x3f, 0xb9, 0xf3, 0xad, 0xe6, 0x64, 0x5f, 0x73, 0x72, 0x53,
	0x73, 0xf2, 0xa3, 0xe6, 0xe4, 0x57, 0xcd, 0x9d, 0xdf, 0x35, 0x27, 0x5f, 0x0f, 0xdc, 0xd9, 0x1f,
	0x38, 0xb9, 0x39, 0x70, 0xe7, 0xf6, 0xc0, 0x9d, 0xb7, 0x6e, 0xb0, 0x6c, 0xc8, 0xd6, 0x03, 0x2b,
	0xfc, 0xe4, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x86, 0xe5, 0x5e, 0x27, 0x61, 0x02, 0x00, 0x00,
}

func (x ErrorType) String() string {
	s, ok := ErrorType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *ErrorInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ErrorInfo)
	if !ok {
		that2, ok := that.(ErrorInfo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.TipName != that1.TipName {
		return false
	}
	if this.TipDesc != that1.TipDesc {
		return false
	}
	if this.ErrorCode != that1.ErrorCode {
		return false
	}
	return true
}
func (this *ErrorMessage) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ErrorMessage)
	if !ok {
		that2, ok := that.(ErrorMessage)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ErrCode != that1.ErrCode {
		return false
	}
	if this.ErrMsg != that1.ErrMsg {
		return false
	}
	if this.ErrInternalInfo != that1.ErrInternalInfo {
		return false
	}
	if len(this.StackStace) != len(that1.StackStace) {
		return false
	}
	for i := range this.StackStace {
		if this.StackStace[i] != that1.StackStace[i] {
			return false
		}
	}
	if this.ErrExtraInfo != that1.ErrExtraInfo {
		return false
	}
	return true
}
func (m *ErrorInfo) ProtoReflect() google_golang_org_protobuf_reflect_protoreflect.Message {
	return nil
}
func (m *ErrorMessage) ProtoReflect() google_golang_org_protobuf_reflect_protoreflect.Message {
	return nil
}
func (m *ErrorInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ErrorInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ErrorInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ErrorCode != 0 {
		i = encodeVarintError(dAtA, i, uint64(m.ErrorCode))
		i--
		dAtA[i] = 0x18
	}
	if len(m.TipDesc) > 0 {
		i -= len(m.TipDesc)
		copy(dAtA[i:], m.TipDesc)
		i = encodeVarintError(dAtA, i, uint64(len(m.TipDesc)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.TipName) > 0 {
		i -= len(m.TipName)
		copy(dAtA[i:], m.TipName)
		i = encodeVarintError(dAtA, i, uint64(len(m.TipName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ErrorMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ErrorMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ErrorMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ErrExtraInfo) > 0 {
		i -= len(m.ErrExtraInfo)
		copy(dAtA[i:], m.ErrExtraInfo)
		i = encodeVarintError(dAtA, i, uint64(len(m.ErrExtraInfo)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.StackStace) > 0 {
		dAtA2 := make([]byte, len(m.StackStace)*10)
		var j1 int
		for _, num := range m.StackStace {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintError(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ErrInternalInfo) > 0 {
		i -= len(m.ErrInternalInfo)
		copy(dAtA[i:], m.ErrInternalInfo)
		i = encodeVarintError(dAtA, i, uint64(len(m.ErrInternalInfo)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ErrMsg) > 0 {
		i -= len(m.ErrMsg)
		copy(dAtA[i:], m.ErrMsg)
		i = encodeVarintError(dAtA, i, uint64(len(m.ErrMsg)))
		i--
		dAtA[i] = 0x12
	}
	if m.ErrCode != 0 {
		i = encodeVarintError(dAtA, i, uint64(m.ErrCode))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintError(dAtA []byte, offset int, v uint64) int {
	offset -= sovError(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

var _ = github_com_ravinggo_tools_jsonany.Any{}

func (m *ErrorInfo) MarshalEasyJSON(w *github_com_mailru_easyjson_jwriter.Writer) {
	if m == nil {
		w.RawString("null")
		return
	}

	w.RawByte('{')
	needWriteComma := false
	if m.TipName != "" {
		w.RawByte('"')
		w.RawString("tip_name")
		w.RawByte('"')
		w.RawByte(':')
		w.String(m.TipName)
		needWriteComma = true
	}
	if m.TipDesc != "" {
		if needWriteComma {
			w.RawByte(',')
		}
		w.RawByte('"')
		w.RawString("tip_desc")
		w.RawByte('"')
		w.RawByte(':')
		w.String(m.TipDesc)
		needWriteComma = true
	}
	if m.ErrorCode != 0 {
		if needWriteComma {
			w.RawByte(',')
		}
		w.RawByte('"')
		w.RawString("error_code")
		w.RawByte('"')
		w.RawByte(':')
		w.Int64(int64(m.ErrorCode))
		needWriteComma = true
	}
	_ = needWriteComma
	w.RawByte('}')

}

func (m *ErrorMessage) MarshalEasyJSON(w *github_com_mailru_easyjson_jwriter.Writer) {
	if m == nil {
		w.RawString("null")
		return
	}

	w.RawByte('{')
	needWriteComma := false
	if m.ErrCode != 0 {
		w.RawByte('"')
		w.RawString("err_code")
		w.RawByte('"')
		w.RawByte(':')
		w.Int64(int64(m.ErrCode))
		needWriteComma = true
	}
	if m.ErrMsg != "" {
		if needWriteComma {
			w.RawByte(',')
		}
		w.RawByte('"')
		w.RawString("err_msg")
		w.RawByte('"')
		w.RawByte(':')
		w.String(m.ErrMsg)
		needWriteComma = true
	}
	if m.ErrInternalInfo != "" {
		if needWriteComma {
			w.RawByte(',')
		}
		w.RawByte('"')
		w.RawString("err_internal_info")
		w.RawByte('"')
		w.RawByte(':')
		w.String(m.ErrInternalInfo)
		needWriteComma = true
	}
	if needWriteComma {
		w.RawByte(',')
	}
	w.RawByte('"')
	w.RawString("stack_stace")
	w.RawByte('"')
	w.RawByte(':')
	if m.StackStace == nil {
		w.RawString("null")
	} else if len(m.StackStace) == 0 {
		w.RawString("[]")
	} else {
		w.RawByte('[')
		for i, v := range m.StackStace {
			w.Uint64(uint64(v))
			if i != len(m.StackStace)-1 {
				w.RawByte(',')
			}
		}
		w.RawByte(']')
	}
	needWriteComma = true
	if m.ErrExtraInfo != "" {
		if needWriteComma {
			w.RawByte(',')
		}
		w.RawByte('"')
		w.RawString("err_extra_info")
		w.RawByte('"')
		w.RawByte(':')
		w.String(m.ErrExtraInfo)
		needWriteComma = true
	}
	_ = needWriteComma
	w.RawByte('}')

}

func (m *ErrorInfo) MarshalJSON() ([]byte, error) {
	w := github_com_mailru_easyjson_jwriter.Writer{Buffer: github_com_mailru_easyjson_buffer.Buffer{Buf: make([]byte, 0, 2048)}}
	m.MarshalEasyJSON(&w)
	return w.BuildBytes()
}
func (m *ErrorInfo) String() string {
	d, _ := m.MarshalJSON()
	return *(*string)(unsafe.Pointer(&d))
}
func (m *ErrorInfo) GoString() string {
	return m.String()
}

func (m *ErrorMessage) MarshalJSON() ([]byte, error) {
	w := github_com_mailru_easyjson_jwriter.Writer{Buffer: github_com_mailru_easyjson_buffer.Buffer{Buf: make([]byte, 0, 2048)}}
	m.MarshalEasyJSON(&w)
	return w.BuildBytes()
}
func (m *ErrorMessage) String() string {
	d, _ := m.MarshalJSON()
	return *(*string)(unsafe.Pointer(&d))
}
func (m *ErrorMessage) GoString() string {
	return m.String()
}

func (m *ErrorInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TipName)
	if l > 0 {
		n += 1 + l + sovError(uint64(l))
	}
	l = len(m.TipDesc)
	if l > 0 {
		n += 1 + l + sovError(uint64(l))
	}
	if m.ErrorCode != 0 {
		n += 1 + sovError(uint64(m.ErrorCode))
	}
	return n
}

func (m *ErrorMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ErrCode != 0 {
		n += 1 + sovError(uint64(m.ErrCode))
	}
	l = len(m.ErrMsg)
	if l > 0 {
		n += 1 + l + sovError(uint64(l))
	}
	l = len(m.ErrInternalInfo)
	if l > 0 {
		n += 1 + l + sovError(uint64(l))
	}
	if len(m.StackStace) > 0 {
		l = 0
		for _, e := range m.StackStace {
			l += sovError(uint64(e))
		}
		n += 1 + sovError(uint64(l)) + l
	}
	l = len(m.ErrExtraInfo)
	if l > 0 {
		n += 1 + l + sovError(uint64(l))
	}
	return n
}

func sovError(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozError(x uint64) (n int) {
	return sovError(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ErrorInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowError
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ErrorInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ErrorInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TipName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowError
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthError
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthError
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TipName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TipDesc", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowError
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthError
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthError
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TipDesc = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrorCode", wireType)
			}
			m.ErrorCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowError
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ErrorCode |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipError(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthError
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ErrorMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowError
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ErrorMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ErrorMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrCode", wireType)
			}
			m.ErrCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowError
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ErrCode |= ErrorType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrMsg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowError
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthError
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthError
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ErrMsg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrInternalInfo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowError
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthError
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthError
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ErrInternalInfo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowError
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.StackStace = append(m.StackStace, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowError
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthError
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthError
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.StackStace) == 0 {
					m.StackStace = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowError
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.StackStace = append(m.StackStace, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field StackStace", wireType)
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrExtraInfo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowError
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthError
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthError
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ErrExtraInfo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipError(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthError
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipError(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowError
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowError
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowError
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthError
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupError
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthError
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthError        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowError          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupError = fmt.Errorf("proto: unexpected end of group")
)
