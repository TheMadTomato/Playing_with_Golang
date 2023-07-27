package main

import "fmt"

func main() {
	i := 1
	for i <= 10 { // kinda different from other languages but still nonetheless understandable
		// intresting go unlike other programs only has one type of loop "for" which can be used in a variety of ways
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
		i = i + 1
	}
	// a more familiar syntax
	for j := 10; j != 0; j-- {
		if j%2 == 0 { // as we can see the if-else conditionals statement are the saem in C/C++
			fmt.Println(j, "even")
		} else {
			fmt.Println(j, "odd")
		}

	}

}

/*The switch-case in go is also similar to C/C++ but we do not need to add break statemetns after each case apprently*/
