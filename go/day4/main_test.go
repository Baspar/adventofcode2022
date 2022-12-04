package main

import (
	"testing"

	. "github.com/baspar/adventofcode2022/internal"
	"github.com/stretchr/testify/assert"
)

var input = SanitizeInput(`2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`)

func TestPart1(t *testing.T) {
	assert := assert.New(t)

	d := &DayImpl{}
	d.Init(input)

	res, err := d.Part1()

	assert.Equal("2", res)
	assert.Nil(err)
}

func TestPart2(t *testing.T) {
	assert := assert.New(t)

	d := &DayImpl{}
	d.Init(input)

	res, err := d.Part2()

	assert.Equal("4", res)
	assert.Nil(err)
}
