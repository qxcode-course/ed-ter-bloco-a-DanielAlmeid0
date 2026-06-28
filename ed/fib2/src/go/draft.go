package main
import "fmt"

func f(n int) int {
    if n == 1 || n == 2 {
        return 1
    }

    if n == 3 {
        return 2
    }

    return f(n-2) + f(n-3)
}
func main() {
    var n int

    _, err := fmt.Scan(&n)
    if err != nil {
        return
    }

    resultado := f(n)
    fmt.Println(resultado)
}