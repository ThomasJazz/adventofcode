package day05

import (
	"strconv"
	"strings"

	"github.com/thomasjazz/adventofcode/lib"
)

var (
	filename1 = "../../input/input-day05-1.txt"
	filename2 = "../../input/input-day05-2.txt"
)

func PartOne() int {
	rules := lib.ReadInput(filename1)
	pageUpdates := lib.ReadInput(filename2)
	middleSum := 0

	// Each key must go before any of the values in the page update sequence
	rulesMap := getRulesMap(rules)
	pageUpdateSequences := getPageUpdateSequences(pageUpdates)

	correctSequences := getCorrectSequences(pageUpdateSequences, rulesMap)
	for _, sequence := range correctSequences {
		middleIndex := len(sequence) / 2
		middleNum, _ := strconv.Atoi(sequence[middleIndex])

		middleSum += middleNum
	}
	return middleSum
}

// getRulesMap returns a map of rules from the input
func getRulesMap(rules []string) map[string]map[string]struct{} {
	rulesMap := make(map[string]map[string]struct{})

	for _, rule := range rules {
		ruleParts := strings.Split(rule, "|")
		if rulesMap[ruleParts[0]] == nil {
			rulesMap[ruleParts[0]] = make(map[string]struct{})
		}

		rulesMap[ruleParts[0]][ruleParts[1]] = struct{}{}
	}

	return rulesMap
}

// getPageUpdateSequences returns a slice of page update sequences from the input
func getPageUpdateSequences(pageUpdates []string) [][]string {
	pageUpdateSequences := make([][]string, 0)

	for _, pageUpdate := range pageUpdates {
		pageUpdateSequences = append(pageUpdateSequences, strings.Split(pageUpdate, ","))
	}

	return pageUpdateSequences
}

// getCorrectSequences returns a slice of sequences that satisfy the rules
// A sequence satisfies the rules if each key goes before any of it's values
func getCorrectSequences(sequences [][]string, rulesMap map[string]map[string]struct{}) [][]string {
	correctSequences := make([][]string, 0)
	for _, updateSequence := range sequences {
		addSequence := true

	sequenceLoop:
		for i := len(updateSequence) - 1; i >= 0; i-- {
			for j := i - 1; j >= 0; j-- {
				if _, ok := rulesMap[updateSequence[i]][updateSequence[j]]; ok {
					addSequence = false
					break sequenceLoop
				}
			}
		}

		if addSequence {
			correctSequences = append(correctSequences, updateSequence)
		}
	}

	return correctSequences
}
