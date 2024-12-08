package day7

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetCalibrationResultFromFile(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totCalRes := 0
	for scanner.Scan() {
		line := scanner.Text()

		lineParts := strings.Split(line, ":")
		testRes, err := strconv.Atoi(lineParts[0])
		if err != nil {
			return 0, err
		}
		eqParts := strings.Split(strings.Trim(lineParts[1], " "), " ")
		partInts := []int{}
		for _, part := range eqParts {
			partInt, err := strconv.Atoi(part)
			if err != nil {
				return 0, err
			}
			partInts = append(partInts, partInt)
		}
		if getTestResFromNumbers(testRes, partInts) {
			totCalRes += testRes
		}
	}
	return totCalRes, nil
}

func getTestResFromNumbers(testRes int, numList []int) bool {
	// Go through all the numbers
	results := []int{}
	for index, num := range numList {
		// Have an array with each partial result growing
		if len(results) == 0 {
			results = append(results, num)
		} else {
			// after each num in list we double qty of partial results on list (one for sum other for *  for every current one)
			for resIndex, result := range results {
				results[resIndex] += num
				concat, _ := strconv.Atoi(strconv.FormatInt(int64(result), 10) + strconv.FormatInt(int64(num), 10))
				results = append(results, []int{concat, result * num}...)
				if index == len(numList)-1 {
					if result*num == testRes || result+num == testRes || concat == testRes {
						return true
					}
				}
			}
			// so for list [ 1, 2 , 3 ]
			// resList = [1]
			// resList = [3, 2]
			// resList = [9, 6, 6, 5]
			// Maybe make it a map (mapres := map[int]boolso we dont compute repeated results
		}
		// once we went through all nums we return mapRes[testRes]
	}
	return false
}
