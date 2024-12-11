package day06

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	expected := 4789
	actual := PartOne()
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func TestPartOneExample(t *testing.T) {
	expected := 41
	input := make([]string, 0)
	input = append(input, "....#.....")
	input = append(input, ".........#")
	input = append(input, "..........")
	input = append(input, "..#.......")
	input = append(input, ".......#..")
	input = append(input, "..........")
	input = append(input, ".#..^.....")
	input = append(input, "........#.")
	input = append(input, "#.........")
	input = append(input, "......#...")

	grid := MakeBoard(input)
	guardLocation, err := LocateGuard(grid)

	if err != nil {
		t.Errorf("Error locating guard: %v", err)
	}

	actual := MarkVisited(grid, guardLocation)
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}
