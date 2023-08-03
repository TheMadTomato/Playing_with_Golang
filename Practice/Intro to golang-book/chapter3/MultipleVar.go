package main

import "fmt"

func main() {
	fmt.Print("Enter a number: ")
	var input float64       // we did not use the := here for i beleive we wanted to make sure that the float be of 64bit capacity thue float64
	fmt.Scanf("%f", &input) // the exact same functino in C
	output := input * 2
	fmt.Println(output)
}

/*
	Go can also define differet varibale at once:
	var (
		a = 1
		b = 2
		c = 3
	)
*/
