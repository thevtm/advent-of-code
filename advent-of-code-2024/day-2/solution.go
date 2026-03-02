package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func validate(report []int, removed_index int) bool {
	direction := "unknown"

	for i := 0; i < len(report)-1; i++ {
		if i == removed_index {
			continue
		}

		j := i + 1

		if j == removed_index {
			j++
		}

		if j >= len(report) {
			break
		}

		diff := report[i] - report[j]
		diff_abs := absInt(diff)

		if diff_abs < 1 || diff_abs > 3 {
			return false
		}

		if direction == "unknown" {
			if diff > 0 {
				direction = "increasing"
			} else {
				direction = "decreasing"
			}
		} else {
			if direction == "increasing" && diff < 0 {
				return false
			} else if direction == "decreasing" && diff > 0 {
				return false
			}
		}
	}

	return true
}

func main() {
	// INPUTS

	input_path := "./input.txt"
	// input_path := "./input-sample.txt"
	// input_path := "./input-sample-2.txt"

	_, source_file_path, _, _ := runtime.Caller(0)
	source_dir := filepath.Dir(source_file_path)

	file_content := string(lo.Must(os.ReadFile(filepath.Join(source_dir, input_path))))

	fmt.Println(file_content)

	file_lines := strings.Split(file_content, "\n")
	file_lines = file_lines[:len(file_lines)-1] // Last element is blank

	fmt.Println(file_lines)
	fmt.Println()

	reports := make([][]int, len(file_lines))

	for i, line := range file_lines {
		line_split := strings.Split(line, " ")

		fmt.Println(line_split)

		reports[i] = lo.Map(line_split, func(s string, _ int) int { return lo.Must(strconv.Atoi(s)) })

	}
	fmt.Println()

	fmt.Println(reports)
	fmt.Println()

	// PROBLEM 1

	safe_count := 0

	for _, report := range reports {
		fmt.Println(report)

		if validate(report, -1) {
			fmt.Println("Safe")
			safe_count++
		} else {
			fmt.Println("Unsafe")
		}
	}
	fmt.Println()

	fmt.Println("Problem 1 Result:", safe_count) // 624
	fmt.Println()

	// PROBLEM 2

	damped_safe_count := 0

DampedReportLoop:
	for _, report := range reports {
		fmt.Println(report)

		if validate(report, -1) {
			fmt.Println("Safe")
			damped_safe_count++
			continue
		}

		for i := range len(report) {
			if validate(report, i) {
				fmt.Println("Safe without", report[i])
				damped_safe_count++
				continue DampedReportLoop
			}
		}

		fmt.Println("Unsafe")
	}
	fmt.Println()

	fmt.Println("Problem 2 Result:", damped_safe_count) // 658
	fmt.Println()
}
