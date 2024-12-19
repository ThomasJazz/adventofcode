package day08

import (
	"fmt"

	"github.com/thomasjazz/adventofcode/lib"
)

var (
	filename = "../../input/input-day08.txt"
)

func PartOne() int {
	fmt.Println("Day 08 solution pt 1")
	lines := lib.ReadInput(filename)
	antennas := GetAntennaFrequencies(lines)
	antiNodes := CalculateAntiNodes(antennas, len(lines), len(lines[0]))
	runes := InsertAntiNodes(lines, antiNodes)
	return CountAntiNodes(runes)
}

func PartTwo() {
	fmt.Println("Day 08 solution pt 2")
	//lines := lib.ReadInput(filename)
}

func GetAntennaFrequencies(lines []string) map[byte][]lib.Coordinate {
	antennas := make(map[byte][]lib.Coordinate)

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] == '.' {
				continue
			}
			if antennas[lines[i][j]] == nil {
				antennas[lines[i][j]] = make([]lib.Coordinate, 0)
			}
			antennas[lines[i][j]] = append(antennas[lines[i][j]], lib.Coordinate{Row: i, Col: j})
		}

	}

	return antennas
}

func CalculateAntiNodes(frequencies map[byte][]lib.Coordinate, maxRow int, maxCol int) []lib.Coordinate {
	var antiNodes []lib.Coordinate

	for _, coord := range frequencies {
		for i := 0; i < len(coord); i++ {
			for j := i + 1; j < len(coord); j++ {
				diff := getCoordinateDiff(coord[i], coord[j])
				nodea, nodeb := calculateAntiNode(coord[i], coord[j], diff)

				// If out of bounds, skip
				if !(nodea.Row >= maxRow || nodea.Col >= maxCol || nodea.Row < 0 || nodea.Col < 0) {
					antiNodes = append(antiNodes, nodea)
				}

				if !(nodeb.Row >= maxRow || nodeb.Col >= maxCol || nodeb.Row < 0 || nodeb.Col < 0) {
					antiNodes = append(antiNodes, nodeb)
				}
			}
		}
	}

	return antiNodes
}

func calculateAntiNode(a lib.Coordinate, b lib.Coordinate, diff lib.Coordinate) (lib.Coordinate, lib.Coordinate) {
	antiA := lib.Coordinate{Row: a.Row + diff.Row, Col: a.Col + diff.Col}
	antiB := lib.Coordinate{Row: b.Row - diff.Row, Col: b.Col - diff.Col}

	return antiA, antiB
}

func getCoordinateDiff(a lib.Coordinate, b lib.Coordinate) lib.Coordinate {
	rowDiff := a.Row - b.Row
	colDiff := a.Col - b.Col
	return lib.Coordinate{Row: rowDiff, Col: colDiff}
}

func InsertAntiNodes(lines []string, antiNodes []lib.Coordinate) [][]rune {
	runes := lib.ConvertListToRunes(lines)

	for _, node := range antiNodes {
		runes[node.Row][node.Col] = '#'
	}
	lib.PrintGrid(runes)
	return runes
}

func CountAntiNodes(runes [][]rune) int {
	count := 0
	for i := 0; i < len(runes); i++ {
		for j := 0; j < len(runes[i]); j++ {
			if runes[i][j] == '#' {
				count++
			}
		}
	}
	return count
}
