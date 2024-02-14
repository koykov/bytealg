package bytealg

import (
	"bytes"

	"github.com/koykov/byteseq"
)

const (
	bfmt4space = ' '
	bfmt4tab   = '\t'
	bfmt4nl    = '\n'
	bfmt4cr    = '\r'
)

// Default formatting bytes: space, tab, new line and caret return.
var bfmt4 = []byte{bfmt4space, bfmt4tab, bfmt4nl, bfmt4cr}

// group: generic versions

// TrimFmt4 removes default formatting bytes from both side of p.
func TrimFmt4[T byteseq.Byteseq](p T) T {
	return trimFmt4(p, trimBoth)
}

// TrimLeftFmt4 removes default formatting bytes from left size of p.
func TrimLeftFmt4[T byteseq.Byteseq](p T) T {
	return trimFmt4(p, trimLeft)
}

// TrimRightFmt4 removes default formatting bytes from right size of p.
func TrimRightFmt4[T byteseq.Byteseq](p T) T {
	return trimFmt4(p, trimRight)
}

// Generic trimFmt4.
func trimFmt4[T byteseq.Byteseq](p T, dir int) T {
	l, r := 0, len(p)-1
	pb := byteseq.Q2B(p)
	if dir == trimBoth || dir == trimLeft {
		for ; l < len(pb); l++ {
			if !bytes.Contains(bfmt4, []byte{pb[l]}) {
				break
			}
		}
	}
	if dir == trimBoth || dir == trimRight {
		for ; r >= l; r-- {
			if !bytes.Contains(bfmt4, []byte{pb[r]}) {
				break
			}
		}
	}
	return byteseq.B2Q[T](pb[l : r+1])
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
