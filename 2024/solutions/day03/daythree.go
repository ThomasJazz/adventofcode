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
	totalMultiplied := 0
	dontMultiplied := 0
	//doEnabled := true

	regexp := regexp.MustCompile(regexpMul)

	for _, line := range lines {
		doArr := strings.Split(line, "do()")

		for _, doLine := range doArr {
			excludeDont := strings.Split(doLine, "don't()")

			for _, dontLine := range excludeDont {
				matches := regexp.FindAllStringSubmatch(dontLine, -1)

				for _, match := range matches {
					leftNum, _ := strconv.Atoi(match[1])
					rightNum, _ := strconv.Atoi(match[2])

					dontMultiplied += leftNum * rightNum
				}
			}

			totalMatches := regexp.FindAllStringSubmatch(doLine, -1)

			for _, match := range totalMatches {
				leftNum, _ := strconv.Atoi(match[1])
				rightNum, _ := strconv.Atoi(match[2])

				totalMultiplied += leftNum * rightNum
			}
		}

	}

	return totalMultiplied - dontMultiplied
}
