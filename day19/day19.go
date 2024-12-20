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

	patterns, designs := readFile(input)

	pt1 := 0
	pt2 := 0
	for _, design := range designs {
		memo := make(map[string]int)
		memo[""] = 1

		c := countDesigns(design, patterns, memo)
		if c > 0 {
			pt1 += 1
		}
		pt2 += c
	}

	fmt.Println("Pt1:", pt1)
	fmt.Println("Pt2:", pt2)
}

func countDesigns(design string, patterns []string, memo map[string]int) int {
	if res, ok := memo[design]; ok {
		return res
	}
	count := 0
	for _, pattern := range patterns {
		if strings.HasPrefix(design, pattern) {
			count += countDesigns(design[len(pattern):], patterns, memo)
		}
	}
	memo[design] = count
	return count
}

func readFile(path string) (patterns []string, designs []string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if len(patterns) == 0 {
			patterns = strings.Split(line, ", ")
		} else if len(strings.TrimSpace(line)) > 0 {
			designs = append(designs, line)
		}
	}
	return
}
