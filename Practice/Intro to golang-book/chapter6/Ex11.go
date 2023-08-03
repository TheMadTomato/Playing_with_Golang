/*11. Write a program that can swap two integers (x := 1; y := 2; swap(&x, &y)
should give you x=2 and y=1).*/

package main

import "fmt"

// if int64 were used here uncommnet the snippet of code where i delcare x an y in main as int64
func swap(x *int, y *int) {
	temp := *x

	*x = *y
	*y = temp

	fmt.Println("Swapacadabra!! x: ", *x, " & y: ", *y)
}

func main() {
	x, y := 1, 2
	// var x int64
	// var y int64
	// x = 1
	// y = 2

	fmt.Println("Before the Swap: x: ", x, " & y: ", y)
	swap(&x, &y)
}

/*when i used the *int64 the porgram did not work gavme me this error:
# command-line-arguments
.\Ex11.go:21:7: cannot use &x (value of type *int) as *int64 value in argument to swap
.\Ex11.go:21:11: cannot use &y (value of type *int) as *int64 value in argument to swap
ig it is because i did not specifcally declare x and y in main as int64 type
ye after proper testing turned out i am right.*/
