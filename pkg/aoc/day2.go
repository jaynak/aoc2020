package aoc

import (
	"regexp"
	"strconv"

	"github.com/jaynak/aoc2020/pkg/util"
)

func Day2(path string) (int, int) {
	lines := util.ReadToStrings(path)

	r := regexp.MustCompile(`^([0-9]+)-([0-9]+) ([a-z]{1}): ([a-z]+)$`)

	retVal := 0
	retValPartTwo := 0

	for _, x := range lines {

		matches := r.FindStringSubmatch(x)

		min, _ := strconv.Atoi(matches[1])
		max, _ := strconv.Atoi(matches[2])
		char := rune(matches[3][0])

		count := 0
		for _, r := range matches[4] {
			if r == char {
				count++
			}
		}

		if count >= min && count <= max {
			retVal++
		}

		// Second condition
		if len(matches[4]) < max {
			continue
		}

		a := matches[4][min-1] == byte(char)
		b := matches[4][max-1] == byte(char)

		if a != b {
			retValPartTwo++
		}
	}

	return retVal, retValPartTwo
}
