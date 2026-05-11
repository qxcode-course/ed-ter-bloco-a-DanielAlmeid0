package main
import (
    "bufio"
    "fmt"
    "os"
)

func subsetSum(nums []int, target int, i, current int) bool {
   if i >= len(nums) {
       return current == target
   }

   if current > target {
        return false
   }

   if subsetSum(nums, target, i+1, current+nums[i]) {
       return true
   } else {
       return subsetSum(nums, target, i+1, current)
   }
}
func main() {
    reader := bufio.NewReader(os.Stdin)
    var n, k int
    fmt.Fscan(reader, &n, &k)
    nums := make([]int, n)
    for i := range nums {
        fmt.Fscan(reader, &nums[i])
    }
    if subsetSum(nums, k, 0, 0) {
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }
}
