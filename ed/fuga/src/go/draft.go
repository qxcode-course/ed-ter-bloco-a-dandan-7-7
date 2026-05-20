package main
import "fmt"
func main() {
    var h, p, f, d int
    fmt.Scanf("%d %d %d %d", &h, &p, &f, &d)

    for {
         f += d
         if f >15{
             f = 0
         }
         if f < 0{
             f = 15
         }
        if f == h{
            fmt.Println("S")
            break
        }
        if f == p{
            fmt.Println("N")
            break
        }
       

    }
}
