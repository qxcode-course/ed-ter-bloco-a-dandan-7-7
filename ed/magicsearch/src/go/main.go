package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MagicSearch(slice []int, value int) int {
	_, _ = slice, value
	val := 0
	b := false
	for i := 0; i < len(slice);i++{
		if slice[i] == value{
			val = i
			b = true
		}
	}
	if b{
		return val
	}
	
	for i:= 1; i < len(slice) ;i++{
		if slice[i-1] < value && slice[i] > value{ 
			return i
		}
	}
		if value == 0{
			return 0
		}
	return len(slice)


	//return 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	slice := make([]int, 0, 1)
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}

	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	result := MagicSearch(slice, value)
	fmt.Println(result)
}
