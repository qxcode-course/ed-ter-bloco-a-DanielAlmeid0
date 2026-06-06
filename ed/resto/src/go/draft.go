package main
import "fmt"

func divisaoEresto(n int) {
    if n == 0 {
        return
    }

    quociente := n / 2
    resto := n % 2

    divisaoEresto(quociente)

    fmt.Printf("%d %d\n", quociente, resto)
}
func main() {
    var n int

    _, err := fmt.Scan(&n)

    if err == nil {
        divisaoEresto(n)
    }
}