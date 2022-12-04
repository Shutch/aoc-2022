package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func part_1(input []byte) (string, error) {
	// split input into lines
	lines := bytes.Split(input, []byte("\n"))

	// iterate through lines and find total calorie count
	// keep maximum calorie count found
	calories := 0
	max_calories := 0
	for _, line := range lines[:len(lines)-1] {
		// if line is blank, reset calories
		if len(line) == 0 {
			if calories > max_calories {
				max_calories = calories
			}
			calories = 0
		} else {
			// convert line to int and add to calories
			cal, err := strconv.Atoi(string(line))
			if err != nil {
				return "", err
			}
			calories += cal
		}
	}
	return strconv.Itoa(max_calories), nil
}

func part_2(input []byte) (string, error) {
	// split input into lines
	lines := bytes.Split(input, []byte("\n"))

	// iterate through lines and find total calorie count
	// of top 3 elves and sum them
	calories := 0
	calories_top_3 := make([]int, 3)
	calories_top_3[0] = 0
	calories_top_3[1] = 0
	calories_top_3[2] = 0
	for _, line := range lines {
		// if line is blank, reset calories
		if len(line) == 0 {
			// find smallest calorie count
			min := 0
			for i := 1; i < 3; i++ {
				if calories_top_3[i] < calories_top_3[min] {
					min = i
				}
			}
			// if calories is greater than smallest calorie count, replace it
			if calories > calories_top_3[min] {
				calories_top_3[min] = calories
			}
			calories = 0
		} else {
			// convert line to int and add to calories
			cal, err := strconv.Atoi(string(line))
			if err != nil {
				return "", err
			}
			calories += cal
		}
	}

	return strconv.Itoa(calories_top_3[0] + calories_top_3[1] + calories_top_3[2]), nil
}

func main() {

	args := os.Args

	if len(args) < 2 {
		fmt.Println("Missing input filename argument.")
		os.Exit(1)
	}

	input, err := os.ReadFile(args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	answer_1, err := part_1(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Part 1: %s\n", answer_1)

	answer_2, err := part_2(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Part 2: %s\n", answer_2)
}
