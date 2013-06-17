package unsafeslice

import (
	"reflect"
	"unsafe"
)

var (
	Uint64Size = 8
	Uint32Size = 4
	Uint16Size = 2
	Uint8Size  = 1
)

func newSliceHeader(b []byte, stride int) unsafe.Pointer {
	sh := &reflect.SliceHeader{}
	sh.Len = len(b) / stride
	sh.Cap = len(b) / stride
	sh.Data = (uintptr)(unsafe.Pointer(&b[0]))
	return unsafe.Pointer(sh)
}

func Uint64SliceFromByteSlice(b []byte) []uint64 {
	return *(*[]uint64)(newSliceHeader(b, Uint64Size))
}

func Int64SliceFromByteSlice(b []byte) []int64 {
	return *(*[]int64)(newSliceHeader(b, Uint64Size))
}

func Uint32SliceFromByteSlice(b []byte) []uint32 {
	return *(*[]uint32)(newSliceHeader(b, Uint32Size))
}

func Int32SliceFromByteSlice(b []byte) []int32 {
	return *(*[]int32)(newSliceHeader(b, Uint32Size))
}

func Uint16SliceFromByteSlice(b []byte) []uint16 {
	return *(*[]uint16)(newSliceHeader(b, Uint16Size))
}

func Int16SliceFromByteSlice(b []byte) []int16 {
	return *(*[]int16)(newSliceHeader(b, Uint16Size))
}

func Uint8SliceFromByteSlice(b []byte) []uint8 {
	return *(*[]uint8)(newSliceHeader(b, Uint8Size))
}

func Int8SliceFromByteSlice(b []byte) []int8 {
	return *(*[]int8)(newSliceHeader(b, Uint8Size))
}
