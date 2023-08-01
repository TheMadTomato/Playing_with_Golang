package main

import (
	"fmt"
	"strings"
)

func main() {
	//intitializing the core list
	var alphabet []string
	for i := 'A'; i <= 'Z'; i++ {
		alphabet = append(alphabet, string(i))
	}
	//take the input of the diffirating factor
	var diffFactor int
	fmt.Println("Enter the diff factor")
	fmt.Scanf("%d", &diffFactor)

	//appending the changes
	var Pos_alphabet []string
	var Neg_alphabet []string
	Pos_alphabet = alphabet
	Neg_alphabet = alphabet
	//initializing the appended lists
	if diffFactor > 0 && diffFactor < 25 {
		//if +diffFactor
		firstLetters := Pos_alphabet[:diffFactor]
		remaining := Pos_alphabet[diffFactor:]
		Pos_alphabet = append(remaining, firstLetters...)
		fmt.Println(Pos_alphabet)
	} else if diffFactor < 0 && diffFactor > -25 {
		//if -diffFactor
		lastLetters := Neg_alphabet[len(Neg_alphabet)-(-diffFactor):]
		remaining := Neg_alphabet[:len(Neg_alphabet)-(-diffFactor)]
		Neg_alphabet = append(lastLetters, remaining...)
		fmt.Println(Neg_alphabet)
	}
	fmt.Println(alphabet)

	var input string
	fmt.Println("Enter input: ")
	fmt.Scanln(&input)
	fmt.Println("Original input: ", input)

	var indexes []int
	for i, char := range alphabet {
		if strings.ContainsRune(input, rune(char)) {
			indexes = append(indexes, i)
		}
	}

	fmt.Println("Indexes:", indexes)
}
