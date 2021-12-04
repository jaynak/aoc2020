package aoc

import (
	"sort"

	"github.com/jaynak/aoc2020/pkg/util"
)

func Day5(path string) (int, int) {

	passes := util.ReadToStrings(path)
	seats := []int{}

	max := -1
	for _, pass := range passes {
		thisID := CalcSeatID(pass)
		seats = append(seats, thisID)
		if thisID > max {
			max = thisID
		}
	}

	sort.Ints(seats)

	offset := seats[0]
	mySeat := -1

	for i, v := range seats {
		if i+offset != v {
			mySeat = i + offset
			break
		}
	}

	return max, mySeat
}

func CalcSeatID(bp string) int {
	row := BinaryFind(bp, 127, 0)
	col := BinaryFind(bp, 7, 7)

	return row*8 + col
}

func BinaryFind(bp string, max int, startAt int) int {
	low := 0
	high := max

	idx := startAt

	for low != high {
		mid := (high - low + 1) / 2
		switch bp[idx] {
		case 'F', 'L':
			high -= mid
		case 'B', 'R':
			low += mid
		default:
			return -1
		}

		idx++
	}

	// low = high at this point
	return low
}
