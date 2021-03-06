// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: examples/proxy/messages/proxy.proto

package messages

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import bytes "bytes"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ID struct {
	PublicKey []byte `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	Address   string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Id        []byte `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *ID) Reset()      { *m = ID{} }
func (*ID) ProtoMessage() {}
func (*ID) Descriptor() ([]byte, []int) {
	return fileDescriptor_proxy_fb373c5375e98a1e, []int{0}
}
func (m *ID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ID.Merge(dst, src)
}
func (m *ID) XXX_Size() int {
	return m.Size()
}
func (m *ID) XXX_DiscardUnknown() {
	xxx_messageInfo_ID.DiscardUnknown(m)
}

var xxx_messageInfo_ID proto.InternalMessageInfo

func (m *ID) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *ID) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ID) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

type ProxyMessage struct {
	Message     string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Destination *ID    `protobuf:"bytes,2,opt,name=destination" json:"destination,omitempty"`
}

func (m *ProxyMessage) Reset()      { *m = ProxyMessage{} }
func (*ProxyMessage) ProtoMessage() {}
func (*ProxyMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_proxy_fb373c5375e98a1e, []int{1}
}
func (m *ProxyMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProxyMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProxyMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ProxyMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyMessage.Merge(dst, src)
}
func (m *ProxyMessage) XXX_Size() int {
	return m.Size()
}
func (m *ProxyMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyMessage proto.InternalMessageInfo

func (m *ProxyMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ProxyMessage) GetDestination() *ID {
	if m != nil {
		return m.Destination
	}
	return nil
}

func init() {
	proto.RegisterType((*ID)(nil), "messages.ID")
	proto.RegisterType((*ProxyMessage)(nil), "messages.ProxyMessage")
}
func (this *ID) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*ID)
	if !ok {
		that2, ok := that.(ID)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *ID")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *ID but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *ID but is not nil && this == nil")
	}
	if !bytes.Equal(this.PublicKey, that1.PublicKey) {
		return fmt.Errorf("PublicKey this(%v) Not Equal that(%v)", this.PublicKey, that1.PublicKey)
	}
	if this.Address != that1.Address {
		return fmt.Errorf("Address this(%v) Not Equal that(%v)", this.Address, that1.Address)
	}
	if !bytes.Equal(this.Id, that1.Id) {
		return fmt.Errorf("Id this(%v) Not Equal that(%v)", this.Id, that1.Id)
	}
	return nil
}
func (this *ID) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ID)
	if !ok {
		that2, ok := that.(ID)
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
	if !bytes.Equal(this.PublicKey, that1.PublicKey) {
		return false
	}
	if this.Address != that1.Address {
		return false
	}
	if !bytes.Equal(this.Id, that1.Id) {
		return false
	}
	return true
}
func (this *ProxyMessage) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*ProxyMessage)
	if !ok {
		that2, ok := that.(ProxyMessage)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *ProxyMessage")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *ProxyMessage but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *ProxyMessage but is not nil && this == nil")
	}
	if this.Message != that1.Message {
		return fmt.Errorf("Message this(%v) Not Equal that(%v)", this.Message, that1.Message)
	}
	if !this.Destination.Equal(that1.Destination) {
		return fmt.Errorf("Destination this(%v) Not Equal that(%v)", this.Destination, that1.Destination)
	}
	return nil
}
func (this *ProxyMessage) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ProxyMessage)
	if !ok {
		that2, ok := that.(ProxyMessage)
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
	if this.Message != that1.Message {
		return false
	}
	if !this.Destination.Equal(that1.Destination) {
		return false
	}
	return true
}
func (this *ID) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&messages.ID{")
	s = append(s, "PublicKey: "+fmt.Sprintf("%#v", this.PublicKey)+",\n")
	s = append(s, "Address: "+fmt.Sprintf("%#v", this.Address)+",\n")
	s = append(s, "Id: "+fmt.Sprintf("%#v", this.Id)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ProxyMessage) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&messages.ProxyMessage{")
	s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",\n")
	if this.Destination != nil {
		s = append(s, "Destination: "+fmt.Sprintf("%#v", this.Destination)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringProxy(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *ID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ID) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.PublicKey) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintProxy(dAtA, i, uint64(len(m.PublicKey)))
		i += copy(dAtA[i:], m.PublicKey)
	}
	if len(m.Address) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintProxy(dAtA, i, uint64(len(m.Address)))
		i += copy(dAtA[i:], m.Address)
	}
	if len(m.Id) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintProxy(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	return i, nil
}

func (m *ProxyMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProxyMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintProxy(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	if m.Destination != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintProxy(dAtA, i, uint64(m.Destination.Size()))
		n1, err := m.Destination.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func encodeVarintProxy(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedID(r randyProxy, easy bool) *ID {
	this := &ID{}
	v1 := r.Intn(100)
	this.PublicKey = make([]byte, v1)
	for i := 0; i < v1; i++ {
		this.PublicKey[i] = byte(r.Intn(256))
	}
	this.Address = string(randStringProxy(r))
	v2 := r.Intn(100)
	this.Id = make([]byte, v2)
	for i := 0; i < v2; i++ {
		this.Id[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedProxyMessage(r randyProxy, easy bool) *ProxyMessage {
	this := &ProxyMessage{}
	this.Message = string(randStringProxy(r))
	if r.Intn(10) != 0 {
		this.Destination = NewPopulatedID(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyProxy interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneProxy(r randyProxy) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringProxy(r randyProxy) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RuneProxy(r)
	}
	return string(tmps)
}
func randUnrecognizedProxy(r randyProxy, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldProxy(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldProxy(dAtA []byte, r randyProxy, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateProxy(dAtA, uint64(key))
		v4 := r.Int63()
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		dAtA = encodeVarintPopulateProxy(dAtA, uint64(v4))
	case 1:
		dAtA = encodeVarintPopulateProxy(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateProxy(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateProxy(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateProxy(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateProxy(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *ID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PublicKey)
	if l > 0 {
		n += 1 + l + sovProxy(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovProxy(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovProxy(uint64(l))
	}
	return n
}

func (m *ProxyMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovProxy(uint64(l))
	}
	if m.Destination != nil {
		l = m.Destination.Size()
		n += 1 + l + sovProxy(uint64(l))
	}
	return n
}

func sovProxy(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozProxy(x uint64) (n int) {
	return sovProxy(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ID) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ID{`,
		`PublicKey:` + fmt.Sprintf("%v", this.PublicKey) + `,`,
		`Address:` + fmt.Sprintf("%v", this.Address) + `,`,
		`Id:` + fmt.Sprintf("%v", this.Id) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ProxyMessage) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ProxyMessage{`,
		`Message:` + fmt.Sprintf("%v", this.Message) + `,`,
		`Destination:` + strings.Replace(fmt.Sprintf("%v", this.Destination), "ID", "ID", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringProxy(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProxy
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProxy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthProxy
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PublicKey = append(m.PublicKey[:0], dAtA[iNdEx:postIndex]...)
			if m.PublicKey == nil {
				m.PublicKey = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProxy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProxy
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProxy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthProxy
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = append(m.Id[:0], dAtA[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProxy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProxy
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
func (m *ProxyMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProxy
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProxyMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProxyMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProxy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProxy
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Destination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProxy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProxy
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Destination == nil {
				m.Destination = &ID{}
			}
			if err := m.Destination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProxy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProxy
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
func skipProxy(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProxy
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
					return 0, ErrIntOverflowProxy
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowProxy
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthProxy
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowProxy
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipProxy(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthProxy = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProxy   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("examples/proxy/messages/proxy.proto", fileDescriptor_proxy_fb373c5375e98a1e)
}

var fileDescriptor_proxy_fb373c5375e98a1e = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0x2d, 0xd6, 0x2f, 0x28, 0xca, 0xaf, 0xa8, 0xd4, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e,
	0x4c, 0x87, 0x71, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x38, 0x60, 0xa2, 0x52, 0x4a, 0xe9,
	0xf9, 0xe9, 0xf9, 0xfa, 0x60, 0xd1, 0xa4, 0xd2, 0x34, 0x7d, 0x10, 0x0f, 0xcc, 0x01, 0xb3, 0x20,
	0xaa, 0x95, 0x7c, 0xb9, 0x98, 0x3c, 0x5d, 0x84, 0x64, 0xb9, 0xb8, 0x0a, 0x4a, 0x93, 0x72, 0x32,
	0x93, 0xe3, 0xb3, 0x53, 0x2b, 0x25, 0x18, 0x15, 0x18, 0x35, 0x78, 0x82, 0x38, 0x21, 0x22, 0xde,
	0xa9, 0x95, 0x42, 0x12, 0x5c, 0xec, 0x89, 0x29, 0x29, 0x45, 0xa9, 0xc5, 0xc5, 0x12, 0x4c, 0x0a,
	0x8c, 0x1a, 0x9c, 0x41, 0x30, 0xae, 0x10, 0x1f, 0x17, 0x53, 0x66, 0x8a, 0x04, 0x33, 0x58, 0x03,
	0x53, 0x66, 0x8a, 0x52, 0x04, 0x17, 0x4f, 0x00, 0xc8, 0x2d, 0xbe, 0x10, 0x37, 0x80, 0x74, 0x42,
	0x9d, 0x03, 0x36, 0x95, 0x33, 0x08, 0xc6, 0x15, 0xd2, 0xe3, 0xe2, 0x4e, 0x49, 0x2d, 0x2e, 0xc9,
	0xcc, 0x4b, 0x2c, 0xc9, 0xcc, 0xcf, 0x03, 0x9b, 0xcb, 0x6d, 0xc4, 0xa3, 0x07, 0x73, 0xbc, 0x9e,
	0xa7, 0x4b, 0x10, 0xb2, 0x02, 0x27, 0x93, 0x1b, 0x0f, 0xe5, 0x18, 0x1e, 0x3c, 0x94, 0x63, 0xfc,
	0xf0, 0x50, 0x8e, 0xf1, 0xc7, 0x43, 0x39, 0xc6, 0x86, 0x47, 0x72, 0x8c, 0x2b, 0x1e, 0xc9, 0x31,
	0xee, 0x78, 0x24, 0xc7, 0x78, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9,
	0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x49, 0x6c,
	0x60, 0x5f, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x54, 0xe6, 0xc7, 0xa3, 0x3a, 0x01, 0x00,
	0x00,
}
