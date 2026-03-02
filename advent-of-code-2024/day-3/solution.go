package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func problem_1(file_content string) {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	str := file_content

	result := 0

	for i, match := range re.FindAllSubmatch([]byte(str), -1) {
		fmt.Println(i, "=>", "0:", string(match[0]), "1:", string(match[1]), "2:", string(match[2]))

		a := lo.Must(strconv.Atoi(string(match[1])))
		b := lo.Must(strconv.Atoi(string(match[2])))

		result += a * b
	}
	fmt.Println()

	fmt.Println("Problem 1 Result:", result) // 167090022
}

func problem_2(file_content string) {
	re := regexp.MustCompile(`(mul\((\d+),(\d+)\))|(do\(\))|(don't\(\))`)
	str := file_content

	result := 0
	mul_enabled := true

	for _, submatches := range re.FindAllSubmatch([]byte(str), -1) {
		whole_match := string(submatches[0])

		if mul_enabled && strings.HasPrefix(whole_match, "mul") {
			a := lo.Must(strconv.Atoi(string(submatches[2])))
			b := lo.Must(strconv.Atoi(string(submatches[3])))
			result += a * b

		} else if strings.HasPrefix(whole_match, "don't") {
			mul_enabled = false
		} else if strings.HasPrefix(whole_match, "do") {
			mul_enabled = true
		}
	}

	fmt.Println("Problem 2 Result:", result) // 167090022
}

func main() {
	// INPUTS

	input_path := "./input.txt"
	// input_path := "./input-sample.txt"
	// input_path := "./input-sample-2.txt"

	_, source_file_path, _, _ := runtime.Caller(0)
	source_dir := filepath.Dir(source_file_path)

	file_content := string(lo.Must(os.ReadFile(filepath.Join(source_dir, input_path))))

	fmt.Println(file_content)
	fmt.Println()

	problem_1(file_content)

	fmt.Println()
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println()

	problem_2(file_content)
}
