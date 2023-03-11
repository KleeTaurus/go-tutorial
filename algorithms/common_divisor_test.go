package algorithms

import "testing"

func TestCommonDivisor(t *testing.T) {
	want := 25
	got := FindCommonDivisor(25, 50)
	if got != want {
		t.Errorf("want: %v, got: %v\n", want, got)
	}

	want = 80
	got = FindCommonDivisor(1680, 640)
	if got != want {
		t.Errorf("want: %v, got: %v\n", want, got)
	}
}
