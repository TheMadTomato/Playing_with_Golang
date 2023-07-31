/*4. Using makeEvenGenerator as an example, write a makeOddGenerator function
that generates odd numbers*/

package main

import "fmt"

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func main() {
	nextOdd := makeOddGenerator()
	// for the sake of automation
	var amount int64
	fmt.Println("Enter the amount to be generated: ")
	fmt.Scanf("%d", &amount)

	for i := 0; i <= int(amount); i++ {
		fmt.Println(i+1, "- ", nextOdd())
	}
}
