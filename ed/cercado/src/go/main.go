package main

import (
	"bufio"
	"fmt"
	"os"
)

// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}

	linhas := len(board)
	colunas := len(board[0])

	var dfs func(l, c int)
	dfs = func(l, c int) {
		if l < 0 || l >= linhas || c < 0 || c >= colunas || board[l][c] != 'O' {
			return
		}

		board[l][c] = 'S'

		dfs(l-1, c)
		dfs(l+1, c)
		dfs(l, c-1)
		dfs(l, c+1)
	}

	for c := 0; c < linhas; c++ {
		if board[0][c] == 'O' {
			dfs(0, c)
		}
		if board[linhas-1][c] == 'O' {
			dfs(linhas-1, c)
		}
	}

	for l := 0; l < linhas; l++ {
		for c := 0; c < colunas; c++ {
			if board[l][c] == 'O' {
				board[l][c] = 'X'
			} else if board[l][c] == 'S' {
				board[l][c] = 'O'
			}
		}
	}
}

// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}
