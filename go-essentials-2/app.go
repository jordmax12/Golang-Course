package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount, expectedReturnRate, years float64

	fmt.Print("Enter the investment amount: ")
	_, err := fmt.Scanln(&investmentAmount)
	if err != nil {
		fmt.Print("Error reading input:", err)
		return
	}

	fmt.Print("Enter the number of years: ")
	_, err = fmt.Scanln(&years)
	if err != nil {
		fmt.Print("Error reading input:", err)
		return
	}

	fmt.Print("Enter the expected return rate: ")
	_, err = fmt.Scanln(&expectedReturnRate)
	if err != nil {
		fmt.Print("Error reading input:", err)
		return
	}

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Printf("Future Value: %.2f\n", futureValue)
	fmt.Printf("Future Real Value: %.2f\n", futureRealValue)
}
