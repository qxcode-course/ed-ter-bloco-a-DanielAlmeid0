package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	rows := len(matrix)
	cols := len(matrix[0])


	memo := make([][]int, rows)
	for i := range memo {
		memo[i] = make([]int, cols)
	}

	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var dfs func(r, c int) int 
	dfs = func(r, c int) int {
		if memo[r][c] != 0 {
			return memo[r][c]
		}

		maxLen := 1

		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]

			if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
				if matrix[nr][nc] > matrix[r][c] {
					length := 1 + dfs(nr, nc)
					if length > maxLen {
						maxLen = length
					}
				}
			}
		}

		memo[r][c] = maxLen
		return maxLen
	}

	globalMax := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			length := dfs(i, j)
			if length > globalMax {
				globalMax = length
			}
		}
	}
	return globalMax
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}
	parts := strings.Fields(scanner.Text())
	if len(parts) < 2 {
		return
	}
	nl, _ := strconv.Atoi(parts[0])
	nc, _ := strconv.Atoi(parts[1])

	matrix := make([][]int, nl)
	for i := 0; i < nl; i++ {
		if !scanner.Scan() {
			return
		}
		tokens := strings.Fields(scanner.Text())
		row := make([]int, nc)
		for j := 0; j < nc && j < len(tokens); j++ {
			v, _ := strconv.Atoi(tokens[j])
			row[j] = v
		}
		matrix[i] = row
	}

	fmt.Println(longestIncreasingPath(matrix))
}
