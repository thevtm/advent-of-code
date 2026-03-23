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
	"time"

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

	// fmt.Println(file_content)
	// fmt.Println()

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

	// fmt.Println(equations)

	// fmt.Println()

	// PROBLEM 1

	start := time.Now()
	total := 0

	for _, equation := range equations {
		result := equation[0]
		coefficients := equation[1:]

		num_possibilities := int(math.Pow(2, float64(len(coefficients)-1)))

		// fmt.Println(equation, num_possibilities)

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

	fmt.Println()
	fmt.Println("Problem 1 Result:", total, "●", time.Since(start)) // 198526852446 ● 27.446577ms

	fmt.Println()
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println()

	// PROBLEM 2

	start = time.Now()
	total = 0

	for _, equation := range equations {
		result := equation[0]
		coefficients := equation[1:]

		num_possibilities := int(math.Pow(3, float64(len(coefficients)-1)))

		// fmt.Println(equation, num_possibilities)

		for permutation := range num_possibilities {
			acc := coefficients[0]

			for _, coefficient := range coefficients[1:] {
				operation := permutation % 3
				permutation /= 3

				if operation == 0 {
					acc += coefficient
				} else if operation == 1 {
					acc *= coefficient
				} else {
					digits := math.Floor(math.Log10(float64(coefficient))) + 1
					acc = acc*int(math.Pow(10, digits)) + coefficient
				}

				// fmt.Println("\t", "operation", operation, "permutation", permutation, "coefficient", coefficient, "acc", acc)
			}

			if acc == result {
				total += result
				break
			}
		}
	}

	fmt.Println("Problem 2 Result:", total, "●", time.Since(start)) // 150077710195188 ● 1.111534997s
}
