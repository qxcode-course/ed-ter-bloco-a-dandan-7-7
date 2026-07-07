package main

import (
	"fmt"
	
)

func BT(idx, somaAtual, obj int, moedas[] int) bool{
    if somaAtual == obj{
        return true
    }
    if somaAtual > obj{
        return false
    }
    if idx >= len(moedas){
        return false
    }

    a:= BT(idx+1, somaAtual + moedas[idx], obj, moedas)
    
    if a{
        return true
    }
    b:= BT(idx+1, somaAtual, obj, moedas)
    if b{
        return true
    }


    return false
}

func main() {
    var total int
    fmt.Scan(&total)
    var numero int
    fmt.Scan(&numero)
    vetor := make([]int, total)

    for i := 0; i < total; i++ {
        fmt.Scan(&vetor[i])
    }

    fmt.Println(BT(0, 0, numero, vetor))

}