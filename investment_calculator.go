package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64
	var years float64
	var expectedReturn float64
	const inflationRate float64 = 8

	fmt.Print("Enter the Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter the number of Years: ")
	fmt.Scan(&years)

	fmt.Print("Enter the Expected Return Rate: ")
	fmt.Scan(&expectedReturn)

	var futureValue float64 = investmentAmount * math.Pow(1+expectedReturn/100, years)
	var futureRealValue float64 = futureValue / math.Pow(1+inflationRate/100, years)

	// fmt.Println("Future Value: ", futureValue)
	// fmt.Println("Future Real Value adjusted for Inflation: ", futureRealValue)
	fmt.Printf("Future Value: %v\nFuture Value adjusted for Inflation: %v\n", futureValue, futureRealValue)
}
