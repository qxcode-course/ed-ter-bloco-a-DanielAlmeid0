package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	data []int
	size int
	capacity int
}

func NewSet(capacity int) *Set {
	if capacity <= 0 {
		capacity = 2
	}

	return &Set{
		data:   make([]int, capacity),
		size:   0,
		capacity: capacity,
	}
}

func (s *Set) reserve(newCapacity int) {
	if newCapacity > s.capacity {
		newData := make([]int, newCapacity)
		copy(newData, s.data[:s.size])
		s.data = newData
		s.capacity = newCapacity
	}
}

func (s *Set) binarySearchInsertionPoint(value int) (int, bool) {
	left, right := 0, s.size-1
	for left <= right {
		mid := left + (right-left)/2

		if s.data[mid] == value {
			return mid, true
		} else if s.data[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left, false
}

func (s *Set) Contains(value int) bool {
	_, found := s.binarySearchInsertionPoint(value)
	return found
}

func (s *Set) Insert(value int) {
	idx, found := s.binarySearchInsertionPoint(value)

	if found {
		return
	}

	if s.size == s.capacity {
		s.reserve(s.capacity * 2)
	}

	for i := s.size; i > idx; i-- {
		s.data[i] = s.data[i-1]
	}

	s.data[idx] = value
	s.size++
}

func (s *Set) Erase(value int) bool {
	idx, found := s.binarySearchInsertionPoint(value)
	if !found {
		return false
	}

	for i := idx; i < s.size-1; i++ {
		s.data[i] = s.data[i+1]
	}
	s.size--
	return true
}

func (s *Set) String() string {
	strVals := make([]string, s.size)
	for i := 0; i < s.size; i++ {
		strVals[i] = strconv.Itoa(s.data[i])
	}
	return "[" + strings.Join(strVals, ", ") + "]"
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	var v *Set
	for scanner.Scan() {
		fmt.Print("$")

		line = scanner.Text()

		fmt.Println(line)

		parts := strings.Fields(line)

		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			capacity := 0
			if len(parts) > 1 {
				capacity, _= strconv.Atoi(parts[1])
			}
			v = NewSet(capacity)
		case "insert":
			if v != nil {
				for _, part := range parts[1:] {
					value, _:= strconv.Atoi(part)
					v.Insert(value)
				}
			}
		case "show":
			if v != nil {
				fmt.Println(v.String())
			}
		case "erase":
			if v != nil && len(parts) > 1 {
				value, _:= strconv.Atoi(parts[1])
				if !v.Erase(value) {
					fmt.Println("value not found")
				}
			}
		case "contains":
			if v != nil && len(parts) > 1{
				value, _:= strconv.Atoi(parts[1])
				fmt.Println(v.Contains(value))
			}
		case "clear":
			if v != nil {
				v = NewSet(v.capacity)
			}
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
