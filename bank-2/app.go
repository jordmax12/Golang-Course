package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Account struct {
	Owner    string  `json:"owner"`
	Balance  float64 `json:"balance"`
	Currency string  `json:"currency"`
}

const accountFile = "account.json"

var account Account

func saveAccount(acc Account) error {
	outputText(fmt.Sprintf("\nSaving account for %s", acc.Owner))
	data, err := json.MarshalIndent(acc, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(accountFile, data, 0644) // 0644 sets read/write permissions for owner, read-only for others
}

func loadAccount() (Account, error) {
	outputText("\nLoading account from file")
	data, err := os.ReadFile(accountFile)
	if err != nil {
		if os.IsNotExist(err) {
			return Account{
				Owner:    "Globo Customer",
				Balance:  1000.0,
				Currency: "USD",
			}, nil
		}
		return Account{}, err
	}

	var acc Account
	if err := json.Unmarshal(data, &acc); err != nil {
		return Account{}, err
	}
	return acc, nil
}

func main() {
	var err error
	account, err = loadAccount()
	if err != nil {
		log.Fatal(err)
	}

	outputText("Welcome to Globo Bank")
	outputText(fmt.Sprintf("Account owner: %s", account.Owner))
	outputText(fmt.Sprintf("Your balance is: %f %s", account.Balance, account.Currency))

	for {
		mainMenu()
	}
}

func checkBalance() {
	outputText(fmt.Sprintf("\nYour balance is: %f %s\n", account.Balance, account.Currency))
	outputText("--------------------------------")
}

func deposit() {
	checkBalance()
	amount := getUserInput("Enter the amount you want to deposit: ")
	account.Balance += amount
	if err := saveAccount(account); err != nil {
		log.Fatal(err)
	}
	outputText(fmt.Sprintf("\nYour new balance is: %f %s\n", account.Balance, account.Currency))
}

func withdraw() {
	checkBalance()
	amount := getUserInput("Enter the amount you want to withdraw: ")
	account.Balance -= amount
	if err := saveAccount(account); err != nil {
		log.Fatal(err)
	}
	outputText(fmt.Sprintf("\nYour new balance is: %f %s\n", account.Balance, account.Currency))
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
