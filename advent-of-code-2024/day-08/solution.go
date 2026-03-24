package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/samber/lo"
)

type Point struct {
	x, y int
}

type Line struct {
	a, b Point
}

type Antenna struct {
	position  Point
	frequency rune
}

func pointInLine(point Point, line Line) bool {
	// cross product of (B-A) x (P-A) == 0 means collinear
	cross := (line.b.x-line.a.x)*(point.y-line.a.y) - (line.b.y-line.a.y)*(point.x-line.a.x)
	return cross == 0
}

func main() {
	// INPUTS

	input_path := "./input.txt"
	// input_path = "./input-sample.txt"

	_, source_file_path, _, _ := runtime.Caller(0)
	source_dir := filepath.Dir(source_file_path)

	file_content := string(lo.Must(os.ReadFile(filepath.Join(source_dir, input_path))))

	fmt.Println(file_content)
	fmt.Println()

	file_lines := strings.Split(file_content, "\n")
	file_lines = file_lines[:len(file_lines)-1] // Last line is blank

	antennas := make([]Antenna, 0)
	taken_positions := make(map[Point]bool)

	for y, line := range file_lines {
		for x, cell := range line {
			if cell != '.' {
				taken_positions[Point{x, y}] = true
				antennas = append(antennas, Antenna{Point{x, y}, cell})
			}
		}
	}

	map_width := len(file_lines)
	map_height := len(file_lines[0])

	fmt.Println("width", map_width, "height", map_height)

	fmt.Println()

	fmt.Println(antennas)

	fmt.Println()
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println()

	// PROBLEM 1

	start := time.Now()

	is_in_bounds := func(pos Point) bool {
		if pos.x < 0 || pos.x >= map_width || pos.y < 0 || pos.y >= map_height {
			return false
		}

		return true
	}

	anti_node_unique_positions := make(map[Point]bool)

	antennas_by_frequency := lo.GroupBy(antennas, func(antenna Antenna) rune {
		return antenna.frequency
	})

	fmt.Println(antennas_by_frequency)

	for _, f := range antennas_by_frequency {
		for i := 0; i < len(f)-1; i++ {
			a := f[i]

			for j := i + 1; j < len(f); j++ {
				b := f[j]

				x := a.position.x - b.position.x
				y := a.position.y - b.position.y

				aa := Point{a.position.x + x, a.position.y + y}
				ab := Point{b.position.x - x, b.position.y - y}

				if is_in_bounds(aa) {
					anti_node_unique_positions[aa] = true
				}

				if is_in_bounds(ab) {
					anti_node_unique_positions[ab] = true
				}
			}
		}
	}

	fmt.Println()

	fmt.Println("Problem 1 Result:", len(anti_node_unique_positions), "●", time.Since(start)) // 256 ● 117.573µs

	fmt.Println()
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println()

	// Problem 2

	start = time.Now()

	lines := make([]Line, 0)

	for _, antennas := range antennas_by_frequency {
		for i := 0; i < len(antennas)-1; i++ {
			for j := i + 1; j < len(antennas); j++ {
				lines = append(lines, Line{antennas[i].position, antennas[j].position})
			}
		}
	}

	anti_node_unique_positions = make(map[Point]bool)

	for i := 0; i < map_width; i++ {
		for j := 0; j < map_height; j++ {
			point := Point{i, j}

			for _, line := range lines {
				if pointInLine(point, line) {
					anti_node_unique_positions[point] = true
					break
				}
			}
		}
	}

	fmt.Println("Problem 2 Result:", len(anti_node_unique_positions), "●", time.Since(start)) // 1005 ● 832.328µs
}
