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
	Left  *Node
	Right *Node
}


func find(node *Node, value int) *Node {
	if node == nil {
		return nil
	}

	if node.Value == value {
		return node
	}

	foundLeft := find(node.Left, value)
	if foundLeft != nil {
		return foundLeft
	}

	return find(node.Right, value)
}


func getHeight(node *Node) int {
	if node == nil {
		return 0
	}

	leftHeight := getHeight(node.Left)
	rightHeight := getHeight(node.Right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}

	return rightHeight + 1 
}


func calcNodeDepth(node *Node, level int, value int) int {
	if node == nil {
		return 0
	}

	if node.Value == value {
		return level
	}

	leftDepth := calcNodeDepth(node.Left, level+1, value)
	if leftDepth != 0 {
		return leftDepth
	}

	return calcNodeDepth(node.Right, level+1, value)
}

// --------------------------------------------------------------------
// Don't change from here
func BShow(node *Node, heranca string) {
	if node != nil && (node.Left != nil || node.Right != nil) {
		BShow(node.Left, heranca+"l")
	}
	for i := 0; i < len(heranca)-1; i++ {
		if heranca[i] != heranca[i+1] {
			fmt.Print("│   ")
		} else {
			fmt.Print("    ")
		}
	}
	if heranca != "" {
		if heranca[len(heranca)-1] == 'l' {
			fmt.Print("╭───")
		} else {
			fmt.Print("╰───")
		}
	}
	if node == nil {
		fmt.Println("#")
		return
	}
	fmt.Println(node.Value)
	if node.Left != nil || node.Right != nil {
		BShow(node.Right, heranca+"r")
	}
}

func create(parts *[]string) *Node {
	elem := (*parts)[0]
	*parts = (*parts)[1:]
	if elem == "#" {
		return nil
	}
	value, _ := strconv.Atoi(elem)
	node := &Node{Value: value}
	node.Left = create(parts)
	node.Right = create(parts)
	return node
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	line := scanner.Text()
	parts := strings.Split(line, " ")
	root := create(&parts)

	scanner.Scan()
	line = scanner.Text()
	fmt.Println("Arvore:")
	BShow(root, "")

	 values := strings.FieldsSeq(line)
	 for s := range values {
	 value, _ := strconv.Atoi(s)
	 node := find(root, value)
	 if node != nil {
	 		fmt.Printf("Altura: %d, Profundidade: %d\n", getHeight(node), calcNodeDepth(root, 1, value))
	 	} else {
	 		fmt.Println("-1")
	 	}
	 }
}
