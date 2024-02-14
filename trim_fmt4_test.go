package bytealg

import (
	"bytes"
	"testing"

	"github.com/koykov/byteconv"
)

var (
	trimOriginFmt4  = []byte("    \t\t\t\n\nfoobar and lorem ipsum \t\r   \n\t\t")
	trimExpectFmt4  = []byte("foobar and lorem ipsum")
	ltrimExpectFmt4 = []byte("foobar and lorem ipsum \t\r   \n\t\t")
	rtrimExpectFmt4 = []byte("    \t\t\t\n\nfoobar and lorem ipsum")
)

func TestTrimFmt4(t *testing.T) {
	t.Run("generic/trim", func(t *testing.T) {
		r := TrimFmt4(trimOriginFmt4)
		if !bytes.Equal(r, trimExpectFmt4) {
			t.Errorf(`Trim: mismatch result %s and expectation %s`, byteconv.B2S(r), byteconv.B2S(trimExpectFmt4))
		}
	})
	t.Run("generic/ltrim", func(t *testing.T) {
		r := TrimLeftFmt4(trimOriginFmt4)
		if !bytes.Equal(r, ltrimExpectFmt4) {
			t.Errorf(`Trim: mismatch result %s and expectation %s`, byteconv.B2S(r), byteconv.B2S(ltrimExpectFmt4))
		}
	})
	t.Run("generic/rtrim", func(t *testing.T) {
		r := TrimRightFmt4(trimOriginFmt4)
		if !bytes.Equal(r, rtrimExpectFmt4) {
			t.Errorf(`Trim: mismatch result %s and expectation %s`, byteconv.B2S(r), byteconv.B2S(rtrimExpectFmt4))
		}
	})
	t.Run("bytes/trim", func(t *testing.T) {
		r := TrimBytesFmt4(trimOriginFmt4)
		if !bytes.Equal(r, trimExpectFmt4) {
			t.Errorf(`Trim: mismatch result %s and expectation %s`, byteconv.B2S(r), byteconv.B2S(trimExpectFmt4))
		}
	})
	t.Run("bytes/ltrim", func(t *testing.T) {
		r := TrimLeftBytesFmt4(trimOriginFmt4)
		if !bytes.Equal(r, ltrimExpectFmt4) {
			t.Errorf(`Trim: mismatch result %s and expectation %s`, byteconv.B2S(r), byteconv.B2S(ltrimExpectFmt4))
		}
	})
	t.Run("bytes/rtrim", func(t *testing.T) {
		r := TrimRightBytesFmt4(trimOriginFmt4)
		if !bytes.Equal(r, rtrimExpectFmt4) {
			t.Errorf(`Trim: mismatch result %s and expectation %s`, byteconv.B2S(r), byteconv.B2S(rtrimExpectFmt4))
		}
	})
}

func BenchmarkTrimFmt4(b *testing.B) {
	b.Run("generic/trim", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r := TrimFmt4(trimOriginFmt4)
			if !bytes.Equal(r, trimExpectFmt4) {
				b.Errorf(`Trim: mismatch result %s and expectation %s`, byteconv.B2S(r), byteconv.B2S(trimExpectFmt4))
			}
		}
	})
	b.Run("generic/ltrim", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r := TrimLeftFmt4(trimOriginFmt4)
			if !bytes.Equal(r, ltrimExpectFmt4) {
				b.Errorf(`Trim: mismatch result %s and expectation %s`, byteconv.B2S(r), byteconv.B2S(ltrimExpectFmt4))
			}
		}
	})
	b.Run("generic/rtrim", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r := TrimRightFmt4(trimOriginFmt4)
			if !bytes.Equal(r, rtrimExpectFmt4) {
				b.Errorf(`Trim: mismatch result %s and expectation %s`, byteconv.B2S(r), byteconv.B2S(rtrimExpectFmt4))
			}
		}
	})
	b.Run("bytes/trim", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r := TrimBytesFmt4(trimOriginFmt4)
			if !bytes.Equal(r, trimExpectFmt4) {
				b.Errorf(`Trim: mismatch result %s and expectation %s`, byteconv.B2S(r), byteconv.B2S(trimExpectFmt4))
			}
		}
	})
	b.Run("bytes/ltrim", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r := TrimLeftBytesFmt4(trimOriginFmt4)
			if !bytes.Equal(r, ltrimExpectFmt4) {
				b.Errorf(`Trim: mismatch result %s and expectation %s`, byteconv.B2S(r), byteconv.B2S(ltrimExpectFmt4))
			}
		}
	})
	b.Run("bytes/rtrim", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r := TrimRightBytesFmt4(trimOriginFmt4)
			if !bytes.Equal(r, rtrimExpectFmt4) {
				b.Errorf(`Trim: mismatch result %s and expectation %s`, byteconv.B2S(r), byteconv.B2S(rtrimExpectFmt4))
			}
		}
	})
}
