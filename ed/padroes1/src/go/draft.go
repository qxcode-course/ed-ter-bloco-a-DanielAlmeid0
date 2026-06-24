package main
import "fmt"

func calculaBlocos(n int) int {
    if n == 1 {
        return 20
    }

    return calculaBlocos(n-1) + 8
}
func main() {
    var n int

    _, err := fmt.Scan(&n)
    if err != nil {
        return
    }

    resultado := calculaBlocos(n)

    fmt.Println(resultado)
}