package day3

import (
	"fmt"
	"strconv"

	"github.com/thomasjazz/adventofcode/lib"
)

var (
	filename = "input-day3.txt"
)

func PartOne() int {
	fmt.Println("Day 3 solution pt 1")
	lines := lib.ReadInput(filename)

	var maxJoltages []int

	for _, line := range lines {
		maxJoltages = append(maxJoltages, GetJoltagePtOne(line))
	}

	return sumMaxJoltages(maxJoltages)
}

func GetJoltagePtOne(line string) int {
	maxLeft, maxRight := -1, -1

	for i := 0; i < len(line)-1; i++ {
		leftNum, _ := strconv.Atoi(string(line[i]))
		rightNum, _ := strconv.Atoi(string(line[i+1]))

		if leftNum > maxLeft {
			maxLeft = leftNum
			maxRight = -1
		}

		if rightNum > maxRight {
			maxRight = rightNum
		}
	}
	concatted, _ := strconv.Atoi(strconv.Itoa(maxLeft) + strconv.Itoa(maxRight))
	return concatted
}

func PartTwo() int64 {
	fmt.Println("Day 3 solution pt 2")
	lines := lib.ReadInput(filename)
	var maxJoltages []int64

	for _, line := range lines {
		maxJoltages = append(maxJoltages, GetJoltagePtTwo(line))
	}

	return sumMaxJoltages64(maxJoltages)
}

func GetJoltagePtTwo(line string) int64 {
	var joltage string

	k := 0
	lineLength := len(line)
	for i := 0; i < 12; i++ {
		currentMax := -1
		currentMaxIndex := 0
		maxScanIndex := lineLength - (11 - i)

		for k < maxScanIndex {
			current, _ := strconv.Atoi(string(line[k]))

			if current > currentMax {
				currentMax = current
				currentMaxIndex = k
			}
			k++
		}
		// Restart search from where we found last max number
		k = currentMaxIndex + 1

		joltage = joltage + strconv.Itoa(currentMax)
	}

	i64, _ := strconv.ParseInt(joltage, 10, 64)
	return i64
}

func sumMaxJoltages64(joltages []int64) int64 {
	sum := int64(0)
	for _, val := range joltages {
		sum += val
	}
	return sum
}

func sumMaxJoltages(joltages []int) int {
	sum := 0
	for _, val := range joltages {
		sum += val
	}
	return sum
}
