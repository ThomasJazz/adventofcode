package day03

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/thomasjazz/adventofcode/lib"
)

var (
	filepath string = "../../input/day_3_input.txt"
	symbols  string = "*&/%=#-+$@"
	numbers  string = "0123456789"
)

func GetCoordinate(row int, col int) string {
	return strconv.Itoa(row) + "," + strconv.Itoa(col)
}

func GetAdjacent(row int, col int) map[string]bool {
	adjacent := make(map[string]bool)

	for i := 0; i < 3; i++ {
		for k := 0; k < 3; k++ {
			coordinate := GetCoordinate(row-1+i, col-1+k)
			adjacent[coordinate] = true
		}
	}

	return adjacent
}

func IsSymbolAdjacent(adjacentSquares [][]string) bool {
	return false
}

func MarkAllSymbolPerimiters(input []string) map[string]bool {
	adjacentSquares := make(map[string]bool)

	// Read input
	for row, line := range input {
		for col, char := range line {
			if strings.Contains(symbols, string(char)) {
				adjacent := GetAdjacent(row, col)

				// Merge map
				for k, v := range adjacent {
					adjacentSquares[k] = v
				}
			}
		}
	}

	return adjacentSquares
}

func GearRatios() {
	partSum := 0
	// TODO: Scan the input for all symbols and add adjacent coordinates to
	// the map
	input := lib.ReadInput(filepath)
	adjacentSquares := MarkAllSymbolPerimiters(input)
	currNum := -1
	numberIsAdjacent := false

	for row, line := range input {
		for col, char := range line {
			// If we are currently examining an int
			if unicode.IsDigit(char) {

				// Check if this is a new number or if it needs to be appended to existing
				if currNum == -1 {
					currNum, _ = strconv.Atoi(string(char))
				} else {
					currNum, _ = strconv.Atoi(strconv.Itoa(currNum) + string(char))
				}

				// If this coordinate is in our map of coordinates adjacent to symbols
				if adjacentSquares[GetCoordinate(row, col)] {
					numberIsAdjacent = true
				}

			} else if currNum != -1 {
				// If it is not a digit and currNum != -1, we have reached the end of the number
				if numberIsAdjacent {
					partSum += currNum
				}

				currNum = -1
				numberIsAdjacent = false
			}

		}
	}

	println(partSum)
	println()
}
