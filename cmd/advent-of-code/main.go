package main

import (
	"fmt"

	"github.com/Manuel1Aguilar/advent-of-code2024/internal/day1"
)

func main() {
	fmt.Println("Advent of code 2024")
	fmt.Println("First problem: Difference between integer lists")
	path := "assets/problem1lists.txt"
	diff, err := day1.GetListsDifference(path)
	if err != nil {
		fmt.Printf("Error getting the difference between the lists: %v", err)
		return
	}

	fmt.Printf("The difference is: %d\n", diff)
	similarity, err := day1.GetListsSimilarity(path)

	fmt.Printf("The similarity is: %d\n", similarity)
}
