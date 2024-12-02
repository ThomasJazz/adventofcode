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
	expected := 692
	safeLevels := PartTwo()

	println(safeLevels)

	if safeLevels != expected {
		t.Fatalf(`Expected value: %d, Actual value: %d`, expected, safeLevels)
	}
}

func TestCheckLine(t *testing.T) {
	levels := []int{1, 2, 3, 4, 5}
	if !CheckLine(levels) {
		t.Fatalf(`Expected value: %t, Actual value: %t`, true, false)
	}

	levels = []int{5, 4, 3, 2, 1}
	if !CheckLine(levels) {
		t.Fatalf(`Expected value: %t, Actual value: %t`, true, false)
	}

	levels = []int{1, 2, 3, 5, 6}
	if !CheckLine(levels) {
		t.Fatalf(`Expected value: %t, Actual value: %t`, true, false)
	}

	levels = []int{1, 2, 3, 5, 4}
	if CheckLine(levels) {
		t.Fatalf(`Expected value: %t, Actual value: %t`, false, true)
	}

	levels = []int{1, 2, 3, 4, 8}
	if CheckLine(levels) {
		t.Fatalf(`Expected value: %t, Actual value: %t`, false, true)
	}
}

func TestApplyDampener(t *testing.T) {
	levels := []int{1, 2, 3, 4, 5}
	if !ApplyDampener(levels) {
		t.Fatalf(`Expected value: %t, Actual value: %t`, true, false)
	}

	levels = []int{5, 4, 3, 2, 1}
	if !ApplyDampener(levels) {
		t.Fatalf(`Expected value: %t, Actual value: %t`, true, false)
	}

	levels = []int{1, 1, 2, 3, 4}
	if !ApplyDampener(levels) {
		t.Fatalf(`Expected value: %t, Actual value: %t`, true, false)
	}

	levels = []int{1, 2, 3, 1, 5}
	if !ApplyDampener(levels) {
		t.Fatalf(`Expected value: %t, Actual value: %t`, true, false)
	}
}
