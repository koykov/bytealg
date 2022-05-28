package bytealg

import (
	"strings"

	fc "github.com/koykov/fastconv"
)

var (
	// Suppress go vet warnings.
	_, _, _, _, _ = TrimLeftStr, TrimRightStr, CopyStr, ToTitleStr, IndexAnyAtStr
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

// Split s to buf using sep as separator.
func AppendSplitStr(buf []string, s, sep string, n int) []string {
	if len(s) == 0 {
		return buf
	}
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

// IndexAnyAtStr is equal to strings.IndexAny() but doesn't consider occurrences of sep in s[:at].
func IndexAnyAtStr(s, sep string, at int) int {
	return IndexAnyAt(fc.S2B(s), fc.S2B(sep), at)
}

// ToUpperStr is an alloc-free replacement of bytes.ToUpperStr() function.
func ToUpperStr(s string) string { return fc.B2S(ToUpper(fc.S2B(s))) }

// ToLowerStr is an alloc-free replacement of bytes.ToLowerStr() function.
func ToLowerStr(s string) string { return fc.B2S(ToLower(fc.S2B(s))) }

// ToTitleStr is an alloc-free replacement of bytes.ToTitleStr() function.
func ToTitleStr(s string) string { return fc.B2S(ToTitle(fc.S2B(s))) }

// Make a copy of string.
func CopyStr(s string) string {
	return fc.B2S(append([]byte(nil), s...))
}
