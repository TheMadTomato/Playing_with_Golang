/* Write a function with one variadic parameter that finds the greatest number in a
list of numbers.*/

package main

import "fmt"

func FindTheGreatest(args ...int) int {
	max := 0
	for _, v := range args {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	numList := []int{28, 57, 73, 61, 42, 19, 83, 94, 37, 52, 101}
	fmt.Println("The greatest number in numList is: ", FindTheGreatest(numList...))
}
