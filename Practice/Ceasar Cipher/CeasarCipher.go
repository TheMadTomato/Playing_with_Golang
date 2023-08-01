package main

import (
	"fmt"
	"strings"
)

func checkAlphabet(input string, alphabet []string) []int {
	//declare where the index will be saved
	indexes := []int{}
	//for each char in the input
	for _, char := range input {
		//check if said char is in the alphabet array
		for i, letter := range alphabet {
			if strings.ToUpper(string(char)) == letter {
				indexes = append(indexes, i)
				break
			}
		}
	}
	return indexes
}

func main() {

	//Welcome msg + menu
	fmt.Println("\t****Caesar Cipher****")
	fmt.Println("1. Encrypt\n2. Decrypt\n3. Exit")

	//user's choice
	var choice int
	fmt.Scanln(&choice)

	switch choice {

	case 1:
		//initializing the core list
		var alphabet []string
		for i := 'A'; i <= 'Z'; i++ {
			alphabet = append(alphabet, string(i))
		}

		//take the input of the differentiating factor
		var diffFactor int
		fmt.Println("Enter the diff factor")
		fmt.Scanln(&diffFactor)

		//Print the Original Alphabet table
		fmt.Println(alphabet)

		//Initialize the arrays to store the edited array, Pos for when the diffFactor>0 and Neg is for when diffFactor<0
		var Pos_alphabet []string
		var Neg_alphabet []string
		Pos_alphabet = alphabet
		Neg_alphabet = alphabet

		//appending the arrays
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

		//Taking the input from user
		var input string
		fmt.Println("Enter input: ")
		fmt.Scanln(&input)
		fmt.Println("Original input: ", input)

		//Storing the indexes of the input's each char in the original alphabet table
		indexes := checkAlphabet(strings.ToUpper(input), alphabet)
		fmt.Println("Indexes:", indexes)

		//Using the indexes value in the appended alphabet table to get the encrypted input
		new_input := make([]string, len(indexes))
		if diffFactor > 0 && diffFactor < 25 {
			for i := 0; i < len(indexes); i++ {
				new_input[i] = Pos_alphabet[indexes[i]]
			}
			fmt.Println("Encrypted input: ", strings.ToLower(strings.Join(new_input, "")))
		} else if diffFactor < 0 && diffFactor > -25 {
			for i := 0; i < len(indexes); i++ {
				new_input[i] = Neg_alphabet[indexes[i]]
			}
			fmt.Println("Encrypted input: ", strings.ToLower(strings.Join(new_input, "")))
		}
	case 2:
		fmt.Println("Not finished yet, What a lazy tomato!")
	case 3:
		fmt.Println("Exit. Thanks for Using.")
	default:
		fmt.Println("Invalid Input! Idiotaka!")
	}
}
