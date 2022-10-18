package bytealg

// Loop rolling versions of several functions.

var _ = HasByteLR

// HasByteLR checks if p contains b.
//
// This function designed to use with the largest input.
func HasByteLR(p []byte, b byte) bool {
	s := p
	for len(s) >= 8 {
		if s[0] == b {
			return true
		}
		if s[1] == b {
			return true
		}
		if s[2] == b {
			return true
		}
		if s[3] == b {
			return true
		}
		if s[4] == b {
			return true
		}
		if s[5] == b {
			return true
		}
		if s[6] == b {
			return true
		}
		if s[7] == b {
			return true
		}
		s = s[8:]
	}
	for len(s) >= 4 {
		if s[0] == b {
			return true
		}
		if s[1] == b {
			return true
		}
		if s[2] == b {
			return true
		}
		if s[3] == b {
			return true
		}
		s = s[4:]
	}
	for len(s) >= 2 {
		if s[0] == b {
			return true
		}
		if s[1] == b {
			return true
		}
		s = s[2:]
	}
	if len(s) == 1 {
		if s[0] == b {
			return true
		}
	}
	return false
}

// IndexByteAtLR is a loop rolling version of IndexAt().
func IndexByteAtLR(p []byte, b byte, at int) int {
	if at < 0 || at >= len(p) {
		return -1
	}

	n := 0
	s := p[at:]
	for len(s) >= 8 {
		if s[0] == b {
			return at + n
		}
		if s[1] == b {
			return at + n + 1
		}
		if s[2] == b {
			return at + n + 2
		}
		if s[3] == b {
			return at + n + 3
		}
		if s[4] == b {
			return at + n + 4
		}
		if s[5] == b {
			return at + n + 5
		}
		if s[6] == b {
			return at + n + 6
		}
		if s[7] == b {
			return at + n + 7
		}
		s = s[8:]
		n += 8
	}
	for len(s) >= 4 {
		if s[0] == b {
			return at + n
		}
		if s[1] == b {
			return at + n + 1
		}
		if s[2] == b {
			return at + n + 2
		}
		if s[3] == b {
			return at + n + 3
		}
		s = s[4:]
		n += 4
	}
	for len(s) >= 2 {
		if s[0] == b {
			return at + n
		}
		if s[1] == b {
			return at + n + 1
		}
		s = s[2:]
		n += 2
	}
	if len(s) > 0 && s[0] == b {
		return at + n
	}
	return -1
}
