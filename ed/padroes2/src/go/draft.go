 package main
import "fmt"

func calculaPecas(n int) int {
    if n == 1 {
        return 3
    }

    return calculaPecas(n-1) + (2*n + 1)
}


func main() {
    var n int

    _ , err := fmt.Scan(&n)
    if err != nil {
        return
    }

    resultado := calculaPecas(n)
    fmt.Println(resultado)
}