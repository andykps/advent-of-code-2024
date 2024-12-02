package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := "input.txt"
	if len(os.Args) > 1 {
		input = os.Args[1]
	}

	f, err := os.Open(input)
	if err != nil {
		log.Fatal("Use -input to provide input file\n", err)
	}
	defer f.Close()

	total := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		levels := make([]int, len(fields))
		for i := 0; i < len(fields); i++ {
			levels[i], _ = strconv.Atoi(fields[i])
		}
		// fmt.Println(levels)
		safe := true
		for i := 1; i < len(levels); i++ {
			step := levels[i] - levels[i-1]
			if step < -3 || step > 3 || step == 0 {
				// fmt.Println(step)
				safe = false
				break
			}
			if i > 1 {
				prevstep := levels[i-1] - levels[i-2]
				if float64(prevstep)/float64(step) < 0.0 {
					// fmt.Println(i, prevstep, step, prevstep/step)
					safe = false
					break
				}
			}
		}
		if safe {
			total += 1
		}
		fmt.Println(levels, safe)
	}
	fmt.Println(total)
}
