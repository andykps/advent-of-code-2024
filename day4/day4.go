package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
)

var re = regexp.MustCompile(`XMAS`)

func main() {
	flag.Parse()
	input := "input.txt"
	if len(flag.Args()) > 0 {
		input = flag.Args()[0]
	}

	data := [][]byte{}

	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data = append(data, scanner.Bytes())
	}
	fmt.Println(len(data), len(data[0]))

	total := 0

	// horizontal
	for i := 0; i < len(data); i++ {
		total += count(data[i])
	}

	// vertical
	for x := 0; x < len(data[0]); x++ { // assumes all same length
		col := []byte{}
		for y := 0; y < len(data); y++ {
			col = append(col, data[y][x])
		}
		total += count(col)
	}

	// diag
	// totally stolen from https://stackoverflow.com/questions/20420065/loop-diagonally-through-two-dimensional-array#answer-20422854
	for k := 0; k < len(data)*2; k++ {
		col := []byte{}
		for j := 0; j <= k; j++ {
			i := k - j
			if i < len(data) && j < len(data) {
				col = append(col, data[i][j])
			}
		}
		// fmt.Println(string(col))
		total += count(col)
	}
	for k := 0; k < len(data)*2; k++ {
		col := []byte{}
		for j := 0; j <= k; j++ {
			i := k - j
			if i < len(data) && j < len(data) {
				col = append(col, data[len(data)-i-1][j])
			}
		}
		// fmt.Println(string(col))
		total += count(col)
	}

	fmt.Println(total)
}

func reverse(arr []byte) (ret []byte) {
	size := len(arr)
	for i := 0; i < size; i++ {
		ret = append(ret, arr[size-i-1])
	}
	return
}

func count(arr []byte) (ret int) {
	matches := re.FindAll(arr, -1)
	ret += len(matches)
	matches = re.FindAll(reverse(arr), -1)
	ret += len(matches)
	return
}
