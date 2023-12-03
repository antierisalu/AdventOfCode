package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func parseNumbersFromString(s string) (firstNum, lastNum int, err error) {
	var firstNumberStr, lastNumberStr string
	var foundFirstNumber, foundLastNumber bool
	var firstNumberIdx, lastNumberIdx int

	// Substrings representing numbers
	substrings := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	// Check for number substrings in the string and store the closest to the start
	for _, substr := range substrings {
		if idx := strings.Index(s, substr); idx != -1 && (!foundFirstNumber || idx < firstNumberIdx) {
			firstNumberStr = substr
			firstNumberIdx = idx
			foundFirstNumber = true
		}
	}
	// Check for digits in the string and store the closest to the start
	for _, ch := range s {
		if unicode.IsDigit(ch) {
			digitStr := string(ch)
			idx := strings.IndexRune(s, ch)
			if !foundFirstNumber || idx < firstNumberIdx {
				firstNumberStr = digitStr
				firstNumberIdx = idx
				foundFirstNumber = true
			}
		}
	}

	// Check for number substrings in the string and store the closest to the end
	for i := len(substrings) - 1; i >= 0; i-- {
		if idx := strings.LastIndex(s, substrings[i]); idx != -1 && (!foundLastNumber || idx > lastNumberIdx) {
			lastNumberStr = substrings[i]
			lastNumberIdx = idx
			foundLastNumber = true
		}
	}

	// Check for digits in the string and store the closest to the end
	for i := len(s) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(s[i])) {
			digitStr := string(s[i])
			idx := strings.LastIndexByte(s, s[i])
			if !foundLastNumber || idx > lastNumberIdx {
				lastNumberStr = digitStr
				lastNumberIdx = idx
				foundLastNumber = true
			}
		}
	}

	// If no numbers found, return an error
	if !foundFirstNumber && !foundLastNumber {
		return 0, 0, fmt.Errorf("no numbers found in the string")
	}

	// Convert the found substrings or digits to integers
	firstNum, err = replaceSubstrings(firstNumberStr)
	if err != nil {
		return 0, 0, fmt.Errorf("error converting first number: %v", err)
	}

	lastNum, err = replaceSubstrings(lastNumberStr)
	if err != nil {
		return 0, 0, fmt.Errorf("error converting last number: %v", err)
	}

	return firstNum, lastNum, err
}

// Function to replace substrings with corresponding numbers
func replaceSubstrings(s string) (int, error) {
	// Mapping of number words to their integer representations
	numbersMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	if val, ok := numbersMap[s]; ok {
		return val, nil
	}

	// If the input string is a single digit, convert it to an integer
	if len(s) == 1 && unicode.IsDigit(rune(s[0])) {
		num, err := strconv.Atoi(s)
		if err != nil {
			return 0, fmt.Errorf("error converting digit to integer: %v", err)
		}
		return num, nil
	}

	return 0, fmt.Errorf("invalid input")
}

func main() {
	staringTime := time.Now()

	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatalf("Error opening the file: %v", file)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalSum := 0
	for scanner.Scan() {
		line := scanner.Text()

		// Example usage
		firstResult, lastResult, err := parseNumbersFromString(line)
		a, b := firstResult, lastResult
		a = a*10 + b
		fmt.Println(a)
		totalSum += a

		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("First Parsed number:", firstResult)
			fmt.Println("Last Parsed number:", lastResult)
		}
		log.Println(totalSum)
	}
	log.Println(time.Since(staringTime))
}
