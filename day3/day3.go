package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	pt2 := flag.Bool("pt2", false, "Is 'Part 2' enabled?")
	flag.Parse()
	input := "input.txt"
	if len(flag.Args()) > 0 {
		input = flag.Args()[0]
	}

	bytes, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}

	if *pt2 {
		pt2re := regexp.MustCompile(`(?s)don't\(\).*?(do\(\)|$)`)
		bytes = pt2re.ReplaceAll(bytes, []byte(""))
	}

	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	total := 0
	match := re.FindAllSubmatch(bytes, -1)
	for _, element := range match {
		a, _ := strconv.Atoi(string(element[1]))
		b, _ := strconv.Atoi(string(element[2]))
		total += a * b
	}
	fmt.Println(total)
}
