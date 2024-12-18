package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type vec struct {
	dx int
	dy int
}

const (
	EMPTY byte = 46 // .
	BYTE  byte = 35 // #
	PATH  byte = 79 // O
)

var (
	NORTH = vec{0, -1}
	EAST  = vec{1, 0}
	SOUTH = vec{0, 1}
	WEST  = vec{-1, 0}
	dirs  = []vec{NORTH, EAST, SOUTH, WEST}
	grid [][]byte
)


func main() {
	debug := flag.Bool("debug", false, "Output extra debug info")
	bytes := flag.Int("bytes", 1024, "Number of bytes to read from file")
	width := flag.Int("w", 71, "Width of grid")
	height := flag.Int("h", 71, "Height of grid")
	flag.Parse()
	input := "input.txt"
	if len(flag.Args()) > 0 {
		input = flag.Args()[0]
	}

	coords := readFile(input)

	grid = buildGrid(coords[0:*bytes], *width, *height)
	if *debug {
		printGrid()
	}

	fmt.Println("Part 1:", solveMaze(point{0, 0}, point{*width - 1, *height - 1}))

	// part 2
	upper := len(coords)
	lower := 0
	for upper != lower {
		mid := lower + (upper - lower) / 2
		grid = buildGrid(coords[0:mid], *width, *height)
		steps := solveMaze(point{0, 0}, point{*width - 1, *height - 1})
		if steps == -1 {
			upper = mid - 1
		} else {
			lower = mid + 1
		}
	}
	fmt.Printf("Part 2: %d,%d\n", coords[upper-1].x, coords[upper-1].y)
}

func solveMaze(start point, goal point) (steps int) {
	openSet := make(map[point]int)
	openSet[start] = 0

	cameFrom := make(map[point]point)

	gScore := make(map[point]int)
	gScore[start] = 0

	fScore := make(map[point]int)
	fScore[start] = heuristic(start, goal)

	for len(openSet) > 0 {
		current := lowest(openSet)
		if current == goal {
			return gScore[current]
		}
		delete(openSet, current)
		for _, dir := range dirs {
			neighbour := point{current.x+dir.dx, current.y+dir.dy}
			if validMove(grid, neighbour.x, neighbour.y) {
				if score, ok := gScore[neighbour]; !ok || gScore[current]+1 < score {
					cameFrom[neighbour] = current
					gScore[neighbour] = gScore[current] + 1
					fScore[neighbour] = gScore[current] + 1 + heuristic(neighbour, goal)
					openSet[neighbour] = fScore[neighbour]
				}
			}
		}
	}

	return -1 // no route
}

func heuristic(p1 point, p2 point) int {
	return (max(p1.x, p2.x) - min(p1.x, p2.x)) + (max(p1.y, p2.y) - min(p1.y, p2.y))
}

func lowest(scores map[point]int) (lowPoint point) {
	lowScore := math.MaxInt
	for p, s := range scores {
		if s < lowScore {
			lowScore = s
			lowPoint = p
		}
	}

	return
}

func validMove(grid [][]byte, x int, y int) bool {
	return y >= 0 && y < len(grid) && x >= 0 && x < len(grid[y]) && grid[y][x] != BYTE
}

func readFile(path string) (bytes []point) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(s[0])
		y, _ := strconv.Atoi(s[1])
		bytes = append(bytes, point{x, y})
	}
	return
}

func buildGrid(bytes []point, width int, height int) (grid [][]byte) {
	grid = make([][]byte, height)
	for y := range grid {
		grid[y] = slices.Repeat([]byte{EMPTY}, width)
	}
	for _, b := range bytes {
		grid[b.y][b.x] = BYTE
	}

	return
}

func printGrid() {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			fmt.Print(string(grid[y][x]))
		}
		fmt.Println()
	}
}