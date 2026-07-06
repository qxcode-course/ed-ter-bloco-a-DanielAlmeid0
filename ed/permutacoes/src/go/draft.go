package main

import (
    "bufio"
	"fmt"
	"os"
    "sort"
)

func backtrack(path []rune, runes []rune, visited []bool) {
    if len(path) == len(runes) {
        fmt.Println(string(path))
        return
    }

    for i := 0; i < len(runes); i++ {
        if !visited[i] {
            visited[i] = true
            path = append(path, runes[i])

            backtrack(path, runes, visited)

            path = path[:len(path)-1]
            visited[i] = false 
        }
    }
}
func main() {
    scanner := bufio.NewScanner(os.Stdin)

    if scanner.Scan() {
        text := scanner.Text()

        runes := []rune(text)
        

        sort.Slice(runes, func(i, j int) bool {
            return runes[i] < runes[j]
        })

        visited := make([]bool, len(runes))

        backtrack([]rune{}, runes, visited)
    }
}