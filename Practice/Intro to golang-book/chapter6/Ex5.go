/*
The Fibonacci sequence is defined as: fib(0) = 0, fib(1) = 1, fib(n) =
fib(n-1) + fib(n-2). Write a recursive function that can find fib(n).
*/
package main

import "fmt"

func Fibonacci(n int64) int64 {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return Fibonacci(n-1) + Fibonacci(n-2)
	}

}

func main() {
	var num int64
	fmt.Println("Enter number: ")
	fmt.Scanf("%d", &num)

	fmt.Println(Fibonacci(num))
}
