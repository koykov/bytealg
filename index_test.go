package bytealg

import (
	"fmt"
	"testing"
)

type tc struct {
	a string
	b string
	i int
}

var (
	indexTC = []tc{
		{"", "", 0},
		{"", "a", -1},
		{"", "foo", -1},
		{"fo", "foo", -1},
		{"foo", "baz", -1},
		{"foo", "foo", 0},
		{"oofofoofooo", "f", 2},
		{"oofofoofooo", "foo", 4},
		{"barfoobarfoo", "foo", 3},
		{"foo", "", 0},
		{"foo", "o", 1},
		{"abcABCabc", "A", 3},
		// cases with one byte strings - test IndexByte and special case in Index()
		{"", "a", -1},
		{"x", "a", -1},
		{"x", "x", 0},
		{"abc", "a", 0},
		{"abc", "b", 1},
		{"abc", "c", 2},
		{"abc", "x", -1},
		{"barfoobarfooyyyzzzyyyzzzyyyzzzyyyxxxzzzyyy", "x", 33},
		{"fofofofooofoboo", "oo", 7},
		{"fofofofofofoboo", "ob", 11},
		{"fofofofofofoboo", "boo", 12},
		{"fofofofofofoboo", "oboo", 11},
		{"fofofofofoooboo", "fooo", 8},
		{"fofofofofofoboo", "foboo", 10},
		{"fofofofofofoboo", "fofob", 8},
		{"fofofofofofofoffofoobarfoo", "foffof", 12},
		{"fofofofofoofofoffofoobarfoo", "foffof", 13},
		{"fofofofofofofoffofoobarfoo", "foffofo", 12},
		{"fofofofofoofofoffofoobarfoo", "foffofo", 13},
		{"fofofofofoofofoffofoobarfoo", "foffofoo", 13},
		{"fofofofofofofoffofoobarfoo", "foffofoo", 12},
		{"fofofofofoofofoffofoobarfoo", "foffofoob", 13},
		{"fofofofofofofoffofoobarfoo", "foffofoob", 12},
		{"fofofofofoofofoffofoobarfoo", "foffofooba", 13},
		{"fofofofofofofoffofoobarfoo", "foffofooba", 12},
		{"fofofofofoofofoffofoobarfoo", "foffofoobar", 13},
		{"fofofofofofofoffofoobarfoo", "foffofoobar", 12},
		{"fofofofofoofofoffofoobarfoo", "foffofoobarf", 13},
		{"fofofofofofofoffofoobarfoo", "foffofoobarf", 12},
		{"fofofofofoofofoffofoobarfoo", "foffofoobarfo", 13},
		{"fofofofofofofoffofoobarfoo", "foffofoobarfo", 12},
		{"fofofofofoofofoffofoobarfoo", "foffofoobarfoo", 13},
		{"fofofofofofofoffofoobarfoo", "foffofoobarfoo", 12},
		{"fofofofofoofofoffofoobarfoo", "ofoffofoobarfoo", 12},
		{"fofofofofofofoffofoobarfoo", "ofoffofoobarfoo", 11},
		{"fofofofofoofofoffofoobarfoo", "fofoffofoobarfoo", 11},
		{"fofofofofofofoffofoobarfoo", "fofoffofoobarfoo", 10},
		{"fofofofofoofofoffofoobarfoo", "foobars", -1},
		{"foofyfoobarfoobar", "y", 4},
		{"oooooooooooooooooooooo", "r", -1},
		{"oxoxoxoxoxoxoxoxoxoxoxoy", "oy", 22},
		{"oxoxoxoxoxoxoxoxoxoxoxox", "oy", -1},
		// test fallback to Rabin-Karp.
		{"000000000000000000000000000000000000000000000000000000000000000000000001", "0000000000000000000000000000000000000000000000000000000000000000001", 5},
	}

	idxAt     = []byte("some # string with # tokens")
	idxExpect = 19
)

func TestIndexAt(t *testing.T) {
	t.Run("generic", func(t *testing.T) {
		r := IndexAtBytes(idxAt, []byte("#"), 8)
		if r != idxExpect {
			t.Error("IndexAtBytes: mismatch result and expectation")
		}
	})
	t.Run("bytes", func(t *testing.T) {
		r := IndexAtBytes(idxAt, []byte("#"), 8)
		if r != idxExpect {
			t.Error("IndexAtBytes: mismatch result and expectation")
		}
	})
}

func BenchmarkIndexAt(b *testing.B) {
	sep := []byte("#")
	b.Run("generic", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r := IndexAt(idxAt, sep, 8)
			_ = r
		}
	})
	b.Run("bytes", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r := IndexAtBytes(idxAt, sep, 8)
			_ = r
		}
	})
}

func TestHasByte(t *testing.T) {
	for _, tc_ := range indexTC {
		if len(tc_.b) > 1 || len(tc_.b) == 0 {
			continue
		}
		p := []byte(tc_.a)
		t.Run(fmt.Sprintf("generic/%s/%s", tc_.a, tc_.b), func(t *testing.T) {
			r := HasByte(p, tc_.b[0])
			if (tc_.i == -1 && r) || (tc_.i >= 0 && !r) {
				t.FailNow()
			}
		})
		t.Run(fmt.Sprintf("bytes/%s/%s", tc_.a, tc_.b), func(t *testing.T) {
			r := HasByteBytes(p, tc_.b[0])
			if (tc_.i == -1 && r) || (tc_.i >= 0 && !r) {
				t.FailNow()
			}
		})
	}
}

func BenchmarkHasByte(b *testing.B) {
	for _, tc_ := range indexTC {
		if len(tc_.b) > 1 || len(tc_.b) == 0 {
			continue
		}
		p := []byte(tc_.a)
		b.Run(fmt.Sprintf("generic/%s/%s", tc_.a, tc_.b), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				r := HasByte(p, tc_.b[0])
				_ = r
			}
		})
		b.Run(fmt.Sprintf("bytes/%s/%s", tc_.a, tc_.b), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				r := HasByteBytes(p, tc_.b[0])
				_ = r
			}
		})
	}
}

func TestIndexByte(t *testing.T) {
	for _, tc_ := range indexTC {
		if len(tc_.b) > 1 || len(tc_.b) == 0 {
			continue
		}
		t.Run(fmt.Sprintf("generic/%s/%s", tc_.a, tc_.b), func(t *testing.T) {
			r := IndexByteAt([]byte(tc_.a), tc_.b[0], 0)
			if r != tc_.i {
				t.FailNow()
			}
		})
		t.Run(fmt.Sprintf("bytes/%s/%s", tc_.a, tc_.b), func(t *testing.T) {
			r := IndexByteAtBytes([]byte(tc_.a), tc_.b[0], 0)
			if r != tc_.i {
				t.FailNow()
			}
		})
	}
}

func BenchmarkIndexByte(b *testing.B) {
	for _, tc_ := range indexTC {
		if len(tc_.b) > 1 || len(tc_.b) == 0 {
			continue
		}
		p := []byte(tc_.a)
		b.Run(fmt.Sprintf("generic/%s/%s", tc_.a, tc_.b), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				r := IndexByteAt(p, tc_.b[0], 0)
				_ = r
			}
		})
		b.Run(fmt.Sprintf("bytes/%s/%s", tc_.a, tc_.b), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				r := IndexByteAtBytes(p, tc_.b[0], 0)
				_ = r
			}
		})
	}
}
