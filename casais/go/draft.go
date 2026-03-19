package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)
    descasados := make([]int, n)
    for i := 0; i < n; i++{
        fmt.Scan(&descasados[i])
    }
    k := 0
    for i := 0; i < n; i++{
        for j := i + 1; j < n; j++{
            if descasados[i] == (-1)*descasados[j] && descasados[i] != 0 && descasados[j] != 0{
                k++
                descasados[i] = 0
                descasados[j] = 0
                break
            }
        }
}
    fmt.Println(k)
}