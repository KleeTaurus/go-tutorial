package main

import "testing"

func TestCommonDivisor(t *testing.T) {
	want := 25
	got, ok := findCommonDivisor(25, 50)
	if !ok || got != want {
		t.Errorf("want: %v, got: %v\n", want, got)
	}

	want = 8
	got, ok = findCommonDivisor(168, 64)
	if !ok || got != want {
		t.Errorf("want: %v, got: %v\n", want, got)
	}
}
