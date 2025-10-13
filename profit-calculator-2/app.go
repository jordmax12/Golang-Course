package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello, World!")

	var revenue float64
	var expenses float64
	var taxRate float64

	revenue = getUserInput("Enter the revenue: ")
	expenses = getUserInput("Enter the expenses: ")
	taxRate = getUserInput("Enter the tax rate: ")

	ebt, tax, profit, ebtRatio := calculateFinancials(revenue, expenses, taxRate)

	finalEbtString := fmt.Sprintf("EBT: %.2f", ebt)
	finalTaxString := fmt.Sprintf("Tax: %.2f", tax)
	finalProfitString := fmt.Sprintf("Net Profit: %.2f", profit)
	finalEbtRatioString := fmt.Sprintf("EBT Ratio: %.2f", ebtRatio)

	outputText(finalEbtString)
	outputText(finalTaxString)
	outputText(finalProfitString)
	outputText(finalEbtRatioString)
}

func outputText (text string) {
	fmt.Println(text)
}

func getUserInput(prompt string) float64 {
	outputText(prompt)
	var input string
	fmt.Scanln(&input)
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Invalid input, please enter a valid number")
		return getUserInput(prompt) // Recursively ask again
	}
	return value
}

func calculateFinancials(revenue float64, expenses float64, taxRate float64) (ebt float64, tax float64, profit float64, ebtRatio float64) {
	ebt = revenue - expenses
	tax = ebt * (taxRate / 100)
	profit = ebt - tax
	ebtRatio = ebt / profit
	return ebt, tax, profit, ebtRatio
}