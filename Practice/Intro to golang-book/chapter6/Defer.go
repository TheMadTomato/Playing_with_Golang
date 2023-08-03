package main

import "fmt"

// Go has a special statement called defer that schedules a function call to be run after the function completes.
func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}
func main() {
	defer second()
	first()
}

/*Defer isd often used with opening and closing folders

Defer gives three main advantages:
1- contribute to a cleaner more readble code
2- if the code contain multiple retrun statements defer will happen befor any of them
3- defer fuinction will run even if a runtime panic occurs*/
