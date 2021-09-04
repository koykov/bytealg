package bytealg

import (
	"reflect"
	"unsafe"

	"github.com/koykov/fastconv"
	"github.com/koykov/x2bytes"
)

// Accumulative byte buffer.
//
// Extended version of ChainBuf.
// todo merge with ChainBuf.
type AccumulativeBuffer struct {
	buf []byte
	off int
}

// Get contents of the buffer.
func (b *AccumulativeBuffer) Bytes() []byte {
	return b.buf
}

// Get copy of the buffer.
func (b *AccumulativeBuffer) BytesCopy() []byte {
	return Copy(b.buf)
}

// Get contents of the buffer as string.
func (b *AccumulativeBuffer) String() string {
	return fastconv.B2S(b.buf)
}

// Get copy of the buffer as string.
func (b *AccumulativeBuffer) StringCopy() string {
	return CopyStr(fastconv.B2S(b.buf))
}

// Write bytes to the buffer.
func (b *AccumulativeBuffer) Write(p []byte) *AccumulativeBuffer {
	b.buf = append(b.buf, p...)
	return b
}

// Write single byte.
func (b *AccumulativeBuffer) WriteByte(p byte) *AccumulativeBuffer {
	b.buf = append(b.buf, p)
	return b
}

// Write string to the buffer.
func (b *AccumulativeBuffer) WriteStr(s string) *AccumulativeBuffer {
	b.buf = append(b.buf, s...)
	return b
}

// Write integer value to the buffer.
func (b *AccumulativeBuffer) WriteInt(i int64) *AccumulativeBuffer {
	b.buf, _ = x2bytes.IntToBytes(b.buf, i)
	return b
}

// Write unsigned integer value to the buffer.
func (b *AccumulativeBuffer) WriteUint(u uint64) *AccumulativeBuffer {
	b.buf, _ = x2bytes.UintToBytes(b.buf, u)
	return b
}

// Write float value to the buffer.
func (b *AccumulativeBuffer) WriteFloat(f float64) *AccumulativeBuffer {
	b.buf, _ = x2bytes.FloatToBytes(b.buf, f)
	return b
}

// Write boolean value to the buffer.
func (b *AccumulativeBuffer) WriteBool(v bool) *AccumulativeBuffer {
	b.buf, _ = x2bytes.BoolToBytes(b.buf, v)
	return b
}

func (b *AccumulativeBuffer) WriteX(v interface{}) *AccumulativeBuffer {
	b.buf, _ = x2bytes.ToBytes(b.buf, v)
	return b
}

// Replace old to new bytes in buffer.
func (b *AccumulativeBuffer) Replace(old, new []byte, n int) *AccumulativeBuffer {
	if b.Len() == 0 || n == 0 {
		return b
	}
	var i, at, c int
	// Use the same byte buffer to make replacement and avoid alloc.
	dst := (b.buf)[b.Len():]
	for {
		if i = IndexAt(b.buf, old, at); i < 0 || c == n {
			dst = append(dst, b.buf[at:]...)
			break
		}
		dst = append(dst, b.buf[at:i]...)
		dst = append(dst, new...)
		at = i + len(old)
		c++
	}
	// Move result to the beginning of buffer.
	b.Reset().Write(dst)
	return b
}

// Replace old to new strings in buffer.
func (b *AccumulativeBuffer) ReplaceStr(old, new string, n int) *AccumulativeBuffer {
	return b.Replace(fastconv.S2B(old), fastconv.S2B(new), n)
}

// Replace all old to new bytes in buffer.
func (b *AccumulativeBuffer) ReplaceAll(old, new []byte) *AccumulativeBuffer {
	return b.Replace(old, new, -1)
}

// Replace all old to new strings in buffer.
func (b *AccumulativeBuffer) ReplaceStrAll(old, new string) *AccumulativeBuffer {
	return b.Replace(fastconv.S2B(old), fastconv.S2B(new), -1)
}

// Get length of the buffer.
func (b *AccumulativeBuffer) Len() int {
	return len(b.buf)
}

// Get capacity of the buffer.
func (b *AccumulativeBuffer) Cap() int {
	return cap(b.buf)
}

// Grow length of the buffer.
func (b *AccumulativeBuffer) Grow(newLen int) *AccumulativeBuffer {
	if newLen <= 0 {
		return b
	}
	// Get buffer header.
	h := *(*reflect.SliceHeader)(unsafe.Pointer(b))
	if newLen < h.Cap {
		// Just increase header's length if capacity allows
		h.Len = newLen
		// .. and restore the buffer from the header.
		b.buf = *(*[]byte)(unsafe.Pointer(&h))
	} else {
		// Append necessary space.
		b.buf = append(b.buf, make([]byte, newLen-b.Len())...)
	}
	return b
}

// Grow length of the buffer to actual length + delta.
//
// See Grow().
func (b *AccumulativeBuffer) GrowDelta(delta int) *AccumulativeBuffer {
	return b.Grow(b.Len() + delta)
}

// Reset length of the buffer.
func (b *AccumulativeBuffer) Reset() *AccumulativeBuffer {
	b.buf = b.buf[:0]
	b.off = 0
	return b
}

// Stake out current offset for further use.
func (b *AccumulativeBuffer) StakeOut() *AccumulativeBuffer {
	b.off = b.Len()
	return b
}

// Get staked offset.
func (b *AccumulativeBuffer) StakedOffset() int {
	return b.off
}

// Get accumulated bytes from staked offset.
func (b *AccumulativeBuffer) StakedBytes() []byte {
	if b.off >= b.Len() {
		return nil
	}
	return b.buf[b.off:]
}

// Get copy of accumulated bytes from staked offset.
func (b *AccumulativeBuffer) StakedBytesCopy() []byte {
	return Copy(b.StakedBytes())
}

// Get accumulated bytes as string.
func (b *AccumulativeBuffer) StakedString() string {
	if b.off >= b.Len() {
		return ""
	}
	return b.String()[b.off:]
}

// Get copy of accumulated bytes as string.
func (b *AccumulativeBuffer) StakedStringCopy() string {
	return CopyStr(b.StakedString())
}
