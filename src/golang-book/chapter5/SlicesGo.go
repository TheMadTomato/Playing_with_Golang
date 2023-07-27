package main

import "fmt"

func main() {
	//append block
	// slice1 := []int{1, 2, 3}
	// slice2 := append(slice1, 4, 5) // to add elemetns to a slice, slice can be increase unlike arrays
	// fmt.Println(slice1, slice2)

	//copy block
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2) // we create a slice with the make function we can also creat slice using [low : high] expression
	copy(slice2, slice1)
	fmt.Println(slice1, slice2) // slice2 only contains 2 slots so it only copy  the first 2 elemnts in slice1. unlike append it cant change the size  of the underlying array nor creat a new array and coppy all elements to it if there was no eonugh space
}

// Mostly understandable nothing to obscure it is a clear concepts, very close to lists in python. Nevertheless i have to work with them a little to a get a storng hold
