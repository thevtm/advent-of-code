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

type ClawMachine struct {
	ax, ay int
	bx, by int
	x, y   int
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

	re := regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)\nButton B: X\+(\d+), Y\+(\d+)\nPrize: X=(\d+), Y=(\d+)`)
	claw_machines := make([]ClawMachine, 0)

	for _, match := range re.FindAllSubmatch([]byte(file_content), -1) {
		cm := ClawMachine{
			lo.Must(strconv.Atoi(string(match[1]))),
			lo.Must(strconv.Atoi(string(match[2]))),
			lo.Must(strconv.Atoi(string(match[3]))),
			lo.Must(strconv.Atoi(string(match[4]))),
			lo.Must(strconv.Atoi(string(match[5]))),
			lo.Must(strconv.Atoi(string(match[6]))),
		}

		claw_machines = append(claw_machines, cm)
	}

	fmt.Println(claw_machines)

	fmt.Println()
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println()

	// PROBLEM 1

	tokens := 0

	for _, cm := range claw_machines {
		l := (cm.bx * cm.ay) - (cm.by * cm.ax)
		r := (cm.x * cm.ay) - (cm.y * cm.ax)
		b := r / l
		a := (cm.x - (cm.bx * b)) / cm.ax

		fmt.Println("ax", cm.ax, "ay", cm.ay, "bx", cm.bx, "by", cm.by, "x", cm.x, "y", cm.y)
		fmt.Println("l", l, "r", r, "a", a, "b", b)

		if (a > 100) || (b > 100) {
			fmt.Println("over 100")
			fmt.Println()
			continue
		}

		x := a*cm.ax + b*cm.bx
		y := a*cm.ay + b*cm.by

		if cm.x != x && cm.y != y {
			fmt.Println("impossible", x, y)
			fmt.Println()
			continue
		}

		t := a*3 + b

		fmt.Println("Yay!", t)
		fmt.Println()

		tokens += t
	}

	fmt.Println("Problem 1 Result:", tokens) // 34787
}
