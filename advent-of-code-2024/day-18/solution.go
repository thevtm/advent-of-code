package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"

	"github.com/samber/lo"
)

type Point struct {
	x, y int
}

func (p Point) plus(other Point) Point {
	return Point{p.x + other.x, p.y + other.y}
}

func main() {
	// INPUTS

	input_path := "./input.txt"
	width := 70 + 1
	height := 70 + 1
	time_frame := 1024

	// input_path := "./input-sample.txt"
	// width := 6 + 1
	// height := 6 + 1
	// time_frame := 12

	_, source_file_path, _, _ := runtime.Caller(0)
	source_dir := filepath.Dir(source_file_path)

	file_content := string(lo.Must(os.ReadFile(filepath.Join(source_dir, input_path))))

	fmt.Println(file_content)
	fmt.Println()

	re := regexp.MustCompile(`(\d+),(\d+)`)

	falling_bytes := make([]Point, 0)

	for _, match := range re.FindAllSubmatch([]byte(file_content), -1) {
		falling_bytes = append(falling_bytes, Point{
			lo.Must(strconv.Atoi(string(match[1]))),
			lo.Must(strconv.Atoi(string(match[2]))),
		})
	}

	fmt.Printf("width: %d height: %d\n", width, height)

	fmt.Println()
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println()

	// PART 1

	part_1(falling_bytes, time_frame, width, height)

	fmt.Println()
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println()

	// PART 2

	part_2(falling_bytes, width, height)

}

func part_1(falling_bytes []Point, time_frame int, width int, height int) {
	memory := make(map[Point]bool, 0)

	for _, pos := range falling_bytes[:time_frame] {
		memory[pos] = true
	}

	fmt.Printf("memory: %+v\n\n", memory)

	is_blocked := func(pos Point) bool {
		// Is within bounds
		if pos.x < 0 || pos.x >= width || pos.y < 0 || pos.y >= height {
			return true
		}

		_, ok := memory[pos]
		return ok
	}

	visited := make(map[Point]bool)
	stack := make([]Point, 0)
	scores := make(map[Point]int)

	start_position := Point{0, 0}
	end_position := Point{width - 1, height - 1}

	stack = append(stack, start_position)
	scores[start_position] = 0

	for len(stack) > 0 {
		pos := stack[0]
		stack = stack[1:]

		if _, ok := visited[pos]; ok {
			continue
		}
		visited[pos] = true

		score := scores[pos]

		for _, movement_direction := range []Point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}} {
			next_pos := pos.plus(movement_direction)
			next_score := score + 1

			if is_blocked(next_pos) {
				continue
			}

			if current_score, ok := scores[next_pos]; ok && current_score <= next_score {
				continue
			}

			scores[next_pos] = next_score

			if next_pos == end_position {
				continue
			}

			stack = append(stack, next_pos)
		}
	}

	fmt.Println("Part 1 Result:", scores[end_position]) // 356
}

func part_2(falling_bytes []Point, width int, height int) {
	start_position := Point{0, 0}
	end_position := Point{width - 1, height - 1}

	memory := make(map[Point]bool, 0)

	is_blocked := func(pos Point) bool {
		// Is within bounds
		if pos.x < 0 || pos.x >= width || pos.y < 0 || pos.y >= height {
			return true
		}

		_, ok := memory[pos]
		return ok
	}

	is_reachable := func() bool {
		visited := make(map[Point]bool)
		stack := make([]Point, 0)
		scores := make(map[Point]int)

		stack = append(stack, start_position)
		scores[start_position] = 0

		for len(stack) > 0 {
			pos := stack[0]
			stack = stack[1:]

			if _, ok := visited[pos]; ok {
				continue
			}
			visited[pos] = true

			score := scores[pos]

			for _, movement_direction := range []Point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}} {
				next_pos := pos.plus(movement_direction)
				next_score := score + 1

				if is_blocked(next_pos) {
					continue
				}

				if current_score, ok := scores[next_pos]; ok && current_score <= next_score {
					continue
				}

				scores[next_pos] = next_score

				if next_pos == end_position {
					continue
				}

				stack = append(stack, next_pos)
			}
		}

		_, reachable := scores[end_position]
		return reachable
	}

	blocking_byte := Point{}

	for i, falling_byte := range falling_bytes {
		memory[falling_byte] = true

		if is_reachable() {
			fmt.Printf("[%d] %+v is reachable\n", i, falling_byte)
			continue
		}

		blocking_byte = falling_byte
		break
	}
	fmt.Println()

	fmt.Printf("Part 2 Result: %d,%d\n", blocking_byte.x, blocking_byte.y) // 22,33
}
