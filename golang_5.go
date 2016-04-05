package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func add(a, b int) int {
	return a + b
}

func add(a, b int) (int, int) {
	return a + b, a - b
}

func add(a, b int) (x, y int) {
	x = a + b
	y = a - b
	return
}

func sum(args ...int) int {
	final := 0
	for _, value := range args {
		final += value
	}
	return final
}

func add(a, b int) int {
	num := 1

	sub := func() int {
		num -= 1
		return num

	}

	fmt.Println(sub())

	return a + b
}

func clean() {
	fmt.Println("do something in clean")
}

func safeDivision(a, b int) int {
	defer func() {
		fmt.Println(recover())
	}()
	v := a / b
	return v
}

func safeDivision(a, b int) int {
	defer func() {
		fmt.Println("defer")
	}()
	v := a / b
	return v
}

func demPanic() {

	defer func() {
		fmt.Println(recover())

	}()
	panic("PANIC")

}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(sum(1, 2, 3, 4, 5, 6))
	defer clean()
	safeDivision(1, 0)
	demPanic()
	fmt.Println("end main")
}
