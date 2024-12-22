package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

	numbers := []int{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		n, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		numbers = append(numbers, n)
	}

	// pt 1
	total := 0
	for _, n := range numbers {
		for i := 0; i < 2000; i++ {
			n = prune(mix(n, n*64))
			n = prune(mix(n, n/32))
			n = prune(mix(n, n*2048))
		}
		total += n
	}
	fmt.Println("Part 1:", total)

	// pt 2
	var prev int
	allbuyers := [][][2]int{}
	for _, n := range numbers {
		changes := [][2]int{}
		for i := 0; i <= 2000; i++ {
			ones := n % 10
			if i > 0 {
				diff := ones - prev
				changes = append(changes, [2]int{ones, diff})
			}
			prev = ones

			n = prune(mix(n, n*64))
			n = prune(mix(n, n/32))
			n = prune(mix(n, n*2048))
		}
		allbuyers = append(allbuyers, changes)
	}

	// find all possible sequences in our data (much less than 4^19)
	allSeqs := make(map[[4]int]bool)
	for _, buyer := range allbuyers {
		for i := 3; i < len(buyer); i++ {
			seq := [4]int{buyer[i-3][1], buyer[i-2][1], buyer[i-1][1], buyer[i][1]}
			allSeqs[seq] = true
		}
	}

	// I'm not sure if modifying the map allSeqs (if I made it int value instead of bool)
	// would change the iteration, so create duplicate map to keep track of price per seq
	allPrices := make(map[[4]int]int)
	for seq := range allSeqs {
		BUYER: for _, buyer := range allbuyers {
			for i := 3; i < len(buyer); i++ {
				if seq == [4]int{buyer[i-3][1], buyer[i-2][1], buyer[i-1][1], buyer[i][1]} {
					allPrices[seq] += buyer[i][0]
					continue BUYER
				}
			}
		}
	}

	// just find the max price
	max := 0
	for _, price := range allPrices {
		if price > max {
			max = price
		}
	}

	fmt.Println("Part 2:", max)
}

func mix(sec int, n int) int {
	return n ^ sec
}

func prune(sec int) int {
	return sec % 16777216
}
