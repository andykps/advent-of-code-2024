package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
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

	total := 0

	for x := 1; x < len(data)-1; x++ {
		for y := 1; y < len(data)-1; y++ {
			// fmt.Println(x, y)
			if data[y][x] == 65 {
				// fmt.Println("Is A")
				if checkMas(data, x, y) {
					// fmt.Println("Is MAS")
					total += 1
				}
			}
		}
	}

	fmt.Println(total)
}

func checkMas(data [][]byte, x int, y int) (isMas bool) {
	tl := data[y-1][x-1]
	tr := data[y-1][x+1]
	bl := data[y+1][x-1]
	br := data[y+1][x+1]
	chars := string([]byte{tl, tr, bl, br})
	// fmt.Println(chars)

	isMas = chars == "MMSS" || chars == "MSMS" || chars == "SMSM" || chars == "SSMM"
	return
}
