package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
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
type node struct {
	pos    point
	cost   int
	length int
	prev   *node
}

const EMPTY byte = 46 // .
const WALL byte = 35  // #
const START byte = 83 // S
const END byte = 69   // E
const VISIT byte = 79 // O

var NORTH = vec{0, -1}
var EAST = vec{1, 0}
var SOUTH = vec{0, 1}
var WEST = vec{-1, 0}
var dirs = []vec{NORTH, EAST, SOUTH, WEST}

var grid = [][]byte{}
var nodes = []point{}
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

	// walk nodes to find costs between
	paths := []node{{start, 0, 0, nil}}
	visited := make(map[point]int)
	for i := 0; i < len(paths); i++ {
		for head := paths[i]; head.pos != end; head = paths[i] {
			adj := findAdjacent(head)
			if len(adj) == 0 {
				// can't reach end on this path
				break
			}
			if i, ok := visited[head.pos]; ok && i < head.length {
				break
			}
			for j, other := range adj {
				dx := head.pos.x - other.x
				dy := head.pos.y - other.y
				if dx < 0 {
					dx = -dx
				}
				if dy < 0 {
					dy = -dy
				}
				cost := dx + dy + head.cost
				next := node{other, cost, head.length + 1, &head}
				if j == 0 {
					paths[i] = next
				} else {
					paths = append(paths, next)
				}
			}
			visited[head.pos] = head.length
		}
	}

	if *debug {
		fmt.Println(nodes)
		// fmt.Println(paths)
	}

	// pt1 look for shortest path in paths
	shortestPaths := []node{}
	for _, path := range paths {
		if path.pos == end {
			if len(shortestPaths) == 0 || path.length == shortestPaths[0].length {
				shortestPaths = append(shortestPaths, path)
			} else if path.length < shortestPaths[0].length {
				shortestPaths = []node{path}
			}
		}
	}
	slices.SortFunc(shortestPaths, func(p1 node, p2 node) int {
		return p1.cost - p2.cost
	})
	// work out how many turns to point in the initial direction
	var firstNode node
	for node := shortestPaths[0]; node.prev != nil; node = *node.prev {
		firstNode = node
	}
	firstDir := slices.Index(dirs, vec{unit(firstNode.pos.x - firstNode.prev.pos.x), unit(firstNode.pos.y - firstNode.prev.pos.y)})
	startDir := slices.Index(dirs, EAST)
	initialTurns := 0
	if startDir == (firstDir+1)%4 || startDir == (firstDir-1)%4 {
		initialTurns = 1 // turn 90 degrees
	} else if startDir != firstDir {
		initialTurns = 2 // turn 180 degress
	}

	// pt2 mark all VISITed points
	for _, path := range shortestPaths {
		if path.cost == shortestPaths[0].cost {
			for node := path; node.prev != nil; node = *node.prev {
				for y := min(node.pos.y, node.prev.pos.y); y <= max(node.pos.y, node.prev.pos.y); y++ {
					for x := min(node.pos.x, node.prev.pos.x); x <= max(node.pos.x, node.prev.pos.x); x++ {
						grid[y][x] = VISIT
					}
				}
			}
		}
	}

	// count the VISITs in the grid
	count := 0
	for y := 1; y < len(grid)-1; y++ {
		for x := 1; x < len(grid[y])-1; x++ {
			if grid[y][x] == VISIT {
				count += 1
			}
		}
	}

	if *debug {
		printGrid()
	}
	fmt.Println("Pt1:", (shortestPaths[0].length-1+initialTurns)*1000+shortestPaths[0].cost)
	fmt.Println("Pt2:", count)
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

func findAdjacent(p node) (adj []point) {
	for _, node := range nodes {
		dx := node.x - p.pos.x
		dy := node.y - p.pos.y
		if (dx == 0 || dy == 0) && noWallsBetween(p.pos, node) && !inPath(p, node) {
			adj = append(adj, node)
		}
	}
	return adj
}

func inPath(head node, p point) bool {
	for node := head; node.prev != nil; node = *node.prev {
		if node.pos == p {
			return true
		}
	}
	return false
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
		if grid[start.y+dir.dy][start.x+dir.dx] == EMPTY {
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
			if !slices.Contains(nodes, queue[i].pos) {
				nodes = append(nodes, queue[i].pos)
			}
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
			if slices.Contains(nodes, point{x, y}) {
				fmt.Print("X")
			} else {
				fmt.Print(string(grid[y][x]))
			}
		}
		fmt.Println()
	}
}
