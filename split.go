package bytealg

import (
	"bytes"
	"strings"

	"github.com/koykov/byteseq"
	"github.com/koykov/entry"
)

// group: generic versions

// AppendSplit splits x to buf using sep as separator.
//
// This function if an alloc-free replacement of bytes.Split() function.
func AppendSplit[T byteseq.Byteseq](buf []T, x, sep T, n int) []T {
	if len(x) == 0 {
		return buf
	}
	sb, pb := byteseq.Q2B(x), byteseq.Q2B(sep)
	var i int
	for {
		m := bytes.Index(sb, pb)
		if m < 0 {
			break
		}
		buf = append(buf, byteseq.B2Q[T](sb[:m:m]))
		sb = sb[m+len(sep):]
		i++
		if n >= 0 && i >= n {
			break
		}
	}
	buf = append(buf, byteseq.B2Q[T](sb))
	return buf[:i+1]
}

// AppendSplitEntry splits x to buf using sep as separator.
//
// buf contains entry.Entry64 records instead of substrings.
func AppendSplitEntry[T byteseq.Byteseq](buf []entry.Entry64, s, sep T, n int) []entry.Entry64 {
	if len(s) == 0 {
		return buf
	}
	sb, pb := byteseq.Q2B(s), byteseq.Q2B(sep)
	var off int
	var i int
	for {
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
		if n >= 0 && i >= n {
			break
		}
	}
	var e entry.Entry64
	e.Encode(uint32(off), uint32(off+len(sb)))
	buf = append(buf, e)
	return buf[:i+1]
}

// group: bytes versions

// AppendSplitBytes splits p to buf using sep as separator.
func AppendSplitBytes(buf [][]byte, p, sep []byte, n int) [][]byte {
	if len(p) == 0 {
		return buf
	}
	var i int
	for {
		m := bytes.Index(p, sep)
		if m < 0 {
			break
		}
		buf = append(buf, p[:m:m])
		p = p[m+len(sep):]
		i++
		if n >= 0 && i >= n {
			break
		}
	}
	buf = append(buf, p)
	return buf[:i+1]
}

// AppendSplitEntryBytes splits p to buf using sep as separator.
func AppendSplitEntryBytes(buf []entry.Entry64, p, sep []byte, n int) []entry.Entry64 {
	if len(p) == 0 {
		return buf
	}
	var off int
	var i int
	for {
		m := bytes.Index(p, sep)
		if m < 0 {
			break
		}
		var e entry.Entry64
		e.Encode(uint32(off), uint32(off+m))
		buf = append(buf, e)
		p = p[m+len(sep):]
		off += m + len(sep)
		i++
		if n >= 0 && i >= n {
			break
		}
	}
	var e entry.Entry64
	e.Encode(uint32(off), uint32(off+len(p)))
	buf = append(buf, e)
	return buf[:i+1]
}

// group: string versions

// AppendSplitString splits s to buf using sep as separator.
func AppendSplitString(buf []string, s, sep string, n int) []string {
	if len(s) == 0 {
		return buf
	}
	var i int
	for {
		m := strings.Index(s, sep)
		if m < 0 {
			break
		}
		buf = append(buf, s[:m])
		s = s[m+len(sep):]
		i++
		if n >= 0 && i >= n {
			break
		}
	}
	buf = append(buf, s)
	return buf[:i+1]
}

// AppendSplitEntryString splits s to buf using sep as separator.
func AppendSplitEntryString(buf []entry.Entry64, s, sep string, n int) []entry.Entry64 {
	if len(s) == 0 {
		return buf
	}
	var off int
	var i int
	for {
		m := strings.Index(s, sep)
		if m < 0 {
			break
		}
		var e entry.Entry64
		e.Encode(uint32(off), uint32(off+m))
		buf = append(buf, e)
		s = s[m+len(sep):]
		off += m + len(sep)
		i++
		if n >= 0 && i >= n {
			break
		}
	}
	var e entry.Entry64
	e.Encode(uint32(off), uint32(off+len(s)))
	buf = append(buf, e)
	return buf[:i+1]
}
