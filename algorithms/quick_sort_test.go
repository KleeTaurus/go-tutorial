package main

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	want := []int{1, 2, 3, 4, 5}
	got := quickSort([]int{3, 5, 2, 1, 4})
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v, got: %v\n", want, got)
	}
}
