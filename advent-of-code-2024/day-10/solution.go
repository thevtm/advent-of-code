package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/samber/lo"
)

type Position struct {
	x, y int
}

func (p Position) plus(other Position) Position {
	return Position{p.x + other.x, p.y + other.y}
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

	digit_map := map[rune]int{'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9}
	height_map := make([][]int, len(file_lines))
	starting_positions := make([]Position, 0)

	for i, line := range file_lines {
		heights := make([]int, len(line))

		for j, height_rune := range line {
			heights[j] = digit_map[height_rune]

			if height_rune == '0' {
				starting_positions = append(starting_positions, Position{j, i})
			}
		}

		height_map[i] = heights
	}

	fmt.Println(height_map)
	fmt.Println()

	fmt.Println("starting_positions", starting_positions)
	fmt.Println()

	height_map_width := len(height_map[0])
	height_map_height := len(height_map)

	fmt.Println("height_map_width", height_map_width, "height_map_height", height_map_height)

	fmt.Println()
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println()

	// PROBLEM 1

	get_height := func(position Position) int { return height_map[position.y][position.x] }
	is_in_bounds := func(pos Position) bool {
		if pos.x < 0 || pos.x >= height_map_width || pos.y < 0 || pos.y >= height_map_height {
			return false
		}

		return true
	}
	can_traverse := func(a Position, b Position) bool {
		return (get_height(b) - get_height(a)) == 1
	}

	score := 0

	for _, starting_position := range starting_positions {
		stack := make([]Position, 1)
		stack[0] = starting_position

		visited_positions := make(map[Position]bool)

		for len(stack) > 0 {
			position := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if visited_positions[position] {
				continue
			}

			visited_positions[position] = true

			if get_height(position) == 9 {
				score++
				fmt.Println("score++", starting_position, "->", position)
				continue
			}

			directions := []Position{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

			for _, direction := range directions {
				new_pos := position.plus(direction)

				if !is_in_bounds(new_pos) {
					continue
				}

				if !can_traverse(position, new_pos) {
					continue
				}

				stack = append(stack, new_pos)
			}
		}
	}

	fmt.Println("Problem 1 Result:", score) // 746
}
