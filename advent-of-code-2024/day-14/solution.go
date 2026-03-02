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

type Robot struct {
	x, y   int
	vx, vy int
}

func main() {
	// INPUTS

	width := 101
	height := 103
	input_path := "./input.txt"

	// width := 11
	// height := 7
	// input_path := "./input-sample.txt"

	_, source_file_path, _, _ := runtime.Caller(0)
	source_dir := filepath.Dir(source_file_path)

	file_content := string(lo.Must(os.ReadFile(filepath.Join(source_dir, input_path))))

	fmt.Println(file_content)
	file_content = file_content[:len(file_content)-1] // Remove new line

	fmt.Println()

	re := regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)
	robots := make([]Robot, 0)

	for _, match := range re.FindAllSubmatch([]byte(file_content), -1) {
		cm := Robot{
			lo.Must(strconv.Atoi(string(match[1]))),
			lo.Must(strconv.Atoi(string(match[2]))),
			lo.Must(strconv.Atoi(string(match[3]))),
			lo.Must(strconv.Atoi(string(match[4]))),
		}

		robots = append(robots, cm)
	}

	fmt.Println(robots)
	fmt.Println()

	mid_width := (width / 2)
	mid_height := (height / 2)

	fmt.Println("width", width, "height", height, "mid_width", mid_width, "mid_height", mid_height)

	fmt.Println()
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println()

	// PROBLEM 1

	// tokens := 0
	seconds := 100

	for i := range len(robots) {
		robot := &robots[i]
		robot.x = (robot.x + (robot.vx * seconds) + (width * seconds)) % width
		robot.y = (robot.y + (robot.vy * seconds) + (height * seconds)) % height
	}
	fmt.Println()

	fmt.Println(robots)
	fmt.Println()

	quadrants := []int{0, 0, 0, 0}

	for _, robot := range robots {
		is_left := robot.x < mid_width
		is_right := robot.x > mid_width
		is_top := robot.y < mid_height
		is_bottom := robot.y > mid_height

		fmt.Println(robot, is_left, is_right, is_top, is_bottom)

		if is_top && is_left {
			fmt.Println(0)
			quadrants[0]++
		} else if is_top && is_right {
			fmt.Println(1)
			quadrants[1]++
		} else if is_bottom && is_left {
			fmt.Println(2)
			quadrants[2]++
		} else if is_bottom && is_right {
			fmt.Println(3)
			quadrants[3]++
		}
	}
	fmt.Println()

	safety_factor := quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]

	fmt.Println("quadrants", quadrants, "safety_factor", safety_factor)
	fmt.Println()

	fmt.Println("Problem 1 Result:", safety_factor) // 226236192
}
