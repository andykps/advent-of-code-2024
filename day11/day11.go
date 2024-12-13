package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var rules = []func(string) []string{
	func(s string) []string {
		if s == "0" {
			return []string{"1"}
		}
		return nil
	},
	func(s string) []string {
		re := regexp.MustCompile(`^0+(\d)`)
		if len(s)%2 == 0 {
			left := s[:len(s)/2]
			right := re.ReplaceAllString(s[len(s)/2:], "$1")
			return []string{left, right}
		}
		return nil
	},
	func(s string) []string {
		i, _ := strconv.Atoi(s)
		new := strconv.Itoa(i * 2024)
		return []string{new}
	}}

func main() {
	blinks := flag.Int("blinks", 1, "Number of blinks")
	debug := flag.Bool("debug", false, "Output extra debug info")
	flag.Parse()
	input := "input.txt"
	if len(flag.Args()) > 0 {
		input = flag.Args()[0]
	}
	line := readWords(input)
	if *debug {
		fmt.Println(line)
	}

	pebbles := make(map[string]int)
	for _, word := range line {
		pebbles[word] += 1
	}
	for b := 0; b < *blinks; b++ {
		updates := make(map[string]int)
		for pebble, count := range pebbles {
			for _, f := range rules {
				if result := f(pebble); result == nil {
					continue
				} else {
					updates[pebble] -= count
					updates[result[0]] += count
					if len(result) > 1 {
						updates[result[1]] += count
					}
					break
				}
			}
		}
		for pebble, update := range updates {
			pebbles[pebble] += update
		}
		if *debug {
			sum := 0
			for _, v := range pebbles {
				sum += v
			}
			fmt.Println(b, sum, pebbles)
		}
	}

	sum := 0
	for _, v := range pebbles {
		sum += v
	}
	fmt.Println(sum)

}

func readWords(path string) (words []string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	return
}
