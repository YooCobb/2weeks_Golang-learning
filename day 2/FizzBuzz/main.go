package main

// 1에서 100까지 출력
// 3의 배수는 Fizz출력
// 5의 배수는 Buzz출력
// 3과 5의 공배수는 FizzBuzz출력

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		}

	}
}
