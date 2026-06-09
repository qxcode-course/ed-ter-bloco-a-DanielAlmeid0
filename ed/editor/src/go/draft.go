package main
import ("fmt"
        "bufio"
        "os"
)

type Editor struct {
    left []rune
    right []rune
}

func (e *Editor) processaComando(cmd rune){
    switch cmd {
    case 'R':
        e.left = append(e.left, '\n')

    case 'B':
        if len(e.left) > 0 {
        e.left = e.left[:len(e.left)-1]
    }
    case 'D':
        if len(e.left) > 0 {
            e.right = e.right[:len(e.right)-1]
        }
    case '<':
        if len(e.left) > 0 {
            ch := e.left[len(e.left)-1]
            e.left = e.left[:len(e.left)-1]
            e.right = append(e.right, ch)
        }
    case '>':
        if len(e.right) > 0 {
            ch := e.right[len(e.right)-1]
            e.right = e.right[:len(e.right)-1]
            e.left = append(e.left, ch)
        }
    default:
        e.left = append(e.left, cmd)
    }
}

func (e *Editor) Print() {
    var out []rune

    out = append(out, e.left...)

    out = append(out, '|')

    for i := len(e.right) - 1; i >= 0; i-- {
        out = append(out, e.right[i])
    }

    fmt.Println(string(out))
}


func main() {
    scanner := bufio.NewScanner(os.Stdin)

    if scanner.Scan() {
        input := scanner.Text()
        editor := &Editor{}

        for _, char := range input {
            editor.processaComando(char)
        }

        editor.Print()
    }
}