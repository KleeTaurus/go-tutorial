package algorithms

import "testing"

func TestRecursionSum(t *testing.T) {
	want := 12
	got := RecursionSum([]int{2, 4, 6})

	if got != want {
		t.Errorf("want: %v, got: %v\n", want, got)
	}
}
