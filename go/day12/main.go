package main

import (
	"fmt"

	utils "github.com/baspar/adventofcode2022/internal"
	"github.com/baspar/adventofcode2022/internal/coords"
	. "github.com/baspar/adventofcode2022/internal/pq"
)

type Exploration struct {
	coords.Coords
	distance int
}

type Grid [][]rune

func (g Grid) Get(c coords.Coords) rune {
	return g[c[0]][c[1]]
}
func (g Grid) AreValidCoords(c coords.Coords) bool {
	return c[0] >= 0 && c[1] >= 0 && c[0] < len(g) && c[1] < len(g[c[0]])
}

type DayImpl struct {
	start coords.Coords
	end   coords.Coords
	grid  Grid
}

func (d *DayImpl) Init(lines []string) error {
	for i, line := range lines {
		d.grid = append(d.grid, nil)
		for j, val := range line {
			switch val {
			case 'S':
				d.start = coords.Coords{i, j}
				d.grid[i] = append(d.grid[i], 'a'-'a')
			case 'E':
				d.end = coords.Coords{i, j}
				d.grid[i] = append(d.grid[i], 'z'-'a')
			default:
				d.grid[i] = append(d.grid[i], val-'a')
			}
		}
	}
	return nil
}
func (d DayImpl) dijkstra(
	start coords.Coords,
	isEnd func(Exploration) bool,
	shouldExplore func(oldCoords, newCoords coords.Coords) bool,
) int {
	wasVisited := make(map[coords.Coords]bool)
	pq := NewPriorityQueue(
		func(expl Exploration) int { return expl.distance },
		Exploration{start, 0},
	)
	for !pq.Empty() {
		expl := pq.Pop()

		if wasVisited[expl.Coords] {
			continue
		}
		wasVisited[expl.Coords] = true

		if isEnd(expl) {
			return expl.distance
		}

		for _, direction := range coords.NaturalDirections {
			newCoords := expl.Coords.Add(direction)
			if !d.grid.AreValidCoords(newCoords) {
				continue
			}

			if shouldExplore(expl.Coords, newCoords) {
				pq.Push(Exploration{newCoords, expl.distance + 1})
			}
		}
	}
	return -1
}
func (d *DayImpl) Part1() (string, error) {
	return fmt.Sprint(d.dijkstra(
		d.start,
		func(expl Exploration) bool {
			return expl.Coords == d.end
		},
		func(from, to coords.Coords) bool {
			return d.grid.Get(from)-d.grid.Get(to) >= -1
		},
	)), nil
}
func (d *DayImpl) Part2() (string, error) {
	return fmt.Sprint(d.dijkstra(
		d.end,
		func(expl Exploration) bool {
			return d.grid.Get(expl.Coords) == 0
		},
		func(from, to coords.Coords) bool {
			return d.grid.Get(to)-d.grid.Get(from) >= -1
		},
	)), nil
}

func main() {
	utils.Run(&DayImpl{}, false)
}
