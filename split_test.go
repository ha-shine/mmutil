package mmutil

import (
	"strings"
	"testing"
)

func TestSplitingWords(t *testing.T) {
	testTable := []struct{ x, y string }{
		{"ကိုကြီး", "ကို ကြီး"},
		{"သိက္ခာ", "သိက် ခာ"},
		{"မြက်ကန်း", "မြက် ကန်း"},
	}
	for _, test := range testTable {
		got := strings.Join(SplitWords(test.x), " ")
		if test.y != got {
			t.Errorf("Expected %+q, Got %+q.", test.y, got)
		}
	}
}
