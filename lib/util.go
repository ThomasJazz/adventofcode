// Common utility functions
package lib

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadInput(filepath string) []string {
	var lines []string
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal("Error reading file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func AbsDiffInt(x int, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func RemoveIndex(s []int, i int) []int {
	return append(s[:i], s[i+1:]...)
}

type Coordinate struct {
	Row int
	Col int
}

func PrintLines(lines []string) {
	for i := 0; i < len(lines); i++ {
		fmt.Println(lines[i])
	}
}

func ConvertListToRunes(lines []string) [][]rune {
	var grid [][]rune
	for i := 0; i < len(lines); i++ {
		grid = append(grid, []rune(lines[i]))
	}
	return grid
}

func PrintGrid(grid [][]rune) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			print(string(grid[i][j]))
		}
		println()
	}
}
