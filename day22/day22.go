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

	// numbers = []int{123}

	// pt 2
	var prev int
	allbuyers := [][][2]int{}
	for _, n := range numbers {
		changes := [][2]int{}
		for i := 0; i <= 2000; i++ {
			ones := n - (n / 10 * 10)
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

	max := 0
	for s1 := -9; s1 < 10; s1++ {
		for s2 := -9; s2 < 10; s2++ {
			for s3 := -9; s3 < 10; s3++ {
				for s4 := -9; s4 < 10; s4++ {
					sequence := [4]int{s1, s2, s3, s4}
					prices := make([]int, len(allbuyers))

				BUYER:
					for b, buyer := range allbuyers {
						i := 0
						for _, change := range buyer {
							if change[1] == sequence[i] {
								if i == 3 {
									prices[b] = change[0]
									continue BUYER
								}
								i += 1
							} else {
								i = 0
							}
						}
					}
					total = 0
					for _, p := range prices {
						total += p
					}
					if total > max {
						max = total
					}
				}
			}
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
