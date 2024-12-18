package day07

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	var expected uint64 = 1153997401072
	actual := PartOne()
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	var expected uint64 = 97902809384118
	actual := PartTwo()
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}
