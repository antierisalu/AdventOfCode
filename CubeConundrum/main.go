package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	getData()
}

func getData() {
	dataFile, err := os.Open("data.txt")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer dataFile.Close()

	scanner := bufio.NewScanner(dataFile)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, " ", 2)
		if len(parts) < 2 {
			continue
		}

		gameParts := strings.SplitN(parts[1], ":", 2)
		gameID, _ := strconv.Atoi(gameParts[0])
		data := gameParts[1]

		sets := strings.Split(data, ";")
		maxCounts := map[string]int{"red": 0, "green": 0, "blue": 0} // initialize with a small number
		for _, set := range sets {
			total := map[string]int{"red": 0, "green": 0, "blue": 0}
			cubes := strings.Split(set, ",")
			for _, cube := range cubes {
				parts := strings.Split(strings.TrimSpace(cube), " ")
				num, _ := strconv.Atoi(parts[0]) // get the number (first part)
				color := parts[1]                // get the color (second part)
				if _, ok := total[color]; ok {
					total[color] += num
				}
			}
			for color, count := range total {
				if count > maxCounts[color] {
					maxCounts[color] = count
				}
			}
		}
		product := maxCounts["red"] * maxCounts["green"] * maxCounts["blue"]
		sum += product
		fmt.Println("Game ID:", gameID, "Maximum counts:", maxCounts, "Product:", product)
	}
	fmt.Println("Sum of all products:", sum)
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
}
