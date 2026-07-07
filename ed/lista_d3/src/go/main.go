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
	next  *Node
	prev  *Node
	root  *Node
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	list := &LList{}
	list.root = &Node{root: nil}
	list.root.next = list.root
	list.root.prev = list.root
	list.root.root = list.root // nó sentinela aponta pra si mesmo
	return list
}

func (l *LList) PushBack(value int) {
	l.insertBefore(l.root, value)
}

func (l *LList) insertBefore(mark *Node, value int) {
	n := &Node{
		Value: value,
		root:  l.root,
	}
	n.prev = mark.prev
	n.next = mark
	mark.prev.next = n
	mark.prev = n
}


func str2list(serial string) *LList {
	serial = serial[1 : len(serial)-1]
	ll := NewLList()
	if serial == "" {
		return ll
	}
	for _, p := range strings.Split(serial, ",") {
		value, _ := strconv.Atoi(p)
		ll.PushBack(value)
	}
	return ll
}

func equals(lla, llb *LList) bool {
	
	n1 := lla.root.next
	n2 := llb.root.next
	for {

		if n1 == lla.root && n2 != llb.root {
			return false
		}
		if n1 != lla.root && n2 == llb.root {
			return false
		}
		if n1 == lla.root && n2 == llb.root {
			break
		}
		if n1.Value != n2.Value {
			return false
		}

		n1 = n1.next
		n2 = n2.next
	}
	return true
}

func merge(lla, llb *LList) *LList {
	l1 := []int{}
	for n := lla.root.next; n != lla.root; n = n.next {
		l1 = append(l1, n.Value)
	}
	l2 := []int{}
	for n := llb.root.next; n != llb.root; n = n.next {
		l2 = append(l2, n.Value)
	}
	l3 := []int{}
	i:= 0
	j:= 0 
	for i<len(l1) && j<len(l2){
	
		if l1[i] <= l2[j]{
			l3 =append(l3, l1[i])
			i++
		} else {
			l3 = append(l3, l2[j])
			j++
		}
		
		
	}

	for ;i<len(l1);i++{
		l3 = append(l3, l1[i])
	}
	for ;j<len(l2);j++{
		l3 = append(l3, l2[j])
	}


	llc := NewLList()
	for x:= 0; x < len(l3);x++{
		llc.PushBack(l3[x])
	}
	return llc

}

func (l *LList) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	for n := l.root.next; n != l.root; n = n.next {
		sb.WriteString(fmt.Sprintf("%d", n.Value))
		if n.next != l.root {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]")
	return sb.String()
}

func reverse(lla *LList) {
	
	n:= NewLList()
	for p := lla.root.prev; p != lla.root; p = p.prev {
		n.PushBack(p.Value)
	}
	*lla = *n
	

}

func addsorted(lla *LList, value int) {

	for n := lla.root.next; n != lla.root; n = n.next {
		if value <= n.Value {
			lla.insertBefore(n, value)
			return
		}
	}
	lla.PushBack(value)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

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
		case "compare":
			lla := str2list(args[1])
			llb := str2list(args[2])
			if equals(lla, llb) {
				fmt.Println("iguais")
			} else {
				fmt.Println("diferentes")
			}
		case "addsorted":
			lla := NewLList()
			for i := 1; i < len(args); i++ {
				value, _ := strconv.Atoi(args[i])
				addsorted(lla, value)
			}
			fmt.Println(lla)
		case "reverse":
			lla := str2list(args[1])
			reverse(lla)
			fmt.Println(lla)
		case "merge":
			lla := str2list(args[1])
			llb := str2list(args[2])
			merged := merge(lla, llb)
			
			fmt.Println(merged)
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
