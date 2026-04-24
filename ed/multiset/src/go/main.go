package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"errors"
)

type MultiSet struct{
  	data []int                            
  	size int                               
  	capacity int  
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}
func NewMultiSet(value int) *MultiSet{
	return &MultiSet{
		data: make([]int, value),
		size: 0,
		capacity: value,
	}
}

func (ms *MultiSet) String() string{
	return "[" + Join(ms.data[0:ms.size], ", ") + "]"
}
func (ms *MultiSet) Insert(value int) {
	if ms.capacity == ms.size{
		if ms.capacity == 0{
			ms.capacity = 1
		} else {
			
			novo:= NewMultiSet(ms.capacity*2)
			for i:= 0; i < ms.size; i++{
				novo.data[i] = ms.data[i]
			}

			ms.data = novo.data
			ms.capacity = novo.capacity

		}
	} 

	indice:= ms.size
	for i:= 0; i < ms.size; i++{
		if ms.data[i] >= value{
			indice = i
			break
		}
	}
	for i:= ms.size; i > indice; i--{
		ms.data[i] = ms.data[i-1]
	}
	ms.size++
	ms.data[indice] = value
}

func (ms *MultiSet) Clear(){
	for i := 0; i < ms.size; i++{
		ms.data[i] = 0
	}
	ms.size = 0
}
func (ms *MultiSet)Contains(value int) bool {
	for i:= 0; i < ms.size; i++{
		if value == ms.data[i]{
			return true
		}
	}
	return false
}

func (ms *MultiSet) Erase(index int) error{
	
	indice := -1
	for i:= 0; i < ms.size; i++{
		if ms.data[i] == index{
			indice = i
			break
		}
	}
	if indice == -1{
		return errors.New("value not found")
	}
	for i:= indice; i < ms.size; i++{
		ms.data[i] = ms.data[i+1]
	}
	ms.size--
	return nil
}

func (ms *MultiSet) Count(value int) int{
	ocorrencias := 0
	for i:= 0; i < ms.size; i++{
		if value == ms.data[i] {
			ocorrencias++
		}
	}
	return ocorrencias
}
func (ms *MultiSet)Unique() int{
	oco := 0
	num := 0
	for i := 0; i < ms.size; i++{
		if ms.data[i] > num{
			num = ms.data[i]
			oco++
		}
	}
	return oco
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
			 	value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
		case "show":
			fmt.Println(ms)
		case "erase":
			
			value, _ := strconv.Atoi(args[1])
			erro := ms.Erase(value)
			if erro != nil{
				fmt.Println(erro)
			} else {
			}
				
		case "contains":
			value, _ := strconv.Atoi(args[1])
			contem := ms.Contains(value)
			fmt.Println(contem)
		case "count":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Count(value))
		case "unique":
			fmt.Println(ms.Unique())
		case "clear":
			ms.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
