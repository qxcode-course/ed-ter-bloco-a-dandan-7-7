package main
import (
    "fmt"
    "math"
    ) 
    
func main() {
    var a, b, n, venc, atual int
    venc = -1
    fmt.Scan(&n)
    melhorDiferenca := 101
    for i := 0; i < n; i++ {
        fmt.Scan(&a, &b)

        if a >= 10 && b >= 10{
            atual = int(math.Abs(float64(a - b)))

            if atual < melhorDiferenca {
                melhorDiferenca = atual
                venc = i 
            }
             
            }
        
    }

    if venc == -1 {
        fmt.Println("sem ganhador")
    } else {
        fmt.Println(venc)
    }


}



