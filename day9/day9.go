package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	path := "input.txt"
	if len(flag.Args()) > 0 {
		path = flag.Args()[0]
	}
	input := readFileOfInts(path)

	blocks := expandToBlocks(input)

	compactBlocks(blocks)

	checksum := calculateChecksum(blocks)

	fmt.Println(checksum)
}

func readFileOfInts(path string) []int {
	ints := []int{}

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanBytes)
	for s.Scan() {
		ints = append(ints, int(s.Bytes()[0]-48))
	}
	return ints
}

func expandToBlocks(input []int) (expanded []int) {
	for i := 0; i < len(input); i += 2 {
		for j := 0; j < input[i]; j++ {
			expanded = append(expanded, i/2)
		}
		for j := 0; i+1 < len(input) && j < input[i+1]; j++ {
			expanded = append(expanded, -1) // using -1 to represent space
		}
	}
	return
}

func blocksToString(input []int) string {
	var sb strings.Builder
	for _, block := range input {
		if block == -1 {
			sb.WriteByte(46)
		} else {
			sb.WriteByte(byte(block + 48))
		}
	}
	return sb.String()
}

func compactBlocks(blocks []int) {
	for {
		if firstSpace, lastDigit := firstSpaceLastDigit(blocks); firstSpace < lastDigit {
			blocks[firstSpace] = blocks[lastDigit]
			blocks[lastDigit] = -1
		} else {
			break
		}
		// fmt.Println(blocksToString(blocks))
	}
}

func firstSpaceLastDigit(blocks []int) (firstSpace int, lastDigit int) {
	firstSpace = -1
	lastDigit = -1
	for i, b := range blocks {
		if b == -1 {
			if firstSpace == -1 {
				firstSpace = i
			}
		} else {
			lastDigit = i
		}
	}
	return firstSpace, lastDigit
}

func calculateChecksum(blocks []int) int {
	total := 0
	for i := 0; i < len(blocks) && blocks[i] > -1; i++ {
		total += i * blocks[i]
	}
	return total
}
