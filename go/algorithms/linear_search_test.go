package algorithms

import "testing"

func TestLinearFind(t *testing.T) {
	if !LinearFind([]int{10, 2, 4}, 2) {
		t.Error("expected to find 2")
	}
	if LinearFind([]string{"val", "anotherval"}, "somestring") {
		t.Error("expected not to find \"somestring\"")
	}
}

func TestIndexOf(t *testing.T) {
	if IndexOf([]int{10, 100, 50}, 50) != 2 {
		t.Error("expected index of 50 to be 2")
	}
	if IndexOf([]int{1, 2, 3}, 40) != -1 {
		t.Error("expected index of 40 to be -1, since it is not in the array")
	}
}
