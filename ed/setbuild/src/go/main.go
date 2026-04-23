package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Set struct{
   data []int                           
   size int                               
   capacity int 
}

func NewSet(valor int) *Set{
	return &Set{
		data : make([]int, valor),
		size : 0,
		capacity : valor,
	}
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	var result strings.Builder
	fmt.Fprintf(&result, "%d", slice[0])
	for _, value := range slice[1:] {
		fmt.Fprintf(&result, "%s%d", sep, value)
	}
	return result.String()
}

func (set *Set) String() string{
	return "[" + Join(set.data[0:set.size], ", ") + "]"
}

func (set *Set) Insert(value int) {
	for i := 0; i < set.size; i++{
		if set.data[i] == value{
			return
		}
	}
	if set.size == set.capacity{
		newCapacity := set.capacity * 2
		if newCapacity == 0 {
			newCapacity = 1
		}
		novo := NewSet(newCapacity)
		for i:= 0; i < set.capacity; i++{
			novo.data[i] = set.data[i]
		}
		set.data = novo.data
		set.capacity = novo.capacity
	}
	
	indice := set.size
	for i:= 0; i < set.size; i++{
		if value < set.data[i]{
			indice = i
			break
		}
	}
	for j := set.size; j > indice; j--{
			set.data[j] = set.data[j - 1]
		}
	set.data[indice] = value
	set.size++
	

	
} 

func (set *Set) Erase(value int) bool{

	indice := -1
	for i := 0; i < set.size; i++{
		if value == set.data[i]{
			indice = i
			break
		}
	}
	if indice == -1{
		return false
	}
	for i:= indice; i <set.size - 1; i++{
		set.data[i] = set.data[i+1]
	}
	set.data[set.size - 1] = 0
	set.size--
	return true

}
func (set *Set) Contains(value int) bool{
	for i := 0; i < set.size; i++{
		if set.data[i] == value{
			return true
		}
	}
	return false
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewSet(value)
		case "insert":
			 for _, part := range parts[1:] {
			 	value, _ := strconv.Atoi(part)
				v.Insert(value)
			}
		case "show":
			fmt.Println(v)
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			sucesso := v.Erase(value)
			if sucesso{
				//fmt.Println("true")
			} else{
				fmt.Println("value not found")
			}
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			contem := v.Contains(value)
			if contem{
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}

			
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
