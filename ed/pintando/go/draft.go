package main
import "fmt" 
import "math"

func main() {
    

    var a, b, c float64


    fmt.Scan(&a, &b, &c)

    s := (a + b + c) / 2.0


    area := math.Sqrt(s * (s - a) * (s - b) * (s - c))
    fmt.Printf("%.2f\n", area)

}
