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

	for b := 0; b < *blinks; b++ {
		newline := []string{}
		for _, word := range line {
			for _, f := range rules {
				if result := f(word); result == nil {
					continue
				} else {
					newline = append(newline, result...)
					break
				}
			}
		}
		if *debug {
			fmt.Println(newline)
		}
		line = newline
	}

	fmt.Println(len(line))

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
