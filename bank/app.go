package main

import (
	"fmt"
	"os"
	"strconv"
)

var accountBalance = 1000.0

func main() {
	outputText("Welcome to Globo Bank")
	outputText(fmt.Sprintf("Your balance is: %f", accountBalance))

	for {
		mainMenu()
	}
}

func checkBalance() {
	outputText(fmt.Sprintf("\nYour balance is: %f\n", accountBalance))
	outputText("--------------------------------")
}

func deposit() {
	checkBalance()
	amount := getUserInput("Enter the amount you want to deposit: ")
	accountBalance += amount
	outputText(fmt.Sprintf("\nYour new balance is: %f\n", accountBalance))
}

func withdraw() {
	checkBalance()
	amount := getUserInput("Enter the amount you want to withdraw: ")
	accountBalance -= amount
	outputText(fmt.Sprintf("\nYour new balance is: %f\n", accountBalance))
}

func exit() {
	outputText("Thank you for using Globo Bank")
	os.Exit(0)
}

func mainMenu() {
	outputText("\n\nWhat would you like to do?")
	outputText("1. Check Balance")
	outputText("2. Deposit")
	outputText("3. Withdraw")
	outputText("4. Exit")

	choice := getUserInput("Enter your choice: ")

	outputText(fmt.Sprintf("You chose: %f", choice))

	switch choice {
	case 1:
		checkBalance()
	case 2:
		deposit()
	case 3:
		withdraw()
	case 4:
		exit()
	default:
		outputText("Invalid choice. Please try again.")
	}
}

func getUserInput(prompt string) float64 {
	outputText(prompt)
	var input string
	fmt.Scanln(&input)
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		outputText("Invalid input. Please try again.")
		return getUserInput(prompt)
	}
	return value
}

func outputText(text string) {
	fmt.Println(text)
}
