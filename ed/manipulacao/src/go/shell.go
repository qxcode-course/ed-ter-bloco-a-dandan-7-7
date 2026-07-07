package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getMen(vet []int) []int {
	_ = vet
	var men []int
	
	for i:= 0; i < len(vet); i++ {
		if vet[i] > 0 {
			men = append(men, vet[i])
		}
	}
	return men
}

	
func getCalmWomen(vet []int) []int {
	_ = vet
	var gcw []int
	for i:= 0; i < len(vet); i++ {
		if vet[i] < 0 && vet[i] > -10 {
			gcw = append(gcw, vet[i])
		}
	}

	return gcw
}

func sortVet(vet []int) []int {
	_ = vet

	
	
	for i := 0; i < len(vet); i++ {
		
		for j := i + 1; j < len(vet); j++ {
			if vet[i] > vet[j] {
				vet[i], vet[j] = vet[j], vet[i]
			}
		
		}
		
	}

	return vet
}

func sortStress(vet []int) []int {
	for i := 0; i < len(vet); i++ {
		
		for j := i + 1; j < len(vet); j++ {
			if abs(vet[i]) > abs(vet[j]) {
				vet[i], vet[j] = vet[j], vet[i]
			}
		
		}
		
	}

	return vet
}

func reverse(vet []int) []int {
	_ = vet
	var rev  []int
	
	for i:= len(vet)-1; i >=0; i--{
		rev = append(rev, vet[i])
		
	}
	return rev
}

func unique(vet []int) []int {
	_ = vet
	var aux []int
	aux = append(aux, vet[0])
	for i:= 1; i < len(vet); i++{
		v := true
		for j:=0 ; j<len(aux); j++{
			if vet[i] == aux[j]{
				v = false
			}

		}
		if v{
			aux = append(aux, vet[i])
		}

	}

	return aux
}

func repeated(vet []int) []int {
	_ = vet
	var rep []int

	for i:= 0; i < len(vet); i++{
		apantes:= false
		for j:=0 ; j<i; j++{
			if vet[i] == vet[j]{
				apantes = true
				break
			}
			
		}
		if apantes{

			rep = append(rep, vet[i])
		}
	}

	
	return rep
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}

