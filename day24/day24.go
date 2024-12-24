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
	"strings"
)

type gate struct {
	in1 string
	in2 string
	op  string
	out string
}

func main() {
	debug := flag.Bool("debug", false, "Should debug mode be turned on?")
	flag.Parse()
	input := "input.txt"
	if len(flag.Args()) > 0 {
		input = flag.Args()[0]
	}

	wires, gates := readFile(input)
	// fmt.Println(wires, gates)

	for next, ok := popNextGate(wires, &gates); ok; next, ok = popNextGate(wires, &gates) {
		if *debug {
			fmt.Println(next)
		}
		simulateGate(*next, wires)
	}

	if *debug {
		wireNames := []string{}
		for w := range wires {
			wireNames = append(wireNames, w)
		}
		slices.Sort(wireNames)
		for _, w := range wireNames {
			fmt.Println(w, wires[w])
		}
	}

	num := calculateZNum(wires)
	fmt.Println(num)
}

func readFile(path string) (wires map[string]int, gates []gate) {
	reWire := regexp.MustCompile(`([a-z0-9]{3}): (\d)`)
	reGates := regexp.MustCompile(`([a-z0-9]{3}) ([A-Z]{2,3}) ([a-z0-9]{3}) -> ([a-z0-9]{3})`)
	wires = make(map[string]int)

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if m := reWire.FindStringSubmatch(line); len(m) > 0 {
			wires[m[1]], _ = strconv.Atoi(m[2])
		} else if m := reGates.FindStringSubmatch(line); len(m) > 0 {
			gates = append(gates, gate{m[1], m[3], m[2], m[4]})
		}
	}
	return
}

func popNextGate(wires map[string]int, gates *[]gate) (next *gate, ok bool) {
	for _, g := range *gates {
		_, ok1 := wires[g.in1]
		if _, ok2 := wires[g.in2]; ok1 && ok2 {
			*gates = slices.DeleteFunc(*gates, func(z gate) bool {
				return z == g
			})
			return &g, true
		}
	}
	return nil, false
}

func simulateGate(g gate, wires map[string]int) {
	switch g.op {
	case "AND":
		wires[g.out] = wires[g.in1] & wires[g.in2]
	case "OR":
		wires[g.out] = wires[g.in1] | wires[g.in2]
	case "XOR":
		wires[g.out] = wires[g.in1] ^ wires[g.in2]
	default:
		panic(fmt.Sprintf("Unrecognised op %s", g.op))
	}
}

func calculateZNum(wires map[string]int) (num int) {
	zeds := []string{}
	for w := range wires {
		if strings.HasPrefix(w, "z") {
			zeds = append(zeds, w)
		}
	}
	slices.Sort(zeds)

	for i, w := range zeds {
		num += wires[w] << i
	}
	return
}
