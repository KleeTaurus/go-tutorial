package main

import (
	"reflect"
	"testing"
)

func TestConvert(t *testing.T) {
	want := [][]int{
		[]int{4, 6, 14, 16, 20},
		[]int{6, 9, 21, 24, 30},
		[]int{14, 21, 49, 56, 70},
		[]int{16, 24, 56, 64, 80},
		[]int{20, 30, 70, 80, 100},
	}

	got := convert([]int{2, 3, 7, 8, 10})
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v, got: %v\n", want, got)
	}
}

func BenchmarkConvert(b *testing.B) {
	for n := 0; n < b.N; n++ {
		convert([]int{2, 3, 7, 8, 10})
	}
}
