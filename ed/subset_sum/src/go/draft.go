package main

import (
	"fmt"
	//"image/jpeg"
)
func main() {
    var total int
    fmt.Scanf("%d", &total)
    var numero int
    fmt.Scanf("%d", &numero)
    vetor := make([]int, total)

    for i := 0; i < total; i++ {
        fmt.Scanf("%d", &vetor[i])
    }

    for i := 0; i < total; i++ {
        for j := 0; j < total; j++ {
            if(vetor[i] + vetor[j] == numero) {
                fmt.Printf("true\n")
                return 
            }
        }
    }

}