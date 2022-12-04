package main

import (
	"fmt"

	. "github.com/baspar/adventofcode2022/internal"
)

type Instruction struct {
	opponent  int
	objective int
}

func (i Instruction) scoreShape(play int) int {
	return play + 1
}
func (i Instruction) scoreOutcome(play int) int {
	return (4 + play - i.opponent) % 3 * 3
}
func (i Instruction) playToWin() int {
	return (2 + i.opponent + i.objective) % 3
}

type DayImpl struct {
	instructions []Instruction
}

func (d *DayImpl) Init(lines []string) error {
	for _, line := range lines {
		d.instructions = append(d.instructions, Instruction{
			opponent:  int(line[0] - 'A'),
			objective: int(line[2] - 'X'),
		})
	}
	return nil
}
func (d *DayImpl) Part1() (string, error) {
	score := 0
	for _, instruction := range d.instructions {
		score += instruction.scoreOutcome(instruction.objective) + instruction.scoreShape(instruction.objective)
	}
	return fmt.Sprint(score), nil
}
func (d *DayImpl) Part2() (string, error) {
	score := 0
	for _, instruction := range d.instructions {
		playToWin := instruction.playToWin()
		score += instruction.scoreOutcome(playToWin) + instruction.scoreShape(playToWin)
	}
	return fmt.Sprint(score), nil
}

func main() {
	Run(&DayImpl{}, false)
}
