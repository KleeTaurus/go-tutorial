package main

import "testing"

func TestFactorial(t *testing.T) {
	want := 120

	got := factorial(5)
	if got != want {
		t.Errorf("want: %v, got: %v\n", want, got)
	}
}
