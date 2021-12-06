package aoc

import "testing"

func TestIsValid(t *testing.T) {
	b := IsValid(62, []int{20, 15, 25, 47, 40})

	if !b {
		t.Fail()
	}
}

func TestFindRange(t *testing.T) {
	b := FindRangeSum(127, []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576})

	if b != 62 {
		t.Fail()
	}
}
