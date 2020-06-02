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

	cpyOrigin = []byte("foobar")
	cpyExpect = []byte("foobar")
)

func TestTrim(t *testing.T) {
	r := Trim(trimOrigin, trimCut)
	if !bytes.Equal(r, trimExpect) {
		t.Errorf(`Trim: mismatch result %s and expectation %s`, fastconv.B2S(r), fastconv.B2S(trimExpect))
	}
}

func BenchmarkTrim(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := Trim(trimOrigin, trimCut)
		if !bytes.Equal(r, trimExpect) {
			b.Errorf(`Trim: mismatch result %s and expectation %s`, fastconv.B2S(r), fastconv.B2S(trimExpect))
		}
	}
}

func BenchmarkTrim_Native(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := bytes.Trim(trimOrigin, trimCutStr)
		if !bytes.Equal(r, trimExpect) {
			b.Errorf(`Trim: mismatch result %s and expectation %s`, fastconv.B2S(r), fastconv.B2S(trimExpect))
		}
	}
}

func TestAppendSplit(t *testing.T) {
	buf := make([][]byte, 0)
	buf = AppendSplit(buf, splitOrigin, splitSep, -1)
	if !EqualSet(buf, splitExpect) {
		t.Error("AppendSplit: mismatch result and expectation")
	}
}

func BenchmarkAppendSplit(b *testing.B) {
	buf := make([][]byte, 0)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		buf = buf[:0]
		buf = AppendSplit(buf, splitOrigin, splitSep, -1)
		if !EqualSet(buf, splitExpect) {
			b.Error("AppendSplit: mismatch result and expectation")
		}
	}
}

func TestIndexAt(t *testing.T) {
	r := IndexAt(idxAt, []byte("#"), 8)
	if r != idxExpect {
		t.Error("IndexAt: mismatch result and expectation")
	}
}

func TestToLower(t *testing.T) {
	r := ToLower(toUpper)
	if !bytes.Equal(r, toLower) {
		t.Error("ToLower: mismatch result and expectation")
	}
}

func TestToUpper(t *testing.T) {
	r := ToUpper(toLower)
	if !bytes.Equal(r, toUpper) {
		t.Error("ToUpper: mismatch result and expectation")
	}
}

func TestCopy(t *testing.T) {
	r := Copy(cpyOrigin)
	if !bytes.Equal(r, cpyExpect) {
		t.Error("Copy: mismatch result and expectation")
	}
}

func BenchmarkCopy(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := Copy(cpyOrigin)
		if !bytes.Equal(r, cpyExpect) {
			b.Error("Copy: mismatch result and expectation")
		}
	}
}
