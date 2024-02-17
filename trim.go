package bytealg

import (
	"unsafe"

	"github.com/koykov/byteseq"
)

const (
	// Trim directions.
	trimBoth  = 0
	trimLeft  = 1
	trimRight = 2
)

// group: generic versions

// Trim makes fast and alloc-free trim over bytes or string.
func Trim[T byteseq.Q](x, cut T) T {
	return trim(x, cut, trimBoth)
}

// TrimLeft is a left version of Trim.
func TrimLeft[T byteseq.Q](p, cut T) T {
	return trim(p, cut, trimLeft)
}

// TrimRight is a right version of Trim.
func TrimRight[T byteseq.Q](p, cut T) T {
	return trim(p, cut, trimRight)
}

func trim[T byteseq.Q](x, cut T, dir int) T {
	if p, ok := byteseq.ToBytes(x); ok {
		pc, _ := byteseq.ToBytes(cut)
		r := btrim(p, pc, dir)
		return *(*T)(unsafe.Pointer(&r))
	}
	if s, ok := byteseq.ToString(x); ok {
		sc, _ := byteseq.ToString(cut)
		r := strim(s, sc, dir)
		return *(*T)(unsafe.Pointer(&r))
	}
	return x
}

// group: bytes versions

// TrimBytes makes fast and alloc-free trim over bytes.
func TrimBytes(p, cut []byte) []byte {
	return btrim(p, cut, trimBoth)
}

// TrimBytesLeft is a left version of TrimBytes.
func TrimBytesLeft(p, cut []byte) []byte {
	return btrim(p, cut, trimLeft)
}

// TrimBytesRight is a right version of TrimBytes.
func TrimBytesRight(p, cut []byte) []byte {
	return btrim(p, cut, trimRight)
}

// Generic trim.
//
// Just calculates trim edges and return sub-slice.
func btrim(p, cut []byte, dir int) []byte {
	l, r := 0, len(p)-1
	if dir == trimBoth || dir == trimLeft {
		for ; l <= r; l++ {
			var brk bool
			for j := 0; j < len(cut); j++ {
				if brk = p[l] == cut[j]; brk {
					break
				}
			}
			if !brk {
				break
			}
		}
	}
	if dir == trimBoth || dir == trimRight {
		for ; r >= l; r-- {
			var brk bool
			for j := 0; j < len(cut); j++ {
				if brk = p[r] == cut[j]; brk {
					break
				}
			}
			if !brk {
				break
			}
		}
	}
	return p[l : r+1]
}

// group: string versions

// TrimString makes fast and alloc-free trim over string.
func TrimString(p, cut string) string {
	return strim(p, cut, trimBoth)
}

// TrimLeftString is a left version of TrimString.
func TrimLeftString(p, cut string) string {
	return strim(p, cut, trimLeft)
}

// TrimRightString is a right version of TrimString.
func TrimRightString(p, cut string) string {
	return strim(p, cut, trimRight)
}

func strim(p, cut string, dir int) string {
	l, r := 0, len(p)-1
	if dir == trimBoth || dir == trimLeft {
		for ; l <= r; l++ {
			var brk bool
			for j := 0; j < len(cut); j++ {
				if brk = p[l] == cut[j]; brk {
					break
				}
			}
			if !brk {
				break
			}
		}
	}
	if dir == trimBoth || dir == trimRight {
		for ; r >= l; r-- {
			var brk bool
			for j := 0; j < len(cut); j++ {
				if brk = p[r] == cut[j]; brk {
					break
				}
			}
			if !brk {
				break
			}
		}
	}
	return p[l : r+1]
}

var _, _, _ = TrimString, TrimLeftString, TrimRightString
