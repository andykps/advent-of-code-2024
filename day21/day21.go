package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

type point struct{ x, y int }

var numpad = [][]byte{
	{'7', '8', '9'},
	{'4', '5', '6'},
	{'1', '2', '3'},
	{'X', '0', 'A'},
}

var dirpad = [][]byte{
	{'X', '^', 'A'},
	{'<', 'v', '>'},
}

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
	total := 0
	for scanner.Scan() {
		code := scanner.Bytes()

		numpadKeys := keysForKeys(numpad, []byte(code))
		// fmt.Println(string(numpadKeys))
		dirKeys1 := keysForKeys(dirpad, numpadKeys)
		// fmt.Println(string(dirKeys1))
		dirKeys2 := keysForKeys(dirpad, dirKeys1)

		// fmt.Println(string(dirKeys2))
		length := len(dirKeys2)
		num := strip(code)
		// fmt.Println(length, num, length*num)
		total += length * num
	}
	fmt.Println(total)

}

func keysForKeys(keypad [][]byte, keys []byte) []byte {
	result := []byte{}

	for i := 0; i < len(keys); i++ {
		var start, end byte
		if i == 0 {
			start = 'A'
			end = keys[i]
		} else {
			start = keys[i-1]
			end = keys[i]
		}
		result = append(result, keysBetweenKeys(keypad, start, end)...)
	}
	return result
}

func keysBetweenKeys(keypad [][]byte, start byte, end byte) []byte {
	keys := []byte{}

	p1 := findButton(keypad, start)
	p2 := findButton(keypad, end)

	dx := p2.x - p1.x
	dy := p2.y - p1.y

	for x := 0; x < abs(dx); x++ {
		if dx < 0 {
			keys = append(keys, '<')
		} else if dx > 0 {
			keys = append(keys, '>')
		}
	}

	for y := 0; y < abs(dy); y++ {
		if dy < 0 {
			keys = append(keys, '^')
		} else if dy > 0 {
			keys = append(keys, 'v')
		}
	}
	order := []byte("<v^>")
	if len(keypad) == 4 {
		// numpad
		if (start == '0' || start == 'A') && (end == '1' || end == '4' || end == '7') {
			order = []byte("^v<>")
		} else if start == '1' || start == '4' || start == '7' {
			order = []byte("><v^")
		}
	} else {
		// dirpad
		if start == '^' || start == 'A' {
			order = []byte("v<^>")
		} else if start == '>' {
			order = []byte("><v^")
		}
	}
	slices.SortFunc(keys, func(a byte, b byte) int {
		return slices.Index(order, a) - slices.Index(order, b)
	})
	keys = append(keys, 'A')
	return keys
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func findButton(grid [][]byte, b byte) point {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == b {
				return point{x, y}
			}
		}
	}
	panic("Button not found")
}

func strip(s []byte) int {
	stripped := []byte{}
	for _, b := range s {
		if b >= '0' && b <= '9' {
			stripped = append(stripped, b)
		}
	}
	i, _ := strconv.Atoi(string(stripped))
	return i
}
