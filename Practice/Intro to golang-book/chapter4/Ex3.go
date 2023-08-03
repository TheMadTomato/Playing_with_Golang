package main

import "fmt"

func main() {

	for j := 1; j <= 100; j++ {
		if j%3 == 0 && j%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if j%5 == 0 {
			fmt.Println("Buzz")
		} else if j%3 == 0 {
			fmt.Println("Fizz")
		}
	}

}
