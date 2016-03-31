package main

import (
	"fmt"
)

func main() {
	var a [10]int
	a[0] = 10
	a[1] = 20
	a[2] = 30
	a[3] = 40
	fmt.Println(a[0], a[1], a[2], a[3])
	fmt.Println(a)

	num := [5]int{1, 2, 3, 4, 5}
	fmt.Println(num)
	for i, value := range num {
		fmt.Println(value, i)
	}

	sliceOne := []int{1, 2, 3, 4, 5}
	fmt.Println(sliceOne[0], sliceOne[1], sliceOne[2])

	sliceTwo := sliceOne[1:3]
	fmt.Println(sliceTwo)

	fmt.Println("numSlice[:3] =", sliceOne[:3])
	fmt.Println("numSlice[1:] =", sliceOne[1:])
	fmt.Println("numSlice[:] =", sliceOne[:])

	sliceThree := make([]int, 5, 10)
	fmt.Println(sliceThree)
	fmt.Println("slice len: ", len(sliceThree))
	fmt.Println("slice cap: ", cap(sliceThree))

	sliceThree = append(sliceThree, 1)
	fmt.Println(sliceThree)
	sliceThree = append(sliceThree, 2, 3, 4)
	fmt.Println(sliceThree)

	m := make(map[string]int)
	m["first"] = 1
	m["first"] = 2
	fmt.Println(m["first"])
	fmt.Println(len(m))

	m["second"] = 2
	fmt.Println(len(m))

	v, ok := m["third"]
	fmt.Println("The value:", v, "Present?", ok)

	delete(m, "first")
	fmt.Println(m)

	var mp = map[string]int{
		"first":  1,
		"second": 2,
	}
	fmt.Println(mp)

	outer := map[string]map[string]int{
		"inner1": map[string]int{
			"first":  1,
			"second": 2,
		},

		"inner2": map[string]int{
			"first":  3,
			"second": 4,
		},
	}
	v1, ok1 := outer["inner1"]
	if ok1 {
		fmt.Println(v1["first"], v1["second"])
	}
}
