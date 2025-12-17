package unit

import "testing"

func TestPlaceholder(t *testing.T) {
	if 1 != 1 {
		t.Fatal("math is broken")
	}
}
