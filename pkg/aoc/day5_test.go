package aoc

import "testing"

func TestDay5(t *testing.T) {
	a, b := Day5("../../data/5-test.txt")

	if a != 820 || b != 0 {
		t.Fail()
	}
}

func TestRows(t *testing.T) {
	row := BinaryFind("FBFBBFFRLR", 127, 0)
	col := BinaryFind("FBFBBFFRLR", 7, 7)

	if row != 44 {
		t.Fatalf("Expected 44 got %v", row)
	}

	if col != 5 {
		t.Fatalf("Expected 5 got %v", row)
	}
}

func TestSeats(t *testing.T) {
	seatid := CalcSeatID("FBFBBFFRLR")

	if seatid != 357 {
		t.Fatalf("Expected 357 got %v", seatid)
	}
}
