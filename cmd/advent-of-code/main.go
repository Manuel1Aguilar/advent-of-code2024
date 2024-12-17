package main

import (
	"fmt"

	"github.com/Manuel1Aguilar/advent-of-code2024/internal/day1"
	"github.com/Manuel1Aguilar/advent-of-code2024/internal/day10"
	"github.com/Manuel1Aguilar/advent-of-code2024/internal/day2"
	"github.com/Manuel1Aguilar/advent-of-code2024/internal/day3"
	"github.com/Manuel1Aguilar/advent-of-code2024/internal/day4"
	"github.com/Manuel1Aguilar/advent-of-code2024/internal/day5"
	"github.com/Manuel1Aguilar/advent-of-code2024/internal/day6"
	"github.com/Manuel1Aguilar/advent-of-code2024/internal/day7"
	"github.com/Manuel1Aguilar/advent-of-code2024/internal/day8"
	"github.com/Manuel1Aguilar/advent-of-code2024/internal/day9"
)

func main() {
	fmt.Println("Advent of code 2024")
	callDay1()
	callDay2()
	callDay3()
	callDay4()
	callDay5()
	//callDay6()
	callDay7()
	callDay8()
	callDay9()
	callDay10()
}

func callDay1() {
	fmt.Println("Day 1:")
	fmt.Println("Get difference between two lists of numbers")

	path := "assets/day1input.txt"

	diff, err := day1.GetListsDifference(path)
	if err != nil {
		fmt.Printf("Error getting the difference between the lists: %v \n", err)
		return
	}

	fmt.Printf("The difference is: %d \n", diff)

	fmt.Println("Second problem: Similarity between integer lists")

	similarity, err := day1.GetListsSimilarity(path)
	if err != nil {
		fmt.Printf("Error getting the similarity between the lists: %v \n", err)
		return
	}

	fmt.Printf("The lists similarity is: %d \n", similarity)
}

func callDay2() {
	fmt.Println("Day 2:")
	fmt.Println("Get quantity of safe reports from input")

	path := "assets/day2input.txt"
	safeReports, err := day2.GetSafeReportQuantityFromFile(path)
	if err != nil {
		fmt.Printf("Error getting the quantity of safe reports from file: %v \n", err)
		return
	}

	fmt.Printf("The number of safe reports with the dampener is to %d \n", safeReports)
}

func callDay3() {
	fmt.Println("Day 3:")
	fmt.Println("Get programs output from text file")

	path := "assets/day3input.txt"
	res, err := day3.GetProgramsOutputFromFile(path)
	if err != nil {
		fmt.Printf("Error getting the programs output from file: %v \n", err)
		return
	}

	fmt.Printf("The programs output is: %d\n", res)
}

func callDay4() {
	fmt.Println("Day 4:")
	fmt.Println("Get count of word appearances on input")

	path := "assets/day4input.txt"
	crossedMasAppearances, xmasAppearances, err := day4.WordSearchSolveFromFile(path)
	if err != nil {
		fmt.Printf("Error getting the word count: %v \n", err)
		return
	}

	fmt.Printf("The word count is: %d\n", xmasAppearances)
	fmt.Printf("The count of times MAS appeared crossed is: %d\n", crossedMasAppearances)
}

func callDay5() {
	fmt.Println("Day 5:")
	fmt.Println("Get sum of the middle of the valid ordered updates")
	path := "assets/day5input.txt"

	res, fixedRes, err := day5.GetCorrectUpdatesMiddleNumberSumFromFile(path)
	if err != nil {
		fmt.Printf("Error getting the sum: %v\n", err)
		return
	}

	fmt.Printf("The sum is: %d, the fixed sum is: %d\n", res, fixedRes)
}

func callDay6() {
	fmt.Println("Day 6:")
	fmt.Println("Get guards path size given factory layout")
	path := "assets/day6input.txt"
	res, loopLocations, err := day6.GetGuardPathSizeFromFile(path)
	if err != nil {
		fmt.Printf("Error getting the path size: %v\n", err)
		return
	}
	fmt.Printf("The path size is: %d\n", res)
	fmt.Printf("Theres %d available loop locations \n", loopLocations)
}

func callDay7() {
	fmt.Println("Day 7:")
	fmt.Println("Get calibration results")
	path := "assets/day7input.txt"
	res, err := day7.GetCalibrationResultFromFile(path)
	if err != nil {
		fmt.Printf("Error getting the calibration result: %v\n", err)
		return
	}
	fmt.Printf("The calibration result is: %d\n", res)
}
func callDay8() {
	fmt.Println("Day 8:")
	fmt.Println("Get antinodes qty")
	path := "assets/day8input.txt"
	res, err := day8.GetValidAntinodesFromFile(path)
	if err != nil {
		fmt.Printf("Error getting the antinode qty: %v\n", err)
		return
	}
	fmt.Printf("The antinode qty is: %d\n", res)
}
func callDay9() {
	fmt.Println("Day 9:")
	fmt.Println("Get disk checksum")
	path := "assets/day9input.txt"
	res, err := day9.GetCorrectDiskChecksumFromFile(path)
	if err != nil {
		fmt.Printf("Error getting the disks checksum: %v\n", err)
		return
	}
	fmt.Printf("The disks checksum is: %d\n", res)
}

func callDay10() {
	fmt.Println("Day 10:")
	fmt.Println("Get trail heads scores")
	path := "assets/day10input.txt"
	res, err := day10.GetTrailMapScoreFromFile(path)
	if err != nil {
		fmt.Printf("Error getting the trails scores: %v\n", err)
		return
	}
	fmt.Printf("The score is: %d\n", res)
}
