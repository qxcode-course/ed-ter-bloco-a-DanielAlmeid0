package main
import "fmt"

func calculaQuadrado(n int) int {
    if n == 1 {
        fmt.Println("1^2 = 1")
        return 1
    }

    fmt.Printf("%d^2 = %d^2 + 2*%d + 1 = ?\n", n, n-1, n-1)

    valorAnterior := calculaQuadrado(n-1)

    resultadoAtual := valorAnterior + 2*(n-1) + 1

    fmt.Printf("%d^2 = %d^2 + 2*%d + 1 = %d\n", n, n-1, n-1, resultadoAtual)

    return resultadoAtual
}
func main() {
    var n int 

    _, err := fmt.Scan(&n)
    if err != nil {
        return
    }

    calculaQuadrado(n)
}