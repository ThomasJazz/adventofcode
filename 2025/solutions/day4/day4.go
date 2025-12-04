package day4

import (
	"fmt"

	"github.com/thomasjazz/adventofcode/lib"
)

var (
	filename = "input-day4.txt"
)

type Tile struct {
	row   int
	col   int
	value string
}

type Board struct {
	tiles [][]Tile
}

func PartOne() int {
	fmt.Println("Day 4 solution pt 1")
	lines := lib.ReadInput(filename)
	board := processBoard(lines)

	accessiblePaper := 0

	for _, row := range board.tiles {
		for _, tile := range row {
			if tile.value == "@" && canBeAccessed(board, tile) {
				accessiblePaper++
			}
		}
	}

	return accessiblePaper
}

func PartTwo() int {
	fmt.Println("Day 4 solution pt 2")
	lines := lib.ReadInput(filename)
	board := processBoard(lines)

	paperRemoved := 0
	boardIter := 0
	for {
		anyPaperRemoved := false
		println("Iterating over board:", boardIter)

		for _, row := range board.tiles {
			for _, tile := range row {
				if tile.value == "@" && canBeAccessed(board, tile) {
					paperRemoved++
					board.tiles[tile.row][tile.col] = Tile{
						row:   tile.row,
						col:   tile.col,
						value: ".",
					}
					anyPaperRemoved = true
				}
			}
		}
		boardIter++
		if !anyPaperRemoved {
			break
		}
	}

	return paperRemoved
}

func canBeAccessed(board Board, tile Tile) bool {
	adjacent := getAdjacentTiles(board, tile.row, tile.col)
	paperAdjacentCount := 0

	for _, tile := range adjacent {
		if tile.value == "@" {
			paperAdjacentCount++
		}
	}

	return paperAdjacentCount < 4
}

func getAdjacentTiles(board Board, row int, col int) []Tile {
	var adjacent []Tile

	startRow := lib.MaxInt(0, row-1)
	startCol := lib.MaxInt(0, col-1)
	endRow := lib.MinInt(len(board.tiles)-1, row+1)
	endCol := lib.MinInt(len(board.tiles[0])-1, col+1)

	for r := startRow; r <= endRow; r++ {
		for c := startCol; c <= endCol; c++ {
			if r == row && c == col {
				continue
			}

			adjacent = append(adjacent, board.tiles[r][c])
		}
	}

	return adjacent
}

func processBoard(lines []string) Board {
	var board Board
	board.tiles = make([][]Tile, len(lines))

	for row := 0; row < len(lines); row++ {
		board.tiles[row] = make([]Tile, len(lines[row]))

		for col := 0; col < len(lines[row]); col++ {
			tile := Tile{
				row:   row,
				col:   col,
				value: string(lines[row][col]),
			}
			board.tiles[row][col] = tile
		}
	}

	return board
}
