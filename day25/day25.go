package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	input := "input.txt"
	if len(flag.Args()) > 0 {
		input = flag.Args()[0]
	}

	keys, locks := readFile(input)

	fitCount := 0
	for _, lock := range locks {
		for _, key := range keys {
			fits := true
			for pin := 0; pin < 5; pin++ {
				if (lock[pin] + key[pin] > 5) {
					fits = false
					break
				}
			}
			if fits {
				fitCount += 1
			}
		}
	}

	fmt.Println(fitCount)

}

func readFile(path string) (keys [][]int, locks [][]int) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	row := 0
	islock := false
	heights := make([]int, 5)
	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		if row == 0 && line[0] == '#' {
			islock = true
		} else if row > 5 {
			if islock {
				locks = append(locks, heights)
			} else {
				keys = append(keys, heights)
			}
			row = -1
			islock = false
			heights = make([]int, 5)
		} else if row > 0 {
			for i := 0; i < len(line); i++ {
				if line[i] == '#' {
					heights[i] += 1
				}
			}
		}
		
		row += 1
	}
	return
}