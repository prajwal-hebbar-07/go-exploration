package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount, years float64 = 1000, 10
	expectedReturns := 5.5

	futureValue := investmentAmount + math.Pow(1+expectedReturns/100, years)
	fmt.Println(futureValue)
}
