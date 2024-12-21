package main

import "fmt"

type point struct{ x, y int }
type vec struct{ dx, dy int }
type path struct {
	prev *path
	point
	dist int
}

var (
	NORTH = vec{0, -1}
	EAST  = vec{1, 0}
	SOUTH = vec{0, 1}
	WEST  = vec{-1, 0}
	dirs  = []vec{NORTH, EAST, SOUTH, WEST}
)

var numpad = [][]byte{
	{'7', '8', '9'},
	{'4', '5', '6'},
	{'1', '2', '3'},
	{'X', '0', 'A'},
}

func main() {
	routes := shortestRoutes(numpad, 'A', '9')
	for _, route := range routes {
		for _, button := range route {
			fmt.Printf("%v ", string(numpad[button.y][button.x]))
		}
		fmt.Println()
	}
}

func shortestRoutes(grid [][]byte, start byte, end byte) [][]point {
	queue := []point{findButton(grid, start)}
	visited := map[point]bool{queue[0]: true}
	parents := map[point][]point{queue[0]: {}}
	distances := map[point]int{queue[0]: 0}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, d := range dirs {
			next := point{curr.x + d.dx, curr.y + d.dy}
			if !visited[next] && isValid(grid, next) {
				visited[next] = true
				queue = append(queue, next)
				distances[next] = distances[curr] + 1
				parents[next] = []point{curr}
			} else if distances[next] == distances[curr] + 1 {
				parents[next] = append(parents[next], curr)
			}
		}
	}

	return findPathDFS(parents, findButton(grid, end))
}

func findPathDFS(parents map[point][]point, goal point) (paths [][]point) {
	if len(parents[goal]) == 0 {
		return [][]point{{goal}}
	}

	for _, parent := range parents[goal] {
		for _, path := range findPathDFS(parents, parent) {
			paths = append(paths, append(path, goal))
		}
	}
	return
}

func isValid(grid [][]byte, p point) bool {
	return p.x >= 0 && p.y >= 0 && p.y < len(grid) && p.x < len(grid[p.y]) && grid[p.y][p.x] != 'X'
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
