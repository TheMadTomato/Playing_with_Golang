package main

import "fmt"

func main() {
	//welcom message
	fmt.Print("Welcom to the Feet to Meters Converter\n.")
	//take input
	fmt.Print("Enter a value in Feet:\n ")
	var Feet float64
	fmt.Scanf("%f", &Feet)
	//calculate
	Meters := Feet * 0.3048
	//output
	fmt.Println("Meters: ", Meters)
}
