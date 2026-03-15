package main
import "fmt"

func main() {
    var totalalbu, qtbBaruel int
    var contagem [51]int

    fmt.Scan(&totalalbu, &qtbBaruel)

    for i := 0; i < qtbBaruel; i++ {
        var figurinha int
        fmt.Scan(&figurinha)

        if figurinha <= totalalbu {
            contagem[figurinha]++
        }
    }

    repetida := false

    for i := 1; i <= totalalbu; i++ {
        if contagem[i] > 1 {

            sobrou := contagem[i] - 1

            for j := 0; j < sobrou; j++ {
                if repetida {
                    fmt.Print(" ")
                }

                 fmt.Print(i)
            repetida = true
            }
        

        }
    }

    if !repetida {
        fmt.Print("N")
    }
    fmt.Println()

    faltando := false
    for i := 1; i <= totalalbu; i++ {
        if contagem[i] == 0 {
            if faltando {
                fmt.Print(" ")
            }
            fmt.Print(i)
            faltando = true
        }
    }

    if !faltando {
    fmt.Print("N")
}

fmt.Println()

}




