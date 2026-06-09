package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não modifique a assinatura da função numIslands
// Ela é a função que será chamada no LeetCode para resolver o problema
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	linhas := len(grid)
	colunas := len(grid[0])
	contadorIlhas := 0

	var afundarIlha func(l, c int)
	afundarIlha = func(l, c int) {

	if l < 0 || l >= linhas || c < 0 || c >= colunas || grid[l][c] == '0'{
		return
	}

	grid[l][c] = '0'

	afundarIlha(l-1, c)
	afundarIlha(l+1, c)
	afundarIlha(l, c-1)
	afundarIlha(l, c+1)
	}

	for i := 0; i < linhas; i++ {
		for j := 0; j < colunas; j++ {
			if grid[i][j] == '1'{
				contadorIlhas++
				afundarIlha(i, j)
			}
		}
	}
	return contadorIlhas
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	grid := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		grid[i] = []byte(scanner.Text())
	}
	result := numIslands(grid)
	fmt.Println(result)
}
