package day2

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetSafeReportQuantityFromFile(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safeCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		report := parseReport(line)

		if isSafe(report) {
			safeCount++
		} else {
			for i := range report {
				reportCopy := make([]int, len(report)-1)
				copy(reportCopy[:i], report[:i])
				copy(reportCopy[i:], report[i+1:])

				if isSafe(reportCopy) {
					safeCount++
					break
				}
			}
		}
	}

	return safeCount, nil
}

func parseReport(line string) []int {
	parts := strings.Fields(line)
	report := make([]int, len(parts))
	for i, part := range parts {
		report[i], _ = strconv.Atoi(part)
	}
	return report
}

func isSafe(report []int) bool {
	if len(report) < 2 {
		return true
	}

	firstDiff := report[1] - report[0]
	if firstDiff == 0 || abs(firstDiff) > 3 {
		return false
	}

	expectedSign := sign(firstDiff)

	for i := 1; i < len(report)-1; i++ {
		diff := report[i+1] - report[i]
		if diff == 0 || abs(diff) > 3 {
			return false
		}

		if sign(diff) != expectedSign {
			return false
		}
	}

	return true
}

func sign(x int) int {
	if x > 0 {
		return 1
	}
	if x < 0 {
		return -1
	}
	return 0
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
