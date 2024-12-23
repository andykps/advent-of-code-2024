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

	connMap := make(map[string][]string)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "-")
		from := parts[0]
		to := parts[1]
		connMap[from] = append(connMap[from], to)
		connMap[to] = append(connMap[to], from)
	}

	conns := [][]string{}
	for from := range connMap {
		conns = append(conns, connections(connMap, from, []string{})...)
	}
	// slices.SortFunc(conns, func(a []string, b []string) int {
	// 	for i := 0; i < 3; i++ {
	// 		c := strings.Compare(a[i], b[i])
	// 		if c == 0 {
	// 			continue
	// 		}
	// 		return c
	// 	}
	// 	return 0
	// })

	deduped := [][]string{}
	for _, c := range conns {
		if !slices.ContainsFunc(deduped, func(c1 []string) bool {
			slices.Sort(c)
			slices.Sort(c1)
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

	fmt.Println(total)

}

func connections(connMap map[string][]string, from string, path []string) [][]string {
	if len(path) > 3 {
		return [][]string{}
	}
	path = append(path, from)
	res := [][]string{}
	for _, neighbour := range connMap[from] {
		if len(path) == 3 && path[0] == neighbour {
			return [][]string{path}
		}
		if !slices.Contains(path, neighbour) {
			res = append(res, connections(connMap, neighbour, path)...)
		}
	}
	return res
}
