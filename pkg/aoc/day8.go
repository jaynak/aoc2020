package aoc

import (
	"strconv"
	"strings"

	"github.com/jaynak/aoc2020/pkg/util"
)

func Day8(path string) (int, int) {

	lines := util.ReadToStrings(path)
	program := ParseInstructions(lines)

	unmod, _ := ExecuteProgram(program)
	mod := TestProgram(program)

	return unmod, mod
}

func TestProgram(program []*instruction) int {

	for _, inst := range program {
		switch inst.code {
		case "nop":
			inst.code = "jmp"
			mod, comp := ExecuteProgram(program)
			if comp {
				return mod
			} else {
				inst.code = "nop"
			}
		case "jmp":
			inst.code = "nop"
			mod, comp := ExecuteProgram(program)
			if comp {
				return mod
			} else {
				inst.code = "jmp"
			}
		}
	}

	return -1
}

func ExecuteProgram(program []*instruction) (int, bool) {

	acc := 0
	seen := make(map[int]bool)

	loop := false
	offset := 0
	size := len(program)

	for !loop && offset != size {
		if offset < 0 || offset > size {
			return -1, false
		}
		seen[offset] = true

		switch program[offset].code {
		case "nop":
			offset++
		case "acc":
			acc += program[offset].mod
			offset++
		case "jmp":
			offset += program[offset].mod
		}

		if _, ok := seen[offset]; ok {
			loop = true
		}
	}

	return acc, offset == size
}

type instruction struct {
	code string
	mod  int
}

func ParseInstructions(lines []string) []*instruction {
	instructions := []*instruction{}

	for _, v := range lines {
		split := strings.Fields(v)
		mod, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}

		i := &instruction{
			code: split[0],
			mod:  mod,
		}

		instructions = append(instructions, i)
	}

	return instructions
}
