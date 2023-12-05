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
}

var (
	IsTop            bool = false
	IsRight          bool = false
	IsLeft           bool = false
	IsBottom         bool = false
	LineCount        int
	CurrentLineCount int
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
		numbers = append(numbers, NumberInfo{
			Number:     data[match[0]:match[1]],
			StartIndex: match[0],
			EndIndex:   match[1] - 1,
		})
	}
	return numbers
}

func countLines() {
	// Assuming "dataFile.txt" is the name of your file
	file, err := os.Open("test.txt")
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
	dataFile, err := os.Open("test.txt")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer dataFile.Close()

	scanner := bufio.NewScanner(dataFile)

	var nextLineSymbols, previousLineSymbols, currentLineSymbols []SymbolInfo
	var currentData string

	for scanner.Scan() {
		nextLineSymbols = previousLineSymbols
		previousLineSymbols = currentLineSymbols
		currentData = scanner.Text()
		currentLineSymbols = getSymbolLocations(currentData)
		numbers := getNumberLocations(currentData)
		lineLength := len(currentData)-1

		for _, numberInfo := range numbers {
			IsLeft = numberInfo.StartIndex == 0
			IsRight = numberInfo.EndIndex == lineLength
			IsTop = CurrentLineCount == 1
			IsBottom = CurrentLineCount == LineCount

			fmt.Printf("\nLine %d - Number: %s, Location: %d-%d\n", CurrentLineCount, numberInfo.Number, numberInfo.StartIndex, numberInfo.EndIndex)
			fmt.Printf("IsTop: %t, IsRight: %t, IsLeft: %t, IsBottom: %t\n", IsTop, IsRight, IsLeft, IsBottom)
		}
		for _, symbolInfo := range nextLineSymbols {
			fmt.Printf("Next line Symbol: %c, Location: %d\n", symbolInfo.Symbol, symbolInfo.Index)
		}
		for _, symbolInfo := range previousLineSymbols {
			fmt.Printf("Previous line Symbol: %c, Location: %d\n", symbolInfo.Symbol, symbolInfo.Index)
		}
		if CurrentLineCount != 0 {
			for _, symbolInfo := range currentLineSymbols {
				fmt.Printf("Current Symbol: %c, Location: %d\n", symbolInfo.Symbol, symbolInfo.Index)
			}
		}
		CurrentLineCount++
	}
}

func main() {
	countLines()
	processFile()


}
