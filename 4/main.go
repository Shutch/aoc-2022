package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part_1(input string) (string, error) {
	// split lines by newline
	lines := strings.Split(input, "\n")

	total := 0
	// iterate through lines to compare pairs
	for _, line := range lines {
		// split string along comma and then hyphens
		elves := strings.Split(line, ",")
		elf_1 := strings.Split(elves[0], "-")
		elf_2 := strings.Split(elves[1], "-")
		elf_1_lower, err := strconv.Atoi(elf_1[0])
		if err != nil {
			return "", err
		}
		elf_1_upper, err := strconv.Atoi(elf_1[1])
		if err != nil {
			return "", err
		}
		elf_2_lower, err := strconv.Atoi(elf_2[0])
		if err != nil {
			return "", err
		}
		elf_2_upper, err := strconv.Atoi(elf_2[1])
		if err != nil {
			return "", err
		}

		// check if elf_1 values fall between elf_2 values and vice versa
		if (elf_1_lower >= elf_2_lower && elf_1_upper <= elf_2_upper) || (elf_1_lower <= elf_2_lower && elf_1_upper >= elf_2_upper) {
			total++
		}

	}

	return fmt.Sprint(total), nil
}

func part_2(input string) (string, error) {
	// split lines by newline
	lines := strings.Split(input, "\n")

	total := 0
	// iterate through lines to compare pairs
	for _, line := range lines {
		// split string along comma and then hyphens
		elves := strings.Split(line, ",")
		elf_1 := strings.Split(elves[0], "-")
		elf_2 := strings.Split(elves[1], "-")
		elf_1_lower, err := strconv.Atoi(elf_1[0])
		if err != nil {
			return "", err
		}
		elf_1_upper, err := strconv.Atoi(elf_1[1])
		if err != nil {
			return "", err
		}
		elf_2_lower, err := strconv.Atoi(elf_2[0])
		if err != nil {
			return "", err
		}
		elf_2_upper, err := strconv.Atoi(elf_2[1])
		if err != nil {
			return "", err
		}

		// check if any value overlaps another
		if (elf_1_upper >= elf_2_lower && elf_1_lower <= elf_2_upper) || (elf_2_upper >= elf_1_lower && elf_2_lower <= elf_1_upper) {
			total++
		}

	}

	return fmt.Sprint(total), nil
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
