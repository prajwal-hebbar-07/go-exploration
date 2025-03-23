package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount = 1000
	var expectedReturnValue = 5.5
	var years = 10

	var futureReturns = float64(investmentAmount) * math.Pow(1 + expectedReturnValue / 100, float64(years))
	fmt.Println(futureReturns)
}