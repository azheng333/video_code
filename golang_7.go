package main

import "fmt"

type Rectangle struct {
    width float64
    height float64
}

func (r Rectangle) area() float64{
    return r.width * r.height
}

func (r *Rectangle) changewidth() {
    r.width = 30
}

func main(){
    var r Rectangle
    r = Rectangle{width:20, height:10}
    r = Rectangle{20, 10}
    fmt.Println("the Rectangle width:", r.width)
    fmt.Println("the area of Rectangle:", r.area())
    r.changewidth()
    fmt.Println("the Rectangle width:", r.width)
}