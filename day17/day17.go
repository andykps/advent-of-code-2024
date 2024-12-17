package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var registers = make(map[string]int)
var program = []int{}

var registerRe = regexp.MustCompile(`Register ([A-Z]): (\d+)`)
var programRe = regexp.MustCompile(`Program: (.*)`)

func main() {
	debug := flag.Bool("debug", false, "Output extra debug info")
	flag.Parse()
	input := "input.txt"
	if len(flag.Args()) > 0 {
		input = flag.Args()[0]
	}

	readProgramFromFile(input)
	if *debug {
		fmt.Println("Initial state:")
		printState()
		fmt.Println(">>> Running program", join(program, byte(',')))
	}

	// pt 1
	output := runProgram()
	fmt.Println("Part 1:", join(output, ','))

	if *debug {
		printState()
	}

	// pt 2
	digit := 1
	a := 1
	for digit <= len(program) {
		registers["A"] = a
		registers["B"] = 0
		registers["C"] = 0
		output := runProgram()

		if reflect.DeepEqual(output[len(output)-digit:], program[len(program)-digit:]) {
			if len(output) == len(program) {
				break
			}
			a *= 8
			digit += 1
			continue
		}
		a += 1
	}
	fmt.Println("Part 2", a)
}

func runProgram() (output []int) {
	for i := 0; i < len(program); i += 2 {
		instruction := program[i]
		operand := program[i+1]

		switch instruction {
		case 0:
			denominator := pow(2, combo(operand))
			registers["A"] = registers["A"] / denominator
		case 1:
			registers["B"] = registers["B"] ^ operand
		case 2:
			registers["B"] = combo(operand) % 8
		case 3:
			if registers["A"] != 0 {
				i = operand - 2
			}
		case 4:
			registers["B"] = registers["B"] ^ registers["C"]
		case 5:
			output = append(output, combo(operand)%8)
		case 6:
			denominator := pow(2, combo(operand))
			registers["B"] = registers["A"] / denominator
		case 7:
			denominator := pow(2, combo(operand))
			registers["C"] = registers["A"] / denominator
		default:
			log.Panic("Unknown opcode: ", instruction)
		}
	}
	return
}

func pow(base int, power int) int {
	if power == 0 {
		return 1
	}
	result := base
	for i := 1; i < power; i++ {
		result *= base
	}
	return result
}

func combo(op int) int {
	switch op {
	case 0, 1, 2, 3:
		return op
	case 4:
		return registers["A"]
	case 5:
		return registers["B"]
	case 6:
		return registers["C"]
	default:
		log.Panic("Unexpected operand: ", op)
	}
	return 0 // never gets here if panic??
}

func printState() {
	fmt.Println("Register A:", registers["A"])
	fmt.Println("Register B:", registers["B"])
	fmt.Println("Register C:", registers["C"])
}

func join(parts []int, sep byte) string {
	var sb strings.Builder
	for i, p := range parts {
		if i != 0 {
			sb.WriteByte(sep)
		}
		sb.WriteString(strconv.Itoa(p))
	}
	return sb.String()
}

// yay! not readGridFromFile
func readProgramFromFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if sub := registerRe.FindStringSubmatch(line); sub != nil {
			registers[sub[1]], _ = strconv.Atoi(sub[2])
		} else if sub := programRe.FindStringSubmatch(line); sub != nil {
			sp := strings.Split(sub[1], ",")
			for _, tok := range sp {
				i, _ := strconv.Atoi(tok)
				program = append(program, i)
			}

		}
	}

}
