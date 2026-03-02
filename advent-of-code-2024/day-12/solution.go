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

	garden := make([][]rune, len(file_lines))

	for i, line := range file_lines {
		garden[i] = []rune(line)
	}

	fmt.Println(garden)
	fmt.Println()

	garden_width := len(garden[0])
	garden_height := len(garden)

	fmt.Println("width", garden_width, "height", garden_height)

	fmt.Println()
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println()

	// PROBLEM 1

	get_plant := func(position Position) rune { return garden[position.y][position.x] }
	is_in_bounds := func(pos Position) bool {
		if pos.x < 0 || pos.x >= garden_width || pos.y < 0 || pos.y >= garden_height {
			return false
		}

		return true
	}

	visited_positions := make(map[Position]bool)
	price := 0

	for x := range garden_width {
		for y := range garden_height {
			position := Position{x, y}

			if visited_positions[position] {
				continue
			}

			stack := make([]Position, 1)
			stack[0] = position

			plant := get_plant(position)
			area := 0
			perimeter := 0

			for len(stack) > 0 {
				stack_position := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if visited_positions[stack_position] {
					continue
				}

				visited_positions[stack_position] = true

				area += 1

				for _, direction := range []Position{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} {
					new_position := stack_position.plus(direction)

					if !is_in_bounds(new_position) || get_plant(new_position) != plant {
						perimeter += 1
						continue
					}

					stack = append(stack, new_position)
				}
			}

			fmt.Println(string(plant), area, perimeter)

			price += area * perimeter
		}
	}
	fmt.Println()

	fmt.Println("Problem 1 Result:", price) // 1457298
}
