package main

import (
	"fmt"

	"github.com/Manuel1Aguilar/advent-of-code2024/internal/day1"
	"github.com/Manuel1Aguilar/advent-of-code2024/internal/day2"
	"github.com/Manuel1Aguilar/advent-of-code2024/internal/day3"
	"github.com/Manuel1Aguilar/advent-of-code2024/internal/day4"
)

func main() {
	fmt.Println("Advent of code 2024")
	callDay1()
	callDay2()
	callDay3()
	callDay4()
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
	}

	fmt.Printf("The word count is: %d\n", xmasAppearances)
	fmt.Printf("The count of times MAS appeared crossed is: %d\n", crossedMasAppearances)
}
