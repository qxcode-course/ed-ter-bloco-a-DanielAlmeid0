package main

import (
	"bufio"
	"fmt"
	"os"
)
func main() {
    scanner := bufio.NewScanner(os.Stdin)
    if !scanner.Scan() {
        return
    }
    input := scanner.Text()

    var text []rune
    cursor := 0

    for _, ch := range input {
        switch ch {
        case 'R':
            text = append(text, 0)
            copy(text[cursor+1:], text[cursor:])
            text[cursor] = '\n'
            cursor++

        case 'B':
            if cursor > 0 {
                text = append(text[:cursor-1], text[cursor:]...)
                cursor--
            }

        case 'D':
            if cursor < len(text) {
                text = append(text[:cursor], text[cursor+1:]...)
            }

        case '<':
            if cursor > 0 {
                cursor--
            }

        case '>':
            if cursor < len(text) {
                cursor++
            }

        default:
            text = append(text, 0)
            copy(text[cursor+1:], text[cursor:])
            text[cursor] = ch
            cursor++
        }
    } 

    finalText := make([]rune, 0, len(text)+1)
    finalText = append(finalText, text[:cursor]...)
    finalText = append(finalText, '|')
    finalText = append(finalText, text[cursor:]...)

    fmt.Println(string(finalText))







}