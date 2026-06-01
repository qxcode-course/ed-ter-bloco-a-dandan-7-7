package main

import (
	"fmt"

	
)

func main() {
	v:=NewQueue[rune]()
	for i:=0;i<16;i++{
		v.Enqueue('A' + rune(i)) 
	}

	//oitavas
	a:= 0
	b:= 0
	
	for i := range 15{
		fmt.Scan(&a, &b)
		t1 := v.Dequeue()
		t2 := v.Dequeue()

		if a > b{
			v.Enqueue(t1)
		} else{
			v.Enqueue(t2)
		}
		_ = i
	}

	fmt.Println(v)
}
