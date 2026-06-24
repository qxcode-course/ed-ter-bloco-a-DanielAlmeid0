package main
import "fmt"

func calculaManeiras(n int) int {
    if n == 1 {
        return 1
    }
    if n == 2 {
        return 1
    }
    if n == 3 {
        return 2
    }

    x1, x2, x3 := 1, 1, 2
    var atual int

    for i := 4; i <= n; i++ {
        atual = x3 + x1

        x1 = x2
        x2 = x3
        x3 = atual
    }
    return atual
}
func main() {
    var n int

    _, err := fmt.Scan(&n)
    if err != nil {
        return
    }

    resultado := calculaManeiras(n)
    fmt.Println(resultado)
}