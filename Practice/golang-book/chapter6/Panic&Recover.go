package main

import "fmt"

// The panic and recover function in go are similar to the cin.clear cin.ignore combination in C++
func main() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}
