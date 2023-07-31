package main

import "fmt"

// func zero(x int) { // the functions shouild change4 the value in main but it did not for suhc task we use pointers
// 	x = 0
// 	}
// 	func main() {
// 	x := 5
// 	zero(x)
// 	fmt.Println(x) // x is still 5
// 	}

func zero(xPtr *int) {
	// a pointer refrence a location in memory where a value is stored rather than the value itself, thus by using a pointer, s function csn modify the original variable.
	*xPtr = 0
}

// in go pointers are rperesented using an * followed by the type.
// didin't really ge a strong grasp pf the derfencing a pointer yet( maybe ehrn i start some projects) but for now i believe it is when access the varible sotred in the pointer memory address
// we use the & operator to to find the address of a variable. i.e &x in main and xPtr in 0 refer to the same memory location
func main() {
	x := 5
	fmt.Println("X before the function zero is, ", x)
	zero(&x)
	fmt.Println("X after the function zero is, ", x) // x is 0
}
