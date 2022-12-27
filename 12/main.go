package main

import (
	"fmt"
	"os"
	"strings"
)

type point struct {
	x, y int
}

// djikstra's algorithm to find the shortest path between my_pos and goal_pos
func djikstras(graph map[point][]point, start_pos point, end_pos point) (int, error) {
	// initialize distance map
	dist := make(map[point]int)
	for k := range graph {
		dist[k] = 1000000
	}
	dist[start_pos] = 0

	// initialize visited map
	visited := make(map[point]bool)

	// initialize queue
	queue := []point{start_pos}

	for len(queue) > 0 {
		// pop queue
		u := queue[0]
		queue = queue[1:]

		// if we have already visited this node, skip it
		if visited[u] {
			continue
		}

		// mark node as visited
		visited[u] = true

		// iterate through neighbors
		for _, v := range graph[u] {
			// if the distance to this neighbor is shorter than the current distance, update it
			if dist[v] > dist[u]+1 {
				dist[v] = dist[u] + 1
			}
			// add neighbor to queue
			queue = append(queue, v)
		}
	}

	return dist[end_pos], nil
}

func part_1(input string) (string, error) {
	var my_pos point
	var goal_pos point

	// parse input into lines by splitting on newlines
	lines := strings.Split(input, "\n")
	// init grid
	height := len(lines)
	width := len(lines[0])
	grid := make([][]int, height)
	for i := range grid {
		grid[i] = make([]int, width)
	}
	// iterate through each character in each line
	for i, line := range lines {
		for j, char := range line {
			switch char {
			case 'S':
				my_pos = point{x: j, y: i}
				grid[i][j] = 0
			case 'E':
				goal_pos = point{x: j, y: i}
				grid[i][j] = 25
			default:
				// parse char into integer, a = 0, z = 25
				grid[i][j] = int(char) - 97
			}
		}
	}
	// represent the grid into a directed graph where each square is a node
	// each node has directional edges to nodes that are 1 or less larger than the current node
	graph := make(map[point][]point)
	for i, line := range grid {
		for j, node := range line {
			// add node to graph
			graph[point{x: j, y: i}] = make([]point, 0)
			// add edges to graph
			// only add edges to nodes that are 1 or less larger than the current node
			for k := -1; k <= 1; k++ {
				for l := -1; l <= 1; l++ {
					if k == 0 && l == 0 { // skip current node
						continue
					}
					if i+k < 0 || i+k >= height || j+l < 0 || j+l >= width { // skip nodes that are out of bounds
						continue
					}
					if k != 0 && l != 0 { // skip nodes that are diagonal
						continue
					}
					if grid[i+k][j+l] <= node+1 {
						graph[point{x: j, y: i}] = append(graph[point{x: j, y: i}], point{x: j + l, y: i + k})
					}
				}
			}
		}
	}

	// pathfinding from S to E using djikstra's algorithm
	dist, err := djikstras(graph, my_pos, goal_pos)
	if err != nil {
		return "", err
	}
	return fmt.Sprint(dist), nil
}

func part_2(input string) (string, error) {
	var start_pos []point
	var goal_pos point

	// parse input into lines by splitting on newlines
	lines := strings.Split(input, "\n")
	// init grid
	height := len(lines)
	width := len(lines[0])
	grid := make([][]int, height)
	for i := range grid {
		grid[i] = make([]int, width)
	}
	// iterate through each character in each line
	for i, line := range lines {
		for j, char := range line {
			switch char {
			case 'S':
				start_pos = append(start_pos, point{x: j, y: i})
				grid[i][j] = 0
			case 'E':
				goal_pos = point{x: j, y: i}
				grid[i][j] = 25
			default:
				// parse char into integer, a = 0, z = 25
				grid[i][j] = int(char) - 97
				if grid[i][j] == 0 {
					start_pos = append(start_pos, point{x: j, y: i})
				}
			}
		}
	}
	// represent the grid into a directed graph where each square is a node
	// each node has directional edges to nodes that are 1 or less larger than the current node
	graph := make(map[point][]point)
	for i, line := range grid {
		for j, node := range line {
			// add node to graph
			graph[point{x: j, y: i}] = make([]point, 0)
			// add edges to graph
			// only add edges to nodes that are 1 or less larger than the current node
			for k := -1; k <= 1; k++ {
				for l := -1; l <= 1; l++ {
					if k == 0 && l == 0 { // skip current node
						continue
					}
					if i+k < 0 || i+k >= height || j+l < 0 || j+l >= width { // skip nodes that are out of bounds
						continue
					}
					if k != 0 && l != 0 { // skip nodes that are diagonal
						continue
					}
					if grid[i+k][j+l] <= node+1 {
						graph[point{x: j, y: i}] = append(graph[point{x: j, y: i}], point{x: j + l, y: i + k})
					}
				}
			}
		}
	}

	// pathfinding from S to E using djikstra's algorithm
	min_dist := 100000000
	for _, pos := range start_pos {
		dist, err := djikstras(graph, pos, goal_pos)
		if err != nil {
			return "", err
		}
		if dist < min_dist {
			min_dist = dist
		}
	}
	return fmt.Sprint(min_dist), nil
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
