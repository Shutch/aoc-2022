package main

import (
	"fmt"
	"os"
	"strings"
)

func part_1(input string) (string, error) {
	//split input into lines by newline
	lines := strings.Split(input, "\n")

	// rock paper scissors win map
	// key is opponent's move, value is my move
	win_map := make(map[int]int)
	win_map[0] = 1
	win_map[1] = 2
	win_map[2] = 0

	// ascii A=65, X=88
	// parse lines
	score := 0
	for _, line := range lines[:len(lines)-1] {
		opp_shape := int(line[0]) - 65
		my_shape := int(line[2]) - 88
		// check for tie
		if opp_shape == my_shape {
			score += 3
		}
		// check for win
		if win_map[opp_shape] == my_shape {
			score += 6
		}
		// add played shape value (1 for R, 2 for P, 3 for S)
		score += my_shape + 1
	}

	return fmt.Sprint(score), nil

}

func part_2(input string) (string, error) {
	// split input into lines by newline
	lines := strings.Split(input, "\n")

	// rock paper scissors win map
	// key is opponent's move, value is my move
	win_shape := func(opp_shape int) int { return (opp_shape + 1) % 3 }

	lose_shape := func(opp_shape int) int { return (opp_shape + 2) % 3 }

	score := 0
	for _, line := range lines[:len(lines)-1] {
		opp_shape := int(line[0]) - 65
		outcome := int(line[2]) - 88
		if outcome == 0 {
			score += lose_shape(opp_shape) + 1
		} else if outcome == 1 {
			score += opp_shape + 1
			score += 3
		} else {
			score += win_shape(opp_shape) + 1
			score += 6
		}
	}

	return fmt.Sprint(score), nil
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
