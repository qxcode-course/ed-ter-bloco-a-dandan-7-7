package main
import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
    )
func main() {
    scanner := bufio.NewScanner(os.Stdin)

    scanner.Scan()
    linha1 := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(linha1[0])
    
    bombs := make([][]int, n)
    
   
}