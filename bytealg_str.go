package bytealg

import (
	"strings"

	fc "github.com/koykov/fastconv"
)

// Alloc-free string trim.
func TrimStr(p, cut string) string {
	return fc.B2S(trim(fc.S2B(p), fc.S2B(cut), trimBoth))
}

// String left trim.
func TrimLeftStr(p, cut string) string {
	return fc.B2S(trim(fc.S2B(p), fc.S2B(cut), trimLeft))
}

// String right trim.
func TrimRightStr(p, cut string) string {
	return fc.B2S(trim(fc.S2B(p), fc.S2B(cut), trimRight))
}

func AppendSplitStr(buf []string, s, sep string, n int) []string {
	if n < 0 {
		n = strings.Count(s, sep) + 1
	}

	n--
	i := 0
	for i < n {
		m := strings.Index(s, sep)
		if m < 0 {
			break
		}
		buf = append(buf, s[:m])
		s = s[m+len(sep):]
		i++
	}
	buf = append(buf, s)
	return buf[:i+1]
}

// IndexAtStr is equal to strings.Index() but doesn't consider occurrences of sep in s[:at].
func IndexAtStr(s, sep string, at int) int {
	return IndexAt(fc.S2B(s), fc.S2B(sep), at)
}

func CopyStr(s string) string {
	return fc.B2S(append([]byte(nil), s...))
}
