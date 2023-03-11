package algorithms

import (
	"reflect"
	"testing"
)

func TestConvert(t *testing.T) {
	want := [][]int{
		{4, 6, 14, 16, 20},
		{6, 9, 21, 24, 30},
		{14, 21, 49, 56, 70},
		{16, 24, 56, 64, 80},
		{20, 30, 70, 80, 100},
	}

	got := Convert([]int{2, 3, 7, 8, 10})
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v, got: %v\n", want, got)
	}
}

func BenchmarkConvert(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Convert([]int{2, 3, 7, 8, 10})
	}
}
