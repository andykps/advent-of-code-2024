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

type Rule struct {
	left  int
	right int
}

var rules = []Rule{}

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

	scanner := bufio.NewScanner(f)
	readRules(scanner)

	// now we've read the rules, scanner is ready to read page numbers
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		pages := []int{}
		for i := 0; i < len(parts); i++ {
			page, _ := strconv.Atoi(parts[i])
			pages = append(pages, page)
		}
		if checkRules(pages) {
			total += pages[len(pages)/2]
		}
	}

	fmt.Println(total)
}

func readRules(scanner *bufio.Scanner) {
	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.TrimSpace(line)) == 0 {
			return
		}
		parts := strings.Split(line, "|")
		left, _ := strconv.Atoi(parts[0])
		right, _ := strconv.Atoi(parts[1])
		rule := Rule{left, right}
		rules = append(rules, rule)
	}
}

func checkRules(pages []int) bool {
	for i := 0; i < len(rules); i++ {
		rule := rules[i]
		matches := []int{}
		for p := 0; p < len(pages); p++ {
			if pages[p] == rule.left || pages[p] == rule.right {
				matches = append(matches, pages[p])
				continue
			}
		}
		if len(matches) == 2 && matches[0] != rule.left {
			return false
		}
	}
	return true
}
