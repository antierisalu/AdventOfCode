package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type SymbolInfo struct {
	Symbol rune
	Index  int
}

type NumberInfo struct {
	Number     string
	StartIndex int
	EndIndex   int
	Locations  []int
}

type LineInfo struct {
	Symbols []SymbolInfo
	Numbers []NumberInfo
	LineLen int
}

var (
	TextFile                         string = "data.txt"
	IsTop, IsLeft, IsRight, IsBottom bool
	LineCount                        int
	CurrentLineCount                 int
)

func getSymbolLocations(data string) []SymbolInfo {
	var symbols []SymbolInfo

	for i, char := range data {
		if char == '*' {
			symbols = append(symbols, SymbolInfo{Symbol: char, Index: i})
		}
	}
	return symbols
}

func getNumberLocations(data string) []NumberInfo {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllStringIndex(data, -1)
	var numbers []NumberInfo

	for _, match := range matches {
		indexes := make([]int, match[1]-match[0])
		for i := range indexes {
			indexes[i] = match[0] + i
		}
		numbers = append(numbers, NumberInfo{
			Number:     data[match[0]:match[1]],
			StartIndex: match[0],
			EndIndex:   match[1] - 1,
			Locations:  indexes,
		})
	}
	return numbers
}

func countLines() {
	file, err := os.Open(TextFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	CurrentLineCount = 1

	for scanner.Scan() {
		LineCount++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Total lines in the file:", LineCount)
}

func processFile() {
	dataFile, err := os.Open(TextFile)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer dataFile.Close()

	scanner := bufio.NewScanner(dataFile)

	var lineInfos []LineInfo
	var totalSum int

	// First pass: read all lines and store symbol and number information
	for scanner.Scan() {
		line := scanner.Text()
		symbols := getSymbolLocations(line)
		numbers := getNumberLocations(line)
		lineInfos = append(lineInfos, LineInfo{Symbols: symbols, Numbers: numbers, LineLen: len(line)})
	}

	// Second pass: process each line with access to next and previous line's symbols
	for i, lineInfo := range lineInfos {

		// Access previous, current and next line's numbers
		var previousLineNumbers, nextLineNumbers []NumberInfo
		if i > 0 {
			previousLineNumbers = lineInfos[i-1].Numbers
		}
		currentLineNumbers := lineInfo.Numbers
		if i < len(lineInfos)-1 {
			nextLineNumbers = lineInfos[i+1].Numbers
		}

		// Process each symbol in the current line
		for _, symbolInfo := range lineInfo.Symbols {
			IsLeft = symbolInfo.Index == 0
			IsRight = symbolInfo.Index == lineInfo.LineLen-1
			IsTop = i == 0
			IsBottom = i == len(lineInfos)-1

			// Check for adjacent numbers
			adjacentNumbers := make([]string, 0)
			for _, numberInfo := range currentLineNumbers {
				if numberInfo.StartIndex == symbolInfo.Index+1 || numberInfo.EndIndex == symbolInfo.Index-1 {
					adjacentNumbers = append(adjacentNumbers, numberInfo.Number)
				}
			}
			if !IsTop {
				for _, numberInfo := range previousLineNumbers {
					if numberInfo.StartIndex <= symbolInfo.Index+1 && numberInfo.EndIndex >= symbolInfo.Index-1 {
						adjacentNumbers = append(adjacentNumbers, numberInfo.Number)
					}
				}
			}
			if !IsBottom {
				for _, numberInfo := range nextLineNumbers {
					if numberInfo.StartIndex <= symbolInfo.Index+1 && numberInfo.EndIndex >= symbolInfo.Index-1 {
						adjacentNumbers = append(adjacentNumbers, numberInfo.Number)
					}
				}
			}

			// If symbol has two or more adjacent numbers, print a message
			if len(adjacentNumbers) >= 2 {
				fmt.Printf("Symbol %c has two or more adjacent numbers: %v\n", symbolInfo.Symbol, adjacentNumbers)
				product := 1
				for _, numStr := range adjacentNumbers {
					num, err := strconv.Atoi(numStr)
					if err != nil {
						log.Fatalf("Error converting string to integer: %v", err)
					}
					product *= num
				}
				totalSum += product
			}
			fmt.Printf("The total sum of all numbers is %d\n", totalSum)
		}
		
	}
}

func main() {
	
	processFile()
	countLines()
}
