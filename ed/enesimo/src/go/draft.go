package main
import "fmt"

func eh_primo(x int, div int) bool {
    if x <= 1 {
        return false
    }

    if div == x {
        return true
    }

    if x%div == 0 {
        return false
    }

    return eh_primo(x, div+1)
}

func busca_primo(alvo int, atual int, contador int) int {

    if eh_primo(atual, 2) {
        if contador == alvo {
            return atual
        }

        return busca_primo(alvo, atual+1, contador+1)
    }

    return busca_primo(alvo, atual+1, contador)
}
func main() {
    var n int 

    _, err := fmt.Scan(&n)
    if err != nil {
        return
    }

    resultado := busca_primo(n, 2, 1)

    fmt.Println(resultado)
}