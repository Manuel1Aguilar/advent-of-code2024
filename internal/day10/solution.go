package day10

import (
	"bufio"
	"os"
	"strconv"
)

type Pos struct {
	X int
	Y int
}

func GetTrailMapScoreFromFile(path string) (int, error) {
	trailMap, err := getTrailMapFromFile(path)
	if err != nil {
		return 0, nil
	}

	sum := getSumForMap(trailMap)
	return sum, nil
}

func getTrailMapFromFile(path string) ([][]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var trailMap [][]int
	for scanner.Scan() {
		chars := []rune(scanner.Text())

		var ints []int
		for _, char := range chars {
			if char == '.' {
				char = '7'
			}
			num, err := strconv.Atoi(string(char))
			if err != nil {
				return nil, err
			}
			ints = append(ints, num)
		}
		trailMap = append(trailMap, ints)
	}
	return trailMap, nil
}

func getSumForMap(trailMap [][]int) int {
	scores := 0

	for rI, row := range trailMap {
		for cI, col := range row {
			if col == 0 {
				pos := Pos{X: rI, Y: cI}
				headPosMap := advanceTrail(trailMap, pos)
				headScore := 0
				for _, val := range headPosMap {
					headScore += val
				}
				scores += headScore
			}
		}
	}
	return scores
}

func advanceTrail(trailMap [][]int, currPos Pos) map[Pos]int {
	var directions [][]int = [][]int{
		{0, -1},
		{-1, 0},
		{1, 0},
		{0, 1},
	}
	positions := make(map[Pos]int)
	currVal := trailMap[currPos.X][currPos.Y]
	for _, dir := range directions {
		if currPos.X+dir[0] >= 0 && currPos.X+dir[0] < len(trailMap) && currPos.Y+dir[1] >= 0 && currPos.Y+dir[1] < len(trailMap[0]) {
			dirVal := trailMap[currPos.X+dir[0]][currPos.Y+dir[1]]

			if dirVal == currVal+1 {
				if dirVal == 9 {
					// add to positions
					positions[Pos{X: currPos.X + dir[0], Y: currPos.Y + dir[1]}]++
				} else {
					// advance trail and add results to positions
					downPos := advanceTrail(trailMap, Pos{X: currPos.X + dir[0], Y: currPos.Y + dir[1]})
					for endPos, val := range downPos {
						positions[endPos] += val
					}
				}
			}
		}
	}
	return positions
}
