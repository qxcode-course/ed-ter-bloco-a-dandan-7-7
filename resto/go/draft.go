package main
import "fmt"

func div(num, rest int)  {

    if num == 0 {
        fmt.Printf("%d %d", num, rest)
        fmt.Println()
        return
    }
    div(num/2, num % 2)
    
    fmt.Print(num, " ", rest)
    fmt.Println()

}

func main() {
    var n int
    fmt.Scan(&n)
    
    div(n/2, n%2)

}
