package main
import "fmt"

func diagonal(s string, k int) {
    if len(s) == 0 {
        return
    }

    for i := 0; i < k; i ++ {
        fmt.Print(" ")
    }

    fmt.Printf("%c\n", s[0])

    diagonal(s[1:], k+1)

}
func main() {
    var palavra string

    _, err := fmt.Scan(&palavra)
    if err != nil {
        return
    }

    diagonal(palavra, 0)
}