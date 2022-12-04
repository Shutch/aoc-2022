package main

import (
	"fmt"
	"os"
	"strings"
)

func part_1(input string) (string, error) {
	// split lines by newline
	lines := strings.Split(input, "\n")

	priority := 0
	for _, line := range lines {
		// first compartment is first half of line
		first_compartment := line[:len(line)/2]
		second_compartment := line[len(line)/2:]

		// find the common character between the two compartments
	out:
		for i := 0; i < len(first_compartment); i++ {
			for j := 0; j < len(second_compartment); j++ {
				if first_compartment[i] == second_compartment[j] {
					// priority a = 1, z = 26, A = 27, Z = 52
					letter := first_compartment[i]
					if letter >= 'a' && letter <= 'z' {
						p := int(letter - 'a' + 1)
						priority += p
					} else {
						p := int(letter - 'A' + 27)
						priority += p
					}
					break out
				}
			}
		}
	}

	return fmt.Sprint(priority), nil
}

func part_2(input string) (string, error) {
	// split lines by newline
	lines := strings.Split(input, "\n")

	priority := 0
	// parse lines in groups of 3
	for i := 0; i < len(lines); i += 3 {
	out:
		// iterate through all three lists to find common character
		for j := 0; j < len(lines[i]); j++ {
			for k := 0; k < len(lines[i+1]); k++ {
				for l := 0; l < len(lines[i+2]); l++ {
					if lines[i][j] == lines[i+1][k] && lines[i][j] == lines[i+2][l] {
						// priority a = 1, z = 26, A = 27, Z = 52
						letter := lines[i][j]
						if letter >= 'a' && letter <= 'z' {
							p := int(letter - 'a' + 1)
							priority += p
						} else {
							p := int(letter - 'A' + 27)
							priority += p
						}
						break out
					}
				}
			}
		}
	}

	return fmt.Sprint(priority), nil
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
