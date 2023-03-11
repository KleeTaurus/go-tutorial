package algorithms

import "testing"

func TestBinarySearch(t *testing.T) {
	want := 1
	got := BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}, 2)
	if got != want {
		t.Errorf("want %v, got %v", want, got)
	}
}
