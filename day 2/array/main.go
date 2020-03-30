package main

import "fmt"

func main() {
	var a []int = []int{32, 29, 78, 16, 81}
	var b = [5]int{1, 2, 3, 4, 5}
	c := [5]int{5, 6, 7, 89, 10}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	d := [...]int{11, 12, 13, 14, 15, 16, 17, 189, 19}

	fmt.Println(len(d))
	fmt.Println(d)

	for _, value := range d {
		fmt.Println(value)
	}

	e := d

	for _, value := range e {
		fmt.Println(value)
	}
}
