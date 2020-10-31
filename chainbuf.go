package bytealg

import (
	"reflect"
	"unsafe"

	"github.com/koykov/any2bytes"
	"github.com/koykov/fastconv"
)

// Primitive byte buffer with chain call support.
type ChainBuf []byte

// Get contents of the buffer.
func (b *ChainBuf) Bytes() []byte {
	return *b
}

// Get contents of the buffer as string.
func (b *ChainBuf) String() string {
	return fastconv.B2S(*b)
}

// Write bytes to the buffer.
func (b *ChainBuf) Write(p []byte) *ChainBuf {
	*b = append(*b, p...)
	return b
}

// Write single byte.
func (b *ChainBuf) WriteByte(p byte) *ChainBuf {
	*b = append(*b, p)
	return b
}

// Write string to the buffer.
func (b *ChainBuf) WriteStr(s string) *ChainBuf {
	*b = append(*b, s...)
	return b
}

// Write integer value to the buffer.
func (b *ChainBuf) WriteInt(i int64) *ChainBuf {
	*b, _ = any2bytes.IntToBytes(*b, i)
	return b
}

// Write float value to the buffer.
func (b *ChainBuf) WriteFloat(f float64) *ChainBuf {
	*b, _ = any2bytes.FloatToBytes(*b, f)
	return b
}

// Write boolean value to the buffer.
func (b *ChainBuf) WriteBool(v bool) *ChainBuf {
	*b, _ = any2bytes.BoolToBytes(*b, v)
	return b
}

// Replace old to new bytes in buffer.
func (b *ChainBuf) Replace(old, new []byte, n int) *ChainBuf {
	if b.Len() == 0 || n == 0 {
		return b
	}
	var i, at, c int
	// Use the same byte buffer to make replacement and avoid alloc.
	dst := (*b)[b.Len():]
	for {
		if i = IndexAt(*b, old, at); i < 0 || c == n {
			dst = append(dst, (*b)[at:]...)
			break
		}
		dst = append(dst, (*b)[at:i]...)
		dst = append(dst, new...)
		at = i + len(old)
		c++
	}
	// Move result to the beginning of buffer.
	b.Reset().Write(dst)
	return b
}

// Replace old to new strings in buffer.
func (b *ChainBuf) ReplaceStr(old, new string, n int) *ChainBuf {
	return b.Replace(fastconv.S2B(old), fastconv.S2B(new), n)
}

// Get length of the buffer.
func (b *ChainBuf) Len() int {
	return len(*b)
}

// Get capacity of the buffer.
func (b *ChainBuf) Cap() int {
	return cap(*b)
}

// Grow length and capacity of the buffer.
func (b *ChainBuf) Grow(cap int) *ChainBuf {
	if cap < 0 {
		return b
	}
	// Get buffer header.
	h := *(*reflect.SliceHeader)(unsafe.Pointer(b))
	if cap < h.Cap {
		// Just increase header's length if capacity allows
		h.Len = cap
		// .. and restore the buffer from the header.
		*b = *(*[]byte)(unsafe.Pointer(&h))
	} else {
		// Append necessary space.
		*b = append(*b, make([]byte, cap-b.Len())...)
	}
	return b
}

// Reset length of the buffer.
func (b *ChainBuf) Reset() *ChainBuf {
	*b = (*b)[:0]
	return b
}

// Conversion to bytes function.
func ChainBufToBytes(dst []byte, val interface{}) ([]byte, error) {
	if b, ok := val.(*ChainBuf); ok {
		dst = append(dst, *b...)
		return dst, nil
	}
	return dst, any2bytes.ErrUnknownType
}
