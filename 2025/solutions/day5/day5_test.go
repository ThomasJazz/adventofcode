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

	sum := SumRanges(merged)
	if sum != 14 {
		t.Errorf("Expected sum of 14, got %d", sum)
	}
}

func TestCombineRangesInner(t *testing.T) {
	idRange1 := IdRange{
		10, 20,
	}
	idRange2 := IdRange{
		12, 15,
	}

	ranges := []IdRange{idRange2, idRange1}
	merged := CombineRanges(ranges)
	// expect one range of 10-20
	if len(merged) != 1 {
		t.Errorf("Expected 1 merged range, got %d", len(merged))
	}
	if merged[0].start != 10 || merged[0].end != 20 {
		t.Errorf("Merged range incorrect, got %d-%d", merged[0].start, merged[0].end)
	}

	sum := SumRanges(merged)
	if sum != 11 {
		t.Errorf("Expected sum of 11, got %d", sum)
	}
}

func TestCombineRangesFromInput(t *testing.T) {
	// 73622261592470-78368864856562
	idRange1 := IdRange{
		73622261592470, 78368864856562,
	}
	//78368864856562-78368864856562
	idRange2 := IdRange{
		78368864856562, 78368864856562,
	}

	ranges := []IdRange{idRange1, idRange2}
	merged := CombineRanges(ranges)
	// expect one range of 73622261592470-78368864856562
	if len(merged) != 1 {
		t.Errorf("Expected 1 merged range, got %d", len(merged))
	}
	if merged[0].start != 73622261592470 || merged[0].end != 78368864856562 {
		t.Errorf("Merged range incorrect, got %d-%d", merged[0].start, merged[0].end)
	}

	sum := SumRanges(merged)
	if sum != 4746603264093 {
		t.Errorf("Expected sum of 4746603264093, got %d", sum)
	}
	println("Sum:", sum)
}

func TestCombineIdenticalRanges(t *testing.T) {
	idRange1 := IdRange{
		100, 200,
	}
	idRange2 := IdRange{
		100, 200,
	}

	ranges := []IdRange{idRange1, idRange2}
	merged := CombineRanges(ranges)
	// expect one range of 100-200
	if len(merged) != 1 {
		t.Errorf("Expected 1 merged range, got %d", len(merged))
	}
	if merged[0].start != 100 || merged[0].end != 200 {
		t.Errorf("Merged range incorrect, got %d-%d", merged[0].start, merged[0].end)
	}

	sum := SumRanges(merged)
	if sum != 101 {
		t.Errorf("Expected sum of 101, got %d", sum)
	}
}
