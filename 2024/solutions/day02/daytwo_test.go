package day02

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	expected := 663
	safeLevels := PartOne()

	println(safeLevels)

	if expected != safeLevels {
		t.Fatalf(`Expected value: %d, Actual value: %d`, expected, safeLevels)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 663
	safeLevels := PartTwo()

	println(safeLevels)

	if safeLevels > expected {
		t.Fatalf(`Expected value: %d, Actual value: %d`, expected, safeLevels)
	}
}
