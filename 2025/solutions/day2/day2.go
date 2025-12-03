package day2

import (
	"fmt"
	"strconv"

	"github.com/thomasjazz/adventofcode/lib"
)

var (
	filename = "../../input/input-day2.txt"
)

type IdRange struct {
	bottom int
	top    int
}

func PartOne() int {
	fmt.Println("Day 2 solution pt 1")
	lines := lib.ReadInputWithDelimiter(filename, ",")
	idRanges := parseRanges(lines)
	var invalidNumbers []int

	for _, idRange := range idRanges {
		for count := idRange.bottom; count <= idRange.top; count++ {
			numberStr := strconv.Itoa(count)

			// Only even length IDs can be invalid
			if len(numberStr)%2 == 0 {
				isValid := false

				// Loop through half the string
				halfLength := len(numberStr) / 2
				for i := 0; i < halfLength; i++ {
					if !(numberStr[i] == numberStr[halfLength+i]) {
						isValid = true
						break
					}
				}

				if !isValid {
					invalidNumbers = append(invalidNumbers, count)
				}
			}
		}
	}

	return sumInvalidNums(invalidNumbers)
}

func PartTwo() int {
	fmt.Println("Day 2 solution pt 2")
	lines := lib.ReadInputWithDelimiter(filename, ",")
	idRanges := parseRanges(lines)
	var invalidNumbers []int

	// I think we need to do sliding window here
	for _, idRange := range idRanges {
		for count := idRange.bottom; count <= idRange.top; count++ {
			numberStr := strconv.Itoa(count)

			// Examples
			// 1215 1215 1215   = 6x1s, 3x2s, 3x5s
			// 1111             = 4x1s
			// 7868 7868        = 2x7s, 4x8s, 2x6s
			var substrings []string

			// Collect all substrings up until half way through the number
			for right := 1; right <= (len(numberStr) / 2); right++ {
				// Can't be a repeating sequence if it doesn't evenly fit into the parent string length
				if len(numberStr)%(right) == 0 {
					substrings = append(substrings, numberStr[0:right])
				}
			}

			for _, substring := range substrings {
				isValid := false

				// Check if the substring repeats. If any don't match, it is valid
				for i := len(substring); i < len(numberStr); i += len(substring) {
					if substring != numberStr[i:i+len(substring)] {
						isValid = true
						break
					}
				}

				if !isValid {
					invalidNumbers = append(invalidNumbers, count)
					break
				}
			}
		}
	}

	return sumInvalidNums(invalidNumbers)
}

func sumInvalidNums(invalidNums []int) int {
	var sum int
	for _, invalidNum := range invalidNums {
		sum += invalidNum
	}

	return sum
}

func parseRanges(lines []string) []IdRange {
	var ranges []IdRange
	for _, val := range lines {
		var bottom, top int
		n, err := fmt.Sscanf(val, "%d-%d", &bottom, &top)
		if err == nil && n == 2 {
			ranges = append(ranges, IdRange{bottom: bottom, top: top})
		}
	}
	return ranges
}
