package main
import "fmt"

func calculaPontos(n int, m int) int {
    if m == 1 {
        return 1
    }

    pontosNovos := (m - 1) * (n - 2) + 1

    return calculaPontos(n, m-1) + pontosNovos
}
func main() {
    var n, m int

    _, err := fmt.Scan(&n, &m)
    if err != nil {
        return
    }

    resultado := calculaPontos(n, m)
    fmt.Println(resultado)
}