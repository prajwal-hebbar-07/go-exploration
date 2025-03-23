package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 1000
	expectedReturns := 5.5
	var years float64 = 10

	futureValue := investmentAmount + math.Pow(1+expectedReturns/100, years)
	fmt.Println(futureValue)
}
