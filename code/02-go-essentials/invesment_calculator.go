package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Starting calculation")
	fmt.Println()

	var investmentAmount float64 = 1000
	years := 10.0 // `:=` infer type
	// var years float64 // initialize with zero value
	// var years float64 = 10.0
	// var years = 10.0 // infer type
	expectedReturnRate := 5.5

	// var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	fmt.Println("Future Value: ", futureValue)
	fmt.Println("\nStarting calculation")

}
