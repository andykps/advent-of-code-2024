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
	
	count := 0
	for _, design := range designs {
		memo := make(map[string]bool)
		memo[""] = true
		if matchDesign(design, patterns, memo) {
			count += 1
		}
	}

	fmt.Println(count)
}

func matchDesign(design string, patterns []string, memo map[string]bool) bool {
	if res, ok := memo[design]; ok {
		return res
	}
	for _, pattern := range patterns {
		if strings.HasPrefix(design, pattern) {
			if (matchDesign(design[len(pattern):], patterns, memo)) {
				memo[design] = true
				return true
			}
		}
	}
	memo[design] = false
	return false
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
