package unsafeslice

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnsafeSliceUint64(t *testing.T) {
	w := &bytes.Buffer{}
	d := []uint64{0xdead, 0xbeef, 0xb334}
	binary.Write(w, binary.LittleEndian, d)
	v := Uint64SliceFromByteSlice(w.Bytes())
	require.Equal(t, d, v)
	require.Equal(t, w.Bytes(), ByteSliceFromUint64Slice(d))
}

func TestUnsafeSliceInt64(t *testing.T) {
	w := &bytes.Buffer{}
	d := []int64{0xdead, 0xbeef, 0xb334}
	binary.Write(w, binary.LittleEndian, d)
	v := Int64SliceFromByteSlice(w.Bytes())
	require.Equal(t, d, v)
	require.Equal(t, w.Bytes(), ByteSliceFromInt64Slice(d))
}

func TestUnsafeSliceUint32(t *testing.T) {
	w := &bytes.Buffer{}
	d := []uint32{0xdead, 0xbeef, 0xb334}
	binary.Write(w, binary.LittleEndian, d)
	v := Uint32SliceFromByteSlice(w.Bytes())
	require.Equal(t, d, v)
	require.Equal(t, w.Bytes(), ByteSliceFromUint32Slice(d))
}

func TestUnsafeSliceInt32(t *testing.T) {
	w := &bytes.Buffer{}
	d := []int32{0xdead, 0xbeef, 0xb334}
	binary.Write(w, binary.LittleEndian, d)
	v := Int32SliceFromByteSlice(w.Bytes())
	require.Equal(t, d, v)
	require.Equal(t, w.Bytes(), ByteSliceFromInt32Slice(d))
}

func TestUnsafeSliceUint16(t *testing.T) {
	w := &bytes.Buffer{}
	d := []uint16{0xdead, 0xbeef, 0xb334}
	binary.Write(w, binary.LittleEndian, d)
	v := Uint16SliceFromByteSlice(w.Bytes())
	require.Equal(t, d, v)
	require.Equal(t, w.Bytes(), ByteSliceFromUint16Slice(d))
}

func TestUnsafeSliceInt16(t *testing.T) {
	w := &bytes.Buffer{}
	d := []int16{0xde, 0xad, 0xbe, 0xef, 0xb3, 0x34}
	binary.Write(w, binary.LittleEndian, d)
	v := Int16SliceFromByteSlice(w.Bytes())
	require.Equal(t, d, v)
	require.Equal(t, w.Bytes(), ByteSliceFromInt16Slice(d))
}

func TestUnsafeSliceUint8(t *testing.T) {
	w := &bytes.Buffer{}
	d := []uint8{0xde, 0xad, 0xbe, 0xef, 0xb3, 0x34}
	binary.Write(w, binary.LittleEndian, d)
	v := Uint8SliceFromByteSlice(w.Bytes())
	require.Equal(t, d, v)
	require.Equal(t, w.Bytes(), ByteSliceFromUint8Slice(d))
}

func TestUnsafeSliceInt8(t *testing.T) {
	w := &bytes.Buffer{}
	d := []int8{0xd, 0xe, 0xa, 0xd, 0xb, 0xe, 0xe, 0xf, 0xb, 0x3, 0x3, 0x4}
	binary.Write(w, binary.LittleEndian, d)
	v := Int8SliceFromByteSlice(w.Bytes())
	require.Equal(t, d, v)
	require.Equal(t, w.Bytes(), ByteSliceFromInt8Slice(d))
}

type Struct struct {
	A uint8
	B uint32
}

func makeTestStructBuffer() []byte {
	w := &bytes.Buffer{}
	a := &Struct{0xab, 0xdead}
	b := &Struct{0xce, 0xbeef}
	// Write struct values with padding
	binary.Write(w, binary.LittleEndian, a.A)
	w.Write([]byte{0, 0, 0})
	binary.Write(w, binary.LittleEndian, a.B)
	binary.Write(w, binary.LittleEndian, b.A)
	w.Write([]byte{0, 0, 0})
	binary.Write(w, binary.LittleEndian, b.B)
	return w.Bytes()
}

func TestUnsafeSliceStruct(t *testing.T) {
	var v []Struct
	b := makeTestStructBuffer()
	require.Nil(t, v)
	StructSliceFromByteSlice(b, &v)
	require.NotNil(t, v)
	require.Equal(t, len(v), 2)
	require.Equal(t, v[0].A, uint8(0xab))
	require.Equal(t, v[0].B, uint32(0xdead))
	require.Equal(t, v[1].A, uint8(0xce))
	require.Equal(t, v[1].B, uint32(0xbeef))
}

func TestByteSliceFromStructSlice(t *testing.T) {
	a := []Struct{
		Struct{0xab, 0xdead},
		Struct{0xce, 0xbeef},
	}
	b := ByteSliceFromStructSlice(a)
	require.Equal(t, 16, len(b))
	require.Equal(t, makeTestStructBuffer(), b)
	require.True(t, bytes.Compare(makeTestStructBuffer(), b) == 0)

	b = ByteSliceFromStructSlice([]Struct{})
	require.Equal(t, len(b), 0)
}

func TestByteSliceFromString(t *testing.T) {
	s := "life after 💀"
	b := ByteSliceFromString(s)
	require.Equal(t, s, string(b))
}

func TestStringFromByteSlice(t *testing.T) {
	b := []byte("life after 💀")
	s := StringFromByteSlice(b)
	require.Equal(t, s, string(b))
}
