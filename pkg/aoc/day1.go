package aoc

import (
	"github.com/jaynak/aoc2020/pkg/util"
)

func Day1(path string) (int, int) {

	entries := util.ReadToInts(path)

	two := 0
	m := make(map[int]bool)

	for _, v := range entries {
		if _, ok := m[v]; ok {
			two = v * (2020 - v)
			break
		}

		m[2020-v] = true
	}

	three := 0
	double := make(map[int]map[int]int)

	// build a double grid
	for i, v := range entries {
		double[i] = make(map[int]int)

		for j, v2 := range entries {
			double[i][j] = v + v2
		}
	}

	for i, v := range entries {
		for j := range entries {
			if i == j {
				continue
			}

			for k := range entries {
				if i == k {
					continue
				}

				if double[j][k] == 2020-v {
					// fmt.Printf("%v, %v, %v\n", entries[i], entries[j], entries[k])
					three = v * entries[j] * entries[k]
					break
				}

			}

			if three != 0 {
				break
			}
		}

		if three != 0 {
			break
		}
	}

	// return the results
	return two, three
}
