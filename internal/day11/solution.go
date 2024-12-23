package day11

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetRockSizeAfterNBlinksFromFile(path string, blinkQty int) (int, error) {
	rocks, err := getRocksFromFile(path)
	if err != nil {
		return 0, err
	}

	return blinkNTimes(rocks, blinkQty), nil
}
func getRocksFromFile(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rocks := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		vals := strings.Fields(line)
		for _, val := range vals {
			intVal, err := strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
			rocks = append(rocks, intVal)
		}
	}
	return rocks, nil
}
func blinkNTimes(rocks []int, blinks int) int {
	valCount := make(map[int]int)
	for _, val := range rocks {
		valCount[val]++
	}

	for i := 0; i < blinks; i++ {
		rockCount := make(map[int]int)
		for stone, num := range valCount {
			if stone == 0 {
				rockCount[1] += num
			} else if s := strconv.Itoa(stone); len(s)%2 == 0 {
				// if len(string(rock)) % 2 == 0 then string(rock)[:len(string(rock))/2] && string(rock)[len(string(rock))/2:]
				firstHalf, secondHalf := s[:len(s)/2], s[len(s)/2:]
				fHalfInt, _ := strconv.Atoi(firstHalf)
				sHalfInt, _ := strconv.Atoi(secondHalf)
				rockCount[fHalfInt] += num
				rockCount[sHalfInt] += num
			} else {
				// else rock * 2024
				rockCount[stone*2024] += num
			}
		}
		valCount = rockCount
	}
	count := 0
	for _, val := range valCount {
		count += val
	}
	return count
}
