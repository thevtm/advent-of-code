package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"sort"

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
	// input_path = "./input-sample-small.txt"
	// input_path = "./input-sample-large.txt"
	// input_path = "./input-sample-problem-2.txt"

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

	start := time.Now()

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

	fmt.Println("Problem 1 Result:", boxes_coordinates_sum, "●", time.Since(start)) // 1514333 ● 2.279269435s

	fmt.Println()
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println()

	// PROBLEM 1

	start = time.Now()

	var push_boxes_horizontal func(position Position, direction rune) bool
	push_boxes_horizontal = func(position Position, direction rune) bool {
		tile := get_warehouse_tile(position)

		if tile == '#' {
			return false
		}

		if tile == '.' {
			return true
		}

		if tile == '[' || tile == ']' {
			next_position := position.plus(directions[direction])

			if push_boxes_horizontal(next_position, direction) == false {
				return false
			}

			set_warehouse_tile(position, '.')
			set_warehouse_tile(next_position, tile)

			return true
		}

		panic(fmt.Sprintf("Unknown tile: %s", string(tile)))
	}

	var push_boxes_vertical func(starting_position Position, direction rune) bool
	push_boxes_vertical = func(starting_position Position, direction rune) bool {
		fmt.Println("[push_boxes_vertical]", "starting_position", starting_position, "direction", string(direction))

		// 1. Check if we can move
		visited_positions := make(map[Position]bool)

		stack := make([]Position, 0)
		stack = append(stack, starting_position)

		boxes_positions := make([]Position, 0)

		for len(stack) > 0 {
			position := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if visited_positions[position] {
				fmt.Println("\t", "position", position, "already visited")
				continue
			}

			tile := get_warehouse_tile(position)

			fmt.Println("\t", "position", position, "tile", string(tile))

			if tile == '.' {
				continue
			}

			if tile == '#' {
				fmt.Println("\t", "Blocked!")
				return false
			}

			if tile == ']' || tile == '[' {
				var left_side_pos Position
				var right_side_pos Position

				if tile == '[' {
					left_side_pos = position
					right_side_pos = position.plus(directions['>'])
				} else {
					left_side_pos = position.plus(directions['<'])
					right_side_pos = position
				}

				visited_positions[left_side_pos] = true
				visited_positions[right_side_pos] = true

				stack = append(stack, left_side_pos.plus(directions[direction]))
				stack = append(stack, right_side_pos.plus(directions[direction]))

				boxes_positions = append(boxes_positions, left_side_pos)

				continue
			}

			panic(fmt.Sprintf("Unknown tile: %s", string(tile)))
		}

		// 2. Move boxes (must process boxes furthest in the direction of movement first)
		sort.Slice(boxes_positions, func(i, j int) bool {
			if direction == '^' {
				return boxes_positions[i].y < boxes_positions[j].y
			}
			return boxes_positions[i].y > boxes_positions[j].y
		})

		for i := 0; i < len(boxes_positions); i++ {
			left_position := boxes_positions[i]
			right_position := left_position.plus(directions['>'])

			next_left_position := left_position.plus(directions[direction])
			next_right_position := next_left_position.plus(directions['>'])

			set_warehouse_tile(next_left_position, get_warehouse_tile(left_position))
			set_warehouse_tile(next_right_position, get_warehouse_tile(right_position))

			set_warehouse_tile(left_position, '.')
			set_warehouse_tile(right_position, '.')
		}

		return true
	}

	push_boxes := func(robot_position Position, direction rune) bool {
		switch direction {
		case '<', '>':
			return push_boxes_horizontal(robot_position, direction)
		case '^', 'v':
			return push_boxes_vertical(robot_position, direction)
		}

		panic(fmt.Sprintf("Unknown direction: %s", string(direction)))
	}

	// 1. Modify Warehouse
	warehouse = make([][]rune, len(warehouse_lines))

	for i, line := range warehouse_lines {
		new_line := strings.ReplaceAll(line, "O", "[]")
		new_line = strings.ReplaceAll(new_line, ".", "..")
		new_line = strings.ReplaceAll(new_line, "@", "@.")
		new_line = strings.ReplaceAll(new_line, "#", "##")

		warehouse[i] = []rune(new_line)
	}

	warehouse_width *= 2

	print_rune_grid(warehouse)
	println()

	// 2. Execute moves
	robot_position = robot_starting_position
	robot_position.x *= 2

	for i, direction := range moves {
		fmt.Println("Move", i, string(direction), "robot_position", robot_position)

		next_robot_position := robot_position.plus(directions[direction])

		if push_boxes(next_robot_position, direction) {
			set_warehouse_tile(robot_position, '.')
			set_warehouse_tile(next_robot_position, '@')

			fmt.Printf("Moved from %v to %v\n", robot_position, next_robot_position)

			robot_position = next_robot_position
		} else {
			fmt.Println("Unable to move.")
		}

		print_rune_grid(warehouse)

		fmt.Println()
	}

	// 3. Calculate GPS coordinates
	boxes_coordinates_sum = 0

	for y, row := range warehouse {
		for x, cell := range row {
			if cell == '[' {
				boxes_coordinates_sum += y*100 + x
			}
		}
	}

	fmt.Println()

	fmt.Println("Problem 2 Result:", boxes_coordinates_sum, "●", time.Since(start)) // 1528453 ● 2.705151983s
}
