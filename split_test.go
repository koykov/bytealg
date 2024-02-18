package bytealg

import (
	"bytes"
	"testing"

	"github.com/koykov/entry"
)

func TestSplit(t *testing.T) {
	t.Run("generic/split", func(t *testing.T) {
		buf := make([][]byte, 0)
		buf = AppendSplit(buf, splitOrigin, splitSep, -1)
		if !EqualSet(buf, splitExpect) {
			t.Error("AppendSplit: mismatch result and expectation")
		}
	})
	t.Run("generic/split entry", func(t *testing.T) {
		buf := make([]entry.Entry64, 0)
		buf = AppendSplitEntry(buf, splitOrigin, splitSep, -1)
		for i := 0; i < len(buf); i++ {
			lo, hi := buf[i].Decode()
			if !bytes.Equal(splitOrigin[lo:hi], splitExpect[i]) {
				t.Error("AppendSplit: mismatch result and expectation")
				break
			}
		}
	})

	t.Run("bytes/split", func(t *testing.T) {
		buf := make([][]byte, 0)
		buf = AppendSplitBytes(buf, splitOrigin, splitSep, -1)
		if !EqualSet(buf, splitExpect) {
			t.Error("AppendSplit: mismatch result and expectation")
		}
	})
	t.Run("bytes/split entry", func(t *testing.T) {
		buf := make([]entry.Entry64, 0)
		buf = AppendSplitEntryBytes(buf, splitOrigin, splitSep, -1)
		for i := 0; i < len(buf); i++ {
			lo, hi := buf[i].Decode()
			if !bytes.Equal(splitOrigin[lo:hi], splitExpect[i]) {
				t.Error("AppendSplit: mismatch result and expectation")
				break
			}
		}
	})
}

func BenchmarkSplit(b *testing.B) {
	b.Run("generic/split", func(b *testing.B) {
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
	b.Run("generic/split entry", func(b *testing.B) {
		b.ReportAllocs()
		buf := make([]entry.Entry64, 0)
		for i := 0; i < b.N; i++ {
			buf = buf[:0]
			buf = AppendSplitEntry(buf, splitOrigin, splitSep, -1)
		}
		_ = buf
	})

	b.Run("bytes/split", func(b *testing.B) {
		b.ReportAllocs()
		buf := make([][]byte, 0)
		for i := 0; i < b.N; i++ {
			buf = buf[:0]
			buf = AppendSplitBytes(buf, splitOrigin, splitSep, -1)
			if !EqualSet(buf, splitExpect) {
				b.Error("AppendSplit: mismatch result and expectation")
			}
		}
	})
	b.Run("bytes/split entry", func(b *testing.B) {
		b.ReportAllocs()
		buf := make([]entry.Entry64, 0)
		for i := 0; i < b.N; i++ {
			buf = buf[:0]
			buf = AppendSplitEntryBytes(buf, splitOrigin, splitSep, -1)
		}
		_ = buf
	})
}
