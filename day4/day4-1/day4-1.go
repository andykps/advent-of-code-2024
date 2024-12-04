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

	data := [][]byte{}

	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Bytes()
		// create a copy of the line here because Scan() can overwrite the bufio buffer of which my slice is just a view?
		// I guess I could also have called `scanner.Text()`
		tmp := make([]byte, len(line))
		copy(tmp, line)
		data = append(data, tmp)
	}
	// fmt.Println(len(data), len(data[0]))

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

func count(arr []byte) (ret int) {
	str := string(arr)
	// fmt.Println(str)
	ret += strings.Count(str, "XMAS")
	ret += strings.Count(str, "SAMX")
	return
}
