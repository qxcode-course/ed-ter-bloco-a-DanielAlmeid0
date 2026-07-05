package main

import (
	
	// "container/list"
	"fmt"
)

func main() {
	q := NewQueue[string]()

	for i := 0; i < 16; i++ {
		q.Enqueue(string(rune('A' + i)))
	}

	for i := 0; i < 15; i++ {
		var gols1, gols2 int

		if _, err := fmt.Scan(&gols1, &gols2); err != nil {
			break
		}

		time1 := q.Dequeue()
		time2 := q.Dequeue()

		if gols1 > gols2 {
			q.Enqueue(time1)
		} else {
			q.Enqueue(time2)
		}
	}

	campeao := q.Dequeue()
	fmt.Println(campeao)
}
