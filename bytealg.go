package bytealg

import (
	"bytes"
	"unicode"
	"unicode/utf8"
	"unsafe"

	"github.com/koykov/byteconv"
	"github.com/koykov/byteseq"
)

// EqualSet checks if two slices of bytes slices is equal.
func EqualSet[T byteseq.Byteseq](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !bytes.Equal(byteseq.Q2B(a[i]), byteseq.Q2B(b[i])) {
			return false
		}
	}
	return true
}

// ToUpper is an alloc-free replacement of bytes.ToUpper() function.
func ToUpper[T byteseq.Byteseq](p T) T { return Map(unicode.ToUpper, p) }

func ToUpperBytes(p []byte) []byte { return MapBytes(unicode.ToUpper, p) }

func ToUpperString(p string) string { return MapString(unicode.ToUpper, p) }

// ToLower is an alloc-free replacement of bytes.ToLower() function.
func ToLower[T byteseq.Byteseq](p T) T { return Map(unicode.ToLower, p) }

func ToLowerBytes(p []byte) []byte { return MapBytes(unicode.ToLower, p) }

func ToLowerString(p string) string { return MapString(unicode.ToLower, p) }

// ToTitle is an alloc-free replacement of bytes.ToTitle() function.
func ToTitle[T byteseq.Byteseq](p T) T { return Map(unicode.ToTitle, p) }

func ToTitleBytes(p []byte) []byte { return MapBytes(unicode.ToTitle, p) }

func ToTitleString(p string) string { return MapString(unicode.ToTitle, p) }

// Map returns modified x with all its characters modified according to the mapping function.
//
// See bytes.Map()/strings.Map() function for details.
func Map[T byteseq.Byteseq](mapping func(r rune) rune, x T) T {
	if p, ok := byteseq.ToBytes(x); ok {
		r := MapBytes(mapping, p)
		return *(*T)(unsafe.Pointer(&r))
	}
	if s, ok := byteseq.ToString(x); ok {
		r := MapString(mapping, s)
		return *(*T)(unsafe.Pointer(&r))
	}
	return x
}

func MapBytes(mapping func(r rune) rune, p []byte) []byte {
	maxbytes := len(p)
	nbytes := 0
	pb := p
	for i := 0; i < len(p); {
		wid := 1
		r := rune(p[i])
		if r >= utf8.RuneSelf {
			r, wid = utf8.DecodeRune(pb[i:])
		}
		r = mapping(r)
		if r >= 0 {
			rl := utf8.RuneLen(r)
			if rl < 0 {
				rl = len(string(utf8.RuneError))
			}
			if nbytes+rl > maxbytes {
				maxbytes = maxbytes*2 + utf8.UTFMax
				pb = Grow(pb, maxbytes)
			}
			nbytes += utf8.EncodeRune(pb[nbytes:maxbytes], r)
		}
		i += wid
	}
	return p[:nbytes]
}

func MapString(mapping func(r rune) rune, s string) string {
	buf := make([]byte, 0, len(s)*2)
	buf = append(buf, s...)
	buf = MapBytes(mapping, buf)
	return byteconv.B2S(buf)
}

// Copy makes a copy of byte sequence.
func Copy[T byteseq.Byteseq](p T) T {
	cpy := append([]byte(nil), p...)
	return byteseq.B2Q[T](cpy)
}

// CopyBytes makes a copy of byte slice.
func CopyBytes(p []byte) (r []byte) {
	return append(r, p...)
}

// CopyString makes a copy of string.
func CopyString(s string) (r string) {
	var buf []byte
	buf = append(buf, s...)
	return byteconv.B2S(buf)
}

// Grow increases length of the byte array.
//
// Two cases are possible:
// * byte array has enough space;
// * need to add extra space to the array.
func Grow(p []byte, newLen int) []byte {
	if newLen <= 0 {
		return p
	}
	// Get byte array header.
	h := *(*byteconv.SliceHeader)(unsafe.Pointer(&p))
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

// GrowDelta increases length of byte array to actual length + delta.
//
// See Grow().
func GrowDelta(p []byte, delta int) []byte {
	return Grow(p, len(p)+delta)
}
