// in Golang functions are also knowsn as procedures or suibroutines. intrsting :D

package main

import "fmt"

// to declare functions we first type func, then write the name of the function. after that insidde the parantheses we set the parameters or inputs.
// laslty after closign the parantheseswe put the return type. quick note: the parameters and the return type are known as the fucntion's signature.
func average(xs []float64) float64 {
	// her is the functions body.
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs)) // multiple value can be returned
}

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))
}

// func f() (int, int) {
// 	return 5, 6
// 	}
// 	func main() {
// 	x, y := f()
// 	}
