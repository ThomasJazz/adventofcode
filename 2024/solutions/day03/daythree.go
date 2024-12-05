package day03

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/thomasjazz/adventofcode/lib"
)

var (
	filename  = "../../input/input-day03.txt"
	regexpMul = `mul\((\d+),(\d+)\)`
)

func PartOne() int {
	lines := lib.ReadInput(filename)
	totalMultiplied := 0

	// Generate regex to capture the values mul(613,600) but not mul( 613, 600 ) or mul(asd, 600)
	regex := regexp.MustCompile(regexpMul)

	for _, line := range lines {
		matches := regex.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			leftNum, _ := strconv.Atoi(match[1])
			rightNum, _ := strconv.Atoi(match[2])

			totalMultiplied += leftNum * rightNum
		}
	}

	return totalMultiplied
}

// Scan for do() and don't(). Only include mul() values after do() and before don't()
func PartTwo() int {
	lines := lib.ReadInput(filename)
	concatenatedString := strings.Join(lines, "\n")
	totalMultiplied := 0

	regexp := regexp.MustCompile(regexpMul)

	parts := strings.Split(concatenatedString, "do()")
	for _, part := range parts {
		dontParts := strings.Split(part, "don't()")

		// Only multiply the mul() values between do() and don't()
		usePart := dontParts[0]
		matches := regexp.FindAllStringSubmatch(usePart, -1)

		for _, match := range matches {
			leftNum, _ := strconv.Atoi(match[1])
			rightNum, _ := strconv.Atoi(match[2])
			totalMultiplied += leftNum * rightNum
		}
	}

	return totalMultiplied
}
