package main

import "fmt"

func main() {
	//|| 逻辑或运算符 如果任何两个操作数是真，则条件为真。
	// &&逻辑与运算符 如果两个操作数是真，则条件为真
	bool_one := true
	bool_two := false
	fmt.Println("bool_one || bool_two:", bool_one || bool_two)
	fmt.Println("bool_one && bool_two:", bool_one && bool_two)

	// >, >=, <, <=, !=, ==
	fmt.Println("6 > 4:", 6 > 4)
	fmt.Println("6 >= 4:", 6 >= 4)
	fmt.Println("6 < 4:", 6 < 4)
	fmt.Println("6 <= 4:", 6 <= 4)
	fmt.Println("6 == 4:", 6 == 4)

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

	switch age {
	case 16:
		fmt.Println("your age is 16")
	case 18:
		fmt.Println("your age is 18")
	default:
		fmt.Println("your age is not 18 or 16")
	}

	//age := 18
	switch {
	case age < 18:
		fmt.Println("our age is smaller than 18")
	case age > 18:
		fmt.Println("your age is bigger than 18")
	default:
		fmt.Println("your age is 18")
	}

	for i := 0; i < 10; i++ {
		println(i)
	}
	i := 0
	for ; i < 10; i++ {
		println(i)
	}

	i = 0
	for i < 10 {
		println(i)
		i++
	}
	/*
		for {
			println(" ")
		}
	*/
}
