package day4

import (
	"bufio"
	"os"
)

func WordSearchSolveFromFile(path string) (int, int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, 0, nil
	}

	scanner := bufio.NewScanner(file)
	var runeMatrix [][]rune

	for scanner.Scan() {
		row := []rune(scanner.Text())
		runeMatrix = append(runeMatrix, row)
	}
	crossedMASappearances, err := countCrossedMasAppearencesInMatrix(runeMatrix)
	if err != nil {
		return 0, 0, nil
	}
	wordApp, err := countWordAppearencesInMatrix(runeMatrix, "XMAS")
	if err != nil {
		return 0, 0, nil
	}
	return crossedMASappearances, wordApp, nil
}
func countWordAppearencesInMatrix(input [][]rune, word string) (int, error) {
	mWidth := len(input[0])
	mHeight := len(input)
	wordCount := 0
	for i := 0; i < mHeight; i++ {
		for j := 0; j < mWidth; j++ {

			wordCount += searchAllDirections(input, j, i, word)
		}
	}
	return wordCount, nil
}
func countCrossedMasAppearencesInMatrix(input [][]rune) (int, error) {
	mWidth := len(input[0])
	mHeight := len(input)
	wordCount := 0
	for i := 0; i < mHeight; i++ {
		for j := 0; j < mWidth; j++ {
			if searchCrossedMAS(input, i, j) {
				wordCount++
			}
		}
	}
	return wordCount, nil
}
func searchCrossedMAS(input [][]rune, startRow int, startCol int) bool {
	if startCol < 1 || startRow < 1 || startCol >= len(input[0])-1 || startRow >= len(input)-1 ||
		input[startRow][startCol] != 'A' {
		return false
	}
	words := []string{
		string(input[startRow-1][startCol-1]) + string(input[startRow+1][startCol+1]),
		string(input[startRow-1][startCol+1]) + string(input[startRow+1][startCol-1]),
	}

	return (words[0] == "SM" || words[0] == "MS") && (words[1] == "SM" || words[1] == "MS")

}

func searchAllDirections(input [][]rune, startRow int, startCol int, word string) int {
	directions := [][]int{
		{0, 1},   // right
		{0, -1},  // left
		{1, 0},   // down
		{-1, 0},  // up
		{1, 1},   // down-right
		{1, -1},  // down-left
		{-1, 1},  // up-right
		{-1, -1}, // up-left
	}
	finds := 0

	mWidth := len(input[0])
	mHeight := len(input)
	for _, dir := range directions {
		currI := startCol
		currJ := startRow
		findIndex := 0
		for currI < mWidth && currI >= 0 && currJ < mHeight && currJ >= 0 &&
			findIndex < len(word) &&
			rune(word[findIndex]) == input[currJ][currI] {

			currI += dir[0]
			currJ += dir[1]
			findIndex++
		}
		if findIndex == len(word) {
			finds++
		}
	}
	return finds
}
