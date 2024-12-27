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

	runSimulation(wires, gates)

	if *debug {
		printWireState(wires)
	}

	num := calculateZNum(wires)
	fmt.Println("Part 1:", num)

	// pt2
	swapGates(gates, gate{"qfs", "whh", "AND", "z05"}, gate{"qfs", "whh", "XOR", "bpf"})
	swapGates(gates, gate{"skn", "spp", "OR", "z11"}, gate{"cgn", "cjh", "XOR", "hcc"})
	swapGates(gates, gate{"x35", "y35", "AND", "z35"}, gate{"khk", "tgs", "XOR", "fdw"})
	swapGates(gates, gate{"x24", "y24", "AND", "hqc"}, gate{"y24", "x24", "XOR", "qcw"})

	for i := 0; i < 45; i++ {
		setWires(wires, 1<<i, 1<<i)
		// printWireState(wires)
		runSimulation(wires, gates)
		num = calculateZNum(wires)
		if num != (1<<i)*2 {
			fmt.Printf("z%02d ", i+1)
		}
	}
	fmt.Println()

	// dumpGate(gates, "z25", 2, 0)

	// for i := 0; i < 6; i++ {
	// 	dumpGate(gates, fmt.Sprintf("z%02d", i), 99, 0)
	// 	fmt.Println()
	// }

	// fmt.Println()
	// for _, g := range gates {
	// 	if g.op == "XOR" {
	// 		if !strings.HasPrefix(g.out, "z") && !strings.HasPrefix(g.in1, "x") && !strings.HasPrefix(g.in1, "y") && !strings.HasPrefix(g.in2, "x") && !strings.HasPrefix(g.in2, "y") {
	// 			fmt.Printf("%s %s %s -> %s\n", g.in1, g.op, g.in2, g.out)
	// 		}
	// 	}
	// }

	// fmt.Println()
	// for _, g := range gates {
	// 	if strings.HasPrefix(g.out, "z") {
	// 		fmt.Printf("%s %s %s -> %s\n", g.in1, g.op, g.in2, g.out)
	// 	}
	// }

}

func dumpGate(gates []gate, wire string, maxdepth int, depth int) {
	if depth > maxdepth {
		return
	}
	for _, g := range gates {
		if g.out == wire {
			fmt.Printf("%*s %s %s -> %s\n", depth*2, g.in1, g.op, g.in2, g.out)
			dumpGate(gates, g.in1, maxdepth, depth+1)
			dumpGate(gates, g.in2, maxdepth, depth+1)
		}
	}
}

func swapGates(gates []gate, g1 gate, g2 gate) {
	// var gate1, gate2 *gate
	// for i, g := range gates {
	// 	if g1.in1 == g.in1 && g1.in2 == g.in2 && g1.op == g.op {
	// 		gate1 = &gates[i]
	// 	}
	// 	if g2.in1 == g.in1 && g2.in2 == g.in2 && g2.op == g.op {
	// 		gate2 = &gates[i]
	// 	}
	// }

	// tmp := gate2.out
	// (*gate2).out = gate1.out
	// (*gate1).out = tmp

	i1 := slices.Index(gates, g1)
	i2 := slices.Index(gates, g2)

	gates[i2].out = g1.out
	gates[i1].out = g2.out
}

func runSimulation(wires map[string]int, gates []gate) {
	cp := make([]gate, len(gates))
	copy(cp, gates)
	for next, ok := popNextGate(wires, &cp); ok; next, ok = popNextGate(wires, &cp) {
		simulateGate(*next, wires)
	}
}

func printWireState(wires map[string]int) {
	wireNames := []string{}
	for w := range wires {
		wireNames = append(wireNames, w)
	}
	slices.Sort(wireNames)
	for _, w := range wireNames {
		fmt.Println(w, wires[w])
	}
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

func setWires(wires map[string]int, x int, y int) {
	clear(wires)
	for i := 0; i < 45; i++ {
		if x&(1<<i) > 0 {
			wires[fmt.Sprintf("x%02d", i)] = 1
		} else {
			wires[fmt.Sprintf("x%02d", i)] = 0
		}
		if y&(1<<i) > 0 {
			wires[fmt.Sprintf("y%02d", i)] = 1
		} else {
			wires[fmt.Sprintf("y%02d", i)] = 0
		}
	}
}
