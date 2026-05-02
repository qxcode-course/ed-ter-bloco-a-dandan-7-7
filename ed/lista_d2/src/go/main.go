package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Node struct{
    value int 
    next *Node
    prev *Node
    root *Node
    
}    

//Next() *Node
//Prev() *Node

func (node *Node) Next() *Node{
	if node.next == node.root {
		return nil
	}
	return node.next
}

func (node *Node) Prev() *Node{
	if node.prev == node.root {
		return nil
	}
	return node.prev
}


type LList struct {
    root *Node
}

func (ll *LList) Front() *Node{
	return ll.root.Next()
}

func (ll *LList) Back() *Node{
	return ll.root.Prev()
}

func NewLList() *LList{
	newNode := &Node{
		value:0,
	}

	newNode.next = newNode
	newNode.prev = newNode
	newNode.root = newNode
	ll := &LList{}
	ll.root = newNode
	return ll
}

func (ll *LList) Size() int{
	ac := 0
	for it := ll.root.next; it != ll.root; it = it.next{
		ac++
	}
	return ac
}


func (ll *LList) Join(sep string) string {
	str := ""
	
	for it := ll.root.next; it != ll.root; it = it.next{
		str+= fmt.Sprintf("%d", it.value)
		
		if it.next != ll.root{
			str += sep
		}
	}
	
	return str
}

func (ll *LList) String() string {
	return "[" + ll.Join(", ") + "]"
}

func (ll *LList) Search(value int) *Node{

	for it := ll.root.next; it != ll.root; it = it.next{
		if it.value == value{
			return it
		}
	}

	return nil
}

func (ll *LList) Remove(node *Node) *Node{
	node.next.prev = node.prev
	node.prev.next = node.next 

	return node.Next()
}

func (ll *LList) PushBack(valor int) {
	novo := &Node{
		value: valor,
		root: ll.root,
	}

	novo.prev = ll.root.prev
	novo.next = ll.root
	novo.prev.next = novo
	ll.root.prev = novo

}

func (ll *LList) PushFront(valor int) {
	novo := &Node{
		value: valor,
		root: ll.root,
	}

	novo.prev = ll.root
	novo.next = ll.root.next
	novo.next.prev = novo
	ll.root.next = novo

}

func (ll *LList) Insert(node *Node, valor int){

	novo := &Node{
		value: valor,
		root: ll.root,
	}

	novo.next = node
	novo.prev = node.prev
	node.prev = novo
	novo.prev.next = novo

}

func (ll *LList) Clear(){
	ll.root.next = ll.root
	ll.root.prev = ll.root
}
func (ll *LList) PopBack(){

	//temp := ll.root.prev
	ll.root.prev = ll.root.prev.prev
	ll.root.prev.next = ll.root
	
}

func (ll *LList) PopFront(){

	//temp := ll.root.prev
	ll.root.next = ll.root.next.next
	ll.root.next.prev = ll.root
	
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)
 	ll := NewLList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			fmt.Println(ll.String())
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
			num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.Next() {
				fmt.Printf("%v ", node.value)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.Prev() {
				fmt.Printf("%v ", node.value)
			}
			fmt.Println("]")
		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				node.value = newvalue
			} else {
				fmt.Println("fail: not found")
			}
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Insert(node, newvalue)
			} else {
				fmt.Println("fail: not found")
			}
		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Remove(node)
			} else {
				fmt.Println("fail: not found")
			}
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
