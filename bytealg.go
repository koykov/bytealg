package bytealg

import (
	"bytes"
	"reflect"
	"unicode"
	"unicode/utf8"
	"unsafe"

	"github.com/koykov/byteseq"
	"github.com/koykov/entry"
)

const (
	// Trim directions.
	trimBoth  = 0
	trimLeft  = 1
	trimRight = 2
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

// Trim makes fast and alloc-free trim.
func Trim[T byteseq.Byteseq](p, cut T) T {
	return trim(p, cut, trimBoth)
}

func TrimLeft[T byteseq.Byteseq](p, cut T) T {
	return trim(p, cut, trimLeft)
}

func TrimRight[T byteseq.Byteseq](p, cut T) T {
	return trim(p, cut, trimRight)
}

// Generic trim.
//
// Just calculates trim edges and return sub-slice.
func trim[T byteseq.Byteseq](p, cut T, dir int) T {
	l, r := 0, len(p)-1
	pb, cb := byteseq.Q2B(p), byteseq.Q2B(cut)
	if dir == trimBoth || dir == trimLeft {
		for ; l < len(pb); l++ {
			if !bytes.Contains(cb, []byte{pb[l]}) {
				break
			}
		}
	}
	if dir == trimBoth || dir == trimRight {
		for ; r >= l; r-- {
			if !bytes.Contains(cb, []byte{pb[r]}) {
				break
			}
		}
	}
	return byteseq.B2Q[T](pb[l : r+1])
}

// AppendSplit splits s to buf using sep as separator.
//
// This function if an alloc-free replacement of bytes.Split() function.
func AppendSplit[T byteseq.Byteseq](buf []T, s, sep T, n int) []T {
	if len(s) == 0 {
		return buf
	}
	sb, pb := byteseq.Q2B(s), byteseq.Q2B(sep)
	if n < 0 {
		n = bytes.Count(sb, pb) + 1
	}
	i := 0
	for i < n {
		m := bytes.Index(sb, pb)
		if m < 0 {
			break
		}
		buf = append(buf, byteseq.B2Q[T](sb[:m:m]))
		sb = sb[m+len(sep):]
		i++
	}
	buf = append(buf, byteseq.B2Q[T](sb))
	return buf[:i+1]
}

// AppendSplitEntry splits s to buf using sep as separator.
//
// buf contains entry.Entry64 records instead of substrings.
func AppendSplitEntry[T byteseq.Byteseq](buf []entry.Entry64, s, sep T, n int) []entry.Entry64 {
	if len(s) == 0 {
		return buf
	}
	sb, pb := byteseq.Q2B(s), byteseq.Q2B(sep)
	if n < 0 {
		n = bytes.Count(sb, pb) + 1
	}
	var off int
	i := 0
	for i < n {
		m := bytes.Index(sb, pb)
		if m < 0 {
			break
		}
		var e entry.Entry64
		e.Encode(uint32(off), uint32(off+m))
		buf = append(buf, e)
		sb = sb[m+len(sep):]
		off += m + len(sep)
		i++
	}
	var e entry.Entry64
	e.Encode(uint32(off), uint32(off+len(sb)))
	buf = append(buf, e)
	return buf[:i+1]
}

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

// ToUpper is an alloc-free replacement of bytes.ToUpper() function.
func ToUpper[T byteseq.Byteseq](p T) T { return Map(unicode.ToUpper, p) }

// ToLower is an alloc-free replacement of bytes.ToLower() function.
func ToLower[T byteseq.Byteseq](p T) T { return Map(unicode.ToLower, p) }

// ToTitle is an alloc-free replacement of bytes.ToTitle() function.
func ToTitle[T byteseq.Byteseq](p T) T { return Map(unicode.ToTitle, p) }

// Map returns modified p with all its characters modified according to the mapping function.
//
// See bytes.Map() function for details.
func Map[T byteseq.Byteseq](mapping func(r rune) rune, p T) T {
	maxbytes := len(p)
	nbytes := 0
	pb := byteseq.Q2B(p)
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

// Copy makes a copy of byte array.
func Copy[T byteseq.Byteseq](p T) T {
	cpy := append([]byte(nil), p...)
	return byteseq.B2Q[T](cpy)
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

// GrowDelta increases length of byte array to actual length + delta.
//
// See Grow().
func GrowDelta(p []byte, delta int) []byte {
	return Grow(p, len(p)+delta)
}
