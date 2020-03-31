package main

import "fmt"

func main() {
	var a []int // int형 슬라이스 선언

	var b []int = make([]int, 5)
	var c = make([]int, 5)
	d := make([]int, 5)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	e := make([]int, 5, 10)
	fmt.Println(len(e))
	e = append(a, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1)
	fmt.Println(len(e))
	fmt.Println(e)

	a = append(a, e...)
	fmt.Println(a)
}
