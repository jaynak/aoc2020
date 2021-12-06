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

	directions := 0
	next := make([]rune, 8)

	// 1 2 3
	// 0   4
	// 7 6 5

	for directions < 8 {
		r, c := row, col
		cont := true

		for cont == true {
			r, c, cont = rm.Next(r, c, directions)

			if !cont {
				// Don't need to continue - there's no more places to go
				next[directions] = '.'
			}

			if rm[r][c] == '.' {
				continue
			}

			next[directions] = rm[r][c]
		}

		directions++
	}

	count := 0
	for _, ru := range next {
		if ru == '#' {
			count++
		}
	}

	return count
}

func (rm room) Next(row int, col int, dir int) (int, int, bool) {

	// 1 2 3
	// 0   4
	// 7 6 5

	newrow := row
	newcol := col

	switch dir {
	case 0:
		newcol--
	case 1:
		newrow--
		newcol--
	case 2:
		newrow--
	case 3:
		newrow--
		newcol++
	case 4:
		newrow++
	case 5:
		newrow++
		newcol++
	case 6:
		newrow++
	case 7:
		newrow++
		newcol--
	}

	if newcol < 0 || newrow < 0 || newrow > len(rm) || newcol > len(rm[0]) {
		return row, col, false
	}

	return newrow, newcol, true
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

// func DeepCopyRoom(s [][]rune) [][]rune {
// 	d := make([][]rune, len(s))
// 	copy(d, s)
// 	return d
// }

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

func (rm room) RunGenSecondRuleset() (room, int) {

	changes := 0

	next := make(room, len(rm))

	for i := range next {
		next[i] = make([]rune, len(rm[0]))
	}

	for i, row := range rm {
		for j, seat := range row {
			occ := rm.VisibleNeighbours(i, j)
			if seat == 'L' && occ == 0 {
				next[i][j] = '#'
				changes++
			} else if seat == '#' && occ >= 5 {
				next[i][j] = 'L'
				changes++
			} else {
				next[i][j] = rm[i][j]
			}
		}
	}

	return next, changes
}
