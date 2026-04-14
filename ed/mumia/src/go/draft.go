package main
import "fmt"
func main() {
    var nome string
    var idd int
    fmt.Scanf("%s", &nome)
    fmt.Scanf("%d", &idd)

    fmt.Printf("%s eh ", nome)


    if(idd < 12){
        fmt.Printf("crianca\n")
    } else if(idd >= 12 && idd < 18){
        fmt.Printf("jovem\n")
    } else if(idd >= 18 && idd < 65){
        fmt.Printf("adulto\n")
    } else if(idd >= 65 && idd < 1000){
        fmt.Printf("idoso\n")
    } else{
        fmt.Printf("mumia\n")
    }
}