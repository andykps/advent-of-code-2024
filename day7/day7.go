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

func main() {
	pt2 := flag.Bool("pt2", false, "Run part 2 instead of part 1")
	flag.Parse()
	input := "input.txt"
	if len(flag.Args()) > 0 {
		input = flag.Args()[0]
	}
	symbols := []byte("+*")
	if *pt2 {
		symbols = append(symbols, 124)
	}
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	total := 0
	scanner := bufio.NewScanner(f)
out:
	for scanner.Scan() {
		re := regexp.MustCompile(`:? `)
		line := scanner.Text()
		parts := re.Split(line, -1)
		answer, _ := strconv.Atoi(parts[0])
		operands := []int{}
		for _, op := range parts[1:] {
			iop, _ := strconv.Atoi(op)
			operands = append(operands, iop)
		}
		operatorCount := len(operands) - 1
		perms := permutations(symbols, operatorCount)
		for _, perm := range perms {
			subtotal := operands[0]
			for i, operand := range operands[1:] {
				switch perm[i] {
				case 42: // *
					subtotal *= operand
				case 43: // +
					subtotal += operand
				case 124: // \
					subtotal, _ = strconv.Atoi(strconv.Itoa(subtotal) + strconv.Itoa(operand))
				}
			}
			if subtotal == answer {
				total += answer
				continue out
			}
		}
	}
	fmt.Println(total)
}

func permutations(symbols []byte, length int) (perms [][]byte) {
	if length == 0 {
		return [][]byte{{}}
	}
	for _, sym := range symbols {
		for _, perm := range permutations(symbols, length-1) {
			newPerm := append([]byte{}, perm...)
			newPerm = append(newPerm, sym)
			perms = append(perms, newPerm)
		}
	}
	return
}
