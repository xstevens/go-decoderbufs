// Code generated by protoc-gen-gogo.
// source: pg_logicaldec.proto
// DO NOT EDIT!

/*
Package decoderbufs is a generated protocol buffer package.

It is generated from these files:
	pg_logicaldec.proto

It has these top-level messages:
	Point
	DatumMessage
	RowMessage
*/
package decoderbufs

import proto "github.com/gogo/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Op int32

const (
	Op_INSERT Op = 0
	Op_UPDATE Op = 1
	Op_DELETE Op = 2
)

var Op_name = map[int32]string{
	0: "INSERT",
	1: "UPDATE",
	2: "DELETE",
}
var Op_value = map[string]int32{
	"INSERT": 0,
	"UPDATE": 1,
	"DELETE": 2,
}

func (x Op) Enum() *Op {
	p := new(Op)
	*p = x
	return p
}
func (x Op) String() string {
	return proto.EnumName(Op_name, int32(x))
}
func (x *Op) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Op_value, data, "Op")
	if err != nil {
		return err
	}
	*x = Op(value)
	return nil
}

type Point struct {
	X                *float64 `protobuf:"fixed64,1,req,name=x" json:"x,omitempty"`
	Y                *float64 `protobuf:"fixed64,2,req,name=y" json:"y,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Point) Reset()         { *m = Point{} }
func (m *Point) String() string { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()    {}

func (m *Point) GetX() float64 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *Point) GetY() float64 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

type DatumMessage struct {
	ColumnName       *string  `protobuf:"bytes,1,opt,name=column_name" json:"column_name,omitempty"`
	ColumnType       *int64   `protobuf:"varint,2,opt,name=column_type" json:"column_type,omitempty"`
	DatumInt32       *int32   `protobuf:"varint,3,opt,name=datum_int32" json:"datum_int32,omitempty"`
	DatumInt64       *int64   `protobuf:"varint,4,opt,name=datum_int64" json:"datum_int64,omitempty"`
	DatumFloat       *float32 `protobuf:"fixed32,5,opt,name=datum_float" json:"datum_float,omitempty"`
	DatumDouble      *float64 `protobuf:"fixed64,6,opt,name=datum_double" json:"datum_double,omitempty"`
	DatumBool        *bool    `protobuf:"varint,7,opt,name=datum_bool" json:"datum_bool,omitempty"`
	DatumString      *string  `protobuf:"bytes,8,opt,name=datum_string" json:"datum_string,omitempty"`
	DatumBytes       []byte   `protobuf:"bytes,9,opt,name=datum_bytes" json:"datum_bytes,omitempty"`
	DatumPoint       *Point   `protobuf:"bytes,10,opt,name=datum_point" json:"datum_point,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *DatumMessage) Reset()         { *m = DatumMessage{} }
func (m *DatumMessage) String() string { return proto.CompactTextString(m) }
func (*DatumMessage) ProtoMessage()    {}

func (m *DatumMessage) GetColumnName() string {
	if m != nil && m.ColumnName != nil {
		return *m.ColumnName
	}
	return ""
}

func (m *DatumMessage) GetColumnType() int64 {
	if m != nil && m.ColumnType != nil {
		return *m.ColumnType
	}
	return 0
}

func (m *DatumMessage) GetDatumInt32() int32 {
	if m != nil && m.DatumInt32 != nil {
		return *m.DatumInt32
	}
	return 0
}

func (m *DatumMessage) GetDatumInt64() int64 {
	if m != nil && m.DatumInt64 != nil {
		return *m.DatumInt64
	}
	return 0
}

func (m *DatumMessage) GetDatumFloat() float32 {
	if m != nil && m.DatumFloat != nil {
		return *m.DatumFloat
	}
	return 0
}

func (m *DatumMessage) GetDatumDouble() float64 {
	if m != nil && m.DatumDouble != nil {
		return *m.DatumDouble
	}
	return 0
}

func (m *DatumMessage) GetDatumBool() bool {
	if m != nil && m.DatumBool != nil {
		return *m.DatumBool
	}
	return false
}

func (m *DatumMessage) GetDatumString() string {
	if m != nil && m.DatumString != nil {
		return *m.DatumString
	}
	return ""
}

func (m *DatumMessage) GetDatumBytes() []byte {
	if m != nil {
		return m.DatumBytes
	}
	return nil
}

func (m *DatumMessage) GetDatumPoint() *Point {
	if m != nil {
		return m.DatumPoint
	}
	return nil
}

type RowMessage struct {
	TransactionId    *uint32         `protobuf:"varint,1,opt,name=transaction_id" json:"transaction_id,omitempty"`
	CommitTime       *uint64         `protobuf:"varint,2,opt,name=commit_time" json:"commit_time,omitempty"`
	Table            *string         `protobuf:"bytes,3,opt,name=table" json:"table,omitempty"`
	Op               *Op             `protobuf:"varint,4,opt,name=op,enum=decoderbufs.Op" json:"op,omitempty"`
	NewTuple         []*DatumMessage `protobuf:"bytes,5,rep,name=new_tuple" json:"new_tuple,omitempty"`
	OldTuple         []*DatumMessage `protobuf:"bytes,6,rep,name=old_tuple" json:"old_tuple,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *RowMessage) Reset()         { *m = RowMessage{} }
func (m *RowMessage) String() string { return proto.CompactTextString(m) }
func (*RowMessage) ProtoMessage()    {}

func (m *RowMessage) GetTransactionId() uint32 {
	if m != nil && m.TransactionId != nil {
		return *m.TransactionId
	}
	return 0
}

func (m *RowMessage) GetCommitTime() uint64 {
	if m != nil && m.CommitTime != nil {
		return *m.CommitTime
	}
	return 0
}

func (m *RowMessage) GetTable() string {
	if m != nil && m.Table != nil {
		return *m.Table
	}
	return ""
}

func (m *RowMessage) GetOp() Op {
	if m != nil && m.Op != nil {
		return *m.Op
	}
	return Op_INSERT
}

func (m *RowMessage) GetNewTuple() []*DatumMessage {
	if m != nil {
		return m.NewTuple
	}
	return nil
}

func (m *RowMessage) GetOldTuple() []*DatumMessage {
	if m != nil {
		return m.OldTuple
	}
	return nil
}

func init() {
	proto.RegisterEnum("decoderbufs.Op", Op_name, Op_value)
}
