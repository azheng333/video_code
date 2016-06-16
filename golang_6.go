package main

import "fmt"


func changeValue(x int){
    x = 20
}

func changeValueNow(x *int){
    *x = 20
}

func main(){
    var p *int
    i := 2
    p = &i
    fmt.Println("memory address i:", p)
    fmt.Println(*p)  
    *p = 3
    fmt.Println(*p)
    fmt.Println(i)
    changeValue(i)
    fmt.Println("i =", i)
    changeValueNow(p)
    fmt.Println("i =", i)    
  
    pj := new(int)
    changeValueNow(pj)
    fmt.Println("pj =", *pj)
      
}