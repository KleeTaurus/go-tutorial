package algorithms

import "testing"

func TestFindMax(t *testing.T) {
	want := 21
	got := FindMax([]int{2, 8, 3, 6, 15, 18, 10, 11, 15, 21, 16})

	if got != want {
		t.Errorf("want: %v, got: %v\n", want, got)
	}
}
