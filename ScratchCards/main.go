package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fileContent, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	lines := strings.Split(string(fileContent), "\n")
	totalPoints := 0
	for idx, line := range lines {
		totalPoints += processLine(line, lines, idx)
	}

	fmt.Println("Total:", totalPoints)
}

func processLine(line string, lines []string, idx int) int {
	parts := strings.Split(line, "|")
	if len(parts) != 2 {
		fmt.Println("Invalid line:", line)
		return 0
	}

	left := strings.Fields(parts[0])
	right := strings.Fields(parts[1])

	matches := findMatches(left, right, lines, idx)
	a := 1
	for i := 1; i < matches+1; i++ {
		a += processLine(lines[idx+i], lines, idx+i)
	}
	return a
}

func findMatches(left, right, lines []string, idx int) int {
	matches := 0
	for _, l := range left {
		for _, r := range right {
			if l == r {
				matches++
			}
		}
	}
	if matches == 0 {
		return 0
	}
	return matches
}
