package main

import (
	"fmt"
	"strings"

	utils "github.com/baspar/adventofcode2022/internal"
	"github.com/baspar/adventofcode2022/internal/math"
)

type Instruction struct {
	op       string
	duration int
	value    int
}

type DayImpl struct {
	instructions []Instruction
}

func RunCPU(instructions []Instruction, callback func(cycle int, register int)) {
	cycle, register := 1, 1

	for _, instruction := range instructions {
		execInstructionAtCycle := cycle + instruction.duration
		for {
			callback(cycle, register)

			cycle++

			if execInstructionAtCycle == cycle {
				switch instruction.op {
				case "addx":
					register += instruction.value
				}
				break
			}
		}
	}
}

func (d *DayImpl) Init(lines []string) error {
	for _, line := range lines {
		var instruction Instruction
		if _, err := fmt.Sscanf(line, "addx %d", &instruction.value); err == nil {
			instruction.op = "addx"
			instruction.duration = 2
		} else {
			instruction.op = "noop"
			instruction.duration = 1
		}
		d.instructions = append(d.instructions, instruction)
	}
	return nil
}
func (d *DayImpl) Part1() (string, error) {
	strength := 0

	RunCPU(d.instructions, func(cycle, register int) {
		if (cycle-20)%40 == 0 {
			strength += cycle * register
		}
	})

	return fmt.Sprint(strength), nil
}
func (d *DayImpl) Part2() (string, error) {
	screen := ""

	RunCPU(d.instructions, func(cycle, spriteCenter int) {
		pixel := (cycle - 1) % 40
		if math.Abs(pixel-spriteCenter) < 2 {
			screen += "▒"
		} else {
			screen += "."
		}

		if pixel == 39 {
			screen += "\n"
		}
	})

	return strings.TrimSuffix(screen, "\n"), nil
}

func main() {
	utils.Run(&DayImpl{}, false)
}
