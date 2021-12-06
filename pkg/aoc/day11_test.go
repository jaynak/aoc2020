package aoc

import "testing"

func TestDay11(t *testing.T) {
	a, _ := Day11("../../data/11-test.txt")

	if a != 37 {
		t.Fatalf("Expecting 37 got %v", a)
	}
}
