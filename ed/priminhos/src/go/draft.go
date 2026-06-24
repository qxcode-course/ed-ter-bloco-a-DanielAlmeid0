package main

import (
	"fmt"
	"strconv"
    "strings"
)

func eh_primo(x int) bool {
    if x <= 1 {
        return false
    }

    for i := 2; i*i <= x; i++ {
        if x%i == 0 {
            return false
        }
    }
    return true
}
func main() {
    var n int
    _, err := fmt.Scan(&n)
    if err != nil {
        return
    }

    primos := make([]string, 0, n)
    atual := 2

    for len(primos) < n {
        if eh_primo(atual) {
            primos = append(primos, strconv.Itoa(atual))
        }
        atual++
    }

    textoFinal := strings.Join(primos, ", ")

    fmt.Printf("[%s]\n", textoFinal)
}