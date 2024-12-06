package day5

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetCorrectUpdatesMiddleNumberSumFromFile(path string) (int, int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	isUpdatesPart := false
	sum := 0
	fixedSum := 0
	var rules [][]int
	for scanner.Scan() {

		line := scanner.Text()
		if len(line) == 0 {
			isUpdatesPart = true
			continue
		}

		if isUpdatesPart {

			parts := strings.Split(line, ",")
			var partInts = make([]int, len(parts))
			for index, element := range parts {
				num, err := strconv.Atoi(element)
				if err != nil {
					return 0, 0, err
				}

				partInts[index] = num
			}

			ruleMap := make(map[int][]int)
			for _, element := range rules {
				ruleMap[element[1]] = append(ruleMap[element[1]], element[0])
			}
			res, err := verifyUpdate(ruleMap, partInts)
			if err != nil {
				return 0, 0, err
			}
			if res {
				sum += partInts[len(partInts)/2]
			} else {
				fixedUpdate := fixUpdate(ruleMap, partInts)
				fixedSum += fixedUpdate[len(fixedUpdate)/2]
			}
		} else {
			parts := strings.Split(line, "|")
			var partInts = make([]int, len(parts))

			for index, element := range parts {
				partInts[index], err = strconv.Atoi(element)
				if err != nil {
					return 0, 0, err
				}
			}

			rules = append(rules, partInts)
		}
	}
	return sum, fixedSum, nil
}

func fixUpdate(ruleMap map[int][]int, update []int) []int {
	graph := make(map[int][]int)
	inDegree := make(map[int]int)
	for _, element := range update {
		inDegree[element] = 0
		graph[element] = []int{}
	}
	for _, element := range update {
		conflicts := getRelevantRulesForUpdateElement(ruleMap, update, element)
		for _, conflict := range conflicts {
			graph[conflict] = append(graph[conflict], element)
			inDegree[element]++
		}

	}
	queue := []int{}
	fixedUpdate := []int{}
	for _, element := range update {
		if inDegree[element] == 0 {
			queue = append(queue, element)
		}
	}

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		fixedUpdate = append(fixedUpdate, front)
		conflicts := graph[front]
		for _, element := range conflicts {
			inDegree[element]--
			if inDegree[element] == 0 {
				queue = append(queue, element)
			}
		}
	}
	return fixedUpdate
}

func getRelevantRulesForUpdateElement(ruleMap map[int][]int, update []int, element int) []int {
	var conflicts []int
	rules := ruleMap[element]
	for _, predec := range rules {
		isInUpdate := false
		for _, updateVal := range update {
			if predec == updateVal {
				isInUpdate = true
			}
		}
		if isInUpdate {
			conflicts = append(conflicts, predec)
		}
	}
	return conflicts
}
func getUpdateElementRuleConflicts(ruleMap map[int][]int, update []int, happenedMap map[int]bool, element int) []int {
	var conflicts []int
	rules := ruleMap[element]
	for _, predec := range rules {
		if happenedMap[predec] == false {
			isInUpdate := false
			for _, updateVal := range update {
				if predec == updateVal {
					isInUpdate = true
				}
			}
			if isInUpdate {
				conflicts = append(conflicts, predec)
			}
		}
	}
	return conflicts
}
func verifyUpdate(ruleMap map[int][]int, update []int) (bool, error) {
	happenedMap := make(map[int]bool)

	for _, element := range update {
		conflicts := getUpdateElementRuleConflicts(ruleMap, update, happenedMap, element)
		if len(conflicts) > 0 {
			return false, nil
		}
		happenedMap[element] = true
	}
	return true, nil
}
