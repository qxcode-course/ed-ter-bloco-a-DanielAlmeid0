 package main
import "fmt"

func printSufixos(s string) {
    if len(s) == 0 {
        return
    }

    printSufixos(s[1:])
    fmt.Println(s)
}
func main() {
    var palavra string

    _, err := fmt.Scan(&palavra)
    if err != nil {
        return
    }

    printSufixos(palavra)
}