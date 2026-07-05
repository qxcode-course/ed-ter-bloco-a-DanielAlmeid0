package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
    r, c int
}

func orangeRotting(grid [][]int) int {
    if len(grid) == 0 {
        return 0
    }

    rows := len(grid)
    cols := len(grid[0])

    queue := []Pos{}
    freshCount := 0

    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if grid[r][c] == 2 {
                queue = append(queue, Pos{r, c})
            } else if grid[r][c] == 1 {
                freshCount++
            }
        }
    }

    if freshCount == 0 {
        return 0
    }

    minutes := 0
    directions := []Pos{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

    for len(queue) > 0 && freshCount > 0 {
        size := len(queue)

        for i := 0; i < size; i++ {
            size := len(queue)

            for i := 0; i < size; i++ {
                curr := queue[0]
                queue = queue[1:]

                for _, d := range directions {
                    nr, nc := curr.r+d.r, curr.c+d.c

                    if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == 1 {
                        grid[nr][nc] = 2 
                        freshCount--
                        queue = append(queue, Pos{nr, nc})
                    }
                }
            }
            minutes++
        }
    }

    if freshCount > 0 {
        return -1
    }

    return minutes
}
func main() {
    scanner := bufio.NewScanner(os.Stdin)
    if !scanner.Scan() {
        return
    }

    parts := strings.Fields(scanner.Text())
    if len(parts) < 2 {
        return
    }

    rows, _ := strconv.Atoi(parts[0])
    cols, _ := strconv.Atoi(parts[1])

    grid := make([][]int, rows)

    for i:= 0; i < rows; i++ {
        scanner.Scan()
        line := strings.Fields(scanner.Text())
        grid[i] = make([]int, cols)

        for j := 0; j < cols; j++ {
            grid[i][j], _= strconv.Atoi(line[j])
        }
    }

    fmt.Println(orangeRotting(grid))
}