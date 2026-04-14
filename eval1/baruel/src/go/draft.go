package main
import "fmt"
func main() {
    var total, possui int
    fmt.Scanf("%d", &total)
    fmt.Scanf("%d", &possui)
    var figurinhas = make([]int, possui)
    for i := 0; i < possui; i++ {
        fmt.Scanf("%d", &figurinhas[i])
    }

    b:= 0
    for i := 1; i<possui;i++{
        if figurinhas[i - 1] == figurinhas[i]{
            fmt.Printf("%d ", figurinhas[i])
            b = 1
        }
    }
    if b == 0{
        fmt.Printf("N")
    }
    fmt.Println()

    
    var pos int = 0
    ver := 0
    faltantes := make([]int, total - possui)
    k := 0
    for i := 0; i <= total; i++{
        pos = 0
        for j := 0; j < possui; j++{
            if i == figurinhas[j]{
                pos = 1
                ver = 1
                

            }

        }
        if pos == 0 && i != 0{
            faltantes[k] = i
            k++
            
        }
    }
    if ver == 0{
        fmt.Printf("N")
    } else{
    for i := 0; i < total - possui; i++{
        fmt.Printf("%d ", faltantes[i])
        if( i < total - possui - 1){
            fmt.Printf(" ")
        }
    }
}
    
    fmt.Println()

}
