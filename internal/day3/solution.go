package day3

import (
	"log"
	"os"
	"strconv"
	"unicode"
)

func GetProgramsOutputFromFile(path string) (int, error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}

	res, err := getProgramsOutput(string(contents))
	if err != nil {
		return 0, err
	}

	return res, nil
}

func getProgramsOutput(input string) (int, error) {

	runeArray := []rune(input)
	startString := "mul("
	wordIndex := 0
	total := 0
	for i := 0; i < len(input); i++ {
		if wordIndex == len(startString) {
			wordIndex = 0
			isSecondNumber := false
			firstNumber := ""
			secondNumber := ""
			isValid := true
			for (unicode.IsDigit(runeArray[i+wordIndex]) || runeArray[i+wordIndex] == ',') && isValid {
				character := runeArray[i+wordIndex]
				if character == ',' && !isSecondNumber && len(firstNumber) > 0 {
					isSecondNumber = true
				} else if unicode.IsDigit(character) {
					if !isSecondNumber {
						firstNumber += string(character)
					} else {
						secondNumber += string(character)
					}
				} else {
					isValid = false
				}
				wordIndex++
			}
			if runeArray[i+wordIndex] == ')' && isValid && (len(firstNumber) > 0 && len(secondNumber) > 0) {
				firstNumberInt, err := strconv.Atoi(firstNumber)
				if err != nil {
					log.Fatalf("Error converting string to int: %v; Input pos: %d, wordIndex: %d, string: %s", err, i, wordIndex,
						firstNumber)
				}
				secondNumberInt, err := strconv.Atoi(secondNumber)
				if err != nil {
					log.Fatalf("Error converting string to int: %v; Input pos: %d, wordIndex: %d, string: %s", err, i, wordIndex,
						secondNumber)
				}
				total += firstNumberInt * secondNumberInt
			}
			wordIndex = 0
		}

		if input[i] == startString[wordIndex] {
			wordIndex++
		} else if input[i] == startString[0] {
			wordIndex = 1
		} else {
			wordIndex = 0
		}
	}

	return total, nil
}