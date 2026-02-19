import (
    "strconv"
    "fmt"
)

func hasAlternatingBits(n int) bool {
    str := strconv.FormatInt(int64(n), 2)
    for i, char := range str {
        if i > 0 && char == rune(str[i-1]) {
            return false
        }
    }
	return true
}