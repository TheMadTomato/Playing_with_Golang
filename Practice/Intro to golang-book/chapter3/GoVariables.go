package main

import "fmt"

func main() {
	//intrstingly enough in go we first creat the variable using the var keyword, then the variable name and lastly variable type.
	var x string = "Hello, Tomato"
	fmt.Println(x)

	//another way of writing it
	var y string
	y = "Tomato again?"
	fmt.Println(y)

	//another intrsting feature, ":=". when we use it we do not need to declare var and type because the compiler is smart enough to infer the type based on the literal value we assign to the variable. pretty Impressive for a rodent :v
	z := "still here?"
	fmt.Println(z)
}

/*generarly we should always used the shorter form whenver possible.
Local and global scopes concepts in Go are the same in C/C++ so nothing much to say about them.
we can make function in Go by saying "func <fucntionName>(){...}.
Constant variable are the same in other languages as well but to properly declare them we should first declare the const keyword, the variable name and lastly its type.
*/
