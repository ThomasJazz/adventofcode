// https://adventofcode.com/2023/day/1
package day01

import (
	"strconv"
	"strings"

	"github.com/thomasjazz/adventofcode/lib"
)

var (
	filename string         = "2023/input/day_1_input.txt"
	alphaMap map[string]int = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
)

func DayOnePartOne() {
	lines := lib.ReadInput(filename)

	var calibrationSum int = 0

	for _, line := range lines {
		calibrationSum += FindCalibrationNumbers(line)
	}

	println(calibrationSum)
}

func DayOnePartTwo() {
	lines := lib.ReadInput(filename)

	var calibrationSum int = 0

	for _, line := range lines {
		calibrationVal, _ := FindCalibrationNumbersTwo(line)

		calibrationSum += calibrationVal
	}

	println(calibrationSum)
}

func FindCalibrationNumbers(line string) int {
	lineArray := strings.Split(line, "")
	var left, right int

	// Search ->
	for _, c := range lineArray {
		val, err := strconv.Atoi(c)

		if err == nil {
			left = val
			break
		}
	}

	// Search <-
	for i := len(lineArray) - 1; i >= 0; i-- {
		val, err := strconv.Atoi(lineArray[i])
		if err == nil {
			right = val
			break
		}
	}

	calibrationVal, _ := strconv.Atoi(strconv.Itoa(left) + strconv.Itoa(right))

	return int(calibrationVal)
}

func FindCalibrationNumbersTwo(line string) (int, error) {
	var left, right int
	leftMinIndex, rightMaxIndex := len(line), -1

	// Find left & right most spelled-out number
	for substring, num := range alphaMap {
		// Check if string contains string representations of numbers; "one"
		leftFound := strings.Index(line, substring)
		rightFound := strings.LastIndex(line, substring)

		if leftFound != -1 && leftFound < leftMinIndex {
			left = num
			leftMinIndex = leftFound
		}

		if rightFound > rightMaxIndex {
			right = num
			rightMaxIndex = rightFound
		}

		// Check if string contains int from map
		leftFound = strings.Index(line, strconv.Itoa(num))
		rightFound = strings.LastIndex(line, strconv.Itoa(num))

		if leftFound != -1 && leftFound < leftMinIndex {
			left = num
			leftMinIndex = leftFound
		}

		if rightFound > rightMaxIndex {
			right = num
			rightMaxIndex = rightFound
		}
	}

	calibrationVal, _ := strconv.Atoi(strconv.Itoa(left) + strconv.Itoa(right))

	return int(calibrationVal), nil
}
