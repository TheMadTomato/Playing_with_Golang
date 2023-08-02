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

func initializeAlphabet() []string {
	//initializing the core list
	alphabet := []string{}
	for i := 'A'; i <= 'Z'; i++ {
		alphabet = append(alphabet, string(i))
	}
	return alphabet
}

func modifyAlphabet(diffFactor int, alphabet []string) []string {
	//Initialize the arrays to store the edited array, Pos for when the diffFactor>0 and Neg is for when diffFactor<0
	var Pos_alphabet []string
	var Neg_alphabet []string
	Pos_alphabet = alphabet
	Neg_alphabet = alphabet

	//appending the arrays
	if diffFactor > 0 {
		//if +diffFactor
		firstLetters := Pos_alphabet[:diffFactor%26]
		remaining := Pos_alphabet[diffFactor%26:]
		Pos_alphabet = append(remaining, firstLetters...)
		return Pos_alphabet
	} else if diffFactor < 0 {
		//if -diffFactor
		lastLetters := Neg_alphabet[len(Neg_alphabet)-(-diffFactor):]
		remaining := Neg_alphabet[:len(Neg_alphabet)-(-diffFactor)]
		Neg_alphabet = append(lastLetters, remaining...)
		return Neg_alphabet
	}
	return alphabet
}

func encryptInput(input string, alphabet []string, diffFactor int) string {
	//Storing the indexes of the input's each char in the original alphabet table
	indexes := checkAlphabet(strings.ToUpper(input), alphabet)

	//Using the indexes value in the appended alphabet table to get the encrypted input
	new_input := make([]string, len(indexes))
	if diffFactor > 0 {
		for i := 0; i < len(indexes); i++ {
			//I used the `%26` to wrap around to the beginning of the alphabet. i.e if diffFactor = 30 it will become 4 after the %26
			//thus no more need for the diffFactor<26 & diffFactor>26 // tho one problem now while positive numbers are wrapping around while negative are making errors
			new_input[i] = alphabet[(indexes[i]+diffFactor)%26]
		}
		return strings.ToLower(strings.Join(new_input, ""))
	} else if diffFactor < 0 {
		for i := 0; i < len(indexes); i++ {
			new_input[i] = alphabet[(indexes[i]-(-diffFactor))%26]
		}
		return strings.ToLower(strings.Join(new_input, ""))
	}
	return input
}

func decryptInput(input string, alphabet []string, diffFactor int) string {
	//Storing the indexes of the input's each char in the modified alphabet table
	indexes := checkAlphabet(strings.ToUpper(input), alphabet)

	//Using the indexes value in the original alphabet table to get the decrypted input
	new_input := make([]string, len(indexes))
	if diffFactor > 0 {
		for i := 0; i < len(indexes); i++ {
			new_input[i] = alphabet[(indexes[i]-diffFactor+26)%26]
		}
		return strings.ToLower(strings.Join(new_input, ""))
	} else if diffFactor < 0 {
		for i := 0; i < len(indexes); i++ {
			new_input[i] = alphabet[(indexes[i]-diffFactor-26)%26]
		}
		return strings.ToLower(strings.Join(new_input, ""))
	}
	return input
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

			//take the input of the differentiating factor
			var diffFactor int
			fmt.Println("Enter the diff factor")
			fmt.Scanln(&diffFactor)

			//Print the Original Alphabet table
			alphabet := initializeAlphabet()
			fmt.Println(alphabet)

			//Modify the alphabet based on the diffFactor
			modifiedAlphabet := modifyAlphabet(diffFactor, alphabet)
			fmt.Println(modifiedAlphabet)

			//Taking the input from user
			var input string
			fmt.Println("Enter input: ")
			fmt.Scanln(&input)
			fmt.Println("\n-------------------------")
			fmt.Println("Original input: ", input)

			//Encrypt the input using the modified alphabet
			encryptedInput := encryptInput(input, modifiedAlphabet, diffFactor)
			fmt.Println("Encrypted input: ", encryptedInput)
			fmt.Println("-------------------------")
		case 2:
			//to decrypt this type of encryption multiple attempts will be needed to guess the right diffFactor
			for {
				//take the input of the differentiating factor
				var diffFactor int
				fmt.Println("Enter the diff factor, \033[33mto exit enter 0\033[0m")
				fmt.Scanln(&diffFactor)

				//break when user is done
				if diffFactor == 0 {
					break
				}

				//Print the Original Alphabet table
				alphabet := initializeAlphabet()
				fmt.Println(alphabet)

				//Modify the alphabet based on the diffFactor
				modifiedAlphabet := modifyAlphabet(diffFactor, alphabet)
				fmt.Println(modifiedAlphabet)

				fmt.Println("Enter encrypted input: ")
				var encryptedInput string
				fmt.Scanln(&encryptedInput)
				fmt.Println("Encrypted input: ", encryptedInput)

				//Decrypt the input using the modified alphabet
				decryptedInput := decryptInput(encryptedInput, modifiedAlphabet, diffFactor%26)
				fmt.Println("Decrypted input: ", decryptedInput)

			}
		case 3:
			fmt.Println("\033[33mExit. Thanks for Using.\033[0m")
		default:
			fmt.Println("\033[31mInvalid Input! Idiotka!\033[0m")
		}

		//beak condition (this syntax of for{//code if{break}} is the only way we can achieve something like do-while in golang)
		if choice == 3 {
			break
		}
	}
}
