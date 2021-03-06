package mmutil

import (
	"testing"
)

func TestIsZawgyi(t *testing.T) {
	testTables := []struct {
		x string
		y bool
	}{
		{"ဟုတ်ပြီလေ", false},
		{"ဟုတ်ြပီေလ", true},
	}
	for _, test := range testTables {
		got := StringIsZawgyi(test.x)
		if test.y != got {
			t.Errorf("Expected %v, Got %v.", test.y, got)
		}
	}
}
