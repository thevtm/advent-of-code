package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/samber/lo"
)

type Turn = [2]rune

type Position struct {
	x, y int
}

type ReindeerStep struct {
	position  Position
	direction rune
	score     int
}

func (p Position) plus(other Position) Position {
	return Position{p.x + other.x, p.y + other.y}
}

func print_rune_grid(grid [][]rune) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func main() {
	// INPUTS

	// input_path := "./input.txt"
	input_path := "./input-sample-small.txt"
	// input_path := "./input-sample-large.txt"

	_, source_file_path, _, _ := runtime.Caller(0)
	source_dir := filepath.Dir(source_file_path)

	file_content := string(lo.Must(os.ReadFile(filepath.Join(source_dir, input_path))))

	fmt.Println(file_content)
	fmt.Println()

	file_lines := strings.Split(file_content, "\n")
	file_lines = file_lines[:len(file_lines)-1] // Last line is blank

	maze := make([][]rune, len(file_lines))

	for i, line := range file_lines {
		maze[i] = []rune(line)
	}

	fmt.Println(maze)
	fmt.Println()

	maze_width := len(maze[0])
	maze_height := len(maze)

	fmt.Printf("maze => width: %d, height: %d\n\n", maze_width, maze_height)

	start_position := Position{0, 0}
	start_direction := '>'
	end_position := Position{0, 0}

	for y, row := range maze {
		for x, cell := range row {
			if cell == 'S' {
				start_position.x = x
				start_position.y = y
			} else if cell == 'E' {
				end_position.x = x
				end_position.y = y
			}
		}
	}

	fmt.Printf("starting_position %+v, end_position %+v\n\n", start_position, end_position)

	fmt.Println()
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println()

	// PROBLEM 1

	get_maze_tile := func(position Position) rune { return maze[position.y][position.x] }
	movement_for_direction := map[rune]Position{'>': {1, 0}, 'v': {0, 1}, '<': {-1, 0}, '^': {0, -1}}
	turns_for_direction := map[rune]Turn{'>': {'^', 'v'}, 'v': {'<', '>'}, '<': {'^', 'v'}, '^': {'<', '>'}}

	maze_score_by_position := make(map[Position]int)

	stack := make([]ReindeerStep, 0)
	stack = append(stack, ReindeerStep{start_position, start_direction, 0})

	maze_score_by_position[start_position] = 0

	for len(stack) > 0 {
		step := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		turn_directions := turns_for_direction[step.direction]

		for i, next_direction := range []rune{step.direction, turn_directions[0], turn_directions[1]} {
			is_turning := i != 0

			next_position := step.position.plus(movement_for_direction[next_direction])
			next_position_tile := get_maze_tile(next_position)

			if next_position_tile == '#' || next_position_tile == 'S' {
				continue
			}

			if next_position_tile != '.' && next_position_tile != 'E' {
				panic(fmt.Sprintf("Unknown tile: %c", next_position_tile))
			}

			next_score := step.score
			if is_turning {
				next_score += 1001
			} else {
				next_score += 1
			}

			if current_score, ok := maze_score_by_position[next_position]; ok && current_score <= next_score {
				continue
			}

			maze_score_by_position[next_position] = next_score

			if next_position_tile == 'E' {
				continue
			}

			stack = append(stack, ReindeerStep{next_position, next_direction, next_score})
		}
	}

	fmt.Printf("%+v\n\n", maze_score_by_position)

	final_score := maze_score_by_position[end_position]

	fmt.Println("Problem 1 Result:", final_score) // 105508
}
