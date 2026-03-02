package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/emirpasic/gods/sets/treeset"
	"github.com/emirpasic/gods/stacks/arraystack"
	"github.com/samber/lo"
)

type TrieNode struct {
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
	file_lines = file_lines[:len(file_lines)-1] // Last element is blank

	// Patterns
	patterns := make(map[string]bool)
	re := regexp.MustCompile(`\w+`)

	for _, match := range re.FindAllString(file_lines[0], -1) {
		patterns[match] = true
	}

	fmt.Printf("patterns: %+v\n\n", patterns)

	// Designs
	designs := make([]string, 0)

	for _, line := range file_lines[2:] {
		designs = append(designs, line)
	}

	fmt.Printf("designs: %+v\n", designs)

	fmt.Println()
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println()

	// PART 1

	longest_pattern := 0

	for pattern := range patterns {
		longest_pattern = int(math.Max(float64(longest_pattern), float64(len(pattern))))
	}

	valid_patterns_count := 0

OuterLoopValidPatterns:
	for _, design := range designs {
		fmt.Printf("%s\n", design)

		set := treeset.NewWithIntComparator()
		set.Add(0)

		it := set.Iterator()

		for it.Next() {
			value := it.Value().(int)

			for i := 1; i <= longest_pattern && (value+i) <= len(design); i++ {
				end_index := value + i
				design_sub := design[value:end_index]

				_, ok := patterns[design_sub]

				if !ok {
					continue
				}

				if end_index == len(design) {
					fmt.Printf("\tGood!\n")
					valid_patterns_count++
					continue OuterLoopValidPatterns
				}

				set.Add(end_index)
			}
		}

		fmt.Printf("\tset: %+v\n", set)
	}
	fmt.Println()

	fmt.Printf("Part 1 Answer: %d\n", valid_patterns_count) // 369

	fmt.Println()
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println()

	// PART 2

	total_arrangements_count := 0

	for _, design := range designs {
		fmt.Printf("%s\n", design)

		stack := arraystack.New()
		stack.Push(0)

		for !stack.Empty() {
			value, _ := stack.Pop()
			start_index := value.(int)

			for i := 1; i <= longest_pattern && (start_index+i) <= len(design); i++ {
				end_index := start_index + i
				design_sub := design[start_index:end_index]

				_, found_pattern := patterns[design_sub]

				if !found_pattern {
					continue
				}

				if end_index == len(design) {
					// fmt.Printf("\tGood!\n")
					total_arrangements_count++
				}

				stack.Push(end_index)
			}
		}
	}
	fmt.Println()

	fmt.Printf("Part 2 Answer: %d\n", total_arrangements_count)
}
