package main
import "fmt"



type Pessoas struct{
    val int
    viv bool
}

func  Proxviv(in int, p []Pessoas) int {
    n := len(p)
    i := (in + 1)% n
    for {
        if p[i].viv {
            return i
        }
        i = (i + 1) % n
    }

}

func main() {
    var n, e int
    fmt.Scan(&n, &e)


    pessoas := make([]Pessoas, n)
    for i := 0; i < n; i++ {
        pessoas[i].val = i+1
        pessoas[i].viv = true
    }
    
    total := len(pessoas) - 1
    espadaIdx := e - 1
    for total > 0 {
        
        
        fmt.Print("[ ")
        for k := 0; k < len(pessoas); k++{
            if pessoas[k].viv{
                fmt.Print(pessoas[k].val)
                if k == espadaIdx {
                    fmt.Print("> ")
                } else{
                    fmt.Print(" ")
                }
            }
        }
        fmt.Print("]\n")


        vitimaIdx := Proxviv(espadaIdx, pessoas)
		pessoas[vitimaIdx].viv = false 
		total--
        espadaIdx = Proxviv(vitimaIdx, pessoas)
        
      
        
    }
    fmt.Print("[")
	for k := 0; k < n; k++ {
		if pessoas[k].viv {
			fmt.Printf(" %d> ", pessoas[k].val)
		}
	}
	fmt.Println("]")

}