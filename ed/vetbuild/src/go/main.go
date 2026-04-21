package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
	"strconv"
	"errors"
)

type Vector struct {
	data     []int
	size     int
	capacity int
}

func NewVector(capacity int) *Vector {
	return &Vector{
		data:     make([]int, capacity), // nunca use len(data) ou cap(data) ou qq método do go de manipulação de array
		size:     0,
		capacity: capacity,
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

func (v *Vector) String() string {
	return "[" + Join(v.data[0:v.size], ", ") + "]"
}

func (v *Vector) Status() string{
	return fmt.Sprintf("size:%d capacity:%d", v.size, v.capacity)
}


func (v *Vector)PushBack (valor int) {
	if v.size == v.capacity {
		
		Vetor := NewVector(v.capacity * 2)

		for i := 0; i < v.capacity ; i++{
			Vetor.data[i] = v.data[i]
		}
		v.data = Vetor.data
		v.capacity = Vetor.capacity

	}
	v.data[v.size] = valor
	v.size++
}

func (v *Vector) At(indice int)  (int, error){

	if indice >= v.capacity || indice < 0{
		return 0, errors.New("index out of range")
	}
	return v.data[indice], nil
}
func (v *Vector) Capacity() int{
	return v.capacity
}

func (v *Vector) Clear() {
	for i:= 0; i < v.capacity; i++{
		v.data[i] = 0
	}
	v.size = 0

}

func (v *Vector) Reserve(valor int){
	v.capacity = valor
}
func (v *Vector) Insert(indice int, valor int) error{
	if indice >= v.capacity || indice < 0{
		return errors.New("index out of range")
	}

	if v.size == v.capacity {
		
		Vetor := NewVector(v.capacity * 2)

		for i := 0; i < v.capacity ; i++{
			Vetor.data[i] = v.data[i]
		}
		v.data = Vetor.data
		v.capacity = Vetor.capacity

	}

	if indice < v.size{
		v.size++
		for i:= v.size - 1; i > indice; i--{
			v.data[i] = v.data[i-1]
		}
	}
	v.data[indice] = valor
	return nil
}
func (v *Vector) PopBack() error{
	if v.size >0{
		v.data[v.size-1] = 0
		v.size--
		return nil
	} else {
		return errors.New("vector is empty")

	}
}
 
func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	 v := NewVector(0)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
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
			v = NewVector(value)
		case "push":
			for _, part := range parts[1:] {
			 	value, _ := strconv.Atoi(part)
			 	v.PushBack(value)
			}
		case "show":
			fmt.Println(v)
		case "status":
			fmt.Println(v.Status())
		case "pop":
			err := v.PopBack()
			if err != nil {
			 	fmt.Println(err)
			}
		case "insert":
			 index, _ := strconv.Atoi(parts[1])
			 value, _ := strconv.Atoi(parts[2])
			 err := v.Insert(index, value)
			if err != nil {
				fmt.Println(err)
			}
		case "erase":
			//index, _ := strconv.Atoi(parts[1])
			//err := v.Erase(index)
			//if err != nil {
			//	fmt.Println(err)
			//}
		case "indexOf":
			// value, _ := strconv.Atoi(parts[1])
			// index := v.IndexOf(value)
			// fmt.Println(index)
		case "contains":
			// value, _ := strconv.Atoi(parts[1])
			// if v.Contains(value) {
			// 	fmt.Println("true")
			// } else {
			// 	fmt.Println("false")
			// }
		case "clear":
			v.Clear()
		case "capacity":
			 fmt.Println(v.Capacity())
		case "get":
			index, _ := strconv.Atoi(parts[1])
			value, err := v.At(index)
			 if err != nil {
			 	fmt.Println(err)
			 } else {
			 	fmt.Println(value)
			}
		case "set":
			// index, _ := strconv.Atoi(parts[1])
			// value, _ := strconv.Atoi(parts[2])
			// err := v.Set(index, value)
			// if err != nil {
			// 	fmt.Println(err)
			// }
			// 
		case "reserve":
			newCapacity, _ := strconv.Atoi(parts[1])
			v.Reserve(newCapacity)
		case "slice":
			// start, _ := strconv.Atoi(parts[1])
			// end, _ := strconv.Atoi(parts[2])
			// slice := v.Slice(start, end)
			// fmt.Println(slice)
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
