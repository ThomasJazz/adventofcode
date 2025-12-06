package day6

import (
	"fmt"
	"regexp"
	"strconv"
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

func PartTwo() int {
	fmt.Println("Day 6 solution pt 2")
	lines := lib.ReadInput(filename)
	start := time.Now()
	problems := getProblemsMap(lines)

	total := 0

	for _, problem := range problems {
		fmt.Println(problem.nums)

		print(problem.operand)
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
