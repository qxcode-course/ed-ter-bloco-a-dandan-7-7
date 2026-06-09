package main
import "fmt"
func main() {
    qtd := 0

    n := 0
    fmt.Scan(&n)

    descasados := map[int]int{}

    
    for i:= 0; i < n; i++{
        animal := 0
        fmt.Scan(&animal)

        par := animal * -1

        if descasados[par] >0{
            qtd++
            descasados[par]--
        } else{
            descasados[animal]++
            
        }


    }

    
    fmt.Println(qtd)
}
