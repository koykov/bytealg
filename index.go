package bytealg

import (
	"bytes"

	"github.com/koykov/byteconv"
)

// IndexAt is equal to bytes.Index() but doesn't consider occurrences of sep in p[:at].
func IndexAt(p, sep []byte, at int) int {
	if at < 0 || at >= len(p) {
		return -1
	}
	i := bytes.Index(p[at:], sep)
	if i < 0 {
		return -1
	}
	return i + at
}

// IndexAnyAt is equal to bytes.IndexAny() but doesn't consider occurrences of sep in p[:at].
func IndexAnyAt(p, sep []byte, at int) int {
	if at < 0 || at >= len(p) {
		return -1
	}
	i := bytes.IndexAny(p[at:], byteconv.B2S(sep))
	if i < 0 {
		return -1
	}
	return i + at
}

// IndexByteAt returns the index of the first instance of c in p (from position at), or -1 if c is not present in p.
func IndexByteAt(p []byte, c byte, at int) int {
	if at < 0 || at >= len(p) {
		return -1
	}
	i := bytes.IndexByte(p[at:], c)
	if i < 0 {
		return -1
	}
	return i + at
}

// HasByte checks if c is present in p.
func HasByte(p []byte, c byte) bool {
	return bytes.IndexByte(p, c) != -1
}

// HasByteAt checks if c is present in p (from position at).
func HasByteAt(p []byte, c byte, at int) bool {
	return IndexByteAt(p, c, at) != -1
}

var _, _ = HasByte, HasByteAt
