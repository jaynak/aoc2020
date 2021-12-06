package aoc

import (
	"sort"

	"github.com/jaynak/aoc2020/pkg/util"
)

// NOTE: Added the preamble length as the first line in the file

func Day9(path string) (int, int) {
	lines := util.ReadToInts(path)

	preamble := lines[0]
	noncompliant := -1

	for i := 1; i < len(lines); i++ {
		if !IsValid(lines[i+preamble], lines[i:i+preamble]) {
			noncompliant = lines[i+preamble]
			break
		}
	}

	return noncompliant, FindRangeSum(noncompliant, lines[1:])
}

func IsValid(target int, prior []int) bool {

	seen := make(map[int]bool)

	for _, v := range prior {
		if _, ok := seen[v]; ok {
			return true
		}
		seen[target-v] = true
	}

	return false
}

func FindRangeSum(target int, values []int) int {
	sum := values[0]
	low := 0
	high := 0

	for sum != target && high < len(values) {
		if sum < target {
			high++
			sum += values[high]
		} else {
			sum -= values[low]
			low++
		}
	}

	// Return the sum of high and low
	found := values[low:high]
	sort.Ints(found)

	return found[0] + found[len(found)-1]
}
