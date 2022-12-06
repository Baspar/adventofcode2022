package main

import (
	"fmt"

	. "github.com/baspar/adventofcode2022/internal"
)

type Assignement struct {
	from int
	to   int
}

func (a Assignement) partialOverlapWith(a_ Assignement) bool {
	return a_.from <= a.from && a.from <= a_.to
}
func (a Assignement) fullOverlapWith(a_ Assignement) bool {
	return a.from <= a_.from && a_.to <= a.to
}

type DayImpl struct {
	assignments [][2]Assignement
}

func (d *DayImpl) Init(lines []string) error {
	for _, line := range lines {
		var assign1, assign2 Assignement
		fmt.Sscanf(line, "%d-%d,%d-%d", &assign1.from, &assign1.to, &assign2.from, &assign2.to)
		d.assignments = append(d.assignments, [2]Assignement{assign1, assign2})
	}
	return nil
}
func (d *DayImpl) Part1() (string, error) {
	numberOverlappedAssigns := 0
	for _, assigns := range d.assignments {
		if assigns[0].fullOverlapWith(assigns[1]) || assigns[1].fullOverlapWith(assigns[0]) {
			numberOverlappedAssigns += 1
		}
	}

	return fmt.Sprint(numberOverlappedAssigns), nil
}
func (d *DayImpl) Part2() (string, error) {
	numberOverlappedAssigns := 0
	for _, assigns := range d.assignments {
		if assigns[0].partialOverlapWith(assigns[1]) || assigns[1].partialOverlapWith(assigns[0]) {
			numberOverlappedAssigns += 1
		}
	}

	return fmt.Sprint(numberOverlappedAssigns), nil
}

func main() {
	Run(&DayImpl{}, false)
}
