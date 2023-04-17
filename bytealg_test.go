package bytealg

import (
	"bytes"
	"testing"

	"github.com/koykov/fastconv"
)

var (
	trimOrigin = []byte("..foo bar!!???")
	trimExpect = []byte("foo bar")
	trimCutStr = "?!."
	trimCut    = []byte(trimCutStr)

	splitOrigin = []byte("foo bar string")
	splitExpect = [][]byte{[]byte("foo"), []byte("bar"), []byte("string")}
	splitSep    = []byte(" ")

	idxAt     = []byte("some # string with # tokens")
	idxExpect = 19

	toUpper = []byte("FOOBAR")
	toLower = []byte("foobar")
	toTitle = []byte("FOOBAR")

	cpyOrigin = []byte("foobar")
	cpyExpect = []byte("foobar")
)

func TestBytealg(t *testing.T) {
	t.Run("trim", func(t *testing.T) {
		r := Trim(trimOrigin, trimCut)
		if !bytes.Equal(r, trimExpect) {
			t.Errorf(`Trim: mismatch result %s and expectation %s`, fastconv.B2S(r), fastconv.B2S(trimExpect))
		}
	})
	t.Run("append split", func(t *testing.T) {
		buf := make([][]byte, 0)
		buf = AppendSplit(buf, splitOrigin, splitSep, -1)
		if !EqualSet(buf, splitExpect) {
			t.Error("AppendSplit: mismatch result and expectation")
		}
	})
	t.Run("index at", func(t *testing.T) {
		r := IndexAt(idxAt, []byte("#"), 8)
		if r != idxExpect {
			t.Error("IndexAt: mismatch result and expectation")
		}
	})
	t.Run("to lower", func(t *testing.T) {
		cpy := Copy(toUpper)
		r := ToLower(cpy)
		if !bytes.Equal(r, toLower) {
			t.Error("ToLower: mismatch result and expectation")
		}
	})
	t.Run("to upper", func(t *testing.T) {
		cpy := Copy(toLower)
		r := ToUpper(cpy)
		if !bytes.Equal(r, toUpper) {
			t.Error("ToUpper: mismatch result and expectation")
		}
	})
	t.Run("to title", func(t *testing.T) {
		cpy := Copy(toLower)
		r := ToTitle(cpy)
		if !bytes.Equal(r, toTitle) {
			t.Error("ToTitle: mismatch result and expectation")
		}
	})
	t.Run("copy", func(t *testing.T) {
		r := Copy(cpyOrigin)
		if !bytes.Equal(r, cpyExpect) {
			t.Error("Copy: mismatch result and expectation")
		}
	})
}

func BenchmarkBytealg(b *testing.B) {
	b.Run("trim", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r := Trim(trimOrigin, trimCut)
			if !bytes.Equal(r, trimExpect) {
				b.Errorf(`Trim: mismatch result %s and expectation %s`, fastconv.B2S(r), fastconv.B2S(trimExpect))
			}
		}
	})
	b.Run("append split", func(b *testing.B) {
		b.ReportAllocs()
		buf := make([][]byte, 0)
		for i := 0; i < b.N; i++ {
			buf = buf[:0]
			buf = AppendSplit(buf, splitOrigin, splitSep, -1)
			if !EqualSet(buf, splitExpect) {
				b.Error("AppendSplit: mismatch result and expectation")
			}
		}
	})
	b.Run("index at", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r := IndexAt(idxAt, []byte("#"), 8)
			if r != idxExpect {
				b.Error("IndexAt: mismatch result and expectation")
			}
		}
	})
	b.Run("index byte at (lur)", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r := IndexByteAtLUR(idxAt, '#', 8)
			if r != idxExpect {
				b.Error("IndexByteAtLUR: mismatch result and expectation")
			}
		}
	})
	b.Run("to lower", func(b *testing.B) {
		b.ReportAllocs()
		buf := make([]byte, 0, len(cpyOrigin))
		for i := 0; i < b.N; i++ {
			buf = append(buf[:0], cpyOrigin...)
			r := ToLower(buf)
			if !bytes.Equal(r, cpyExpect) {
				b.Error("ToLower: mismatch result and expectation")
			}
		}
	})
	b.Run("to upper", func(b *testing.B) {
		b.ReportAllocs()
		buf := make([]byte, 0, len(toLower))
		for i := 0; i < b.N; i++ {
			buf = append(buf[:0], toLower...)
			r := ToUpper(buf)
			if !bytes.Equal(r, toUpper) {
				b.Error("ToUpper: mismatch result and expectation")
			}
		}
	})
	b.Run("to title", func(b *testing.B) {
		b.ReportAllocs()
		buf := make([]byte, 0, len(toLower))
		for i := 0; i < b.N; i++ {
			buf = append(buf[:0], toLower...)
			r := ToTitle(buf)
			if !bytes.Equal(r, toTitle) {
				b.Error("ToTitle: mismatch result and expectation")
			}
		}
	})
	b.Run("copy", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r := Copy(cpyOrigin)
			if !bytes.Equal(r, cpyExpect) {
				b.Error("Copy: mismatch result and expectation")
			}
		}
	})
}
