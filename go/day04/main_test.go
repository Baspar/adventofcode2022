package main

import (
	"testing"

	utils "github.com/baspar/adventofcode2022/internal"
	"github.com/stretchr/testify/assert"
)

var input = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

var inputs = map[string][2]string{
	input: {"2", "4"},
}


func TestPart1(t *testing.T) {
	assert := assert.New(t)

	for input, expectedRes := range inputs {
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
		d := &DayImpl{}
		d.Init(utils.SanitizeInput(input))

		res, err := d.Part2()

		assert.Equal(expectedRes[1], res)
		assert.Nil(err)
	}
}
