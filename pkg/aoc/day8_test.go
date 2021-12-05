package aoc

import "testing"

func TestDay8(t *testing.T) {
	a, b := Day8("../../data/8-test.txt")

	if a != 5 {
		t.Fatalf("Expected 5 got %v\n", a)
	}

	if b != 8 {
		t.Fatalf("Expected 8 got %v\n", b)
	}
}
