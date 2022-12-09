package main

import (
	"testing"

	utils "github.com/baspar/adventofcode2022/internal"
	"github.com/stretchr/testify/assert"
)

var shortInput = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`
var longInput = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`

var inputs = map[string][2]string{
	shortInput: {"13", "1"},
	longInput:  {"", "36"},
}

func TestPart1(t *testing.T) {
	assert := assert.New(t)

	for input, expectedRes := range inputs {
		if expectedRes[0] == "" {
			continue
		}
		d := &DayImpl{}
		d.Init(utils.SanitizeInput(input))

		res, err := d.Part1()

		assert.Equal(expectedRes[0], res)
		assert.Nil(err)
	}
}

func TestPart2(t *testing.T) {
	assert := assert.New(t)

	for input, expectedRes := range inputs {
		if expectedRes[1] == "" {
			continue
		}
		d := &DayImpl{}
		d.Init(utils.SanitizeInput(input))

		res, err := d.Part2()

		assert.Equal(expectedRes[1], res)
		assert.Nil(err)
	}
}
