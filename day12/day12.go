package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

type dir struct {
	dx int
	dy int
}
type plot struct {
	plant byte
	borders int
	region int
	x int
	y int
}

var maxregion int = 0
var dirs = []dir{{0,-1},{1,0},{0,1},{-1,0}} // N, E, S, W

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
		printGrid(grid, func(p plot) string {return string(byte(48+p.region))})
	}

	countAllBorders(grid)
	if *debug {
		// print borders
		printGrid(grid, func(p plot) string {return strconv.Itoa(p.borders)})
	}

	fmt.Println("Part 1:", calculatePrice(grid))
	fmt.Println("Part 2:", part2(grid))

}

func readGridFromFile(path string) (grid [][]plot) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	y := 0
	for scanner.Scan() {
		row := []plot{}
		x := 0
		for _, b := range scanner.Bytes() {
			row = append(row, plot{b, 0, -1, x, y})
			x += 1
		}
		grid = append(grid, row)
		y += 1
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
	fmt.Println()
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
		if x1, y1 := x+dir.dx, y+dir.dy;
			validCoords(grid, x1, y1) && grid[y1][x1].plant == grid[y][x].plant {
			neighbours += 1
		}
	}
	return
}

func calculatePrice(grid [][]plot) (price int) {
	regions := getRegions(grid)

	for _, region := range regions {
		plotCount := len(region)
		borders := 0
		for _, plot := range region {
			borders += plot.borders
		}
		price += plotCount * borders
	}
	return
}

func part2(grid [][]plot) (price int) {
	regions := getRegions(grid)

	for _, region := range regions {
		corners := 0
		for _, p := range region {
			for i := 0; i < len(dirs); i++ {
				dx1 := dirs[i].dx
				dy1 := dirs[i].dy
				dx2 := dirs[(i+1)%4].dx
				dy2 := dirs[(i+1)%4].dy
				x1 := p.x + dx1
				y1 := p.y + dy1
				x2 := p.x + dx2
				y2 := p.y + dy2
				if (!validCoords(grid, x1, y1) || grid[y1][x1].region != p.region) && (!validCoords(grid, x2, y2) || grid[y2][x2].region != p.region) {
					corners += 1
				}
				if validCoords(grid, x1, y1) && grid[y1][x1].region == p.region && validCoords(grid, x2, y2) && grid[y2][x2].region == p.region {
					// we don't need to validate coords because it's the corner between 2 known valid coords
					if grid[p.y + nonzero(dy1, dy2)][p.x + nonzero(dx1, dx2)].region != p.region {
						corners += 1
					}
				}
			}
		}
		price += corners * len(region)
	}
	return
}

func nonzero(v1 int, v2 int) int {
	if (v1 == 0) {
		return v2
	}
	return v1
}

func getRegions(grid [][]plot) (regions map[int][]plot) {
	regions = make(map[int][]plot)
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			plot := grid[y][x]
			regions[plot.region] = append(regions[plot.region], plot)
		}
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
		if x1, y1 := x+dir.dx, y+dir.dy; validCoords(grid, x1, y1) && grid[y1][x1].region == -1 && grid[y1][x1].plant == grid[y][x].plant {
			grid[y1][x1].region = grid[y][x].region
			groupNeighbours(grid, x1, y1)
		}
	}
}

func validCoords(grid [][]plot, x int, y int) bool {
	return x >= 0 && y >= 0 && y < len(grid) && x < len(grid[y])
}