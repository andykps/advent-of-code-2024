package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

const GRID byte = 0
const MOVES byte = 1
const UP byte = 94    // ^
const DOWN byte = 118 // v
const LEFT byte = 60  // <
const RIGHT byte = 62 // >
const ROBOT byte = 64 // @
const EMPTY byte = 46 // .
const BOX byte = 79   // O
const WALL byte = 35  // #

type vec struct {
	dx int
	dy int
}
type point struct {
	x int
	y int
}

var grid = [][]byte{}
var moves = []byte{}
var robot = point{-1, -1}

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
		fmt.Println(string(moves), "\n")
	}
	robot = findRobot()

	for _, move := range moves {
		processMove(move)
		if *debug {
			fmt.Println("Move:", string(move))
			printGrid()
			fmt.Println()
		}
	}

	gps := calcBoxesGPS()
	fmt.Println(gps)
}

func readGridFromFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	mode := GRID
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := make([]byte, len(scanner.Bytes()))
		copy(line, scanner.Bytes())
		if len(line) == 0 {
			mode = MOVES
			continue
		}
		if mode == GRID {
			grid = append(grid, line)
		} else if mode == MOVES {
			moves = append(moves, line...)
		}
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

func processMove(move byte) {
	dir := getDirection(move)
	newX := robot.x + dir.dx
	newY := robot.y + dir.dy
	if grid[newY][newX] == EMPTY {
		grid[robot.y][robot.x] = EMPTY
		grid[newY][newX] = ROBOT
		robot = point{newX, newY}
	} else if grid[newY][newX] == BOX {
		// keep looking in dir for an empty space or wall
		for i := 1; true; i++ {
			if grid[newY+i*dir.dy][newX+i*dir.dx] == EMPTY {
				// we can push into this space
				for j := i; j >= 0; j-- {
					grid[newY+j*dir.dy][newX+j*dir.dx] = grid[newY+(j-1)*dir.dy][newX+(j-1)*dir.dx]
				}
				grid[robot.y][robot.x] = EMPTY
				grid[newY][newX] = ROBOT
				robot = point{newX, newY}
				break
			} else if grid[newY+i*dir.dy][newX+i*dir.dx] == WALL {
				// can't move, give up
				break
			} // else it can only be a box so keep looking for a space
		}
	}
}

func getDirection(move byte) (dir vec) {
	switch move {
	case UP:
		dir = vec{0, -1}
	case DOWN:
		dir = vec{0, 1}
	case LEFT:
		dir = vec{-1, 0}
	case RIGHT:
		dir = vec{1, 0}
	default:
		log.Fatal("Invalid move", move, string(move))
	}
	return
}

func findRobot() (coords point) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == ROBOT {
				return point{x, y}
			}
		}
	}
	panic("Can't find robot")
}

func calcBoxesGPS() (gps int) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == BOX {
				gps += 100*y + x
			}
		}
	}
	return
}
