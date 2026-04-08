package main
import "fmt"

func divisaoRecursiva(n int) {
    
    if n == 0 {
        return
    }

    resultado := n / 2
    resto := n % 2

    divisaoRecursiva(resultado)

    fmt.Printf("%d %d\n", resultado, resto)

}


func main() {
   var numero int
    
    fmt.Scan(&numero)
    
    divisaoRecursiva(numero)
}
