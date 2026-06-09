package main
import "fmt"
func main() {
    n := 0; f := 0
    fmt.Scan(&f, &n)
    
    fig := map[int]int{}

    for i :=0; i < n; i++{
        c := 0
        fmt.Scan(&c)
        
        
        fig[c]++
        

    }

    for i := 1; i<=f; i++{
        if fig[i]>1{
            fmt.Printf("%d ", i)
            fig[i]--
        }
    }

}
