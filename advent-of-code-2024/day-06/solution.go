package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/samber/lo"
)

type Step struct {
	x, y      int
	direction rune
}

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

	matrix := make([][]rune, len(file_lines))
	for i, line := range file_lines {
		matrix[i] = []rune(line)
	}

	matrix_width := len(matrix[0])
	matrix_height := len(matrix)

	fmt.Println("matrix_width", matrix_width, "matrix_height", matrix_height)

	start_x := 0
	start_y := 0
	start_direction := '^' // It's always ^

OuterLoop:
	for y, row := range matrix {
		for x, cell := range row {
			if cell == '^' {
				start_x = x
				start_y = y
				break OuterLoop
			}
		}
	}

	println("start_x", start_x, "start_y", start_y)

	fmt.Println()

	// PROBLEM 1

	outside_step := Step{-1, -1, 'o'}
	unique_positions := make(map[int]bool)

	get_cell := func(x int, y int) rune {
		if x < 0 || x >= matrix_width || y < 0 || y >= matrix_height {
			return 'o'
		}

		return matrix[y][x]
	}

	next_direction := func(direction rune) rune {
		if direction == '^' {
			return '>'
		} else if direction == '>' {
			return 'v'
		} else if direction == 'v' {
			return '<'
		} else if direction == '<' {
			return '^'
		}

		panic("Unreachable!")
	}

	next_step := func(step Step) Step {
		if step.direction == '^' {
			next_cell := get_cell(step.x, step.y-1)

			if next_cell == '#' {
				return Step{step.x, step.y, next_direction(step.direction)}
			} else if next_cell == 'o' {
				return outside_step
			} else {
				return Step{step.x, step.y - 1, step.direction}
			}

		} else if step.direction == '>' {
			next_cell := get_cell(step.x+1, step.y)

			if next_cell == '#' {
				return Step{step.x, step.y, next_direction(step.direction)}
			} else if next_cell == 'o' {
				return outside_step
			} else {
				return Step{step.x + 1, step.y, step.direction}
			}

		} else if step.direction == 'v' {
			next_cell := get_cell(step.x, step.y+1)

			if next_cell == '#' {
				return Step{step.x, step.y, next_direction(step.direction)}
			} else if next_cell == 'o' {
				return outside_step
			} else {
				return Step{step.x, step.y + 1, step.direction}
			}

		} else if step.direction == '<' {
			next_cell := get_cell(step.x-1, step.y)

			if next_cell == '#' {
				return Step{step.x, step.y, next_direction(step.direction)}
			} else if next_cell == 'o' {
				return outside_step
			} else {
				return Step{step.x - 1, step.y, step.direction}
			}
		} else {
			panic("Unknown direction")
		}
	}

	pos_to_key := func(x int, y int) int {
		return x*1000 + y
	}

	for step := (Step{start_x, start_y, start_direction}); step != outside_step; step = next_step(step) {
		unique_positions[pos_to_key(step.x, step.y)] = true
	}

	fmt.Println("Problem 1 Result:", len(unique_positions)) // 4663

	fmt.Println()
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println()

	// PROBLEM 2

	set_cell := func(x int, y int, val rune) {
		if x < 0 || x >= matrix_width || y < 0 || y >= matrix_height {
			panic("Outside")
		}

		matrix[y][x] = val
	}

	possible_obstacle_positions_set := make(map[int]bool)

	starting_step := Step{start_x, start_y, start_direction}

	for step := starting_step; step != outside_step; step = next_step(step) {
		if get_cell(step.x, step.y) == '^' {
			continue
		}

		// Block
		set_cell(step.x, step.y, '#')

		path := make(map[Step]bool)

		for alt_step := starting_step; alt_step != outside_step; alt_step = next_step(alt_step) {
			if path[alt_step] == true {
				possible_obstacle_positions_set[pos_to_key(step.x, step.y)] = true
				fmt.Println("found", step)
				break
			}

			path[alt_step] = true
		}

		// Unblock
		set_cell(step.x, step.y, '.')
	}
	fmt.Println()

	fmt.Println("Problem 2 Result:", len(possible_obstacle_positions_set)) // 1530
}
