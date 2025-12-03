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
		maxJoltages = append(maxJoltages, concatted)
	}

	return sumMaxJoltages(maxJoltages)
}

func sumMaxJoltages(joltages []int) int {
	sum := 0
	for _, val := range joltages {
		sum += val
	}
	return sum
}

func PartTwo() {
	fmt.Println("Day 3 solution pt 2")
	// lines := lib.ReadInput(filename)
}
