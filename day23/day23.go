package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

var connMap = make(map[string][]string)

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

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "-")
		from := parts[0]
		to := parts[1]
		connMap[from] = append(connMap[from], to)
		connMap[to] = append(connMap[to], from)
	}

	// pt1
	conns := [][]string{}
	for from := range connMap {
		conns = append(conns, connections(from, []string{})...)
	}
	deduped := [][]string{}
	for _, c := range conns {
		if !slices.ContainsFunc(deduped, func(c1 []string) bool {
			return slices.Compare(c, c1) == 0
		}) {
			deduped = append(deduped, c)
		}
	}
	total := 0
	for _, c := range deduped {
		for i := 0; i < 3; i++ {
			if strings.HasPrefix(c[i], "t") {
				total += 1
				break
			}
		}
	}

	// pt 2
	allCliques := [][]string{}
	for from := range connMap {
		cliques := createCliques(from)
		for _, c := range cliques {
			if !slices.ContainsFunc(allCliques, func(c1 []string) bool {
				return slices.Compare(c, c1) == 0
			}) {
				allCliques = append(allCliques, c)
			}
		}
	}

	maxlen := 0
	var maxlan []string
	for _, c := range allCliques {
		if len(c) > maxlen {
			maxlen = len(c)
			maxlan = c
		}
	}

	fmt.Println("Part 1:", total)
	fmt.Println("Part 2:", strings.Join(maxlan, ","))

}

func connections(from string, path []string) [][]string {
	if len(path) > 3 {
		return [][]string{}
	}
	path = append(path, from)
	res := [][]string{}
	for _, neighbour := range connMap[from] {
		if len(path) == 3 && path[0] == neighbour {
			slices.Sort(path)
			return [][]string{path}
		}
		if !slices.Contains(path, neighbour) {
			res = append(res, connections(neighbour, path)...)
		}
	}
	return res
}

func createCliques(from string) (cliques [][]string) {
	for _, neighbour := range connMap[from] {
		clique := []string{from, neighbour}
		for node, neighbours := range connMap {
			if node == from {
				continue
			}
			if containsAll(neighbours, clique) {
				clique = append(clique, node)
			}
		}
		slices.Sort(clique)
		cliques = append(cliques, clique)
	}
	return
}

func containsAll(haystack []string, needles []string) (containsAll bool) {
	for _, needle := range needles {
		if !slices.Contains(haystack, needle) {
			return false
		}
	}
	return true
}
