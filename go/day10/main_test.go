package main

import (
	"testing"

	utils "github.com/baspar/adventofcode2022/internal"
	"github.com/stretchr/testify/assert"
)

var input = `addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop`

var output = `▒▒..▒▒..▒▒..▒▒..▒▒..▒▒..▒▒..▒▒..▒▒..▒▒..
▒▒▒...▒▒▒...▒▒▒...▒▒▒...▒▒▒...▒▒▒...▒▒▒.
▒▒▒▒....▒▒▒▒....▒▒▒▒....▒▒▒▒....▒▒▒▒....
▒▒▒▒▒.....▒▒▒▒▒.....▒▒▒▒▒.....▒▒▒▒▒.....
▒▒▒▒▒▒......▒▒▒▒▒▒......▒▒▒▒▒▒......▒▒▒▒
▒▒▒▒▒▒▒.......▒▒▒▒▒▒▒.......▒▒▒▒▒▒▒.....`

var inputs = map[string][2]string{
	input: {"13140", output},
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
