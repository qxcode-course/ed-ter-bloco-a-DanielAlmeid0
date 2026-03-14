package main
import "fmt"
func main() {

    var nome string
    var idade int

    fmt.Scan(&nome, &idade)

    var classificao string

    if idade < 12 {
        classificao = "crianca"

    } else if idade < 18 {
        classificao = "jovem"

    } else if idade < 65 {
        classificao = "adulto"

    } else if idade < 1000 {
        classificao = "idoso"

    } else {
        classificao = "mumia"
    }

    fmt.Printf("%s eh %s\n", nome, classificao)

}
