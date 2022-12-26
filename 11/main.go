package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type monkey struct {
	items           []int
	op              string
	opVal           int
	test            int
	ifTrue          int
	ifFalse         int
	inspectionCount int
}

func part_1(input string) (string, error) {
	// parse input into lines
	lines := strings.Split(input, "\n")
	// parse lines into monkeys
	monkeys := make([]monkey, 0)
	for i := 0; i < len(lines); i++ {
		if strings.HasPrefix(lines[i], "Monkey") {
			// parse monkey
			monkey := monkey{}
			// parse starting items
			i++
			monkey.items = make([]int, 0)
			for _, item := range strings.Split(lines[i][18:], ", ") {
				val, err := strconv.Atoi(item)
				if err != nil {
					return "", err
				}
				monkey.items = append(monkey.items, val)
			}
			// parse operation
			i++
			monkey.op = string(lines[i][23])
			// parse operation value
			if lines[i][25:] == "old" {
				monkey.opVal = -1
			} else {
				val, err := strconv.Atoi(lines[i][25:])
				if err != nil {
					return "", err
				}
				monkey.opVal = val
			}
			// parse test
			i++
			val, err := strconv.Atoi(lines[i][21:])
			if err != nil {
				return "", err
			}
			monkey.test = val
			// parse if true
			i++
			val, err = strconv.Atoi(lines[i][29:])
			if err != nil {
				return "", err
			}
			monkey.ifTrue = val
			// parse if false
			i++
			val, err = strconv.Atoi(lines[i][30:])
			if err != nil {
				return "", err
			}
			monkey.ifFalse = val

			monkeys = append(monkeys, monkey)
		}
	}

	inspections := make([]int, len(monkeys))
	// run 20 rounds total
	for round := 1; round <= 20; round++ {
		// single round iterating through each monkey
		for i := 0; i < len(monkeys); i++ {
			// iterate through each item
			for j := 0; j < len(monkeys[i].items); j++ {
				// inspection
				if monkeys[i].op == "+" {
					if monkeys[i].opVal == -1 {
						monkeys[i].items[j] += monkeys[i].items[j]
					} else {
						monkeys[i].items[j] += monkeys[i].opVal
					}
				} else {
					if monkeys[i].opVal == -1 {
						monkeys[i].items[j] *= monkeys[i].items[j]
					} else {
						monkeys[i].items[j] *= monkeys[i].opVal
					}
				}
				inspections[i]++

				// worry level divided by 3 and rounded down to nearest integer
				monkeys[i].items[j] = monkeys[i].items[j] / 3

				// divisibility test
				var monkey_index int
				if monkeys[i].items[j]%monkeys[i].test == 0 {
					monkey_index = monkeys[i].ifTrue
				} else {
					monkey_index = monkeys[i].ifFalse
				}
				// throw to next monkey
				monkeys[monkey_index].items = append(monkeys[monkey_index].items, monkeys[i].items[j])
			}
			// clear items
			monkeys[i].items = make([]int, 0)
		}
	}

	// sort inspections highest to lowest
	for i := 0; i < len(inspections); i++ {
		for j := i + 1; j < len(inspections); j++ {
			if inspections[i] < inspections[j] {
				inspections[i], inspections[j] = inspections[j], inspections[i]
			}
		}
	}

	monkey_business := inspections[0] * inspections[1]

	return fmt.Sprint(monkey_business), nil
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func part_2(input string) (string, error) {
	tests := make([]int, 0)
	// parse input into lines
	lines := strings.Split(input, "\n")
	// parse lines into monkeys
	monkeys := make([]monkey, 0)
	for i := 0; i < len(lines); i++ {
		if strings.HasPrefix(lines[i], "Monkey") {
			// parse monkey
			monkey := monkey{}
			// parse starting items
			i++
			monkey.items = make([]int, 0)
			for _, item := range strings.Split(lines[i][18:], ", ") {
				val, err := strconv.Atoi(item)
				if err != nil {
					return "", err
				}
				monkey.items = append(monkey.items, val)
			}
			// parse operation
			i++
			monkey.op = string(lines[i][23])
			// parse operation value
			if lines[i][25:] == "old" {
				monkey.opVal = -1
			} else {
				val, err := strconv.Atoi(lines[i][25:])
				if err != nil {
					return "", err
				}
				monkey.opVal = val
			}
			// parse test
			i++
			val, err := strconv.Atoi(lines[i][21:])
			if err != nil {
				return "", err
			}
			monkey.test = val
			tests = append(tests, val)
			// parse if true
			i++
			val, err = strconv.Atoi(lines[i][29:])
			if err != nil {
				return "", err
			}
			monkey.ifTrue = val
			// parse if false
			i++
			val, err = strconv.Atoi(lines[i][30:])
			if err != nil {
				return "", err
			}
			monkey.ifFalse = val

			monkeys = append(monkeys, monkey)
		}
	}

	// find the supermodulo by multiplying all tests together
	// cheat found here https://www.reddit.com/r/adventofcode/comments/zih7gf/2022_day_11_part_2_what_does_it_mean_find_another/izr79go/?context=3
	supermodulo := 1
	for i := 0; i < len(tests); i++ {
		supermodulo *= tests[i]
	}

	inspections := make([]int, len(monkeys))
	// run 20 rounds total
	for round := 1; round <= 10000; round++ {
		// single round iterating through each monkey
		for i := 0; i < len(monkeys); i++ {
			// iterate through each item
			for j := 0; j < len(monkeys[i].items); j++ {
				// inspection
				if monkeys[i].op == "+" {
					if monkeys[i].opVal == -1 {
						monkeys[i].items[j] += monkeys[i].items[j]
					} else {
						monkeys[i].items[j] += monkeys[i].opVal
					}
				} else {
					if monkeys[i].opVal == -1 {
						monkeys[i].items[j] *= monkeys[i].items[j]
					} else {
						monkeys[i].items[j] *= monkeys[i].opVal
					}
				}
				inspections[i]++

				// worry level divided by 3 and rounded down to nearest integer
				// monkeys[i].items[j] = monkeys[i].items[j] / 3

				// worry level mod supermodulo
				monkeys[i].items[j] = monkeys[i].items[j] % supermodulo

				// divisibility test
				var monkey_index int
				if monkeys[i].items[j]%monkeys[i].test == 0 {
					monkey_index = monkeys[i].ifTrue
				} else {
					monkey_index = monkeys[i].ifFalse
				}
				// throw to next monkey
				monkeys[monkey_index].items = append(monkeys[monkey_index].items, monkeys[i].items[j])
			}
			// clear items
			monkeys[i].items = make([]int, 0)
		}
	}

	// sort inspections highest to lowest
	for i := 0; i < len(inspections); i++ {
		for j := i + 1; j < len(inspections); j++ {
			if inspections[i] < inspections[j] {
				inspections[i], inspections[j] = inspections[j], inspections[i]
			}
		}
	}

	monkey_business := inspections[0] * inspections[1]

	return fmt.Sprint(monkey_business), nil
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
