package main
import (
	"bufio"
	"fmt"
	"os"
)

func isSafe(s []byte, pos, L int, ch byte) bool {
    for j := 1; j <= L; j++ {

        if pos-j >= 0 && s[pos-j] == ch{
            return false
        }

        if pos+j < len(s) && s[pos+j] == ch {
            return false
        }
    }
    return true
}

func bt(s []byte, i, L int) bool {
    if i == len(s) {
        return true
    }
    if s[i] != '.'{
        return bt(s, i+1, L)
    }
    for d := 0; d <= L; d++ {
        ch := byte('0' + d)
        if isSafe(s, i, L, ch) {
            s[i] = ch
            if bt(s, i+1, L) {
                return true
            }
            s[i] = '.'
        }
    }

    return false
}
func main() {
    reader := bufio.NewReader(os.Stdin)
    var seq string
    var L int
    fmt.Fscan(reader, &seq, &L)
    s := []byte(seq)
    bt(s, 0, L)
    fmt.Println(string(s))
}
