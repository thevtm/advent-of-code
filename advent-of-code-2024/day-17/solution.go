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

	re := regexp.MustCompile(`\d+`)

	register_a := 0
	register_b := 0
	register_c := 0

	program := make([]int, 0)

	for i, match := range re.FindAllString(file_content, -1) {
		if i == 0 {
			register_a = lo.Must(strconv.Atoi(match))
		} else if i == 1 {
			register_b = lo.Must(strconv.Atoi(match))
		} else if i == 2 {
			register_c = lo.Must(strconv.Atoi(match))
		} else {
			program = append(program, lo.Must(strconv.Atoi(match)))
		}
	}

	fmt.Printf("A: %d B: %d C: %d\n", register_a, register_b, register_c)
	fmt.Printf("Program: %+v\n", program)

	fmt.Println()
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println()

	// PART 1

	instruction_pointer := 0
	output := make([]int, 0)

	combo_operand := func(operand int) int {
		if operand <= 3 {
			return operand
		} else if operand == 4 {
			return register_a
		} else if operand == 5 {
			return register_b
		} else if operand == 6 {
			return register_c
		} else {
			panic(fmt.Sprintf("Unknown combo operand %d", operand))
		}
	}

	for {
		if instruction_pointer >= len(program) {
			break
		}

		operator := program[instruction_pointer]
		operand := program[instruction_pointer+1]

		fmt.Printf("A: %d B: %d C: %d IP: %d Operator: %d Operand: %d\n", register_a, register_b, register_c, instruction_pointer, operator, operand)

		// 0 axl
		if operator == 0 {
			register_a = register_a / int(math.Pow(2, float64(combo_operand(operand))))

			instruction_pointer += 2
			continue
		}

		// 1 bxl
		if operator == 1 {
			register_b = register_b ^ operand

			instruction_pointer += 2
			continue
		}

		// 2 bst
		if operator == 2 {
			register_b = combo_operand(operand) % 8

			instruction_pointer += 2
			continue
		}

		// 3 jnz
		if operator == 3 {
			if register_a != 0 {
				instruction_pointer = operand
			} else {
				instruction_pointer += 2
			}

			continue
		}

		// 4 bxc
		if operator == 4 {
			register_b = register_b ^ register_c

			instruction_pointer += 2
			continue
		}

		// 5 out
		if operator == 5 {
			output = append(output, combo_operand(operand)%8)

			instruction_pointer += 2
			continue
		}

		// 6 bdv
		if operator == 6 {
			register_b = register_a / int(math.Pow(2, float64(combo_operand(operand))))

			instruction_pointer += 2
			continue
		}

		// 7 cdv
		if operator == 7 {
			register_c = register_a / int(math.Pow(2, float64(combo_operand(operand))))

			instruction_pointer += 2
			continue
		}

		panic(fmt.Sprintf("Unknown operator %d", operator))
	}
	fmt.Printf("A: %d B: %d C: %d IP: %d\n", register_a, register_b, register_c, instruction_pointer)
	fmt.Println()

	output_str := strings.Join(lo.Map(output, func(x int, _ int) string { return strconv.Itoa(x) }), ",")

	fmt.Println("Part 1 Result:", output_str) // 3,7,1,7,2,1,0,6,3
}
