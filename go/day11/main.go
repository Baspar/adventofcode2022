package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	utils "github.com/baspar/adventofcode2022/internal"
)

type Operation struct {
	op  byte
	val int
}

func (operation Operation) Call(old int) int {
	n := operation.val
	if n == 0 {
		n = old
	}

	if operation.op == '*' {
		return old * n
	} else {
		return old + n
	}
}

type Monkey struct {
	items           []int
	operation       Operation
	testDivisibleBy int
	throwIfTrue     int
	throwIfFalse    int
}

func (m *Monkey) InspectAndThrow(monkeyGetsBored bool) (throwToMonkey int, value int) {
	item := m.items[0]
	m.items = m.items[1:]

	item = m.operation.Call(item)
	if monkeyGetsBored {
		item /= 3
	}

	if item%m.testDivisibleBy == 0 {
		return m.throwIfTrue, item
	} else {
		return m.throwIfFalse, item
	}
}

func NewMonkey(lines []string) (monkey Monkey) {
	// Items
	for _, item := range strings.Split(strings.Split(lines[1], ": ")[1], ", ") {
		n, _ := strconv.Atoi(item)
		monkey.items = append(monkey.items, int(n))
	}

	// Operation
	var value string
	fmt.Sscanf(lines[2], "  Operation: new = old %c %s", &monkey.operation.op, &value)
	if value != "old" {
		monkey.operation.val, _ = strconv.Atoi(value)
	}

	// Test
	fmt.Sscanf(lines[3], "  Test: divisible by %d", &monkey.testDivisibleBy)
	fmt.Sscanf(lines[4], "    If true: throw to monkey %d", &monkey.throwIfTrue)
	fmt.Sscanf(lines[5], "    If false: throw to monkey %d", &monkey.throwIfFalse)
	return monkey
}

type DayImpl struct {
	monkeys []Monkey
}

func (d *DayImpl) getMonkeyBusiness(numberOfRound int, monkeyGetsBored bool) int {
	objectsInspected := make([]int, len(d.monkeys))

	leastCommonMultiplier := 1
	for _, m := range d.monkeys {
		leastCommonMultiplier *= m.testDivisibleBy
	}

	for round := 0; round < numberOfRound; round++ {
		for monkeyIndex := 0; monkeyIndex < len(d.monkeys); monkeyIndex++ {
			objectsInspected[monkeyIndex] += len(d.monkeys[monkeyIndex].items)
			for len(d.monkeys[monkeyIndex].items) > 0 {
				throwToMonkey, item := d.monkeys[monkeyIndex].InspectAndThrow(monkeyGetsBored)
				d.monkeys[throwToMonkey].items = append(d.monkeys[throwToMonkey].items, item%leastCommonMultiplier)
			}
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(objectsInspected)))

	return objectsInspected[0] * objectsInspected[1]
}

func (d *DayImpl) Init(lines []string) error {
	d.monkeys = nil
	for i := 0; i < len(lines); i += 7 {
		d.monkeys = append(d.monkeys, NewMonkey(lines[i:i+6]))
	}
	return nil
}
func (d *DayImpl) Part1() (string, error) {
	return fmt.Sprint(d.getMonkeyBusiness(20, true)), nil
}
func (d *DayImpl) Part2() (string, error) {
	return fmt.Sprint(d.getMonkeyBusiness(10000, false)), nil
}

func main() {
	utils.Run(&DayImpl{}, true)
}
