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

func main() {
	// INPUTS

	input_path := "./input.txt"
	// input_path := "./input-sample.txt"

	_, source_file_path, _, _ := runtime.Caller(0)
	source_dir := filepath.Dir(source_file_path)

	file_content := string(lo.Must(os.ReadFile(filepath.Join(source_dir, input_path))))

	fmt.Println(file_content)
	fmt.Println()

	file_lines := strings.Split(file_content, "\n")
	file_lines = file_lines[:len(file_lines)-1] // Last line is blank

	first_blank_line_index := slices.Index(file_lines, "")

	page_ordering_rule_lines := file_lines[:first_blank_line_index]

	before_map := make(map[string]bool)

	for _, rule_line := range page_ordering_rule_lines {
		before_map[rule_line] = true
	}

	page_update_lines := file_lines[first_blank_line_index+1:]
	page_updates := make([][]string, len(page_update_lines))

	for i, page_update_line := range page_update_lines {
		pages := strings.Split(page_update_line, ",")
		page_updates[i] = pages
	}

	fmt.Println(page_updates)

	fmt.Println()

	// PROBLEM 1

	sum_middle := 0
	invalid_page_updates := make([][]string, 0)

PageUpdatesLoop:
	for _, page_update := range page_updates {
		fmt.Println(page_update)

		for i, page := range page_update {

			for j := 0; j < i; j++ {
				before_page := page_update[j]
				key := fmt.Sprintf("%s|%s", page, before_page)
				if _, ok := before_map[key]; ok {
					fmt.Println(" ", "invalid", key)
					invalid_page_updates = append(invalid_page_updates, page_update)
					continue PageUpdatesLoop
				}
			}
		}

		middle_index := len(page_update) / 2
		middle_page_int := lo.Must(strconv.Atoi(page_update[middle_index]))
		fmt.Println(" ", "valid", "middle_page:", middle_page_int)
		sum_middle += middle_page_int
	}

	fmt.Println("Problem 1 Result:", sum_middle) // 4689

	fmt.Println()
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println()

	// PROBLEM 2

	invalid_sum_middle := 0

	for _, invalid_page_update := range invalid_page_updates {
		fmt.Println(invalid_page_update)

		slices.SortFunc(invalid_page_update, func(a string, b string) int {
			if _, ok := before_map[fmt.Sprintf("%s|%s", a, b)]; ok {
				return -1
			}

			if _, ok := before_map[fmt.Sprintf("%s|%s", b, a)]; ok {
				return 1
			}

			return 0
		})

		fmt.Println(invalid_page_update)

		middle_index := len(invalid_page_update) / 2
		middle_page_int := lo.Must(strconv.Atoi(invalid_page_update[middle_index]))
		fmt.Println(" ", "valid", "middle_page:", middle_page_int)
		invalid_sum_middle += middle_page_int
	}
	fmt.Println()

	fmt.Println("Problem 2 Result:", invalid_sum_middle) // 6336
}
