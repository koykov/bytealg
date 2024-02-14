package bytealg

import (
	"bytes"

	"github.com/koykov/byteseq"
)

const (
	// Trim directions.
	trimBoth  = 0
	trimLeft  = 1
	trimRight = 2
)

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
