package main
import "fmt"

func calculaCombinacao(n, k int) int {

    if k < 0 || k > n {
        return 0
    }

    if k > n-k {
        k = n - k 
    }

    dp := make([]int, k+1)
    dp[0] = 1

    for i := 1; i <= n; i++ {
        for j := k; j > 0; j-- {
            dp[j] = dp[j] + dp[j-1]
        }
    }
    return dp[k]
}
func main() {
    var n, k int

    _, err := fmt.Scan(&n, &k)
    if err != nil {
        return
    }

    resultado := calculaCombinacao(n, k)
    fmt.Println(resultado)
}