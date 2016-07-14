package main

//import "os"
import "fmt"
//import "io/ioutil"
import "strconv"

func main(){
    //file, _ := os.Create("test.txt")  
    /*
    file, _ := os.OpenFile("test.txt", os.O_RDWR, 0666) 
    file.WriteString("a string in a line")
    file.Close()
    */
    /*
    stream, _ := ioutil.ReadFile("test.txt")     
    readString := string(stream)
    fmt.Println(readString)
    */
    /*
    file, _ := os.Open("test.txt")
    buf := make([]byte, 1024)
    n, _ := file.Read(buf)
    readString := string(buf[:n])
    fmt.Println(readString)
    */
    /*
    fmt.Println("what is your name? ")
    var name string
    fmt.Scan(&name)
    fmt.Println(name)
    */
    num1 := 5
    num2 := 11.5
    str1 := "100"
    str2 := "120.5"
    
    fmt.Println(float64(num1))
    fmt.Println(int(num2))
    fmt.Println(strconv.ParseInt(str1, 0, 64))   
    fmt.Println(strconv.ParseFloat(str2, 64))
}