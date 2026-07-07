package main
import "fmt"
func main() {
    n :=0
    v:= 0
    fmt.Scan(&n, &v)
    vetor := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&vetor[i])
    }

    v = v % n

    resultado := make([]int, n)

    for i := 0; i < n;
     i++ {

        vP := (i+v)%n
        resultado[vP] = vetor[i]

    }
    fmt.Print("[ ")
    for i := 0; i < n; i++ {
        fmt.Print(resultado[i], " ")
    }
    fmt.Println("]")
}
