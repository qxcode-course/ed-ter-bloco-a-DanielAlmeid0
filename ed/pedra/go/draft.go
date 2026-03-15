package main

import (
	"fmt"
)
func main() {
    var competidores, A, B int
    indicevencedor := -1
    recorde := 90000

    fmt.Scan(&competidores)


    for i := 0; i < competidores; i++ {
        fmt.Scan(&A, &B)
        
       if A >= 10 && B >= 10 {

        pontuacao := A - B
        if pontuacao < 0 {
            pontuacao = -pontuacao
        }

        if pontuacao < recorde {
            recorde = pontuacao
            indicevencedor = i

        }
       }

    }

    if indicevencedor != -1 {
        fmt.Println(indicevencedor)
    } else {
        fmt.Println("sem ganhador")
    }
}
