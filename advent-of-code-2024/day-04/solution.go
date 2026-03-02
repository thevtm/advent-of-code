package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/samber/lo"
)

func main() {
	// INPUTS

	input_path := "./input.txt"
	// input_path := "./input-sample.txt"
	// input_path := "./input-sample-2.txt"

	_, source_file_path, _, _ := runtime.Caller(0)
	source_dir := filepath.Dir(source_file_path)

	file_content := string(lo.Must(os.ReadFile(filepath.Join(source_dir, input_path))))

	fmt.Println(file_content)
	fmt.Println()

	file_lines := strings.Split(file_content, "\n")
	file_lines = file_lines[:len(file_lines)-1] // Last line is blank

	row_count := len(file_lines)
	col_count := len(file_lines[0])

	fmt.Println("row count:", row_count, "column count:", col_count)

	matrix := make([][]rune, row_count)

	for i, line := range file_lines {
		matrix[i] = []rune(line)
	}

	get_cell := func(x int, y int) rune {
		if x < 0 || y < 0 {
			return '.'
		}

		if x >= row_count || y >= col_count {
			return '.'
		}

		return matrix[y][x]
	}

	cell_match := func(x int, y int, value rune) bool {
		return get_cell(x, y) == value
	}

	fmt.Println()

	// PROBLEM 1

	xmas_count := 0

	for i, row := range matrix {
		for j, cell := range row {
			if cell != 'X' {
				continue
			}

			// Horizontal
			if cell_match(j+1, i, 'M') && cell_match(j+2, i, 'A') && cell_match(j+3, i, 'S') {
				// fmt.Println(j, i, "horizontal")
				xmas_count++
			}

			// Reverse Horizontal
			if cell_match(j-1, i, 'M') && cell_match(j-2, i, 'A') && cell_match(j-3, i, 'S') {
				// fmt.Println(j, i, "reverse horizontal")
				xmas_count++
			}

			// Vertical
			if cell_match(j, i+1, 'M') && cell_match(j, i+2, 'A') && cell_match(j, i+3, 'S') {
				// fmt.Println(j, i, "vertical")
				xmas_count++
			}

			// Reverse Vertical
			if cell_match(j, i-1, 'M') && cell_match(j, i-2, 'A') && cell_match(j, i-3, 'S') {
				// fmt.Println(j, i, "reverse vertical")
				xmas_count++
			}

			// Diagonal
			if cell_match(j+1, i+1, 'M') && cell_match(j+2, i+2, 'A') && cell_match(j+3, i+3, 'S') {
				xmas_count++
			}

			if cell_match(j+1, i-1, 'M') && cell_match(j+2, i-2, 'A') && cell_match(j+3, i-3, 'S') {
				xmas_count++
			}

			// Reverse Diagonal
			if cell_match(j-1, i+1, 'M') && cell_match(j-2, i+2, 'A') && cell_match(j-3, i+3, 'S') {
				xmas_count++
			}

			if cell_match(j-1, i-1, 'M') && cell_match(j-2, i-2, 'A') && cell_match(j-3, i-3, 'S') {
				xmas_count++
			}
		}
	}

	fmt.Println("Problem 1 Result:", xmas_count) // 2545

	fmt.Println()
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println()

	// PROBLEM 2

	mas_count := 0

	for i, row := range matrix {
		for j, cell := range row {
			if cell != 'A' {
				continue
			}

			left_to_right := cell_match(j-1, i-1, 'M') && cell_match(j+1, i+1, 'S')
			reverse_left_to_right := cell_match(j-1, i-1, 'S') && cell_match(j+1, i+1, 'M')

			right_to_left := cell_match(j+1, i-1, 'M') && cell_match(j-1, i+1, 'S')
			reverse_right_to_left := (cell_match(j+1, i-1, 'S') && cell_match(j-1, i+1, 'M'))

			if (left_to_right || reverse_left_to_right) && (right_to_left || reverse_right_to_left) {
				mas_count++
			}
		}
	}

	fmt.Println("Problem 2 Result:", mas_count) // 1886
}
