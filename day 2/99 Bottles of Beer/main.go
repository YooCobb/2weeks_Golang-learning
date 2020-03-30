package main

import "fmt"

func main() {
	s1 := "bottles of beer"
	s1_1 := "bottle of beer"
	s2 := " on the wall, "

	for i := 99; i > 0; i-- {
		if i != 1 {
			fmt.Printf("%d %s%s%d %s.\n", i, s1, s2, i, s1)
		} else {
			fmt.Printf("%d %s%s%d %s.\n", i, s1_1, s2, i, s1_1)
		}
	}
}
