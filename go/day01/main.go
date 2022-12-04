package main

import (
	"fmt"
	"sort"
	"strconv"

	. "github.com/baspar/adventofcode2022/internal"
)

type DayImpl struct {
	stacks []int
}

func (d *DayImpl) Init(lines []string) error {
	stack := 0
	for _, line := range lines {
		if len(line) == 0 {
			d.stacks = append(d.stacks, stack)
			stack = 0
		} else if calories, err := strconv.Atoi(line); err == nil {
			stack += calories
		} else {
			return err
		}
	}
	if stack > 0 {
		d.stacks = append(d.stacks, stack)
	}
	sort.Ints(d.stacks)
	return nil
}
func (d *DayImpl) Part1() (string, error) {
	return fmt.Sprint(d.stacks[len(d.stacks)-1]), nil
}
func (d *DayImpl) Part2() (string, error) {
	return fmt.Sprint(d.stacks[len(d.stacks)-1] + d.stacks[len(d.stacks)-2] + d.stacks[len(d.stacks)-3]), nil
}

func main() {
	Run(&DayImpl{}, false)
}
