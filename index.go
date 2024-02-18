package bytealg

import (
	"bytes"
	"strings"

	"github.com/koykov/byteconv"
	"github.com/koykov/byteseq"
)

// group: generic versions

// IndexAt is equal to bytes.Index() but doesn't consider occurrences of sep in p[:at].
func IndexAt[T byteseq.Q](x, sep T, at int) int {
	if p, ok := byteseq.ToBytes(x); ok {
		ps, _ := byteseq.ToBytes(sep)
		return IndexAtBytes(p, ps, at)
	}
	if s, ok := byteseq.ToString(x); ok {
		ss, _ := byteseq.ToString(sep)
		return IndexAtString(s, ss, at)
	}
	return -1
}

// IndexAnyAt is equal to bytes.IndexAny() but doesn't consider occurrences of sep in p[:at].
func IndexAnyAt[T byteseq.Q](x, sep T, at int) int {
	if p, ok := byteseq.ToBytes(x); ok {
		ps, _ := byteseq.ToBytes(sep)
		return IndexAnyAtBytes(p, ps, at)
	}
	if s, ok := byteseq.ToString(x); ok {
		ss, _ := byteseq.ToString(sep)
		return IndexAnyAtString(s, ss, at)
	}
	return -1
}

// IndexByteAt returns the index of the first instance of c in p (from position at), or -1 if c is not present in p.
func IndexByteAt[T byteseq.Q](x T, c byte, at int) int {
	if p, ok := byteseq.ToBytes(x); ok {
		return IndexByteAtBytes(p, c, at)
	}
	if s, ok := byteseq.ToString(x); ok {
		return IndexByteAtString(s, c, at)
	}
	return -1
}

// HasByte checks if c is present in p.
func HasByte[T byteseq.Q](x T, c byte) bool {
	if p, ok := byteseq.ToBytes(x); ok {
		return HasByteBytes(p, c)
	}
	if s, ok := byteseq.ToString(x); ok {
		return HasByteString(s, c)
	}
	return false
}

// HasByteAt checks if c is present in p (from position at).
func HasByteAt[T byteseq.Q](x T, c byte, at int) bool {
	if p, ok := byteseq.ToBytes(x); ok {
		return HasByteAtBytes(p, c, at)
	}
	if s, ok := byteseq.ToString(x); ok {
		return HasByteAtString(s, c, at)
	}
	return false
}

var _ = HasByteAt[string]

// group: bytes versions

// IndexAtBytes is equal to bytes.Index() but doesn't consider occurrences of sep in p[:at].
func IndexAtBytes(p, sep []byte, at int) int {
	if at < 0 || at >= len(p) {
		return -1
	}
	i := bytes.Index(p[at:], sep)
	if i < 0 {
		return -1
	}
	return i + at
}

// IndexAnyAtBytes is equal to bytes.IndexAny() but doesn't consider occurrences of sep in p[:at].
func IndexAnyAtBytes(p, sep []byte, at int) int {
	if at < 0 || at >= len(p) {
		return -1
	}
	i := bytes.IndexAny(p[at:], byteconv.B2S(sep))
	if i < 0 {
		return -1
	}
	return i + at
}

// IndexByteAtBytes returns the index of the first instance of c in p (from position at), or -1 if c is not present in p.
func IndexByteAtBytes(p []byte, c byte, at int) int {
	if at < 0 || at >= len(p) {
		return -1
	}
	i := bytes.IndexByte(p[at:], c)
	if i < 0 {
		return -1
	}
	return i + at
}

// HasByteBytes checks if c is present in p.
func HasByteBytes(p []byte, c byte) bool {
	return bytes.IndexByte(p, c) != -1
}

// HasByteAtBytes checks if c is present in p (from position at).
func HasByteAtBytes(p []byte, c byte, at int) bool {
	return IndexByteAtBytes(p, c, at) != -1
}

// group: string versions

// IndexAtString is equal to bytes.Index() but doesn't consider occurrences of sep in p[:at].
func IndexAtString(s, sep string, at int) int {
	if at < 0 || at >= len(s) {
		return -1
	}
	i := strings.Index(s[at:], sep)
	if i < 0 {
		return -1
	}
	return i + at
}

// IndexAnyAtString is equal to bytes.IndexAny() but doesn't consider occurrences of sep in p[:at].
func IndexAnyAtString(s, sep string, at int) int {
	if at < 0 || at >= len(s) {
		return -1
	}
	i := strings.IndexAny(s[at:], sep)
	if i < 0 {
		return -1
	}
	return i + at
}

// IndexByteAtString returns the index of the first instance of c in p (from position at), or -1 if c is not present in p.
func IndexByteAtString(s string, c byte, at int) int {
	if at < 0 || at >= len(s) {
		return -1
	}
	i := strings.IndexByte(s[at:], c)
	if i < 0 {
		return -1
	}
	return i + at
}

// HasByteString checks if c is present in p.
func HasByteString(s string, c byte) bool {
	return HasByteAtString(s, c, 0)
}

// HasByteAtString checks if c is present in p (from position at).
func HasByteAtString(s string, c byte, at int) bool {
	return IndexByteAtString(s, c, at) != -1
}
