package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

type plot struct {
	plant byte
	borders int
	region int
}

var maxregion int = 0
var dirs = [][2]int{{0,-1},{1,0},{0,1},{-1,0}}

func main() {
	debug := flag.Bool("debug", false, "Output extra debug info")
	flag.Parse()
	input := "input.txt"
	if len(flag.Args()) > 0 {
		input = flag.Args()[0]
	}

	grid := readGridFromFile(input)
	if *debug {
		// print plants
		printGrid(grid, func(p plot) string {return string(p.plant)})
	}

	groupRegions(grid)
	if *debug {
		// print plants
		printGrid(grid, func(p plot) string {return string(p.region)})
	}

	countAllBorders(grid)
	if *debug {
		// print borders
		printGrid(grid, func(p plot) string {return strconv.Itoa(p.borders)})
	}

	price := calculatePrice(grid)
	fmt.Println(price)

}

func readGridFromFile(path string) (grid [][]plot) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := []plot{}
		for _, b := range scanner.Bytes() {
			row = append(row, plot{b, 0, -1})
		}
		grid = append(grid, row)
	}
	return
}

func printGrid(grid [][]plot, f func(plot) string) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			fmt.Print(f(grid[y][x]))
		}
		fmt.Println()
	}
}

func countAllBorders(grid [][]plot) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			grid[y][x].borders = 4 - countNeighbours(grid, x, y)
		}
	}
}

func countNeighbours(grid [][]plot, x int, y int) (neighbours int) {
	for _, dir := range dirs {
		if x1, y1 := x+dir[0], y+dir[1];
			x1 >= 0 && y1 >= 0 && y1 < len(grid) && x1 < len(grid[y1]) && grid[y1][x1].plant == grid[y][x].plant {
			neighbours += 1
			grid[y1][x1].region = grid[y][x].region
		}
	}
	return
}

func calculatePrice(grid [][]plot) (price int) {
	regions := make(map[int][]plot)
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			plot := grid[y][x]
			regions[plot.region] = append(regions[plot.region], plot)
		}
	}

	for _, region := range regions {
		plotCount := len(region)
		borders := 0
		for _, plot := range region {
			borders += plot.borders
		}
		price += plotCount * borders
		fmt.Println(region[0].region, string(region[0].plant), borders, plotCount, price)
	}
	return
}

func groupRegions(grid [][]plot) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			plot := &grid[y][x]
			if plot.region == -1 {
				plot.region = maxregion
				maxregion += 1
				groupNeighbours(grid, x, y)
			}
		}
	}
}

func groupNeighbours(grid [][]plot, x int, y int) {
	for _, dir := range dirs {
		if x1, y1 := x+dir[0], y+dir[1]; x1 >= 0 && y1 >= 0 && y1 < len(grid) && x1 < len(grid[y1]) && grid[y1][x1].region == -1 && grid[y1][x1].plant == grid[y][x].plant {
			grid[y1][x1].region = grid[y][x].region
			groupNeighbours(grid, x1, y1)
		}
	}
}