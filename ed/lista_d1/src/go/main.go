package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	
)

type Node struct{
    Value int            // Valor é público
    next *Node           // o próximo nó da lista
    prev *Node           // o nó anterior
}

type LList struct {
    root *Node
}

func NewLList() *LList{
	ll := &LList{}

	node := &Node{
		Value: 0,
	}

	node.next = node
	node.prev = node

	ll.root = node

	return ll

}

func (ll *LList) Join(sep string) string {
	str := ""

	for it := ll.root.next; it != ll.root; it = it.next{
		str+= fmt.Sprintf("%d", it.Value)

		if it.next != ll.root{
			str += sep
		}
	}

	return str
}

func (ll *LList) String() string {
    return "[" + ll.Join(", ") + "]"
}

func (ll *LList) PushBack(num int) {

	newNode := &Node{
		Value: num,
	}

	newNode.next = ll.root
	newNode.prev = ll.root.prev
	newNode.prev.next = newNode
	ll.root.prev = newNode

}

func (ll *LList) PushFront(num int) {
	
	newNode := &Node{
		Value: num,
	}

	newNode.next = ll.root.next
	newNode.prev = ll.root
	ll.root.next = newNode
	newNode.next.prev = newNode

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
func (ll *LList) Clear(){

	ll.root.next = ll.root
    ll.root.prev = ll.root
}

func (ll *LList) Size() int {

	ac := 0
	for it := ll.root.next; it != ll.root; it = it.next{
		ac++
	}
	return ac
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
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
