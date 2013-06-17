package unsafeslice

import (
	"bytes"
	"encoding/binary"
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestUnsafeSlice64(t *testing.T) {
	w := &bytes.Buffer{}
	d := []uint64{0xdead, 0xbeef, 0xb334}
	binary.Write(w, binary.LittleEndian, d)
	v := Uint64SliceFromByteSlice(w.Bytes())
	assert.Equal(t, d, v)
}

func TestUnsafeSlice32(t *testing.T) {
	w := &bytes.Buffer{}
	d := []uint32{0xdead, 0xbeef, 0xb334}
	binary.Write(w, binary.LittleEndian, d)
	v := Uint32SliceFromByteSlice(w.Bytes())
	assert.Equal(t, d, v)
}

func TestUnsafeSlice16(t *testing.T) {
	w := &bytes.Buffer{}
	d := []uint16{0xdead, 0xbeef, 0xb334}
	binary.Write(w, binary.LittleEndian, d)
	v := Uint16SliceFromByteSlice(w.Bytes())
	assert.Equal(t, d, v)
}
