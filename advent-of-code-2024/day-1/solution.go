package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"slices"
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

func main() {
	// INPUTS

	input_path := "./input.txt"
	// input_path := "./input-sample.txt"

	_, source_file_path, _, _ := runtime.Caller(0)
	source_dir := filepath.Dir(source_file_path)

	file_content := string(lo.Must(os.ReadFile(filepath.Join(source_dir, input_path))))

	fmt.Println(file_content)

	file_lines := strings.Split(file_content, "\n")
	file_lines = file_lines[:len(file_lines)-1] // Last element is blank

	fmt.Println(file_lines)

	left, right := make([]int, len(file_lines)), make([]int, len(file_lines))

	for i, line := range file_lines {
		line_split := strings.Split(line, "   ")

		fmt.Println(line_split)

		left[i] = lo.Must(strconv.Atoi(line_split[0]))
		right[i] = lo.Must(strconv.Atoi(line_split[1]))
	}

	// PROBLEM 1

	fmt.Println(left)
	fmt.Println(right)

	slices.Sort(left)
	slices.Sort(right)

	fmt.Println(left)
	fmt.Println(right)

	total_distance := 0

	for i := range len(left) {
		total_distance += absInt(left[i] - right[i])
	}

	fmt.Println("Problem 1 Result:", total_distance) // 1341714

	// PROBLEM 2

	left_freq, right_freq := make(map[int]int), make(map[int]int)

	for i := range len(left) {
		left_freq[left[i]] += 1
		right_freq[right[i]] += 1
	}

	fmt.Println(left_freq)
	fmt.Println(right_freq)

	similarity_score := 0

	for left_key, left_val := range left_freq {
		similarity_score += right_freq[left_key] * left_val * left_key
	}

	fmt.Println("Problem 2 Result:", similarity_score) // 27384707
}
