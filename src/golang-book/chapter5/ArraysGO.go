package main

import "fmt"

func main() {
	// var x [5]float64 // so that's how we declare arrays in go, nothing too special
	// x[0] = 98
	// x[1] = 93
	// x[2] = 77
	// x[3] = 82
	// x[4] = 83
	x := [5]float64{98, 93, 77, 82, 83} // a shorter way
	// var total float64 = 0
	// for i := 0; i < len(x); i++ {
	// 	total += x[i]
	// }
	// fmt.Println(total / float64(len(x))) // that's how we can convert types in order to do error free compilation
	var total float64 = 0
	for _, value := range x { //another shorter way, and because Go compiler doesnt let us keep unused variable in our code we use the underscore
		total += value
	}
	fmt.Println(total / float64(len(x)))
}
