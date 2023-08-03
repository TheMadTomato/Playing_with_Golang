/*
Write a function that takes an integer and halves it and returns true if it was even
or false if it was odd. For example, half(1) should return (0, false) and
half(2) should return (1, true).
*/
package main

import "fmt"

func half(num int64) {
	if num%2 == 0 {
		fmt.Println("(", num/2, ", true)")
	} else {
		fmt.Println("(", num/2, ", false)")
	}
}

func main() {
	var num int64
	fmt.Println("Enter a number: ")
	fmt.Scanf("%d", &num)

	half(num)
}

/*INtrstingly enough the % operator did not work with the float64 type it would give this errror (# command-line-arguments
.\Ex2.go:11:5: invalid operation: operator % not defined on num (variable of type float64)),
 but after i changed it to int64 it worked fine. i am offlin nw will search it up when i connect back to wifi*/
