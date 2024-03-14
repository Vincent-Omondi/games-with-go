cd package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	low, high := 1, 100
	attempt := 0

	fmt.Println("Think of a number between", low, "and", high)
	fmt.Println("If you are ready press ENTER to continue")
	scanner.Scan()

	userNumber := getUserNumber(scanner)

	for {

		// bunary search

		guess := (low + high) / 2
		fmt.Println("My guess is:", guess)
		fmt.Println("Is the guess:")

		fmt.Println("(a) Too high?")
		fmt.Println("(b) Too low?")
		fmt.Println("(c) correct?")
		scanner.Scan()

		attempt++

		responce := scanner.Text()

		if responce == "a" {
			high = guess - 1
		} else if responce == "b" {
			low = guess + 1
		} else if responce == "c" {
			if guess != userNumber {
				fmt.Println(" You said the number is", userNumber, "but now you say it is", guess)
				fmt.Println("You lied. Game over")
			} else {
				fmt.Println("i Won in", attempt, " trials. Wooah!!")
			}
			return

		} else {
			fmt.Println("Invalid Input. Try again")
		}

	}
}

func getUserNumber(scanner *bufio.Scanner) int {
	fmt.Println("Please enter your choosen number")
	scanner.Scan()
	var num int
	fmt.Sscan(scanner.Text(), &num)
	return num
}
