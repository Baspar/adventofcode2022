package main

import (
	"encoding/json"
	"fmt"
	. "math"

	utils "github.com/baspar/adventofcode2022/internal"
	"github.com/baspar/adventofcode2022/internal/math"
)

const (
	CorrectOrder = iota
	ContinueChecking
	IncorrectOrder
)

type DayImpl struct {
	comparisons []Comparison
}

type Node struct {
	val    *int
	array  []Node
	marked bool
}

func (n Node) IsArray() bool {
	return n.val == nil
}
func (n Node) ToArrayNode() Node {
	if n.IsArray() {
		return n
	} else {
		return Node{array: []Node{n}}
	}
}

func (left Node) isInOrderWith(right Node) int {
	if left.IsArray() || right.IsArray() {
		leftArray, rightArray := left.ToArrayNode(), right.ToArrayNode()
		for i := 0; i < math.Min(len(leftArray.array), len(rightArray.array)); i++ {
			if status := leftArray.array[i].isInOrderWith(rightArray.array[i]); status == ContinueChecking {
				continue
			} else {
				return status
			}
		}

		if len(rightArray.array) == len(leftArray.array) {
			return ContinueChecking
		} else if len(rightArray.array) > len(leftArray.array) {
			return CorrectOrder
		} else {
			return IncorrectOrder
		}
	}

	if *left.val == *right.val {
		return ContinueChecking
	} else if *left.val < *right.val {
		return CorrectOrder
	} else {
		return IncorrectOrder
	}
}

func (left Node) IsInOrderWith(right Node) bool {
	return left.isInOrderWith(right) != IncorrectOrder
}

func from(items []any) (node Node) {
	for _, item := range items {
		var child Node
		if arr, isArr := item.([]any); isArr {
			child.array = append(node.array, from(arr))
		} else if n, isNumber := item.(float64); isNumber {
			asInt := int(Round(n))
			child.val = &asInt
		}
		node.array = append(node.array, child)
	}

	return node
}

func From(line string) Node {
	var items []any
	json.Unmarshal([]byte(line), &items)
	return from(items)
}

type Comparison struct {
	left  Node
	right Node
}

func (d *DayImpl) Init(lines []string) error {
	for i := 0; i < len(lines); i += 3 {
		d.comparisons = append(d.comparisons, Comparison{
			left:  From(lines[i]),
			right: From(lines[i+1]),
		})
	}
	return nil
}
func (d *DayImpl) Part1() (string, error) {
	sum := 0
	for i, comparison := range d.comparisons {
		if comparison.left.IsInOrderWith(comparison.right) {
			sum += i + 1
		}
	}
	return fmt.Sprint(sum), nil
}
func (d *DayImpl) Part2() (string, error) {
	specialNode1, positionNode1 := From("[[2]]"), 1
	specialNode2, positionNode2 := From("[[6]]"), 2
	for _, comparison := range d.comparisons {
		for _, node := range []Node{comparison.left, comparison.right} {
			if node.IsInOrderWith(specialNode1) {
				positionNode1++
			}
			if node.IsInOrderWith(specialNode2) {
				positionNode2++
			}

		}
	}
	return fmt.Sprint(positionNode1 * positionNode2), nil
}

func main() {
	utils.Run(&DayImpl{}, false)
}
