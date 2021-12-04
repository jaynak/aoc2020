package aoc

import (
	"github.com/jaynak/aoc2020/pkg/util"
)

func Day3(path string) (int, int) {

	lines := util.ReadToStrings(path)

	// trees := CountTrees(lines, 3, 1)

	slopes := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	trees := []int{}
	product := 0

	for _, slope := range slopes {
		trees = append(trees, CountTrees(lines, slope[0], slope[1]))
	}

	for _, tree := range trees {
		if product == 0 {
			product = tree
		} else {
			product = product * tree
		}
	}

	return trees[1], product
}

func CountTrees(lines []string, right int, down int) int {

	trees := 0
	lastcol := -right
	wrap := len(lines[0])

	for i := 0; i < len(lines); i = i + down {

		lastcol = (lastcol + right) % wrap
		if lines[i][lastcol] == '#' {
			trees++
		}
	}

	return trees
}
