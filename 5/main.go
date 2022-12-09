package main

import (
	"fmt"
	"os"
    "strings"
    "strconv"
)

type Stack struct {
    stack  string
    height int
}

func part_1(input string) (string, error) {
    // split lines by newline
    lines := strings.Split(input, "\n")

    // scan for bottom of crates (blank line)
    crate_end_line_index := 0
    for i, line := range lines {
        if line == "" {
            crate_end_line_index = i - 2  // blank line and column label
            break
        } 
    }

    // build stacks from bottom to top
    // number of stacks can be determined by length of column label row
    // 1 == 2, 2 == 6, 3 == 10
    num_stacks := (len(lines[crate_end_line_index + 1]) + 2) / 4

    // index of each stack (in the string) is known
    // 1 == 1, 2 == 5, 3 == 9
    stacks := make([]string, num_stacks)
    for i := crate_end_line_index; i >= 0; i-- {
        for j := 0; j < num_stacks; j++ {
            column_index := ((j + 1) * 4) - 3
            char := string(lines[i][column_index])
            if char != " " {
                stacks[j] = stacks[j] + char
            }
        }
    }

    // execute moves
    for i := crate_end_line_index + 3; i < len(lines) - 1; i++ {
        line := lines[i]
        fields := strings.Fields(line)
        qty, err := strconv.Atoi(fields[1])
        if err != nil {
            return "", err
        }
        start_stack, err := strconv.Atoi(fields[3])
        if err != nil {
            return "", err
        }
        start_stack--
        end_stack, err := strconv.Atoi(fields[5])
        if err != nil {
            return "", err
        }
        end_stack--

        // move onto stack, remove from previous stack
        for j := 0; j < qty; j++ {
            stacks[end_stack] = stacks[end_stack] + stacks[start_stack][len(stacks[start_stack]) - 1:]
            stacks[start_stack] = stacks[start_stack][:len(stacks[start_stack]) - 1]
        }
    }

    ans := ""
    for i := 0; i < num_stacks; i++ {
        ans = ans + stacks[i][len(stacks[i]) - 1:]
    }

	return ans, nil
}

func part_2(input string) (string, error) {
    // split lines by newline
    lines := strings.Split(input, "\n")

    // scan for bottom of crates (blank line)
    crate_end_line_index := 0
    for i, line := range lines {
        if line == "" {
            crate_end_line_index = i - 2  // blank line and column label
            break
        } 
    }

    // build stacks from bottom to top
    // number of stacks can be determined by length of column label row
    // 1 == 2, 2 == 6, 3 == 10
    num_stacks := (len(lines[crate_end_line_index + 1]) + 2) / 4

    // index of each stack (in the string) is known
    // 1 == 1, 2 == 5, 3 == 9
    stacks := make([]string, num_stacks)
    for i := crate_end_line_index; i >= 0; i-- {
        for j := 0; j < num_stacks; j++ {
            column_index := ((j + 1) * 4) - 3
            char := string(lines[i][column_index])
            if char != " " {
                stacks[j] = stacks[j] + char
            }
        }
    }

    // execute moves
    for i := crate_end_line_index + 3; i < len(lines) - 1; i++ {
        line := lines[i]
        fields := strings.Fields(line)
        qty, err := strconv.Atoi(fields[1])
        if err != nil {
            return "", err
        }
        start_stack, err := strconv.Atoi(fields[3])
        if err != nil {
            return "", err
        }
        start_stack--
        end_stack, err := strconv.Atoi(fields[5])
        if err != nil {
            return "", err
        }
        end_stack--

        // move onto stack, remove from previous stack
        stacks[end_stack] = stacks[end_stack] + stacks[start_stack][len(stacks[start_stack]) - qty:]
        stacks[start_stack] = stacks[start_stack][:len(stacks[start_stack]) - qty]
    }

    ans := ""
    for i := 0; i < num_stacks; i++ {
        ans = ans + stacks[i][len(stacks[i]) - 1:]
    }

	return ans, nil
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
