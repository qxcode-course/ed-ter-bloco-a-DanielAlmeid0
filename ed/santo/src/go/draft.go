package main
import "fmt"

func milagre(n int, c float64) float64 {
   if n == 0 {
    return 0.0
   }

   dinheiroSaida := milagre(n-1, c)

   dinheiroEntrada := (dinheiroSaida + c) / 2.0

   return dinheiroEntrada
}
func main() {
    var n int
    var c float64

    _, err := fmt.Scan(&n, &c)
    if err != nil {
        return
    }

    resultado := milagre(n, c)

    fmt.Printf("%.2f\n", resultado)
}