package main

import (
	"testing"

	utils "github.com/baspar/adventofcode2022/internal"
	"github.com/stretchr/testify/assert"
)

var input = utils.SanitizeInput(`A Y
B X
C Z`)

func TestPart1(t *testing.T) {
	assert := assert.New(t)

	d := &DayImpl{}
	d.Init(input)

	res, err := d.Part1()

	assert.Equal("15", res)
	assert.Nil(err)
}

func TestPart2(t *testing.T) {
	assert := assert.New(t)

	d := &DayImpl{}
	d.Init(input)

	res, err := d.Part2()

	assert.Equal("12", res)
	assert.Nil(err)
}
