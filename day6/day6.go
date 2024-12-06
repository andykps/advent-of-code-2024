package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
)

var GUARD = []byte{94, 62, 118, 60} // ^>V<
const (
	WALL  = 35
	FLOOR = 46
	X     = 88
)

func main() {
	flag.Parse()
	input := "input.txt"
	if len(flag.Args()) > 0 {
		input = flag.Args()[0]
	}

	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	grid := [][]byte{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Bytes()
		tmp := make([]byte, len(line))
		copy(tmp, line)
		grid = append(grid, tmp)
	}

	// fmt.Println([]byte("X"))
	x, y := findGuard(grid)
	for x > -1 && x < len(grid[0]) && y > -1 && y < len(grid) {
		x, y = moveGuard(grid, x, y)
		// printGrid(grid)
		// fmt.Println("", x, y)
	}
	// fmt.Println(x, y)
	printGrid(grid)

	total := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == X {
				total += 1
			}
		}
	}
	fmt.Println(total)
}

func findGuard(grid [][]byte) (x int, y int) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if slices.Contains(GUARD, grid[y][x]) {
				return x, y
			}
		}
	}
	return -1, -1
}

func printGrid(grid [][]byte) {
	for y := 0; y < len(grid); y++ {
		fmt.Println(string(grid[y]))
	}
}

func moveGuard(grid [][]byte, x int, y int) (int, int) {
	dir := slices.Index(GUARD, grid[y][x])
	switch dir {
	case 0: // up
		grid[y][x] = X
		y = y - 1
		if y < 0 {
			return x, y
		}
		if grid[y][x] == WALL {
			dir = 1
			y = y + 1
		}
	case 1: // right
		grid[y][x] = X
		x = x + 1
		if x >= len(grid[y]) {
			return x, y
		}
		if grid[y][x] == WALL {
			dir = 2
			x = x - 1
		}
	case 2: // down
		grid[y][x] = X
		y = y + 1
		if y >= len(grid) {
			return x, y
		}
		if grid[y][x] == WALL {
			dir = 3
			y = y - 1
		}
	case 3: // left
		grid[y][x] = X
		x = x - 1
		if x < 0 {
			return x, y
		}
		if grid[y][x] == WALL {
			dir = 0
			x = x + 1
		}
	default:
		log.Fatal("Guard not at expected position")
		return -1, -1
	}
	grid[y][x] = GUARD[dir]
	return x, y
}
