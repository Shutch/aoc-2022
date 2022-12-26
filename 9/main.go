package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

func part_1(input string) (string, error) {
	// parse input into list of lines
	lines := strings.Split(input, "\n")

	// points is a map of point structs to bools
	// when the tail touches a grid point it's added to the map
	points := make(map[point]bool)

	// starting at 0
	current_head_point := point{0, 0}
	current_tail_point := point{0, 0}

	points[current_tail_point] = true

	// for each line, move the head
	for _, line := range lines {
		//fmt.Println(line)
		// parse line into direction and distance
		direction := line[0]
		distance, err := strconv.Atoi(line[2:])
		if err != nil {
			return "", err
		}
		switch direction {
		case 'R':
			for i := 0; i < distance; i++ {
				current_head_point.x += 1
				// check that the head doesn't move more than +1 x away from the tail
				if current_head_point.x-current_tail_point.x > 1 {
					current_tail_point.x = current_head_point.x - 1
					current_tail_point.y = current_head_point.y
				}
				points[current_tail_point] = true
				//fmt.Println(current_head_point, current_tail_point, len(points))
			}
		case 'L':
			for i := 0; i < distance; i++ {
				current_head_point.x -= 1
				// check that the head doesn't move more than -1 x away from the tail
				if current_head_point.x-current_tail_point.x < -1 {
					current_tail_point.x = current_head_point.x + 1
					current_tail_point.y = current_head_point.y
				}
				points[current_tail_point] = true
				//fmt.Println(current_head_point, current_tail_point, len(points))
			}
		case 'U':
			for i := 0; i < distance; i++ {
				current_head_point.y += 1
				// check that the head doesn't move more than +1 y away from the tail
				if current_head_point.y-current_tail_point.y > 1 {
					current_tail_point.y = current_head_point.y - 1
					current_tail_point.x = current_head_point.x
				}
				points[current_tail_point] = true
				//fmt.Println(current_head_point, current_tail_point, len(points))
			}
		case 'D':
			for i := 0; i < distance; i++ {
				current_head_point.y -= 1
				// check that the head doesn't move more than -1 y away from the tail
				if current_head_point.y-current_tail_point.y < -1 {
					current_tail_point.y = current_head_point.y + 1
					current_tail_point.x = current_head_point.x
				}
				points[current_tail_point] = true
				//fmt.Println(current_head_point, current_tail_point, len(points))
			}
		}
	}
	return fmt.Sprint(len(points)), nil
}

// prints the points on an ascii grid
func printPoints(current_rope_points []point, points map[point]bool, width, height int) {
	//pregenerate grid before printing
	grid := make([][]string, height*2+1)
	for i := range grid {
		grid[i] = make([]string, width*2+1)
	}

	for y := height; y >= -height; y-- {
		for x := -width; x <= width; x++ {
			// default is a dot
			grid[y+height][x+width] = "."

			// check if x, y coordinates are in points, if so print #
			_, ok := points[point{x, y}]
			if ok {
				grid[y+height][x+width] = "#"
			}
		}
	}
	// assigning current rope points starting at the last one
	for i := len(current_rope_points) - 1; i >= 0; i-- {
		grid[current_rope_points[i].y+height][current_rope_points[i].x+width] = strconv.Itoa(i)
	}

	// printing the grid
	for y := height; y >= -height; y-- {
		for x := -width; x <= width; x++ {
			fmt.Print(grid[y+height][x+width])
		}
		fmt.Println()
	}
}

func part_2(input string) (string, error) {
	// parse input into list of lines
	lines := strings.Split(input, "\n")

	// points is a map of point structs to bools
	// when the tail touches a grid point it's added to the map
	points := make(map[point]bool)

	// ten points all starting at 0, 0
	current_rope_points := []point{
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
	}

	points[current_rope_points[0]] = true

	// for each line, move the head
	for _, line := range lines {
		//fmt.Println(line)
		// parse line into direction and distance
		direction := line[0]
		distance, err := strconv.Atoi(line[2:])
		if err != nil {
			return "", err
		}
		switch direction {
		case 'R':
			for i := 0; i < distance; i++ {
				current_rope_points[0].x += 1
				// check that each rope point doesn't move more than +1 x away from the previous rope point
				for j := 1; j < 10; j++ {
					head := current_rope_points[j-1]
					tail := current_rope_points[j]

					// if the head and tail are more than 1 away from each other
					// and in different column and row, move the tail diagonally towards the head
					if head.x-1 > tail.x && head.y-1 > tail.y {
						tail.x = head.x - 1
						tail.y = head.y - 1
					}
					if head.x+1 < tail.x && head.y-1 > tail.y {
						tail.x = head.x + 1
						tail.y = head.y - 1
					}
					if head.x-1 > tail.x && head.y+1 < tail.y {
						tail.x = head.x - 1
						tail.y = head.y + 1
					}
					if head.x+1 < tail.x && head.y+1 < tail.y {
						tail.x = head.x + 1
						tail.y = head.y + 1
					}

					// if the tail is more than 1 away from the head, move it one step towards the head
					if head.x-1 > tail.x {
						tail.x = head.x - 1
						tail.y = head.y
					}
					if head.x+1 < tail.x {
						tail.x = head.x + 1
						tail.y = head.y
					}
					if head.y-1 > tail.y {
						tail.x = head.x
						tail.y = head.y - 1
					}
					if head.y+1 < tail.y {
						tail.x = head.x
						tail.y = head.y + 1
					}
					current_rope_points[j-1] = head
					current_rope_points[j] = tail
				}
				//fmt.Println(current_rope_points)
				//printPoints(current_rope_points, points, 15, 15)
				//println()
				points[current_rope_points[9]] = true
			}
		case 'L':
			for i := 0; i < distance; i++ {
				current_rope_points[0].x -= 1
				// check that each rope point doesn't move more than -1 x away from the previous rope point
				for j := 1; j < 10; j++ {
					head := current_rope_points[j-1]
					tail := current_rope_points[j]

					// if the head and tail are more than 1 away from each other
					// and in different column and row, move the tail diagonally towards the head
					if head.x-1 > tail.x && head.y-1 > tail.y {
						tail.x = head.x - 1
						tail.y = head.y - 1
					}
					if head.x+1 < tail.x && head.y-1 > tail.y {
						tail.x = head.x + 1
						tail.y = head.y - 1
					}
					if head.x-1 > tail.x && head.y+1 < tail.y {
						tail.x = head.x - 1
						tail.y = head.y + 1
					}
					if head.x+1 < tail.x && head.y+1 < tail.y {
						tail.x = head.x + 1
						tail.y = head.y + 1
					}

					// if the tail is more than 1 away from the head, move it one step towards the head
					if head.x-1 > tail.x {
						tail.x = head.x - 1
						tail.y = head.y
					}
					if head.x+1 < tail.x {
						tail.x = head.x + 1
						tail.y = head.y
					}
					if head.y-1 > tail.y {
						tail.x = head.x
						tail.y = head.y - 1
					}
					if head.y+1 < tail.y {
						tail.x = head.x
						tail.y = head.y + 1
					}
					current_rope_points[j-1] = head
					current_rope_points[j] = tail
				}
				//fmt.Println(current_rope_points)
				//printPoints(current_rope_points, points, 15, 15)
				//println()
				points[current_rope_points[9]] = true
			}
		case 'U':
			for i := 0; i < distance; i++ {
				current_rope_points[0].y += 1
				// check that each rope point doesn't move more than +1 y away from the previous rope point
				for j := 1; j < 10; j++ {
					head := current_rope_points[j-1]
					tail := current_rope_points[j]

					// if the head and tail are more than 1 away from each other
					// and in different column and row, move the tail diagonally towards the head
					if head.x-1 > tail.x && head.y-1 > tail.y {
						tail.x = head.x - 1
						tail.y = head.y - 1
					}
					if head.x+1 < tail.x && head.y-1 > tail.y {
						tail.x = head.x + 1
						tail.y = head.y - 1
					}
					if head.x-1 > tail.x && head.y+1 < tail.y {
						tail.x = head.x - 1
						tail.y = head.y + 1
					}
					if head.x+1 < tail.x && head.y+1 < tail.y {
						tail.x = head.x + 1
						tail.y = head.y + 1
					}

					// if the tail is more than 1 away from the head, move it one step towards the head
					if head.x-1 > tail.x {
						tail.x = head.x - 1
						tail.y = head.y
					}
					if head.x+1 < tail.x {
						tail.x = head.x + 1
						tail.y = head.y
					}
					if head.y-1 > tail.y {
						tail.x = head.x
						tail.y = head.y - 1
					}
					if head.y+1 < tail.y {
						tail.x = head.x
						tail.y = head.y + 1
					}
					current_rope_points[j-1] = head
					current_rope_points[j] = tail
				}
				//fmt.Println(current_rope_points)
				//printPoints(current_rope_points, points, 15, 15)
				//println()
				points[current_rope_points[9]] = true
			}
		case 'D':
			for i := 0; i < distance; i++ {
				current_rope_points[0].y -= 1
				// check that each rope point doesn't move more than -1 y away from the previous rope point
				for j := 1; j < 10; j++ {
					head := current_rope_points[j-1]
					tail := current_rope_points[j]

					// if the head and tail are more than 1 away from each other
					// and in different column and row, move the tail diagonally towards the head
					if head.x-1 > tail.x && head.y-1 > tail.y {
						tail.x = head.x - 1
						tail.y = head.y - 1
					}
					if head.x+1 < tail.x && head.y-1 > tail.y {
						tail.x = head.x + 1
						tail.y = head.y - 1
					}
					if head.x-1 > tail.x && head.y+1 < tail.y {
						tail.x = head.x - 1
						tail.y = head.y + 1
					}
					if head.x+1 < tail.x && head.y+1 < tail.y {
						tail.x = head.x + 1
						tail.y = head.y + 1
					}

					// if the tail is more than 1 away from the head, move it one step towards the head
					if head.x-1 > tail.x {
						tail.x = head.x - 1
						tail.y = head.y
					}
					if head.x+1 < tail.x {
						tail.x = head.x + 1
						tail.y = head.y
					}
					if head.y-1 > tail.y {
						tail.x = head.x
						tail.y = head.y - 1
					}
					if head.y+1 < tail.y {
						tail.x = head.x
						tail.y = head.y + 1
					}
					current_rope_points[j-1] = head
					current_rope_points[j] = tail
				}
				//fmt.Println(current_rope_points)
				//printPoints(current_rope_points, points, 15, 15)
				//println()
				points[current_rope_points[9]] = true
			}
		}
		//fmt.Println(len(points))
	}
	return fmt.Sprint(len(points)), nil
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
