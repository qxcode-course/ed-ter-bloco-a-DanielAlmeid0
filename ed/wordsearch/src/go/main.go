package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(board [][]byte, word string) bool {
	rows := len(board)
	cols := len(board[0])

	var dfs func(r, c, index int) bool
	dfs = func(r, c, index int) bool {

		if index == len(word) {
			return true
		}

		if r < 0 || r >= rows || c < 0 || c >= cols || board[r][c] != word[index] {
			return false
		}

		temp := board[r][c]

		board[r][c] = '#'

		found := dfs(r-1, c, index+1) || 
				 dfs(r+1, c, index+1) ||
				 dfs(r, c-1, index+1)   ||
				 dfs(r, c+1, index+1)

		board[r][c] = temp

		return found
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == word[0] {
				if dfs(i, j, 0) {
					return true
				}
			}
		}
	}

	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
