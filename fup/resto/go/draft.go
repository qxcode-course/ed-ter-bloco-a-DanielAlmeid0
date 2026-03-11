package main
import "fmt"

func main() {

    var n, d int

    fmt.Scan(&n, &d)

    quociente := n / d
    resto := n % d

    fmt.Println(quociente, resto);
}
