package main

import (
	"fmt"

	. "github.com/baspar/adventofcode2022/internal"
	"github.com/baspar/adventofcode2022/internal/set"
)

type x [2]string
type DayImpl struct {
	rucksacks []string
}

func badgeToScore(badge rune) int {
	if 'a' <= badge && badge <= 'z' {
		return int(badge - 'a' + 1)
	} else {
		return int(badge - 'A' + 27)
	}
}

func (d *DayImpl) Init(lines []string) error {
	d.rucksacks = lines
	return nil
}
func (d *DayImpl) Part1() (string, error) {
	score := 0
	for _, rucksacks := range d.rucksacks {
		n := len(rucksacks)
		rucksack1, rucksack2 := rucksacks[:n/2], rucksacks[n/2:]

		badge := set.Intersect([]rune(rucksack1), []rune(rucksack2))
		score += badgeToScore(badge[0])
	}
	return fmt.Sprint(score), nil
}
func (d *DayImpl) Part2() (string, error) {
	score := 0
	for i := 0; i < len(d.rucksacks); i += 3 {
		rucksack1, rucksack2, rucksack3 := d.rucksacks[i], d.rucksacks[i+1], d.rucksacks[i+2]

		badge := set.Intersect([]rune(rucksack1), []rune(rucksack2), []rune(rucksack3))[0]
		score += badgeToScore(badge)
	}
	return fmt.Sprint(score), nil
}

func main() {
	Run(&DayImpl{}, false)
}
