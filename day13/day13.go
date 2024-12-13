package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Coord struct {
	x int
	y int
}

type Machine struct {
	buttonA Coord
	buttonB Coord
	prize   Coord
}

func main() {
	debug := flag.Bool("debug", false, "Output extra debug info")
	pt2 := flag.Bool("pt2", false, "Should we offset the prizes for part 2?")
	flag.Parse()
	input := "input.txt"
	if len(flag.Args()) > 0 {
		input = flag.Args()[0]
	}

	var machines []Machine
	if (*pt2) {
		machines = readFromFile(input, 10000000000000)
	} else {
		machines = readFromFile(input, 0)
	}
	if *debug {
		fmt.Println(machines)
	}

	total := 0
	for _, m := range machines {
		a, b := calculateIntersection(m.buttonA.x, m.buttonB.x, -m.prize.x, m.buttonA.y, m.buttonB.y, -m.prize.y)
		if a == math.Trunc(a) && b == math.Trunc(b) {
			if *debug {
				fmt.Println(a*3 + b)
			}
			total += int(a*3) + int(b)
		}
	}
	fmt.Println(total)
}

func calculateIntersection(a1 int, b1 int, c1 int, a2 int, b2 int, c2 int) (x float64, y float64) {
	// finding intersection of 2 lines
	// https://www.cuemath.com/geometry/intersection-of-two-lines/
	x = (float64(b1)*float64(c2) - float64(b2)*float64(c1)) / (float64(a1)*float64(b2) - float64(a2)*float64(b1))
	y = (float64(c1)*float64(a2) - float64(c2)*float64(a1)) / (float64(a1)*float64(b2) - float64(a2)*float64(b1))
	return
}

func readFromFile(path string, prizeOffset int) (machines []Machine) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	re := regexp.MustCompile(`([^:]+): X[+=](\d+), Y[+=](\d+)`)

	m := Machine{Coord{-1, -1}, Coord{-1, -1}, Coord{-1, -1}}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			machines = append(machines, m)
			m = Machine{Coord{-1, -1}, Coord{-1, -1}, Coord{-1, -1}}
			continue
		}
		matches := re.FindStringSubmatch(scanner.Text())
		x, _ := strconv.Atoi(matches[2])
		y, _ := strconv.Atoi(matches[3])
		switch matches[1] {
		case "Button A":
			m.buttonA = Coord{x, y}
		case "Button B":
			m.buttonB = Coord{x, y}
		case "Prize":
			m.prize = Coord{x + prizeOffset, y + prizeOffset}
		}
	}
	machines = append(machines, m)
	return
}
