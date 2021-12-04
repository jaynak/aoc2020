package aoc

import "testing"

func TestDay7(t *testing.T) {
	a, b := Day7("../../data/7-test.txt")

	if a != 4 {
		t.Fatalf("Expected 4 got %v\n", a)
	}

	if b != 32 {
		t.Fatalf("Expected 32 got %v\n", b)
	}
}
