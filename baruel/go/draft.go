package main
import "fmt"
func main() {
    var qt, qf int
    fmt.Scan(&qt, &qf)
    figurinhas := make([]int, qt)
    rep := 0
    frp := make([]int, 50)
    for i := 0; i < qf; i++ {
        fmt.Scan(&figurinhas[i])
        if(i > 0) {
            if figurinhas[i] == figurinhas[i-1] {
                rep++
                frp[i] = figurinhas[i]
            }
        }
    }
    if rep == 0 {
        fmt.Println("N")
    } else {
        
    }
}
