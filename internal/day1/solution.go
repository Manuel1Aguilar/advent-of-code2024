package day1

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetListsSimilarity(path string) (int, error) {
	list1, list2, err := getListsFromPath(path)
	if err != nil {
		return 0, err
	}
	appearances := make(map[int]int)

	for _, element := range list2 {
		appearances[element]++
	}

	similarity := 0
	for _, element := range list1 {
		similarity += element * appearances[element]
	}

	return similarity, nil
}
func GetListsDifference(path string) (int, error) {
	list1, list2, err := getListsFromPath(path)
	if err != nil {
		return 0, err
	}
	quicksort(list1)
	quicksort(list2)
	diff := 0

	for index, element := range list1 {
		if element < list2[index] {
			diff += list2[index] - element
		} else {
			diff += element - list2[index]
		}
	}
	return diff, nil
}

func quicksort(list []int) {
	if len(list) < 2 {
		return
	}
	pivotIndex := getMedianOfThreeIndex(list)
	list[pivotIndex], list[len(list)-1] = list[len(list)-1], list[pivotIndex]
	pivotValue := list[len(list)-1]

	ltPointer := 0
	for i := 0; i < len(list)-1; i++ {
		if list[i] <= pivotValue {
			list[i], list[ltPointer] = list[ltPointer], list[i]
			ltPointer++
		}
	}
	list[ltPointer], list[len(list)-1] = list[len(list)-1], list[ltPointer]

	quicksort(list[:ltPointer])
	quicksort(list[ltPointer+1:])
}

func getMedianOfThreeIndex(list []int) int {
	first := list[0]
	middle := list[len(list)/2]
	end := list[len(list)-1]
	var pivotIndex int
	if (first <= middle && first >= end) || (first >= middle && first <= end) {
		pivotIndex = 0
	} else if (middle <= first && middle >= end) || (middle >= first && middle <= end) {
		pivotIndex = len(list) / 2
	} else {
		pivotIndex = len(list) - 1
	}
	return pivotIndex
}

func getListsFromPath(path string) ([]int, []int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var list1, list2 []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		for index, element := range parts {
			num, err := strconv.Atoi(element)
			if err != nil {
				log.Printf("Number is invalid, skipping: %v", err)
				continue
			}
			if index == 0 {
				list1 = append(list1, num)
			}
			if index == 1 {
				list2 = append(list2, num)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return list1, list2, nil
}
