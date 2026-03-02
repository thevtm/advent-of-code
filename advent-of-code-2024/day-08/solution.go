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

type Antenna struct {
	position  Position
	frequency rune
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
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

	antennas := make([]Antenna, 0)
	taken_positions := make(map[Position]bool)

	for y, line := range file_lines {
		for x, cell := range line {
			if cell != '.' {
				taken_positions[Position{x, y}] = true
				antennas = append(antennas, Antenna{Position{x, y}, cell})
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

	is_in_bounds := func(pos Position) bool {
		if pos.x < 0 || pos.x >= map_width || pos.y < 0 || pos.y >= map_height {
			return false
		}

		return true
	}

	anti_node_unique_positions := make(map[Position]bool)

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

				aa := Position{a.position.x + x, a.position.y + y}
				ab := Position{b.position.x - x, b.position.y - y}

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

	fmt.Println("Problem 1 Result:", len(anti_node_unique_positions)) // 256
}
