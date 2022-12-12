package main

import (
	"fmt"

	utils "github.com/baspar/adventofcode2022/internal"
	"github.com/baspar/adventofcode2022/internal/coords"
	"github.com/baspar/adventofcode2022/internal/math"
)

type Instruction struct {
	direction coords.Coords
	distance  int
}

type Knot struct{ coords.Coords }

func (knot *Knot) MoveToward(destination Knot) {
	if knot.Distance(destination.Coords) > 1 {
		diff := knot.Diff(destination.Coords)
		knot.Coords[0] -= math.Sign(diff[0])
		knot.Coords[1] -= math.Sign(diff[1])
	}
}

type Rope []Knot

func (r Rope) Tail() Knot {
	return r[len(r)-1]
}

type Exploration struct {
	seen map[string]bool
	rope Rope
}

func (e *Exploration) Move(direction coords.Coords) {
	e.rope[0].AddMut(direction)
	for knotIndex := 1; knotIndex < len(e.rope); knotIndex++ {
		e.rope[knotIndex].MoveToward(e.rope[knotIndex-1])
	}
}
func (e *Exploration) Run(instructions []Instruction) {
	for _, instruction := range instructions {
		for i := 0; i < instruction.distance; i++ {
			e.Move(instruction.direction)
			e.seen[e.rope.Tail().Str()] = true
		}
	}
}

func NewExploration(length int) Exploration {
	return Exploration{
		seen: map[string]bool{coords.Zero.Str(): true},
		rope: make([]Knot, length),
	}
}

type DayImpl struct {
	instructions []Instruction
}

func (d *DayImpl) Init(lines []string) error {
	for _, line := range lines {
		var (
			direction   byte
			instruction Instruction
		)
		fmt.Sscanf(line, "%c %d", &direction, &instruction.distance)
		instruction.direction = coords.NaturalDirections[direction]
		d.instructions = append(d.instructions, instruction)
	}
	return nil
}
func (d *DayImpl) Part1() (string, error) {
	exploration := NewExploration(2)
	exploration.Run(d.instructions)
	return fmt.Sprint(len(exploration.seen)), nil
}
func (d *DayImpl) Part2() (string, error) {
	exploration := NewExploration(10)
	exploration.Run(d.instructions)
	return fmt.Sprint(len(exploration.seen)), nil
}

func main() {
	utils.Run(&DayImpl{}, false)
}
