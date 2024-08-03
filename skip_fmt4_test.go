package bytealg

import (
	"bytes"
	"testing"
)

var (
	skipFmt4Origin = []byte(`{
                "medium": {
                  "w": 600,
                  "h": 450,
                  "resize": "fit"
                }



`)
	testFmt4Expect = []byte(`{"medium":{"w":600,"h":450,"resize":"fit"}`)
)

func TestSkipFmt4(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		r := testSkipFmt4(nil, skipFmt4Origin)
		if !bytes.Equal(r, testFmt4Expect) {
			t.FailNow()
		}
	})
}

func BenchmarkSkipFmt4(b *testing.B) {
	b.Run("0", func(b *testing.B) {
		b.ReportAllocs()
		var buf []byte
		for i := 0; i < b.N; i++ {
			buf = testSkipFmt4(buf[:0], skipFmt4Origin)
		}
	})
}

func testSkipFmt4(buf, src []byte) []byte {
	var offset, p int
	var eof bool
	for {
		offset, eof = SkipBytesFmt4(src, offset)
		if eof {
			break
		}
		if offset == p {
			offset++
			buf = append(buf, src[p:offset]...)
		}
		p = offset
	}
	return buf
}
