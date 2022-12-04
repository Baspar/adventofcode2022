package main

import (
	"fmt"
	"strconv"
	"strings"

	. "github.com/baspar/adventofcode2022/internal"
)

type Assignement struct {
	from int
	to   int
}

func toAssignement(s string) Assignement {
	var (
		from, to int
		err      error
	)
	sections := strings.Split(s, "-")
	if from, err = strconv.Atoi(sections[0]); err != nil {
		panic(err)
	}
	if to, err = strconv.Atoi(sections[1]); err != nil {
		panic(err)
	}

	return Assignement{from, to}
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
		assignments := strings.Split(line, ",")
		assign1, assign2 := toAssignement(assignments[0]), toAssignement(assignments[1])
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
