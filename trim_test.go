package bytealg

import (
	"bytes"
	"testing"

	"github.com/koykov/byteconv"
)

var (
	trimOrigin = []byte("..foo bar!!???")
	trimExpect = []byte("foo bar")
	trimCutStr = "?!."
	trimCut    = []byte(trimCutStr)
)

func TestTrim(t *testing.T) {
	t.Run("left right", func(t *testing.T) {
		r := Trim(trimOrigin, trimCut)
		if !bytes.Equal(r, trimExpect) {
			t.Errorf(`Trim: mismatch result %s and expectation %s`, byteconv.B2S(r), byteconv.B2S(trimExpect))
		}
	})
}

func BenchmarkTrim(b *testing.B) {
	b.Run("generic", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r := Trim(trimOrigin, trimCut)
			_ = r
		}
	})
	b.Run("bytes", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r := TrimBytes(trimOrigin, trimCut)
			_ = r
		}
	})
	b.Run("string", func(b *testing.B) {
		b.ReportAllocs()
		so := string(trimOrigin)
		for i := 0; i < b.N; i++ {
			r := TrimString(so, trimCutStr)
			_ = r
		}
	})
}
