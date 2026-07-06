package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func burnTrees(grid [][]rune, l, c int) {
	if l < 0 || l >= len(grid) || c < 0 || c >= len(grid[0]) {
		return
	}

	if grid[l][c] != '#' {
		return
	}

	stack := NewStack[Pos]()
	stack.Push(Pos{l, c})

	direcoes := []Pos{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for !stack.IsEmpty() {
		curr := stack.Pop()

		if grid[curr.l][curr.c] == '#' {
			grid[curr.l][curr.c]= 'o'

			for _, d := range direcoes {
				nl := curr.l + d.l
				nc := curr.c + d.c

				if nl >= 0 && nl < len(grid) && nc >= 0 && nc < len(grid[0]) && grid[nl][nc] == '#' {
					stack.Push(Pos{nl, nc})
				}
			}
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	burnTrees(grid, lfire, cfire)
	showGrid(grid)
}

func showGrid(mat [][]rune) {
	for _, linha := range mat {
		fmt.Println(string(linha))
	}
}
