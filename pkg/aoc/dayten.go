package aoc

import (
	"sort"

	"github.com/jaynak/aoc2020/pkg/util"
)

func Day10(path string) (int, int) {

	lines := util.ReadToInts(path)

	sort.Ints(lines)

	diffs := make([]int, 3)
	diffs[2] = 1
	last := 0

	for _, v := range lines {
		diff := v - last
		if diff > 3 {
			break
		}

		diffs[v-last-1]++
		last = v
	}

	paths := Day10DP(lines)

	return diffs[0] * diffs[2], paths
}

func Day10DP(chargers []int) int {

	// fmt.Println(chargers)

	revChargers := chargers
	for i, j := 0, len(revChargers)-1; i < j; i, j = i+1, j-1 {
		revChargers[i], revChargers[j] = revChargers[j], revChargers[i]
	}

	mem := make([]int, revChargers[0]+3)

	for i, v := range revChargers {
		if i == 0 {
			mem[v] = 1
		} else {
			mem[v] = mem[v+1] + mem[v+2] + mem[v+3]
		}
	}

	return mem[1] + mem[2] + mem[3]
}
