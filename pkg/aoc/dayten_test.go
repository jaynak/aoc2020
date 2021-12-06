package aoc

import "testing"

func TestDay10(t *testing.T) {
	x, _ := Day10("../../data/10-test.txt")

	if x != 35 {
		t.Fatalf("Expected 35, got %v", x)
	}
}
