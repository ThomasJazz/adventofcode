package day08

import (
	"fmt"
	"testing"
)

func TestPartOne(t *testing.T) {
	// Add tests here
	antiNodes := PartOne()
	println(antiNodes)
}

// ............
// ........0...
// .....0......
// .......0....
// ....0.......
// ......A.....
// ............
// ............
// ........A...
// .........A..
// ............
// ............
func TestGetAntennaFrequencies(t *testing.T) {
	sample := make([]string, 0)
	sample = append(sample, "............")
	sample = append(sample, "........0...")
	sample = append(sample, ".....0......")
	sample = append(sample, ".......0....")
	sample = append(sample, "....0.......")
	sample = append(sample, "......A.....")
	sample = append(sample, "............")
	sample = append(sample, "............")
	sample = append(sample, "........A...")
	sample = append(sample, ".........A..")
	sample = append(sample, "............")
	sample = append(sample, "............")

	antennas := GetAntennaFrequencies(sample)
	for k, v := range antennas {
		println(fmt.Sprintf("Frequency: %c", k))
		for i := 0; i < len(v); i++ {
			println(fmt.Sprintf("%d, %d", v[i].Row, v[i].Col))
		}
	}
}

// Desired:
// ..........
// ...#......
// ..........
// ....a.....
// ..........
// .....a....
// ..........
// ......#...
// ..........
// ..........
func TestGetFreqSimple(t *testing.T) {
	sample := make([]string, 0)
	sample = append(sample, "..........")
	sample = append(sample, "..........")
	sample = append(sample, "..........")
	sample = append(sample, "....a.....")
	sample = append(sample, "..........")
	sample = append(sample, ".....a....")
	sample = append(sample, "..........")
	sample = append(sample, "..........")
	sample = append(sample, "..........")
	sample = append(sample, "..........")

	antennas := GetAntennaFrequencies(sample)
	antiNodes := CalculateAntiNodes(antennas, len(sample), len(sample[0]))
	runes := InsertAntiNodes(sample, antiNodes)
	value := CountAntiNodes(runes)
	println(value)
}

// Desired:
// ..........
// ...#......
// #.........
// ....a.....
// ........a.
// .....a....
// ..#.......
// ......#...
// ..........
// ..........
func TestGetFreqSimple2(t *testing.T) {
	sample := make([]string, 0)
	sample = append(sample, "..........")
	sample = append(sample, "..........")
	sample = append(sample, "..........")
	sample = append(sample, "....a.....")
	sample = append(sample, "........a.")
	sample = append(sample, ".....a....")
	sample = append(sample, "..........")
	sample = append(sample, "..........")
	sample = append(sample, "..........")
	sample = append(sample, "..........")

	antennas := GetAntennaFrequencies(sample)
	antiNodes := CalculateAntiNodes(antennas, len(sample), len(sample[0]))
	runes := InsertAntiNodes(sample, antiNodes)
	value := CountAntiNodes(runes)
	println(value)
}

func TestPartTwo(t *testing.T) {
	// Add tests here
}
