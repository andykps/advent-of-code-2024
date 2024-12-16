package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
)

type vec struct {
	dx int
	dy int
}
type point struct {
	x int
	y int
}

const EMPTY byte = 46 // .
const WALL byte = 35  // #
const START byte = 83 // S
const END byte = 69   // E

var NORTH = vec{0, -1}
var EAST = vec{1, 0}
var SOUTH = vec{0, 1}
var WEST = vec{-1, 0}
var dirs = []vec{NORTH, EAST, SOUTH, WEST}

var grid = [][]byte{}
var nodes = make(map[point]int)
var start point
var end point

func main() {
	debug := flag.Bool("debug", false, "Output extra debug info")
	flag.Parse()
	input := "input.txt"
	if len(flag.Args()) > 0 {
		input = flag.Args()[0]
	}

	readGridFromFile(input)
	if *debug {
		printGrid()
		fmt.Println("Start:", start, "End:", end)
		fmt.Println()
	}

	findNodes()
	if *debug {
		printGridWithNodes()
		fmt.Println()
	}

	nodes[start] = 0
	// walk nodes to find costs between
	queue := []point{start}
	for i := 0; i < len(queue); i++ {
		node := queue[i]
		adj := findAdjacent(node)
		for _, other := range adj {
			if slices.Contains(queue, other) {
				// if it's before us in the queue then it's already processed
				continue
			}
			dx := node.x - other.x
			dy := node.y - other.y
			if dx < 0 {
				dx = -dx
			}
			if dy < 0 {
				dy = -dy
			}
			cost := dx + dy
			nodes[other] = cost + nodes[node]
			queue = append(queue, other)
		}
	}

	if *debug {
		fmt.Println(nodes)
	}

	// walk backwards from end to start to count paths (turns)
	count := 0
	queue = []point{end}
	for i := 0; i < len(queue); i++ {
		node := queue[i]
		min := math.MaxInt
		var best point
		for _, other := range findAdjacent(node) {
			if nodes[other] < min {
				min = nodes[other]
				best = other
			}
		}
		if best == start {
			firstDir := slices.Index(dirs, vec{unit(node.x - start.x), unit(node.y - start.y)})
			startDir := slices.Index(dirs, EAST)
			if startDir == firstDir {
				count += 0 // no need to turn
			} else if startDir == (firstDir + 1)%4 || startDir == (firstDir - 1)%4 {
				count += 1 // turn 90 degress
			} else {
				count += 2 // turn 180 degrees
			}

			break
		}
		count += 1
		queue = append(queue, best)
	}

	fmt.Println(count, nodes[end], count*1000 + nodes[end])
}

func unit(i int) int {
	if i < 0 {
		return -1
	} else if i > 0 {
		return 1
	} else {
		return 0
	}
}

func findAdjacent(p point) (adj []point) {
	for node := range nodes {
		dx := node.x - p.x
		dy := node.y - p.y
		if (dx == 0 || dy == 0) && noWallsBetween(p, node) {
			adj = append(adj, node)
		}
	}
	return adj
}

func noWallsBetween(p1 point, p2 point) bool {
	for y := min(p1.y, p2.y); y <= max(p1.y, p2.y); y++ {
		for x := min(p1.x, p2.x); x <= max(p1.x, p2.x); x++ {
			if grid[y][x] == WALL {
				return false
			}
		}
	}
	return true
}

func findNodes() {
	type todo struct {
		pos point
		dir vec
	}
	queue := []todo{}

	for _, dir := range dirs {
		if grid[start.y + dir.dy][start.x + dir.dx] == EMPTY {
			queue = append(queue, todo{start, dir})
		}
	}

	// find all the nodes in the map
	for i := 0; i < len(queue); i++ {
		exits := []todo{}
		for _, dir := range dirs {
			x := queue[i].pos.x + dir.dx
			y := queue[i].pos.y + dir.dy
			cell := grid[y][x]
			if cell == EMPTY || cell == END {
				exits = append(exits, todo{point{x, y}, dir})
			}
		}
		for _, exit := range exits {
			if !slices.Contains(queue, exit) {
				queue = append(queue, exit)
			}
		}
		if len(exits) > 2 || grid[queue[i].pos.y][queue[i].pos.x] == END || len(exits) == 2 && !(exits[0].dir.dy == 0 && exits[1].dir.dy == 0 || exits[0].dir.dx == 0 && exits[1].dir.dx == 0) {
			// this must be a junction
			nodes[queue[i].pos] = math.MaxInt
		}
	}
}

func readGridFromFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	y := 0
	for scanner.Scan() {
		line := make([]byte, len(scanner.Bytes()))
		copy(line, scanner.Bytes())
		grid = append(grid, line)
		if x := slices.Index(line, START); x > -1 {
			start = point{x, y}
		}
		if x := slices.Index(line, END); x > -1 {
			end = point{x, y}
		}
		y += 1
	}
}

func printGrid() {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			fmt.Print(string(grid[y][x]))
		}
		fmt.Println()
	}
}

func printGridWithNodes() {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if _, ok := nodes[point{x, y}]; ok {
				fmt.Print("X")
			} else {
				fmt.Print(string(grid[y][x]))
			}
		}
		fmt.Println()
	}
}