// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"regexp"
// 	"unicode"
// )

// type SymbolInfo struct {
// 	Symbol rune
// 	Index  int
// }

// type NumberInfo struct {
// 	Number string
// 	StartIndex  int
// 	EndIndex	int
// }

// func getSymbolLocations(data string) []SymbolInfo {
// 	var symbols []SymbolInfo

// 	for i, char := range data {
// 		if !unicode.IsDigit(char) && !unicode.IsLetter(char) && !unicode.IsSpace(char) && char != '.' {
// 			symbols = append(symbols, SymbolInfo{Symbol: char, Index: i})
// 		}
// 	}

// 	return symbols
// }

// func getNumberLocations(data string) []NumberInfo {
// 	re := regexp.MustCompile(`\d+`)
// 	matches := re.FindAllStringIndex(data, -1)
// 	var numbers []NumberInfo

// 	for _, match := range matches {
// 		numbers = append(numbers, NumberInfo{
// 			Number:     data[match[0]:match[1]],
// 			StartIndex: match[0],
// 			EndIndex:   match[1] - 1,
// 		})
// 	}

// 	return numbers
// }

// func processFile() {
// 	dataFile, err := os.Open("test.txt")
// 	if err != nil {
// 		log.Fatalf("Error opening file: %v", err)
// 	}
// 	defer dataFile.Close()

// 	scanner := bufio.NewScanner(dataFile)
// 	lineNumber := 1
// 	var nextLineSymbols, previousLineSymbols, currentLineSymbols []SymbolInfo
// 	var currData string

// 	for scanner.Scan() {
// 		nextLineSymbols = previousLineSymbols
// 		previousLineSymbols = currentLineSymbols
// 		currData = scanner.Text()
// 		currentLineSymbols = getSymbolLocations(currData)
// 		numbers := getNumberLocations(currData)

// 		fmt.Printf("\nLine %d:\n", lineNumber)
// 		for _, numberInfo := range numbers {
// 			fmt.Printf("Number: %s, Location: %d-%d\n", numberInfo.Number, numberInfo.StartIndex, numberInfo.EndIndex)
// 		}
// 		for _, symbolInfo := range nextLineSymbols {
// 			fmt.Printf("Previous line Symbol: %c, Location: %d\n", symbolInfo.Symbol, symbolInfo.Index)
// 		}
// 		for _, symbolInfo := range previousLineSymbols {
// 			fmt.Printf("Current line Symbol: %c, Location: %d\n", symbolInfo.Symbol, symbolInfo.Index)
// 		}
// 		for _, symbolInfo := range currentLineSymbols {
// 			fmt.Printf("Next line Symbol: %c, Location: %d\n", symbolInfo.Symbol, symbolInfo.Index)
// 		}

// 		lineNumber++
// 	}
// }

// func main() {
// 	processFile()
// }

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

type NumberInfo struct {
	Number     string
	StartIndex int
	EndIndex   int
}

func processFile() {
	dataFile, err := os.Open("test.txt")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer dataFile.Close()

	scanner := bufio.NewScanner(dataFile)
	lineNumber := 1
	var nextLineSymbols, previousLineSymbols, currentLineSymbols []SymbolInfo
	var currData string

	for scanner.Scan() {
		nextLineSymbols = previousLineSymbols
		previousLineSymbols = currentLineSymbols
		currData = scanner.Text()
		currentLineSymbols = getSymbolLocations(currData)
		numbers := getNumberLocations(currData)

		fmt.Printf("\nLine %d:\n", lineNumber)
		for _, numberInfo := range numbers {
			fmt.Printf("Number: %s, Location: %d-%d\n", numberInfo.Number, numberInfo.StartIndex, numberInfo.EndIndex)
		}
		for _, symbolInfo := range nextLineSymbols {
			fmt.Printf("Next line Symbol: %c, Location: %d\n", symbolInfo.Symbol, symbolInfo.Index)
		}
		for _, symbolInfo := range previousLineSymbols {
			fmt.Printf("Previous line Symbol: %c, Location: %d\n", symbolInfo.Symbol, symbolInfo.Index)
		}
		if lineNumber != 1 {
			for _, symbolInfo := range currentLineSymbols {
				fmt.Printf("Current line Symbol: %c, Location: %d\n", symbolInfo.Symbol, symbolInfo.Index)
			}
		}

		lineNumber++
	}
}

func main() {
	processFile()
}
