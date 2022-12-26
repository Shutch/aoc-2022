package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part_1(input string) (string, error) {
	signal_strength_cycles := []int{20, 60, 100, 140, 180, 220}
	X := 1
	cycle_num := 0
	signal_strengths := []int{0, 0, 0, 0, 0, 0}

	// parse input into list
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		// parse instruction
		fields := strings.Fields(line)
		if len(fields) == 1 {
			cycle_num += 1
			// check if cycle is in signaler strength cycle
			for i, cycle := range signal_strength_cycles {
				if cycle_num == cycle {
					signal_strengths[i] = X * cycle_num
				}
			}
		} else {
			cycle_num += 1
			// check if cycle is in signaler strength cycle
			for i, cycle := range signal_strength_cycles {
				if cycle_num == cycle {
					signal_strengths[i] = X * cycle_num
				}
			}

			cycle_num += 1
			// check if cycle is in signaler strength cycle
			for i, cycle := range signal_strength_cycles {
				if cycle_num == cycle {
					signal_strengths[i] = X * cycle_num
				}
			}

			// parse number in second field
			num, err := strconv.Atoi(fields[1])
			if err != nil {
				return "", err
			}
			X += num
		}
	}
	// sum signal strengths
	sum := 0
	for _, strength := range signal_strengths {
		sum += strength
	}
	return fmt.Sprint(sum), nil
}

func part_2(input string) (string, error) {
	X := 1
	cycle_num := 1

	// parse input into list
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		// parse instruction
		fields := strings.Fields(line)
		if len(fields) == 1 {
			// draw pixel if cycle_num is within 1 of X
			if (cycle_num-1)%40-X <= 1 && (cycle_num-1)%40-X >= -1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
			if cycle_num%40 == 0 {
				fmt.Println()
			}
			cycle_num += 1

		} else {
			// draw pixel if cycle_num is within 1 of X
			if (cycle_num-1)%40-X <= 1 && (cycle_num-1)%40-X >= -1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
			if cycle_num%40 == 0 {
				fmt.Println()
			}
			cycle_num += 1

			// draw pixel if cycle_num is within 1 of X
			if (cycle_num-1)%40-X <= 1 && (cycle_num-1)%40-X >= -1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
			if cycle_num%40 == 0 {
				fmt.Println()
			}
			cycle_num += 1

			// parse number in second field
			num, err := strconv.Atoi(fields[1])
			if err != nil {
				return "", err
			}
			X += num
		}
	}

	return "", nil
}

func main() {
	// get input
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input_str := string(input)

	// part 1
	ans, err := part_1(input_str)
	if err != nil {
		panic(err)
	}
	fmt.Println(ans)

	// part 2
	ans, err = part_2(input_str)
	if err != nil {
		panic(err)
	}
	fmt.Println(ans)
}
