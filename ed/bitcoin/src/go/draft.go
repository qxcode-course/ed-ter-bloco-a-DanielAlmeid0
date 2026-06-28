package main
import "fmt"

func numAtivos(n int, k int) int {
    if n <= k {
        return 1
    }

    metadeCima := (n + 1) / 2
    metadeBaixo := n / 2

    return numAtivos(metadeCima, k) + numAtivos(metadeBaixo, k)
}
func main() {
    var n, k int

    _, err := fmt.Scan(&n, &k)
    if err != nil {
        return
    }

    resultado := numAtivos(n, k)
    fmt.Println(resultado)
}