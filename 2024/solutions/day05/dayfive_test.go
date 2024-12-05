package day05

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	result := PartOne()
	if result != 6612 {
		t.Errorf("Expected 0, got %d", result)
	}
}

func TestPartTwo(t *testing.T) {
	result := PartTwo()
	if result != 4944 {
		t.Errorf("Expected 0, got %d", result)
	}
}
