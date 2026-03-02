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

	next_pos := func(step Step) (int, int) {
		if step.direction == '^' {
			return step.x + 0, step.y - 1
		} else if step.direction == '>' {
			return step.x + 1, step.y + 0
		} else if step.direction == 'v' {
			return step.x + 0, step.y + 1
		} else if step.direction == '<' {
			return step.x - 1, step.y + 0
		}

		panic("Unreachable!")
	}

	possible_obstacle_positions := 0

	path := make(map[Step]bool)

	for step := (Step{start_x, start_y, start_direction}); step != outside_step; step = next_step(step) {
		path[step] = true

		alternative_path := make(map[Step]bool)
		alternative_step := Step{step.x, step.y, next_direction(step.direction)}

		obstacle_pos_x, obstacle_pos_y := next_pos(step)

		if obstacle_pos_x < 0 || obstacle_pos_x >= matrix_width || obstacle_pos_y < 0 || obstacle_pos_y >= matrix_height {
			continue
		}

		if matrix[obstacle_pos_y][obstacle_pos_x] == '#' {
			continue
		}

		obstacle_position_value := matrix[obstacle_pos_y][obstacle_pos_x]
		matrix[obstacle_pos_y][obstacle_pos_x] = '#'

		for {
			if path[alternative_step] || alternative_path[alternative_step] {
				possible_obstacle_positions++
				break
			}

			if alternative_step == outside_step {
				break
			}

			alternative_path[alternative_step] = true
			alternative_step = next_step(alternative_step)
		}

		matrix[obstacle_pos_y][obstacle_pos_x] = obstacle_position_value
	}
	fmt.Println()

	fmt.Println("Problem 2 Result:", possible_obstacle_positions) // 1729
}
