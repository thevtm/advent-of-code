package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"time"

	"github.com/samber/lo"
)

type ClawMachine struct {
	ax, ay int
	bx, by int
	px, py int
}

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

	start := time.Now()

	tokens := 0

	for _, cm := range claw_machines {
		// Cramer's rule
		l := (cm.bx * cm.ay) - (cm.by * cm.ax)
		r := (cm.px * cm.ay) - (cm.py * cm.ax)
		b := r / l
		a := (cm.px - (cm.bx * b)) / cm.ax

		fmt.Println("ax", cm.ax, "ay", cm.ay, "bx", cm.bx, "by", cm.by, "px", cm.px, "py", cm.py)
		fmt.Println("l", l, "r", r, "a", a, "b", b)

		if (a > 100) || (b > 100) {
			fmt.Println("over 100")
			fmt.Println()
			continue
		}

		x := a*cm.ax + b*cm.bx
		y := a*cm.ay + b*cm.by

		if cm.px != x || cm.py != y {
			fmt.Println("impossible", x, y)
			fmt.Println()
			continue
		}

		t := a*3 + b

		fmt.Println("Yay!", t)
		fmt.Println()

		tokens += t
	}

	fmt.Println("Problem 1 Result:", tokens, "●", time.Since(start)) // 34787 ● 6.58389ms

	fmt.Println()
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println()

	// PROBLEM 2

	start = time.Now()

	tokens = 0

	for _, cm := range claw_machines {
		cm.px += 10000000000000
		cm.py += 10000000000000

		// Cramer's rule
		l := (cm.bx * cm.ay) - (cm.by * cm.ax)
		r := (cm.px * cm.ay) - (cm.py * cm.ax)
		b := r / l
		a := (cm.px - (cm.bx * b)) / cm.ax

		fmt.Println("ax", cm.ax, "ay", cm.ay, "bx", cm.bx, "by", cm.by, "px", cm.px, "py", cm.py)
		fmt.Println("l", l, "r", r, "a", a, "b", b)

		if a < 0 || b < 0 {
			fmt.Println("negative")
			fmt.Println()
			continue
		}

		x := a*cm.ax + b*cm.bx
		y := a*cm.ay + b*cm.by

		if cm.px != x || cm.py != y {
			fmt.Println("impossible", "x", x, "y", y)
			fmt.Println()
			continue
		}

		t := a*3 + b

		fmt.Println("Yay!", "t", t)
		fmt.Println()

		tokens += t
	}

	fmt.Println("Problem 2 Result:", tokens, "●", time.Since(start)) // 85644161121698 ● 2.445132ms
}
