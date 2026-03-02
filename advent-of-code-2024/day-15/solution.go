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

func print_rune_grid(grid [][]rune) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func main() {
	// INPUTS

	input_path := "./input.txt"
	// input_path := "./input-sample-small.txt"
	// input_path := "./input-sample-large.txt"

	_, source_file_path, _, _ := runtime.Caller(0)
	source_dir := filepath.Dir(source_file_path)

	file_content := string(lo.Must(os.ReadFile(filepath.Join(source_dir, input_path))))

	fmt.Println(file_content)
	fmt.Println()

	file_lines := strings.Split(file_content, "\n")
	file_lines = file_lines[:len(file_lines)-1] // Last line is blank

	warehouse_lines := file_lines[:len(file_lines)-2]
	warehouse := make([][]rune, len(warehouse_lines))

	for i, line := range warehouse_lines {
		warehouse[i] = []rune(line)
	}

	fmt.Println(warehouse)
	fmt.Println()

	warehouse_width := len(warehouse[0])
	warehouse_height := len(warehouse)

	fmt.Printf("warehouse => width: %d, height: %d\n\n", warehouse_width, warehouse_height)

	robot_starting_position := Position{0, 0}

	for y, row := range warehouse {
		for x, cell := range row {
			if cell == '@' {
				robot_starting_position.x = x
				robot_starting_position.y = y

				break
			}
		}
	}

	fmt.Printf("starting_position %+v\n\n", robot_starting_position)

	moves := []rune(file_lines[len(file_lines)-1])

	fmt.Printf("moves %+v\n", moves)

	fmt.Println()
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println()

	// PROBLEM 1

	get_warehouse_tile := func(position Position) rune { return warehouse[position.y][position.x] }
	set_warehouse_tile := func(position Position, value rune) { warehouse[position.y][position.x] = value }

	directions := map[rune]Position{'^': {0, -1}, 'v': {0, 1}, '<': {-1, 0}, '>': {1, 0}}

	robot_position := robot_starting_position

	for i, move := range moves {
		fmt.Println(i, string(move))
		print_rune_grid(warehouse)
		fmt.Println()

		direction := directions[move]

		next_robot_position := robot_position.plus(direction)
		tile := get_warehouse_tile(next_robot_position)

		if tile == '.' {
			set_warehouse_tile(robot_position, '.')
			set_warehouse_tile(next_robot_position, '@')
			robot_position = next_robot_position
			continue
		}

		if tile == '#' {
			continue
		}

		if tile != 'O' { // Box should be the only valid option left
			panic(fmt.Sprintf("Unknown tile: %c", tile))
		}

		invalid_next_empty_spot := Position{-1, -1}
		next_empty_spot := next_robot_position.plus(direction)

		for {
			next_empty_spot_tile := get_warehouse_tile(next_empty_spot)

			if next_empty_spot_tile == '.' {
				break
			}

			if next_empty_spot_tile == 'O' {
				next_empty_spot = next_empty_spot.plus(direction)
				continue
			}

			if next_empty_spot_tile == '#' {
				next_empty_spot = invalid_next_empty_spot
				break
			}
		}

		if next_empty_spot == invalid_next_empty_spot {
			continue
		}

		set_warehouse_tile(robot_position, '.')
		set_warehouse_tile(next_robot_position, '@')
		set_warehouse_tile(next_empty_spot, 'O')

		robot_position = next_robot_position
	}

	boxes_coordinates_sum := 0

	for y, row := range warehouse {
		for x, cell := range row {
			if cell == 'O' {
				boxes_coordinates_sum += y*100 + x
			}
		}
	}

	fmt.Println("Problem 1 Result:", boxes_coordinates_sum) // 1514333
}
