package lib

import (
	"bufio"
	"log"
	"os"
)

func ReadInput(filepath string) []string {
	var lines []string
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal("Error reading file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
