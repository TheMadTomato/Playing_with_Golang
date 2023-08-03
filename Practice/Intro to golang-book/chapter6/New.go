package main

import "fmt"

func one(xPtr *int) {
	*xPtr = 1
}
func main() {
	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr) // x is 1
}

// not sure i understand it fully yet, tho we can say it another way to get a pointer. and it has many advantages in go bcs go is a garbage-collected languagge
