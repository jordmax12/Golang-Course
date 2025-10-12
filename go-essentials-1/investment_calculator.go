package main

import (
	"fmt"
	"math"
)

func main() {
	// varinvestmentAmount, years float64 = 1000.0, 10.0
	// expectedReturnRate := 5.5

	// investmentAmount, years := 1000.0, 10.0
	// expectedReturnRate := 5.5

	investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	fmt.Printf("Future Value: %.2f\n", futureValue)
}
