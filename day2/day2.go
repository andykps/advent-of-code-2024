package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

func isSafe(levels []int) bool {
	for i := 1; i < len(levels); i++ {
		step := levels[i] - levels[i-1]
		if step < -3 || step > 3 || step == 0 {
			return false
		}
		if i > 1 {
			prevstep := levels[i-1] - levels[i-2]
			// if we have a -ve result then the 2 steps have opposite sign
			// need to convert to float64 or fails on results < 1
			if float64(prevstep)/float64(step) < 0.0 {
				return false
			}
		}
	}
	return true
}

func main() {
	dampenerEnabled := flag.Bool("dampener", false, "Is the dampener enabled?")
	debug := flag.Bool("debug", false, "Output debug logging?")
	flag.Parse()

	if *debug {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	input := "input.txt"
	if len(flag.Args()) > 0 {
		input = flag.Args()[0]
	}

	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	total := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		levels := make([]int, len(fields))
		for i := 0; i < len(fields); i++ {
			levels[i], _ = strconv.Atoi(fields[i])
		}

		safe := isSafe(levels)
		if !safe && *dampenerEnabled {
			for i := 0; i < len(levels) && !safe; i++ {
				newLine := []int{}
				newLine = append(newLine, levels[:i]...)
				newLine = append(newLine, levels[i+1:]...)
				safe = isSafe(newLine)
			}
		}
		if safe {
			total += 1
		}

	}
	fmt.Println(total)
}
