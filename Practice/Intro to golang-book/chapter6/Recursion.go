package main

// recursion concepts is all th same accros all lnaguages, nothing specila to see here
import "fmt"

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {
	fmt.Println(factorial(5))
}
