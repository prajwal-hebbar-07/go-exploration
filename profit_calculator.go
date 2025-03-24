package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Enter the Revenue Generated: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter the Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter the Tax Rate: ")
	fmt.Scan(&taxRate)

	var earningsBeforeTax float64 = revenue - expenses
	var earningsAfterTax float64 = earningsBeforeTax * (1 - taxRate)
	var ratio float64 = earningsBeforeTax / earningsAfterTax

	fmt.Println("Earnings Before Tax: ", earningsBeforeTax)
	fmt.Println("Earnings After Tax: ", earningsAfterTax)
	fmt.Println("Ratio: ", ratio)
}

