package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Enter the revenue: ")
	fmt.Scanln(&revenue)

	fmt.Print("Enter the expenses: ")
	fmt.Scanln(&expenses)

	fmt.Print("Enter the tax rate: ")
	fmt.Scanln(&taxRate)

	ebt := revenue - expenses
	tax := ebt * taxRate
	profit := ebt * (1 - taxRate/100)
	ebtRatio := ebt / profit

	fmt.Printf("EBT: %d\n", ebt)	
	fmt.Printf("Tax: %f\n", tax)
	fmt.Printf("Net Profit: %f\n", profit)
	fmt.Printf("EBT Ratio: %f\n", ebtRatio)
}