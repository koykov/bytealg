package bytealg

import (
	"strings"
	"testing"

	"github.com/koykov/fastconv"
)

var (
	trimOriginS = "..foo bar!!???"
	trimExpectS = "foo bar"
	trimCutS    = "?!."

	idxAtStr = fastconv.B2S(idxAt)
)

func TestTrimStr(t *testing.T) {
	r := TrimStr(trimOriginS, trimCutS)
	if r != trimExpectS {
		t.Errorf(`Trim: mismatch result %s and expectation %s`, r, trimExpectS)
	}
}

func BenchmarkTrimStr(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := TrimStr(trimOriginS, trimCutS)
		if r != trimExpectS {
			b.Errorf(`Trim: mismatch result %s and expectation %s`, r, trimExpectS)
		}
	}
}

func BenchmarkTrimStr_Native(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := strings.Trim(trimOriginS, trimCutS)
		if r != trimExpectS {
			b.Errorf(`Trim: mismatch result %s and expectation %s`, r, trimExpectS)
		}
	}
}

func TestIndexAtStr(t *testing.T) {
	r := IndexAtStr(idxAtStr, "#", 8)
	if r != idxExpect {
		t.Error("IndexAtStr: mismatch result and expectation")
	}
}
