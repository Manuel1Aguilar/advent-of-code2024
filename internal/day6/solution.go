package day6

import (
	"bufio"
	"os"
)

type Guard struct {
	XPos int
	YPos int
	Dir  rune
}

var directions map[rune][]int = map[rune][]int{
	'>': {0, 1},
	'<': {0, -1},
	'^': {-1, 0},
	'v': {1, 0},
}

func GetGuardPathSizeFromFile(path string) (int, int, error) {
	factoryMap, err := getRunesMatrixFromFile(path)
	if err != nil {
		return 0, 0, err
	}
	guard := getGuardFromMap(factoryMap)
	originalMap := make([][]rune, len(factoryMap))
	for i := range factoryMap {
		originalMap[i] = make([]rune, len(factoryMap[i]))
		copy(originalMap[i], factoryMap[i]) // copies the runes of this row
	}

	var originalGuard Guard = *guard
	factoryMap[guard.XPos][guard.YPos] = 'X'
	pathSize := findGuardPathSize(guard, factoryMap)
	loopLocations := checkPathForLoopObstLocations(&originalGuard, factoryMap)
	return pathSize, loopLocations, nil

}

func checkPathForLoopObstLocations(originalGuard *Guard, factoryMap [][]rune) int {
	loopLocations := 0
	for rowIndex, row := range factoryMap {
		for colIndex, cell := range row {
			if cell == 'X' {
				// Check for loop if obst here
				isLoop := checkIfMapHasLoop(originalGuard, factoryMap, []int{rowIndex, colIndex})
				if isLoop {
					loopLocations++
				}
			}
		}
	}
	return loopLocations
}

func checkIfMapHasLoop(originalGuard *Guard, factoryMap [][]rune, obstacleLocation []int) bool {
	guardFollower := &Guard{
		XPos: originalGuard.XPos,
		YPos: originalGuard.YPos,
		Dir:  originalGuard.Dir,
	}
	visitedStates := make(map[Guard]bool)
	for {
		direction := directions[guardFollower.Dir]
		newxPos := guardFollower.XPos + direction[0]
		newyPos := guardFollower.YPos + direction[1]

		if newxPos < 0 || newxPos >= len(factoryMap) || newyPos < 0 || newyPos >= len(factoryMap[0]) {
			return false
		}
		state := Guard{XPos: guardFollower.XPos, YPos: guardFollower.YPos, Dir: guardFollower.Dir}
		if visitedStates[state] {
			return true
		}
		visitedStates[state] = true
		if factoryMap[newxPos][newyPos] == '#' || (newxPos == obstacleLocation[0] && obstacleLocation[1] == newyPos) {
			// rotate logic
			guardFollower.Dir = rotateGuard(guardFollower.Dir)
		} else {
			guardFollower.XPos = newxPos
			guardFollower.YPos = newyPos
		}
	}
}
func findGuardPathSize(guard *Guard, factoryMap [][]rune) int {

	factoryMap[guard.XPos][guard.YPos] = 'X'
	pathSize := 1
	for {
		direction := directions[guard.Dir]
		newxPos := guard.XPos + direction[0]
		newyPos := guard.YPos + direction[1]

		if newxPos < 0 || newxPos >= len(factoryMap) || newyPos < 0 || newyPos >= len(factoryMap[0]) {
			break
		}
		if factoryMap[newxPos][newyPos] == '#' {
			// rotate logic
			guard.Dir = rotateGuard(guard.Dir)
		} else {
			if factoryMap[newxPos][newyPos] != 'X' {
				factoryMap[newxPos][newyPos] = 'X'
				pathSize++
			}
			guard.XPos = newxPos
			guard.YPos = newyPos
		}
	}
	return pathSize
}

func rotateGuard(oldDirection rune) rune {
	switch oldDirection {
	case '>':
		return 'v'
	case '<':
		return '^'
	case '^':
		return '>'
	case 'v', 'V':
		return '<'
	default:
		return '^'
	}
}

func getRunesMatrixFromFile(path string) ([][]rune, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	runes := [][]rune{}
	for scanner.Scan() {
		line := scanner.Text()

		lineRunes := []rune(line)
		runes = append(runes, lineRunes)
	}
	return runes, nil
}
func getGuardFromMap(factoryMap [][]rune) *Guard {

	guardPos := &Guard{}
	for i := 0; i < len(factoryMap); i++ {
		for j := 0; j < len(factoryMap[0]); j++ {

			if factoryMap[i][j] != '#' && factoryMap[i][j] != '.' {
				guardPos.Dir = factoryMap[i][j]
				guardPos.XPos = i
				guardPos.YPos = j
			}
		}
	}
	return guardPos
}
