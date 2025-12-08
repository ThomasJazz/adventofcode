package day7

import (
	"fmt"

	"github.com/thomasjazz/adventofcode/lib"
)

var (
	// filename = "input-day7.txt"
	filename = "test-input.txt"
	grid     [][]lib.Tile
)

func PartOne() int {
	fmt.Println("Day 7 solution pt 1")
	lines := lib.ReadInput(filename)
	grid = lib.ReadLinesToGrid(lines)

	for _, tile := range grid[0] {
		if tile.Value == "S" {
			return SplitManifolds(grid, tile.Row+1, tile.Column)
		}
	}

	return 0
}

func SplitManifolds(grid [][]lib.Tile, row, col int) int {
	if row >= len(grid)-1 || col >= len(grid[0])-1 {
		return 0
	} else if grid[row][col].Value == "|" { // Already traversed here
		return 0
	} else if grid[row][col].Value == "." {
		grid[row][col].Value = "|"
		return SplitManifolds(grid, row+1, col)
	} else if grid[row][col].Value == "^" {
		return SplitManifolds(grid, row+1, col-1) + SplitManifolds(grid, row+1, col+1) + 1
	}

	return 0
}

func PartTwo() int {
	fmt.Println("Day 7 solution pt 2")
	lines := lib.ReadInput(filename)
	grid = lib.ReadLinesToGrid(lines)

	for _, tile := range grid[0] {
		if tile.Value == "S" {
			return QuantumSplit(tile.Row+1, tile.Column)
		}
	}

	return 0
}

// This is the correct algorithm, but we need to memoize or the tree will be too big to complete
func QuantumSplit(row, col int) int {
	if row >= len(grid)-1 || col >= len(grid[0])-1 {
		return 1
	} else if grid[row][col].Value == "." {
		// grid[row][col].Value = "|"
		return QuantumSplit(row+1, col)
	} else if grid[row][col].Value == "^" {
		return QuantumSplit(row+1, col-1) + QuantumSplit(row+1, col+1)
	}

	return 0
}
