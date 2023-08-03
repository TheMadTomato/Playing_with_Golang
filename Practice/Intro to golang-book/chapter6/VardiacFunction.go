package main

import "fmt"

// when we use elipsis "..." we make a variadic function where add can support multutiple integers in its parameter
// it can also support slices
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}
func main() {
	fmt.Println(add(1, 2, 3))
}
