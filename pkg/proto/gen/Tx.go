// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package gen

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Tx struct {
	_tab flatbuffers.Table
}

func GetRootAsTx(buf []byte, offset flatbuffers.UOffsetT) *Tx {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Tx{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsTx(buf []byte, offset flatbuffers.UOffsetT) *Tx {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Tx{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Tx) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Tx) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Tx) Hash(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *Tx) HashLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Tx) HashBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Tx) MutateHash(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *Tx) Kind() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Tx) MutateKind(n byte) bool {
	return rcv._tab.MutateByteSlot(6, n)
}

func (rcv *Tx) Ts() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Tx) MutateTs(n int64) bool {
	return rcv._tab.MutateInt64Slot(8, n)
}

func (rcv *Tx) Payload(obj *Entry) *Entry {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Entry)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func TxStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func TxAddHash(builder *flatbuffers.Builder, hash flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(hash), 0)
}
func TxStartHashVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func TxAddKind(builder *flatbuffers.Builder, kind byte) {
	builder.PrependByteSlot(1, kind, 0)
}
func TxAddTs(builder *flatbuffers.Builder, ts int64) {
	builder.PrependInt64Slot(2, ts, 0)
}
func TxAddPayload(builder *flatbuffers.Builder, payload flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(payload), 0)
}
func TxEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
