package main

import (
	"bufio"
	"fmt"
	"os"
)

type State struct {
    text []rune
    cursor int
}
func main() {
    scanner := bufio.NewScanner(os.Stdin)
    if !scanner.Scan() {
        return
    }
    input := scanner.Text()

    var text []rune
    cursor := 0
    var undoStack []State
    var redoStack []State

    saveState := func() {

        tCopy := make([]rune, len(text))
        copy(tCopy, text)

        undoStack = append(undoStack, State{text: tCopy, cursor: cursor})
        redoStack = []State{}
    }

    for _, ch := range input {
        switch ch {
        case 'R': //O enter
            saveState()
            text = append(text, 0)
            copy(text[cursor+1:], text[cursor:])
            text[cursor] = '\n'
            cursor++

        case 'B':
            if cursor > 0 {
                saveState()
                text = append(text[:cursor-1], text[cursor:]...)
                cursor--
            }

        case 'D':
            if cursor < len(text) {
                saveState()

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
        
        case 'Z':
            if len(undoStack) > 0 {

                tCopy := make([]rune, len(text))
                copy(tCopy, text)
                redoStack = append(redoStack, State{text: tCopy, cursor: cursor})

                last := undoStack[len(undoStack)-1]
                undoStack = undoStack[:len(undoStack)-1]

                text = make([]rune, len(last.text))
                copy(text, last.text)
                cursor = last.cursor
            }

        case 'Y':
            if len(redoStack) > 0 {
               tCopy := make([]rune, len(text))
               copy(tCopy, text)
               undoStack = append(undoStack, State{text: tCopy, cursor: cursor})

               last := redoStack[len(redoStack)-1]
               redoStack = redoStack[:len(redoStack)-1]

               text = make([]rune, len(last.text))
               copy(text, last.text)
               cursor = last.cursor
            }
        default:
            saveState()
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