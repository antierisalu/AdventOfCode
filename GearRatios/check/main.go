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
	Number string
	StartIndex  int
	EndIndex	int
}

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
		numbers = append(numbers, NumberInfo{
			Number:     data[match[0]:match[1]],
			StartIndex: match[0],
			EndIndex:   match[1] - 1,
		})
	}

	return numbers
}

func processFile() {
	dataFile, err := os.Open("test.txt")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer dataFile.Close()

	scanner := bufio.NewScanner(dataFile)
	lineNumber := 1
	var prevSymbols, currSymbols, nextSymbols []SymbolInfo
	var currData string

	for scanner.Scan() {
		prevSymbols = currSymbols
		currSymbols = nextSymbols
		currData = scanner.Text()
		nextSymbols = getSymbolLocations(currData)
		numbers := getNumberLocations(currData)



		fmt.Printf("\nLine %d:\n", lineNumber)
		for _, numberInfo := range numbers {
			fmt.Printf("Number: %s, Location: %d-%d\n", numberInfo.Number, numberInfo.StartIndex, numberInfo.EndIndex)
		}
		for _, symbolInfo := range prevSymbols {
			fmt.Printf("Previous line Symbol: %c, Location: %d\n", symbolInfo.Symbol, symbolInfo.Index)
		}
		for _, symbolInfo := range currSymbols {
			fmt.Printf("Current line Symbol: %c, Location: %d\n", symbolInfo.Symbol, symbolInfo.Index)
		}
		for _, symbolInfo := range nextSymbols {
			fmt.Printf("Next line Symbol: %c, Location: %d\n", symbolInfo.Symbol, symbolInfo.Index)
		}

		lineNumber++
	}
}

func main() {
	processFile()
}
