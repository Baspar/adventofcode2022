package main

import (
	"testing"

	. "github.com/baspar/adventofcode2022/internal"
	"github.com/stretchr/testify/assert"
)

var input = SanitizeInput(``)

func TestPart1(t *testing.T) {
	assert := assert.New(t)

	d := &DayImpl{}
	d.Init(input)

	res, err := d.Part1()

	assert.Equal("", res)
	assert.Nil(err)
}

func TestPart2(t *testing.T) {
	assert := assert.New(t)

	d := &DayImpl{}
	d.Init(input)

	res, err := d.Part2()

	assert.Equal("", res)
	assert.Nil(err)
}
