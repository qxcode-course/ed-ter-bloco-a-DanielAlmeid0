package main
import "fmt"

func calculaCoelhos(n int, k int) int {
    if n == 1 || n == 2 {
        return 1
    }

    return calculaCoelhos(n-1, k) + k*calculaCoelhos(n-2, k)
}
func main() {
    var n, k int

    _, err := fmt.Scan(&n, &k)
    if err != nil {
        return
    }

    resultado := calculaCoelhos(n, k)
    fmt.Println(resultado)
}