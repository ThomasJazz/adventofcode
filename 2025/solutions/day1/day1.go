package day1

import (
	"fmt"
	"math"
	"strconv"

	"github.com/thomasjazz/adventofcode/lib"
)

var (
	filename = "../../input/input-day1.txt"
	//filename = "../../input/sample-input.txt"
)

type Instruction struct {
	direction string
	clicks    int
}

func PartOne() int {
	fmt.Println("Day 1 solution pt 1")
	lines := lib.ReadInput(filename)
	timesHitZero := 0

	position := 50

	for _, line := range lines {
		instruction := ParseInstruction(line)

		switch instruction.direction {
		case "L":
			position = (position - instruction.clicks) % 100
		case "R":
			position = (position + instruction.clicks) % 100
		}

		if position < 0 {
			position += 100
		}

		if position == 0 {
			timesHitZero++
		}
	}

	return timesHitZero
}

func PartTwo() int {
	fmt.Println("Day 1 solution pt 2")
	lines := lib.ReadInput(filename)
	timesHitZero := 0

	position := 50

	for _, line := range lines {
		println(position)
		instruction := ParseInstruction(line)
		rawPosition := position

		// If we started at zero and then go into the negatives, we dont want to count the starting 0
		startedAtZero := false
		if position == 0 {
			startedAtZero = true
		}

		switch instruction.direction {
		case "L":
			rawPosition = (position - instruction.clicks)
			position = rawPosition % 100
		case "R":
			rawPosition = (position + instruction.clicks)
			position = rawPosition % 100
		}

		// Always want a positive position
		if position < 0 {
			position += 100
		}

		timesHitZeroThisTime := int(math.Abs(float64(rawPosition)) / 100)

		if rawPosition <= 0 && !startedAtZero {
			timesHitZeroThisTime++
		}

		timesHitZero += timesHitZeroThisTime

		println("\t", line, rawPosition, timesHitZeroThisTime, timesHitZero)
	}

	return timesHitZero
}

func ParseInstruction(line string) Instruction {
	direction := line[:1]
	clicks, _ := strconv.Atoi(line[1:])
	return Instruction{
		direction: direction,
		clicks:    clicks,
	}
}
