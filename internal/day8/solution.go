package day8

import (
	"bufio"
	"os"
)

type Pos struct {
	X int
	Y int
}

func GetValidAntinodesFromFile(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	var cityMap [][]rune
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		cityMap = append(cityMap, []rune(line))
	}
	groups, antennas := getAntennaGroups(cityMap)
	res := getNodesQtyFromAntennaGroups(groups, antennas, len(cityMap), len(cityMap[0]))
	return res, nil
}

// Make a function that goes through all pos in an antenna and gets the diff between each coords
// E.g. {1, 0} diff {3, 2} = {-1, -2 } and { 5, 4 }
// Discard all outside the bounds of the map (<0 >len row and col)
func getNodesQtyFromAntennaGroups(groups map[rune][]Pos, antennas []rune, xLimit int, yLimit int) int {
	antinodeLocations := make(map[Pos]bool)
	antinodePosCount := 0
	for _, antenna := range antennas {
		group := groups[antenna]
		for index, pos := range group {
			if !antinodeLocations[pos] {
				antinodeLocations[pos] = true
				antinodePosCount++
			}
			for i := index + 1; i < len(group); i++ {
				deltaX := pos.X - group[i].X
				deltaY := pos.Y - group[i].Y
				antinodePos := Pos{X: pos.X + deltaX, Y: pos.Y + deltaY}
				for antinodePos.X >= 0 && antinodePos.X < xLimit && antinodePos.Y >= 0 && antinodePos.Y < yLimit {

					if !antinodeLocations[antinodePos] {
						antinodeLocations[antinodePos] = true
						antinodePosCount++
					}
					antinodePos.X += deltaX
					antinodePos.Y += deltaY
				}
				antinodePosBack := Pos{X: group[i].X - deltaX, Y: group[i].Y - deltaY}
				for antinodePosBack.X >= 0 && antinodePosBack.X < xLimit && antinodePosBack.Y >= 0 && antinodePosBack.Y < yLimit {
					if !antinodeLocations[antinodePosBack] {
						antinodeLocations[antinodePosBack] = true
						antinodePosCount++
					}
					antinodePosBack.X -= deltaX
					antinodePosBack.Y -= deltaY
				}
			}
		}
	}
	return antinodePosCount
}

func getAntennaGroups(cityMap [][]rune) (map[rune][]Pos, []rune) {
	groups := make(map[rune][]Pos)
	var antennas []rune
	for rowI, row := range cityMap {
		for colI, cell := range row {
			if cell != '.' {
				if len(groups[cell]) == 0 {
					antennas = append(antennas, cell)
				}
				groups[cell] = append(groups[cell], Pos{X: rowI, Y: colI})
			}
		}
	}

	return groups, antennas
}
