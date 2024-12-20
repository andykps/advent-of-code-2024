package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"sync"
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
	WALL  byte = 35 // #
	START byte = 83 // S
	END   byte = 69 // E
	PATH  byte = 79 // O
)

var (
	NORTH = vec{0, -1}
	EAST  = vec{1, 0}
	SOUTH = vec{0, 1}
	WEST  = vec{-1, 0}
	dirs  = []vec{NORTH, EAST, SOUTH, WEST}
)

func main() {
	flag.Parse()
	input := "input.txt"
	if len(flag.Args()) > 0 {
		input = flag.Args()[0]
	}

	grid := readGridFromFile(input)
	path := solveMaze(grid)
	fullTime := len(path) - 1
	fmt.Printf("Non cheater path takes %d picoseconds\n", fullTime)

	walls := findCheatableWalls(grid)
	min := 100
	count := 0
	concurrencyLimit := 40

	var wg sync.WaitGroup
	var mu sync.Mutex

	worker := func() {
		for {
			mu.Lock()
			if len(walls) == 0 {
				mu.Unlock()
				break
			}
			cheat := walls[0]
			walls = walls[1:]
			mu.Unlock()

			grid := readGridFromFile(input)
			grid[cheat.y][cheat.x] = EMPTY
			path := solveMaze(grid)
	
			if len(path)+min-2 < fullTime {
				fmt.Println(len(path)-1)
				mu.Lock()
				count += 1
				mu.Unlock()
			}
		}

		wg.Done()
	}

	for i := 0; i < concurrencyLimit; i++ {
		wg.Add(1)
		go worker()
	}

	wg.Wait()

	fmt.Printf("%d cheats save at least %d picoseconds", count, min)
	// plotPath(grid, path)
	// printGrid(grid)
}

func solveMaze(grid [][]byte) (path []point) {
	start := findInGrid(grid, START)
	goal := findInGrid(grid, END)

	openSet := make(map[point]int)
	openSet[start] = heuristic(start, goal)

	cameFrom := make(map[point]point)

	gScore := make(map[point]int)
	gScore[start] = 0

	for len(openSet) > 0 {
		current := lowest(openSet)
		if current == goal {
			return buildPath(cameFrom, current)
		}
		delete(openSet, current)
		for _, dir := range dirs {
			neighbour := point{current.x + dir.dx, current.y + dir.dy}
			if validMove(grid, neighbour.x, neighbour.y) {
				if score, ok := gScore[neighbour]; !ok || gScore[current]+1 < score {
					cameFrom[neighbour] = current
					gScore[neighbour] = gScore[current] + 1
					openSet[neighbour] = gScore[current] + 1 + heuristic(neighbour, goal)
				}
			}
		}
	}

	return nil
}

func heuristic(p1 point, p2 point) int {
	x := p1.x - p2.x
	y := p1.y - p2.y
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	return x + y
	// return (max(p1.x, p2.x) - min(p1.x, p2.x)) + (max(p1.y, p2.y) - min(p1.y, p2.y))
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
	return grid[y][x] != WALL && y >= 0 && y < len(grid) && x >= 0 && x < len(grid[y])
}

func buildPath(cameFrom map[point]point, current point) (path []point) {
	path = append(path, current)
	for from, ok := cameFrom[current]; ok; from, ok = cameFrom[current] {
		path = append([]point{from}, path...)
		current = from
	}
	return
}

func plotPath(grid [][]byte, path []point) {
	for _, p := range path {
		grid[p.y][p.x] = PATH
	}
}

func findInGrid(grid [][]byte, b byte) point {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == b {
				return point{x, y}
			}
		}
	}
	log.Fatalf("Can't find '%c' in grid\n", b)
	return point{-1, -1} // would never get here?
}

func findCheatableWalls(grid [][]byte) (walls []point) {
	for y := 1; y < len(grid)-1; y++ {
		for x := 1; x < len(grid[y])-1; x++ {
			neighbours := 0
			for _, d := range dirs {
				if grid[y+d.dy][x+d.dx] == WALL {
					neighbours += 1
				}
			}
			if neighbours < 3 {
				walls = append(walls, point{x, y})
			}
		}
	}
	return
}

func readGridFromFile(path string) (grid [][]byte) {
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
		y += 1
	}
	return
}

func printGrid(grid [][]byte) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			fmt.Print(string(grid[y][x]))
		}
		fmt.Println()
	}
}
