package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func maximumDetonation(bombs [][]int) int {
    n := len(bombs)
    adj := make([][]int, n)

    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if i == j {
                continue
            }

            dx := bombs[i][0] - bombs[j][0]
            dy := bombs[i][1] - bombs[j][1]
            r := bombs[i][2]

            if dx*dx + dy*dy <= r*r {
                adj[i] = append(adj[i], j)
            }
        }
    }

    maxDetonated := 0

    for i := 0; i < n; i++ {
        visited := make([]bool, n)
        count := 0

        var dfs func(node int)
        dfs = func(node int) {
            visited[node] = true
            count++

            for _, neighbor := range adj[node] {
                if !visited[neighbor] {
                    dfs(neighbor)
                }
            }
        }
        dfs(i)

        if count > maxDetonated {
            maxDetonated = count
        }
    }
    return maxDetonated
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

    n, _ := strconv.Atoi(parts[0])

    bombs := make([][]int, n)
    for i := 0; i < n; i++ {
        scanner.Scan()
        line := strings.Fields(scanner.Text())
        b := make([]int, 3)
        b[0], _ = strconv.Atoi(line[0])
        b[1], _ = strconv.Atoi(line[1])
        b[2], _ = strconv.Atoi(line[2])
        bombs[i] = b
    }

    fmt.Println(maximumDetonation(bombs))
}