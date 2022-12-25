package main

import (
	"fmt"
	"os"
)

func part_1(input string) (string, error) {
	// the input is a single line
	// parse input as a sliding window of 4 characters
	for i := 0; i < len(input)-3; i++ {
		// check if the 4 characters are unique
		unique := true
		for j := 0; j < 4; j++ {
		out:
			for k := j + 1; k < 4; k++ {
				if input[i+j] == input[i+k] {
					unique = false
					break out
				}
			}
		}
		// if the 4 characters are unique, return the index of
		// the last character
		if unique {
			return fmt.Sprint(i + 4), nil
		}
	}
	// if no unique 4 characters are found, raise error
	return "", fmt.Errorf("no unique 4 characters found")
}

func part_2(input string) (string, error) {
	// the input is a single line
	// parse input as a sliding window of 14 characters
	for i := 0; i < len(input)-13; i++ {
		// check if the 4 characters are unique
		unique := true
		for j := 0; j < 14; j++ {
		out:
			for k := j + 1; k < 14; k++ {
				if input[i+j] == input[i+k] {
					unique = false
					break out
				}
			}
		}
		// if the 14 characters are unique, return the index of
		// the last character
		if unique {
			return fmt.Sprint(i + 14), nil
		}
	}
	// if no unique 4 characters are found, raise error
	return "", fmt.Errorf("no unique 4 characters found")
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
