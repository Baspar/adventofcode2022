package main

import (
	"fmt"
	"math"

	utils "github.com/baspar/adventofcode2022/internal"
)

type File struct {
	name string
	size int
}
type Dir struct {
	name  string
	size  int
	Files []File
	Dirs  []*Dir
}

func (d Dir) sumDirsSmallerThan(maxSize int) (sum int) {
	if d.size < maxSize {
		sum += d.size
	}
	for _, dir := range d.Dirs {
		sum += dir.sumDirsSmallerThan(maxSize)
	}
	return sum
}
func (d Dir) sizeOfSmallestDirBiggerThan(minSize int) int {
	smallestDirSize := math.MaxInt

	if d.size > minSize {
		smallestDirSize = d.size
	}
	for _, dir := range d.Dirs {
		if smallestSubDirSize := dir.sizeOfSmallestDirBiggerThan(minSize); smallestSubDirSize < smallestDirSize {
			smallestDirSize = smallestSubDirSize
		}
	}
	return smallestDirSize
}
func (d *Dir) readFrom(lines []string, lineIndex int) (newLineIndex int) {
	for ; lineIndex < len(lines) && lines[lineIndex] != "$ cd .."; lineIndex++ {
		var (
			file File
			dir Dir
		)
		if _, err := fmt.Sscanf(lines[lineIndex], "%d %s", &file.size, &file.name); err == nil {
			d.Files = append(d.Files, file)
			d.size += file.size
		} else if _, err := fmt.Sscanf(lines[lineIndex], "$ cd %s", &dir.name); err == nil {
			lineIndex = dir.readFrom(lines, lineIndex+1)
			d.size += dir.size
			d.Dirs = append(d.Dirs, &dir)
		}
	}
	return lineIndex
}

type DayImpl struct {
	root Dir
}

func (d *DayImpl) Init(lines []string) error {
	d.root.readFrom(lines, 0)
	return nil
}
func (d *DayImpl) Part1() (string, error) {
	return fmt.Sprint(d.root.sumDirsSmallerThan(100000)), nil
}
func (d *DayImpl) Part2() (string, error) {
	return fmt.Sprint(d.root.sizeOfSmallestDirBiggerThan(d.root.size - 40000000)), nil
}

func main() {
	utils.Run(&DayImpl{}, false)
}
