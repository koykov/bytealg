package bytealg

import (
	"reflect"
	"unsafe"
)

// Byte sequence.
type Byteptr struct {
	// Offset in virtual memory.
	offset uint64
	// Limit of bytes array.
	limit int
}

// Set new offset and limit.
func (p *Byteptr) Set(offset uint64, limit int) {
	p.offset, p.limit = offset, limit
}

func (p *Byteptr) SetOffset(offset uint64) {
	p.offset = offset
}

// Gen offset in virtual memory.
func (p *Byteptr) Offset() uint64 {
	return p.offset
}

func (p *Byteptr) SetLimit(limit int) {
	p.limit = limit
}

// Get limit of underlying byte array.
func (p *Byteptr) Limit() int {
	return p.limit
}

// Convert byte sequence to string.
func (p *Byteptr) String() string {
	h := reflect.StringHeader{
		Data: uintptr(p.offset),
		Len:  p.limit,
	}
	return *(*string)(unsafe.Pointer(&h))
}

// Convert byte sequence to byte array.
func (p *Byteptr) Bytes() []byte {
	h := reflect.SliceHeader{
		Data: uintptr(p.offset),
		Len:  p.limit,
		Cap:  p.limit,
	}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func (p *Byteptr) Reset() {
	p.Set(0, 0)
}
