package main
import "fmt"
func main() {
	var n int 
	if _, err := fmt.Scan(&n); err != nil {
		return
	}

	fila := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&fila[i])
	}

	var m int
	fmt.Scan(&m)

	saiu := make(map[int]bool)
	for i := 0; i < m; i++ {
		var id int 
		fmt.Scan(&id)
		saiu[id] = true

	}

	for _, pessoa := range fila {
		if !saiu[pessoa] {
			fmt.Printf("%d ", pessoa)
		}
	}

	fmt.Println()
}
