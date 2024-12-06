// Day 4: Ceres Search
package day04

import (
	"github.com/thomasjazz/adventofcode/lib"
)

var (
	filename = "../../input/input-day04.txt"
)

func PartOne() int {
	lines := lib.ReadInput(filename)
	totalXmas := 0

	for row := 0; row < len(lines); row++ {
		for col := 0; col < len(lines[row]); col++ {
			char := lines[row][col]
			if char == 'X' {
				totalXmas += ScanForXmas(lines, row, col)
			}
		}
	}

	return totalXmas
}

// Look for XMAS in all directions
func ScanForXmas(lines []string, row int, col int) int {
	xmasFound := 0
	maxRows := len(lines)
	maxCols := len(lines[0])

	// Scan in all directions
	if row < maxRows-3 {
		if lines[row+1][col] == 'M' && lines[row+2][col] == 'A' && lines[row+3][col] == 'S' {
			xmasFound++
		}
	}
	if row > 2 {
		if lines[row-1][col] == 'M' && lines[row-2][col] == 'A' && lines[row-3][col] == 'S' {
			xmasFound++
		}
	}
	if col < maxCols-3 {
		if lines[row][col+1] == 'M' && lines[row][col+2] == 'A' && lines[row][col+3] == 'S' {
			xmasFound++
		}
	}
	if col > 2 {
		if lines[row][col-1] == 'M' && lines[row][col-2] == 'A' && lines[row][col-3] == 'S' {
			xmasFound++
		}
	}
	if row < maxRows-3 && col < maxCols-3 {
		if lines[row+1][col+1] == 'M' && lines[row+2][col+2] == 'A' && lines[row+3][col+3] == 'S' {
			xmasFound++
		}
	}
	if row > 2 && col > 2 {
		if lines[row-1][col-1] == 'M' && lines[row-2][col-2] == 'A' && lines[row-3][col-3] == 'S' {
			xmasFound++
		}
	}
	if row < maxRows-3 && col > 2 {
		if lines[row+1][col-1] == 'M' && lines[row+2][col-2] == 'A' && lines[row+3][col-3] == 'S' {
			xmasFound++
		}
	}
	if row > 2 && col < maxCols-3 {
		if lines[row-1][col+1] == 'M' && lines[row-2][col+2] == 'A' && lines[row-3][col+3] == 'S' {
			xmasFound++
		}
	}

	return xmasFound
}

func PartTwo() int {
	lines := lib.ReadInput(filename)
	totalXmas := 0

	for row := 1; row < len(lines)-1; row++ {
		for col := 1; col < len(lines[row])-1; col++ {
			char := lines[row][col]
			if char == 'A' {
				totalXmas += ScanForMasMas(lines, row, col)
			}
		}
	}

	return totalXmas
}

// Look for pattern that looks like this (can be rotated):
// S.S
// .A.
// M.M
func ScanForMasMas(lines []string, row int, col int) int {
	// S's on top left and top right, M's on bottom left and bottom right
	if lines[row-1][col-1] == 'S' && lines[row-1][col+1] == 'S' && lines[row+1][col-1] == 'M' && lines[row+1][col+1] == 'M' {
		return 1
		// M's on top left and top right, S's on bottom left and bottom right
	} else if lines[row-1][col-1] == 'M' && lines[row-1][col+1] == 'M' && lines[row+1][col-1] == 'S' && lines[row+1][col+1] == 'S' {
		return 1
		// M's on top left and bottom right, S's on top right and bottom left
	} else if lines[row-1][col-1] == 'M' && lines[row-1][col+1] == 'S' && lines[row+1][col-1] == 'M' && lines[row+1][col+1] == 'S' {
		return 1
		// S's on top left and bottom right, M's on top right and bottom left
	} else if lines[row-1][col-1] == 'S' && lines[row-1][col+1] == 'M' && lines[row+1][col-1] == 'S' && lines[row+1][col+1] == 'M' {
		return 1
	}
	return 0
}
