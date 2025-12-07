package day6

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/thomasjazz/adventofcode/lib"
)

var (
	filename = "input-day6.txt"
)

type CephaloProblem struct {
	index   int
	nums    []int
	operand string
}

func PartOne() int {
	fmt.Println("Day 6 solution pt 1")
	lines := lib.ReadInput(filename)
	start := time.Now()
	problems := getProblemsMap(lines)

	total := 0

	for _, problem := range problems {
		switch problem.operand {
		case "+":
			problemAnswer := 0
			for _, num := range problem.nums {
				problemAnswer += num
			}
			total += problemAnswer
		case "*":
			problemAnswer := 1
			for _, num := range problem.nums {
				problemAnswer *= num
			}

			total += problemAnswer
		}
	}

	fmt.Printf("Finished in %s\n", time.Since(start))
	return total
}

// May want to re-parse entirely?
func PartTwo() int64 {
	fmt.Println("Day 6 solution pt 2")
	lines := lib.ReadInput(filename)
	start := time.Now()

	var total int64 = 0
	operandIndex := 0
	nextOperandIndex := -1
	currentOperand := ""
	i := 0
	for i < len(lines[4]) {
		currentBottomChar := string(lines[4][i])

		// Get the range we want to calculate over
		if currentBottomChar == "*" || currentBottomChar == "+" {
			operandIndex = i
			currentOperand = currentBottomChar

			// find next operand
			for k := i + 1; k < len(lines[4]); k++ {
				if string(lines[4][k]) == "*" || string(lines[4][k]) == "+" {
					nextOperandIndex = k
					break
				} else if k == len(lines[4])-1 {
					nextOperandIndex = k + 2
					break
				}
			}
		}

		// Parse and calculate
		var calculated int

		if currentOperand == "*" {
			calculated = 1
		} else {
			calculated = 0
		}

		for column := nextOperandIndex - 2; column >= operandIndex; column-- {
			numString := ""
			for row := 0; row < 4; row++ {
				numString += string(lines[row][column])
			}
			num, _ := strconv.Atoi(strings.Replace(numString, " ", "", -1))

			if currentOperand == "*" {
				calculated *= num
			} else {
				calculated += num
			}
		}

		total += int64(calculated)

		i = nextOperandIndex
	}

	fmt.Printf("Finished in %s\n", time.Since(start))
	return total
}

func getProblemsMap(lines []string) map[int]*CephaloProblem {
	problems := make(map[int]*CephaloProblem)
	re := regexp.MustCompile(" ")

	for i, line := range lines {
		split := re.Split(line, -1)

		problemNumber := 0

		for _, element := range split {
			if element == "" {
				continue
			}

			if _, exists := problems[problemNumber]; !exists {
				problems[problemNumber] = &CephaloProblem{
					problemNumber,
					make([]int, 0),
					"",
				}
			}

			// Operands
			if i == 4 {
				problems[problemNumber].operand = element
			} else { // Numbers
				num, _ := strconv.Atoi(element)
				problems[problemNumber].nums = append(problems[problemNumber].nums, num)
			}

			problemNumber++
		}
	}

	return problems
}
