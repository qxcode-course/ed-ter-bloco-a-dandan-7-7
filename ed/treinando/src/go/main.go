package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func tostr(vet []int) string {
	_ = vet
	

	str:= "["

	for i:= 0; i< len(vet);i++{
		str+= fmt.Sprintf("%d", vet[i])
		if i < len(vet)-1{
			str+= fmt.Sprint(", ")			
		}
	}
	str+= "]"
	return str
}

func tostrrev(vet []int) string {
	str:= "["

	for i:= len(vet)-1; i>= 0;i--{
		str+= fmt.Sprintf("%d", vet[i])
		if i > 0{
			str+= fmt.Sprint(", ")			
		}
	}
	str+= "]"
	return str
}

// reverse: inverte os elementos do slice
func reverse(vet []int) {
	_ = vet
	reverse := make([]int, len(vet))
	i, j := 0, len(vet)-1
	for i < len(vet){
		reverse[i] = vet[j]
		i++
		j--
	}
	
	
	for i = 0; i < len(vet);i++{
		vet[i] = reverse[i]

	}

}

// sum: soma dos elementos do slice
func sum(vet []int) int {
	_ = vet
	sum:= 0
	for i:= 0;i<len(vet);i++{
		sum =vet[i] + sum
	}
	return sum
}

// mult: produto dos elementos do slice
func mult(vet []int) int {
	_ = vet
	m:= 1

	for i:= 0; i < len(vet);i++{
		m *= vet[i]
	}

	return m
}

// min: retorna o índice e valor do menor valor
// crie uma função recursiva interna do modelo
// var rec func(v []int) (int, int)
// para fazer uma recursão que retorna valor e índice
func min(vet []int) int {
	_ = vet
	if len(vet) == 0 {
		return -1
	}

	var rec func(v []int) (int, int)
	
	rec = func(v []int) (int, int) {
		if len(v) == 1 {
			return v[0], 0
		}

		val, ind:= rec(v[1:])

		if val < v[0] {
			return val, ind + 1
		}
		return v[0], 0
		
		
	}

	_ , ind := rec(vet)

	return ind
}

func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			vet = nil
			for _, arg := range args[1:] {
				if val, err := strconv.Atoi(arg); err == nil {
					vet = append(vet, val)
				}
			}
		case "tostr":
			fmt.Println(tostr(vet))
		case "torev":
			fmt.Println(tostrrev(vet))
		case "reverse":
			reverse(vet)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			fmt.Println(min(vet))
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
