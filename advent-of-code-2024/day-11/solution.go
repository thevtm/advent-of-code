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
	file_content = file_content[:len(file_content)-1] // Remove new line

	fmt.Println()

	stones := make([]int, 0)
	re := regexp.MustCompile(`\d+`)

	for _, number_str := range re.FindAllString(file_content, -1) {
		stones = append(stones, lo.Must(strconv.Atoi(number_str)))
	}

	fmt.Println(stones)

	fmt.Println()
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println()

	// PROBLEM 1

	stones_map := make(map[int]int)

	for _, stone := range stones {
		stones_map[stone] += 1
	}

	for blink := range 25 {
		new_stones_map := make(map[int]int)

		for stone, count := range stones_map {

			// 0 becomes 1
			if stone == 0 {
				new_stones_map[1] += count
				continue
			}

			// Split if even number of digits
			stone_str := fmt.Sprintf("%d", stone)
			stone_str_len := len(stone_str)

			if stone_str_len%2 == 0 {
				half_point := stone_str_len / 2
				left_stone := lo.Must(strconv.Atoi(stone_str[:half_point]))
				right_stone := lo.Must(strconv.Atoi(stone_str[half_point:]))

				new_stones_map[left_stone] += count
				new_stones_map[right_stone] += count

				continue
			}

			// Multiply by 2024
			new_stones_map[stone*2024] += count
		}

		fmt.Println("blink", blink, len(new_stones_map), new_stones_map)

		stones_map = new_stones_map
	}

	stone_count := 0

	for _, count := range stones_map {
		stone_count += count
	}

	fmt.Println("Problem 1 Result:", stone_count) // 198075
}
