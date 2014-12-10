// Code generated by protoc-gen-gogo.
// source: unmarshalmerge.proto
// DO NOT EDIT!

/*
	Package unmarshalmerge is a generated protocol buffer package.

	It is generated from these files:
		unmarshalmerge.proto

	It has these top-level messages:
		Big
		BigUnsafe
		Sub
*/
package unmarshalmerge

import proto "github.com/gogo/protobuf/proto"
import math "math"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto/gogo.pb"

import io "io"
import fmt "fmt"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"

import io1 "io"
import fmt1 "fmt"
import github_com_gogo_protobuf_proto1 "github.com/gogo/protobuf/proto"

import fmt2 "fmt"
import strings "strings"
import reflect "reflect"

import fmt3 "fmt"
import strings1 "strings"
import github_com_gogo_protobuf_proto2 "github.com/gogo/protobuf/proto"
import sort "sort"
import strconv "strconv"
import reflect1 "reflect"

import fmt4 "fmt"
import bytes "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Big struct {
	Sub              *Sub   `protobuf:"bytes,1,opt" json:"Sub,omitempty"`
	Number           *int64 `protobuf:"varint,2,opt" json:"Number,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Big) Reset()      { *m = Big{} }
func (*Big) ProtoMessage() {}

func (m *Big) GetSub() *Sub {
	if m != nil {
		return m.Sub
	}
	return nil
}

func (m *Big) GetNumber() int64 {
	if m != nil && m.Number != nil {
		return *m.Number
	}
	return 0
}

type BigUnsafe struct {
	Sub              *Sub   `protobuf:"bytes,1,opt" json:"Sub,omitempty"`
	Number           *int64 `protobuf:"varint,2,opt" json:"Number,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *BigUnsafe) Reset()      { *m = BigUnsafe{} }
func (*BigUnsafe) ProtoMessage() {}

func (m *BigUnsafe) GetSub() *Sub {
	if m != nil {
		return m.Sub
	}
	return nil
}

func (m *BigUnsafe) GetNumber() int64 {
	if m != nil && m.Number != nil {
		return *m.Number
	}
	return 0
}

type Sub struct {
	SubNumber        *int64 `protobuf:"varint,1,opt" json:"SubNumber,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Sub) Reset()      { *m = Sub{} }
func (*Sub) ProtoMessage() {}

func (m *Sub) GetSubNumber() int64 {
	if m != nil && m.SubNumber != nil {
		return *m.SubNumber
	}
	return 0
}

func init() {
}
func (m *Big) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sub", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Sub == nil {
				m.Sub = &Sub{}
			}
			if err := m.Sub.Unmarshal(data[index:postIndex]); err != nil {
				return err
			}
			index = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Number", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Number = &v
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := github_com_gogo_protobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *Sub) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubNumber", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.SubNumber = &v
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := github_com_gogo_protobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *BigUnsafe) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io1.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt1.Errorf("proto: wrong wireType = %d for field Sub", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io1.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + msglen
			if postIndex > l {
				return io1.ErrUnexpectedEOF
			}
			if m.Sub == nil {
				m.Sub = &Sub{}
			}
			if err := m.Sub.Unmarshal(data[index:postIndex]); err != nil {
				return err
			}
			index = postIndex
		case 2:
			if wireType != 0 {
				return fmt1.Errorf("proto: wrong wireType = %d for field Number", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io1.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Number = &v
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := github_com_gogo_protobuf_proto1.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io1.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (this *Big) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Big{`,
		`Sub:` + strings.Replace(fmt2.Sprintf("%v", this.Sub), "Sub", "Sub", 1) + `,`,
		`Number:` + valueToStringUnmarshalmerge(this.Number) + `,`,
		`XXX_unrecognized:` + fmt2.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *BigUnsafe) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&BigUnsafe{`,
		`Sub:` + strings.Replace(fmt2.Sprintf("%v", this.Sub), "Sub", "Sub", 1) + `,`,
		`Number:` + valueToStringUnmarshalmerge(this.Number) + `,`,
		`XXX_unrecognized:` + fmt2.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Sub) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Sub{`,
		`SubNumber:` + valueToStringUnmarshalmerge(this.SubNumber) + `,`,
		`XXX_unrecognized:` + fmt2.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringUnmarshalmerge(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt2.Sprintf("*%v", pv)
}
func NewPopulatedBig(r randyUnmarshalmerge, easy bool) *Big {
	this := &Big{}
	if r.Intn(10) != 0 {
		this.Sub = NewPopulatedSub(r, easy)
	}
	if r.Intn(10) != 0 {
		v1 := r.Int63()
		if r.Intn(2) == 0 {
			v1 *= -1
		}
		this.Number = &v1
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedUnmarshalmerge(r, 3)
	}
	return this
}

func NewPopulatedBigUnsafe(r randyUnmarshalmerge, easy bool) *BigUnsafe {
	this := &BigUnsafe{}
	if r.Intn(10) != 0 {
		this.Sub = NewPopulatedSub(r, easy)
	}
	if r.Intn(10) != 0 {
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		this.Number = &v2
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedUnmarshalmerge(r, 3)
	}
	return this
}

func NewPopulatedSub(r randyUnmarshalmerge, easy bool) *Sub {
	this := &Sub{}
	if r.Intn(10) != 0 {
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		this.SubNumber = &v3
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedUnmarshalmerge(r, 2)
	}
	return this
}

type randyUnmarshalmerge interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneUnmarshalmerge(r randyUnmarshalmerge) rune {
	res := rune(r.Uint32() % 1112064)
	if 55296 <= res {
		res += 2047
	}
	return res
}
func randStringUnmarshalmerge(r randyUnmarshalmerge) string {
	v4 := r.Intn(100)
	tmps := make([]rune, v4)
	for i := 0; i < v4; i++ {
		tmps[i] = randUTF8RuneUnmarshalmerge(r)
	}
	return string(tmps)
}
func randUnrecognizedUnmarshalmerge(r randyUnmarshalmerge, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldUnmarshalmerge(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldUnmarshalmerge(data []byte, r randyUnmarshalmerge, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateUnmarshalmerge(data, uint64(key))
		v5 := r.Int63()
		if r.Intn(2) == 0 {
			v5 *= -1
		}
		data = encodeVarintPopulateUnmarshalmerge(data, uint64(v5))
	case 1:
		data = encodeVarintPopulateUnmarshalmerge(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateUnmarshalmerge(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateUnmarshalmerge(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateUnmarshalmerge(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateUnmarshalmerge(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
func (this *Big) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings1.Join([]string{`&unmarshalmerge.Big{` +
		`Sub:` + fmt3.Sprintf("%#v", this.Sub),
		`Number:` + valueToGoStringUnmarshalmerge(this.Number, "int64"),
		`XXX_unrecognized:` + fmt3.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func (this *BigUnsafe) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings1.Join([]string{`&unmarshalmerge.BigUnsafe{` +
		`Sub:` + fmt3.Sprintf("%#v", this.Sub),
		`Number:` + valueToGoStringUnmarshalmerge(this.Number, "int64"),
		`XXX_unrecognized:` + fmt3.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func (this *Sub) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings1.Join([]string{`&unmarshalmerge.Sub{` +
		`SubNumber:` + valueToGoStringUnmarshalmerge(this.SubNumber, "int64"),
		`XXX_unrecognized:` + fmt3.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func valueToGoStringUnmarshalmerge(v interface{}, typ string) string {
	rv := reflect1.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect1.Indirect(rv).Interface()
	return fmt3.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringUnmarshalmerge(e map[int32]github_com_gogo_protobuf_proto2.Extension) string {
	if e == nil {
		return "nil"
	}
	s := "map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings1.Join(ss, ",") + "}"
	return s
}
func (this *Big) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt4.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Big)
	if !ok {
		return fmt4.Errorf("that is not of type *Big")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt4.Errorf("that is type *Big but is nil && this != nil")
	} else if this == nil {
		return fmt4.Errorf("that is type *Bigbut is not nil && this == nil")
	}
	if !this.Sub.Equal(that1.Sub) {
		return fmt4.Errorf("Sub this(%v) Not Equal that(%v)", this.Sub, that1.Sub)
	}
	if this.Number != nil && that1.Number != nil {
		if *this.Number != *that1.Number {
			return fmt4.Errorf("Number this(%v) Not Equal that(%v)", *this.Number, *that1.Number)
		}
	} else if this.Number != nil {
		return fmt4.Errorf("this.Number == nil && that.Number != nil")
	} else if that1.Number != nil {
		return fmt4.Errorf("Number this(%v) Not Equal that(%v)", this.Number, that1.Number)
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt4.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *Big) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Big)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Sub.Equal(that1.Sub) {
		return false
	}
	if this.Number != nil && that1.Number != nil {
		if *this.Number != *that1.Number {
			return false
		}
	} else if this.Number != nil {
		return false
	} else if that1.Number != nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *BigUnsafe) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt4.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*BigUnsafe)
	if !ok {
		return fmt4.Errorf("that is not of type *BigUnsafe")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt4.Errorf("that is type *BigUnsafe but is nil && this != nil")
	} else if this == nil {
		return fmt4.Errorf("that is type *BigUnsafebut is not nil && this == nil")
	}
	if !this.Sub.Equal(that1.Sub) {
		return fmt4.Errorf("Sub this(%v) Not Equal that(%v)", this.Sub, that1.Sub)
	}
	if this.Number != nil && that1.Number != nil {
		if *this.Number != *that1.Number {
			return fmt4.Errorf("Number this(%v) Not Equal that(%v)", *this.Number, *that1.Number)
		}
	} else if this.Number != nil {
		return fmt4.Errorf("this.Number == nil && that.Number != nil")
	} else if that1.Number != nil {
		return fmt4.Errorf("Number this(%v) Not Equal that(%v)", this.Number, that1.Number)
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt4.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *BigUnsafe) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*BigUnsafe)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Sub.Equal(that1.Sub) {
		return false
	}
	if this.Number != nil && that1.Number != nil {
		if *this.Number != *that1.Number {
			return false
		}
	} else if this.Number != nil {
		return false
	} else if that1.Number != nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Sub) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt4.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Sub)
	if !ok {
		return fmt4.Errorf("that is not of type *Sub")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt4.Errorf("that is type *Sub but is nil && this != nil")
	} else if this == nil {
		return fmt4.Errorf("that is type *Subbut is not nil && this == nil")
	}
	if this.SubNumber != nil && that1.SubNumber != nil {
		if *this.SubNumber != *that1.SubNumber {
			return fmt4.Errorf("SubNumber this(%v) Not Equal that(%v)", *this.SubNumber, *that1.SubNumber)
		}
	} else if this.SubNumber != nil {
		return fmt4.Errorf("this.SubNumber == nil && that.SubNumber != nil")
	} else if that1.SubNumber != nil {
		return fmt4.Errorf("SubNumber this(%v) Not Equal that(%v)", this.SubNumber, that1.SubNumber)
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt4.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *Sub) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Sub)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.SubNumber != nil && that1.SubNumber != nil {
		if *this.SubNumber != *that1.SubNumber {
			return false
		}
	} else if this.SubNumber != nil {
		return false
	} else if that1.SubNumber != nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
