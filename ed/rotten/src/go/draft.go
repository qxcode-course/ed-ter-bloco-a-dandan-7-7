package main
import (
    "fmt"
    "bufio" 
    "os"
    "strings"
    "strconv"
)

type pos struct {
    x int
    y int
}  


func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    
    scanner.Scan()
    primeiraLinha := scanner.Text()
    parts := strings.Fields(primeiraLinha)
    ms := parts[0]
    ns := parts[1]

    m, _:= strconv.Atoi(ms)
    n, _:= strconv.Atoi(ns)

    grid := make([][]int, m)
    for i := 0; i < m; i++ {
        grid[i] = make([]int, n)
        scanner.Scan()
        line := scanner.Text()
        parts := strings.Fields(line)
        for j := 0; j < n; j++ {
            grid[i][j], _ = strconv.Atoi(parts[j])
        }
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            fmt.Print(grid[i][j])
            if j < n-1 {
                fmt.Print(" ")
            }
        }
        fmt.Println()
    }

}