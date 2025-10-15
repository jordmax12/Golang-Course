package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

type Results struct {
	EBT float64
	Tax float64
	Profit float64
	EBTRatio float64
}

func validateUserInput(input float64) bool {
	if math.IsNaN(input) || input <= 0 {
		return false
	}
	return true
}

func storeResults(results Results) error {
	file_name := fmt.Sprintf("results-%s.json", time.Now().Format("2006-01-02_15-04-05"))
	data, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(file_name, data, 0644)
}

func main() {
	fmt.Println("Hello, World!")

	var revenue float64
	var expenses float64
	var taxRate float64

	revenue = getUserInput("Enter the revenue: ", validateUserInput)
	expenses = getUserInput("Enter the expenses: ", validateUserInput)
	taxRate = getUserInput("Enter the tax rate: ", validateUserInput)

	ebt, tax, profit, ebtRatio := calculateFinancials(revenue, expenses, taxRate)

	results := Results{
		EBT:      ebt,
		Tax:      tax,
		Profit:   profit,
		EBTRatio: ebtRatio,
	}

	err := storeResults(results)
	if err != nil {
		fmt.Println("Error storing results:", err)
	}

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

func getUserInput(prompt string, validate func(float64) bool) float64 {
	outputText(prompt)
	var input string
	fmt.Scanln(&input)
	value, _ := strconv.ParseFloat(input, 64)
	if !validate(value) {
		fmt.Println("Invalid input, please enter a valid number")
		return getUserInput(prompt, validate)
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