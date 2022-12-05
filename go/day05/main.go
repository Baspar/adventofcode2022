package main

import (
	"fmt"
	"github.com/gammazero/deque"
	"strings"

	utils "github.com/baspar/adventofcode2022/internal"
)

type Instruction struct {
	n    int
	from int
	to   int
}

type Stacks map[int]*deque.Deque[byte]

func formatStacks(stacks Stacks) (f string) {
	for slot := 1; slot <= len(stacks); slot++ {
		f += fmt.Sprintf("%c", stacks[slot].Front())
	}
	return
}

type DayImpl struct {
	stacks       Stacks
	instructions []Instruction
}

func (d *DayImpl) Init(lines []string) error {
	d.stacks = make(Stacks)
	d.instructions = nil

	var (
		lineNumber int
		instr      Instruction
		width      = (len(lines[0]) + 1) / 4
	)

	// Read stacks
	for lineNumber = 0; !strings.HasPrefix(lines[lineNumber], " 1 "); lineNumber++ {
		for slot := 1; slot <= width; slot++ {
			if value := lines[lineNumber][slot*4-3]; value != ' ' {
				if d.stacks[slot] == nil {
					d.stacks[slot] = deque.New[byte]()
				}
				d.stacks[slot].PushBack(value)
			}
		}
	}

	// Skip empty line
	lineNumber += 2

	// Read instructions
	for ; lineNumber < len(lines); lineNumber++ {
		fmt.Sscanf(lines[lineNumber], "move %d from %d to %d", &instr.n, &instr.from, &instr.to)
		d.instructions = append(d.instructions, instr)
	}

	return nil
}

func (d *DayImpl) Part1() (string, error) {
	for _, instruction := range d.instructions {
		for i := 0; i < instruction.n; i++ {
			item := d.stacks[instruction.from].PopFront()
			d.stacks[instruction.to].PushFront(item)
		}
	}

	return formatStacks(d.stacks), nil
}

func (d *DayImpl) Part2() (string, error) {
	for _, instruction := range d.instructions {
		buffer := deque.New[byte]()

		for i := 0; i < instruction.n; i++ {
			buffer.PushFront(d.stacks[instruction.from].PopFront())
		}

		for i := 0; i < instruction.n; i++ {
			d.stacks[instruction.to].PushFront(buffer.PopFront())
		}
	}

	return formatStacks(d.stacks), nil
}

func main() {
	utils.Run(&DayImpl{}, true)
}
