package day02

import (
	"strconv"
	"strings"

	"github.com/thomasjazz/adventofcode/lib"
)

var (
	filename = "../../input/input-day02.txt"
	maxDiff  = 3
	minDiff  = 1
)

type LevelInput struct {
	PrevNum int
	CurrNum int
	Order   int
}

type LevelInfo struct {
	Diff  int
	Order int
}

// The levels are either all increasing or all decreasing.
// Any two adjacent levels differ by at least one and at most three.
func PartOne() int {
	lines := lib.ReadInput(filename)

	safeLines := 0

	for _, line := range lines {
		levels := strings.Split(line, " ")

		if len(levels) == 1 {
			safeLines++
			continue
		}

		ordering := -1 // 0 will be decreasing, 1 will be increasing
		matchesCriteria := true

		// Work through the line and check if it matches criteria
		for i := 1; i < len(levels); i++ {
			prevNum, _ := strconv.Atoi(levels[i-1])
			num, _ := strconv.Atoi(levels[i])

			// Set max and min
			info := getLevelInfo(LevelInput{PrevNum: prevNum, CurrNum: num, Order: ordering})
			matchesCriteria = evaluateLevelInfo(ordering, info)

			if !matchesCriteria {
				break
			}

			ordering = info.Order
		}

		if matchesCriteria {
			safeLines++
		}
	}

	return safeLines
}

func PartTwo() int {
	lines := lib.ReadInput(filename)
	safeLines := 0

	for _, line := range lines {
		levelsRemoved := 0
		levels := strings.Split(line, " ")

		if len(levels) == 1 {
			safeLines++
			continue
		}

		ordering := -1
		matchesCriteria := true
		skipLast := 0

		for i := 1; i < len(levels); i++ {
			// If we've made it this far without removing a level, we are good to go
			if i == len(levels)-1 && levelsRemoved == 0 {
				break
			}

			lastNum, _ := strconv.Atoi(levels[i-(1+skipLast)])
			currNum, _ := strconv.Atoi(levels[i])

			info := getLevelInfo(LevelInput{PrevNum: lastNum, CurrNum: currNum, Order: ordering})
			matchesCriteria = evaluateLevelInfo(ordering, info)

			if !matchesCriteria {
				if levelsRemoved > 0 {
					break
				}
				levelsRemoved, skipLast = 1, 1
				matchesCriteria = true
			} else {
				skipLast = 0
				ordering = info.Order
			}
		}

		if matchesCriteria {
			safeLines++
		}
	}

	return safeLines
}

func getLevelInfo(input LevelInput) LevelInfo {
	order := -1

	if input.PrevNum < input.CurrNum {
		order = 1
	} else if input.PrevNum > input.CurrNum {
		order = 0
	}

	diff := lib.AbsDiffInt(input.CurrNum, input.PrevNum)

	return LevelInfo{Diff: diff, Order: order}
}

func evaluateLevelInfo(lastOrder int, info LevelInfo) bool {
	if lastOrder != -1 && lastOrder != info.Order {
		return false
	}

	if info.Diff > maxDiff || info.Diff < minDiff {
		return false
	}

	return true
}
