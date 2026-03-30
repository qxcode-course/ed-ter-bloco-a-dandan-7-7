package main
import "fmt"
func main() {

    var tament, tambus int
    fmt.Scan(&tament)
    
    entradas := make([]string, tament)

    for i := 0; i < tament; i++ {
        fmt.Scan(&entradas[i])
    }

    fmt.Scan(&tambus)

    consultas := make([]string, tambus)
    for i := 0; i < tambus; i++ {
        fmt.Scan(&consultas[i])
    }
    
    ocorrencias := make([]int, tambus)

    for i := 0; i < tambus; i++ {
        for j := 0; j < tament; j++ {
            if consultas[i] == entradas[j] {
                ocorrencias[i]++
            }
        }
    }
    
    for i := 0; i < tambus; i++ {
        fmt.Printf("%d", ocorrencias[i])
        if(i < tambus-1) {
            fmt.Print(" ")
        }
    }
    fmt.Println()

    


}
