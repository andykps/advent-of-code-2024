package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
)

type Bot struct {
	id int
	x  int
	y  int
	dx int
	dy int
}

var grid [][][]*Bot
var bots = []*Bot{}

var gridWidth *int
var gridHeight *int

var maxId int = 0

func main() {
	debug := flag.Bool("debug", false, "Output extra debug info")
	gridWidth = flag.Int("w", 101, "Width of grid (default=101)")
	gridHeight = flag.Int("h", 103, "Height of grid (default=103)")
	iterations := flag.Int("it", 100, "Number of iterations/seconds (default=100)")
	pt2 := flag.Bool("pt2", false, "Should we be doing part 2?")
	flag.Parse()
	input := "input.txt"
	if len(flag.Args()) > 0 {
		input = flag.Args()[0]
	}

	grid = make([][][]*Bot, *gridHeight)
	for y := 0; y < *gridHeight; y++ {
		grid[y] = make([][]*Bot, *gridWidth)
	}

	readBotsFromFile(input)
	if *debug {
		printGrid()
		fmt.Println()
	}

	for i := 0; i < *iterations; i++ {
		moveBots()
		if *debug {
			printGrid()
			fmt.Println()
		}
		if (*pt2 && checkForLines(10)) {
			printGrid()
			fmt.Println("Tree after", i+1, "seconds")
			break
		}
	}

	sf := calcSafetyFactor()
	fmt.Println(sf)
}

var re = regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)

func readBotsFromFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		matches := re.FindStringSubmatch(scanner.Text())
		x, _ := strconv.Atoi(matches[1])
		y, _ := strconv.Atoi(matches[2])
		dx, _ := strconv.Atoi(matches[3])
		dy, _ := strconv.Atoi(matches[4])
		bot := Bot{maxId, x, y, dx, dy}
		maxId += 1
		bots = append(bots, &bot)
		grid[y][x] = append(grid[y][x], &bot)
	}

}

func printGrid() {
	for y := 0; y < *gridHeight; y++ {
		for x := 0; x < *gridWidth; x++ {
			c := len(grid[y][x])
			if c > 0 {
				fmt.Print(c)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	// for _, bot := range bots {
	// 	fmt.Printf("(%d, %d) ", bot.x, bot.y)
	// }
	fmt.Println()
}

func moveBots() {
	for _, bot := range bots {
		newX := bot.x + bot.dx
		newY := bot.y + bot.dy
		if newX < 0 {
			newX = *gridWidth + newX
		}
		if newX > *gridWidth-1 {
			newX = newX - *gridWidth
		}
		if newY < 0 {
			newY = *gridHeight + newY
		}
		if newY > *gridHeight-1 {
			newY = newY - *gridHeight
		}
		i := slices.IndexFunc(grid[bot.y][bot.x], func(p *Bot) bool {
			return p.id == bot.id
		})
		if i == -1 {
			log.Fatal("Bot not found in grid at expected position")
		}
		// remove from old position
		grid[bot.y][bot.x] = append(grid[bot.y][bot.x][:i], grid[bot.y][bot.x][(i+1):]...)
		// add to new position
		grid[newY][newX] = append(grid[newY][newX], bot)
		// update bot
		bot.x = newX
		bot.y = newY
	}
}

func calcSafetyFactor() (sf int) {
	nw := 0
	ne := 0
	for y := 0; y < *gridHeight / 2; y++ {
		for x := 0; x < *gridWidth / 2; x++ {
			nw += len(grid[y][x])
		}
		for x := *gridWidth / 2 + 1; x < *gridWidth; x++ {
			ne += len(grid[y][x])
		}
	}
	sw := 0
	se := 0
	for y := *gridHeight / 2 + 1; y < *gridHeight; y++ {
		for x := 0; x < *gridWidth / 2; x++ {
			sw += len(grid[y][x])
		}
		for x := *gridWidth / 2 + 1; x < *gridWidth; x++ {
			se += len(grid[y][x])
		}
	}
	fmt.Println(nw, ne, sw, se)
	return nw * ne * sw * se
}

func checkForLines(minLength int) bool {
	for _, bot := range bots {
		if (bot.x + minLength < *gridWidth) {
			var i int
			for i = 1; i < minLength; i++ {
				if len(grid[bot.y][bot.x + i]) == 0 {
					break
				}
			}
			if i == minLength - 1 {
				return true
			}
		}
	}
	return false
}