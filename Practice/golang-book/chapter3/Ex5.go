package main

import "fmt"

func main() {
	//welcom message
	fmt.Print("Welcom to the Fehrenheit to Celsius Converter\n.")
	//take input
	fmt.Print("Enter a value in Fahrenheit:\n ")
	var Fehrenheit float64
	fmt.Scanf("%f", &Fehrenheit)
	//calculate
	Celsius := ((Fehrenheit - 32) * 5 / 9)
	//output
	fmt.Println("Celsius: ", Celsius)
}
