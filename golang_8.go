package main

import "fmt"
import "math"

type Shape interface {
    area() float64
}

type MultiShape interface {
    Shape
}

type Rectangle struct {
    height float64
    width float64
}

type Circle struct {
    radius float64
}

func (r Rectangle) area() float64{
    return r.height * r.width
}

func (c Circle) area() float64 {
    return math.Pi * math.Pow(c.radius, 2)
}

func getArea(shape MultiShape) float64 {
    return shape.area()
}

func main(){
    r := Rectangle{20, 10}
    c := Circle{4}
    fmt.Println("Rectangle Area=", getArea(r))
    fmt.Println("Circle Area=", getArea(c))    
}
/*type Person struct{
    name string
}

type Students struct{
    class int
    Person
}

func main(){
    s := Students{1, Person{"weizheng"}}
    fmt.Println("name:", s.name)
}*/