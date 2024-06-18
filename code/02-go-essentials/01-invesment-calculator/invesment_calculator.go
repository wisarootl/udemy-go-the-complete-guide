package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Starting calculation")
	fmt.Println()

	const inflationRate = 6.5
	var investmentAmount float64
	var years float64
	// years := 10.0 // `:=` infer type
	// var years float64 // initialize with zero value
	// var years float64 = 10.0
	// var years = 10.0 // infer type
	expectedReturnRate := 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount) // user input, `&` is pointer

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	// var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println("Future Value: ", futureValue)
	fmt.Println("futureRealValue: ", futureRealValue)
	fmt.Println("\nComplete calculation")

}
