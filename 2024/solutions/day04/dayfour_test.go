package day04

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	expected := 2646
	response := PartOne()
	println(response)

	if response != expected {
		t.Fatalf(`Expected value: %d, Actual value: %d`, expected, response)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 0
	response := PartTwo()
	println(response)

	if response != expected {
		t.Fatalf(`Expected value: %d, Actual value: %d`, expected, response)
	}
}

func TestScanForMasMas_oneX(t *testing.T) {
	lines := []string{
		"XMXSX",
		"XXAXX",
		"XMXSX",
	}

	expected := 1
	response := ScanForMasMas(lines, 1, 2)

	if response != expected {
		t.Fatalf(`Expected value: %d, Actual value: %d`, expected, response)
	}
}

func TestScanForMasMas_noX(t *testing.T) {
	lines := []string{
		"XMXSX",
		"XMASX",
		"XXXXX",
	}

	expected := 0
	response := ScanForMasMas(lines, 1, 0)

	if response != expected {
		t.Fatalf(`Expected value: %d, Actual value: %d`, expected, response)
	}
}
