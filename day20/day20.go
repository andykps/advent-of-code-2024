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

type point struct {
	x int
	y int
}

type vec struct {
	dx int
	dy int
}

type cellCost struct {
	x    int
	y    int
	cost int
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

	costs, path := fillMaze(grid)

	fullTime := len(path) - 1
	fmt.Printf("Non cheater path takes %d picoseconds\n", fullTime)

	{
		count := 0
		for _, cell := range path {
			shortcuts := calculateShortcuts(2)
			count += countCheats(costs, cell.x, cell.y, shortcuts)
		}
		fmt.Println("Part 1:", count)
	}

	{
		count := 0
		for _, cell := range path {
			shortcuts := calculateShortcuts(20)
			count += countCheats(costs, cell.x, cell.y, shortcuts)
		}
		fmt.Println("Part 2:", count)
	}

}

func calculateShortcuts(cheatLength int) (costs []cellCost) {
	for y := -cheatLength; y < cheatLength+1; y++ {
		for x := -cheatLength; x < cheatLength+1; x++ {
			// md is manhattan distance
			if md := abs(y) + abs(x); md <= cheatLength && !(x == 0 && y == 0) {
				costs = append(costs, cellCost{x, y, md})
			}
		}
	}
	return
}

func abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}

func countCheats(costs [][]int, x int, y int, shortcuts []cellCost) (cheats int) {
	minSaving := 100

	for _, sc := range shortcuts {
		x1, y1 := x+sc.x, y+sc.y
		if y1 >= 0 && y1 < len(costs) && x1 >= 0 && x1 < len(costs[y1]) && costs[y][x]-minSaving-sc.cost >= costs[y1][x1] {
			cheats += 1
		}
	}
	return
}

func fillMaze(grid [][]byte) (costs [][]int, coords []point) {
	end := findInGrid(grid, END)
	costs = make([][]int, len(grid))
	for y := 0; y < len(grid); y++ {
		costs[y] = slices.Repeat([]int{math.MaxInt}, len(grid[y]))
	}

	queue := []cellCost{{end.x, end.y, 0}}
	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]

		if costs[q.y][q.x] < math.MaxInt {
			continue
		}

		costs[q.y][q.x] = q.cost
		coords = append(coords, point{q.x, q.y})

		for _, d := range dirs {
			x1 := q.x + d.dx
			y1 := q.y + d.dy

			if grid[y1][x1] == EMPTY || grid[y1][x1] == START {
				queue = append(queue, cellCost{x1, y1, q.cost + 1})
			}
		}
	}
	return
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
