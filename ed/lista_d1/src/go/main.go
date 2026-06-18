package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	next *Node
	prev *Node
}
type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	root := &Node{}
	root.next = root
	root.prev = root
	return &LList{
		root: root,
		size: 0,
	}
}

func (ll *LList) Size() int {
	return ll.size
}

func (ll *LList) Clear() {
	ll.root.next = ll.root
	ll.root.prev = ll.root
	ll.size = 0
}

func (ll *LList) insertBefore(at *Node, value int) {
	newNode := &Node{
		Value: value,
		prev: at.prev,
		next: at,
	}
	at.prev.next = newNode
	at.prev = newNode
	ll.size++
}

func (ll *LList) PushFront(value int) {
	ll.insertBefore(ll.root.next, value)
}

func (ll *LList) PushBack(value int) {
	ll.insertBefore(ll.root, value)
}

func (ll *LList) remove(node *Node) {
	if node == ll.root {
		return
	}
	node.prev.next = node.next
	node.next.prev = node.prev
	ll.size--
}

func (ll *LList) PopFront() {
	if ll.size > 0 {
		ll.remove(ll.root.next)
	}
}

func (ll *LList) PopBack() {
	if ll.size > 0 {
		ll.remove(ll.root.prev)
	}
}

func (ll *LList) String() string {
	if ll.size == 0 {
		return "[]"
	}

	var b strings.Builder
	b.WriteString("[")

	curr := ll.root.next
	b.WriteString(strconv.Itoa(curr.Value))

	curr = curr.next
	for curr != ll.root {
		b.WriteString(", ")
		b.WriteString(strconv.Itoa(curr.Value))
		curr = curr.next
	}

	b.WriteString("]")
	return b.String()
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
 	ll := NewLList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			 fmt.Println(ll.String())
		case "size":
			 fmt.Println(ll.Size())
		case "push_back":
			 for _, v := range args[1:] {
			 	num, _ := strconv.Atoi(v)
			 	ll.PushBack(num)
			 }
		case "push_front":
			 for _, v := range args[1:] {
			 	num, _ := strconv.Atoi(v)
			 	ll.PushFront(num)
			 }
		case "pop_back":
			 ll.PopBack()
		case "pop_front":
			 ll.PopFront()
		case "clear":
			ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
