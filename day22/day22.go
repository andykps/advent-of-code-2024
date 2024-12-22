package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
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

	total := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		sec, _ := strconv.Atoi(scanner.Text())

		for i := 0; i < 2000; i++ {
			sec = prune(mix(sec, sec * 64))
			sec = prune(mix(sec, sec/32))
			sec = prune(mix(sec, sec*2048))
		}
		total += sec
	}
	fmt.Println(total)
}

func mix(sec int, n int) int {
	return n ^ sec
}

func prune(sec int) int {
	return sec % 16777216
}