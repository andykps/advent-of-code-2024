package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	left := []int{}
	right := []int{}
	total := 0

	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		leftInt, _ := strconv.Atoi(words[0])
		rightInt, _ := strconv.Atoi(words[1])
		left = append(left, leftInt)
		right = append(right, rightInt)
	}

	for i := 0; i < len(left); i++ {
		counter := 0
		for j := 0; j < len(right); j++ {
			if left[i] == right[j] {
				counter += 1
			}
		}
		total += counter * left[i]
		//		fmt.Println(counter, total)
	}

	fmt.Println(total)
}
