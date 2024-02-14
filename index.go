package bytealg

import (
	"bytes"

	"github.com/koykov/byteseq"
)

// IndexAt is equal to bytes.Index() but doesn't consider occurrences of sep in p[:at].
func IndexAt[T byteseq.Byteseq](p, sep T, at int) int {
	if at < 0 || at >= len(p) {
		return -1
	}
	pb, sb := byteseq.Q2B(p), byteseq.Q2B(sep)
	i := bytes.Index(pb[at:], sb)
	if i < 0 {
		return -1
	}
	return i + at
}

// IndexAnyAt is equal to bytes.IndexAny() but doesn't consider occurrences of sep in p[:at].
func IndexAnyAt[T byteseq.Byteseq](p, sep T, at int) int {
	if at < 0 || at >= len(p) {
		return -1
	}
	pb, ss := byteseq.Q2B(p), byteseq.Q2S(sep)
	i := bytes.IndexAny(pb[at:], ss)
	if i < 0 {
		return -1
	}
	return i + at
}

// IndexByteAt returns the index of the first instance of c in p (from position at), or -1 if c is not present in p.
func IndexByteAt[T byteseq.Byteseq](p T, c byte, at int) int {
	pb := byteseq.Q2B(p)
	if at < 0 || at >= len(pb) {
		return -1
	}
	i := bytes.IndexByte(pb[at:], c)
	if i < 0 {
		return -1
	}
	return i + at
}

// HasByte checks if c is present in p.
func HasByte[T byteseq.Byteseq](p T, c byte) bool {
	pb := byteseq.Q2B(p)
	return bytes.IndexByte(pb, c) != -1
}

// HasByteAt checks if c is present in p (from position at).
func HasByteAt[T byteseq.Byteseq](p T, c byte, at int) bool {
	return IndexByteAt(p, c, at) != -1
}

var _, _ = HasByte[string], HasByteAt[string]
