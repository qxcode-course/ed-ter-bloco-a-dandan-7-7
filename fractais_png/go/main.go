package main


import (
	"fmt"
	
)




func rotacao(pen *Pen, dist float64, x float64, y float64, z float64) {
	
	
    if dist < 1 {
        return
    }

    


    for i := 0; i < 3; i++ {
        pen.Walk(dist)
        dist = dist - 8
        pen.Right(90)
        if x + 75 == 255 {
            x = 0
        } else {
            x = x + 75
        }
        if y + 20 == 255 {
            y = 0
        } else {
            y = y + 20
        }
        if z + 40 == 255 {
            z = 0
        } else {
            z = z + 40
        }
        pen.SetRGB(x, y, z)
    }
    
    
    rotacao(pen, dist, x, y, z)
    
    

    
    
}

func main() {
	pen := NewPen(700, 700)
	pen.SetHeading(0)
	
	pen.SetPosition(100, 100)
	pen.SetLineWidth(1)
    x := 255.0
    y := 0.0
    z := 0.0
    pen.SetRGB(x, y, z)
	rotacao(pen, 450, x, y, z)
	
	pen.SavePNG("rotacao.png")
	fmt.Println("PNG file created successfully.")
    
}
