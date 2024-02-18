package bytealg

import (
	"unsafe"

	"github.com/koykov/byteseq"
)

const (
	bfmt4space = ' '
	bfmt4tab   = '\t'
	bfmt4nl    = '\n'
	bfmt4cr    = '\r'
)

// group: generic versions

// TrimFmt4 removes default formatting bytes from both side of x.
func TrimFmt4[T byteseq.Q](x T) T {
	return trimFmt4(x, trimBoth)
}

// TrimLeftFmt4 is a left version of TrimFmt4.
func TrimLeftFmt4[T byteseq.Q](x T) T {
	return trimFmt4(x, trimLeft)
}

// TrimRightFmt4 is a right version of TrimFmt4.
func TrimRightFmt4[T byteseq.Q](x T) T {
	return trimFmt4(x, trimRight)
}

// Generic trimFmt4.
func trimFmt4[T byteseq.Q](x T, dir int) T {
	if p, ok := byteseq.ToBytes(x); ok {
		r := btrimFmt4(p, dir)
		return *(*T)(unsafe.Pointer(&r))
	}
	if s, ok := byteseq.ToString(x); ok {
		r := strimFmt4(s, dir)
		return *(*T)(unsafe.Pointer(&r))
	}
	return x
}

// group: bytes versions

// TrimBytesFmt4 removes default formatting bytes from both side of p.
func TrimBytesFmt4(p []byte) []byte {
	return btrimFmt4(p, trimBoth)
}

// TrimLeftBytesFmt4 removes default formatting bytes from left size of p.
func TrimLeftBytesFmt4(p []byte) []byte {
	return btrimFmt4(p, trimLeft)
}

// TrimRightBytesFmt4 removes default formatting bytes from right size of p.
func TrimRightBytesFmt4(p []byte) []byte {
	return btrimFmt4(p, trimRight)
}

// Generic btrimFmt4.
func btrimFmt4(p []byte, dir int) []byte {
	l, r := 0, len(p)-1
	if dir == trimBoth || dir == trimLeft {
		for ; l < len(p); l++ {
			c := p[l]
			if c != bfmt4space && c != bfmt4tab && c != bfmt4nl && c != bfmt4cr {
				break
			}
		}
	}
	if dir == trimBoth || dir == trimRight {
		for ; r >= l; r-- {
			c := p[r]
			if c != bfmt4space && c != bfmt4tab && c != bfmt4nl && c != bfmt4cr {
				break
			}
		}
	}
	return p[l : r+1]
}

// group: string versions

// TrimStringFmt4 removes default formatting bytes from both side of p.
func TrimStringFmt4(p string) string {
	return strimFmt4(p, trimBoth)
}

// TrimLeftStringFmt4 removes default formatting bytes from left size of p.
func TrimLeftStringFmt4(p string) string {
	return strimFmt4(p, trimLeft)
}

// TrimRightStringFmt4 removes default formatting bytes from right size of p.
func TrimRightStringFmt4(p string) string {
	return strimFmt4(p, trimRight)
}

// Generic strimFmt4.
func strimFmt4(p string, dir int) string {
	l, r := 0, len(p)-1
	if dir == trimBoth || dir == trimLeft {
		for ; l < len(p); l++ {
			c := p[l]
			if c != bfmt4space && c != bfmt4tab && c != bfmt4nl && c != bfmt4cr {
				break
			}
		}
	}
	if dir == trimBoth || dir == trimRight {
		for ; r >= l; r-- {
			c := p[r]
			if c != bfmt4space && c != bfmt4tab && c != bfmt4nl && c != bfmt4cr {
				break
			}
		}
	}
	return p[l : r+1]
}

var _, _, _ = TrimStringFmt4, TrimLeftStringFmt4, TrimRightStringFmt4
