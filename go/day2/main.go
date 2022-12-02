package main

import (
	"fmt"

	utils "github.com/baspar/adventofcode2022/internal"
)

type Instruction struct {
	opponent  int
	objective int
}

func (i Instruction) scoreShape(play int) int {
	return play + 1
}
func (i Instruction) scoreOutcome(play int) int {
	result := (3 + play - i.opponent) % 3
	switch result {
	case 0: // Tie
		return 3
	case 1: // Win
		return 6
	default: // Lost
		return 0
	}
}
func (i Instruction) playToWin() int {
	var delta int
	switch i.objective {
	case 0: // Lose
		delta = -1
	case 1: // Tie
		delta = 0
	case 2: // Win
		delta = 1
	}
	return (3 + i.opponent + delta) % 3
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
	utils.Run(&DayImpl{}, false)
}
