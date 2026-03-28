package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/samber/lo"
)

func main() {
	// INPUTS

	input_path := "./input.txt"
	// input_path = "./input-sample.txt"

	_, source_file_path, _, _ := runtime.Caller(0)
	source_dir := filepath.Dir(source_file_path)

	file_content := string(lo.Must(os.ReadFile(filepath.Join(source_dir, input_path))))

	fmt.Println(file_content)
	file_content = file_content[:len(file_content)-1] // Remove new line

	fmt.Println()

	digit_map := map[byte]int{'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9}

	disk := make([]int, len(file_content))

	for i := 0; i < len(file_content); i++ {
		disk[i] = digit_map[file_content[i]]
	}

	fmt.Println("disk", disk)

	fmt.Println()
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println()

	// PROBLEM 1

	start := time.Now()

	index := 0
	left_block_index := 0
	right_file_block_index := len(disk) - ((len(disk) - 1) % 2) - 1

	disk_used := make([]int, len(disk))
	copy(disk_used, disk)

	checksum := int64(0)

	for right_file_block_index >= left_block_index {
		// fmt.Println("disk_used", disk_used)
		// fmt.Println("\t", "index", index, "left_block_index", left_block_index, "right_file_block_index", right_file_block_index)

		if disk_used[left_block_index] == 0 {
			left_block_index++
			continue
		}

		if disk_used[right_file_block_index] == 0 {
			right_file_block_index -= 2
			continue
		}

		is_empty_sector := (left_block_index % 2) == 1

		if is_empty_sector {
			disk_used[left_block_index]--
			disk_used[right_file_block_index]--

			file_id := right_file_block_index / 2

			checksum += int64(index * file_id)
		} else {
			disk_used[left_block_index]--

			file_id := left_block_index / 2

			checksum += int64(index * file_id)
		}

		index++
	}

	fmt.Println()

	fmt.Println("Problem 1 Result:", checksum, "●", time.Since(start)) // 6399153661894 ● 209.369µs

	fmt.Println()
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println()

	start = time.Now()

	index = 0
	left_block_index = 0
	last_file_block_index := len(disk) - ((len(disk) - 1) % 2) - 1

	disk_used = make([]int, len(disk))
	copy(disk_used, disk)

	checksum = int64(0)

	for left_block_index <= last_file_block_index {
		is_empty_sector := (left_block_index % 2) == 1

		if disk_used[left_block_index] < 0 {
			fmt.Println(">", index, "skip", -disk_used[left_block_index])
			index += -disk_used[left_block_index]
			left_block_index++
			continue
		}

		if is_empty_sector {
			empty_block_size := disk_used[left_block_index]
			right_file_block_index := last_file_block_index

			for right_file_block_index > left_block_index &&
				(disk_used[right_file_block_index] < 0 || disk_used[right_file_block_index] > empty_block_size) {
				right_file_block_index -= 2
			}

			no_file_found := right_file_block_index <= left_block_index

			// fmt.Println("no_file_found", no_file_found, "right_file_block_index", right_file_block_index)

			if no_file_found {
				fmt.Println(">", index, "no_file_found")
				left_block_index++
				index += empty_block_size
			} else {
				file_id := right_file_block_index / 2
				file_size := disk_used[right_file_block_index]

				fmt.Println(">", index, "file_id", file_id, "file_size", file_size)

				for i := 0; i < file_size; i++ {
					checksum += int64(index * file_id)
					index++
				}

				disk_used[left_block_index] = empty_block_size - file_size
				disk_used[right_file_block_index] = -file_size

				if disk_used[left_block_index] == 0 {
					left_block_index++
				}
			}
		} else {
			file_id := left_block_index / 2
			file_size := disk_used[left_block_index]

			fmt.Println(">", index, "file_id", file_id, "file_size", file_size)

			for i := 0; i < file_size; i++ {
				checksum += int64(index * file_id)
				index++
			}

			left_block_index++
		}
	}

	fmt.Println("Problem 2 Result:", checksum, "●", time.Since(start)) // 6421724645083 ● 60.769ms
}
