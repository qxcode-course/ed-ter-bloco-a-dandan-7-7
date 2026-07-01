package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func existe(node *Node, value int) bool {
	
	if node.Value == value{
		return true
	}
	a:= false
	b:= false
	if node.Left != nil{
		a = existe(node.Left, value)
		
	} 
	
	if node.Right != nil{
		b= existe(node.Right, value)
		
	} 
	
	
	if a || b{
		return true
	}
	return false

}

func BstInsert(values []int) *Node {
	// TODO
	no:= &Node{Value: values[0]}
	index:=1
	for i:=1;i<len(values);i++{
		if !existe(no, values[i]){
			if index % 2 != 0{
				no.Left = &Node{Value: values[i]}
				index++
			} else{
				no.Right = &Node{Value: values[i]}
				index++

			}

		}
	}
	
	_ = values

	return no
}

// -----------------------------------------------------------------------------------
func BShow(node *Node, history string) {
	if node != nil && (node.Left != nil || node.Right != nil) {
		BShow(node.Left, history+"l")
	}
	for i := 0; i < len(history)-1; i++ {
		if history[i] != history[i+1] {
			fmt.Print("│   ")
		} else {
			fmt.Print("    ")
		}
	}
	if history != "" {
		if history[len(history)-1] == 'l' {
			fmt.Print("╭───")
		} else {
			fmt.Print("╰───")
		}
	}
	if node == nil {
		fmt.Println("#")
		return
	}
	fmt.Println(node.Value)
	if node.Left != nil || node.Right != nil {
		BShow(node.Right, history+"r")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	values := make([]int, 0, len(parts))
	for _, elem := range parts {
		v, err := strconv.Atoi(elem)
		if err == nil {
			values = append(values, v)
		}
	}
	root := BstInsert(values)
	BShow(root, "") // Chama a função de impressão formatada
}
