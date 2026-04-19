package main
import "fmt"
func main() {
    var n, e int
    fmt.Scan(&n, &e)


    pessoas := make([]int, n)

    for i:= 0; i<n;i++ {
        pessoas[i] = i+1
    }

    var achou bool
    var valor int
    for i := 0; i < n - 1; i++ {
        achou = false
        fmt.Print("[ ")
        for j := 0; j < n; j++ {
            if pessoas[j] == 0{
                continue
            }
            if pessoas[j] != 0 {
                fmt.Print(pessoas[j])
            }
            if pessoas[j] == e {
                fmt.Print("> ")
            } else {
                fmt.Print(" ")
            }
        }
        
        fmt.Println("]")
        for j := 1; j <= n; j++ {
            if pessoas[j] == e {
                for k := j + 1; k < n; k++ {
                    if pessoas[k] != 0 {
                        pessoas[k] = 0
                        achou = true
                        valor = k + 1
                        break 
                    }
                }
                if !achou {
                    for k := 1; k <= j; k++ {
                        if pessoas[k] != 0 {
                            pessoas[k] = 0
                            achou = true
                            valor = k + 1
                            break 
                        }
                    }
                }
                if(e == n - 1) {
                    for k := 1; k <= n; k++ {
                        if pessoas[k] != 0 {
                            e = k
                            break
                        }
                    }
                } else {
                    e = valor
                }
                break
            }
        }
    }


}