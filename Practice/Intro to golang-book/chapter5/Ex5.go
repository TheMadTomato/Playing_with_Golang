package main

import "fmt"

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	var min int
	min = x[0]
	for i := 0; i <= len(x)-1; i++ {
		if min >= x[i] {
			min = x[i]
		} else {
			continue
		}
	}
	fmt.Println("The smalles number in the array x is: ", min)
}
