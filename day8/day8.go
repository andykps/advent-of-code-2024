package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

const (
	EMPTY    byte = 46
	ANTINODE byte = 35
)

type point struct {
	x int
	y int
}

func main() {
	flag.Parse()
	input := "input.txt"
	if len(flag.Args()) > 0 {
		input = flag.Args()[0]
	}
	grid := readGridFromFile(input)
	antennas := locateAntennas(grid)

	for _, coords := range antennas {
		for i := 0; i < len(coords); i++ {
			for j := i + 1; j < len(coords); j++ {
				x1 := coords[i].x + coords[i].x - coords[j].x
				y1 := coords[i].y + coords[i].y - coords[j].y
				x2 := coords[j].x + coords[j].x - coords[i].x
				y2 := coords[j].y + coords[j].y - coords[i].y
				if x1 >= 0 && x1 < len(grid[0]) && y1 >= 0 && y1 < len(grid) {
					grid[y1][x1] = ANTINODE
				}
				if x2 >= 0 && x2 < len(grid[0]) && y2 >= 0 && y2 < len(grid) {
					grid[y2][x2] = ANTINODE
				}
			}
		}
	}

	total := countInGrid(grid, ANTINODE)
	fmt.Println(total)

	// printGrid(grid)
	// fmt.Println(antennas)
}

func readGridFromFile(path string) (grid [][]byte) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Bytes()
		tmp := make([]byte, len(line))
		copy(tmp, line)
		grid = append(grid, tmp)
	}
	return
}

func printGrid(grid [][]byte) {
	for y := 0; y < len(grid); y++ {
		fmt.Println(string(grid[y]))
	}
}

func locateAntennas(grid [][]byte) (antennas map[byte][]point) {
	antennas = make(map[byte][]point)
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			cell := grid[y][x]
			if cell != EMPTY {
				if elem, ok := antennas[cell]; ok {
					antennas[cell] = append(elem, point{x, y})
				} else {
					antennas[cell] = []point{{x, y}}
				}
			}
		}
	}
	return
}

func countInGrid(grid [][]byte, target byte) (count int) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == target {
				count += 1
			}
		}
	}
	return
}
