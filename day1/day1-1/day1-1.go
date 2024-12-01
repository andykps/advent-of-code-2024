package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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
	// dists := []int{}
	total := 0

	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		leftInt, _ := strconv.Atoi(words[0])
		rightInt, _ := strconv.Atoi(words[1])
		left = append(left, leftInt)
		right = append(right, rightInt)
	}

	slices.Sort(left)
	slices.Sort(right)

	for i := 0; i < len(left); i++ {
		dist := left[i] - right[i]
		if dist < 0 {
			dist = -dist
		}
		//	dists = append(dists, dist)
		total += dist
	}

	//	fmt.Println(left, right, dists)
	fmt.Println(total)
}
