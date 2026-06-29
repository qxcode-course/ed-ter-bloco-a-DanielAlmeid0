package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MagicSearch(slice []int, value int) int {
	esquerda := 0
	direita := len(slice) - 1
	foundIndex := -1

	for esquerda <= direita {
		meio := esquerda + (direita-esquerda)/2

		if slice[meio] == value {
			foundIndex = meio
			break
		} else if slice[meio] < value {
			esquerda = meio + 1
		} else {
			direita = meio - 1
		}
	}


	if foundIndex != -1 {
		ultimo := foundIndex
		for ultimo+1 < len(slice) && slice[ultimo+1] == value {
			ultimo++
		}
		return ultimo
	}
	
	return esquerda
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	slice := make([]int, 0, 1)

	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}

	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())

	result := MagicSearch(slice, value)
	fmt.Println(result)
}
