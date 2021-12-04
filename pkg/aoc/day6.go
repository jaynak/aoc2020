package aoc

import "github.com/jaynak/aoc2020/pkg/util"

func Day6(path string) (int, int) {

	lines := util.ReadToStrings(path)

	q := make(map[rune]int)
	total := 0
	total2 := 0
	people := 0

	for _, line := range lines {
		if line == "" {
			total += len(q)

			for _, r := range q {
				if r == people {
					total2++
				}
			}

			q = make(map[rune]int)
			people = 0
		} else {
			people++
			for _, r := range line {
				q[r]++
			}
		}
	}

	total += len(q)

	for _, r := range q {
		if r == people {
			total2++
		}
	}

	return total, total2
}
