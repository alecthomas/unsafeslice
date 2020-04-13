// Package unsafeslice contains functions for zero-copy casting between typed slices and byte slices.
package unsafeslice

import (
	"reflect"
	"unsafe"
)

// Useful constants.
const (
	Uint64Size = 8
	Uint32Size = 4
	Uint16Size = 2
	Uint8Size  = 1
)

func newRawSliceHeader(sh *reflect.SliceHeader, b []byte, stride int) *reflect.SliceHeader {
	sh.Len = len(b) / stride
	sh.Cap = len(b) / stride
	sh.Data = (uintptr)(unsafe.Pointer(&b[0]))
	return sh
}

func newSliceHeaderFromBytes(b []byte, stride int) unsafe.Pointer {
	sh := &reflect.SliceHeader{}
	return unsafe.Pointer(newRawSliceHeader(sh, b, stride))
}

func newSliceHeader(p unsafe.Pointer, size int) unsafe.Pointer {
	return unsafe.Pointer(&reflect.SliceHeader{
		Len:  size,
		Cap:  size,
		Data: uintptr(p),
	})
}

func Uint64SliceFromByteSlice(b []byte) []uint64 {
	return *(*[]uint64)(newSliceHeaderFromBytes(b, Uint64Size))
}

func ByteSliceFromUint64Slice(b []uint64) []byte {
	return *(*[]byte)(newSliceHeader(unsafe.Pointer(&b[0]), len(b)*Uint64Size))
}

func Int64SliceFromByteSlice(b []byte) []int64 {
	return *(*[]int64)(newSliceHeaderFromBytes(b, Uint64Size))
}

func ByteSliceFromInt64Slice(b []int64) []byte {
	return *(*[]byte)(newSliceHeader(unsafe.Pointer(&b[0]), len(b)*Uint64Size))
}

func Uint32SliceFromByteSlice(b []byte) []uint32 {
	return *(*[]uint32)(newSliceHeaderFromBytes(b, Uint32Size))
}

func ByteSliceFromUint32Slice(b []uint32) []byte {
	return *(*[]byte)(newSliceHeader(unsafe.Pointer(&b[0]), len(b)*Uint32Size))
}

func Int32SliceFromByteSlice(b []byte) []int32 {
	return *(*[]int32)(newSliceHeaderFromBytes(b, Uint32Size))
}

func ByteSliceFromInt32Slice(b []int32) []byte {
	return *(*[]byte)(newSliceHeader(unsafe.Pointer(&b[0]), len(b)*Uint32Size))
}

func Uint16SliceFromByteSlice(b []byte) []uint16 {
	return *(*[]uint16)(newSliceHeaderFromBytes(b, Uint16Size))
}

func ByteSliceFromUint16Slice(b []uint16) []byte {
	return *(*[]byte)(newSliceHeader(unsafe.Pointer(&b[0]), len(b)*Uint16Size))
}

func Int16SliceFromByteSlice(b []byte) []int16 {
	return *(*[]int16)(newSliceHeaderFromBytes(b, Uint16Size))
}

func ByteSliceFromInt16Slice(b []int16) []byte {
	return *(*[]byte)(newSliceHeader(unsafe.Pointer(&b[0]), len(b)*Uint16Size))
}

func Uint8SliceFromByteSlice(b []byte) []uint8 {
	return b
}

func ByteSliceFromUint8Slice(b []uint8) []byte {
	return b
}

func Int8SliceFromByteSlice(b []byte) []int8 {
	return *(*[]int8)(newSliceHeaderFromBytes(b, Uint8Size))
}

func ByteSliceFromInt8Slice(b []int8) []byte {
	return *(*[]byte)(newSliceHeader(unsafe.Pointer(&b[0]), len(b)*Uint8Size))
}

func ByteSliceFromString(s string) []byte {
	h := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return *(*[]byte)(newSliceHeader(unsafe.Pointer(h.Data), len(s)*Uint8Size))
}

func StringFromByteSlice(b []byte) string {
	h := &reflect.StringHeader{
		Data: uintptr(unsafe.Pointer(&b[0])),
		Len:  len(b),
	}
	return *(*string)(unsafe.Pointer(h))
}

// Create a slice of structs from a slice of bytes.
//
// 		var v []Struct
// 		StructSliceFromByteSlice(bytes, &v)
//
// Elements in the byte array must be padded correctly. See unsafe.AlignOf, et al.
//
// Note that this is slower than the scalar primitives above as it uses reflection.
func StructSliceFromByteSlice(b []byte, out interface{}) {
	ptr := reflect.ValueOf(out)
	if ptr.Kind() != reflect.Ptr {
		panic("expected pointer to a slice of structs (*[]X)")
	}
	slice := ptr.Elem()
	if slice.Kind() != reflect.Slice {
		panic("expected pointer to a slice of structs (*[]X)")
	}
	// TODO: More checks, such as ensuring that:
	// - elements are NOT pointers
	// - structs do not contain pointers, slices or maps
	stride := int(slice.Type().Elem().Size())
	if len(b)%stride != 0 {
		panic("size of byte buffer is not a multiple of struct size")
	}
	sh := (*reflect.SliceHeader)(unsafe.Pointer(slice.UnsafeAddr()))
	newRawSliceHeader(sh, b, stride)
}

// ByteSliceFromStructSlice does what you would expect.
//
// Note that this is slower than the scalar primitives above as it uses reflection.
func ByteSliceFromStructSlice(s interface{}) []byte {
	slice := reflect.ValueOf(s)
	if slice.Kind() != reflect.Slice {
		panic("expected a slice of structs (*[]X)")
	}
	var length int
	var data uintptr
	if slice.Len() != 0 {
		elem := slice.Index(0)
		length = int(elem.Type().Size()) * slice.Len()
		data = elem.UnsafeAddr()
	}
	out := &reflect.SliceHeader{
		Len:  length,
		Cap:  length,
		Data: data,
	}
	return *(*[]byte)(unsafe.Pointer(out))
}
