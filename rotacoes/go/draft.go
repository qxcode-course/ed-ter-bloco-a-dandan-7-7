package main
import "fmt"
func main() {
    var tam, rot int
    fmt.Scan(&tam)
    fmt.Scan(&rot)
    t := make([]int, tam)
    for i := 0; i < tam; i++ {
        fmt.Scan(&t[i])
    }
    var ult int
    for i:= 0; i < rot; i++ {
        ult = t[tam-1]
        for j := tam - 1; j > 0; j-- {
            t[j] = t[j-1]
        }
        t[0] = ult
    }
    fmt.Print("[ ")
    for i:= 0; i < tam; i++ {
        fmt.Print(t[i])
        if(i < tam - 1){
            fmt.Print(" ")
        }
    }
    fmt.Print(" ]")
    fmt.Println()
}
