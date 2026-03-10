package main

import "fmt"

func main() {
	var n, e int
	fmt.Scan(&n, &e)
	//var j int

	vivos := []int{}

	for i := 0; i < n; i++ {

		vivos = append(vivos, i+1)
	}

	fmt.Print("[")
    var j int
	for i := 0; i < n; i++ {

        if(vivos[i] == e){
            vivos[i + 1] = 0
            if(e >= n) {
                e==1
            } else {
                e +=2
            }

            
        }

        if(vivos[i] != 0){
		fmt.Print(vivos[i], " ")
	}
	fmt.Print("]")

}
