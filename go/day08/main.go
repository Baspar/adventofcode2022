package main

import (
	"fmt"

	utils "github.com/baspar/adventofcode2022/internal"
)

type Grid [][]int

func (g Grid) isOutOfBound(i, j int) bool {
	return i < 0 || j < 0 || i >= len(g) || j >= len(g[i])
}

func (g Grid) getVisibilityDistance(i, j, di, dj int) int {
	for n := 0; ; n++ {
		if g.isOutOfBound(i+n*di, j+n*dj) {
			return n - 1
		}
		if n != 0 && g[i+n*di][j+n*dj] >= g[i][j] {
			return n
		}
	}
}
func (g Grid) isFullyVisible(i, j, di, dj int) bool {
	for n := 1; !g.isOutOfBound(i+n*di, j+n*dj); n++ {
		if g[i][j] <= g[i+n*di][j+n*dj] {
			return false
		}
	}
	return true
}

type DayImpl struct {
	grid Grid
}

func (d *DayImpl) Init(lines []string) error {
	for i, line := range lines {
		d.grid = append(d.grid, nil)
		for _, char := range line {
			d.grid[i] = append(d.grid[i], int(char-'0'))
		}
	}
	return nil
}

func (d *DayImpl) Part1() (string, error) {
	visibles := 0
	for i := 0; i < len(d.grid); i++ {
		for j := 0; j < len(d.grid); j++ {
			isVisible := false
			isVisible = isVisible || d.grid.isFullyVisible(i, j, 0, 1)
			isVisible = isVisible || d.grid.isFullyVisible(i, j, 0, -1)
			isVisible = isVisible || d.grid.isFullyVisible(i, j, 1, 0)
			isVisible = isVisible || d.grid.isFullyVisible(i, j, -1, 0)
			if isVisible {
				visibles++
			}
		}
	}
	return fmt.Sprint(visibles), nil
}
func (d *DayImpl) Part2() (string, error) {
	maxScore := 0
	for i := 0; i < len(d.grid); i++ {
		for j := 0; j < len(d.grid); j++ {
			score := 1
			score *= d.grid.getVisibilityDistance(i, j, 0, 1)
			score *= d.grid.getVisibilityDistance(i, j, 0, -1)
			score *= d.grid.getVisibilityDistance(i, j, 1, 0)
			score *= d.grid.getVisibilityDistance(i, j, -1, 0)
			if score > maxScore {
				maxScore = score
			}
		}
	}
	return fmt.Sprint(maxScore), nil
}

func main() {
	utils.Run(&DayImpl{}, false)
}
