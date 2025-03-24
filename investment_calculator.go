package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5
	var investAmount float64
	var years float64
	var expectedReturn float64

	fmt.Print("Enter the Investment Amount: ")
	fmt.Scan(&investAmount)

	fmt.Print("Enter the number of Years: ")
	fmt.Scan(&years)

	fmt.Print("Enter the Expected Return Rate: ")
	fmt.Scan(&expectedReturn)

	var futureValue float64 = investAmount * math.Pow(1+expectedReturn/100, years)
	var futureRealValue float64 = futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println("Future Value: ", futureValue)
	fmt.Println("Future Value including the Inflation: ", futureRealValue)
}
