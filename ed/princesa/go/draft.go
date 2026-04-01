package main

import (
    "fmt"
)
    
    func josephusMarcacao(n, e int) {
        vivos := make([]bool, n)
        for i := range vivos {
            vivos[i] = true
        }

        idx := e - 1
        vivosCount := n

        for vivosCount > 0 {
            fmt.Print("[ ")
            
            for i := 0; i < n; i++ {
                if vivos[i] {
                    if i == idx {
                        fmt.Printf("%d> ", i+1)
                    } else {
                        fmt.Printf("%d ", i+1)
                    }
                }
            }
            fmt.Println("]")

            if vivosCount == 1 {
                break
            }

            killIdx := (idx + 1) % n
            for !vivos[killIdx] {
                killIdx = (killIdx + 1) % n
            }

            vivos[killIdx] = false
            vivosCount--

            idx = (killIdx + 1) % n
            for !vivos[idx] {
                idx = (idx + 1) % n
            }
        }
    }

    func josephusRemocaoSlice(n, e int) {
        arr := make([]int, n)
        for i := 0; i < n; i++ {
            arr[i] = i + 1
        }

        idx := -1

        for i, val := range arr {
            if val == e {
                idx = i
                break
            }
        }

        for len(arr) > 0 {
            fmt.Print("[ ")
            
            for i, val := range arr {
                if i == idx {
                    fmt.Printf("%d> ", val)
                } else {
                    fmt.Printf("%d ", val)
                }
            }
            fmt.Println("]")

            if len(arr) == 1 {
                break
            }

            killIdx := (idx + 1) % len(arr)
            arr = append(arr[:killIdx], arr[killIdx+1:]...)

            idx = killIdx % len(arr)
        }
    }

    func main() {
        var n, e int
        fmt.Scan(&n, &e)
        josephusRemocaoSlice(n, e)
    }