package bytealg

const (
	bfmt4space = ' '
	bfmt4tab   = '\t'
	bfmt4nl    = '\n'
	bfmt4cr    = '\r'
)

// group: bytes versions

// TrimFmt4 removes default formatting bytes from both side of p.
func TrimFmt4(p []byte) []byte {
	return btrimFmt4(p, trimBoth)
}

// TrimLeftFmt4 removes default formatting bytes from left size of p.
func TrimLeftFmt4(p []byte) []byte {
	return btrimFmt4(p, trimLeft)
}

// TrimRightFmt4 removes default formatting bytes from right size of p.
func TrimRightFmt4(p []byte) []byte {
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
