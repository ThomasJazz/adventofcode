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
