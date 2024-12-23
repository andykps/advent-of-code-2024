package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
)

func main() {
	debug := flag.Bool("debug", false, "Output extra debug info")
	flag.Parse()
	input := "input.txt"
	if len(flag.Args()) > 0 {
		input = flag.Args()[0]
	}
	grid := readGridFromFile(input)
	if *debug {
		printGrid(grid)
	}

	trailheads := findTrailheads(grid)
	pt1Total := 0
	pt2Total := 0
	for _, start := range trailheads {
		visited9s := followTrail(grid, start)

		if *debug {
			fmt.Println(visited9s)
		}

		deduped9s := []point{}
		for _, p := range visited9s {
			if !slices.Contains(deduped9s, p) {
				deduped9s = append(deduped9s, p)
			}
		}

		pt1Total += len(deduped9s)
		pt2Total += len(visited9s)
	}
	fmt.Println("Part 1:", pt1Total)
	fmt.Println("Part 1:", pt2Total)
}

type point struct {
	x int
	y int
}

func readGridFromFile(path string) (grid [][]int) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Bytes()
		ints := make([]int, len(line))
		for i, b := range line {
			ints[i] = int(b - 48)
		}
		grid = append(grid, ints)
	}
	return
}

func printGrid(grid [][]int) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			fmt.Print(grid[y][x])
		}
		fmt.Println()
	}
}

func findTrailheads(grid [][]int) (trailheads []point) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 0 {
				trailheads = append(trailheads, point{x, y})
			}
		}
	}
	return
}

var DIR = []point{{0,-1},{1,0},{0,1},{-1,0}}

func followTrail(grid [][]int, start point) (visited9s []point) {
	queue := []point{start}
	for len(queue) > 0 {
		p := queue[0]
		height := grid[p.y][p.x]
		for _, d := range DIR {
			x := p.x + d.x
			y := p.y + d.y
			if x >= 0 && x < len(grid[0]) && y >= 0 && y < len(grid) {
				if grid[y][x] == height+1 {
					p2 := point{x, y}
					if grid[y][x] == 9 {
						visited9s = append(visited9s, p2)
					} else {
						queue = append(queue, p2)
					}
				}
			}
		}
		queue = queue[1:]
	}
	return
}
