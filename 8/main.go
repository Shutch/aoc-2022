package main

import (
	"fmt"
	"os"
	"strings"
)

// a struct to represent a point on the map
// the map contains two dimensions (x and y) along with a value
// and a boolean to indicate if the point has been seen before
type Point struct {
	x, y, val int
	seen      bool
}

func part_1(input string) (string, error) {
	// split the input into lines
	lines := strings.Split(input, "\n")

	// get length of grid from number of lines and length of each line
	height := len(lines)
	width := len(lines[0])
	grid := make([][]Point, height)

	// loop through each line, filling in the grid
	for y, line := range lines {
		grid[y] = make([]Point, width)
		for x, char := range line {
			// convert the character to an integer
			val := int(char - '0')
			// create a new point and add it to the grid
			grid[y][x] = Point{x, y, val, false}
		}
	}
	// looking left to right (x == 0 to x == width - 1)
	// finding max value in each row
	// when a max value is found, set Point.seen to true
	total_seen := 0
	for y := 0; y < height; y++ {
		max := -1
		for x := 0; x < width; x++ {
			if grid[y][x].val > max {
				max = grid[y][x].val
				if !grid[y][x].seen {
					grid[y][x].seen = true
					total_seen += 1
				}
			}
		}
	}
	// right to left (x == width - 1 to x == 0)
	for y := 0; y < height; y++ {
		max := -1
		for x := width - 1; x >= 0; x-- {
			if grid[y][x].val > max {
				max = grid[y][x].val
				if !grid[y][x].seen {
					grid[y][x].seen = true
					total_seen += 1
				}
			}
		}
	}
	// top to bottom (y == 0 to y == height - 1)
	for x := 0; x < width; x++ {
		max := -1
		for y := 0; y < height; y++ {
			if grid[y][x].val > max {
				max = grid[y][x].val
				if !grid[y][x].seen {
					grid[y][x].seen = true
					total_seen += 1
				}
			}
		}
	}
	// bottom to top (y == height - 1 to y == 0)
	for x := 0; x < width; x++ {
		max := -1
		for y := height - 1; y >= 0; y-- {
			if grid[y][x].val > max {
				max = grid[y][x].val
				if !grid[y][x].seen {
					grid[y][x].seen = true
					total_seen += 1
				}
			}
		}
	}

	return fmt.Sprint(total_seen), nil
}

func part_2(input string) (string, error) {
	// split the input into lines
	lines := strings.Split(input, "\n")

	// get length of grid from number of lines and length of each line
	height := len(lines)
	width := len(lines[0])
	grid := make([][]Point, height)

	// loop through each line, filling in the grid
	for y, line := range lines {
		grid[y] = make([]Point, width)
		for x, char := range line {
			// convert the character to an integer
			val := int(char - '0')
			// create a new point and add it to the grid
			grid[y][x] = Point{x, y, val, false}
		}
	}
	max_scenic_score := 0
	// iterate through each grid point
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			treehouse_height := grid[y][x].val
			north_trees_seen := 0
			// look up and count all trees that are lower than the treehouse
			for i := y - 1; i >= 0; i-- {
				north_trees_seen++
				if grid[i][x].val >= treehouse_height {
					break
				}
			}
			// look down and count all trees that are lower than the treehouse
			south_trees_seen := 0
			for i := y + 1; i < height; i++ {
				south_trees_seen++
				if grid[i][x].val >= treehouse_height {
					break
				}
			}
			// look left and count all trees that are lower than the treehouse
			west_trees_seen := 0
			for i := x - 1; i >= 0; i-- {
				west_trees_seen++
				if grid[y][i].val >= treehouse_height {
					break
				}
			}
			// look right and count all trees that are lower than the treehouse
			east_trees_seen := 0
			for i := x + 1; i < width; i++ {
				east_trees_seen++
				if grid[y][i].val >= treehouse_height {
					break
				}
			}
			scenic_score := north_trees_seen * south_trees_seen * west_trees_seen * east_trees_seen
			if scenic_score > max_scenic_score {
				max_scenic_score = scenic_score
			}
			// fmt.Printf("x: %d, y: %d, val: %d, north: %d, south: %d, west: %d, east: %d, score: %d\n", x, y, grid[y][x].val, north_trees_seen, south_trees_seen, west_trees_seen, east_trees_seen, scenic_score)
		}
	}

	return fmt.Sprint(max_scenic_score), nil
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
