package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type Step struct {
	acc       int
	operation rune
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

	equations := make([][]int, len(file_lines))
	re := regexp.MustCompile(`\d+`)

	for i, line := range file_lines {
		equation := make([]int, 0)

		for _, number_str := range re.FindAllString(line, -1) {
			equation = append(equation, lo.Must(strconv.Atoi(number_str)))
		}

		equations[i] = equation
	}

	fmt.Println(equations)

	fmt.Println()

	// PROBLEM 1

	total := 0

	for _, equation := range equations {
		result := equation[0]
		coefficients := equation[1:]

		num_possibilities := int(math.Pow(2, float64(len(coefficients)-1)))

		fmt.Println(equation, num_possibilities)

		for permutation := range num_possibilities {
			acc := coefficients[0]

			for coefficient_index, coefficient := range coefficients[1:] {
				bit := (permutation >> coefficient_index) & 1 // Right shift and mask with 1

				if bit == 1 {
					acc += coefficient
				} else {
					acc *= coefficient
				}

				// fmt.Println("bit", bit, "permutation", permutation, "coefficient", coefficient, "acc", acc)
			}

			if acc == result {
				total += result
				break
			}
		}
	}

	fmt.Println("Problem 1 Result:", total) // 1985268524462
}
