package main

import (
	"bufio"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strings"
)

func main() {
	pt2 := flag.Bool("pt2", false, "Perform part 2?")
	debug := flag.Bool("debug", false, "Output debug logging?")
	flag.Parse()

	if *debug {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	path := "input.txt"
	if len(flag.Args()) > 0 {
		path = flag.Args()[0]
	}

	input := readFileOfInts(path)

	blocks := expandToBlocks(input)
	slog.Debug(blocksToString((blocks)))

	if *pt2 {
		defragFiles(blocks)
	} else {
		compactBlocks(blocks)
	}

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
		byte := s.Bytes()[0]
		if byte >= 48 && byte <= 57 {
			ints = append(ints, int(byte-48))
		}
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
			sb.WriteRune(46)
		} else {
			sb.WriteRune(rune(block + 48))
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
		slog.Debug(blocksToString(blocks))
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
	for i := 0; i < len(blocks); i++ {
		if blocks[i] > -1 {
			total += i * blocks[i]
		}
	}
	return total
}

func defragFiles(blocks []int) {
	firstSpace, lastDigit := firstSpaceLastDigit(blocks)
	fileId := -1
	size := 0
	for i := lastDigit; i >= firstSpace; i-- {
		if blocks[i] != fileId {
			if size > 0 {
				if availableSpace := findSpace(blocks, size, i+2); availableSpace > -1 {
					for j := 0; j < size; j++ {
						blocks[availableSpace+j] = blocks[i+1+j]
						blocks[i+1+j] = -1
					}
					slog.Debug(blocksToString(blocks))
				}
			}
			fileId = blocks[i]
			size = 0
		}
		if fileId != -1 {
			size += 1
		}
	}
	slog.Debug(blocksToString(blocks))
}

func findSpace(blocks []int, space int, limit int) (index int) {
	index = -1
	for i := 0; i < limit; i++ {
		if blocks[i] == -1 && index == -1 {
			index = i
		} else if index > -1 && i-index >= space {
			return
		} else if index > -1 && blocks[i] > -1 {
			index = -1
		}
	}
	return -1
}
