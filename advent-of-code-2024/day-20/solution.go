package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/samber/lo"
	pq "github.com/ugurcsen/gods-generic/queues/priorityqueue"
	"github.com/ugurcsen/gods-generic/sets/hashset"
	"github.com/ugurcsen/gods-generic/utils"
)

type Turn = [2]rune

type Point struct {
	x, y int
}

type ReindeerStep struct {
	position  Point
	direction rune
	score     int
}

func (p Point) plus(other Point) Point {
	return Point{p.x + other.x, p.y + other.y}
}

func print_rune_grid(grid [][]rune) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
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

	race_track := make([][]rune, len(file_lines))

	for i, line := range file_lines {
		race_track[i] = []rune(line)
	}

	fmt.Println(race_track)
	fmt.Println()

	race_track_width := len(race_track[0])
	race_track_height := len(race_track)

	fmt.Printf("width: %d, height: %d\n\n", race_track_width, race_track_height)

	start_position := Point{0, 0}
	end_position := Point{0, 0}

	for y, row := range race_track {
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

	// PART 1

	get_maze_tile := func(position Point) rune { return race_track[position.y][position.x] }
	movement_directions := []Point{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	racetrack_distance_by_position := make(map[Point]int)

	visited := hashset.New[Point]()

	queue_priority := func(a, b Point) int {
		priorityA := racetrack_distance_by_position[a]
		priorityB := racetrack_distance_by_position[b]
		return utils.NumberComparator[int](priorityA, priorityB)
	}

	queue := pq.NewWith[Point](queue_priority)
	queue.Enqueue(start_position)

	racetrack_distance_by_position[start_position] = 0

	for !queue.Empty() {
		position, _ := queue.Dequeue()

		if visited.Contains(position) {
			continue
		}
		visited.Add(position)

		distance := racetrack_distance_by_position[position]

		for _, movement := range movement_directions {
			next_position := position.plus(movement)
			next_position_tile := get_maze_tile(next_position)

			if next_position_tile == '#' || next_position_tile == 'S' {
				continue
			}

			if next_position_tile != '.' && next_position_tile != 'E' {
				panic(fmt.Sprintf("Unknown tile: %c", next_position_tile))
			}

			next_distance := distance + 1

			if current_distance, ok := racetrack_distance_by_position[next_position]; ok && current_distance <= next_distance {
				continue
			}

			racetrack_distance_by_position[next_position] = next_distance

			if next_position_tile == 'E' {
				continue
			}

			queue.Enqueue(next_position)
		}
	}
	fmt.Println()

	fmt.Printf("%+v\n\n", racetrack_distance_by_position)

	uber_cheat_count := 0

	for position, distance := range racetrack_distance_by_position {
		for _, movement_1 := range movement_directions {
			for _, movement_2 := range movement_directions {
				cheat_position := position.plus(movement_1).plus(movement_2)

				cheat_position_distance, ok := racetrack_distance_by_position[cheat_position]

				if !ok {
					continue
				}

				if distance >= cheat_position_distance {
					continue
				}

				delta_distance := (cheat_position_distance - 2) - distance

				if delta_distance < 100 {
					continue
				}

				fmt.Printf("%+v => %+v = %d\n", position, cheat_position, delta_distance)
				uber_cheat_count++
			}
		}
	}
	fmt.Println()

	fmt.Println("Part 1 Answer:", uber_cheat_count) // 105508
}
