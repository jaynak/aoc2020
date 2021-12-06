package aoc

import (
	"github.com/jaynak/aoc2020/pkg/util"
)

func Day11(path string) (int, int) {
	lines := util.ReadToStrings(path)

	cols := len(lines[0])
	rows := len(lines)

	rm := make(room, rows)

	for i := range rm {
		rm[i] = make([]rune, cols)
	}

	for i, line := range lines {
		for j, r := range line {
			rm[i][j] = r
		}
	}

	changed := true
	for changed {
		next, changes := rm.RunGen()
		rm = next

		if changes == 0 {
			changed = false
		}
	}

	return rm.CountOccupied(), 0

}

type room [][]rune

func (rm room) CountOccupied() int {
	count := 0

	for _, row := range rm {
		for _, col := range row {
			if col == '#' {
				count++
			}
		}
	}

	return count
}

func (rm room) VisibleNeighbours(row int, col int) int {

	// Brute force!
	count := 0

	// Build a structure of arrays

	for i := row - 1; i >= 0; i-- {
		if rm[i][col] == '#' {
			count++
		}

		if rm[i][col] != '.' {
			break
		}
	}
}

func (rm room) OccupiedNeighbours(row int, col int) int {
	count := 0

	top, left, bottom, right := row-1, col-1, row+1, col+1
	if top < 0 {
		top = 0
	}
	if left < 0 {
		left = 0
	}
	if bottom >= len(rm) {
		bottom = len(rm) - 1
	}
	if right >= len(rm[0]) {
		right = len(rm[0]) - 1
	}

	for r := top; r <= bottom; r++ {
		for c := left; c <= right; c++ {
			if c == col && r == row {
				continue
			}

			if rm[r][c] == '#' {
				count++
			}
		}
	}

	return count

}

func DeepCopyRoom(s [][]rune) [][]rune {
	d := make([][]rune, len(s))
	copy(d, s)
	return d
}

func (rm room) RunGen() (room, int) {

	changes := 0

	next := make(room, len(rm))

	for i := range next {
		next[i] = make([]rune, len(rm[0]))
	}

	for i, row := range rm {
		for j, seat := range row {
			occ := rm.OccupiedNeighbours(i, j)
			if seat == 'L' && occ == 0 {
				next[i][j] = '#'
				changes++
			} else if seat == '#' && occ >= 4 {
				next[i][j] = 'L'
				changes++
			} else {
				next[i][j] = rm[i][j]
			}
		}
	}

	return next, changes
}
