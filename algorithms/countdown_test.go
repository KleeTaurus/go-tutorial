package main

import (
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	want := []int{5, 4, 3, 2, 1}

	got := []int{}
	countdown(5, got)
	if reflect.DeepEqual(want, got) {
		t.Errorf("want: %v, got: %v", want, got)
	}
}
