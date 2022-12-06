package main

import (
	"fmt"

	utils "github.com/baspar/adventofcode2022/internal"
)

type DayImpl struct {
	stream string
}

func (d *DayImpl) findRollingWindow(size int) int {
	hist := make(map[byte]int)

	for i, upper := range d.stream {
		if i >= size {
			lower := d.stream[i-size]
			if hist[lower]--; hist[lower] == 0 {
				delete(hist, lower)
			}
		}

		hist[byte(upper)]++
		if len(hist) == size {
			return i + 1
		}
	}
	return -1
}
func (d *DayImpl) Init(lines []string) error {
	d.stream = lines[0]
	return nil
}
func (d *DayImpl) Part1() (string, error) {
	return fmt.Sprint(d.findRollingWindow(4)), nil
}
func (d *DayImpl) Part2() (string, error) {
	return fmt.Sprint(d.findRollingWindow(14)), nil
}

func main() {
	utils.Run(&DayImpl{}, false)
}
