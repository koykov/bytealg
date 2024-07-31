package bytealg

import (
	"math"
	"unsafe"

	"github.com/koykov/byteconv"
	"github.com/koykov/byteseq"
)

// SkipFmt4 moves offset to first non-fmt4 byte in x.
// Returns new offset and EOF flag.
func SkipFmt4[T byteseq.Q](x T, offset int) (int, bool) {
	if p, ok := byteseq.ToBytes(x); ok {
		return skipFmt4(p, len(p), offset)
	}
	if s, ok := byteseq.ToString(x); ok {
		return skipFmt4(byteconv.S2B(s), len(s), offset)
	}
	return offset, false
}

// SkipBytesFmt4 moves offset to first non-fmt4 byte in bytes p.
// Returns new offset and EOF flag.
func SkipBytesFmt4(p []byte, offset int) (int, bool) {
	return skipFmt4(p, len(p), offset)
}

// SkipStringFmt4 moves offset to first non-fmt4 byte in string s.
// Returns new offset and EOF flag.
func SkipStringFmt4(s string, offset int) (int, bool) {
	return skipFmt4(byteconv.S2B(s), len(s), offset)
}

const skipFmt4TableThreshold = 512

// Table based approach of fmt skip.
func skipFmt4(src []byte, n, offset int) (int, bool) {
	_ = src[n-1]
	_ = skipTable4[math.MaxUint8]
	if n-offset > skipFmt4TableThreshold {
		offset, _ = skipFmtBin8(src, n, offset)
	}
	for ; offset < n && skipTable4[src[offset]]; offset++ {
	}
	return offset, offset == n
}

// Binary based approach of fmt skip.
func skipFmtBin8(src []byte, n, offset int) (int, bool) {
	_ = src[n-1]
	_ = skipTable4[math.MaxUint8]
	if *(*uint64)(unsafe.Pointer(&src[offset])) == binNlSpace7 {
		offset += 8
		for offset < n && *(*uint64)(unsafe.Pointer(&src[offset])) == binSpace8 {
			offset += 8
		}
	}
	return offset, false
}

var (
	skipTable4  = [math.MaxUint8 + 1]bool{}
	binNlSpace7 uint64
	binSpace8   uint64
)

func init() {
	skipTable4[bfmt4space] = true
	skipTable4[bfmt4tab] = true
	skipTable4[bfmt4nl] = true
	skipTable4[bfmt4cr] = true

	binNlSpace7Bytes, binSpace8Bytes := []byte("\n       "), []byte("        ")
	binNlSpace7, binSpace8 = *(*uint64)(unsafe.Pointer(&binNlSpace7Bytes[0])), *(*uint64)(unsafe.Pointer(&binSpace8Bytes[0]))
}

var _, _, _ = SkipFmt4[string], SkipBytesFmt4, SkipStringFmt4
