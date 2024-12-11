package day06

import (
	"errors"
	"log"

	"github.com/thomasjazz/adventofcode/lib"
)

type Coordinate struct {
	row int
	col int
}

var (
	filename = "../../input/input-day06.txt"
)

func PartOne() int {
	lines := lib.ReadInput(filename)
	grid := MakeBoard(lines)

	guardStart, err := LocateGuard(grid)
	if err != nil {
		log.Fatal(err)
	}

	return MarkVisited(grid, guardStart)
}

func MarkVisited(grid [][]rune, startCoord Coordinate) int {
	inRoom := true
	row := startCoord.row
	col := startCoord.col

	for inRoom {
		//PrintGrid(grid)
		// Move up
		if grid[row][col] == '^' {
			if row == 0 {
				grid[row][col] = 'X'
				inRoom = false
			} else if grid[row-1][col] == '#' {
				grid[row][col] = '>'
			} else {
				grid[row][col] = 'X'
				row--
				grid[row][col] = '^'
			}
		}
		// Move down
		if grid[row][col] == 'v' {
			if row == len(grid)-1 {
				grid[row][col] = 'X'
				inRoom = false
			} else if grid[row+1][col] == '#' {
				grid[row][col] = '<'
			} else {
				grid[row][col] = 'X'
				row++
				grid[row][col] = 'v'
			}
		}
		// Move left
		if grid[row][col] == '<' {
			if col == 0 {
				grid[row][col] = 'X'
				inRoom = false
			} else if grid[row][col-1] == '#' {
				grid[row][col] = '^'
			} else {
				grid[row][col] = 'X'
				col--
				grid[row][col] = '<'
			}
		}
		// Move right
		if grid[row][col] == '>' {
			if col == len(grid[row])-1 {
				grid[row][col] = 'X'
				inRoom = false
			} else if grid[row][col+1] == '#' {
				grid[row][col] = 'v'
			} else {
				grid[row][col] = 'X'
				col++
				grid[row][col] = '>'
			}
		}
	}

	grid[row][col] = 'X'

	PrintGrid(grid)
	return countXs(grid)
}

func MakeBoard(lines []string) [][]rune {
	grid := make([][]rune, len(lines))

	for i, line := range lines {
		grid[i] = []rune(line)
	}

	return grid
}

func LocateGuard(grid [][]rune) (Coordinate, error) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '^' {
				return Coordinate{i, j}, nil
			}
		}
	}

	return Coordinate{-1, -1}, errors.New("guard not found")
}

func PrintGrid(grid [][]rune) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			print(string(grid[i][j]))
		}
		println()
	}
}

func GridToString(grid [][]rune) string {
	str := ""
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			str += string(grid[i][j])
		}
		str += "\n"
	}
	return str
}

func countXs(grid [][]rune) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 'X' {
				count++
			}
		}
	}
	return count
}
