package main

import "testing"

func TestCommonDivisor(t *testing.T) {
	want := 25
	got := findCommonDivisor(25, 50)
	if got != want {
		t.Errorf("want: %v, got: %v\n", want, got)
	}

	want = 80
	got = findCommonDivisor(1680, 640)
	if got != want {
		t.Errorf("want: %v, got: %v\n", want, got)
	}
}
