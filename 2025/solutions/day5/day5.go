package day5

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/thomasjazz/adventofcode/lib"
)

var (
	filename = "input-day5.txt"
)

type IdRange struct {
	start int64
	end   int64
}
type IngredientDatabase struct {
	freshIdRanges        []IdRange
	availableIngredients []int64
}

func PartOne() int {
	start := time.Now()
	fmt.Println("Day 5 solution pt 1")
	lines := lib.ReadInput(filename)
	db := processLines(lines)
	var freshIngredients []int64

	for _, ingredient := range db.availableIngredients {
		isFresh := false
		for _, idRange := range db.freshIdRanges {
			if ingredient >= idRange.start && ingredient <= idRange.end {
				isFresh = true
				break
			}
		}

		if isFresh {
			freshIngredients = append(freshIngredients, ingredient)
		}
	}
	fmt.Printf("Finished in %s\n", time.Since(start))
	return len(freshIngredients)
}

// 541.875µs
func PartTwo() int64 {
	start := time.Now()
	fmt.Println("Day 5 solution pt 2")
	// It would obviously be faster to process as we read the file instead of reading into memory then processing
	// but i kinda like it this way since we dont care about speed that much
	lines := lib.ReadInput(filename)
	ranges := processRanges(lines)

	mergedRanges := CombineRanges(ranges)

	fmt.Printf("Finished in %s\n", time.Since(start))
	return SumRanges(mergedRanges)
}

func CombineRanges(ranges []IdRange) []IdRange {
	var mergedRanges []IdRange

	anyMerges := false
	for _, idRange := range ranges {
		if len(mergedRanges) == 0 {
			mergedRanges = append(mergedRanges, idRange)
			continue
		}

		index := slices.IndexFunc(mergedRanges, func(r IdRange) bool {
			return (r.start >= idRange.start && r.start <= idRange.end) || (r.end >= idRange.start && r.end <= idRange.end) ||
				(idRange.start >= r.start && idRange.start <= r.end) || (idRange.end >= r.start && idRange.end <= r.end)
		})

		if index == -1 {
			mergedRanges = append(mergedRanges, idRange)
		} else {
			mergedRange := IdRange{
				start: lib.MinInt64(mergedRanges[index].start, idRange.start),
				end:   lib.MaxInt64(mergedRanges[index].end, idRange.end),
			}
			mergedRanges[index] = mergedRange
			anyMerges = true
		}
	}

	if anyMerges {
		return CombineRanges(mergedRanges)
	} else {
		return mergedRanges
	}
}

func SumRanges(ranges []IdRange) int64 {
	var total int64 = 0
	for _, r := range ranges {
		total += (r.end - r.start + 1)
	}
	return total
}

func processLines(lines []string) IngredientDatabase {
	var rangeLines, ingredients []string

	// fileSegment 0 = ranges. fileSegment 1 = individual IDs
	fileSegment := 0
	for _, line := range lines {
		if line == "" {
			fileSegment = 1
			continue
		}

		if fileSegment == 0 {
			rangeLines = append(rangeLines, line)
		} else {
			ingredients = append(ingredients, line)
		}
	}

	return IngredientDatabase{
		freshIdRanges:        processRanges(rangeLines),
		availableIngredients: processIngredients(ingredients),
	}
}

func processRanges(lines []string) []IdRange {
	var idRanges []IdRange
	for _, line := range lines {
		if line == "" {
			return idRanges
		}

		ranges := strings.Split(line, "-")

		idRanges = append(idRanges, IdRange{
			start: lib.StringToInt64(ranges[0]),
			end:   lib.StringToInt64(ranges[1]),
		})

	}

	return idRanges
}

func processIngredients(lines []string) []int64 {
	var ingredients []int64
	for _, line := range lines {
		ingredients = append(ingredients, lib.StringToInt64(line))
	}
	return ingredients
}
