package day01

import (
	"sort"
	"strconv"
	"strings"

	"github.com/thomasjazz/adventofcode/lib"
)

var (
	filename string = "../../input/input-day01.txt"
)

// Calculate total difference between list 1 and 2 sorted
func PartOne() int {
	lines := lib.ReadInput(filename)

	var listOne []int
	var listTwo []int

	for _, line := range lines {
		arr := strings.Split(line, "   ")

		value1, _ := strconv.Atoi(arr[0])
		value2, _ := strconv.Atoi(arr[1])

		listOne = append(listOne, value1)
		listTwo = append(listTwo, value2)
	}

	sort.Ints(listOne)
	sort.Ints(listTwo)

	totalDiff := 0

	for i := 0; i < len(listOne); i++ {
		diff := lib.AbsDiffInt(listOne[i], listTwo[i])
		totalDiff += diff
	}

	return totalDiff
}

// Calculate similarity score of list 1 & 2
// Similarity = list 1 value * number of occurrences in list 2
func PartTwo() int {
	lines := lib.ReadInput(filename)

	similarityScore := 0
	var freqMap map[int]int = make(map[int]int)
	var list1 []int

	for _, line := range lines {
		values := strings.Split(line, "   ")
		value1, _ := strconv.Atoi(values[0])
		value2, _ := strconv.Atoi(values[1])

		if _, ok := freqMap[value2]; ok {
			freqMap[value2]++
		} else {
			freqMap[value2] = 1
		}

		list1 = append(list1, value1)
	}

	for i := 0; i < len(list1); i++ {
		if _, ok := freqMap[list1[i]]; ok {
			// Multiply number from list 1 by number of occurrences in list 2
			similarityScore += list1[i] * freqMap[list1[i]]
		}
	}

	return similarityScore
}
