package bytealg

import (
	"reflect"
	"unsafe"

	"github.com/koykov/fastconv"
)

// Byte sequence.
type Byteptr struct {
	// Offset in virtual memory.
	o uint64
	// Length of bytes array.
	l int
}

// Set new offset and length.
func (m *Byteptr) Set(o uint64, l int) {
	m.o, m.l = o, l
}

func (m *Byteptr) SetOffset(offset uint64) {
	m.o = offset
}

// Gen offset in virtual memory.
func (m *Byteptr) Offset() uint64 {
	return m.o
}

func (m *Byteptr) SetLen(len int) {
	m.l = len
}

// Get length of underlying byte array.
func (m *Byteptr) Len() int {
	return m.l
}

// Convert byte sequence to string.
func (m *Byteptr) String() string {
	return fastconv.B2S(m.Bytes())
}

// Convert byte sequence to byte array.
func (m *Byteptr) Bytes() []byte {
	h := reflect.SliceHeader{
		Data: uintptr(m.o),
		Len:  m.l,
		Cap:  m.l,
	}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func (m *Byteptr) Reset() {
	m.Set(0, 0)
}
