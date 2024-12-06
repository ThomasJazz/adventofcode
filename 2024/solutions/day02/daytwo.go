// Day 2: Red-Nosed Reports
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
		intLevels := make([]int, len(levels))

		for i, level := range levels {
			intLevel, _ := strconv.Atoi(level)
			intLevels[i] = intLevel
		}

		if CheckLine(intLevels) {
			safeLines++
		}
	}

	return safeLines
}

func PartTwo() int {
	lines := lib.ReadInput(filename)
	safeLines := 0

	for _, line := range lines {
		levels := strings.Split(line, " ")
		intLevels := make([]int, len(levels))

		for i, level := range levels {
			intLevel, _ := strconv.Atoi(level)
			intLevels[i] = intLevel
		}

		if CheckLine(intLevels) || ApplyDampener(intLevels) {
			safeLines++
		}
	}

	return safeLines
}

func ApplyDampener(levels []int) bool {
	for i := 0; i < len(levels); i++ {
		levelsCopy := make([]int, len(levels))
		copy(levelsCopy, levels)
		levelsCopy = lib.RemoveIndex(levelsCopy, i)

		if CheckLine(levelsCopy) {
			return true
		}
	}

	return false
}

func CheckLine(levels []int) bool {
	if len(levels) == 1 {
		return true
	}

	ordering := -1 // 0 will be decreasing, 1 will be increasing
	matchesCriteria := true

	// Work through the line and check if it matches criteria
	for i := 1; i < len(levels); i++ {
		// Set max and min
		info := getLevelInfo(LevelInput{PrevNum: levels[i-1], CurrNum: levels[i], Order: ordering})
		matchesCriteria = evaluateLevelInfo(ordering, info)

		if !matchesCriteria {
			return false
		}

		ordering = info.Order
	}

	return matchesCriteria
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
