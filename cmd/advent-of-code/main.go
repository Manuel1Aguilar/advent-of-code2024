package main

import (
	"fmt"

	"github.com/Manuel1Aguilar/advent-of-code2024/internal/day1"
)

func main() {
	fmt.Println("Advent of code 2024")
	fmt.Println("First problem: Difference between integer lists")

	diff, err := day1.GetListsDifference("assets/problem1lists.txt")
	if err != nil {
		fmt.Printf("Error getting the difference between the lists: %v", err)
		return
	}

	fmt.Printf("The difference is: %d\n", diff)
}
