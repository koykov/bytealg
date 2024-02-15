package bytealg

const (
	// Trim directions.
	trimBoth  = 0
	trimLeft  = 1
	trimRight = 2
)

// group: bytes versions

// Trim makes fast and alloc-free trim.
func Trim(p, cut []byte) []byte {
	return trim(p, cut, trimBoth)
}

func TrimLeft(p, cut []byte) []byte {
	return trim(p, cut, trimLeft)
}

func TrimRight(p, cut []byte) []byte {
	return trim(p, cut, trimRight)
}

// Generic trim.
//
// Just calculates trim edges and return sub-slice.
func trim(p, cut []byte, dir int) []byte {
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

func TrimLeftString(p, cut string) string {
	return strim(p, cut, trimLeft)
}

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
