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
	t.Run("both", func(t *testing.T) {
		r := Trim(trimOrigin, trimCut)
		if !bytes.Equal(r, trimExpect) {
			t.Errorf(`Trim: mismatch result %s and expectation %s`, byteconv.B2S(r), byteconv.B2S(trimExpect))
		}
	})
}

func BenchmarkTrim(b *testing.B) {
	b.Run("both", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r := Trim(trimOrigin, trimCut)
			if !bytes.Equal(r, trimExpect) {
				b.Errorf(`Trim: mismatch result %s and expectation %s`, byteconv.B2S(r), byteconv.B2S(trimExpect))
			}
		}
	})
}
