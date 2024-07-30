package bytealg

import (
	"unsafe"

	"github.com/koykov/byteconv"
	"github.com/koykov/byteseq"
)

func SkipFmt4[T byteseq.Q](x T, offset int) (int, bool) {
	if p, ok := byteseq.ToBytes(x); ok {
		return skipFmt4(p, len(p), offset)
	}
	if s, ok := byteseq.ToString(x); ok {
		return skipFmt4(byteconv.S2B(s), len(s), offset)
	}
	return offset, false
}

// Table based approach of fmt skip.
func skipFmt4(src []byte, n, offset int) (int, bool) {
	_ = src[n-1]
	_ = skipTable4[255]
	if n-offset > 512 {
		offset, _ = skipFmtBin8(src, n, offset)
	}
	for ; skipTable4[src[offset]]; offset++ {
	}
	return offset, offset == n
}

// Binary based approach of fmt skip.
func skipFmtBin8(src []byte, n, offset int) (int, bool) {
	_ = src[n-1]
	_ = skipTable4[255]
	if *(*uint64)(unsafe.Pointer(&src[offset])) == binNlSpace7 {
		offset += 8
		for offset < n && *(*uint64)(unsafe.Pointer(&src[offset])) == binSpace8 {
			offset += 8
		}
	}
	return offset, false
}

var (
	skipTable4  = [256]bool{}
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

var _ = SkipFmt4[string]
