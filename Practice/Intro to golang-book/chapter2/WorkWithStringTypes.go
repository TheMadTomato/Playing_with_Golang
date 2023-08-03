package main

import "fmt"

func main() { //got to say this syntax reminds me a lot of python :v
	fmt.Println("The Mad Tomato")                // the control
	fmt.Println(len("The Mad Tomato"))           //the typical len function that counts all characters in a str
	fmt.Println("The Mad Tomato"[1])             // indexing. tho intrstingly here we see 104 instead of 'h', the byte representation this is because characters are represented by bytes
	fmt.Println("The Mad Tomato " + " Using GO") // typical concatenation process
	fmt.Println(`The \nMad \nTomato`)            // we see that \n gets printed when between ``
	fmt.Println("The \nMad \nTomato")            // we see newlines printed when between ""
}

/* Unlike C/C++, in Go, strings can be created using double quotes "" or back ticks `` .
but there is a difference. backticks quoted strings cannot use newline characters such as \n and \t
*/
