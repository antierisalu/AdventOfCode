package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"unicode"
)

type SymbolInfo struct {
	Symbol rune
	Index  int
}

type NumberInfo struct {
	Number     string
	StartIndex int
	EndIndex   int
	Locations    []int
}

type LineInfo struct {
	Symbols []SymbolInfo
	Numbers []NumberInfo
	LineLen int
}

var (
	TextFile                         string = "test.txt"
	IsTop, IsLeft, IsRight, IsBottom bool
	LineCount                        int
	CurrentLineCount                 int
	SumOfNumbers                     int
)

func getSymbolLocations(data string) []SymbolInfo {
	var symbols []SymbolInfo

	for i, char := range data {
		if !unicode.IsDigit(char) && !unicode.IsLetter(char) && !unicode.IsSpace(char) && char != '.' {
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
			Locations:    indexes,
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

	// First pass: read all lines and store symbol and number information
	for scanner.Scan() {
		line := scanner.Text()
		symbols := getSymbolLocations(line)
		numbers := getNumberLocations(line)
		lineInfos = append(lineInfos, LineInfo{Symbols: symbols, Numbers: numbers, LineLen: len(line)})
		fmt.Println(line)
	}
	
	// Second pass: process each line with access to next and previous line's symbols
	for i, lineInfo := range lineInfos {

		// Access previous, current and next line's symbols
		var previousLineSymbols, nextLineSymbols []SymbolInfo
		if i > 0 {
			previousLineSymbols = lineInfos[i-1].Symbols
		}
		currentLineSymbols := lineInfo.Symbols
		if i < len(lineInfos)-1 {
			nextLineSymbols = lineInfos[i+1].Symbols
		}

		// Process each number in the current line
		for _, numberInfo := range lineInfo.Numbers {
			IsLeft = numberInfo.StartIndex == 0
			IsRight = numberInfo.EndIndex == lineInfo.LineLen-1
			IsTop = i == 0
			IsBottom = i == len(lineInfos)-1

			// Check for adjacent symbols
			for _, symbolInfo := range currentLineSymbols {
				if symbolInfo.Index == numberInfo.StartIndex-1 || symbolInfo.Index == numberInfo.EndIndex+1 {
					fmt.Printf("Number %s has an adjacent symbol %c\n", numberInfo.Number, symbolInfo.Symbol)
				}
			}
			if !IsTop {
				for _, symbolInfo := range previousLineSymbols {
					if symbolInfo.Index >= numberInfo.StartIndex && symbolInfo.Index <= numberInfo.EndIndex {
						fmt.Printf("Number %s has an adjacent symbol %c on the previous line\n", numberInfo.Number, symbolInfo.Symbol)
					}
				}
			}
			if !IsBottom {
				for _, symbolInfo := range nextLineSymbols {
					if symbolInfo.Index >= numberInfo.StartIndex && symbolInfo.Index <= numberInfo.EndIndex {
						fmt.Printf("Number %s has an adjacent symbol %c on the next line\n", numberInfo.Number, symbolInfo.Symbol)
					}
				}
			}
		}
	}
}


func main() {
	countLines()
	processFile()
}
