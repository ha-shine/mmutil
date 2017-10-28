package mmutil

import (
	"testing"
)

func testIsZawgyi(t *testing.T) {
	testTables := []struct {
		x string
		y bool
	}{
		{"ဟုတ်ပြီလေ", true},
		{"ဟုတ်ြပီေလ", false},
	}
	for _, test := range testTables {
		got := StringIsZawgyi(test.x)
		if test.y != got {
			t.Errorf("Expected %v, Got %v.", test.y, got)
		}
	}
}
