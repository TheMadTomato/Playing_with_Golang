package main

import (
	"fmt"
	"strings"

	"github.com/inancgumus/screen"
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

	//main loop
	for {
		//Welcome msg + menu
		fmt.Println("\033[35m****Caesar Cipher****\033[0m")
		fmt.Println("\033[34m1. Encrypt\033[0m\n\033[36m2. Decrypt\033[0m\n\033[31m3. Exit\033[0m\n")

		//user's choice
		var choice int
		fmt.Scanln(&choice)

		switch choice {

		case 1:
			//Haven't figured out a proper fix for this yet
			fmt.Println("\033[31mWarning! Do not include spaces in your input.\033[0m\n")
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
			fmt.Println("\033[31mNot finished yet, What a lazy tomato!\033[0m")
		case 3:
			fmt.Println("\033[33mExit. Thanks for Using.\033[0m")
		default:
			fmt.Println("\033[31mInvalid Input! Idiotka!\033[0m")
		}

		//only equivalent i found to system("cls")
		screen.Clear()

		//beak condition (this syntax of for{ if{break}} is the way we can achieve something like do-while in golang)
		if choice == 3 {
			break
		}
	}
}
