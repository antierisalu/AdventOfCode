package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalPoints := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			fmt.Println("Invalid line:", line)
			continue
		}

		left := strings.Fields(parts[0])
		right := strings.Fields(parts[1])

		points := calculatePoints(left, right)
		totalPoints += points
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println("Total points:", totalPoints)
}

func calculatePoints(left, right []string) int {
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
	return 1 << (matches - 1) // This line calculates the points as per your rules
}
