package main

import (
	"bufio"
	"fmt"
	"os"
)


type Pos struct {
	l, c int
}

func encontraPosicoes(grid [][]rune) (Pos, Pos) {
	var inicio, fim Pos 
	for i, row := range grid {
		for j, ch := range row {
			if ch == 'I' {
				inicio = Pos{i, j}
			} else if ch == 'F' {
				fim = Pos{i, j}
			}
		}
	}
	return inicio, fim
}

func escapeMaze(grid [][]rune, inicio, fim Pos) {
	nl := len(grid)
	nc := len(grid[0])

	visited := make([][]bool, nl)
	for i := 0; i < nl; i++ {
		visited[i] = make([]bool, nc)
	}

	caminho := NewStack[Pos]()
	becos := NewStack[Pos]()


	direcoes := []Pos{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	caminho.Push(inicio)

	for !caminho.IsEmpty() {
		atual := caminho.Top()
		visited[atual.l][atual.c] = true

		if atual == fim {
			break
		}

		var validos []Pos

		for _, d := range direcoes {
			nlViz := atual.l + d.l
			ncViz := atual.c + d.c

			if nlViz >= 0 && nlViz < nl && ncViz >= 0 && ncViz < nc {
				if grid[nlViz][ncViz] != '#' && !visited[nlViz][ncViz] {
					validos = append(validos, Pos{nlViz, ncViz})
				}
			}
		}

		if len(validos) > 0 {
			caminho.Push(validos[0])
		} else  {
			becos.Push(atual)
			caminho.Pop()
		}
	}

	for !caminho.IsEmpty() {
		p := caminho.Pop()
		grid[p.l][p.c] = '.'
	}
}

func printGrid(grid [][]rune) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan(){
		return
	}

	var nl, nc int
	fmt.Sscanf(scanner.Text(), "%d %d", &nl, &nc)

	grid := make([][]rune, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		grid[i] = []rune(scanner.Text())
	}

	inicio, fim := encontraPosicoes(grid)

	escapeMaze(grid, inicio, fim)

	printGrid(grid)
}