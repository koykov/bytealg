package bytealg

import (
	"bytes"
	"reflect"
	"unicode"
	"unicode/utf8"
	"unsafe"
)

const (
	// Trim directions.
	trimBoth  = 0
	trimLeft  = 1
	trimRight = 2
)

var (
	// Suppress go vet warnings.
	_, _, _ = TrimLeft, TrimRight, ToTitle
)

// Check if two slices of bytes slices is equal.
func EqualSet(a, b [][]byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !bytes.Equal(a[i], b[i]) {
			return false
		}
	}
	return true
}

// Fast and alloc-free trim.
func Trim(p, cut []byte) []byte {
	return trim(p, cut, trimBoth)
}

// Left trim.
func TrimLeft(p, cut []byte) []byte {
	return trim(p, cut, trimLeft)
}

// Right trim.
func TrimRight(p, cut []byte) []byte {
	return trim(p, cut, trimRight)
}

// Generic trim.
//
// Just calculates trim edges and return sub-slice.
func trim(p, cut []byte, dir int) []byte {
	l, r := 0, len(p)-1
	if dir == trimBoth || dir == trimLeft {
		for ; l < len(p); l++ {
			if !bytes.Contains(cut, []byte{p[l]}) {
				break
			}
		}
	}
	if dir == trimBoth || dir == trimRight {
		for ; r >= l; r-- {
			if !bytes.Contains(cut, []byte{p[r]}) {
				break
			}
		}
	}
	return p[l : r+1]
}

// Split s to buf using sep as separator.
//
// This function if an alloc-free replacement of bytes.Split() function.
func AppendSplit(buf [][]byte, s, sep []byte, n int) [][]byte {
	if len(s) == 0 {
		return buf
	}
	if n < 0 {
		n = bytes.Count(s, sep) + 1
	}
	i := 0
	for i < n {
		m := bytes.Index(s, sep)
		if m < 0 {
			break
		}
		buf = append(buf, s[:m:m])
		s = s[m+len(sep):]
		i++
	}
	buf = append(buf, s)
	return buf[:i+1]
}

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

// ToUpper is an alloc-free replacement of bytes.ToUpper() function.
func ToUpper(p []byte) []byte { return Map(unicode.ToUpper, p) }

// ToLower is an alloc-free replacement of bytes.ToLower() function.
func ToLower(p []byte) []byte { return Map(unicode.ToLower, p) }

// ToTitle is an alloc-free replacement of bytes.ToTitle() function.
func ToTitle(p []byte) []byte { return Map(unicode.ToTitle, p) }

// Map returns modified p with all its characters modified according to the mapping function.
//
// See bytes.Map() function for details.
func Map(mapping func(r rune) rune, p []byte) []byte {
	maxbytes := len(p)
	nbytes := 0
	for i := 0; i < len(p); {
		wid := 1
		r := rune(p[i])
		if r >= utf8.RuneSelf {
			r, wid = utf8.DecodeRune(p[i:])
		}
		r = mapping(r)
		if r >= 0 {
			rl := utf8.RuneLen(r)
			if rl < 0 {
				rl = len(string(utf8.RuneError))
			}
			if nbytes+rl > maxbytes {
				maxbytes = maxbytes*2 + utf8.UTFMax
				p = Grow(p, maxbytes)
			}
			nbytes += utf8.EncodeRune(p[nbytes:maxbytes], r)
		}
		i += wid
	}
	return p[:nbytes]
}

// Make a copy of byte array.
func Copy(p []byte) []byte {
	return append([]byte(nil), p...)
}

// Increase length of the byte array.
//
// Two cases are possible:
// * byte array has enough space;
// * need to add extra space to the array.
func Grow(p []byte, newLen int) []byte {
	if newLen <= 0 {
		return p
	}
	// Get byte array header.
	h := *(*reflect.SliceHeader)(unsafe.Pointer(&p))
	if newLen <= h.Cap {
		// p already has enough space.
		h.Len = newLen
		p = *(*[]byte)(unsafe.Pointer(&h))
	} else {
		// Need to add extra space to p.
		p = append(p, make([]byte, newLen-h.Len)...)
	}
	return p
}

// Increase length of byte array to actual length + delta.
//
// See Grow().
func GrowDelta(p []byte, delta int) []byte {
	return Grow(p, len(p)+delta)
}
