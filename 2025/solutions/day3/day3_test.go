package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	// Add tests here
	result := PartOne()
	println(result)
}

func TestPartTwo(t *testing.T) {
	assert := assert.New(t)
	// Add tests here
	testInput := "811111111111119"
	var expected int64 = 811111111119

	result := GetJoltagePtTwo(testInput)

	assert.Equal(expected, result, "Expected should match result")
}

func TestPartTwoAgain(t *testing.T) {
	assert := assert.New(t)
	// Add tests here
	testInput := "234234234234278"
	var expected int64 = 434234234278

	result := GetJoltagePtTwo(testInput)

	assert.Equal(expected, result, "Expected should match result")
}

func TestPartTwoAgainAgain(t *testing.T) {
	assert := assert.New(t)
	// Add tests here
	testInput := "818181911112111"
	var expected int64 = 888911112111

	result := GetJoltagePtTwo(testInput)

	assert.Equal(expected, result, "Expected should match result")
}

func TestPartTwoFull(t *testing.T) {
	result := PartTwo()
	println(result)
}
