package main

import (
	"bufio"
	"fmt"
	"strconv"
    "strings"
	"os"
)

func PodeColocar(vet []rune, ind, L int, digito rune) bool{

    inicio := ind - L
    if inicio < 0{
        inicio = 0
    }

    fim := ind + L
    if fim >= len(vet){
        fim = len(vet) - 1

    }
    for i:= inicio; i <= fim; i++{
        if vet[i] == digito{
            return false
        }
    }

    return true
}

func Substuir(vet []rune, L rune, ind int) bool {

    if ind >= len(vet){
        return true
    }
    if vet[ind] == '.'{
       

        for digito := '0'; digito <= L; digito++{
            if PodeColocar(vet, ind, int(L - '0'), digito){
                vet[ind] = digito
                if Substuir(vet, L, ind + 1) {
                    return true
                }
                vet[ind] = '.'
            }
        }
        return false

        //tenho q fazer iterador de digito/indice
    }
    return Substuir(vet, L, ind + 1)
}


func main() {
    var str string
    num := "0"
    L := 0
    scan := bufio.NewReader(os.Stdin)
    str, _ = scan.ReadString('\n')
    str = strings.TrimSpace(str)
    
    
    num, _ = scan.ReadString('\n')
    num = strings.TrimSpace(num) 
    L, _ = strconv.Atoi(num)
    

    vet := []rune(str)
    
    Substuir(vet, rune(num[0]), 0)
    for i:= 0; i < len(str); i++{
        fmt.Printf("%c", vet[i])
    }
    _ = L
    fmt.Println()

}