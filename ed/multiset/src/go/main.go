 package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


type MultiSet struct {
	data  []int
	size  int
	capacity int 
}


func NewMultiSet(capacity int) *MultiSet {
	return &MultiSet{
		data: make([]int, capacity),
		size: 0,
		capacity: capacity,
	}
}

func (ms *MultiSet) expand() {
	if ms.capacity == 0 {
		ms.capacity = 1
	} else {
		ms.capacity *= 2
	}
	newData := make([]int, ms.capacity)
	copy(newData, ms.data[:ms.size])
	ms.data = newData
}

func (ms *MultiSet) search(value int) (bool, int) {
	left, right := 0, ms.size
	for left < right {
		mid := left + (right-left)/2

		if ms.data[mid] <= value {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if left > 0 && ms.data[left-1] == value {
		return true, left - 1
	}
	return false, left
}

func (ms *MultiSet) Insert(value int) {
	if ms.size == ms.capacity {
		ms.expand()
	}
	_, idx := ms.search(value)

	for i := ms.size; i > idx; i-- {
		ms.data[i] = ms.data[i-1]
	}
	ms.data[idx] = value
	ms.size++
}

func (ms *MultiSet) Erase(value int) error {
	found, idx := ms.search(value)
	
	if !found {
		return fmt.Errorf("value not found")
	}

	for i := idx; i < ms.size-1; i++ {
		ms.data[i] = ms.data[i+1]
	}
	ms.size--
	return nil
}

func (ms *MultiSet) Count(value int) int {
	found, idx := ms.search(value)
	if !found {
		return 0
	}

	count := 0
	for i := idx; i >= 0 && ms.data[i] == value; i-- {
		count++
	}
	return count
}

func (ms *MultiSet) Unique() int {
	if ms.size == 0 {
		return 0
	}
	count := 1
	for i := 1; i < ms.size; i++{
		if ms.data[i] != ms.data[i-1] {
			count++
		}
	}
	return count
}

func (ms *MultiSet) Clear() {
	ms.size = 0
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func (ms *MultiSet) Contains(value int) bool {
    found, _ := ms.search(value)
    return found
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	var ms *MultiSet

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			 value, _ := strconv.Atoi(args[1])
			 ms = NewMultiSet(value)
		case "insert":
			 for _, part := range args[1:] {
			 	value, _ := strconv.Atoi(part)
				ms.Insert(value)
			 }
		case "show":
			fmt.Printf("[%s]\n", Join(ms.data[:ms.size], ", "))
		case "erase":
			 value, _ := strconv.Atoi(args[1])
			 err := ms.Erase(value)
			 if err != nil {
				fmt.Println(err.Error())
			 }
		case "contains":
		 value, _ := strconv.Atoi(args[1])
		 fmt.Println(ms.Contains(value))
		case "count":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Count(value))
		case "unique":
			fmt.Println(ms.Unique())
		case "clear":
			ms.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
