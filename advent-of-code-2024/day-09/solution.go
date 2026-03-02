package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

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

	digit_map := map[byte]int{'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9}

	var disk [100000]int
	empty_sector_id := -1
	index := 0

	for i := 0; i < len(file_content)/2; i++ {
		file_size := digit_map[file_content[i*2]]
		free_size := digit_map[file_content[(i*2)+1]]

		for j := 0; j < file_size; j++ {
			disk[index] = i
			index++
		}

		for j := 0; j < free_size; j++ {
			disk[index] = empty_sector_id
			index++
		}
	}

	fmt.Println(disk[:index])

	fmt.Println()

	fmt.Println("index", index)

	fmt.Println()
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println()

	// PROBLEM 1

	left_index := 0
	right_index := index - 1

	for {
		for ; disk[left_index] != empty_sector_id; left_index++ {
		}

		for ; disk[right_index] == empty_sector_id; right_index-- {
		}

		if left_index < right_index {
			break
		}

		disk[left_index] = disk[right_index]
		disk[right_index] = empty_sector_id

		// fmt.Println(left_index, right_index, "\t", disk[:index])

		left_index++
		right_index--
	}

	fmt.Println()

	fmt.Println(disk[:index])

	fmt.Println()

	var checksum int64 = 0

	for i := 0; disk[i] != empty_sector_id; i++ {
		checksum += int64(i) * int64(disk[i])
	}

	fmt.Println("Problem 1 Result:", checksum) // 6399153661894
}
