package main
import "fmt"
func main() {
    var quantidade int
    var direcao string

    fmt.Scan(&quantidade, &direcao)

    var x [1000]int
    var y [1000]int

    for i := 0; i < quantidade; i++ {
        fmt.Scan(&x[i], &y[i])

    }

    for i := quantidade - 1; i > 0; i-- {
        x[i] = x[i-1]
        y[i] = y[i-1]
    }

    if direcao == "U" {
        y[0] = y[0] - 1
    } else if direcao == "D" {
        y[0] = y[0] + 1
    } else if direcao == "L" {
        x[0] = x[0] - 1
    } else if direcao == "R" {
        x[0] = x[0] + 1
    }

    for i := 0; i < quantidade; i++ {
        fmt.Println(x[i], y[i])
    
    }
}
