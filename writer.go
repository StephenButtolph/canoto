package canoto

import "encoding/binary"

type Writer struct {
	b []byte
}

func (w *Writer) Tag(fieldNumber uint32, wireType WireType) {
	w.Uint64(tagToUint64(fieldNumber, wireType))
}

func (w *Writer) Int32(v int32) {
	w.Uint64(uint64(v))
}

func (w *Writer) Int64(v int64) {
	w.Uint64(uint64(v))
}

func (w *Writer) Uint32(v uint32) {
	w.Uint64(uint64(v))
}

func (w *Writer) Uint64(v uint64) {
	w.b = binary.AppendUvarint(w.b, v)
}

func (w *Writer) Sint32(v int32) {
	w.Uint64(sint32ToUint64(v))
}

func (w *Writer) Sint64(v int64) {
	w.Uint64(sint64ToUint64(v))
}

func (w *Writer) Fixed32(v uint32) {
	var bytes [SizeFixed32]byte
	binary.LittleEndian.PutUint32(bytes[:], v)
	w.b = append(w.b, bytes[:]...)
}

func (w *Writer) Fixed64(v uint64) {
	var bytes [SizeFixed64]byte
	binary.LittleEndian.PutUint64(bytes[:], v)
	w.b = append(w.b, bytes[:]...)
}

func (w *Writer) Sfixed32(v int32) {
	w.Fixed32(uint32(v))
}

func (w *Writer) Sfixed64(v int64) {
	w.Fixed64(uint64(v))
}

func (w *Writer) Bool(v bool) {
	if v {
		w.b = append(w.b, trueByte)
	} else {
		w.b = append(w.b, falseByte)
	}
}

func (w *Writer) String(v string) {
	w.Int32(int32(len(v)))
	w.b = append(w.b, v...)
}

func (w *Writer) Bytes(v []byte) {
	w.Int32(int32(len(v)))
	w.b = append(w.b, v...)
}

func tagToUint64(fieldNumber uint32, wireType WireType) uint64 {
	return uint64(fieldNumber<<wireTypeLength | uint32(wireType))
}

func sint32ToUint64(v int32) uint64 {
	if v >= 0 {
		return uint64(v) << 1
	} else {
		return ^uint64(v)<<1 | 1
	}
}

func sint64ToUint64(v int64) uint64 {
	if v >= 0 {
		return uint64(v) << 1
	} else {
		return ^uint64(v)<<1 | 1
	}
}
