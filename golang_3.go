package main

import (
	"fmt"
)

func main() {
	// ||  &&
	bool_one := true
	bool_two := false
	fmt.Println("bool_one || bool_two:", bool_one || bool_two)
	fmt.Println("bool_one && bool_two:", bool_one && bool_two)

	// >, >=, <, <=, !=, ==>
	fmt.Println("6 > 4:", 6 > 4)
	fmt.Println("6 >= 4:", 6 >= 4)
	fmt.Println("6 < 4:", 6 < 4)
	fmt.Println("6 <= 4:", 6 <= 4)
	fmt.Println("6 == 4:", 6 == 4)
	fmt.Println("6 != 4:", 6 != 4)

	//if else
	age := 18

	if age == 18 {
		fmt.Println("your age is 18")
	} else {
		fmt.Println("your age is smaller than 18")
	}

	if age == 18 {
		fmt.Println("your age is 18")
	} else if age > 18 {
		fmt.Println("your age is bigger than 18")
	} else {
		fmt.Println("your age is smaller than 18")
	}

	//switch
	switch age {
	case 16:
		fmt.Println("your age is 16")
	case 18:
		fmt.Println("your age is 18")
	default:
		fmt.Println("your age is not 18 or 16")
	}

	switch {
	case age < 18:
		fmt.Println("our age is smaller than 18")
	case age > 18:
		fmt.Println("your age is bigger than 18")
	default:
		fmt.Println("your age is 18")
	}

	// for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	i := 0
	for i < 10 {
		println(i)
		i++
	}

	for {
		println(' ')
	}
}
