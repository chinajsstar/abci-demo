// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

/*
Package abci-demo is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	User
	Request
	RequestInitPlatform
	RequestUserInfoOnChain
	Response
	ResponseInitPlatform
	ResponseUserInfoOnChain
	Receipt
*/
package abci-demo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 消息类型
type MessageType int32

const (
	MessageType_MsgError           MessageType = 0
	MessageType_MsgInitPlatform    MessageType = 1
	MessageType_MsgUserInfoOnChain MessageType = 2
)

var MessageType_name = map[int32]string{
	0: "MsgError",
	1: "MsgInitPlatform",
	2: "MsgUserInfoOnChain",
}
var MessageType_value = map[string]int32{
	"MsgError":           0,
	"MsgInitPlatform":    1,
	"MsgUserInfoOnChain": 2,
}

func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}
func (MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// 用户信息
type User struct {
	Username    string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Age         int64  `protobuf:"varint,2,opt,name=age" json:"age,omitempty"`
	Educate     string `protobuf:"bytes,3,opt,name=educate" json:"educate,omitempty"`
	Workstation string `protobuf:"bytes,4,opt,name=workstation" json:"workstation,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *User) GetEducate() string {
	if m != nil {
		return m.Educate
	}
	return ""
}

func (m *User) GetWorkstation() string {
	if m != nil {
		return m.Workstation
	}
	return ""
}

type Request struct {
	// Types that are valid to be assigned to Value:
	//	*Request_InitPlatform
	//	*Request_UserInfoOnChain
	Value         isRequest_Value `protobuf_oneof:"value"`
	InstructionId int64           `protobuf:"varint,27,opt,name=instructionId" json:"instructionId,omitempty"`
	Pubkey        []byte          `protobuf:"bytes,28,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Sign          []byte          `protobuf:"bytes,29,opt,name=sign,proto3" json:"sign,omitempty"`
	ActionId      MessageType     `protobuf:"varint,30,opt,name=actionId,enum=abci-demo.MessageType" json:"actionId,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isRequest_Value interface {
	isRequest_Value()
}

type Request_InitPlatform struct {
	InitPlatform *RequestInitPlatform `protobuf:"bytes,1,opt,name=initPlatform,oneof"`
}
type Request_UserInfoOnChain struct {
	UserInfoOnChain *RequestUserInfoOnChain `protobuf:"bytes,2,opt,name=userInfoOnChain,oneof"`
}

func (*Request_InitPlatform) isRequest_Value()    {}
func (*Request_UserInfoOnChain) isRequest_Value() {}

func (m *Request) GetValue() isRequest_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Request) GetInitPlatform() *RequestInitPlatform {
	if x, ok := m.GetValue().(*Request_InitPlatform); ok {
		return x.InitPlatform
	}
	return nil
}

func (m *Request) GetUserInfoOnChain() *RequestUserInfoOnChain {
	if x, ok := m.GetValue().(*Request_UserInfoOnChain); ok {
		return x.UserInfoOnChain
	}
	return nil
}

func (m *Request) GetInstructionId() int64 {
	if m != nil {
		return m.InstructionId
	}
	return 0
}

func (m *Request) GetPubkey() []byte {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

func (m *Request) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

func (m *Request) GetActionId() MessageType {
	if m != nil {
		return m.ActionId
	}
	return MessageType_MsgError
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Request) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Request_OneofMarshaler, _Request_OneofUnmarshaler, _Request_OneofSizer, []interface{}{
		(*Request_InitPlatform)(nil),
		(*Request_UserInfoOnChain)(nil),
	}
}

func _Request_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Request)
	// value
	switch x := m.Value.(type) {
	case *Request_InitPlatform:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.InitPlatform); err != nil {
			return err
		}
	case *Request_UserInfoOnChain:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UserInfoOnChain); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Request.Value has unexpected type %T", x)
	}
	return nil
}

func _Request_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Request)
	switch tag {
	case 1: // value.initPlatform
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RequestInitPlatform)
		err := b.DecodeMessage(msg)
		m.Value = &Request_InitPlatform{msg}
		return true, err
	case 2: // value.userInfoOnChain
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RequestUserInfoOnChain)
		err := b.DecodeMessage(msg)
		m.Value = &Request_UserInfoOnChain{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Request_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Request)
	// value
	switch x := m.Value.(type) {
	case *Request_InitPlatform:
		s := proto.Size(x.InitPlatform)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Request_UserInfoOnChain:
		s := proto.Size(x.UserInfoOnChain)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// 初始化平台
type RequestInitPlatform struct {
	UserName      string `protobuf:"bytes,1,opt,name=userName" json:"userName,omitempty"`
	Password      string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	UserPublicKey []byte `protobuf:"bytes,2,opt,name=userPublicKey,proto3" json:"userPublicKey,omitempty"`
}

func (m *RequestInitPlatform) Reset()                    { *m = RequestInitPlatform{} }
func (m *RequestInitPlatform) String() string            { return proto.CompactTextString(m) }
func (*RequestInitPlatform) ProtoMessage()               {}
func (*RequestInitPlatform) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RequestInitPlatform) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *RequestInitPlatform) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *RequestInitPlatform) GetUserPublicKey() []byte {
	if m != nil {
		return m.UserPublicKey
	}
	return nil
}

type RequestUserInfoOnChain struct {
	Username    string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Age         int64  `protobuf:"varint,2,opt,name=age" json:"age,omitempty"`
	Educate     string `protobuf:"bytes,3,opt,name=educate" json:"educate,omitempty"`
	Workstation string `protobuf:"bytes,4,opt,name=workstation" json:"workstation,omitempty"`
}

func (m *RequestUserInfoOnChain) Reset()                    { *m = RequestUserInfoOnChain{} }
func (m *RequestUserInfoOnChain) String() string            { return proto.CompactTextString(m) }
func (*RequestUserInfoOnChain) ProtoMessage()               {}
func (*RequestUserInfoOnChain) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RequestUserInfoOnChain) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *RequestUserInfoOnChain) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *RequestUserInfoOnChain) GetEducate() string {
	if m != nil {
		return m.Educate
	}
	return ""
}

func (m *RequestUserInfoOnChain) GetWorkstation() string {
	if m != nil {
		return m.Workstation
	}
	return ""
}

type Response struct {
	// Types that are valid to be assigned to Value:
	//	*Response_InitPlatform
	//	*Response_UserInfoOnChain
	Value isResponse_Value `protobuf_oneof:"value"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type isResponse_Value interface {
	isResponse_Value()
}

type Response_InitPlatform struct {
	InitPlatform *ResponseInitPlatform `protobuf:"bytes,1,opt,name=initPlatform,oneof"`
}
type Response_UserInfoOnChain struct {
	UserInfoOnChain *ResponseUserInfoOnChain `protobuf:"bytes,2,opt,name=userInfoOnChain,oneof"`
}

func (*Response_InitPlatform) isResponse_Value()    {}
func (*Response_UserInfoOnChain) isResponse_Value() {}

func (m *Response) GetValue() isResponse_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Response) GetInitPlatform() *ResponseInitPlatform {
	if x, ok := m.GetValue().(*Response_InitPlatform); ok {
		return x.InitPlatform
	}
	return nil
}

func (m *Response) GetUserInfoOnChain() *ResponseUserInfoOnChain {
	if x, ok := m.GetValue().(*Response_UserInfoOnChain); ok {
		return x.UserInfoOnChain
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Response) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Response_OneofMarshaler, _Response_OneofUnmarshaler, _Response_OneofSizer, []interface{}{
		(*Response_InitPlatform)(nil),
		(*Response_UserInfoOnChain)(nil),
	}
}

func _Response_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Response)
	// value
	switch x := m.Value.(type) {
	case *Response_InitPlatform:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.InitPlatform); err != nil {
			return err
		}
	case *Response_UserInfoOnChain:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UserInfoOnChain); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Response.Value has unexpected type %T", x)
	}
	return nil
}

func _Response_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Response)
	switch tag {
	case 1: // value.initPlatform
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ResponseInitPlatform)
		err := b.DecodeMessage(msg)
		m.Value = &Response_InitPlatform{msg}
		return true, err
	case 2: // value.userInfoOnChain
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ResponseUserInfoOnChain)
		err := b.DecodeMessage(msg)
		m.Value = &Response_UserInfoOnChain{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Response_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Response)
	// value
	switch x := m.Value.(type) {
	case *Response_InitPlatform:
		s := proto.Size(x.InitPlatform)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Response_UserInfoOnChain:
		s := proto.Size(x.UserInfoOnChain)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ResponseInitPlatform struct {
	InstructionId int64 `protobuf:"varint,1,opt,name=instructionId" json:"instructionId,omitempty"`
}

func (m *ResponseInitPlatform) Reset()                    { *m = ResponseInitPlatform{} }
func (m *ResponseInitPlatform) String() string            { return proto.CompactTextString(m) }
func (*ResponseInitPlatform) ProtoMessage()               {}
func (*ResponseInitPlatform) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ResponseInitPlatform) GetInstructionId() int64 {
	if m != nil {
		return m.InstructionId
	}
	return 0
}

type ResponseUserInfoOnChain struct {
	InstructionId int64 `protobuf:"varint,1,opt,name=instructionId" json:"instructionId,omitempty"`
}

func (m *ResponseUserInfoOnChain) Reset()                    { *m = ResponseUserInfoOnChain{} }
func (m *ResponseUserInfoOnChain) String() string            { return proto.CompactTextString(m) }
func (*ResponseUserInfoOnChain) ProtoMessage()               {}
func (*ResponseUserInfoOnChain) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ResponseUserInfoOnChain) GetInstructionId() int64 {
	if m != nil {
		return m.InstructionId
	}
	return 0
}

type Receipt struct {
	IsOk bool   `protobuf:"varint,1,opt,name=isOk" json:"isOk,omitempty"`
	Err  []byte `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (m *Receipt) Reset()                    { *m = Receipt{} }
func (m *Receipt) String() string            { return proto.CompactTextString(m) }
func (*Receipt) ProtoMessage()               {}
func (*Receipt) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Receipt) GetIsOk() bool {
	if m != nil {
		return m.IsOk
	}
	return false
}

func (m *Receipt) GetErr() []byte {
	if m != nil {
		return m.Err
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "abci-demo.User")
	proto.RegisterType((*Request)(nil), "abci-demo.Request")
	proto.RegisterType((*RequestInitPlatform)(nil), "abci-demo.RequestInitPlatform")
	proto.RegisterType((*RequestUserInfoOnChain)(nil), "abci-demo.RequestUserInfoOnChain")
	proto.RegisterType((*Response)(nil), "abci-demo.Response")
	proto.RegisterType((*ResponseInitPlatform)(nil), "abci-demo.ResponseInitPlatform")
	proto.RegisterType((*ResponseUserInfoOnChain)(nil), "abci-demo.ResponseUserInfoOnChain")
	proto.RegisterType((*Receipt)(nil), "abci-demo.Receipt")
	proto.RegisterEnum("abci-demo.MessageType", MessageType_name, MessageType_value)
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 467 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x41, 0x6f, 0x13, 0x3d,
	0x10, 0x8d, 0x93, 0x7c, 0xcd, 0x76, 0x92, 0x8f, 0x46, 0x2e, 0x84, 0x15, 0xd0, 0x6a, 0x59, 0x71,
	0x88, 0x38, 0x04, 0x11, 0xae, 0x48, 0x48, 0xb4, 0x48, 0x89, 0x60, 0x69, 0x65, 0xd1, 0x1f, 0xe0,
	0x6c, 0xa6, 0x8b, 0x95, 0xc4, 0x5e, 0x6c, 0x2f, 0x55, 0xae, 0xfc, 0x1a, 0xc4, 0xaf, 0x44, 0x36,
	0x49, 0xba, 0x49, 0xb6, 0xc0, 0x89, 0xdb, 0xcc, 0xf8, 0x79, 0xfc, 0x66, 0xde, 0xdb, 0x85, 0x43,
	0x9e, 0x8b, 0x41, 0xae, 0x95, 0x55, 0x34, 0xe0, 0x93, 0x54, 0x9c, 0xe3, 0x42, 0xc5, 0x39, 0x34,
	0xaf, 0x0c, 0x6a, 0xfa, 0x08, 0x82, 0xc2, 0xa0, 0x96, 0x7c, 0x81, 0x21, 0x89, 0x48, 0xff, 0x90,
	0x6d, 0x72, 0xda, 0x85, 0x06, 0xcf, 0x30, 0xac, 0x47, 0xa4, 0xdf, 0x60, 0x2e, 0xa4, 0x21, 0xb4,
	0x70, 0x5a, 0xa4, 0xdc, 0x62, 0xd8, 0xf0, 0xe0, 0x75, 0x4a, 0x23, 0x68, 0xdf, 0x28, 0x3d, 0x33,
	0x96, 0x5b, 0xa1, 0x64, 0xd8, 0xf4, 0xa7, 0xe5, 0x52, 0xfc, 0xa3, 0x0e, 0x2d, 0x86, 0x5f, 0x0a,
	0x34, 0x96, 0x9e, 0x41, 0x47, 0x48, 0x61, 0x2f, 0xe7, 0xdc, 0x5e, 0x2b, 0xbd, 0xf0, 0x2f, 0xb7,
	0x87, 0x27, 0x83, 0x35, 0xbd, 0xc1, 0x0a, 0x38, 0x2e, 0x81, 0x46, 0x35, 0xb6, 0x75, 0x89, 0x7e,
	0x80, 0x23, 0x47, 0x75, 0x2c, 0xaf, 0xd5, 0x85, 0x3c, 0xfb, 0xcc, 0x85, 0xf4, 0x54, 0xdb, 0xc3,
	0x68, 0xaf, 0xcf, 0xd5, 0x36, 0x6e, 0x54, 0x63, 0xbb, 0x57, 0xe9, 0x33, 0xf8, 0x5f, 0x48, 0x63,
	0x75, 0x91, 0x3a, 0xb6, 0xe3, 0x69, 0xf8, 0xd8, 0x8f, 0xbd, 0x5d, 0xa4, 0x3d, 0x38, 0xc8, 0x8b,
	0xc9, 0x0c, 0x97, 0xe1, 0x93, 0x88, 0xf4, 0x3b, 0x6c, 0x95, 0x51, 0x0a, 0x4d, 0x23, 0x32, 0x19,
	0x9e, 0xf8, 0xaa, 0x8f, 0xe9, 0x4b, 0x08, 0xf8, 0xba, 0xd9, 0x69, 0x44, 0xfa, 0xf7, 0x86, 0x0f,
	0x6e, 0x89, 0x25, 0x68, 0x0c, 0xcf, 0xf0, 0xd3, 0x32, 0x47, 0xb6, 0x81, 0xbd, 0x6d, 0xc1, 0x7f,
	0x5f, 0xf9, 0xbc, 0xc0, 0xd8, 0xc0, 0x71, 0xc5, 0x0a, 0xd6, 0x6a, 0x7d, 0xdc, 0x51, 0xcb, 0xe5,
	0xee, 0x2c, 0xe7, 0xc6, 0xdc, 0x28, 0x3d, 0x5d, 0x89, 0xb3, 0xc9, 0xdd, 0x70, 0x0e, 0x77, 0x59,
	0x4c, 0xe6, 0x22, 0x7d, 0x8f, 0x4b, 0xbf, 0xa8, 0x0e, 0xdb, 0x2e, 0xc6, 0xdf, 0x08, 0xf4, 0xaa,
	0x17, 0xf6, 0x0f, 0x6d, 0xf2, 0x9d, 0x40, 0xc0, 0xd0, 0xe4, 0x4a, 0x1a, 0xa4, 0xe7, 0x95, 0x3e,
	0x39, 0x2d, 0xeb, 0xfb, 0x0b, 0xf9, 0x5b, 0xa3, 0x24, 0x77, 0x19, 0xe5, 0xe9, 0x7e, 0xa3, 0x3f,
	0x3b, 0xe5, 0x56, 0xa4, 0xd7, 0x70, 0xbf, 0xea, 0xfd, 0x7d, 0x2b, 0x91, 0x0a, 0x2b, 0xc5, 0x6f,
	0xe0, 0xe1, 0x1d, 0x8f, 0xfe, 0x65, 0x83, 0x17, 0xee, 0x7b, 0x4a, 0x51, 0xe4, 0xd6, 0xd9, 0x4f,
	0x98, 0x8b, 0x99, 0xc7, 0x05, 0xcc, 0xc7, 0x4e, 0x16, 0xd4, 0x7a, 0xa5, 0xb4, 0x0b, 0x9f, 0x8f,
	0xa0, 0x5d, 0xb2, 0x1d, 0xed, 0x40, 0x90, 0x98, 0xec, 0x9d, 0xd6, 0x4a, 0x77, 0x6b, 0xf4, 0x18,
	0x8e, 0x12, 0x93, 0x95, 0xe7, 0xe8, 0x12, 0xda, 0x03, 0x9a, 0x98, 0x6c, 0x87, 0x5e, 0xb7, 0x3e,
	0x39, 0xf0, 0xbf, 0x93, 0x57, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x00, 0xfb, 0xfe, 0x5b,
	0x04, 0x00, 0x00,
}
