package main

import "fmt"

func main() { //intrstingly enough if the brackets are not used in this style particularly the porgram does not work .(O.o).
	//will have to investigate for sure
	fmt.Println("1 + 1 = ", 1+1)     // it is kinda like C syntax but less scary :p
	fmt.Println("1 + 1 = ", 1.0+1.0) // same result no change -.-
}

/* So i did some investigation about the prior brackets incidents and apprently it is illigal
to use any other bracket format style because of somthing called the "semicolon injection rule"
or "automatic semicolon insertion". To my understanding this rule just helps go developers
by providing overall consistency accrose different code bases. intresting indeed .(0.0). */
