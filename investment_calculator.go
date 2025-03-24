package main

import (
	"fmt"
	"math"
)

func main() {
	// investmentAmount, years, expectedReturns := 1000.0, 10.0, 5.5
	const inflationRate float64 = 2.5
	var investmentAmount float64 = 1000
	var years float64 = 10
	var expectedReturns float64 = 5.5

	var futureValue float64 = investmentAmount + math.Pow(1+expectedReturns/100, years)
	var futureRealValue float64 = futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
