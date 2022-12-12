package coords

import (
	"fmt"

	"github.com/baspar/adventofcode2022/internal/math"
)

type Coords [2]int

var Zero = Coords{0, 0}

var NaturalDirections = map[byte]Coords{
	'U': {1, 0},
	'D': {-1, 0},
	'R': {0, 1},
	'L': {0, -1},
}

func (c Coords) Str() string {
	return fmt.Sprintf("%d_%d", c[0], c[1])
}
func (c1 Coords) Diff(c2 Coords) Coords {
	return Coords{c1[0] - c2[0], c1[1] - c2[1]}
}
func (c1 Coords) Distance(c2 Coords) int {
	return math.Max(math.Abs(c1[0]-c2[0]), math.Abs(c1[1]-c2[1]))
}
func (coords *Coords) AddMut(c Coords) {
	coords[0] += c[0]
	coords[1] += c[1]
}
func (coords *Coords) Add(c Coords) Coords {
	return Coords{coords[0] + c[0], coords[1] + c[1]}
}
