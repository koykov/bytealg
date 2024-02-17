package bytealg

import (
	"bytes"
	"testing"

	"github.com/koykov/entry"
)

var (
	splitOrigin = []byte("foo bar string")
	splitExpect = [][]byte{[]byte("foo"), []byte("bar"), []byte("string")}
	splitSep    = []byte(" ")

	toUpper = []byte("FOOBAR")
	toLower = []byte("foobar")
	toTitle = []byte("FOOBAR")

	cpyOrigin = []byte("foobar")
	cpyExpect = []byte("foobar")
)

func TestBytealg(t *testing.T) {
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
	b.Run("append split entry", func(b *testing.B) {
		b.ReportAllocs()
		buf := make([]entry.Entry64, 0)
		for i := 0; i < b.N; i++ {
			buf = buf[:0]
			buf = AppendSplitEntry(buf, splitOrigin, splitSep, -1)
		}
		_ = buf
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
