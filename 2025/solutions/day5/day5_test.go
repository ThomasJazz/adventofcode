package day5

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	// Add tests here
	result := PartOne()
	println(result)
}

func TestPartTwo(t *testing.T) {
	// Add tests here
	result := PartTwo()
	println(result)
}

// 3-5
// 10-14
// 16-20
// 12-18
// ->
// 3-5
// 10-20
func TestCombineRanges(t *testing.T) {
	// Add tests here
	// PartTwo()
	idRange1 := IdRange{
		3, 5,
	}
	idRange2 := IdRange{
		10, 14,
	}
	idRange3 := IdRange{
		16, 20,
	}
	idRange4 := IdRange{
		12, 18,
	}
	ranges := []IdRange{idRange1, idRange2, idRange3, idRange4}
	merged := CombineRanges(ranges)
	for _, r := range merged {
		println(r.start, "-", r.end)
	}

	// Expected output:
	// 3 - 5
	// 10 - 20
	if len(merged) != 2 {
		t.Errorf("Expected 2 merged ranges, got %d", len(merged))
	}
	if merged[0].start != 3 || merged[0].end != 5 {
		t.Errorf("First merged range incorrect, got %d-%d", merged[0].start, merged[0].end)
	}
	if merged[1].start != 10 || merged[1].end != 20 {
		t.Errorf("Second merged range incorrect, got %d-%d", merged[1].start, merged[1].end)
	}
}
