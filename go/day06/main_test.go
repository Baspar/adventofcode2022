package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var inputs = map[string][2]string{
	`mjqjpqmgbljsphdztnvjfqwrcgsmlb`:    {"7", "19"},
	`bvwbjplbgvbhsrlpgdmjqwftvncz`:      {"5", "23"},
	`nppdvjthqldpwncqszvftbrmjlhg`:      {"6", "23"},
	`nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg`: {"10", "29"},
	`zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`:  {"11", "26"},
}

func TestPart1(t *testing.T) {
	assert := assert.New(t)

	for input, expectedRes := range inputs {
		d := &DayImpl{}
		d.Init([]string{input})

		res, err := d.Part1()

		assert.Equal(expectedRes[0], res)
		assert.Nil(err)
	}
}

func TestPart2(t *testing.T) {
	assert := assert.New(t)

	for input, expectedRes := range inputs {
		d := &DayImpl{}
		d.Init([]string{input})

		res, err := d.Part2()

		assert.Equal(expectedRes[1], res)
		assert.Nil(err)
	}
}
